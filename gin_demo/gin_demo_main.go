package main

import (
	"encoding/json"
	"fmt"
	"gin_demo/logic"
	"gin_demo/router"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin_demo/middleware"
)

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *ResponseMessage) SetCode(c int) {
	r.Code = c
}
func (r *ResponseMessage) SetMessage(m string) {
	r.Message = m
}
func (r *ResponseMessage) SetData(d any) {
	r.Data = d
}
func (r *ResponseMessage) String() string {
	d, e := json.Marshal(r)
	if e != nil {
		fmt.Println("encode fail, e: ", e)
		return ""
	}
	return string(d)
}

type Test1Response struct {
	A int     `json:"a"`
	B float32 `json:"b"`
	C string  `json:"c"`
}

func NewTest1Response() *Test1Response {
	return &Test1Response{
		A: 123,
		B: 1.111,
		C: "this test1 response",
	}
}

func BasicRun() {
	port := 9091
	fmt.Println("---- run go-gin basic run -----")
	r := gin.New()
	gin.SetMode(gin.DebugMode) //gin.ReleaseMode ; gin.
	r.Use(middleware.PreProcess())

	//path: ip:port/busiV1/
	g1 := r.Group("busiV1")
	{
		// ip:port/busiV1/test1/
		g1.GET("test1", func(c *gin.Context) {
			busiRsp := NewTest1Response()
			resp := &ResponseMessage{}
			resp.SetCode(1000)
			resp.SetMessage("succ")
			resp.SetData(busiRsp)

			c.String(http.StatusOK, resp.String())
		})
	}
	{
		// ip:port/busiV1/11/
		g1.GET("11", func(c *gin.Context) {
			busiRsp := NewTest1Response()
			resp := &ResponseMessage{}
			resp.SetCode(1000)
			resp.SetMessage("is ok")
			resp.SetData(busiRsp)

			c.String(http.StatusOK, resp.String())
		})
	}

	g2 := r.Group("busiV2")
	{
		g2.POST("test2", func(c *gin.Context) {
			c.String(http.StatusOK, "test2.")
		})
	}

	//
	logicNode := new(logic.LogicEntry)
	router.SetUrl(r, logicNode)
	r.Run(":" + strconv.Itoa(port))
}
func main() {
	BasicRun()
}
