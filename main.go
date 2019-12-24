package main

import (
    "log"
    "github.com/kataras/iris"
    config "github.com/spf13/viper"
    "github.com/380566067/GoPanApp/route"
)

func main() {
    app := newApp()

    //应用App设置
    appConfig(app)

    //路由设置
    route.Routes(app)

    app.Favicon("./favicon.ico")
    app.Use(iris.Gzip)

    app.Run(
        iris.Addr(":" + config.GetString("server.Port")),   //进行端口监听
        iris.WithoutServerError(iris.ErrServerClosed),      //无服务错误提示
        iris.WithOptimizations,                             //更快的json序列化和更多优化
    )
}

//构建App
func newApp() *iris.Application {
    app := iris.New()

    //注册静态资源
    app.StaticWeb("/static", "./static")
    app.StaticWeb("/upload", "./upload")

    //读取配置信息
    config.AddConfigPath("./config")
    config.SetConfigName("config")
    if err := config.ReadInConfig(); err != nil {
        log.Fatalf("读取配置文件错误, %s", err)
    }

    //设置日志级别,开发阶段为debug
    if config.GetBool("application.AppDebug") == true {
        //设置debug
        app.Logger().SetLevel("debug")
        //设置模版重载
        app.RegisterView(iris.HTML("./application/view", ".html").Reload(true))
    } else {
        app.RegisterView(iris.HTML("./application/view", ".html"))
    }

    return app
}

/**
 * 项目设置
 */
func appConfig(app *iris.Application) {

    //配置 字符编码
    app.Configure(iris.WithConfiguration(iris.Configuration{
        Charset: "UTF-8",
    }))

    //错误配置
    //未发现错误
    app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "errmsg": iris.StatusNotFound,
            "msg":    " not found ",
            "data":   iris.Map{},
        })
    })

    app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "errmsg": iris.StatusInternalServerError,
            "msg":    " interal error ",
            "data":   iris.Map{},
        })
    })
}