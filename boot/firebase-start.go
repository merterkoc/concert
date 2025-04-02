package boot

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"gilab.com/pragmaticreviews/golang-gin-poc/config"
	"google.golang.org/api/option"
)

func FirebaseStart() *firebase.App {

	env := config.GetConfig()
	var opt option.ClientOption
	if env != "stage" {
		if env != "dev" {
			opt = option.WithCredentialsFile("./firebase-config/prod/serviceAccountKey.json")
		} else {
			opt = option.WithCredentialsFile("./firebase-config/dev/serviceAccountKey.json")
		}

		aapp, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("error initializing app: %v", err)
			panic(err) // "throwing" the error
		}
		return aapp
	} else {
		log.Println("Stage mode: Please export FIREBASE_AUTH_EMULATOR_HOST= for local development")

		conf := &firebase.Config{}
		aapp, err := firebase.NewApp(context.Background(), conf)
		if err != nil {
			log.Printf("error initializing app: %v", err)
			panic(err) // "throwing" the error
		}
		return aapp
	}

}
