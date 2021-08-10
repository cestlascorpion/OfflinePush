package core

import (
	"context"
	"sync"
	"time"

	"github.com/cestlascorpion/push/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type AuthCache struct {
	Client proto.AuthClient
	Cache  map[UniqueId]*AuthToken
	Mutex  sync.RWMutex
}

func NewAuthCache() (*AuthCache, error) {
	conn, err := grpc.Dial(AuthServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Errorf("grpc dial err %+v", err)
		return nil, err
	}
	return &AuthCache{
		Client: proto.NewAuthClient(conn),
		Cache:  make(map[UniqueId]*AuthToken),
	}, nil
}

func (c *AuthCache) ReadAuth(uniqueId UniqueId) *AuthToken {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	token, ok := c.Cache[uniqueId]
	if !ok {
		return nil
	}
	return token
}

func (c *AuthCache) WriteAuth(uniqueId UniqueId, auth *AuthToken) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Cache[uniqueId] = auth
}

func (c *AuthCache) GetAuth(uniqueId UniqueId, oldToken string) (*AuthToken, error) {
	auth := c.ReadAuth(uniqueId)
	if auth != nil && auth.Token != oldToken && auth.ExpireAt > time.Now().Unix()+30 {
		return auth, nil
	}

	resp, err := c.Client.GetToken(context.Background(), &proto.GetTokenReq{
		PushAgent: uniqueId.PushAgent,
		BundleId:  uniqueId.BundleId,
		OldToken:  oldToken,
	})
	if err != nil {
		log.Errorf("client get token er %+v", err)
		return nil, err
	}
	newAuth := &AuthToken{
		Token:    resp.Token,
		ExpireAt: resp.ExpireAt,
	}
	c.WriteAuth(uniqueId, newAuth)
	return newAuth, nil
}

func (c *AuthCache) Start(ctx context.Context) error {
	go func() {
		ticker := time.NewTicker(time.Second * 120)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Infof("Cache Watcher exiting...")
				return
			case <-ticker.C:
				err := c.Check()
				log.Debugf("Cache Watcher check...")
				if err != nil {
					log.Errorf("Cache Watcher failed err %+v", err)
				}
			}
		}
	}()
	return nil
}

func (c *AuthCache) Check() error {
	expired := make(map[UniqueId]*AuthToken)

	{
		c.Mutex.RLock()
		for id, auth := range c.Cache {
			if auth.ExpireAt < time.Now().Unix()+60 {
				expired[id] = auth
			}
		}
		c.Mutex.RUnlock()
	}

	if len(expired) == 0 {
		return nil
	}
	log.Infof("expire in next %d sec, expired %+v", 30, expired)

	for id, auth := range expired {
		var invalidToken string
		if auth != nil {
			invalidToken = auth.Token
		}
		resp, err := c.Client.GetToken(context.Background(), &proto.GetTokenReq{
			PushAgent: id.PushAgent,
			BundleId:  id.BundleId,
			OldToken:  invalidToken,
		})
		if err != nil {
			log.Errorf("client get token failed, push agent %s bundle id %s err %+v", id.PushAgent, id.BundleId, err)
			continue
		}
		expired[id] = &AuthToken{
			Token:    resp.Token,
			ExpireAt: resp.ExpireAt,
		}
	}

	{
		c.Mutex.Lock()
		for key, value := range expired {
			if value != nil {
				c.Cache[key] = value
			}
		}
		c.Mutex.Unlock()
	}

	return nil
}
