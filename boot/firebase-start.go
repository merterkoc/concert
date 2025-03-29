package boot

import (
	"context"
	firebase "firebase.google.com/go"
	"gilab.com/pragmaticreviews/golang-gin-poc/config"
	"google.golang.org/api/option"
	"log"
)

func FirebaseStart() *firebase.App {
	env := config.GetConfig()
	var opt option.ClientOption
	if env == "dev" {
		opt = option.WithCredentialsFile("./firebase-config/dev/serviceAccountKey.json")
	} else {
		opt = option.WithCredentialsFile("./firebase-config/prod/serviceAccountKey.json")
	}
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
		panic(err) // "throwing" the error
	}
	return app
}
