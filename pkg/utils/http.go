package utils

type Response struct {
	Code    int8        `json:"code"`
	Msg     string      `json:"msg"`
	Records interface{} `json:"records"`
}
