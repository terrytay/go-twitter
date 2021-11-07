package db

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Firebase struct {
	c *firestore.Client
}

func NewFirebase() *Firebase {
	f := &Firebase{nil}
	f.init()
	return f
}

func (f *Firebase) init() error {
	ctx := context.Background()

	opt := option.WithCredentialsFile("firebase-key.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	f.c = client
	return nil
}

func (f *Firebase) FindOne(ctx context.Context, collection string, id string) (interface{}, error) {
	doc, err := f.c.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	m := doc.Data()
	return m, nil
}

func (f *Firebase) FindAll(ctx context.Context, collection string) ([]interface{}, error) {
	docs, err := f.c.Collection(collection).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var vals []interface{}
	for _, v := range docs {
		vals = append(vals, v.Data())
	}
	return vals, nil
}

func (f *Firebase) Create(ctx context.Context, collection string, data Data) error {
	_, err := f.c.Collection(collection).Doc(data.GetId()).Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) Update(ctx context.Context, collection string, data Data) error {
	_, err := f.c.Collection(collection).Doc(data.GetId()).Set(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) Delete(ctx context.Context, collection string, id string) error {
	_, err := f.c.Collection(collection).Doc(id).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firebase) Close() error {
	return f.c.Close()
}
