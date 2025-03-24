package internal

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Environment struct {
	TicketMasterAPIToken string
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Environment{
		TicketMasterAPIToken: os.Getenv("TICKET_MASTER_API_TOKEN"),
		AppPort:              os.Getenv("APP_PORT"),
	}
}
