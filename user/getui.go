package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cestlascorpion/offlinepush/core"
	log "github.com/sirupsen/logrus"
)

type GeTuiUser struct {
	apiUrl string
	client *core.RestyClient
}

func NewGeTuiUser(apiUrl, appId string, hc *http.Client) (*GeTuiUser, error) {
	client, err := core.NewRestyClient(hc)
	if err != nil {
		return nil, err
	}
	return &GeTuiUser{
		apiUrl: fmt.Sprintf(apiUrl, appId),
		client: client,
	}, nil
}

func (g *GeTuiUser) BindAlias(ctx context.Context, list *AliasList, token string) error {
	result, err := g.client.POST(ctx, g.url("/user/alias"), token, list)
	if err != nil {
		log.Errorf("BindAlias() POST err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("BindAlias() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("BindAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("BindAlias data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) QueryAliasByCid(ctx context.Context, cid string, token string) (string, error) {
	result, err := g.client.GET(ctx, g.url(fmt.Sprintf("/user/alias/cid/%s", cid)), token, nil)
	if err != nil {
		log.Errorf("QueryAliasByCid() GET err %+v", err)
		return "", err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryAliasByCid() unmarshal err %+v", err)
		return "", err
	}
	if resp.Code != 0 {
		log.Errorf("QueryAliasByCid() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return "", fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryAliasByCid data %+v", resp.Data)

	alias := ""
	if unknown, ok := resp.Data["alias"]; ok {
		ret, ok := unknown.(string)
		if !ok {
			return alias, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		alias = ret
	}
	return alias, nil
}

func (g *GeTuiUser) QueryCidByAlias(ctx context.Context, alias string, token string) ([]string, error) {
	result, err := g.client.GET(ctx, g.url(fmt.Sprintf("/user/cid/alias/%s", alias)), token, nil)
	if err != nil {
		log.Errorf("QueryCidByAlias() GET err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryCidByAlias() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryCidByAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryCidByAlias data %+v", resp.Data)

	cidList := make([]string, 0)
	if unknown, ok := resp.Data["cid"]; ok {
		ret, ok := unknown.([]interface{})
		if !ok {
			return nil, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		for i := range ret {
			cid, ok := ret[i].(string)
			if !ok {
				log.Warnf("unknown type of v %+v %T", ret[i], ret[i])
				continue
			}
			cidList = append(cidList, cid)
		}
	}
	return cidList, nil
}

func (g *GeTuiUser) UnbindAlias(ctx context.Context, list *AliasList, token string) error {
	result, err := g.client.DELETE(ctx, g.url("/user/alias"), token, list)
	if err != nil {
		log.Errorf("UnbindAlias() DELETE err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("UnbindAlias() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("UnbindAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("UnbindAlias data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) RevokeAlias(ctx context.Context, alias string, token string) error {
	result, err := g.client.DELETE(ctx, g.url(fmt.Sprintf("/user/alias/%s", alias)), token, nil)
	if err != nil {
		log.Errorf("RevokeAlias() DELETE err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("RevokeAlias() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("RevokeAlias() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("RevokeAlias data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) BindUserWithTag(ctx context.Context, cid string, list *CustomTagList, token string) error {
	result, err := g.client.POST(ctx, g.url(fmt.Sprintf("/user/custom_tag/cid/%s", cid)), token, list)
	if err != nil {
		log.Errorf("BindUserWithTag() POST err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("BindUserWithTag() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("BindUserWithTag() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("BindUserWithTag data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) BindTagWithUser(ctx context.Context, tag string, list *CidList, token string) (map[string]bool, error) {
	result, err := g.client.PUT(ctx, g.url(fmt.Sprintf("/user/custom_tag/batch/%s", tag)), token, list)
	if err != nil {
		log.Errorf("BindTagWithUser() PUT err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("BindTagWithUser() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("BindTagWithUser() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("BindTagWithUser data %+v", resp.Data)

	success := make(map[string]bool)
	for cid, unknown := range resp.Data {
		ret, ok := unknown.(string)
		if !ok {
			log.Warnf("unknown type %+v %T", unknown, unknown)
			continue
		}
		if ret == "true" {
			success[cid] = true
		} else {
			success[cid] = false
		}
	}
	return success, nil
}

func (g *GeTuiUser) UnbindTagFromUser(ctx context.Context, tag string, list *CidList, token string) (map[string]bool, error) {
	result, err := g.client.DELETE(ctx, g.url(fmt.Sprintf("/user/custom_tag/batch/%s", tag)), token, list)
	if err != nil {
		log.Errorf("UnbindTagFromUser() DELETE err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("UnbindTagFromUser() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("UnbindTagFromUser() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("UnbindTagFromUser data %+v", resp.Data)

	success := make(map[string]bool)
	for cid, unknown := range resp.Data {
		ret, ok := unknown.(string)
		if !ok {
			log.Warnf("unknown type %+v %T", unknown, unknown)
			continue
		}
		if ret == "true" {
			success[cid] = true
		} else {
			success[cid] = false
		}
	}
	return success, nil
}

func (g *GeTuiUser) QueryUserTag(ctx context.Context, cid string, token string) ([]string, error) {
	result, err := g.client.GET(ctx, g.url(fmt.Sprintf("/user/custom_tag/cid/%s", cid)), token, nil)
	if err != nil {
		log.Errorf("QueryUserTag() GET err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryUserTag() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryUserTag() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryUserTag data %+v", resp.Data)

	cidList := make([]string, 0)
	if unknown, ok := resp.Data[cid]; ok {
		ret, ok := unknown.([]interface{})
		if !ok {
			return nil, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		for i := range ret {
			cid, ok := ret[i].(string)
			if !ok {
				log.Warnf("unknown type of v %+v %T", ret[i], ret[i])
				continue
			}
			cidList = append(cidList, cid)
		}
	}
	return cidList, nil
}

func (g *GeTuiUser) AddBlackList(ctx context.Context, cidList []string, token string) error {
	result, err := g.client.POST(ctx, g.url(fmt.Sprintf("/user/black/cid/%s", strings.Join(cidList, ","))), token, nil)
	if err != nil {
		log.Errorf("AddBlackList() POST err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("AddBlackList() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("AddBlackList() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("AddBlackList data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) DelBlackList(ctx context.Context, cidList []string, token string) error {
	result, err := g.client.DELETE(ctx, g.url(fmt.Sprintf("/user/black/cid/%s", strings.Join(cidList, ","))), token, nil)
	if err != nil {
		log.Errorf("DelBlackList() DELETE err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("DelBlackList() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("DelBlackList() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("DelBlackList data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) QueryUserStatus(ctx context.Context, cidList []string, token string) (map[string]map[string]string, error) {
	result, err := g.client.GET(ctx, g.url(fmt.Sprintf("/user/status/%s", strings.Join(cidList, ","))), token, nil)
	if err != nil {
		log.Errorf("QueryUserStatus() GET err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryUserStatus() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryUserStatus() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryUserStatus data %+v", resp.Data)

	status := make(map[string]map[string]string)
	for cid, unknown := range resp.Data {
		detail, ok := unknown.(map[string]interface{})
		if !ok {
			return status, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		if len(detail) == 0 {
			continue
		}
		meta := make(map[string]string)
		for k, v := range detail {
			desc, ok := v.(string)
			if !ok {
				log.Warnf("unknown type of v %+v %T", v, v)
				continue
			}
			meta[k] = desc
		}
		status[cid] = meta
	}
	return status, nil
}

func (g *GeTuiUser) QueryDeviceStatus(ctx context.Context, cidList []string, token string) (map[string]map[string]string, error) {
	result, err := g.client.GET(ctx, g.url(fmt.Sprintf("/user/deviceStatus/%s", strings.Join(cidList, ","))), token, nil)
	if err != nil {
		log.Errorf("QueryDeviceStatus() GET err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryDeviceStatus() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryDeviceStatus() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryDeviceStatus data %+v", resp.Data)

	status := make(map[string]map[string]string)
	for cid, unknown := range resp.Data {
		ret, err := g.decodeNestedResp(unknown)
		if err != nil {
			log.Warnf("decode unknown failed err %+v", err)
			continue
		}
		status[cid] = ret
	}
	return status, nil
}

func (g *GeTuiUser) QueryUserInfo(ctx context.Context, cidList []string, token string) ([]string, map[string]map[string]string, error) {
	result, err := g.client.GET(ctx, g.url(fmt.Sprintf("/user/detail/%s", strings.Join(cidList, ","))), token, nil)
	if err != nil {
		log.Errorf("QueryUserInfo() GET err %+v", err)
		return nil, nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryUserInfo() unmarshal err %+v", err)
		return nil, nil, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryUserInfo() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryUserInfo data %+v", resp.Data)

	invalid := make([]string, 0)
	if unknown, ok := resp.Data["invalidCids"]; ok {
		ret, ok := unknown.([]interface{})
		if !ok {
			return nil, nil, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		for i := range ret {
			cid, ok := ret[i].(string)
			if !ok {
				log.Warnf("unknown type of v %+v %T", ret[i], ret[i])
				continue
			}
			invalid = append(invalid, cid)
		}
	}
	validInfo := make(map[string]map[string]string)
	if unknown, ok := resp.Data["validCids"]; ok {
		ret, ok := unknown.(map[string]interface{})
		if !ok {
			return nil, nil, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		for cid, v := range ret {
			detail, ok := v.(map[string]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("unknown type of v %+v %T", v, v)
			}
			meta := make(map[string]string)
			for k, v := range detail {
				switch k {
				case "client_app_id":
					desc, ok := v.(string)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = desc
				case "package_name":
					desc, ok := v.(string)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = desc
				case "device_token":
					desc, ok := v.(string)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = desc
				case "phone_type":
					desc, ok := v.(float64)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = strconv.FormatInt(int64(desc), 10)
				case "phone_model":
					desc, ok := v.(string)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = desc
				case "notification_switch":
					desc, ok := v.(bool)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = strconv.FormatBool(desc)
				case "create_time":
					desc, ok := v.(string)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = desc
				case "login_freq":
					desc, ok := v.(float64)
					if !ok {
						log.Warnf("unknown type of v %+v %T", v, v)
						break
					}
					meta[k] = strconv.FormatInt(int64(desc), 10)
				}
			}
			validInfo[cid] = meta
		}
	}
	return invalid, validInfo, nil
}

func (g *GeTuiUser) SetPushBadge(ctx context.Context, cidList []string, op *Operation, token string) error {
	result, err := g.client.POST(ctx, g.url(fmt.Sprintf("/user/badge/%s", strings.Join(cidList, ","))), token, op)
	if err != nil {
		log.Errorf("SetPushBadge() POST err %+v", err)
		return err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("SetPushBadge() unmarshal err %+v", err)
		return err
	}
	if resp.Code != 0 {
		log.Errorf("SetPushBadge() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("SetPushBadge data %+v", resp.Data)
	return nil
}

func (g *GeTuiUser) QueryUserCount(ctx context.Context, list *ComplexTagList, token string) (int, error) {
	result, err := g.client.POST(ctx, g.url("/user/count"), token, list)
	if err != nil {
		log.Errorf("QueryUserCount() POST err %+v", err)
		return 0, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("QueryUserCount() unmarshal err %+v", err)
		return 0, err
	}
	if resp.Code != 0 {
		log.Errorf("QueryUserCount() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return 0, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("QueryUserCount data %+v", resp.Data)

	count := 0
	if unknown, ok := resp.Data["user_count"]; ok {
		ret, ok := unknown.(float64)
		if !ok {
			return count, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}
		count = int(ret)
	}
	return count, nil
}

func (g *GeTuiUser) ManageCidAndDeviceToken(ctx context.Context, manufacturer string, dtList *CidAndDeviceTokenList, token string) ([]*DTError, error) {
	result, err := g.client.POST(ctx, g.url(fmt.Sprintf("/user/bind_dt/%s", manufacturer)), token, dtList)
	if err != nil {
		log.Errorf("ManageCidAndDeviceToken() POST err %+v", err)
		return nil, err
	}
	var resp userResp
	err = json.Unmarshal(result, &resp)
	if err != nil {
		log.Errorf("ManageCidAndDeviceToken() unmarshal err %+v", err)
		return nil, err
	}
	if resp.Code != 0 {
		log.Errorf("ManageCidAndDeviceToken() resp.Code %d, resp.Msg %s", resp.Code, resp.Msg)
		return nil, fmt.Errorf("resp.Code %d resp.Msg %s", resp.Code, resp.Msg)
	}
	log.Debugf("ManageCidAndDeviceToken data %+v", resp.Data)

	errorList := make([]*DTError, 0)
	if unknown, ok := resp.Data["errorList"]; ok {
		list, ok := unknown.([]interface{})
		if !ok {
			log.Errorf("unknown type %+v %T", unknown, unknown)
			return nil, fmt.Errorf("unknown type %+v %T", unknown, unknown)
		}

		for i := range list {
			data, ok := list[i].(map[string]interface{})
			if !ok {
				log.Warnf("unknown type of data %+v %T", list[i], list[i])
				return nil, fmt.Errorf("unknown type of data %+v %T", list[i], list[i])
			}

			cid, ok := data["cid"]
			if !ok {
				log.Warnf("cid not found")
				continue
			}
			c, ok := cid.(string)
			if !ok {
				log.Errorf("unknown type of cid %+v %T", cid, cid)
				return nil, fmt.Errorf("unknown type of cid %+v %T", cid, cid)
			}
			deviceToken, ok := data["device_token"]
			if !ok {
				log.Warnf("device_token not found")
				continue
			}
			d, ok := deviceToken.(string)
			if !ok {
				log.Errorf("unknown type of deviceToken %+v %T", cid, cid)
				return nil, fmt.Errorf("unknown type of deviceToken %+v %T", deviceToken, deviceToken)
			}
			errorCode, ok := data["error_code"]
			if !ok {
				log.Warnf("error_code not found")
				continue
			}
			e, ok := errorCode.(float64)
			if !ok {
				log.Errorf("unknown type of errorCode %+v %T", cid, cid)
				return nil, fmt.Errorf("unknown type of errorCode %+v %T", errorCode, errorCode)
			}

			errorList = append(errorList, &DTError{
				Cid:         c,
				DeviceToken: d,
				ErrorCode:   int(e),
			})
		}
	}
	return errorList, nil
}

func (g *GeTuiUser) Close() {

	// do nothing
}

type userResp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (g *GeTuiUser) decodeNestedResp(unknown interface{}) (map[string]string, error) {
	nested, ok := unknown.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unknown type %+v %T", unknown, unknown)
	}

	resp := make(map[string]string)
	if code, ok := nested["code"]; !ok {
		return nil, fmt.Errorf("no code found")
	} else {
		if ret, ok := code.(float64); !ok {
			return nil, fmt.Errorf("unknown type of code %+v %T", code, code)
		} else {
			if ret != 0 {
				log.Warnf("resp.Code %v resp.Msg %v", ret, nested["msg"])
				resp["available"] = "false"
				return resp, nil
			}
			resp["available"] = "true"
		}
	}
	data, ok := nested["data"]
	if !ok {
		return nil, fmt.Errorf("no data found")
	}
	detail, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unknown type of data %+v %T", data, data)
	}
	for k, v := range detail {
		desc, ok := v.(string)
		if !ok {
			log.Warnf("unknown type of v %+v %T", v, v)
			continue
		}
		resp[k] = desc
	}
	return resp, nil
}

func (g *GeTuiUser) url(path string) string {
	return g.apiUrl + path
}
