package user

import (
	"fmt"

	"github.com/cestlascorpion/offlinepush/core"
)

type Agent interface {
	BindAlias(list *AliasList, token string) error
	QueryAliasByCid(cid string, token string) (string, error)
	QueryCidByAlias(alias string, token string) ([]string, error)
	UnbindAlias(list *AliasList, token string) error
	RevokeAlias(alias string, token string) error
	BindUserWithTag(cid string, list *CustomTagList, token string) error
	BindTagWithUser(tag string, list *CidList, token string) (map[string]bool, error)
	UnbindTagFromUser(tag string, list *CidList, token string) (map[string]bool, error)
	QueryUserTag(cid string, token string) ([]string, error)
	AddBlackList(cidList []string, token string) error
	DelBlackList(cidList []string, token string) error
	QueryUserStatus(cidList []string, token string) (map[string]map[string]string, error)
	QueryDeviceStatus(cidList []string, token string) (map[string]map[string]string, error)
	QueryUserInfo(cidList []string, token string) ([]string, map[string]map[string]string, error)
	SetPushBadge(cidList []string, op *Operation, token string) error
	QueryUserCount(list *ComplexTagList, token string) (int, error)
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

func (m *AgentMgr) BindAlias(uniqueId core.UniqueId, list *AliasList, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("BindAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.BindAlias(list, token)
}

func (m *AgentMgr) QueryAliasByCid(uniqueId core.UniqueId, cid string, token string) (string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return "", fmt.Errorf("QueryAliasByCid() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryAliasByCid(cid, token)
}

func (m *AgentMgr) QueryCidByAlias(uniqueId core.UniqueId, alias string, token string) ([]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryCidByAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryCidByAlias(alias, token)
}

func (m *AgentMgr) UnbindAlias(uniqueId core.UniqueId, list *AliasList, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("UnbindAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.UnbindAlias(list, token)
}

func (m *AgentMgr) RevokeAlias(uniqueId core.UniqueId, alias string, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("RevokeAlias() unsupported uniqueId %s", uniqueId)
	}
	return agent.RevokeAlias(alias, token)
}

func (m *AgentMgr) BindUserWithTag(uniqueId core.UniqueId, cid string, list *CustomTagList, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("BindUserWithTag() unsupported uniqueId %s", uniqueId)
	}
	return agent.BindUserWithTag(cid, list, token)
}

func (m *AgentMgr) BindTagWithUser(uniqueId core.UniqueId, tag string, list *CidList, token string) (map[string]bool, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("BindTagWithUser() unsupported uniqueId %s", uniqueId)
	}
	return agent.BindTagWithUser(tag, list, token)
}

func (m *AgentMgr) UnbindTagFromUser(uniqueId core.UniqueId, tag string, list *CidList, token string) (map[string]bool, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("UnbindTagFromUser() unsupported uniqueId %s", uniqueId)
	}
	return agent.UnbindTagFromUser(tag, list, token)
}

func (m *AgentMgr) QueryUserTag(uniqueId core.UniqueId, cid string, token string) ([]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryUserTag() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserTag(cid, token)
}

func (m *AgentMgr) AddBlackList(uniqueId core.UniqueId, cidList []string, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("AddBlackList() unsupported uniqueId %s", uniqueId)
	}
	return agent.AddBlackList(cidList, token)
}

func (m *AgentMgr) DelBlackList(uniqueId core.UniqueId, cidList []string, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("DelBlackList() unsupported uniqueId %s", uniqueId)
	}
	return agent.DelBlackList(cidList, token)
}

func (m *AgentMgr) QueryUserStatus(uniqueId core.UniqueId, cidList []string, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryUserStatus() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserStatus(cidList, token)
}

func (m *AgentMgr) QueryDeviceStatus(uniqueId core.UniqueId, cidList []string, token string) (map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("QueryDeviceStatus() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryDeviceStatus(cidList, token)
}

func (m *AgentMgr) QueryUserInfo(uniqueId core.UniqueId, cidList []string, token string) ([]string, map[string]map[string]string, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, nil, fmt.Errorf("QueryUserInfo() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserInfo(cidList, token)
}

func (m *AgentMgr) SetPushBadge(uniqueId core.UniqueId, cidList []string, op *Operation, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("SetPushBadge() unsupported uniqueId %s", uniqueId)
	}
	return agent.SetPushBadge(cidList, op, token)
}

func (m *AgentMgr) QueryUserCount(uniqueId core.UniqueId, list *ComplexTagList, token string) (int, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return 0, fmt.Errorf("QueryUserCount() unsupported uniqueId %s", uniqueId)
	}
	return agent.QueryUserCount(list, token)
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
