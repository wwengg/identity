// @Title
// @Description
// @Author  Wangwengang  2022/4/14 下午5:43
// @Update  Wangwengang  2022/4/14 下午5:43
package service

import (
	"fmt"
	"testing"

	"github.com/wwengg/proto/common"
)

func TestName(t *testing.T) {
	fmt.Println(getErrMsgZH(common.EnumErr_JoinRoomErr, MapErrMsgZH))
}
