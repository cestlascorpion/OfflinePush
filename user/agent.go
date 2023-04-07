package user

import (
	"context"
	"fmt"

	"github.com/cestlascorpion/offlinepush/core"
)

type Agent interface {
	BindAlias(ctx context.Context, list *AliasList, token string) error
	QueryAliasByCid(ctx context.Context, cid string, token string) (string, error)
	QueryCidByAlias(ctx context.Context, alias string, token string) ([]string, error)
	UnbindAlias(ctx context.Context, list *AliasList, token string) error
	RevokeAlias(ctx context.Context, alias string, token string) error
	BindUserWithTag(ctx context.Context, cid string, list *CustomTagList, token string) error
	BindTagWithUser(ctx context.Context, tag string, list *CidList, token string) (map[string]bool, error)
	UnbindTagFromUser(ctx context.Context, tag string, list *CidList, token string) (map[string]bool, error)
	QueryUserTag(ctx context.Context, cid string, token string) ([]string, error)
	AddBlackList(ctx context.Context, cidList []string, token string) error
	DelBlackList(ctx context.Context, cidList []string, token string) error
	QueryUserStatus(ctx context.Context, cidList []string, token string) (map[string]map[string]string, error)
	QueryDeviceStatus(ctx context.Context, cidList []string, token string) (map[string]map[string]string, error)
	QueryUserInfo(ctx context.Context, cidList []string, token string) ([]string, map[string]map[string]string, error)
	SetPushBadge(ctx context.Context, cidList []string, op *Operation, token string) error
	QueryUserCount(ctx context.Context, list *ComplexTagList, token string) (int, error)
	ManageCidAndDeviceToken(ctx context.Context, manufacturer string, dtList *CidAndDeviceTokenList, token string) ([]*DTError, error)
	Close()
}

type AgentMgr struct {
	agents map[core.UniqueId]Agent
}

func NewUserMgr() (*AgentMgr, error) {
	return &AgentMgr{
		agents: make(map[core.UniqueId]Agent),
	}, nil
}

func (m *AgentMgr) RegisterAgent(uniqueId core.UniqueId, agent Agent) error {
	m.agents[uniqueId] = agent
	return nil
}

func (m *AgentMgr) BindAlias(ctx context.Context, uniqueId core.UniqueId, list *AliasList, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("BindAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.BindAlias(ctx, list, token)
}

func (m *AgentMgr) QueryAliasByCid(ctx context.Context, uniqueId core.UniqueId, cid string, token string) (string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("QueryAliasByCid() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryAliasByCid(ctx, cid, token)
}

func (m *AgentMgr) QueryCidByAlias(ctx context.Context, uniqueId core.UniqueId, alias string, token string) ([]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryCidByAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryCidByAlias(ctx, alias, token)
}

func (m *AgentMgr) UnbindAlias(ctx context.Context, uniqueId core.UniqueId, list *AliasList, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("UnbindAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.UnbindAlias(ctx, list, token)
}

func (m *AgentMgr) RevokeAlias(ctx context.Context, uniqueId core.UniqueId, alias string, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("RevokeAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.RevokeAlias(ctx, alias, token)
}

func (m *AgentMgr) BindUserWithTag(ctx context.Context, uniqueId core.UniqueId, cid string, list *CustomTagList, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("BindUserWithTag() unsupported uniqueId %s", uniqueId)
	}
	return agent.BindUserWithTag(ctx, cid, list, token)
}

func (m *AgentMgr) BindTagWithUser(ctx context.Context, uniqueId core.UniqueId, tag string, list *CidList, token string) (map[string]bool, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("BindTagWithUser() unsupported uniqueId %s", uniqueId)
	}
	return agent.BindTagWithUser(ctx, tag, list, token)
}

func (m *AgentMgr) UnbindTagFromUser(ctx context.Context, uniqueId core.UniqueId, tag string, list *CidList, token string) (map[string]bool, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("UnbindTagFromUser() unsupported uniqueId %s", uniqueId)
	}
	return agent.UnbindTagFromUser(ctx, tag, list, token)
}

func (m *AgentMgr) QueryUserTag(ctx context.Context, uniqueId core.UniqueId, cid string, token string) ([]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryUserTag() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserTag(ctx, cid, token)
}

func (m *AgentMgr) AddBlackList(ctx context.Context, uniqueId core.UniqueId, cidList []string, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("AddBlackList() unsupported uniqueId %s", uniqueId)
	}
	return agent.AddBlackList(ctx, cidList, token)
}

func (m *AgentMgr) DelBlackList(ctx context.Context, uniqueId core.UniqueId, cidList []string, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("DelBlackList() unsupported uniqueId %s", uniqueId)
	}
	return agent.DelBlackList(ctx, cidList, token)
}

func (m *AgentMgr) QueryUserStatus(ctx context.Context, uniqueId core.UniqueId, cidList []string, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryUserStatus() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserStatus(ctx, cidList, token)
}

func (m *AgentMgr) QueryDeviceStatus(ctx context.Context, uniqueId core.UniqueId, cidList []string, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryDeviceStatus() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryDeviceStatus(ctx, cidList, token)
}

func (m *AgentMgr) QueryUserInfo(ctx context.Context, uniqueId core.UniqueId, cidList []string, token string) ([]string, map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, nil, fmt.Errorf("QueryUserInfo() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserInfo(ctx, cidList, token)
}

func (m *AgentMgr) SetPushBadge(ctx context.Context, uniqueId core.UniqueId, cidList []string, op *Operation, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("SetPushBadge() unsupported uniqueId %s", uniqueId)
	}
	return agent.SetPushBadge(ctx, cidList, op, token)
}

func (m *AgentMgr) QueryUserCount(ctx context.Context, uniqueId core.UniqueId, list *ComplexTagList, token string) (int, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return 0, fmt.Errorf("QueryUserCount() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserCount(ctx, list, token)
}

func (m *AgentMgr) ManageCidAndDeviceToken(ctx context.Context, uniqueId core.UniqueId, manufacturer string, dtList *CidAndDeviceTokenList, token string) ([]*DTError, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("ManageCidAndDeviceToken() unsupported uniqueId %s", uniqueId)
	}
	return agent.ManageCidAndDeviceToken(ctx, manufacturer, dtList, token)
}

func (m *AgentMgr) Close() {
	for _, agent := range m.agents {
		agent.Close()
	}
}

type DataList struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}

type AliasList struct {
	DataList []*DataList `json:"data_list"`
}

type CustomTagList struct {
	TagList []string `json:"custom_tag"`
}

type CidList struct {
	CidList []string `json:"cid"`
}

type Operation struct {
	Badge string `json:"badge"`
}

type Tag struct {
	Key     string   `json:"key"`
	Values  []string `json:"values"`
	OptType string   `json:"opt_type"`
}

type ComplexTagList struct {
	Tag []*Tag `json:"tag"`
}

type DT struct {
	Cid         string `json:"cid"`
	DeviceToken string `json:"device_token"`
}

type CidAndDeviceTokenList struct {
	DTList []*DT `json:"dt_list"`
}

type DTError struct {
	Cid         string `json:"cid"`
	DeviceToken string `json:"device_token"`
	ErrorCode   int    `json:"error_code"`
}
