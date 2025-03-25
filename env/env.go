package env

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const(
	EnvLocal = "local"
	EnvDev = "dev"
	EnvProd = "prod"
)

type Env struct{
	Env 				string 					`env:"ENV,required"`
	Http				Http					`env-required:"true"`
	PgSql 				PgSql 					`env-required:"true"`
	Kafka				Kafka					`yaml:"kafka" env-required:"true"`
	OpenaiClient 		OpenaiClient			`yaml:"openai" env-required:"true"`
}

type Http struct {
	Port	string			`env:"PORT,required"`
}

type PgSql struct {
	Host     string `env:"POSTGRES_HOST,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	DbName   string `env:"POSTGRES_DB,required"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	SSLMode  string `env:"POSTGRES_SSLMODE" env-default:"disable"`
	URI 	string 	`env:"POSTGRES_URI"`
}

type Kafka struct{
	TopicsRead 		[]string 	`yaml:"topics_read" env-required:"true"`
	TopicWrite 		[]string 	`yaml:"topics_write" env-required:"true"`
	GroupID 		string 		`yaml:"group_id" env-required:"true"`
	Brokers			[]string	`yaml:"brokers" env-required:"true"`
	Region 			string 		`yaml:"region" env-required:"true"`
	AWSProfile 		string 		`yaml:"awsProfile" env-required:"true"`
	// For other auth type
	User 			string 		`env:"KAFKA_USER,required"`
	Pass 			string 		`env:"KAFKA_PASS,required"`
}

type OpenaiClient struct {
    BaseURL     	string        	`yaml:"base_url" `
	TestURL			string			`yaml:"test_url" `
    APIKey      	string        	`env:"OPENAI_KEY,required"`
    APISecret   	string        	`env:"OPENAI_SECRET,required"`
	Timeout			time.Duration	`yaml:"timeout" env-required:"true"`
	RetriesCount	int				`yaml:"retries_count" env-required:"true"`
}

func (c *Http) GetPort() string {
    if envPort := os.Getenv("PORT"); envPort != "" {
        return envPort
    }
    return c.Port
}


func MustLoad()	*Env {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found")
	}

	path := fetchConfigFlag()

	if path == ""{
		panic("config path is empty")
	}

	if _,err := os.Stat(path); os.IsNotExist(err){
		panic("config path dose not exist: " + path)
	}

	var config Env

	if err:=cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	if err:=cleanenv.ReadEnv(&config); err != nil {
		panic("failed to read env: " + err.Error())
	}

	return &config
}

func fetchConfigFlag() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file") 
	flag.Parse()

	if res == ""{
		res = os.Getenv("CONFIG_PATH")	
	}

	return res
}