package cfg

type Config struct {
	DBName  string `env:"DB_NAME,required"`
	DBUser  string `env:"DB_USER,required"`
	DBPass  string `env:"DB_PASSWORD,required"`
	DBHost  string `env:"DB_HOST,required"`
	DBPort  string `env:"DB_PORT,required"`
	AppPort string `env:"APP_PORT,required"`
}
