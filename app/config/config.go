package config

var (
	GoTest = false
	Path   = "./config"
	config = new(Config)
)

type Config struct {
	e *Environment
	a *Application
}

const (
	Dev  = "dev"
	Pre  = "pre"
	Prod = "prod"
)

type Environment struct {
	isDev  bool
	isPre  bool
	isProd bool
}

type Application struct {
	App   *App
	Http  *Http
	Https *Https
	Pprof *Pprof
	Data  *Data
	Log   *Log
	Mysql *MysqlShards
	Redis *RedisShards
}

type App struct {
	Version  string
	Mode     string
	Language string
}

type Http struct {
	Host string
	Port int
}

type Https struct {
	Host     string
	Port     int
	CertFile string
	KeyFile  string
}

type Data struct {
	Path string
}

type Log struct {
	Level string
	Path  string
}

type Pprof struct {
	Enable bool
	Host   string
	Port   int
}

type MysqlShards struct {
	Debug  bool
	DDL    bool
	Tables int
	Common *MysqlConfig
	Shards []*MysqlConfig
}

type RedisShards struct {
	SessionDB int
	LogicDB   int
	Shards    []*RedisConfig
}

type MysqlConfig struct {
	Dsn          string
	MaxOpen      int // <= 0 means unlimited
	MaxIdleCount int // zero means defaultMaxIdleConns; negative means 0
	MaxIdleTime  int // maximum amount of time a connection may be idle before being closed
	MaxLifetime  int // maximum amount of time a connection may be reused
}

type RedisConfig struct {
	Url            string
	ConnectTimeout int // specifies the timeout for connecting to the Redis server
	KeepAlive      int // specifies the keep-alive period for TCP connections to the Redis server
	ReadTimeout    int // specifies the timeout for reading a single command reply
	WriteTimeout   int // specifies the timeout for writing a single command
	MaxIdle        int // Maximum number of idle connections in the pool.
	MaxActive      int // When zero, there is no limit on the number of connections in the pool
	IdleTimeout    int // the value is zero, then idle connections are not closed
	MaxLifetime    int // If the value is zero, then the pool does not close connections based on age
}
