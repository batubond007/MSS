package main

type Options struct {
	Addr     string
	MongoUri string
}

func NewOptions() *Options {
	return &Options{
		Addr:     "localhost:8080",
		MongoUri: "mongodb://localhost:27017",
	}
}
