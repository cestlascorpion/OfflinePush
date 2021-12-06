package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	log "github.com/sirupsen/logrus"
)

type GeTuiAuth struct {
	ApiUrl       string
	AppKey       string
	MasterSecret string
	client       *core.RestyClient
}

func NewGeTuiAgent(baseUrl, appId, appKey, masterSecret string, timeout time.Duration) (*GeTuiAuth, error) {
	client, err := core.NewRestyClient(timeout)
	if err != nil {
		return nil, err
	}
	return &GeTuiAuth{
		ApiUrl:       fmt.Sprintf(baseUrl, appId),
		AppKey:       appKey,
		MasterSecret: masterSecret,
		client:       client,
	}, nil
}

func NewGeTuiAgentWithClient(baseUrl, appId, appKey, masterSecret string, hc *http.Client, timeout time.Duration) (*GeTuiAuth, error) {
	client, err := core.NewRestyClientWithClient(hc, timeout)
	if err != nil {
		return nil, err
	}
	return &GeTuiAuth{
		ApiUrl:       fmt.Sprintf(baseUrl, appId),
		AppKey:       appKey,
		MasterSecret: masterSecret,
		client:       client,
	}, nil
}

func (g *GeTuiAuth) GetAuth() (*core.AuthToken, error) {
	sign, tp := core.Signature(g.AppKey, g.MasterSecret)
	result, err := g.client.POST(g.url("/auth"), "", authReq{
		Sign:      sign,
		Timestamp: tp,
		AppKey:    g.AppKey,
	})
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
	return g.ApiUrl + path
}
