package validators

import "regexp"

import "database/sql"

var (
	keyPattern  = regexp.MustCompile(`^[0-9a-zA-Z]{128}$`) // 128位固定长度
	bodyPattern = regexp.MustCompile(`^.{0,128}$`)
)

// Key 昵称检查
func Key(value string) (pass bool, msg string) {
	if keyPattern.MatchString(value) {
		return true, ""
	}

	return false, "无效的 Key 值"
}

// Body 用户名
func Body(value string) (pass bool, msg string) {
	if bodyPattern.MatchString(value) {
		return true, ""
	}
	return false, "不合法的 Body 值"
}

//TODO: 判断时间有效的要处理一下

// ExpiryTime 密码
func ExpiryTime(time string) (pass bool, msg string) {
	nullTime := sql.NullTime{}
	nullTime.Scan(time)
	if nullTime.Valid {
		return true, ""
	}

	return false, "无效的生效时间"
}

// DestroyTime 密码
func DestroyTime(time string) (pass bool, msg string) {
	nullTime := sql.NullTime{}
	nullTime.Scan(time)
	if nullTime.Valid {
		return true, ""
	}

	return false, "无效的销毁时间"
}

// EffectiveTime 密码
func EffectiveTime(time string) (pass bool, msg string) {
	nullTime := sql.NullTime{}
	nullTime.Scan(time)
	if nullTime.Valid {
		return true, ""
	}

	return false, "无效的生效时间"
}
