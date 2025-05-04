package firebaseApp

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type Adapter struct {
	firebaseApp *firebase.App
}

func NewFirebaseApp(path string) (*Adapter, error) {
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return &Adapter{firebaseApp: app}, nil
}

func (app *Adapter) GetUser(uid string) (*auth.UserRecord, error) {
	ctx := context.Background()

	client, err := app.firebaseApp.Auth(ctx)
	if err != nil {
		return nil, err
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (app *Adapter) GetUserByEmail(email string) (*auth.UserRecord, error) {
	ctx := context.Background()

	client, err := app.firebaseApp.Auth(ctx)
	if err != nil {
		return nil, err
	}

	u, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
