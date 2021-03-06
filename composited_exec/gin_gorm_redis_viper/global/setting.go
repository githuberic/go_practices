package global

import (
	"go_practices/composited_exec/gin_gorm_redis_viper/pkg/setting"
	"time"
)

//服务器配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//数据库配置
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettingS struct {
	Addr                  string
	Password              string
	DB                    int
	DialConnectionTimeout time.Duration
	DialReadTimeout       time.Duration
	DialWriteTimeout      time.Duration
	IdleTimeout           time.Duration
}

//定义全局变量
var (
	ServerSetting   *ServerSettingS
	DatabaseSetting *DatabaseSettingS
	RedisSetting    *RedisSettingS
)

//读取配置到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}
	return nil
}
