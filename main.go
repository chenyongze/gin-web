package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// func1: 处理最基本的GET
func func1 (c *gin.Context)  {
    // 回复一个200OK,在client的http-get的resp的body中获取数据
    c.String(http.StatusOK, "test1 OK")
}
// func2: 处理最基本的POST
func func2 (c *gin.Context) {
    // 回复一个200 OK, 在client的http-post的resp的body中获取数据
    c.JSON(http.StatusOK, gin.H{
        "status":0,
        "data":gin.H{
            "a":"b",
        },
    })
}

func func3 (c *gin.Context)  {
    //回复一个200OK,在client的http-get的resp的body中获取数据
    //c.String(http.StatusOK, "v1 group")
    c.JSON(http.StatusOK, gin.H{
        "l":"group v1",
    })
}
func main(){
    // 注册一个默认的路由器
    router := gin.Default()
    // 最基本的用法
    router.GET("/test1", func1)
    router.GET("/test2", func2)

    groupx := router.Group("v1")
    groupx.GET("/test1", func3)
    // 绑定端口是8888
    router.Run(":8086")
}