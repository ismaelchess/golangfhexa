package repository

import (
	"context"
	"log"
	"time"

	"github.com/golangfhexa/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func newMongClient(mongoServerURL string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServerURL))
	if err != nil {
		return nil, err
	}
	//We could ping the server to test connectivity if we want

	return client, nil
}

func NewMongoRepository(mongoServerURL, mongoDb string, timeout int) (domain.Repository, error) {
	mongoClient, err := newMongClient(mongoServerURL, timeout)
	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	return repo, nil

}

func (r *mongoRepository) Store(employee *domain.Employee) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("items")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"nouser":    employee.NoUser,
			"email":     employee.Email,
			"full_name": employee.Full_name,
			"data1":     employee.Data1,
			"data2":     employee.Data2,
		},
	)
	if err != nil {
		return errors.Wrap(err, "Error writing to repository")
	}
	return nil

}
func (r *mongoRepository) Update(employee *domain.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("items")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"nouser": employee.NoUser},
		bson.D{
			{"$set", bson.D{{"full_name", employee.Full_name}, {"data1", employee.Data1}, {"data2", employee.data2}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoRepository) Find(nouser string) (*domain.Employee, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	employee := &domain.Employee{}
	collection := r.client.Database(r.db).Collection("items")
	filter := bson.M{"nouser": nouser}
	err := collection.FindOne(ctx, filter).Decode(employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Error Finding a catalogue item")
		}
		return nil, errors.Wrap(err, "repository research")
	}
	return employee, nil

}

func (r *mongoRepository) FindAll() ([]*domain.Employee, error) {

	var items []*domain.Employee

	collection := r.client.Database(r.db).Collection("items")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.TODO()) {

		var item domain.Employee
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil

}
func (r *mongoRepository) Delete(nouser string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"code": nouser}

	collection := r.client.Database(r.db).Collection("items")
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil

}
