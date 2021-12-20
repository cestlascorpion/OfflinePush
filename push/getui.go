package push

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cestlascorpion/offlinepush/core"
	log "github.com/sirupsen/logrus"
)

type GeTuiPush struct {
	apiUrl string
	client *core.RestyClient
}

func NewGeTuiPush(apiUrl, appId string, hc *http.Client) (*GeTuiPush, error) {
	client, err := core.NewRestyClient(hc)
	if err != nil {
		return nil, err
	}
	return &GeTuiPush{
		apiUrl: fmt.Sprintf(apiUrl, appId),
		client: client,
	}, nil
}

func (g *GeTuiPush) PushSingleByCid(request *SingleReq, token string) (map[string]map[string]string, error) {
	result, err := g.client.POST(g.url("/push/single/cid"), token, request)
	if err != nil {
		log.Errorf("PushSingleByCid() POST err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushSingleByCid() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("PushSingleByCid() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushSingleByCid data %+v", resp.Data)
	return g.resp2PushDetail(resp.Data)
}

func (g *GeTuiPush) PushSingleByAlias(request *SingleReq, token string) (map[string]map[string]string, error) {
	result, err := g.client.POST(g.url("/push/single/alias"), token, request)
	if err != nil {
		log.Errorf("PushSingleByAlias() POST err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushSingleByAlias() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("PushSingleByAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushSingleByAlias data %+v", resp.Data)
	return g.resp2PushDetail(resp.Data)
}

func (g *GeTuiPush) PushBatchByCid(request *BatchReq, token string) (map[string]map[string]string, error) {
	result, err := g.client.POST(g.url("/push/single/batch/cid"), token, request)
	if err != nil {
		log.Errorf("PushBatchByCid() POST err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushBatchByCid() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("PushBatchByCid() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushBatchByCid data %+v", resp.Data)
	return g.resp2PushDetail(resp.Data)
}

func (g *GeTuiPush) PushBatchByAlias(request *BatchReq, token string) (map[string]map[string]string, error) {
	result, err := g.client.POST(g.url("/push/single/batch/alias"), token, request)
	if err != nil {
		log.Errorf("PushBatchByAlias() POST err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushBatchByAlias() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("PushBatchByAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushBatchByAlias data %+v", resp.Data)
	return g.resp2PushDetail(resp.Data)
}

func (g *GeTuiPush) CreateMsg(request *CreateReq, token string) (string, error) {
	result, err := g.client.POST(g.url("/push/list/message"), token, request)
	if err != nil {
		log.Errorf("CreateMsg() POST err %+v", err)
		return "", err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("CreateMsg() unmarshal err %+v", err)
		return "", err
	}
	if resp.Code != 0 {
		log.Errorf("CreateMsg() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return "", fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("CreateMsg data %+v", resp.Data)
	return g.resp2PushTaskId(resp.Data)
}

func (g *GeTuiPush) PushListByCid(request *ListReq, token string) (map[string]map[string]string, error) {
	result, err := g.client.POST(g.url("/push/list/cid"), token, request)
	if err != nil {
		log.Errorf("PushListByCid() POST err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushListByCid() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("PushListByCid() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushListByCid data %+v", resp.Data)
	return g.resp2PushDetail(resp.Data)
}

func (g *GeTuiPush) PushListByAlias(request *ListReq, token string) (map[string]map[string]string, error) {
	result, err := g.client.POST(g.url("/push/list/alias"), token, request)
	if err != nil {
		log.Errorf("PushListByAlias() POST err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushListByAlias() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("PushListByAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushListByAlias data %+v", resp.Data)
	return g.resp2PushDetail(resp.Data)
}

func (g *GeTuiPush) PushAll(request *AllReq, token string) (string, error) {
	result, err := g.client.POST(g.url("/push/all"), token, request)
	if err != nil {
		log.Errorf("PushAll() POST err %+v", err)
		return "", err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushAll() unmarshal err %+v", err)
		return "", err
	}
	if resp.Code != 0 {
		log.Errorf("PushAll() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return "", fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushAll data %+v", resp.Data)
	return g.resp2PushTaskId(resp.Data)
}

func (g *GeTuiPush) PushByTag(request *ByTagReq, token string) (string, error) {
	result, err := g.client.POST(g.url("/push/tag"), token, request)
	if err != nil {
		log.Errorf("PushByTag() POST err %+v", err)
		return "", err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushByTag() unmarshal err %+v", err)
		return "", err
	}
	if resp.Code != 0 {
		log.Errorf("PushByTag() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return "", fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushByTag data %+v", resp.Data)
	return g.resp2PushTaskId(resp.Data)
}

func (g *GeTuiPush) PushByFastCustomTag(request *ByTagReq, token string) (string, error) {
	result, err := g.client.POST(g.url("/push/fast_custom_tag"), token, request)
	if err != nil {
		log.Errorf("PushByFastCustomTag() POST err %+v", err)
		return "", err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("PushByFastCustomTag() unmarshal err %+v", err)
		return "", err
	}
	if resp.Code != 0 {
		log.Errorf("PushByFastCustomTag() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return "", fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("PushByFastCustomTag data %+v", resp.Data)
	return g.resp2PushTaskId(resp.Data)
}

func (g *GeTuiPush) StopPush(taskId, token string) (bool, error) {
	result, err := g.client.DELETE(g.url(fmt.Sprintf("/task/%s", taskId)), token, nil)
	if err != nil {
		log.Errorf("StopPush() POST err %+v", err)
		return false, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("StopPush() unmarshal err %+v", err)
		return false, err
	}
	if resp.Code != 0 {
		log.Errorf("StopPush() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return false, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return true, nil
}

func (g *GeTuiPush) DeleteScheduleTask(taskId, token string) (bool, error) {
	result, err := g.client.DELETE(g.url(fmt.Sprintf("/task/schedule/%s", taskId)), token, nil)
	if err != nil {
		log.Errorf("DeleteScheduleTask() POST err %+v", err)
		return false, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("DeleteScheduleTask() unmarshal err %+v", err)
		return false, err
	}
	if resp.Code != 0 {
		log.Errorf("DeleteScheduleTask() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return false, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return true, nil
}

func (g *GeTuiPush) QueryScheduleTask(taskId, token string) (map[string]string, error) {
	result, err := g.client.GET(g.url(fmt.Sprintf("/task/schedule/%s", taskId)), token, nil)
	if err != nil {
		log.Errorf("QueryScheduleTask() GET err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryScheduleTask() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryScheduleTask() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2Schedule(resp.Data)
}

func (g *GeTuiPush) QueryDetail(taskId, cId, token string) ([][2]string, error) {
	result, err := g.client.GET(g.url(fmt.Sprintf("/task/detail/%s/%s", cId, taskId)), token, nil)
	if err != nil {
		log.Errorf("QueryDetail() GET err %+v", err)
		return nil, err
	}
	var resp pushResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryDetail() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryDetail() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2Detail(resp.Data)
}

type pushResp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (g *GeTuiPush) resp2PushDetail(data map[string]interface{}) (map[string]map[string]string, error) {
	result := make(map[string]map[string]string)
	if data == nil {
		return result, errors.New("empty result")
	}
	for taskId, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return result, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		detail := make(map[string]string)
		for k, v := range value {
			status, ok := v.(string)
			if !ok {
				return result, fmt.Errorf("unknown type of v %+v", v)
			}
			detail[k] = status
		}
		result[taskId] = detail
	}
	return result, nil
}

func (g *GeTuiPush) resp2PushTaskId(data map[string]interface{}) (string, error) {
	if data == nil {
		return "", errors.New("empty result")
	}
	for taskId, unknown := range data {
		if taskId != "taskid" {
			continue
		}
		value, ok := unknown.(string)
		if !ok {
			return "", fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		return value, nil
	}
	return "", errors.New("taskid not found")
}

func (g *GeTuiPush) resp2Schedule(data map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string)
	if data == nil {
		return result, errors.New("empty result")
	}
	if len(data) != 1 {
		return result, errors.New("unexpected result")
	}
	for _, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return result, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		for k, v := range value {
			desc, ok := v.(string)
			if !ok {
				return result, fmt.Errorf("unknown type of v %+v", v)
			}
			result[k] = desc
		}
	}
	return result, nil
}

func (g *GeTuiPush) resp2Detail(data map[string]interface{}) ([][2]string, error) {
	result := make([][2]string, 0)
	if data == nil {
		return result, errors.New("empty result")
	}
	if len(data) != 1 {
		return result, errors.New("unexpected result")
	}
	for key, unknown := range data {
		if key != "detail" {
			continue
		}
		value, ok := unknown.([]interface{})
		if !ok {
			return result, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		for _, item := range value {
			item, ok := item.(map[string]interface{})
			if !ok {
				return result, fmt.Errorf("unknown type of item %+v", item)
			}
			ts, ok := item["time"]
			if !ok {
				return result, fmt.Errorf("time field not found in item %+v", item)
			}
			t, ok := ts.(string)
			if !ok {
				return result, fmt.Errorf("unknown type of ts %+v", item)
			}
			event, ok := item["event"]
			if !ok {
				return result, fmt.Errorf("event field not found in item %+v", item)
			}
			e, ok := event.(string)
			if !ok {
				return result, fmt.Errorf("unknown type of event %+v", item)
			}
			result = append(result, [2]string{t, e})
		}
	}
	return result, nil
}

func (g *GeTuiPush) url(path string) string {
	return g.apiUrl + path
}
