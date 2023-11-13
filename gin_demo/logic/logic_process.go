package logic

type LogicReqMessage struct {
	Id string `json:"id"`
}

type LogicRespMessage struct {
	Name string 	`json:"name"`
	Age int 		`json:"age"`
	Address string `json:"addr"`
}