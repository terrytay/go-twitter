package db

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Firebase struct {
	l *log.Logger
	c *firestore.Client
}

func NewFirebase(l *log.Logger) (*Firebase, error) {
	f := &Firebase{l, nil}
	err := f.init()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *Firebase) init() error {
	ctx := context.Background()

	opt := option.WithCredentialsFile("firebase-key.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		f.l.Println("[FIREBASE]", err)
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		f.l.Println("[FIREBASE]", err)
		return err
	}

	f.c = client
	return nil
}

func (f *Firebase) FindOne(ctx context.Context, collection string, id string) (interface{}, error) {
	doc, err := f.c.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		f.l.Println("[FIREBASE]", err)
		return nil, err
	}
	m := doc.Data()
	return m, nil
}

func (f *Firebase) FindAll(ctx context.Context, collection string) ([]interface{}, error) {
	docs, err := f.c.Collection(collection).Documents(ctx).GetAll()
	if err != nil {
		f.l.Println("[FIREBASE]", err)
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
		f.l.Println("[FIREBASE]", err)
		return err
	}
	return nil
}

func (f *Firebase) Update(ctx context.Context, collection string, data Data) error {
	_, err := f.c.Collection(collection).Doc(data.GetId()).Set(ctx, data)
	if err != nil {
		f.l.Println("[FIREBASE]", err)
		return err
	}
	return nil
}

func (f *Firebase) Delete(ctx context.Context, collection string, id string) error {
	_, err := f.c.Collection(collection).Doc(id).Delete(ctx)
	if err != nil {
		f.l.Println("[FIREBASE]", err)
		return err
	}
	return nil
}

func (f *Firebase) Close() error {
	f.l.Println("[FIREBASE]", "closing db connection")
	return f.c.Close()
}
