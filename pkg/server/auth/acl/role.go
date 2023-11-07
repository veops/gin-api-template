package acl

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

	"app/pkg/conf"
)

func GetRoleResources(ctx context.Context, rid int32, resourceTypeId string) (res []*Resource, err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}

	data := &ResourceResult{}
	url := fmt.Sprintf("%v/acl/roles/%v/resources", conf.Cfg.Auth.Acl.Url, rid)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetQueryParams(map[string]string{
			"app_id":           conf.Cfg.Auth.Acl.AppId,
			"resource_type_id": resourceTypeId,
		}).
		SetResult(data).
		Get(url)

	if err = HandleErr(err, resp, func(dt map[string]any) bool { return true }); err != nil {
		return
	}

	res = data.Resources

	return
}

func HasPermission(ctx context.Context, rid int32, resourceName, resourceTypeName, permission string) (res bool, err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return false, err
	}

	data := make(map[string]any)
	url := fmt.Sprintf("%s/acl/roles/has_perm", conf.Cfg.Auth.Acl.Url)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetQueryParams(map[string]string{
			"rid":                cast.ToString(rid),
			"resource_name":      resourceName,
			"resource_type_name": resourceTypeName,
			"perm":               permission,
		}).
		SetResult(&data).
		Get(url)
	if err = HandleErr(err, resp, func(dt map[string]any) bool { return true }); err != nil {
		return
	}

	if v, ok := data["result"]; ok {
		res = v.(bool)
	}
	return
}

func GrantRoleResource(ctx context.Context, roleId int32, resourceId int32, permissions []string) (err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/acl/roles/%d/resources/%d/grant", conf.Cfg.Auth.Acl.Url, roleId, resourceId)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetBody(map[string]any{
			"perms": permissions,
		}).
		Post(url)
	err = HandleErr(err, resp, func(dt map[string]any) bool { return true })
	return
}

func RevokeRoleResource(ctx context.Context, roleId int32, resourceId int, permissions []string) (err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/acl/roles/%d/resources/%d/revoke", conf.Cfg.Auth.Acl.Url, roleId, resourceId)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetBody(map[string]any{
			"perms": permissions,
		}).
		Post(url)
	err = HandleErr(err, resp, func(dt map[string]any) bool { return true })
	return
}
