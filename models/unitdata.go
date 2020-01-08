package models

import (
	"database/sql"
	"math"
	"time"

	"github.com/grpcbrick/swaper/standard"
)

// UnitData 数据
type UnitData struct {
	Body          sql.NullString
	ExpiryTime    sql.NullTime
	DestroyTime   sql.NullTime
	CreatedTime   sql.NullTime
	EffectiveTime sql.NullTime
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
	now := time.Now()
	Data := new(standard.UnitData)
	Data.Body = srv.Body.String
	Data.ExpiryTime = int64(math.Floor(srv.ExpiryTime.Time.Sub(now).Seconds()))
	Data.DestroyTime = int64(math.Floor(srv.DestroyTime.Time.Sub(now).Seconds()))
	Data.CreatedTime = int64(math.Floor(srv.CreatedTime.Time.Sub(now).Seconds()))
	Data.EffectiveTime = int64(math.Floor(srv.EffectiveTime.Time.Sub(now).Seconds()))

	return Data
}
