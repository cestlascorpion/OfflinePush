package auth

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	log "github.com/sirupsen/logrus"
)

type GeTuiAuth struct {
	apiUrl       string
	appKey       string
	masterSecret string
	client       *core.RestyClient
	cache        *core.AuthToken
	cancel       context.CancelFunc
	mutex        sync.RWMutex
}

func NewGeTuiAgent(baseUrl, appId, appKey, masterSecret string, hc *http.Client) (*GeTuiAuth, error) {
	client, err := core.NewRestyClient(hc)
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf(baseUrl, appId)
	token, err := getAuth(apiUrl, appKey, masterSecret, time.Second*5)
	if err != nil {
		return nil, err
	}

	auth := &GeTuiAuth{
		apiUrl:       apiUrl,
		appKey:       appKey,
		masterSecret: masterSecret,
		client:       client,
		cache: &core.AuthToken{
			Token:    token.Token,
			ExpireAt: token.ExpireAt,
		},
		cancel: nil,
	}

	auth.cancel = auth.watch()
	return auth, nil
}

func (g *GeTuiAuth) GetAuth() (*core.AuthToken, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	auth := &core.AuthToken{
		Token:    g.cache.Token,
		ExpireAt: g.cache.ExpireAt,
	}
	return auth, nil
}

func (g *GeTuiAuth) DelAuth(token string) error {
	result, err := g.client.DELETE(g.url(fmt.Sprintf("/auth/%s", token)), "", nil)
	if err != nil {
		log.Errorf("delete failed err %+v", err)
		return err
	}
	var resp authResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("del auth ok token %s", token)
	return nil
}

type authReq struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	AppKey    string `json:"appkey"`
}

type authResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		ExpireTime string `json:"expire_time"`
		Token      string `json:"token"`
	} `json:"data"`
}

func (g *GeTuiAuth) url(path string) string {
	return g.apiUrl + path
}

func (g *GeTuiAuth) watch() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.TODO())

	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				token, err := getAuth(g.apiUrl, g.appKey, g.masterSecret, time.Second*5)
				if err != nil {
					log.Warnf("watch get auth err %+v", err)
					continue
				}

				g.mutex.Lock()
				g.cache.Token = token.Token
				g.cache.ExpireAt = token.ExpireAt
				g.mutex.Unlock()
			}
		}
	}()

	return cancel
}

func signature(appKey string, masterSecret string) (signature string, timestamp string) {
	timestamp = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	original := appKey + timestamp + masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}

func getAuth(apiUrl, appKey, masterSecret string, timeout time.Duration) (*core.AuthToken, error) {
	sign, tp := signature(appKey, masterSecret)
	result, err := core.POST(apiUrl+"/auth", "", authReq{
		Sign:      sign,
		Timestamp: tp,
		AppKey:    appKey,
	}, timeout)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp authResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	if resp.Data == nil {
		log.Errorf("resp.Data is nil")
		return nil, errors.New("resp.Data is nil")
	}
	token := resp.Data.Token
	expireAtMs, err := strconv.ParseInt(resp.Data.ExpireTime, 10, 64)
	if err != nil {
		log.Errorf(" parseInt failed expire time %s ms", resp.Data.ExpireTime)
		return nil, fmt.Errorf("parseInt failed expire time %s ms", resp.Data.ExpireTime)
	}
	log.Debugf("get auth ok token %s expireAtSec %d", token, expireAtMs/1000)
	return &core.AuthToken{
		Token:    token,
		ExpireAt: expireAtMs / 1000,
	}, nil
}
