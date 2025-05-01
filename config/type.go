package config

type DBConfig struct {
	Host     string `josn:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type ServerConfig struct {
	HttpPort          int    `json:"httpPort"`
	Secret            string `json:"secret"`
	AuthExpMin        int    `json:"authExpMin"`
	AuthRefreshExpMin int    `json:"authRefreshExpMin"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type Config struct {
	DBConfig     DBConfig     `json:"dbConfig"`
	ServerConfig ServerConfig `json:"serverConfig"`
	RedisConfig  RedisConfig  `json:"redisConfig"`
}
