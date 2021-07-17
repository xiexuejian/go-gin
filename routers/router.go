package routers

import (
	"github.com/gin-gonic/gin"
	. "go-gin/controller"
	"html/template"
	"log"
	"net/http"
	"time"
)

//	初始化路由
//	对应官方文档的API例子
func InitRouter(request *gin.Engine) {

	//	1、路由参数
	request.GET("/testGet/:name", TestGet)
		//	全匹配
	request.GET("/testGet/:name/*action", TestGet02)
	request.GET("/testGet03/:name/action", TestGet03)

	//	2、查询字符串参数
	request.GET("/testStringParam", TestGet04)

	//	3、Multipart/Urlencoded 表单
	request.POST("/formPost", FormPost)

	//	4、query（get）+post 表单
	request.POST("/getAndPost", GetAndPost)

	//	5、Map 作为查询字符串或 post表单 参数
	request.POST("/mapPost", MapAndPost)

	//	6、上传文件
		// 给表单限制上传大小 (默认是 32 MiB)
	request.MaxMultipartMemory = 8 << 20 // 8 MiB
		//	单个文件
	request.POST("/simpleFile", SimpleFile)
		//	多个文件
	request.POST("/multiFile", MultiFile)
	//	7、路由分组
	group := request.Group("/group")
	{
		group.POST("/url", RouteGroup)
		group.POST("/url2", RouteGroup)
	}

	//	8、默认的没有中间件的空白 Gin
	request.GET("/space",SpaceGin)

	//	9、使用中间件
	g := gin.New()
	UseMiddleware(g)

	//	10、如何写入日志文件-------在gin_log包下

	//	11、模型绑定和验证
	request.POST("/bindJson",BindJson)
	request.POST("/bindXml",BIndXml)
	request.POST("/bindHtml",BindHtml)

	//	12、自定义验证器
	request.GET("/bookable",GetBookAble)

	//	13、只绑定查询字符串（Only Bind Query String）
	request.Any("/onlyBindQueryString",OnlyBindQueryString)

	//	14、绑定Get参数或者Post参数
	request.GET("/bindGetOrPost",BindGetOrPost)

	//	15、绑定URI
	request.GET("/:name/:id",BindUri)

	//	16、绑定Header
	request.GET("/header",BindHeader)

	//	17、绑定HTML复选框
	request.LoadHTMLGlob("views/*")
	request.GET("/header", IndexHandler)
	request.POST("/header", FormHandler)

	//	18、Multipart/Urlencoded绑定
	request.POST("/profile",BindMultipart)

	//	19、XML、JSON、YAML和ProtoBuf 渲染
		// gin.H 是 map[string]interface{} 的快捷写法
	request.GET("/someJson",SomeJson)
	request.GET("/moreJson",MoreJson)
	request.GET("/someXml",SomeXml)
	request.GET("/someYaml",SomeYaml)
	request.GET("/someProtoBuf",SomeProtoBuf)

	//	20、SecureJSON
	request.GET("/secureJson",SecureJson)

	//	21、JSONP
	request.GET("/jsonp",Jsonp)

	//	22、AsciiJSON
	request.GET("/asciiJson",AsciiJSON)

	//	23、PureJSON
	request.GET("/pjson",JsonOne)
	request.GET("/pureJson",PureJSON)

	//	24、提供静态文件
	request.Static("/assets", "./assets")
	request.StaticFS("/more_static", http.Dir("my_file_system"))
	request.StaticFile("/favicon.ico", "./resources/favicon.ico")

	//	25、从文件提供数据
	request.GET("/local/file",LocalFile)
	request.GET("/fs/file",FsFile)

	//	26、从reader提供数据
	request.GET("/someDataFromReader",SomeDataFromReader)

	//	27、HTML 渲染
		//request.LoadHTMLGlob("templates/*")
		//request.LoadHTMLFiles("templates/template1.html","templates/template2.html")
	request.GET("/index",Html)

	//	28、多模板
	//	Gin允许默认只使用一个 html.Template 。查看 多模板渲染 的使用详情，类似 go 1.6 block template 。

	//	29、重定向
	request.GET("/test01",Redirect)
	request.GET("/test02",PostRedirect)
	//	发出路由器重定向，使用HandleContext如下：
	request.GET("/test03", func(c *gin.Context) {
		c.Request.URL.Path = "/test04"
		request.HandleContext(c)
	})
	request.GET("/test04",TestRedirect)

	//	30、自定义中间件
	r := gin.New()
	r.Use(Logger())
	r.GET("/selfDefineMid",SelfDefineMid)

	//	31、使用Using BasicAuth() 中间件

		// 在组中使用 gin.BasicAuth() 中间件
		// gin.Accounts 是 map[string]string 的快捷写法
	routerGroup := request.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
		// /admin/secrets 结尾
		// 点击 "localhost:8080/admin/secrets
	routerGroup.GET("/secrets",UseBasicAuth)

	//	32、中间件中使用Goroutines
	request.GET("/long_async",LongAsync)
	request.GET("/long_sync",LongSync)

	//	33、自定义HTTP配置

	// 第一种
	//_ = http.ListenAndServe(":8080", request)
	// 第二种
	//s := &http.Server{
	//	Addr: ":8080",
	//	Handler: request,
	//	ReadTimeout: 10 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//_ = s.ListenAndServe()

	//	34、支持 Let's Encrypt

	//	35、使用Gin运行多种服务
	//	看文档：https://github.com/gin-gonic/gin#run-multiple-service-using-gin

	//	36、正常的重启或停止
	//	看文档：https://github.com/gin-gonic/gin#graceful-shutdown-or-restart

	//	37、使用模板构建单个二进制文件
	//	看文档:https://github.com/gin-gonic/gin#build-a-single-binary-with-templates

	//	38、使用自定义结构绑定表单数据
	request.GET("/getb", GetDataB)
	request.GET("/getc", GetDataC)
	request.GET("/getd", GetDataD)

	//	39、尝试将 body 绑定到不同的结构中
	request.POST("/shouldBind",ShouldBindDemo)
	request.POST("/shouldBindBodyWith",ShouldBindBodyWithDemo)

	//	40、http2 服务器推送
	request.Static("/assets", "./assets")
	var html = template.Must(template.New("https").Parse(`
		<html>
		<head>
		<title>Https Test</title>
		<script src="/assets/app.js"></script>
		</head>
		<body>
		<h1 style="color:red;">Welcome, Ginner!</h1>
		</body>
		</html>
		`))
	request.SetHTMLTemplate(html)

	request.GET("/http2",HttpPush)
	//request.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")

	//	41、自定义路由日志的格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string,
		nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath,
			handlerName, nuHandlers)
	}

	//	42、设置并获Cookie
	request.GET("/cookie",GainCookie)

	//	43、测试
	//	net/http/httptest包是http测试的首选方式。
	//	官方文档：https://github.com/gin-gonic/gin#testing
}


func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置简单的变量
		c.Set("example", "12345")
		// 在请求之前
		c.Next()
		// 在请求之后
		latency := time.Since(t)
		log.Print(latency)
		// 记录发送状态
		status := c.Writer.Status()
		log.Println(status)
	}
}