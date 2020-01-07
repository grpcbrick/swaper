package dao

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/grpcbrick/switch/models"
	"github.com/yinxulai/goutils/sqldb"
)

const unitDataTableName = "unit-datas"
const unitDataHistoryTableName = "unit-datas-history"

func createDataTable() error {
	// 主表
	matserStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + unitDataTableName + "`",
		"(",
		"`Key` VARCHAR(128) NOT NULL COMMENT 'Key',",
		"`Body` BLOB DEFAULT NULL COMMENT '数据',",
		"`ExpiryTime` DATETIME NOT NULL COMMENT '失效时间', ",
		"`DestroyTime` DATETIME NOT NULL COMMENT '销毁时间',",
		"`EffectiveTime` DATETIME NOT NULL COMMENT '生效时间',",
		"`CreatedTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`Key`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err := matserStmp.Exec()
	if err != nil {
		return err
	}
	// 历史表
	historyStmp := sqldb.CreateStmt(strings.Join([]string{
		"CREATE TABLE IF NOT EXISTS `" + unitDataHistoryTableName + "`",
		"(",
		"`Index` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'Index',",
		"`Key` VARCHAR(128)  COMMENT 'Key',",
		"`Body` BLOB COMMENT '数据',",
		"`ExpiryTime` DATETIME NOT NULL COMMENT '失效时间', ",
		"`DestroyTime` DATETIME NOT NULL COMMENT '销毁时间',",
		"`EffectiveTime` DATETIME NOT NULL COMMENT '生效时间',",
		"`CreatedTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',",
		"`UpdatedTime` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',",
		"PRIMARY KEY (`Index`)",
		")",
		"ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;",
	}, " "))
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}

	return nil
}

func truncateDataTable() error {
	var err error
	masterStmp := sqldb.CreateStmt("truncate table `" + unitDataTableName + "`")
	_, err = masterStmp.Exec()
	if err != nil {
		return err
	}
	historyStmp := sqldb.CreateStmt("truncate table `" + unitDataHistoryTableName + "`")
	_, err = historyStmp.Exec()
	if err != nil {
		return err
	}
	return nil
}

// CreateUnitDataHistory 对指定数据创建一条历史快照
func CreateUnitDataHistory(key string) error {
	var err error
	namedData := map[string]interface{}{
		"Key": key,
	}

	// 插入一条更新历史
	historyStmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + unitDataHistoryTableName + "`",
		"(Key,Data,ExpiryTime,DestroyTime,EffectiveTime,CreatedTime,UpdatedTime)",
		"SELECT",
		"Key,Data,ExpiryTime,DestroyTime,EffectiveTime,CreatedTime,UpdatedTime",
		"FROM",
		"`" + unitDataTableName + "`",
		"WHERE",
		"`Key`=:Key",
	}, " "))
	_, err = historyStmp.Exec(namedData)
	if err != nil {
		return err
	}
	return nil
}

// CountUnitDataByKey 根据 key 统计
func CountUnitDataByKey(key string) (int64, error) {
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + unitDataTableName + "`",
		"WHERE",
		"`Key`=:Key",
	}, " "))

	result := struct{ Count int64 }{}
	namedData := map[string]interface{}{
		"Key": key,
	}
	err := stmp.Get(&result, namedData)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// QueryUnitDatas 查询用户
func QueryUnitDatas(page, limit int64) (totalPage, currentPage int64, datas []*models.UnitData, err error) {
	currentPage = page // 固定当前页

	// 查询数据长度
	countStmp := sqldb.CreateStmt(strings.Join([]string{
		"SELECT COUNT(*) as Count FROM",
		"`" + unitDataTableName + "`",
	}, " "))

	countResult := struct{ Count int64 }{}
	err = countStmp.Get(&countResult)
	if err != nil {
		return totalPage, currentPage, datas, err
	}

	count := countResult.Count
	// 计算总页码数
	totalPage = int64(math.Ceil(float64(count) / float64(limit)))

	// 超出数据总页数
	if page > totalPage {
		// 返回当前页、空数据（当前页确实不存在数据）
		return totalPage, page, datas, err
	}

	// 计算偏移
	offSet := (page - 1) * limit
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + unitDataTableName + "`",
		"LIMIT :Limit",
		"OFFSET :Offset",
	}, " "))

	datas = []*models.UnitData{}
	namedData := map[string]interface{}{
		"Limit":  limit,
		"Offset": offSet,
	}

	err = stmp.Select(&datas, namedData)
	if err != nil {
		return totalPage, currentPage, datas, err
	}

	return totalPage, currentPage, datas, err
}

