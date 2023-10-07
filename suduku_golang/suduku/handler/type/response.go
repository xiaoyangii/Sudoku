package pack

import "github.com/bytedance/sonic"

type SudoGetResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSudoResponse() *SudoGetResponse {
	return &SudoGetResponse{}
}

func (s *SudoGetResponse) SetSudoResponse(code int, message string, data interface{}) {
	s.Code = code
	s.Message = message
	s.Data = data
}

func PackSudoResponse(resp *SudoGetResponse) []byte {
	res, _ := sonic.Marshal(*resp)
	return res
}
