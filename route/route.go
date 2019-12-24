package route

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
    "github.com/380566067/GoPanApp/application/common"
)

func Routes(app *iris.Application) {
    mvc.New(app.Party("/")).
    Register(common.SessManager.Start)
}