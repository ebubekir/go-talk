package firebase

import (
	"context"
	"errors"
	firebaseGo "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"google.golang.org/api/option"
)

type App struct {
	ProjectId   string `json:"project_id"`
	FirebaseApp *firebaseGo.App
	AuthClient  *auth.Client
}

func NewFirebaseApp(projectId string, credentials string) (*App, error) {
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(credentials))

	app, err := firebaseGo.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("firebase.NewApp: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("firebase.Auth: %v", err)
	}

	return &App{
		FirebaseApp: app,
		ProjectId:   projectId,
		AuthClient:  authClient,
	}, nil
}

func (f *App) CreateUser(appUrl string, email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		EmailVerified(false)

	user, err := f.AuthClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	actionCodeSettings := &auth.ActionCodeSettings{
		URL: appUrl + "/auth/verify-email",
	}

	_, err = f.AuthClient.EmailVerificationLinkWithSettings(context.Background(), user.UserInfo.Email, actionCodeSettings)
	if err != nil {
		return nil, fmt.Errorf("error sending verification email: %v", err)
	}

	return user, nil
}

func (f *App) DeleteUser(firebaseId string) error {
	err := f.AuthClient.DeleteUser(context.Background(), firebaseId)
	if err != nil {
		return errors.New("error deleting user")
	}
	return nil
}
