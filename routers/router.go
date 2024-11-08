package routers

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"tndcsc/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 創建頻率限制器
	limiter := tollbooth.NewLimiter(1, nil) // 每秒最多 1 次請求
	limiter.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})
	limiter.SetMessageContentType("application/json; charset=utf-8")
	limiter.SetMessage(`{"http_code": "429", "message": "API 請求頻率過快，請稍後再試！", "status": "error"}`)

	// 設置需要頻率限制的路由
	r.GET("/", tollbooth_gin.LimitHandler(limiter), controllers.GetData)

	return r
}
