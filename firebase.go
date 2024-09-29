package main

import (
	"context"
	"path/filepath"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	firestoreClient *firestore.Client
	authClient      *auth.Client
	storageClient   *storage.Client
)

func configSetup() (*firestore.Client, *auth.Client, *storage.Client, error) {
	// Buscar el archivo de credenciales en el directorio actual
	matchingPattern := "./service_account.json"
	matches, err := filepath.Glob(matchingPattern)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(matches) == 0 {
		return nil, nil, nil, err
	}

	// Utilizar el primer archivo coincidente (puedes ajustar esto según tus necesidades)
	pathToCredentials := matches[0]
	opt := option.WithCredentialsFile(pathToCredentials)

	// Inicializar Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, nil, nil, err
	}

	// Inicializar cliente de Firestore
	firestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	// Inicializar cliente de Auth de Firebase
	authClient, err = app.Auth(context.Background())
	if err != nil {
		return nil, nil, nil, err
	}

	// Inicializar cliente de almacenamiento en la nube (Google Cloud Storage) usando las mismas credenciales
	storageClient, err = storage.NewClient(context.Background(), opt)
	if err != nil {
		return nil, nil, nil, err
	}

	return firestoreClient, authClient, storageClient, nil
}
