package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubelog/cmd"
)

func Hededer(context *gin.Context)  {
	context.HTML(200, "index.html", cmd.PodAllNameder())
}

//func NamePod(context *gin.Context)  {
//	context.HTML(200, "pods.html", cmd.Postname("redis"))
//}

func Getname(context *gin.Context) {
	name := context.Param("name")
	context.HTML(200, "pods.html", cmd.Postname(name))
}

//func Download(context *gin.Context) {
//	name := context.Param("name")
//	context.HTML(200, "Download.html", cmd.LogSeveDownload())
//}

func Download(context *gin.Context)  {
	Pod := context.Param("name")
	cmd.LogSeveDownload(Pod)
	context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", Pod + ".log"))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	context.Writer.Header().Add("Content-Type", "application/octet-stream")
	context.File("./tmp/" + Pod + ".log")
}

func Setup() *gin.Engine {
	r := gin.Default() //返回默认的路由引擎

	r.LoadHTMLGlob("./html/*")

	r.GET("/", Hededer)

	r.GET("/pods/:name", Getname)

	r.GET("/Download/:name", Download)


	//r.POST("/book", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"ms":"POST"})
	//})

	//启动服务
	r.Run(":8686")  //默认不填写端口是8080

	return r
}
