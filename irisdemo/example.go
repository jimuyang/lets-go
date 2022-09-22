package main

import (
	"regexp"
	"strconv"

	"github.com/kataras/iris/v12/hero"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())
	// 或者使用
	// app := iris.Default()

	// Method:   GET
	// Resource: http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// ctx.JSON(iris.Map{"message": "Hello Iris!"})
		ctx.WriteString("Hello Iris!")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello/{name}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("name"))
	})

	// app.Get("/someGet", getting)
	// app.Post("/somePost", posting)
	// app.Put("/somePut", putting)
	// app.Delete("/someDelete", deleting)
	// app.Patch("/somePatch", patching)
	// app.Head("/someHead", head)
	// app.Options("/someOptions", options)

	// demo: 从ctx获取入参
	app.Get("/users/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		app.Logger().Println("id: ", id)
	})

	// demo: 使用内置的macro regexp prefix suffix contains min max range
	app.Get("/profile/{name:alphabetical max(5)}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.WriteString(name)
	})

	// demo：自己注册宏
	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)
	app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)

	app.Get("/coord/{lat:string coordinate()}/{lon:string coordinate()}", func(ctx iris.Context) {
		ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
	})

	// 2个参数的macro
	app.Macros().Get("int").RegisterFunc("range", func(min, max int) func(int) bool {
		return func(i int) bool {
			return i >= min && i <= max
		}
	})
	app.Get("/limitnum/{num:int range(1,100)}", func(ctx iris.Context) {
		num := ctx.Params().GetIntDefault("num", 1)
		ctx.WriteString(strconv.Itoa(num))
	})
	// demo: use hero
	helloHandler := hero.Handler(hello)
	app.Get("/hero/{to:string}", helloHandler)

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func hello(to string) string {
	return "hello " + to
}
