package rocheinteview

const (
	PingMessage    = "message"
	PostmanEchoURL = "https://postman-echo.com/get?message=%s"
)

type PingResponse struct {
	Echo      string `json:"echo"`
	Timestamp int64  `json:"timestamp"`
	Env       string `json:"env"`
	Version   string `json:"version"`
}
