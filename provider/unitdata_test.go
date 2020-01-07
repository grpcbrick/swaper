package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/switch/dao"
	"github.com/grpcbrick/switch/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/sqldb"
)

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境")                                                                       // 日志
	config.Set("encrypt_password", "encrypt_password")                                          // 密码加密中会用到的
	sqldb.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4&parseTime=true") // 测试数据库
	dao.InitTables()                                                                            // 初始化测试数据库
	// InitTestData()                                                                              // 预插入一条用户数据
	fmt.Println("开始执行测试")        // 日志
	exitCode := m.Run()          // 执行测试
	fmt.Println("测试执行完成,清理测试数据") // 日志
	dao.TruncateTables()         // 重置测试数据库
	os.Exit(exitCode)            // 推出
}

func TestService_CreateUnitData(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.CreateUnitDataRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常创建", &standard.CreateUnitDataRequest{Body: "Body", ExpiryTime: 10, DestroyTime: 10, EffectiveTime: 10},
			standard.State_SUCCESS, false},

		{"重复的 Body", &standard.CreateUnitDataRequest{Body: "Body", ExpiryTime: 10, DestroyTime: 10, EffectiveTime: 10},
			standard.State_SUCCESS, false},

		{"空的 ExpiryTime", &standard.CreateUnitDataRequest{Body: "Body", DestroyTime: 10, EffectiveTime: 10},
			standard.State_SUCCESS, false},

		{"空的 DestroyTime", &standard.CreateUnitDataRequest{Body: "Body", ExpiryTime: 10, EffectiveTime: 10},
			standard.State_SUCCESS, false},

		{"空的 EffectiveTime", &standard.CreateUnitDataRequest{Body: "Body", ExpiryTime: 10, DestroyTime: 10},
			standard.State_SUCCESS, false},

		{"空的 ExpiryTime、DestroyTime、EffectiveTime", &standard.CreateUnitDataRequest{Body: "Body"},
			standard.State_SUCCESS, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := srv.CreateUnitData(context.Background(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() 创建失败，错误：%v, 期望得到：%v", err, tt.wantErr)
				return
			}
			if gotResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.CreateUser() 返回状态不符合预期，得到：%v, 期望得到：%v", gotResp, tt.wantState)
				return
			}

			// 如果目标是成功、就去测一下查询接口
			if gotResp.State == standard.State_SUCCESS {
				getResp, err := srv.GetUnitDataByKey(context.Background(), &standard.GetUnitDataByKeyRequest{Key: gotResp.GetKey()})
				if err != nil {
					t.Errorf("Service.GetUnitDataByKey() 根据 Key 查询 UnitData 失败： %v", err)
					return
				}
				if getResp.GetData().GetBody() != tt.args.Body {
					t.Errorf("Service.CreateUser() 失败, 希望创建的 Body： %v, 创建后查询到： %v", tt.args.Body, getResp.GetData().GetBody())
				}
			}
		})
	}
}
