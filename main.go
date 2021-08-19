// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午12:59
// @Update  Wangwengang  2021/8/19 下午12:59
package main

import (
	"flag"

	"github.com/wwengg/arsenal/config"
	"github.com/wwengg/arsenal/logger"
	"github.com/wwengg/arsenal/sdk/rpcx"
	"github.com/wwengg/identity/service"
)

var (
	serverID = flag.Int64("serverID", 1, "unique server ID")
	epoch    = flag.Int64("epoch", 1580601600000, "epoch time for base timeline")
	nodeBits = flag.Int64("nodeBits", 8, "the number of bits to use for Node")
)

func main() {
	// Init config
	config.Viper()

	// Init logger
	logger.Setup()

	s := rpcx.NewRpcxServer()
	// register service
	identity := service.NewSnowFlake(*serverID, *epoch, uint8(*nodeBits), 22-uint8(*nodeBits))
	s.RegisterName("Identity", identity, "")

	// serve rpcx server
	s.Serve()
}
