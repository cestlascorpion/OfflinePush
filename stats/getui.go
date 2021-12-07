package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type GeTuiStats struct {
	ApiUrl string
	client *core.RestyClient
}

func NewGeTuiStats(apiUrl, appId string, hc *http.Client) (*GeTuiStats, error) {
	client, err := core.NewRestyClient(hc)
	if err != nil {
		return nil, err
	}
	return &GeTuiStats{
		ApiUrl: fmt.Sprintf(apiUrl, appId),
		client: client,
	}, nil
}

func (g *GeTuiStats) GetTasks(taskIds []string, token string) (map[string]*proto.BaseStatics, error) {
	result, err := g.client.GET(g.url(fmt.Sprintf("/report/push/task/%s", strings.Join(taskIds, ","))), token, nil)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp reportResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		if resp.Code == 10001 {
			return nil, errors.New(core.InvalidTokenErr)
		}
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2Tasks(resp.Data)
}

func (g *GeTuiStats) GetTaskGroup(group, token string) (map[string]*proto.BaseStatics, error) {
	result, err := g.client.GET(g.url(fmt.Sprintf("/report/push/task_group/%s", group)), token, nil)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp reportResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		if resp.Code == 10001 {
			return nil, errors.New(core.InvalidTokenErr)
		}
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2Group(resp.Data)
}

func (g *GeTuiStats) GetPushCount(token string) ([]*proto.BasePushCount, error) {
	result, err := g.client.GET(g.url("/report/push/count"), token, nil)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp reportResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		if resp.Code == 10001 {
			return nil, errors.New(core.InvalidTokenErr)
		}
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2Count(resp.Data)
}

func (g *GeTuiStats) GetPushDataByDay(date time.Time, token string) (map[string]*proto.BaseStatics, error) {
	result, err := g.client.GET(g.url(fmt.Sprintf("/report/push/date/%s", date.Format("2006-01-02"))), token, nil)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp reportResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		if resp.Code == 10001 {
			return nil, errors.New(core.InvalidTokenErr)
		}
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2PushByDay(resp.Data)
}

func (g *GeTuiStats) GetUserDataByDay(date time.Time, token string) (map[string]map[string]int32, error) {
	result, err := g.client.GET(g.url(fmt.Sprintf("/report/user/date/%s", date.Format("2006-01-02"))), token, nil)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp reportResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		if resp.Code == 10001 {
			return nil, errors.New(core.InvalidTokenErr)
		}
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2UserByDay(resp.Data)
}

func (g *GeTuiStats) GetOnlineUserBy24H(token string) (map[int64]int32, error) {
	result, err := g.client.GET(g.url("/report/online_user"), token, nil)
	if err != nil {
		log.Errorf("post failed err %+v", err)
		return nil, err
	}
	var resp reportResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("unmarshal failed err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	return g.resp2OnlineBy24H(resp.Data)
}

type reportResp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (g *GeTuiStats) resp2Tasks(data map[string]interface{}) (map[string]*proto.BaseStatics, error) {
	stats := make(map[string]*proto.BaseStatics)
	if data == nil {
		return stats, nil
	}
	for taskId, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return stats, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		statics, err := g.data2BaseStatics(value)
		if err != nil {
			return stats, err
		}
		stats[taskId] = statics
	}
	return stats, nil
}

func (g *GeTuiStats) resp2Group(data map[string]interface{}) (map[string]*proto.BaseStatics, error) {
	stats := make(map[string]*proto.BaseStatics)
	if data == nil {
		return stats, nil
	}
	for groupId, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return stats, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		statics, err := g.data2BaseStatics(value)
		if err != nil {
			return stats, err
		}
		stats[groupId] = statics
	}
	return stats, nil
}

func (g *GeTuiStats) resp2Count(data map[string]interface{}) ([]*proto.BasePushCount, error) {
	stats := make([]*proto.BasePushCount, 0)
	if data == nil {
		return stats, nil
	}
	for manufacturer, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return stats, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		for tag, detail := range value {
			count := &proto.BasePushCount{}
			final, ok := detail.(map[string]interface{})
			if !ok {
				return stats, fmt.Errorf("unknown type of %s -> interface{} %+v", tag, detail)
			}
			for k, v := range final {
				switch k {
				case "total_num":
					num, ok := v.(float64)
					if !ok {
						str, ok := v.(string)
						if !ok {
							return stats, fmt.Errorf("unknown type of v %+v", v)
						} else {
							n, err := strconv.ParseInt(str, 10, 64)
							if err != nil {
								return stats, fmt.Errorf("unknown type of v %+v", v)
							}
							count.TotalNum = int32(n)
						}
					} else {
						count.TotalNum = int32(num)
					}
				case "push_num":
					num, ok := v.(float64)
					if !ok {
						str, ok := v.(string)
						if !ok {
							return stats, fmt.Errorf("unknown type of v %+v", v)
						} else {
							n, err := strconv.ParseInt(str, 10, 64)
							if err != nil {
								return stats, fmt.Errorf("unknown type of v %+v", v)
							}
							count.PushNum = int32(n)
						}
					} else {
						count.PushNum = int32(num)
					}
				case "remain_num":
					num, ok := v.(float64)
					if !ok {
						str, ok := v.(string)
						if !ok {
							return stats, fmt.Errorf("unknown type of v %+v", v)
						} else {
							n, err := strconv.ParseInt(str, 10, 64)
							if err != nil {
								return stats, fmt.Errorf("unknown type of v %+v", v)
							}
							count.RemainNum = int32(n)
						}
					} else {
						count.RemainNum = int32(num)
					}
				case "limit":
					limit, ok := v.(bool)
					if !ok {
						return stats, fmt.Errorf("unknown type of v %+v", v)
					}
					count.HasLimit = limit
				default:
					log.Warnf("unknown field %s %+v", k, v)
				}
			}
			count.Manufacturer = manufacturer
			count.DescribeTag = tag
			stats = append(stats, count)
		}
	}
	return stats, nil
}

