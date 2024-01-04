package domain

type LoggerPayload struct {
	Time string
	Loc  string
	Msg  string
	Req  any
}

type Logger interface {
	Error(payload *LoggerPayload)
	Info(payload *LoggerPayload)
}
