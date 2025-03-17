package firebase

import (
	"context"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/gommon/log"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var AuthClient *auth.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("secrets/cassist-eaf16-firebase-adminsdk-fbsvc-f74bc3f9e2.json")

	var err error
	FirebaseApp, err = firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Panicf("Could not initialize firebase app: %v", err)
	}

	AuthClient, err = FirebaseApp.Auth(context.Background())
	if err != nil {
        log.Panicf("Could not initialize firebase auth client: %v", err)
	}

}
