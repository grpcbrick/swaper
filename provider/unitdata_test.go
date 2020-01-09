package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/swaper/dao"
	"github.com/grpcbrick/swaper/standard"
	"github.com/yinxulai/goutils/config"
	"github.com/yinxulai/goutils/sqldb"
)

var createdKays []string
var testEncryptKey = "TESTTEST"

func TestMain(m *testing.M) {
	fmt.Println("准备测试环境")                                                                       // 日志
	config.Set("encrypt_password", "encrypt_password")                                          // 密码加密中会用到的
	sqldb.Init("mysql", "root:root@tcp(localhost:3306)/default?charset=utf8mb4&parseTime=true") // 测试数据库
	dao.InitTables()                                                                            // 初始化测试数据库
	// InitTestData()                                                                              // 预插入一条用户数据
	fmt.Println("开始执行测试")        // 日志
	exitCode := m.Run()          // 执行测试
	fmt.Println("测试执行完成,清理测试数据") // 日志
	// dao.TruncateTables()         // 重置测试数据库
	os.Exit(exitCode) // 推出
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
				// GetUnitDataByKey 查询测试
				key := gotResp.GetKey()
				getResp, err := srv.GetUnitDataByKey(context.Background(), &standard.GetUnitDataByKeyRequest{Key: key})
				if err != nil {
					t.Errorf("Service.GetUnitDataByKey() 根据 Key 查询 UnitData 失败： %v", err)
					return
				}

				if getResp.State != standard.State_SUCCESS {
					t.Errorf("Service.GetUnitDataByKey() 查询失败, key：%v, 状态：%v, 错误信息：%v", key, getResp.State, getResp.Message)
					return
				}

				if getResp.GetData().GetBody() != tt.args.Body {
					t.Errorf("Service.CreateUser() 失败, 希望创建的 Body： %v, 创建后查询到： %v", tt.args.Body, getResp.GetData().GetBody())
					return
				}
				createdKays = append(createdKays, key)
			}
		})
	}
}

func TestService_UpdateUnitDataBodyByKey(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.UpdateUnitDataBodyByKeyRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常更新", &standard.UpdateUnitDataBodyByKeyRequest{Body: "Body1"},
			standard.State_SUCCESS, false},
		{"正常更新", &standard.UpdateUnitDataBodyByKeyRequest{Body: "Body2"},
			standard.State_SUCCESS, false},
		{"空的 Body", &standard.UpdateUnitDataBodyByKeyRequest{},
			standard.State_SUCCESS, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range createdKays {
				tt.args.Key = key
				gotResp, err := srv.UpdateUnitDataBodyByKey(context.Background(), tt.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("Service.UpdateUnitDataBodyByKey() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if gotResp.State.String() != tt.wantState.String() {
					t.Errorf("Service.UpdateUnitDataBodyByKey() = %v, want %v", gotResp, tt.wantState)
					return
				}
			}
		})
	}
}

func TestService_UpdateUnitDataExpiryTimeByKey(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.UpdateUnitDataExpiryTimeByKeyRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常更新", &standard.UpdateUnitDataExpiryTimeByKeyRequest{ExpiryTime: 1000},
			standard.State_SUCCESS, false},
		{"正常更新", &standard.UpdateUnitDataExpiryTimeByKeyRequest{ExpiryTime: 1000},
			standard.State_SUCCESS, false},
		{"空的 ExpiryTime", &standard.UpdateUnitDataExpiryTimeByKeyRequest{},
			standard.State_SUCCESS, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range createdKays {
				tt.args.Key = key
				gotResp, err := srv.UpdateUnitDataExpiryTimeByKey(context.Background(), tt.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("Service.UpdateUnitDataExpiryTimeByKey() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if gotResp.State.String() != tt.wantState.String() {
					t.Errorf("Service.UpdateUnitDataExpiryTimeByKey() = %v, want %v", gotResp, tt.wantState)
					return
				}
			}
		})
	}
}

func TestService_UpdateUnitDataDestroyTimeByKey(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.UpdateUnitDataDestroyTimeByKeyRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常更新", &standard.UpdateUnitDataDestroyTimeByKeyRequest{DestroyTime: 1000},
			standard.State_SUCCESS, false},
		{"正常更新", &standard.UpdateUnitDataDestroyTimeByKeyRequest{DestroyTime: 1000},
			standard.State_SUCCESS, false},
		{"空的 DestroyTime", &standard.UpdateUnitDataDestroyTimeByKeyRequest{},
			standard.State_SUCCESS, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range createdKays {
				tt.args.Key = key
				gotResp, err := srv.UpdateUnitDataDestroyTimeByKey(context.Background(), tt.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("Service.UpdateUnitDataDestroyTimeByKey() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if gotResp.State.String() != tt.wantState.String() {
					t.Errorf("Service.UpdateUnitDataDestroyTimeByKey() = %v, want %v", gotResp, tt.wantState)
					return
				}
			}
		})
	}
}

func TestService_UpdateUnitDataEffectiveTimeByKey(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.UpdateUnitDataEffectiveTimeByKeyRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常更新", &standard.UpdateUnitDataEffectiveTimeByKeyRequest{EffectiveTime: 1000},
			standard.State_SUCCESS, false},
		{"正常更新", &standard.UpdateUnitDataEffectiveTimeByKeyRequest{EffectiveTime: 1000},
			standard.State_SUCCESS, false},
		{"空的 EffectiveTime", &standard.UpdateUnitDataEffectiveTimeByKeyRequest{},
			standard.State_SUCCESS, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range createdKays {
				tt.args.Key = key
				gotResp, err := srv.UpdateUnitDataEffectiveTimeByKey(context.Background(), tt.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("Service.UpdateUnitDataEffectiveTimeByKey() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if gotResp.State.String() != tt.wantState.String() {
					t.Errorf("Service.UpdateUnitDataEffectiveTimeByKey() = %v, want %v", gotResp, tt.wantState)
					return
				}
			}
		})
	}
}

func TestService_DestroyUnitDataByKey(t *testing.T) {
	srv := NewService()
	tests := []struct {
		name      string
		args      *standard.DestroyUnitDataByKeyRequest
		wantState standard.State
		wantErr   bool
	}{
		{"正常销毁", &standard.DestroyUnitDataByKeyRequest{},
			standard.State_SUCCESS, false},
		{"正常销毁", &standard.DestroyUnitDataByKeyRequest{},
			standard.State_SUCCESS, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range createdKays {
				tt.args.Key = key
				gotResp, err := srv.DestroyUnitDataByKey(context.Background(), tt.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("Service.DestroyUnitDataByKey() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if gotResp.State.String() != tt.wantState.String() {
					t.Errorf("Service.DestroyUnitDataByKey() = %v, want %v", gotResp, tt.wantState)
					return
				}
			}
		})
	}
}
