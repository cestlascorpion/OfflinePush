package core

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type RestyClient struct {
	hc *http.Client
}

func NewRestyClient(client *http.Client) (*RestyClient, error) {
	if client == nil {
		return nil, errors.New("invalid paramter")
	}
	return &RestyClient{
		hc: client,
	}, nil
}

func (r *RestyClient) GET(url, token string, body interface{}) ([]byte, error) {
	client := resty.NewWithClient(r.hc).EnableTrace().
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

func (r *RestyClient) POST(url, token string, body interface{}) ([]byte, error) {
	client := resty.NewWithClient(r.hc).EnableTrace().
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

func (r *RestyClient) PUT(url, token string, body interface{}) ([]byte, error) {
	client := resty.NewWithClient(r.hc).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json")
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	log.Debugf("put url %s", url)
	resp, err := client.R().SetBody(body).Put(url)
	if err != nil {
		log.Errorf("resty put err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func (r *RestyClient) DELETE(url, token string, body interface{}) ([]byte, error) {
	client := resty.NewWithClient(r.hc).EnableTrace().
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

func GET(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json").SetCloseConnection(true)
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
		SetHeader("Accept", "application/json").SetCloseConnection(true)
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

func PUT(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json").SetCloseConnection(true)
	if len(token) > 0 {
		client.SetHeader("token", token)
	}
	log.Debugf("put url %s", url)
	resp, err := client.R().SetBody(body).Put(url)
	if err != nil {
		log.Errorf("resty put err %+v", err)
		return nil, err
	}
	log.Debugf("%+v", resp.Request.TraceInfo())
	return resp.Body(), nil
}

func DELETE(url, token string, body interface{}, timeout time.Duration) ([]byte, error) {
	client := resty.New().SetTimeout(timeout).EnableTrace().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Accept", "application/json").SetCloseConnection(true)
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
