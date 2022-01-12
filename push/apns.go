package push

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	log "github.com/sirupsen/logrus"
)

type ApnsPush struct {
	apiUrl string
	topic  string
	client *http.Client
}

func NewApnsPush(env, topic string, hc *http.Client) (*ApnsPush, error) {
	api := apns2.HostDevelopment
	if env != "Production" {
		api = apns2.HostProduction
	}

	return &ApnsPush{
		apiUrl: api,
		topic:  topic,
		client: hc,
	}, nil
}

func (a *ApnsPush) PushSingleByCid(ctx context.Context, request *SingleReq, token string) (map[string]map[string]string, error) {
	notify, err := a.buildNotify(request)
	if err != nil {
		log.Errorf("build notify err %+v", err)
		return nil, err
	}

	content, err := json.Marshal(notify)
	if err != nil {
		log.Errorf("json marshal err %+v", err)
		return nil, err
	}

	url := fmt.Sprintf("%s/3/device/%s", a.apiUrl, notify.DeviceToken)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		log.Errorf("new http request err %+v", err)
		return nil, err
	}

	req.Header.Set("authorization", fmt.Sprintf("bearer %s", token))
	if notify.Topic != "" {
		req.Header.Set("apns-topic", notify.Topic)
	}
	if notify.ApnsID != "" {
		req.Header.Set("apns-id", notify.ApnsID)
	}
	if notify.CollapseID != "" {
		req.Header.Set("apns-collapse-id", notify.CollapseID)
	}
	if notify.Priority > 0 {
		req.Header.Set("apns-priority", fmt.Sprintf("%v", notify.Priority))
	}
	if !notify.Expiration.IsZero() {
		req.Header.Set("apns-expiration", fmt.Sprintf("%v", notify.Expiration.Unix()))
	}
	if notify.PushType != "" {
		req.Header.Set("apns-push-type", string(notify.PushType))
	} else {
		req.Header.Set("apns-push-type", string(apns2.PushTypeAlert))
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	req.WithContext(ctx)
	httpRes, err := a.client.Do(req)
	if err != nil {
		log.Errorf("http client do err %+v", err)
		return nil, err
	}
	defer httpRes.Body.Close()

	response := &apns2.Response{}
	response.StatusCode = httpRes.StatusCode
	response.ApnsID = httpRes.Header.Get("apns-id")

	err = json.NewDecoder(httpRes.Body).Decode(&response)
	if err != nil && err != io.EOF {
		log.Errorf("json decode err %+v", err)
		return nil, err
	}

	if !response.Sent() {
		log.Errorf("push result %+v", response)
		return nil, errors.New(response.Reason)
	}

	status := response.ApnsID
	if len(status) == 0 {
		status = notify.ApnsID
	}
	return map[string]map[string]string{
		"x": {
			notify.DeviceToken: status,
		},
	}, nil
}

func (a *ApnsPush) PushSingleByAlias(ctx context.Context, request *SingleReq, token string) (map[string]map[string]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) PushBatchByCid(ctx context.Context, request *BatchReq, token string) (map[string]map[string]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) PushBatchByAlias(ctx context.Context, request *BatchReq, token string) (map[string]map[string]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) CreateMsg(ctx context.Context, request *CreateReq, token string) (string, error) {
	return "", errors.New("unsupported for now")
}

func (a *ApnsPush) PushListByCid(ctx context.Context, request *ListReq, token string) (map[string]map[string]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) PushListByAlias(ctx context.Context, request *ListReq, token string) (map[string]map[string]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) PushAll(ctx context.Context, request *AllReq, token string) (string, error) {
	return "", errors.New("unsupported for now")
}

func (a *ApnsPush) PushByTag(ctx context.Context, request *ByTagReq, token string) (string, error) {
	return "", errors.New("unsupported for now")
}

func (a *ApnsPush) PushByFastCustomTag(ctx context.Context, request *ByTagReq, token string) (string, error) {
	return "", errors.New("unsupported for now")
}

func (a *ApnsPush) StopPush(ctx context.Context, taskId, token string) (bool, error) {
	return false, errors.New("unsupported for now")
}

func (a *ApnsPush) DeleteScheduleTask(ctx context.Context, taskId, token string) (bool, error) {
	return false, errors.New("unsupported for now")
}

func (a *ApnsPush) QueryScheduleTask(ctx context.Context, taskId, token string) (map[string]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) QueryDetail(ctx context.Context, taskId, cId, token string) ([][2]string, error) {
	return nil, errors.New("unsupported for now")
}

func (a *ApnsPush) Close() {
	// do nothing
}

func (a *ApnsPush) buildNotify(req *SingleReq) (*apns2.Notification, error) {
	if req.PushChannel == nil || req.PushChannel.Ios == nil || req.Audience == nil || len(req.Audience.Cid) != 1 {
		return nil, errors.New("invalid request")
	}

	ttl := time.Duration(req.Settings.GetTtl()) * time.Millisecond
	if ttl == 0 {
		ttl = time.Hour * 2
	}

	priority := apns2.PriorityHigh
	pushType := apns2.PushTypeAlert
	if req.PushChannel.Ios.GetAps().GetContentAvailable() == 1 && req.PushChannel.Ios.GetAps().Alert == nil {
		priority = apns2.PriorityLow
		pushType = apns2.PushTypeBackground
	}
	if req.PushChannel.Ios.Type == "voip" {
		pushType = apns2.PushTypeVOIP
	}

	content := payload.NewPayload()
	if alert := req.PushChannel.Ios.GetAps().GetAlert(); alert != nil {
		content.Alert(alert)
	}
	if badge := req.PushChannel.Ios.GetAutoBadge(); len(badge) != 0 {
		if !strings.ContainsAny(badge, "+-") {
			num, err := strconv.ParseInt(badge, 10, 64)
			if err == nil {
				content.Badge(int(num))
			}
		} else {
			// TODO: +- is not that easy
			content.UnsetBadge()
		}
	}
	if sound := req.PushChannel.Ios.GetAps().GetSound(); len(sound) != 0 {
		content.Sound(sound)
	}
	if ext := req.PushChannel.Ios.GetPayload(); len(ext) != 0 {
		content.Custom("ext", ext)
	}

	return &apns2.Notification{
		ApnsID:      uuid.NewString(),
		CollapseID:  req.PushChannel.Ios.GetApnsCollapseId(),
		DeviceToken: req.Audience.Cid[0],
		Topic:       a.topic,
		Expiration:  time.Now().Add(ttl),
		Priority:    priority,
		Payload:     content,
		PushType:    pushType,
	}, nil
}
