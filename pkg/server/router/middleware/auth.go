// Package middleware
/**
Copyright (c) The Authors.
* @Author: feng.xiang
* @Date: 2023/11/2 19:49
* @Desc:
*/
package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/spf13/viper"

	"app/pkg/conf"
	"app/pkg/logger"
	"app/pkg/server/auth/acl"
)

var (
	basicAuthDb = sync.Map{}
)

func init() {
	basicAuthDb.Store("admin", "admin")
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//session := map[string]any{}
		var err error
		var ok bool
		if conf.Cfg.Auth.Acl != nil {
			err, ok = authAcl(c)
		} else {
			// TODO: add your auth here
			ok = true
		}
		if !ok {
			logger.L.Warn(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorized refused",
			})
			return
		}
		c.Next()
		return
	}
}

func authAcl(ctx *gin.Context) (error, bool) {
	session := &acl.Session{}
	sess, err := ctx.Cookie("session")
	if err == nil {
		s := NewSignature(conf.Cfg.SecretKey, "cookie-session", "", "hmac", nil, nil)
		content, err := s.Unsign(sess)
		if err != nil {
			return err, false
		}
		err = json.Unmarshal(content, &session)
		if err != nil {
			return err, false
		}
		ctx.Set("session", session)
		return nil, true
	}
	return fmt.Errorf("no session"), false
}

func authBasic(ctx *gin.Context) (error, bool) {
	if user, password, ok := ctx.Request.BasicAuth(); ok {
		if p, ok := basicAuthDb.Load(user); ok && p.(string) == password {
			return nil, true
		} else {
			return fmt.Errorf("invalid user or password"), false
		}
	}
	return fmt.Errorf("invalid user or password"), false
}

func authWithWhiteList(ip string) bool {
	return lo.Contains(viper.GetStringSlice("gateway.whiteList"), ip)
}
