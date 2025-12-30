package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
)

const (
	HeaderContentTypeKey  = "Content-Type"
	HeaderContentTypeJSON = "application/json; charset=utf-8"
)

func NewConfig() (*Configuration, error) {
	var envFiles []string
	if _, err := os.Stat(".env"); err == nil {
		log.Println("found .env file, adding it to env config files list")
		envFiles = append(envFiles, ".env")
	}

	if len(envFiles) > 0 {
		err := godotenv.Overload(envFiles...)
		if err != nil {
			return nil, errors.Wrapf(err, "error while opening env config: %s", err)
		}
	}

	cfg := &Configuration{}
	ctx := context.Background()

	err := envconfig.Process(ctx, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error while loading config")
	}
	return cfg, nil
}

type (
	Configuration struct {
		Version     VersionConfig  `env:",prefix=VERSION_"`
		Log         LogConfig      `env:",prefix=LOG_"`
		RunTime     RuntimeConfig  `env:",prefix=RUNTIME_"`
		HTTP        HTTPConfig     `env:",prefix=HTTP_"`
		GRPC        GRPCConfig     `env:",prefix=GRPC_"`
		ClientsGRPC ClientsGRPC    `env:",prefix=CLIENTS_GRPC_"`
		RWDB        DBConfig       `env:",prefix=RWDB_"`
		RDB         DBConfig       `env:",prefix=RDB_"`
		Redis       RedisConfig    `env:",prefix=REDIS_"`
		LifeTime    LifeTimeConfig `env:",prefix=LIFE_TIME_"`
		ApiKey      ApiKeyConfig   `env:",prefix=API_KEY_"`
		Role        RoleConfig     `env:",prefix=ROLE_"`
	}

	VersionConfig struct {
		Number string `env:"NUMBER,default=1.0.0"`
		Build  string `env:"BUILD,default=dev"`
	}
	LogConfig struct {
		Level             string        `env:"LEVEL,default=info"`
		Batch             bool          `env:"BATCH,default=false"`
		BatchSize         int           `env:"BATCH_SIZE,default=1000"`
		BatchPollInterval time.Duration `env:"BATCH_POLL_INTERVAL,default=5s"`
	}
	RuntimeConfig struct {
		UseCPUs    int `env:"USE_CPUS,default=0"`
		MaxThreads int `env:"MAX_THREADS,default=0"`
	}
	HTTPConfig struct {
		RequestLoggingEnabled      bool          `env:"REQUEST_LOGGING_ENABLED,default=false"`
		ResponseTimeLoggingEnabled bool          `env:"RESPONSE_TIME_LOGGING_ENABLED,default=false"`
		ReadTimeout                time.Duration `env:"READ_TIMEOUT,default=30s"`
		WriteTimeout               time.Duration `env:"WRITE_TIMEOUT,default=30s"`
		IdleTimeout                time.Duration `env:"IDLE_TIMEOUT,default=30s"`
		MaxRequestBodySize         int           `env:"MAX_REQUEST_BODY_SIZE,default=4194304"`
		Network                    string        `env:"NETWORK,default=tcp"`
		Address                    string        `env:"ADDRESS,default=:8080"`
	}
	GRPCConfig struct {
		RequestLoggingEnabled      bool          `env:"REQUEST_LOGGING_ENABLED,default=false"`
		ResponseTimeLoggingEnabled bool          `env:"RESPONSE_TIME_LOGGING_ENABLED,default=false"`
		ReadTimeout                time.Duration `env:"READ_TIMEOUT,default=30s"`
		WriteTimeout               time.Duration `env:"WRITE_TIMEOUT,default=30s"`
		IdleTimeout                time.Duration `env:"IDLE_TIMEOUT,default=30s"`
		MaxRequestBodySize         int           `env:"MAX_REQUEST_BODY_SIZE,default=33554432"`
		Network                    string        `env:"NETWORK,default=tcp"`
		Address                    string        `env:"ADDRESS,default=:18080"`
	}

	ClientsGRPC struct {
		TestMessenger ClientGRPC `env:",prefix=TEST_MESSENGER_"`
	}

	ClientGRPC struct {
		Address             string        `env:"ADDRESS"`
		Port                string        `env:"PORT"`
		IdleTimeout         time.Duration `env:"IDLE_TIMEOUT"`
		InsecureSkipVerify  bool          `env:"INSECURE_SKIP_VERIFY"`
		MaxRequestBodySize  int           `env:"MAX_REQUEST_BODY_SIZE,default=4194304"`
		MaxResponseBodySize int           `env:"MAX_RESPONSE_BODY_SIZE,default=4194304"`
	}

	DBConfig struct {
		ConnectionString         string        `env:"CONNECTION_STRING,required"`
		MaxOpenConnection        int32         `env:"MAX_OPEN_CONNECTION,default=25"`
		MaxIdleConnection        int32         `env:"MAX_IDLE_CONNECTION,default=10"`
		MaxIdleConnectionTimeout time.Duration `env:"MAX_IDLE_TIMEOUT,default=300s"`
	}

	RedisConfig struct {
		ConnectionString string        `env:"CONNECTION_STRING"`
		Timeout          time.Duration `env:"TIMEOUT,default=2s"`
		CountPerUseScan  int64         `env:"COUNT_PER_USE_SCAN,default=100"`
	}

	LifeTimeConfig struct {
		Session        time.Duration `env:"SESSION,default=720h"`
		TempUser       time.Duration `env:"TEMP_USER,default=24h"`
		ForgotPassword time.Duration `env:"FORGOT_PASSWORD,default=10m"`
	}

	ApiKeyConfig struct {
		Web       string `env:"WEB,required"`
		Superuser string `env:"SUPERUSER,required"`
	}

	RoleConfig struct {
		Name RoleName `env:",prefix=NAME_"`
	}

	RoleName struct {
		Person    string `env:"PERSON,default=person"`
		Organizer string `env:"ORGANIZER,default=organizer"`
		Moderator string `env:"MODERATOR,default=moderator"`
		Superuser string `env:"SUPERUSER,default=superuser"`
	}
)

func (c ClientGRPC) GetFullAddress() string {
	return c.Address + c.Port
}