// QueryUnitDataByKey 根据 key 查询
func QueryUnitDataByKey(key string) (*models.UnitData, error) {
	var err error
	namedData := map[string]interface{}{"Key": key}
	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"SELECT * FROM",
		"`" + unitDataTableName + "`",
		"WHERE",
		"`Key`=:Key",
	}, " "))

	data := new(models.UnitData)
	err = stmp.Get(data, namedData)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateUnitData 创建数据
func CreateUnitData(body string, expirytime, destroytime, effectivetime time.Time) (int64, error) {
	var err error

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"INSERT INTO",
		"`" + unitDataTableName + "`",
		"(`Key`, `Body`, `ExpiryTime`, `DestroyTime`, `EffectiveTime`)",
		"VALUES",
		"(:Key, :Body, :ExpiryTime, :DestroyTime, :EffectiveTime)",
	}, " "))

	namedData := map[string]interface{}{
		"Body":          body,
		"ExpiryTime":    expirytime,
		"DestroyTime":   destroytime,
		"EffectiveTime": effectivetime,
	}

	result, err := stmp.Exec(namedData)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdataUnitDataFieldByKey 根据 Key 更新指定字段
func UpdataUnitDataFieldByKey(key string, field map[string]interface{}) error {
	var err error

	fieldSQL := []string{}
	for name := range field {
		fieldSQL = append(fieldSQL, fmt.Sprintf("`%s`=:%s", name, name))
	}

	stmp := sqldb.CreateNamedStmt(strings.Join([]string{
		"UPDATE",
		"`" + unitDataTableName + "`",
		"SET",
		strings.Join(fieldSQL, ","),
		"WHERE",
		"`Key`=:Key",
	}, " "))

	// 修改前创建历史
	err = CreateUnitDataHistory(key)
	if err != nil {
		return err
	}

	// 更新
	field["Key"] = key
	_, err = stmp.Exec(field)
	return err
}

// UpdateUnitDataBodyByKey 更新数据
func UpdateUnitDataBodyByKey(key string, body string) error {
	return UpdataUnitDataFieldByKey(key, map[string]interface{}{"Body": body})
}

// UpdateUintDataExpiryTimeByKey 更新过期时间
func UpdateUintDataExpiryTimeByKey(key string, expiryTime time.Time) error {
	return UpdataUnitDataFieldByKey(key, map[string]interface{}{"ExpiryTime": expiryTime})
}

// UpdateUintDataDestroyTimeByKey 更新销毁时间
func UpdateUintDataDestroyTimeByKey(key string, destroyTime time.Time) error {
	return UpdataUnitDataFieldByKey(key, map[string]interface{}{"DestroyTime": destroyTime})
}

// UpdateUintDataEffectiveTimeByKey 更新生效时间
func UpdateUintDataEffectiveTimeByKey(key string, dataEffective time.Time) error {
	return UpdataUnitDataFieldByKey(key, map[string]interface{}{"EffectiveTime": dataEffective})
}

// 生成一个为一个 key
func generateUniqueKey() (string, error) {
	randomString := func(l int) string {
		str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		bytes := []byte(str)
		result := []byte{}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < l; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}
		return string(result)
	}

	for {
		newKey := randomString(128)
		count, err := CountUnitDataByKey(newKey)
		if err != nil {
			return "", err
		}

		if count <= 0 {
			return newKey, nil
		}
	}
}
