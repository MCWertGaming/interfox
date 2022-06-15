package main

import (
	"github.com/PurotoApp/interfox/database"
	"github.com/PurotoApp/interfox/endpoints"
	"github.com/PurotoApp/interfox/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to the PostgreSQL
	pg_conn := database.ConnectSQL()
	if pg_conn.Error != nil {
		tools.ErrorPanic(pg_conn.Error)
	}
	// Connect to Redis
	redisVerify := database.ConnectRedis(1)  // stores the limited session for email verification
	redisSession := database.ConnectRedis(2) // stored the normal session

	// check if redis can be reached
	if err := redisVerify.Ping().Err(); err != nil {
		tools.ErrorPanic(err)
	} else if err := redisSession.Ping().Err(); err != nil {
		tools.ErrorPanic(err)
	}

	// migrate all tables
	database.AutoMigrateSQL(pg_conn) // migrate all tables

	// create router
	router := gin.Default()

	// configure gin
	endpoints.ConfigRouter(router)

	// set routes
	endpoints.SetRoutes(router, pg_conn, redisVerify, redisSession)

	// start
	if err := router.Run("0.0.0.0:3621"); err != nil {
		tools.ErrorPanic(err)
	}

	// clean up
	redisVerify.Close()
	redisSession.Close()
}
