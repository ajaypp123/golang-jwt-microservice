package models

type Config struct {
	HTTPPort int         `json:"http_port"`
	GRPCPort int         `json:"grpc_port"`
	Mongo    MongoConfig `json:"mongo"`
}

type MongoConfig struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}
