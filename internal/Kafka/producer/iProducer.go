package producer

type Producer interface {
	Send(message Message) error
}

type Message struct {
	CmdType CMDType
	Info    string
}

type CMDType int

const (
	CREATE CMDType = iota
	UPDATE
	DELETE
)
