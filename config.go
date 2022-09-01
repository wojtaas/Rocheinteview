package rocheinteview

type Config struct {
	HTTPPort    int    `envconfig:"HTTP_PORT" default:"40000"`
	GRPCPort    int    `envconfig:"GRPC_PORT" default:"30000"`
	ServiceName string `envconfig:"SERVICE_NAME" default:"ROCHE"`

	Env string `envconfig:"env"`
}
