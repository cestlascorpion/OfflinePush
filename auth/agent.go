package auth

import (
	"fmt"

	"github.com/cestlascorpion/offlinepush/core"
)

type Agent interface {
	GetAuth() (auth *core.AuthToken, err error)
	DelAuth(token string) (err error)
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

func (m *AgentMgr) GetAuth(uniqueId core.UniqueId) (*core.AuthToken, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetAuth()
}

func (m *AgentMgr) DelAuth(uniqueId core.UniqueId, token string) error {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.DelAuth(token)
}

func (m *AgentMgr) Close() {
	for _, agent := range m.agents {
		agent.Close()
	}
}
