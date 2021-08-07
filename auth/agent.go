package auth

import (
	"fmt"

	. "github.com/cestlascorpion/push/core"
)

type Agent interface {
	GetAuth() (token string, expireAt int64, err error)
	DelAuth(token string) (err error)
}

type AgentMgr struct {
	Agents map[UniqueId]Agent
}

func NewAgentMgr() (*AgentMgr, error) {
	return &AgentMgr{
		Agents: make(map[UniqueId]Agent),
	}, nil
}

func (m *AgentMgr) RegisterAgent(uniqueId UniqueId, agent Agent) error {
	m.Agents[uniqueId] = agent
	return nil
}

func (m *AgentMgr) GetAuth(uniqueId UniqueId) (*AuthToken, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	token, expireAt, err := agent.GetAuth()
	if err != nil {
		return nil, err
	}
	return &AuthToken{Token: token, ExpireAt: expireAt}, nil
}

func (m *AgentMgr) DelAuth(uniqueId UniqueId, token string) error {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.DelAuth(token)
}
