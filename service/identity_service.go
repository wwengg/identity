// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午10:11
// @Update  Wangwengang  2021/8/19 下午10:11
package service

import (
	"context"
	"strconv"

	"github.com/wwengg/proto/common"
	"github.com/wwengg/proto/identity"

	"github.com/wwengg/arsenal/sdk/snowflake"
)

// SnowFlake twitter snowflake backend.
// It is based on https://github.com/bwmarrin/snowflake and
// - remove some unused methods
// - add GenerateBatch methods
//
// **How it Works**.
// Each time you generate an ID, it works, like this.
//
// A timestamp with millisecond precision is stored using 41 bits of the ID.
// Then the NodeID is added in subsequent bits.
// Then the Sequence Number is added, starting at 0 and incrementing for each ID generated in the same millisecond. If you generate enough IDs in the same millisecond that the sequence would roll over or overfill then the generate function will pause until the next millisecond.
//
// +--------------------------------------------------------------------------+
// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
// +--------------------------------------------------------------------------+
type IdentityService struct {
	serverID string
	node     *snowflake.Node
}

type IdentityAble interface {
	// IdentityAble can be used for interface verification.

	// GetId is server rpc method as defined
	GetId(ctx context.Context, args common.Empty, reply *identity.GetIdReply) (err error)

	// GetIds is server rpc method as defined
	GetIds(ctx context.Context, args identity.GetIdsArgs, reply *identity.GetIdsReply) (err error)
}

// NewSnowFlake created a new SnowFlake with given configs.
func NewSnowFlake(serverID int64, epoch int64, nodeBits, stepBits uint8) IdentityAble {
	node, err := snowflake.NewNode(serverID, epoch, nodeBits, stepBits)
	if err != nil {
		panic("can't create snowflake node: " + err.Error())
	}

	return &IdentityService{
		serverID: strconv.FormatInt(serverID, 10),
		node:     node,
	}
}

// GetId is server rpc method as defined
func (s *IdentityService) GetId(ctx context.Context, args common.Empty, reply *identity.GetIdReply) (err error) {
	*reply = identity.GetIdReply{
		Id: s.node.Generate(),
	}
	return
}

// GetIds is server rpc method as defined
func (s *IdentityService) GetIds(ctx context.Context, args identity.GetIdsArgs, reply *identity.GetIdsReply) (err error) {
	*reply = identity.GetIdsReply{
		Ids: s.node.GenerateBatch(uint16(args.Num)),
	}
	err = GetErr(MapErrMsgZH, common.EnumErr_JoinRoomErr)
	return nil
}
