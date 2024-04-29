package models

type Config struct {
	HTTPPort int         `json:"http_port"`
	GRPCPort int         `json:"grpc_port"`
	Mongo    MongoConfig `json:"mongo"`
	Redis    RedisConfig `json:"redis"`
}

type MongoConfig struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

type RedisConfig struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
