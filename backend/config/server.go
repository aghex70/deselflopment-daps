package config

type ServerConfig struct {
	Grpc *GrpcConfig
	Rest *RestConfig
}

func LoadServerConfig() *ServerConfig {
	return &ServerConfig{
		Grpc: LoadGrpcConfig(),
		Rest: LoadRestConfig(),
	}
}
