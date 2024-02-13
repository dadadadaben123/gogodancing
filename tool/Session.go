package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(engine *gin.Engine) {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mySession", store))
}

func SetSess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		Failed(context, "文件保存失败")
		return err
	}
	return nil
}

func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)

}
