package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// 可以决定在c.Next之前或之后调用指定方法
	r.Use(func(c *gin.Context) {
		// log latency, response code, path
		//logger.Info("incoming request",
		//	zap.String("path",c.Request.URL.Path))
		//c.Next()	// 之前

		s := time.Now()

		c.Next() // 之后
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("time", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int()) // 往context中塞东西
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}

		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
