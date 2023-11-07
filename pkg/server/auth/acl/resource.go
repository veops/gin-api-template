package acl

import (
	"context"
	"fmt"

	"app/pkg/conf"
)

func AddResource(ctx context.Context, uid int, resourceTypeId string, name string) (res *Resource, err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}

	res = &Resource{}
	url := fmt.Sprintf("%s/acl/resources", conf.Cfg.Auth.Acl.Url)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetBody(map[string]any{
			"type_id": resourceTypeId,
			"name":    name,
			"uid":     uid,
		}).
		SetResult(&res).
		Post(url)
	err = HandleErr(err, resp, func(dt map[string]any) bool { return true })
	return
}

func DeleteResource(ctx context.Context, resourceId int32) (err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%v/acl/resources/%v", conf.Cfg.Auth.Acl.Url, resourceId)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		Delete(url)
	err = HandleErr(err, resp, func(dt map[string]any) bool { return true })
	return
}

func UpdateResource(ctx context.Context, resourceId int, updates map[string]string) (err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/acl/resources/%d", conf.Cfg.Auth.Acl.Url, resourceId)
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetFormData(updates).
		Put(url)
	err = HandleErr(err, resp, func(dt map[string]any) bool { return true })
	return
}

func GetResourcePermissions(ctx context.Context, resourceId int32) (res map[string]*ResourcePermissionsRespItem, err error) {
	token, err := GetAclToken(ctx)
	if err != nil {
		return
	}
	res = make(map[string]*ResourcePermissionsRespItem)
	url := fmt.Sprintf("%v/acl/resources/%v/permissions", conf.Cfg.Auth.Acl.Url, resourceId) //TODO conf
	resp, err := RC.R().
		SetHeader("App-Access-Token", token).
		SetResult(&res).
		Get(url)
	err = HandleErr(err, resp, func(dt map[string]any) bool { return true })
	return
}
