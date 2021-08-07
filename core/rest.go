package core

import (
	"errors"
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
	resp, err := client.R().SetBody(body).Get(url)
	if err != nil {
		log.Fatalf("resty get err %+v", err)
		return nil, err
	}
	if !resp.IsSuccess() {
		log.Errorf("resty get failed %+v", resp.Status())
		return nil, errors.New(resp.Status())
	}
	log.Tracef("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func POST(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json")
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	resp, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Fatalf("resty post err %+v", err)
		return nil, err
	}
	if !resp.IsSuccess() {
		log.Errorf("resty post failed %+v", resp.Status())
		return nil, errors.New(resp.Status())
	}
	log.Tracef("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func DELETE(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json")
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	resp, err := client.R().SetBody(body).Delete(url)
	if err != nil {
		log.Fatalf("resty Delete err %+v", err)
		return nil, err
	}
	if !resp.IsSuccess() {
		log.Errorf("resty Delete failed %+v", resp.Status())
		return nil, errors.New(resp.Status())
	}
	log.Tracef("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}
