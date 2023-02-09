package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/xq-libs/go-utils/stringutil"
	"log"
	"strconv"
)

func GetRequestUser(c *gin.Context) User {
	return User{
		ID:       "100001",
		TenantId: "100001",
		Account:  "Admin",
		Name:     "超级管理员",
	}
}

func GetRequestPage(c *gin.Context) Pageable {
	return Pageable{
		Page: GetQueryInt(c, "page", 0),
		Size: GetQueryInt(c, "size", 10),
	}
}

func GetQueryInt(c *gin.Context, key string, df int) int {
	v := c.Query(key)
	if stringutil.IsNotBlank(v) {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Panicf("Value %s is not convert to int", v)
		}
		return i
	}
	return df
}

func GetParamInt(c *gin.Context, key string, df int) int {
	v := c.Param(key)
	if stringutil.IsNotBlank(v) {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Panicf("Value %s is not convert to int", v)
		}
		return i
	}
	return df
}

func GetRequestBody[T any](c *gin.Context, t T) T {
	err := c.ShouldBind(&t)
	if err != nil {
		log.Panicf("Bind obj from body failure: %v", err)
	}
	return t
}

func GetRequestQuery[T any](c *gin.Context, t T) T {
	err := c.ShouldBindQuery(&t)
	if err != nil {
		log.Panicf("Bind obj from query failure: %v", err)
	}
	return t
}
