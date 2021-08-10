package stats

import (
	"fmt"
	"time"

	. "github.com/cestlascorpion/push/core"
	"github.com/cestlascorpion/push/proto"
)

type Agent interface {
	GetTasks(taskIds []string, token string) (map[string]*proto.BaseStatics, error)
	GetTaskGroup(group, token string) (map[string]*proto.BaseStatics, error)
	GetPushCount(token string) ([]*proto.BasePushCount, error)
	GetPushDataByDay(date time.Time, token string) (map[string]*proto.BaseStatics, error)
	GetUserDataByDay(date time.Time, token string) (map[string]map[string]int32, error)
	GetOnlineUserBy24H(token string) (map[int64]int32, error)
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

func (m *AgentMgr) GetTasks(uniqueId UniqueId, taskIds []string, token string) (map[string]*proto.BaseStatics, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetTasks(taskIds, token)
}

func (m *AgentMgr) GetTaskGroup(uniqueId UniqueId, group, token string) (map[string]*proto.BaseStatics, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetTaskGroup(group, token)
}

func (m *AgentMgr) GetPushCount(uniqueId UniqueId, token string) ([]*proto.BasePushCount, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetPushCount(token)
}

func (m *AgentMgr) GetPushDataByDay(uniqueId UniqueId, date time.Time, token string) (map[string]*proto.BaseStatics, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetPushDataByDay(date, token)
}

func (m *AgentMgr) GetUserDataByDay(uniqueId UniqueId, date time.Time, token string) (map[string]map[string]int32, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetUserDataByDay(date, token)
}

func (m *AgentMgr) GetOnlineUserBy24H(uniqueId UniqueId, token string) (map[int64]int32, error) {
	agent, ok := m.Agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetOnlineUserBy24H(token)
}
