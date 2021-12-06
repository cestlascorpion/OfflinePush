package core

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type RestyClient struct {
	client *resty.Client
}

func NewRestyClient(timeout time.Duration) (*RestyClient, error) {
	return &RestyClient{
		client: resty.New().SetTimeout(timeout).EnableTrace().
			SetHeader("Content-Type", "application/json;charset=utf-8").
			SetHeader("Accept", "application/json"),
	}, nil
}

func NewRestyClientWithClient(client *http.Client, timeout time.Duration) (*RestyClient, error) {
	if client == nil {
		return nil, errors.New("invalid parameter")
	}
	return &RestyClient{
		client: resty.NewWithClient(client).SetTimeout(timeout).EnableTrace().
			SetHeader("Content-Type", "application/json;charset=utf-8").
			SetHeader("Accept", "application/json"),
	}, nil
}

func (r *RestyClient) GET(url, token string, body interface{}) ([]byte, error) {
	if len(token) > 0 {
		r.client.SetHeader("token", token)
	}
	log.Debugf("get url %s", url)
	resp, err := r.client.R().SetBody(body).Get(url)
	if err != nil {
		log.Errorf("resty get err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func (r *RestyClient) POST(url, token string, body interface{}) ([]byte, error) {
	if len(token) > 0 {
		r.client.SetHeader("token", token)
	}
	log.Debugf("post url %s", url)
	resp, err := r.client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorf("resty post err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func (r *RestyClient) PUT(url, token string, body interface{}) ([]byte, error) {
	if len(token) > 0 {
		r.client.SetHeader("token", token)
	}
	log.Debugf("put url %s", url)
	resp, err := r.client.R().SetBody(body).Put(url)
	if err != nil {
		log.Errorf("resty put err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func (r *RestyClient) DELETE(url, token string, body interface{}) ([]byte, error) {
	if len(token) > 0 {
		r.client.SetHeader("token", token)
	}
	log.Debugf("del url %s", url)
	resp, err := r.client.R().SetBody(body).Delete(url)
	if err != nil {
		log.Errorf("resty Delete err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}
