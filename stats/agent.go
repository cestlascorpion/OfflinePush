package stats

import (
	"context"
	"fmt"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
)

type Agent interface {
	GetTasks(ctx context.Context, taskIds []string, token string) (map[string]*proto.BaseStatics, error)
	GetTaskGroup(ctx context.Context, group, token string) (map[string]*proto.BaseStatics, error)
	GetPushCount(ctx context.Context, token string) ([]*proto.BasePushCount, error)
	GetPushDataByDay(ctx context.Context, date time.Time, token string) (map[string]*proto.BaseStatics, error)
	GetUserDataByDay(ctx context.Context, date time.Time, token string) (map[string]map[string]int32, error)
	GetOnlineUserBy24H(ctx context.Context, token string) (map[int64]int32, error)
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

func (m *AgentMgr) GetTasks(ctx context.Context, uniqueId core.UniqueId, taskIds []string, token string) (map[string]*proto.BaseStatics, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetTasks(ctx, taskIds, token)
}

func (m *AgentMgr) GetTaskGroup(ctx context.Context, uniqueId core.UniqueId, group, token string) (map[string]*proto.BaseStatics, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetTaskGroup(ctx, group, token)
}

func (m *AgentMgr) GetPushCount(ctx context.Context, uniqueId core.UniqueId, token string) ([]*proto.BasePushCount, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetPushCount(ctx, token)
}

func (m *AgentMgr) GetPushDataByDay(ctx context.Context, uniqueId core.UniqueId, date time.Time, token string) (map[string]*proto.BaseStatics, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetPushDataByDay(ctx, date, token)
}

func (m *AgentMgr) GetUserDataByDay(ctx context.Context, uniqueId core.UniqueId, date time.Time, token string) (map[string]map[string]int32, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetUserDataByDay(ctx, date, token)
}

func (m *AgentMgr) GetOnlineUserBy24H(ctx context.Context, uniqueId core.UniqueId, token string) (map[int64]int32, error) {
	agent, ok := m.agents[uniqueId]
	if !ok {
		return nil, fmt.Errorf("unsupported uniqueId %s", uniqueId)
	}
	return agent.GetOnlineUserBy24H(ctx, token)
}

func (m *AgentMgr) Close() {
	for _, agent := range m.agents {
		agent.Close()
	}
}
