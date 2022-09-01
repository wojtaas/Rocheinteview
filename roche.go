package rocheinteview

const PingMessage = "message"

type PingResponse struct {
	Echo      string `json:"echo"`
	Timestamp int64  `json:"timestamp"`
	Env       string `json:"env"`
	Version   string `json:"version"`
}
