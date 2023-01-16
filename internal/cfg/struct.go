package cfg

type Config struct {
	DBName           string `env:"DB_NAME,required"`
	DBUser           string `env:"DB_USER,required"`
	DBPass           string `env:"DB_PASSWORD,required"`
	DBHost           string `env:"DB_HOST,required"`
	DBPort           string `env:"DB_PORT,required"`
	AppPort          string `env:"APP_PORT,required"`
	RabbitUser       string `env:"RMQ_USER,required"`
	RabbitPass       string `env:"RMQ_PASS,required"`
	RabbitHost       string `env:"RMQ_HOST,required"`
	RabbitPort       string `env:"RMQ_PORT,required"`
	KavehNegarAPIKey string `env:"KAVEH_NEGAR_API_KEY,required"`
	KavehNegarURL    string `env:"KAVEH_NEGAR_URL,required"`
}
