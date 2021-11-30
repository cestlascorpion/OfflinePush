package core

import (
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func GET(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json")
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	log.Debugf("get url %s", url)
	resp, err := client.R().SetBody(body).Get(url)
	if err != nil {
		log.Errorf("resty get err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func POST(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json")
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	log.Debugf("post url %s", url)
	resp, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorf("resty post err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func DELETE(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json")
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	log.Debugf("del url %s", url)
	resp, err := client.R().SetBody(body).Delete(url)
	if err != nil {
		log.Errorf("resty Delete err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}
