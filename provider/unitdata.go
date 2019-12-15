package provider

import (
	"context"

	"github.com/grpcbrick/switch/standard"
)

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
	

	return resp, err
}

// GetUnitDataByKey GetUnitDataByKey
func (srv *Service) GetUnitDataByKey(ctx context.Context, req *standard.GetUnitDataByKeyRequest) (resp *standard.GetUnitDataByKeyResponse, err error) {
	resp = new(standard.GetUnitDataByKeyResponse)
	return resp, err
}

// DestroyUnitDataByKey DestroyUnitDataByKey
func (srv *Service) DestroyUnitDataByKey(ctx context.Context, req *standard.DestroyUnitDataByKeyRequest) (resp *standard.DestroyUnitDataByKeyResponse, err error) {
	resp = new(standard.DestroyUnitDataByKeyResponse)
	return resp, err
}

// UpdateUnitDataBodyByKey UpdateUnitDataBodyByKey
func (srv *Service) UpdateUnitDataBodyByKey(ctx context.Context, req *standard.UpdateUnitDataBodyByKeyRequest) (resp *standard.UpdateUnitDataBodyByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataBodyByKeyResponse)
	return resp, err
}

// UpdateUnitDataExpiryTimeByKey UpdateUnitDataExpiryTimeByKey
func (srv *Service) UpdateUnitDataExpiryTimeByKey(ctx context.Context, req *standard.UpdateUnitDataExpiryTimeByKeyRequest) (resp *standard.UpdateUnitDataExpiryTimeByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataExpiryTimeByKeyResponse)
	return resp, err
}

// UpdateUnitDataDestroyTimeByKey UpdateUnitDataDestroyTimeByKey
func (srv *Service) UpdateUnitDataDestroyTimeByKey(ctx context.Context, req *standard.UpdateUnitDataDestroyTimeByKeyRequest) (resp *standard.UpdateUnitDataDestroyTimeByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataDestroyTimeByKeyResponse)
	return resp, err
}

// UpdateUnitDataEffectiveTimeByKey UpdateUnitDataEffectiveTimeByKey
func (srv *Service) UpdateUnitDataEffectiveTimeByKey(ctx context.Context, req *standard.UpdateUnitDataEffectiveTimeByKeyRequest) (resp *standard.UpdateUnitDataEffectiveTimeByKeyResponse, err error) {
	resp = new(standard.UpdateUnitDataEffectiveTimeByKeyResponse)
	return resp, err
}
