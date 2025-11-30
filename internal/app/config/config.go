package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/lpernett/godotenv"
	"os"
	"sso/internal/app/constants"
	"time"
)

type Values struct {
	Namespace     string        `envconfig:"KUBERNETES_NAMESPACE" required:"false"`
	TokenLifeTime time.Duration `envconfig:"TOKEN_LIFE_TIME" required:"true"`
	GRPCServer    *GRPCServerConfig
	DbPayments    *DBConfig
}

type GRPCServerConfig struct {
	Port             int32         `envconfig:"GRPC_SERVER_PORT" required:"true"`
	KeepaliveTimeout time.Duration `envconfig:"GRPC_SERVER_KEEPALIVE_TIMEOUT" required:"true"`
}

type DBConfig struct {
	DBName   string `envconfig:"DB_AUTH_NAME" required:"true"`
	User     string `envconfig:"DB_AUTH_USER" required:"true"`
	Password string `envconfig:"DB_AUTH_PASSWORD" required:"true"`
	Hostname string `envconfig:"DB_AUTH_HOSTNAME" required:"true"`
	SSLMode  string `envconfig:"DB_AUTH_SSLMODE" required:"false"`
	Port     int32  `envconfig:"DB_AUTH_PORT" required:"true"`

	MaxOpenConns                    int           `envconfig:"DB_AUTH_MAX_OPEN_CONNS" required:"true"`
	MaxIdleConns                    int           `envconfig:"DB_AUTH_MAX_IDLE_CONNS" required:"true"`
	MaxLifeTimeConns                time.Duration `envconfig:"DB_AUTH_MAX_LIFETIME_CONNS" required:"true"`
	StatementTimeout                time.Duration `envconfig:"DB_AUTH_STATEMENT_TIMEOUT" required:"false"`
	IdleInTransactionSessionTimeout time.Duration `envconfig:"DB_AUTH_IDLE_IN_TRANSACTION_SESSION_TIMEOUT" required:"false"`
	LockTimeout                     time.Duration `envconfig:"DB_AUTH_LOCK_TIMEOUT" required:"false"`
}

func New() (*Values, error) {
	err := LoadEnvFile()
	if err != nil {
		return nil, err
	}

	cfg := &Values{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadEnvFile() error {
	if needUseLocalEnvFile() {
		err := godotenv.Load(constants.DefaultEnvFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func needUseLocalEnvFile() bool {
	for _, arg := range os.Args {
		if arg == constants.UseLocalEnvFileArg {
			return true
		}
	}
	return false
}
