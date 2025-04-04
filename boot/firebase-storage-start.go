package boot

import (
	"cloud.google.com/go/storage"
	"context"
	"gilab.com/pragmaticreviews/golang-gin-poc/config"
	"google.golang.org/api/option"
	"log"
)

func FirebaseStorageStart() *storage.Client {
	env := config.GetConfig()
	ctx := context.Background()
	if env != "stage" {
		if env == "dev" {
			opt := option.WithCredentialsFile("./firebase-config/dev/serviceAccountKey.json")
			client, err := storage.NewClient(ctx, opt)
			if err != nil {
				log.Fatalf("Firebase Storage client not created: %v", err)
			}
			return client
		} else {
			opt := option.WithCredentialsFile("./firebase-config/prod/serviceAccountKey.json")
			client, err := storage.NewClient(ctx, opt)
			if err != nil {
				log.Fatalf("Firebase Storage client not created: %v", err)
			}
			return client
		}
	}
	panic("Stage mode: Please export FIREBASE_AUTH_EMULATOR_HOST= for local development")
}
