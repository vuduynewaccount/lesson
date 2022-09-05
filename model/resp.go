package model

type Resp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"detail"`
}



func (s *Resp) Constructor(Status int, Msg string, Data interface{}) {
  s.Status=Status;
  s.Msg=Msg;
  s.Data=Data;
}
