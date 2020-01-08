package provider

import (
	"context"
	"time"

	"github.com/grpcbrick/swaper/dao"
	"github.com/grpcbrick/swaper/standard"
	validators "github.com/grpcbrick/swaper/validators"
)

// NewService NewService
func NewService() *Service {
	service := new(Service)
	service.encryptKey = "IAMGROOT" // 8 byte
	return service
}

// Service Service
type Service struct {
	encryptKey string
}

// EncryptID 加密
// func (srv *Service) EncryptID(id int64) string {
// 	byteKey := []byte(srv.encryptKey)
// 	byteData := []byte(strconv.FormatInt(id, 10))

// 	//将字节秘钥转换成block快
// 	block, _ := des.NewCipher(byteKey)

// 	// PKCS5Padding
// 	//计算出需要补多少位
// 	padding := block.BlockSize() - len(byteData)%block.BlockSize()
// 	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
// 	// 需要补padding位的padding值
// 	origData := bytes.Repeat([]byte{byte(padding)}, padding)

// 	//设置加密方式
// 	blockMode := cipher.NewCBCEncrypter(block, byteKey)
// 	//创建明文长度的字节数组
// 	crypted := make([]byte, len(origData))
// 	//加密明文,加密后的数据放到数组中
// 	blockMode.CryptBlocks(crypted, origData)
// 	//将字节数组转换成字符串
// 	return base64.StdEncoding.EncodeToString(crypted)
// }

// // DecryptID 解密
// func (srv *Service) DecryptID(data string) (int64, error) {
// 	byteData := []byte(data)
// 	byteKey := []byte(srv.encryptKey)

// 	block, _ := des.NewCipher(byteKey)
// 	blockMode := cipher.NewCBCDecrypter(block, byteKey)
// 	origData := make([]byte, len(byteData))
// 	// origData := crypted
// 	blockMode.CryptBlocks(origData, byteData)

// 	// PKCS5UnPadding
// 	length := len(origData)
// 	// 去掉最后一个字节 unpadding 次
// 	unpadding := int(origData[length-1])

// 	return strconv.ParseInt(string(origData[:(length-unpadding)]), 10, 64)
// }

// CreateUnitData 创建一个单位数据 返回一个 key
func (srv *Service) CreateUnitData(ctx context.Context, req *standard.CreateUnitDataRequest) (resp *standard.CreateUnitDataResponse, err error) {
	now := time.Now()
	var expiryTime time.Time
	var destroyTime time.Time
	var effectiveTime time.Time
	resp = new(standard.CreateUnitDataResponse)

	// TODO: 最大时间需要检查

	if req.EffectiveTime == 0 {
		// 默认立即生效
		effectiveTime = now
	} else {
		// 当前时间 + EffectiveTime/秒
		effectiveTime = now.Add(time.Duration(req.EffectiveTime * time.Second.Nanoseconds()))
	}

	if req.ExpiryTime == 0 {
		// 默认生效后五分钟有效期
		expiryTime = effectiveTime.Add(time.Minute * 5)
	} else {
		// 当前时间 + ExpiryTime/秒
		expiryTime = now.Add(time.Duration(req.ExpiryTime * time.Second.Nanoseconds()))
	}

	if req.DestroyTime == 0 {
		// 默认有效期后2小时销毁时间
		destroyTime = expiryTime.Add(time.Hour * 2)
	} else {
		// 当前时间 + DestroyTime/秒
		destroyTime = now.Add(time.Duration(req.DestroyTime * time.Second.Nanoseconds()))
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

// GetUnitDataByKey 通过 key 获取一个单位数据
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

// DestroyUnitDataByKey 立即销毁一个单位数据
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

// UpdateUnitDataBodyByKey 根据 key 更新单位数据的 body 部分
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

// UpdateUnitDataExpiryTimeByKey 根据 key 更新单位数据的 ExpiryTime 部分
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

	expiryTime := time.Now().Add(time.Duration(req.ExpiryTime * time.Second.Nanoseconds()))
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

// UpdateUnitDataDestroyTimeByKey 根据 key 更新单位数据的 DestroyTime 部分
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

	destroyTime := time.Now().Add(time.Duration(req.DestroyTime * time.Second.Nanoseconds()))
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

// UpdateUnitDataEffectiveTimeByKey 根据 key 更新单位数据的 EffectiveTime 部分
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

	effectiveTime := time.Now().Add(time.Duration(req.EffectiveTime * time.Second.Nanoseconds()))
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
