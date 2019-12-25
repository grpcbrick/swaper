package main

import (
	"net"

	"github.com/grpcbrick/swaper/dao"
	"github.com/grpcbrick/swaper/provider"
	"github.com/grpcbrick/swaper/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/grpc/interceptor"
	"github.com/yinxulai/goutils/sqldb"
	"google.golang.org/grpc"
)

func init() {
	config.SetStandard("rpc_port", ":3000", true, "RPC 服务监听的端口")
	config.SetStandard("mysql_url", "", true, "RPC 使用的 MYSQL 数据库配置")
	config.LoadFlag()
}

func main() {
	var err error
	sqldb.Init("mysql", config.MustGet("mysql_url"))
	dao.MustInitTables()
	lis, err := net.Listen("tcp", config.MustGet("rpc_port"))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(interceptor.NewCalllogs()...)
	standard.RegisterSwaperServer(grpcServer, provider.NewService())
	panic(grpcServer.Serve(lis))
}
