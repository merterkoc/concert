package internal

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gilab.com/pragmaticreviews/golang-gin-poc/config"
	"github.com/joho/godotenv"
)

type Environment struct {
	TicketMasterAPIToken string
	DBString             string
	AppPort              string
}

type EnvService struct {
	Env *Environment
	mu  sync.Once
}

var envServiceInstance *EnvService

func GetEnvServiceInstance() *EnvService {
	if envServiceInstance == nil {
		envServiceInstance = &EnvService{}
		envServiceInstance.mu.Do(func() {
			envServiceInstance.Env = loadEnv()
		})
	}
	return envServiceInstance
}

func loadEnv() *Environment {

	env := config.GetConfig()

	fmt.Printf("Selected environment: %e\n", env)
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Environment{
		TicketMasterAPIToken: os.Getenv("TICKET_MASTER_API_TOKEN"),
		AppPort:              os.Getenv("APP_PORT"),
		DBString:             os.Getenv("DB_STRING"),
	}
}
