package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/grpcbrick/switch/dao"
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

// func InitTestData() {
// 	// 预创建一条测试用户数据
// 	// 方便 label、group 测试
// 	srv := NewService()
// 	resp, err := srv.CreateUser(
// 		context.Background(),
// 		&standard.CreateUserRequest{Class: "InitTestData", Inviter: 0, Nickname: "InitTestData", Username: "InitTestData", Password: "InitTestData"},
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if resp.State != standard.State_SUCCESS {
// 		panic(fmt.Errorf("准备测试数据失败、用户创建失败: %v", resp))
// 	}

// 	fmt.Printf("预插入一条用户数据 ID 为: %d \n", resp.Data.ID)
// }

// func TestService_CreateUser(t *testing.T) {
// 	srv := NewService()
// 	tests := []struct {
// 		name      string
// 		args      *standard.CreateUserRequest
// 		wantState standard.State
// 		wantErr   bool
// 	}{
// 		{"正常创建", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password"},
// 			standard.State_SUCCESS, false},

// 		{"重复的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username", Password: "Password"},
// 			standard.State_USER_ALREADY_EXISTS, false},

// 		{"空的 Class", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username1", Password: "Password"},
// 			standard.State_SUCCESS, false},

// 		{"空的 Inviter", &standard.CreateUserRequest{Class: "TEST", Nickname: "Nickname", Username: "Username2", Password: "Password"},
// 			standard.State_SUCCESS, false},

// 		{"空的 Nickname", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Username: "Username3", Password: "Password"},
// 			standard.State_PARAMS_INVALID, false},

// 		{"空的 Username", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Password: "Username4"},
// 			standard.State_PARAMS_INVALID, false},

// 		{"空的 Password", &standard.CreateUserRequest{Class: "TEST", Inviter: 1, Nickname: "Nickname", Username: "Username5"},
// 			standard.State_PARAMS_INVALID, false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotResp, err := srv.CreateUser(context.Background(), tt.args)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotResp.State.String() != tt.wantState.String() {
// 				t.Errorf("Service.CreateUser() = %v, want %v", gotResp, tt.wantState)
// 				return
// 			}
// 			if gotResp.State == standard.State_SUCCESS {
// 				if gotResp.Data.Username != tt.args.Username {
// 					t.Errorf("Service.CreateUser() = %v, want %v", gotResp, tt.wantState)
// 					return
// 				}
// 			}
// 		})
// 	}
// }
