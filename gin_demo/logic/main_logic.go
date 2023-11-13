package logic

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogicEntry struct {
}

func (e *LogicEntry) GetInfo(c *gin.Context, req *LogicReqMessage) (*LogicRespMessage, error) {
	response := &LogicRespMessage{}
	if req == nil {
		return nil, errors.New("input param is nil")
	}
	if len(req.Id) <= 0 {
		fmt.Println("not invalid id")
		return nil, ErrorGetInfo
	}

	// do business......

	response.Address = "shenzhen"
	response.Age = 18
	response.Name = "huawei"
	return response, nil
}
