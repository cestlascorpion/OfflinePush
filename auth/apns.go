package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"sync"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type ApnsAuth struct {
	key    *ecdsa.PrivateKey
	keyId  string
	teamId string
	cache  *core.AuthToken
	cancel context.CancelFunc
	mutex  sync.RWMutex
}

func NewApnsAuth(privateKey []byte, keyId, teamId string) (*ApnsAuth, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		log.Errorf("pem.Decode failed")
		return nil, errors.New("invalid pem")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Errorf("x509 parse err %+v", err)
		return nil, err
	}

	pk, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		log.Errorf("unexpected key type %T", key)
		return nil, err
	}

	token, err := genJWT(pk, keyId, teamId)
	if err != nil {
		log.Errorf("generate jwt err %+v err", err)
		return nil, err
	}

	auth := &ApnsAuth{
		key:    pk,
		keyId:  keyId,
		teamId: teamId,
		cache:  token,
		cancel: nil,
	}

	auth.cancel = auth.watch()
	return auth, nil
}

func (a *ApnsAuth) GetAuth() (*core.AuthToken, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	auth := &core.AuthToken{
		Token:    a.cache.Token,
		ExpireAt: a.cache.ExpireAt,
	}
	return auth, nil
}

func (a *ApnsAuth) DelAuth(token string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if token != a.cache.Token {
		return nil
	}

	auth, err := genJWT(a.key, a.keyId, a.teamId)
	if err != nil {
		log.Errorf("generate jwt err %+v", err)
		return err
	}

	a.cache.Token = auth.Token
	a.cache.ExpireAt = auth.ExpireAt
	return nil
}

func (a *ApnsAuth) Close() {
	a.cancel()
}

func (a *ApnsAuth) watch() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.TODO())

	go func() {
		ticker := time.NewTicker(time.Minute * 30)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				token, err := genJWT(a.key, a.keyId, a.teamId)
				if err != nil {
					log.Warnf("watch get auth err %+v err", err)
					continue
				}

				a.mutex.Lock()
				a.cache.Token = token.Token
				a.cache.ExpireAt = token.ExpireAt
				a.mutex.Unlock()
			}
		}
	}()

	return cancel
}

func genJWT(pk *ecdsa.PrivateKey, keyId, teamId string) (*core.AuthToken, error) {
	ts := time.Now()
	token := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": keyId,
		},
		Claims: jwt.MapClaims{
			"iss": teamId,
			"iat": ts.Unix(),
		},
		Method: jwt.SigningMethodES256,
	}

	signedToken, err := token.SignedString(pk)
	if err != nil {
		log.Errorf("jwt sign with pk err %+v", err)
		return nil, err
	}

	return &core.AuthToken{
		Token:    signedToken,
		ExpireAt: ts.Add(time.Hour).Unix(),
	}, nil
}
