package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	. "github.com/cestlascorpion/push/core"
	log "github.com/sirupsen/logrus"
)

type GeTuiAuth struct {
	ApiUrl       string
	AppKey       string
	MasterSecret string
	Timeout      time.Duration
}

func NewGeTuiAgent(baseUrl, appId, appKey, masterSecret string, timeout time.Duration) (*GeTuiAuth, error) {
	return &GeTuiAuth{
		ApiUrl:       fmt.Sprintf(baseUrl, appId),
		AppKey:       appKey,
		MasterSecret: masterSecret,
		Timeout:      timeout,
	}, nil
}

func (g *GeTuiAuth) GetAuth() (string, int64, error) {
	sign, tp := Signature(g.AppKey, g.MasterSecret)
	result, err := POST(g.url("/auth"), "", authReq{
		Sign:      sign,
		Timestamp: tp,
		AppKey:    g.AppKey,
	}, g.Timeout)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return "", 0, err
	}
	var resp authResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return "", 0, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return "", 0, fmt.Errorf("resp.Code %d reps.Msg %s", resp.Code, resp.Msg)
	}
	if resp.Data == nil {
		log.Errorf("resp.Data is nil")
		return "", 0, errors.New("resp.Data is nil")
	}
	token := resp.Data.Token
	expireAtMs, err := strconv.ParseInt(resp.Data.ExpireTime, 10, 64)
	if err != nil {
		log.Errorf(" parseInt failed expire time %s ms", resp.Data.ExpireTime)
		return "", 0, fmt.Errorf("parseInt failed expire time %s ms", resp.Data.ExpireTime)
	}
	log.Tracef("get auth ok token %s expireAtSec %d", token, expireAtMs/1000)
	return token, expireAtMs / 1000, nil
}

func (g *GeTuiAuth) DelAuth(token string) error {
	result, err := DELETE(g.url(fmt.Sprintf("/auth/%s", token)), "", nil, g.Timeout)
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
		return fmt.Errorf("resp.Code %d reps.Msg %s", resp.Code, resp.Msg)
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
