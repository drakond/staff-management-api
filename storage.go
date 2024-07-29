package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	ID       int    `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Position string `json:"position" bson:"position"`
}

type Department struct {
	ID        int    `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	Employees []int  `json:"employees" bson:"employees"`
}

type MongoStorage struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoStorage(uri string, dbName string) (*MongoStorage, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)
	return &MongoStorage{
		client:   client,
		database: database,
	}, nil
}

func (s *MongoStorage) Insert(e *Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("employees")
	_, err := collection.InsertOne(ctx, e)
	return err
}

func (s *MongoStorage) Get(id int) (Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var employee Employee
	collection := s.database.Collection("employees")
	filter := bson.M{"id": id}
	err := collection.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (s *MongoStorage) GetAll() ([]Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("employees")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var employees []Employee
	for cursor.Next(ctx) {
		var employee Employee
		err := cursor.Decode(&employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (s *MongoStorage) Update(id int, e Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("employees")
	filter := bson.M{"id": id}
	update := bson.M{"$set": e}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *MongoStorage) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("employees")
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (s *MongoStorage) InsertDepartment(d *Department) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("departments")
	_, err := collection.InsertOne(ctx, d)
	return err
}

func (s *MongoStorage) GetDepartment(id int) (Department, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var department Department
	collection := s.database.Collection("departments")
	filter := bson.M{"id": id}
	err := collection.FindOne(ctx, filter).Decode(&department)
	if err != nil {
		return Department{}, err
	}
	return department, nil
}

func (s *MongoStorage) DeleteDepartment(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("departments")
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (s *MongoStorage) AddEmployeeToDepartment(departmentID, employeeID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := s.database.Collection("departments")
	filter := bson.M{"id": departmentID}
	update := bson.M{"$push": bson.M{"employees": employeeID}}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
