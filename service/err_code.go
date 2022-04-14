// @Title
// @Description
// @Author  Wangwengang  2022/4/14 下午5:18
// @Update  Wangwengang  2022/4/14 下午5:18
package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wwengg/proto/common"
	"github.com/wwengg/proto/rainbow"
)

var MapErrMsgZH = map[common.EnumErr]string{
	common.EnumErr_SUCCESS: `请求成功`,
}

func getErrMsgZH(errMap map[common.EnumErr]string, errcode common.EnumErr) string {
	if v, ok := errMap[errcode]; ok {
		return v
	}
	return fmt.Sprintf(`错误(%d)`, errcode)
}

func GetErr(errMap map[common.EnumErr]string, err common.EnumErr) error {
	var reply rainbow.HttpReply
	reply.Code = err
	reply.Message = getErrMsgZH(errMap, err)
	if bytes, err := json.Marshal(reply); err == nil {
		errMessage := base64.StdEncoding.EncodeToString(bytes)
		return errors.New(errMessage)
	}
	return errors.New("")
}
