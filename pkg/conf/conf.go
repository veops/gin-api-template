package conf

var (
	Cfg = &ConfigYaml{
		Mode: "debug",
		Http: &HttpConfig{
			Host: "0.0.0.0",
			Port: 80,
		},
		Log: &LogConfig{
			Level:         "info",
			MaxSize:       100, // megabytes
			MaxBackups:    5,
			MaxAge:        15, // 15 days
			Compress:      true,
			Path:          "app.log",
			ConsoleEnable: true,
		},
		Auth: &Auth{
			Custom: map[string]string{},
		},
	}
)

type HttpConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
	MaxIdle  int    `yaml:"maxIdle"`
	PoolSize int    `yaml:"poolSize"`
}

type MysqlConfig struct {
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type KV struct {
	Key   string
	Value string
}

type AclConfig struct {
	Url           string `yaml:"url"`
	AppId         string `yaml:"appId"`
	SecretKey     string `yaml:"secretKey"`
	ResourceNames []*KV  `yaml:"resourceNames"`
}

type LogConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
	// MaxSize max size of single file, unit is MB
	MaxSize int `yaml:"maxSize"`
	// MaxBackups max number of backup files
	MaxBackups int `yaml:"maxBackups"`
	// MaxAge max days of backup files, unit is day
	MaxAge int `yaml:"maxAge"`
	// Compress whether compress backup file
	Compress bool `yaml:"compress"`
	// Format
	Format string `yaml:"format"`
	// Console output
	ConsoleEnable bool `yaml:"consoleEnable"`
}

type Auth struct {
	Acl    *AclConfig        `yaml:"acl"`
	Custom map[string]string `yaml:"custom"`
}

type ConfigYaml struct {
	Mode      string       `yaml:"mode"`
	Http      *HttpConfig  `yaml:"http"`
	Log       *LogConfig   `yaml:"log"`
	Redis     *RedisConfig `yaml:"redis"`
	Mysql     *MysqlConfig `yaml:"mysql"`
	Auth      *Auth        `yaml:"auth"`
	SecretKey string       `json:"secretKey"`
}
