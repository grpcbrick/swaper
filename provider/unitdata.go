package provider

import (
	"context"
	"time"

	"github.com/grpcbrick/switch/dao"
	"github.com/grpcbrick/switch/standard"
	validators "github.com/grpcbrick/switch/validators"
)

var timeLayout = time.RFC850

// NewService NewService
func NewService() *Service {
	service := new(Service)
	return service
}

// Service Service
type Service struct {
}

// CreateUnitData CreateUnitData
func (srv *Service) CreateUnitData(ctx context.Context, req *standard.CreateUnitDataRequest) (resp *standard.CreateUnitDataResponse, err error) {
	resp = new(standard.CreateUnitDataResponse)
	var effectiveTime time.Time
	var destroyTime time.Time
	var expiryTime time.Time

	if req.EffectiveTime == "" {
		// 默认立即生效
		effectiveTime = time.Now()
	} else {
		effectiveTime, err = time.Parse(timeLayout, req.EffectiveTime)
		if err != nil {
			resp.State = standard.State_PARAMS_INVALID
			resp.Message = err.Error()
			return resp, nil
		}
	}

	if req.ExpiryTime == "" {
		// 默认生效后五分钟有效期
		expiryTime = effectiveTime.Add(time.Minute * 5)
	} else {
		expiryTime, err = time.Parse(timeLayout, req.ExpiryTime)
		if err != nil {
			resp.State = standard.State_PARAMS_INVALID
			resp.Message = err.Error()
			return resp, nil
		}
	}

	if req.DestroyTime == "" {
		// 默认有效期后2小时销毁时间
		destroyTime = expiryTime.Add(time.Hour * 2)
	} else {
		destroyTime, err = time.Parse(timeLayout, req.DestroyTime)
		if err != nil {
			resp.State = standard.State_PARAMS_INVALID
			resp.Message = err.Error()
			return resp, nil
		}
	}

	key, err := dao.CreateUnitData(req.Body, effectiveTime, destroyTime, expiryTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"
	resp.Key = key
	return resp, nil
}

// GetUnitDataByKey GetUnitDataByKey
func (srv *Service) GetUnitDataByKey(ctx context.Context, req *standard.GetUnitDataByKeyRequest) (resp *standard.GetUnitDataByKeyResponse, err error) {
	resp = new(standard.GetUnitDataByKeyResponse)

	if ok, msg := validators.Key(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	uintData, err := dao.QueryUnitDataByKey(req.Key)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.Data = uintData.OutProtoStruct()
	resp.State = standard.State_SUCCESS
	resp.Message = "创建成功"
	return resp, err
}

// DestroyUnitDataByKey DestroyUnitDataByKey
func (srv *Service) DestroyUnitDataByKey(ctx context.Context, req *standard.DestroyUnitDataByKeyRequest) (resp *standard.DestroyUnitDataByKeyResponse, err error) {
	resp = new(standard.DestroyUnitDataByKeyResponse)
	if ok, msg := validators.Key(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	err = dao.UpdateUintDataDestroyTimeByKey(req.Key, time.Now())
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "销毁成功"
	return resp, err
}

// UpdateUnitDataBodyByKey UpdateUnitDataBodyByKey
func (srv *Service) UpdateUnitDataBodyByKey(ctx context.Context, req *standard.UpdateUnitDataBodyByKeyRequest) (resp *standard.UpdateUnitDataBodyByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataBodyByKeyResponse)
	if ok, msg := validators.Key(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	err = dao.UpdateUnitDataBodyByKey(req.Key, req.Body)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, err
}

// UpdateUnitDataExpiryTimeByKey UpdateUnitDataExpiryTimeByKey
func (srv *Service) UpdateUnitDataExpiryTimeByKey(ctx context.Context, req *standard.UpdateUnitDataExpiryTimeByKeyRequest) (resp *standard.UpdateUnitDataExpiryTimeByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataExpiryTimeByKeyResponse)
	if ok, msg := validators.Key(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.ExpiryTime(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	expiryTime, err := time.Parse(timeLayout, req.ExpiryTime)
	if err != nil {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = err.Error()
		return resp, nil
	}

	err = dao.UpdateUintDataExpiryTimeByKey(req.Key, expiryTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, err
}

// UpdateUnitDataDestroyTimeByKey UpdateUnitDataDestroyTimeByKey
func (srv *Service) UpdateUnitDataDestroyTimeByKey(ctx context.Context, req *standard.UpdateUnitDataDestroyTimeByKeyRequest) (resp *standard.UpdateUnitDataDestroyTimeByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataDestroyTimeByKeyResponse)
	if ok, msg := validators.Key(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.DestroyTime(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	destroyTime, err := time.Parse(timeLayout, req.DestroyTime)
	if err != nil {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = err.Error()
		return resp, nil
	}

	err = dao.UpdateUintDataDestroyTimeByKey(req.Key, destroyTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, err
}

// UpdateUnitDataEffectiveTimeByKey UpdateUnitDataEffectiveTimeByKey
func (srv *Service) UpdateUnitDataEffectiveTimeByKey(ctx context.Context, req *standard.UpdateUnitDataEffectiveTimeByKeyRequest) (resp *standard.UpdateUnitDataEffectiveTimeByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataEffectiveTimeByKeyResponse)
	if ok, msg := validators.Key(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	if ok, msg := validators.EffectiveTime(req.Key); ok != true {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = msg
		return resp, nil
	}

	effectiveTime, err := time.Parse(timeLayout, req.EffectiveTime)
	if err != nil {
		resp.State = standard.State_PARAMS_INVALID
		resp.Message = err.Error()
		return resp, nil
	}

	err = dao.UpdateUintDataEffectiveTimeByKey(req.Key, effectiveTime)
	if err != nil {
		resp.State = standard.State_DB_OPERATION_FATLURE
		resp.Message = err.Error()
		return resp, nil
	}

	resp.State = standard.State_SUCCESS
	resp.Message = "更新成功"
	return resp, err
}
