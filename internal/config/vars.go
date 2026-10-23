package config

var (
	IS_DEV        string
	DB_HOST       string
	DB_USER       string
	DB_PORT       string
	DB_NAME       string
	DB_PASS       string
	ALLOW_ORIGINS string
)

func NewVars() {
	IS_DEV = GetEnv("IS_DEV", "true")
	DB_HOST = GetEnv("DB_HOST", "")
	DB_USER = GetEnv("DB_USER", "")
	DB_PORT = GetEnv("DB_PORT", "")
	DB_NAME = GetEnv("DB_NAME", "")
	DB_PASS = GetEnv("DB_PASS", "")
	ALLOW_ORIGINS = GetEnv("ALLOW_ORIGINS", "")

}
