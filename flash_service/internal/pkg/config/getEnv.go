package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	GRPCPort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	LogPath          string
	KafkaUrl         string
	BucketName       string
	MinioUrl         string
	MinioUser        string
	MinioPassword    string
	MinioPath        string
	Email            string
	EmailPassword    string

	DefaultOffset string
	DefaultLimit  string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "flash_sale"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "flash_sale"))
	config.LogPath = cast.ToString(getOrReturnDefaultValue("LOG_PATH", "logs/info.log"))
	config.KafkaUrl = cast.ToString(getOrReturnDefaultValue("KAFKA_URL", "q"))
	config.BucketName = cast.ToString(getOrReturnDefaultValue("BUCKET_NAME", "q"))
	config.MinioUrl = cast.ToString(getOrReturnDefaultValue("MINIO_URL", "q"))
	config.MinioUser = cast.ToString(getOrReturnDefaultValue("MINIO_USER", "q"))
	config.MinioPassword = cast.ToString(getOrReturnDefaultValue("MINIO_PASSWORD", "q"))
	config.MinioPath = cast.ToString(getOrReturnDefaultValue("MINIO_PATH", "q"))

	config.Email = cast.ToString(getOrReturnDefaultValue("EMAIL", "q"))
	config.EmailPassword = cast.ToString(getOrReturnDefaultValue("EMAIL_PASSWORD", "q"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
