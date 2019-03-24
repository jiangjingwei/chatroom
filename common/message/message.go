package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMesType"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息数据
}

type LoginMes struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
