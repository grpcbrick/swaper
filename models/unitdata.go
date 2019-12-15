package models

import (
	"database/sql"

	"github.com/grpcbrick/switch/standard"
)

// UnitData 数据
type UnitData struct {
	Body          sql.NullString
	ExpiryTime    sql.NullTime
	DestroyTime   sql.NullTime
	CreatedTime   sql.NullTime
	EffectiveTime sql.NullTime
}

// EqualPassword EqualPassword
func (srv *UnitData) EqualPassword(target string) bool {
	return false
}

// LoadStringMap 从 string map 中加载数据
func (srv *UnitData) LoadStringMap(data map[string]string) {
	srv.Body.Scan(data["Body"])
	srv.ExpiryTime.Scan(data["ExpiryTime"])
	srv.DestroyTime.Scan(data["DestroyTime"])
	srv.CreatedTime.Scan(data["CreatedTime"])
	srv.EffectiveTime.Scan(data["EffectiveTime"])

}

// LoadProtoStruct LoadProtoStruct
func (srv *UnitData) LoadProtoStruct(Data *standard.UnitData) {
	srv.Body.Scan(Data.Body)
	srv.ExpiryTime.Scan(Data.ExpiryTime)
	srv.DestroyTime.Scan(Data.DestroyTime)
	srv.CreatedTime.Scan(Data.CreatedTime)
	srv.EffectiveTime.Scan(Data.EffectiveTime)
}

// OutProtoStruct OutProtoStruct
func (srv *UnitData) OutProtoStruct() *standard.UnitData {
	Data := new(standard.UnitData)
	Data.Body = srv.Body.String
	Data.ExpiryTime = srv.ExpiryTime.Time.String()
	Data.DestroyTime = srv.DestroyTime.Time.String()
	Data.CreatedTime = srv.CreatedTime.Time.String()
	Data.EffectiveTime = srv.EffectiveTime.Time.String()

	return Data
}
