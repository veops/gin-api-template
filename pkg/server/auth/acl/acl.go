package acl

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	Write  = "write"
	Delete = "delete"
	Read   = "read"
	Grant  = "grant"
)

var (
	AllPermissions = []string{Write, Delete, Read, Grant}
)

type Role struct {
	Permissions []string `json:"permissions"`
}

type ResourceResult struct {
	Groups    []any       `json:"groups"`
	Resources []*Resource `json:"resources"`
}
type Resource struct {
	AppID          int      `json:"app_id"`
	CreatedAt      string   `json:"created_at"`
	Deleted        bool     `json:"deleted"`
	DeletedAt      string   `json:"deleted_at"`
	ResourceId     int32    `json:"id"`
	Name           string   `json:"name"`
	Permissions    []string `json:"permissions"`
	ResourceTypeID int      `json:"resource_type_id"`
	UID            int      `json:"uid"`
	UpdatedAt      string   `json:"updated_at"`
}

type Perm struct {
	Name string `json:"name"`
	Rid  int32  `json:"rid"`
}

type ResourcePermissionsRespItem struct {
	Perms []*Perm `json:"perms"`
}

type Acl struct {
	Uid         int32    `json:"uid"`
	UserName    string   `json:"userName"`
	Rid         int32    `json:"rid"`
	RoleName    string   `json:"roleName"`
	ParentRoles []string `json:"parentRoles"`
	ChildRoles  []string `json:"childRoles"`
	NickName    string   `json:"nickName"`
}

type Session struct {
	Uid int `json:"uid"`
	Acl Acl `json:"acl"`
}

func (s *Session) GetUid() int {
	return s.Uid
}

func (s *Session) GetUserName() string {
	return s.Acl.UserName
}

func (a *Acl) GetUserName(ctx *gin.Context) string {
	res, exist := ctx.Get("session")
	if exist {
		if v, ok := res.(*Session); ok {
			return v.GetUserName()
		}
	}
	return ""
}

func (a *Acl) GetUserInfo(ctx *gin.Context) (any, error) {
	res, exist := ctx.Get("session")
	if exist {
		if v, ok := res.(*Session); ok {
			return v, nil
		}
	}
	return res, fmt.Errorf("no session")
}
