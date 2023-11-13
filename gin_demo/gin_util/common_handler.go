package ginUtil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// WrapBusinessCall use func WrapBusinessCall(handler interface{}) func(ctx *gin.Context) { input parameters (is func obj) as template to define other func.
// in new other func call parameter func obj.
func WrapBusinessCall(handler interface{}) func(ctx *gin.Context) {
	BusiHandleType := reflect.TypeOf(handler)
	BusiHandleValue := reflect.ValueOf(handler)
	//
	realFcType := func(ctx *gin.Context) {}
	//
	wrappedFunc := func(args []reflect.Value) (results []reflect.Value) {
		//because of wrappedFunc is func(*gin.Context), so  args is one, gin.Context.
		beginTime := time.Now().UnixMilli()

		ctx := args[0].Interface().(*gin.Context)
		var realIN []reflect.Value
		realIN = append(realIN, args[0])
		if BusiHandleType.NumIn() == 2 {
			param := BusiHandleType.In(1)
			if param.Kind() == reflect.Ptr {
				param = param.Elem()
			}
			//
			val := reflect.New(param)
			if ctx.Request.Method == http.MethodGet {
				if err := ctx.ShouldBindQuery(val.Interface()); err != nil {
					fmt.Printf("bind to query failed, err: %v\n", err)
					ctx.AbortWithStatus(http.StatusInternalServerError)
					return nil
				}
			} else {
				// allow post body is nil
				body := GetRequestBody(ctx)
				if len(body) != 0 {
					if err := json.Unmarshal([]byte(body), val.Interface()); err != nil {
						fmt.Printf("json unmarshal to struct fail, data: %v, err: %v\n", GetRequestBody(ctx), err)

						ctx.AbortWithStatus(http.StatusInternalServerError)
						return nil
					}
				}

			}
			for _i := 0; _i < val.Elem().NumField(); _i++ {
				if val.Elem().Field(_i).Kind() == reflect.String {
					val.Elem().Field(_i).SetString(strings.Trim(val.Elem().Field(_i).String(), " "))
				}
			}
			realIN = append(realIN, val)
		}

		//call busi logic process handle.
		retValues := BusiHandleValue.Call(realIN)

		// last reponse out parameter is error type.
		retItemNums := BusiHandleType.NumOut()

		if retItemNums == 2 { //business data, and busi error.
			statusCode := http.StatusOK
			var content interface{}
			content = http.StatusText(http.StatusOK) //default is: OK

			if retValues[retItemNums-1].Interface() != nil {
				fmt.Printf("response fail, e: %v\n", StructToJsonString(retValues[retItemNums-1].Interface()))
				//
				nowTm := time.Now().UnixNano() / 1e6
				//reponse fail.
				result := CliError{
					Message:   "",
					MessageId: "",  // utils.GetLocalIp() + "_" + requestid.GetRequestID(ctx),
					Status:    201, // 默认err
					Timestamp: time.Now().UnixNano() / 1e6,
					CostTime:  nowTm - beginTime,
				}

				//result := errors.ErrServer
				errImpl, ok := retValues[retItemNums-1].Interface().(*Error)
				if ok {
					if errImpl.UserErrCode != 0 {
						result.Status = errImpl.UserErrCode
					}
					result.Message = errImpl.UserMsg
					if errImpl.HttpCode != 0 {
						statusCode = errImpl.HttpCode
					}
				} else {
					fmt.Printf("internal server, err: %v\n", retValues[retItemNums-1].Interface().(error))
				}
				content = result

			} else if BusiHandleType.NumOut() != 1 {
				//response succ.
				nowTm := time.Now().UnixNano() / 1e6
				result := CliError{
					Message:   "",
					MessageId: "",   //utils.GetLocalIp() + "_" + requestid.GetRequestID(ctx),
					Status:    1000, //succ.
					Timestamp: nowTm,
					CostTime:  nowTm - beginTime,
				}
				result.Content = retValues[0].Interface() //set buis data to result.content.
				content = result
			}
			if !ctx.IsAborted() {
				ctx.Set("ctx_status", "success")
				ctx.JSON(statusCode, content)
			}
		} else { //if return only one business data
			if !ctx.IsAborted() {
				ctx.JSON(http.StatusOK, retValues[0].Interface())
			}
		}
		ctx.Next()
		return nil
	}
	h := reflect.MakeFunc(reflect.TypeOf(realFcType), wrappedFunc)
	return h.Interface().(func(ctx *gin.Context))
}
