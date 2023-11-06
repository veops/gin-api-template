// Package acl
package acl

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/spf13/cast"

	"app/pkg/conf"
)

func GetSessionFromCtx(ctx *gin.Context) (res *Session) {
	res, _ = ctx.Value("session").(*Session)
	return
}

func HasPerm(resourceId int32, rid int32, action string) bool {
	mapping, err := GetResourcePermissions(context.Background(), resourceId)
	if err != nil {
		return false
	}
	for _, v := range mapping {
		if lo.ContainsBy(v.Perms, func(p *Perm) bool { return p.Rid == rid && p.Name == action }) {
			return true
		}
	}
	return false
}

func IsAdmin(session *Session) bool {
	for _, pr := range session.Acl.ParentRoles {
		if pr == "admin" || pr == "acl_admin" {
			return true
		}
	}
	return false
}

func resourceTypeName(resourceType string) string {
	names := conf.Cfg.Auth.Acl.ResourceNames
	for _, v := range names {
		if v.Key == resourceType {
			return v.Value
		}
	}
	return "NONE"
}

func CreateGrantAcl(ctx context.Context, session *Session, resourceType string, now time.Time) (resourceId int32, err error) {
	s := rand.New(rand.NewSource(time.Now().Unix())).Int31n(1000)
	resourceName := strings.Join([]string{resourceType, now.Format(time.DateTime), cast.ToString(s)}, "-")
	resource, err := AddResource(ctx,
		session.GetUid(),
		resourceTypeName(resourceType),
		resourceName)
	if err != nil {
		return
	}
	if err = GrantRoleResource(ctx, session.Acl.Rid, resource.ResourceId, AllPermissions); err != nil {
		return
	}
	resourceId = resource.ResourceId
	return
}
