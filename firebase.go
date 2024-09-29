package main

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	authClient *auth.Client
)

func configSetup() (*auth.Client, error) {
	// Buscar el archivo de credenciales en el directorio actual
	matchingPattern := "./service_account.json"
	matches, err := filepath.Glob(matchingPattern)
	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return nil, err
	}

	// Utilizar el primer archivo coincidente (puedes ajustar esto según tus necesidades)
	pathToCredentials := matches[0]
	opt := option.WithCredentialsFile(pathToCredentials)

	// Inicializar Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	// Inicializar cliente de Auth de Firebase
	authClient, err = app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return authClient, nil
}
