package acl

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"

	"app/pkg/conf"
	"app/pkg/server/storage/cache/redis"
)

var (
	RC *resty.Client
)

func init() {
	RC = resty.New()
	RC.SetRetryCount(3)
}

func GetAclToken(ctx context.Context) (res string, err error) {
	res, err = redis.RC.Get(ctx, "aclToken").Result()
	if err == nil {
		return
	}
	aclConfig := conf.Cfg.Auth.Acl

	url := fmt.Sprintf("%s%s", aclConfig.Url, "/acl/apps/token")
	secretHash := md5.Sum([]byte(aclConfig.SecretKey))
	secretKey := hex.EncodeToString(secretHash[:])

	data := make(map[string]string)
	resp, err := RC.R().
		SetBody(map[string]any{"app_id": aclConfig.AppId, "secret_key": secretKey}).
		SetResult(&data).
		Post(url)
	if err = HandleErr(err, resp, func(dt map[string]any) bool { return dt["token"] != "" }); err != nil {
		return
	}

	res = data["token"]
	_, err = redis.RC.SetNX(ctx, "aclToken", res, time.Hour).Result()
	return
}

func HandleErr(err error, resp *resty.Response, isOk func(dt map[string]any) bool) error {
	if err != nil {
		return err
	}

	dt := make(map[string]any)
	err = json.Unmarshal(resp.Body(), &dt)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 || !isOk(dt) {
		pc, _, _, _ := runtime.Caller(1)
		return fmt.Errorf("%s failed\n httpcode=%d resp=%s",
			runtime.FuncForPC(pc).Name(), resp.StatusCode(), string(resp.Body()))
	}
	return nil
}