func (g *GeTuiStats) resp2PushByDay(data map[string]interface{}) (map[string]*proto.BaseStatics, error) {
	stats := make(map[string]*proto.BaseStatics, 0)
	if data == nil {
		return stats, nil
	}
	for date, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return stats, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		statics, err := g.data2BaseStatics(value)
		if err != nil {
			return stats, err
		}
		stats[date] = statics
	}
	return stats, nil
}

func (g *GeTuiStats) resp2UserByDay(data map[string]interface{}) (map[string]map[string]int32, error) {
	stats := make(map[string]map[string]int32)
	if data == nil {
		return stats, nil
	}
	for date, unknown := range data {
		statics := make(map[string]int32)
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return stats, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		for k, v := range value {
			switch k {
			case "accumulative_num":
				num, ok := v.(float64)
				if !ok {
					return stats, fmt.Errorf("unknown type of v %+v", v)
				}
				statics["accumulative_num"] = int32(num)
			case "register_num":
				num, ok := v.(float64)
				if !ok {
					return stats, fmt.Errorf("unknown type of v %+v", v)
				}
				statics["register_num"] = int32(num)
			case "active_num":
				num, ok := v.(float64)
				if !ok {
					return stats, fmt.Errorf("unknown type of v %+v", v)
				}
				statics["active_num"] = int32(num)
			case "online_num":
				num, ok := v.(float64)
				if !ok {
					return stats, fmt.Errorf("unknown type of v %+v", v)
				}
				statics["online_num"] = int32(num)
			default:
				log.Warnf("unknown field %s %+v", k, v)
			}
		}
		stats[date] = statics
	}
	return stats, nil
}

func (g *GeTuiStats) resp2OnlineBy24H(data map[string]interface{}) (map[int64]int32, error) {
	stats := make(map[int64]int32)
	for _, unknown := range data {
		value, ok := unknown.(map[string]interface{})
		if !ok {
			return stats, fmt.Errorf("unknown type of interface{} %+v", unknown)
		}
		for t, n := range value {
			ts, err := strconv.ParseInt(t, 10, 64)
			if err != nil {
				return stats, fmt.Errorf("parse uint64 failed t %s", t)
			}
			num, ok := n.(float64)
			if !ok {
				return stats, fmt.Errorf("unknown type of n %+v", n)
			}
			stats[ts] = int32(num)
		}
	}
	return stats, nil
}

func (g *GeTuiStats) url(path string) string {
	return g.ApiUrl + path
}

func (g *GeTuiStats) data2BaseStatics(data map[string]interface{}) (*proto.BaseStatics, error) {
	statics := &proto.BaseStatics{
		Total:      &proto.BaseStatics_Total{},
		DetailList: make([]*proto.BaseStatics_Detail, 0),
		ActionList: make([]*proto.BaseStatics_Action, 0),
	}
	for manufacturer, summary := range data {
		if manufacturer == "total" {
			final, ok := summary.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("unknown type of total->interface{} %+v", summary)
			}
			for k, v := range final {
				switch k {
				case "target_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					statics.Total.Target = int32(num)
				case "receive_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					statics.Total.Receive = int32(num)
				case "display_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					statics.Total.Display = int32(num)
				case "click_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					statics.Total.Click = int32(num)
				case "msg_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					statics.Total.Msg = int32(num)
				default:
					log.Warnf("unknown field %s %+v", k, v)
				}
			}
		} else if manufacturer == "actionCntMap" {
			final, ok := summary.(map[string]interface{})
			for !ok {
				return nil, fmt.Errorf("unknown type of actionCntMap->interface{} %+v", summary)
			}
			for k, v := range final {
				num, ok := v.(float64)
				if !ok {
					return nil, fmt.Errorf("unknown type of v %+v", v)
				}
				statics.ActionList = append(statics.ActionList, &proto.BaseStatics_Action{
					Key:   k,
					Value: int32(num),
				})
			}
		} else {
			final, ok := summary.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("unknown type of %s->interface{} %+v", manufacturer, summary)
			}
			detail := &proto.BaseStatics_Detail{}
			for k, v := range final {
				switch k {
				case "target_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					detail.Target = int32(num)
				case "receive_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					detail.Receive = int32(num)
				case "display_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					detail.Display = int32(num)
				case "click_num":
					num, ok := v.(float64)
					if !ok {
						return nil, fmt.Errorf("unknown type of v %+v", v)
					}
					detail.Click = int32(num)
				default:
					log.Warnf("unknown field %s %+v", k, v)
				}
			}
			detail.Manufacturer = manufacturer
			statics.DetailList = append(statics.DetailList, detail)
		}
	}
	return statics, nil
}
