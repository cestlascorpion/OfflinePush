package push

import (
	"context"
	"fmt"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
)

type Agent interface {
	PushSingleByCid(ctx context.Context, request *SingleReq, token string) (map[string]map[string]string, error)
	PushSingleByAlias(ctx context.Context, request *SingleReq, token string) (map[string]map[string]string, error)
	PushBatchByCid(ctx context.Context, request *BatchReq, token string) (map[string]map[string]string, error)
	PushBatchByAlias(ctx context.Context, request *BatchReq, token string) (map[string]map[string]string, error)
	CreateMsg(ctx context.Context, request *CreateReq, token string) (string, error)
	PushListByCid(ctx context.Context, request *ListReq, token string) (map[string]map[string]string, error)
	PushListByAlias(ctx context.Context, request *ListReq, token string) (map[string]map[string]string, error)
	PushAll(ctx context.Context, request *AllReq, token string) (string, error)
	PushByTag(ctx context.Context, request *ByTagReq, token string) (string, error)
	PushByFastCustomTag(ctx context.Context, request *ByTagReq, token string) (string, error)
	StopPush(ctx context.Context, taskId, token string) (bool, error)
	DeleteScheduleTask(ctx context.Context, taskId, token string) (bool, error)
	QueryScheduleTask(ctx context.Context, taskId, token string) (map[string]string, error)
	QueryDetail(ctx context.Context, taskId, cId, token string) ([][2]string, error)
	Close()
}

type AgentMgr struct {
	agents map[core.UniqueId]Agent
}

func NewAgentMgr() (*AgentMgr, error) {
	return &AgentMgr{
		agents: make(map[core.UniqueId]Agent),
	}, nil
}

func (m *AgentMgr) RegisterAgent(uniqueId core.UniqueId, agent Agent) error {
	m.agents[uniqueId] = agent
	return nil
}

func (m *AgentMgr) PushSingle(ctx context.Context, uniqueId core.UniqueId, request *SingleReq, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.Audience.Cid) != 0 {
		return agent.PushSingleByCid(ctx, request, token)
	}
	return agent.PushSingleByAlias(ctx, request, token)
}

func (m *AgentMgr) PushBatch(ctx context.Context, uniqueId core.UniqueId, request *BatchReq, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.MsgList[0].Audience.Cid) != 0 {
		return agent.PushBatchByCid(ctx, request, token)
	}
	return agent.PushBatchByAlias(ctx, request, token)
}

func (m *AgentMgr) CreateMsg(ctx context.Context, uniqueId core.UniqueId, request *CreateReq, token string) (string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.CreateMsg(ctx, request, token)
}

func (m *AgentMgr) PushList(ctx context.Context, uniqueId core.UniqueId, request *ListReq, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.Audience.Cid) != 0 {
		return agent.PushListByCid(ctx, request, token)
	}
	return agent.PushListByAlias(ctx, request, token)
}

func (m *AgentMgr) PushAll(ctx context.Context, uniqueId core.UniqueId, request *AllReq, token string) (string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.PushAll(ctx, request, token)
}

func (m *AgentMgr) PushByTag(ctx context.Context, uniqueId core.UniqueId, request *ByTagReq, token string) (string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	if len(request.Audience.FastCustomTag) != 0 {
		return agent.PushByFastCustomTag(ctx, request, token)
	}
	return agent.PushByTag(ctx, request, token)
}

func (m *AgentMgr) StopPush(ctx context.Context, uniqueId core.UniqueId, taskId, token string) (bool, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return false, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.StopPush(ctx, taskId, token)
}

func (m *AgentMgr) DeleteScheduleTask(ctx context.Context, uniqueId core.UniqueId, taskId, token string) (bool, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return false, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.DeleteScheduleTask(ctx, taskId, token)
}

func (m *AgentMgr) QueryScheduleTask(ctx context.Context, uniqueId core.UniqueId, taskId, token string) (map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryScheduleTask(ctx, taskId, token)
}

func (m *AgentMgr) QueryDetail(ctx context.Context, uniqueId core.UniqueId, taskId, cId, token string) ([][2]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryDetail(ctx, taskId, cId, token)
}

func (m *AgentMgr) Close() {
	for _, agent := range m.agents {
		agent.Close()
	}
}

type SingleReq struct {
	RequestId   string             `json:"request_id"`             // ???????????????????????????????????????10-32??????????????????request_id??????????????????????????????
	Audience    *proto.Audience    `json:"audience"`               // ???????????????cid????????????????????????cid
	Settings    *proto.Settings    `json:"settings,omitempty"`     // ??????????????????????????????
	PushMessage *proto.PushMessage `json:"push_message"`           // ???????????????????????????????????????
	PushChannel *proto.PushChannel `json:"push_channel,omitempty"` // ?????????????????????????????????????????????ios???????????????android??????????????????
}

type BatchReq struct {
	IsAsync bool         `json:"is_async,omitempty"` // ????????????????????????false????????????????????????????????????????????????data???is_async???false?????????data
	MsgList []*SingleReq `json:"msg_list"`           // ??????????????????????????????????????????????????????????????? 200
}

type CreateReq struct {
	RequestId   string             `json:"request_id,omitempty"`   // ????????????????????????????????????10-32??????????????????request_id??????????????????????????????
	GroupName   string             `json:"group_name,omitempty"`   // ????????????????????????
	Settings    *proto.Settings    `json:"settings,omitempty"`     // ??????????????????????????????
	PushMessage *proto.PushMessage `json:"push_message"`           // ???????????????????????????????????????
	PushChannel *proto.PushChannel `json:"push_channel,omitempty"` // ?????????????????????????????????????????????ios???????????????android??????????????????
}

type ListReq struct {
	Audience *proto.Audience `json:"audience"`           // ??????????????????cid???????????????cid??????????????????????????????????????????200
	IsAsync  bool            `json:"is_async,omitempty"` // ????????????????????????false????????????????????????????????????????????????data???is_async???false?????????data
	TaskId   string          `json:"taskid"`             // ??????????????????????????????????????????????????????????????????taskId?????????????????????
}

type AllReq struct {
	RequestId   string             `json:"request_id"`             // ?????????????????????????????????10-32??????????????????request_id??????????????????????????????
	GroupName   string             `json:"group_name,omitempty"`   // ????????????????????????
	Audience    string             `json:"audience"`               // ????????????????????????all
	Settings    *proto.Settings    `json:"settings,omitempty"`     // ??????????????????????????????
	PushMessage *proto.PushMessage `json:"push_message"`           // ???????????????????????????????????????
	PushChannel *proto.PushChannel `json:"push_channel,omitempty"` // ?????????????????????????????????????????????ios???????????????android??????????????????
}

type ByTagReq struct {
	RequestId   string             `json:"request_id"`             // ?????????????????????????????????10-32??????????????????request_id??????????????????????????????
	GroupName   string             `json:"group_name,omitempty"`   // ????????????????????????
	Settings    *proto.Settings    `json:"settings,omitempty"`     // ??????????????????????????????
	Audience    *proto.Audience    `json:"audience"`               // ???????????????tag??????
	PushMessage *proto.PushMessage `json:"push_message"`           // ???????????????????????????????????????
	PushChannel *proto.PushChannel `json:"push_channel,omitempty"` // ?????????????????????????????????????????????ios???????????????android??????????????????
}
