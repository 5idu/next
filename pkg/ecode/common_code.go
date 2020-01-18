package ecode

// APICode api code
type APICode string

// APICodeMapZH api code map zh
var APICodeMapZH map[APICode]string

func init() {
	APICodeMapZH = make(map[APICode]string)
	APICodeMapZH[Default] = "未处理的错误"
	APICodeMapZH[OK] = "成功"
	APICodeMapZH[MethodNotAllowed] = "不支持该方法"
	APICodeMapZH[ServerErr] = "服务器错误"
	APICodeMapZH[BadParameters] = "参数验证未通过"
	APICodeMapZH[Unauthorized] = "未认证"
	APICodeMapZH[AccessDenied] = "无权访问"
	APICodeMapZH[NotFound] = "数据不存在"
	APICodeMapZH[DBError] = "数据库操作异常"
	APICodeMapZH[CacheError] = "缓存操作异常"
}

// api code define
const (
	Default          APICode = ""
	OK               APICode = "OK"
	MethodNotAllowed APICode = "MethodNotAllowed"
	ServerErr        APICode = "ServerErr"
	BadParameters    APICode = "BadParameters"
	Unauthorized     APICode = "Unauthorized"
	AccessDenied     APICode = "AccessDenied"
	NotFound         APICode = "NotFound"
	DBError          APICode = "DBError"
	CacheError       APICode = "CacheError"
)
