package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	ginUtil "gin_demo/gin_util"
)

const (
	CtxClientIp = "ctx_client_ip"
)
func PreProcess() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(CtxClientIp, c.Request.Header.Get("X-REAL-IP"))
		requestID := c.GetHeader(ginUtil.CtxRequestID) 
		if requestID == "" {
			requestID = ginUtil.GenerateRequestID()
		}

		c.Writer.Header().Set(ginUtil.CtxRequestID, requestID)
		c.Set(ginUtil.CtxRequestID, requestID)


		path := c.Request.URL.Path
		method := c.Request.Method
		if path == "/test" || path == "/metrics" {
			return
		}

		var reqBuf []byte
		if method == http.MethodPost ||
			method == http.MethodPut ||
			method == http.MethodDelete {
			reqBuf, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBuf))
		} else if method == http.MethodGet {
			m := map[string]interface{}{}
			for k, v := range c.Request.URL.Query() {
				m[k] = v[0]
			}
			reqBuf, _ = json.Marshal(m)
		}
		ginUtil.SetRequestBody(c, string(reqBuf))
		c.Next()
	}
	
}