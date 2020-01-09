package models

import (
	"database/sql"

	"github.com/grpcbrick/swaper/standard"
)

// UnitData 数据
type UnitData struct {
	Key           sql.NullString
	Body          sql.NullString
	ExpiryTime    sql.NullTime
	DestroyTime   sql.NullTime
	EffectiveTime sql.NullTime
	DeletedTime   sql.NullTime
	CreatedTime   sql.NullTime
	UpdatedTime   sql.NullTime
}

// LoadProtoStruct LoadProtoStruct
func (srv *UnitData) LoadProtoStruct(Data *standard.UnitData) {
	srv.Body.Scan(Data.Body)
	srv.ExpiryTime.Scan(Data.ExpiryTime)
	srv.DestroyTime.Scan(Data.DestroyTime)
	srv.CreatedTime.Scan(Data.CreatedTime)
	srv.DeletedTime.Scan(Data.DeletedTime)
	srv.UpdatedTime.Scan(Data.UpdatedTime)
	srv.EffectiveTime.Scan(Data.EffectiveTime)
}

// OutProtoStruct OutProtoStruct
func (srv *UnitData) OutProtoStruct() *standard.UnitData {
	Data := new(standard.UnitData)
	Data.Body = srv.Body.String
	Data.ExpiryTime = srv.ExpiryTime.Time.String()
	Data.DestroyTime = srv.DestroyTime.Time.String()
	Data.DeletedTime = srv.DeletedTime.Time.String()
	Data.CreatedTime = srv.CreatedTime.Time.String()
	Data.UpdatedTime = srv.UpdatedTime.Time.String()
	Data.EffectiveTime = srv.EffectiveTime.Time.String()
	return Data
}
