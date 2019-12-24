package common

import (
    "time"
    "math/rand"
    "github.com/kataras/iris/sessions"
    config "github.com/spf13/viper"
)


/*session全局共享定义*/
var SessManager = sessions.New(sessions.Config{
    Cookie:  config.GetString("application.SessionCoolieName"),
    Expires: time.Duration(config.GetInt64("application.SessionExpires")) * time.Hour,
})

/*获取随机整数*/
func GenerateRangeNum(min int, max int) int {
    if min == max {
        return min
    }
    rand.Seed(time.Now().Unix())
    randNum := rand.Intn(max-min) + min
    return randNum
}