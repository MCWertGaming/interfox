package endpoints

import (
	"os"

	"github.com/PurotoApp/interfox/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, pg_conn *gorm.DB, redisVerify, redisSession *redis.Client) {
	//router.POST("/v1/user", registerUser(pg_conn, redisVerify, redisSession))
	//router.POST("/v1/user/login", loginUser(pg_conn, redisVerify, redisSession))
	//router.POST("/v1/user/verify", verifyUser(pg_conn, redisVerify, redisSession))
	//router.POST("/v1/user/validate", validateSession(redisVerify, redisSession))
	//router.PATCH("/v1/user", updatePassword(pg_conn, redisVerify, redisSession))
	//router.POST("/v1/user/delete", accountDeletion(collVerifySession, collSession, collUsers, collProfiles))
	// swagger docs
	router.Static("/swagger", "swagger/")
	// user redirects
	router.GET("/", redirect("/swagger"))
	router.GET("/v1", redirect("/swagger"))
	// static responses
	router.GET("/health", getHealth)
}

func ConfigRouter(router *gin.Engine) {

	if os.Getenv("GIN_MODE") == "release" {
		// turn on proxy support
		// TODO: allow users to specify trusted proxies
		// TODO: what if proxy behind proxy
		// TODO: what if no value specified
		tools.ErrorFatal("Router", router.SetTrustedProxies(nil))
	} else {
		// turn off proxy support for debugging
		tools.ErrorFatal("Router", router.SetTrustedProxies(nil))
	}
}
