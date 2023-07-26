package config

import "errors"

type Constants struct {
	DebugMode                   bool `env:"DEBUG_MODE,default=false"`
	Port                        int  `env:"PORT,default=3000"`
	ExportLog                   bool `env:"EXPORT_LOG,default=true"`
	SystemShutdownTimeOutSecond int  `env:"SYSTEM_SHUTDOWN_TIMEOUT_SECOND,default=30"`
	PostgreSQLUseSSL            bool `env:"POSTGRESQL_USE_SSL,default=false"`
	RedisUseSSL                 bool `env:"REDIS_USE_SSL,default=false"`
	AwsS3PreSignDurationHour    int  `env:"AWS_S3_PRE_SIGN_DURATION_HOUR,default=2"`
}

var (
	ErrInvalidEnv = errors.New("invalid env")
)
