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
		args      standard.CreateUnitDataRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常创建", standard.CreateUnitDataRequest{}, standard.State_SUCCESS, false},
		{"正常创建", standard.CreateUnitDataRequest{Body: "Body2"}, standard.State_SUCCESS, false},
		{"正常创建", standard.CreateUnitDataRequest{}, standard.State_SUCCESS, false},
		{"正常创建", standard.CreateUnitDataRequest{}, standard.State_SUCCESS, false},
		{"正常创建", standard.CreateUnitDataRequest{}, standard.State_SUCCESS, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试创建
			createResp, err := srv.CreateUnitData(context.Background(), &tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUnitData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if createResp.State.String() != tt.wantState.String() {
				t.Errorf("Service.CreateUnitData() = %v, want %v", createResp, tt.wantState)
				return
			}

			if tt.wantState != standard.State_SUCCESS {
				return
			}

			// 测试 query
			getResp, err := srv.GetUnitDataByKey(context.Background(), &standard.GetUnitDataByKeyRequest{Key: createResp.Key})
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUnitDataByKey() error = %v", getResp)
				return
			}
			t.Errorf("11111, %s %v %v", createResp.Key, getResp.Data, tt.args)

			// if getResp.Data.Body != tt.args.Body {
			// 	t.Errorf("Service.GetUnitDataByKey() = %v, want %v", getResp, tt.args)
			// 	return
			// }

		})
	}
}
