package core

import (
	"context"
	"sync"
	"time"

	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type AuthCache struct {
	client proto.AuthClient
	cache  map[UniqueId]*AuthToken
	cancel context.CancelFunc
	mutex  sync.RWMutex
}

func NewAuthCache() (*AuthCache, error) {
	conn, err := grpc.Dial(AuthServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Errorf("grpc dial err %+v", err)
		return nil, err
	}
	cache := &AuthCache{
		client: proto.NewAuthClient(conn),
		cache:  make(map[UniqueId]*AuthToken),
		cancel: nil,
	}

	cache.cancel = cache.watch()
	return cache, nil
}

func (c *AuthCache) GetAuth(uniqueId UniqueId) (*AuthToken, error) {
	auth := c.readAuth(uniqueId)
	if auth != nil && auth.ExpireAt > time.Now().Add(time.Minute).Unix() {
		return auth, nil
	}

	resp, err := c.client.GetToken(context.Background(), &proto.GetTokenReq{
		PushAgent: uniqueId.PushAgent,
		BundleId:  uniqueId.BundleId,
	})
	if err != nil {
		log.Errorf("client get token er %+v", err)
		return nil, err
	}
	newAuth := &AuthToken{
		Token:    resp.Token,
		ExpireAt: resp.ExpireAt,
	}
	c.writeAuth(uniqueId, newAuth)
	return newAuth, nil
}

func (c *AuthCache) Close() {
	c.Close()
}

func (c *AuthCache) readAuth(uniqueId UniqueId) *AuthToken {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	token, ok := c.cache[uniqueId]
	if !ok {
		return nil
	}
	return token
}

func (c *AuthCache) writeAuth(uniqueId UniqueId, auth *AuthToken) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[uniqueId] = auth
}

func (c *AuthCache) watch() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.TODO())

	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Infof("cache Watcher exiting...")
				return
			case <-ticker.C:
				err := c.check()
				log.Debugf("cache Watcher check...")
				if err != nil {
					log.Errorf("cache Watcher failed err %+v", err)
				}
			}
		}
	}()

	return cancel
}

func (c *AuthCache) check() error {
	expired := make(map[UniqueId]*AuthToken)

	{
		c.mutex.RLock()
		for id, auth := range c.cache {
			if auth.ExpireAt < time.Now().Add(time.Minute).Unix() {
				expired[id] = auth
			}
		}
		c.mutex.RUnlock()
	}

	if len(expired) == 0 {
		return nil
	}
	log.Infof("expire in next %d sec, expired %+v", 30, expired)

	for id := range expired {
		resp, err := c.client.GetToken(context.Background(), &proto.GetTokenReq{
			PushAgent: id.PushAgent,
			BundleId:  id.BundleId,
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
		c.mutex.Lock()
		for key, value := range expired {
			if value != nil {
				c.cache[key] = value
			}
		}
		c.mutex.Unlock()
	}

	return nil
}
