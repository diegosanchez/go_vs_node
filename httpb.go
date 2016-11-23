package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func connect() (*mgo.Collection, *mgo.Session) {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(err)
	}

	return s.DB("local").C("startup_log"), s
}

func nextEntry(c *mgo.Collection, e *bson.M) error {
	return c.Find(nil).One(e)
}

func disconnect(s *mgo.Session) {
	s.Close()
}

func main() {
	var entry bson.M
	col, s := connect()
	err := nextEntry(col, &entry)

	if err != nil {
		panic(err)
	}

	bs, err := json.Marshal(entry)

	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	json.Indent(&out, bs, "=", "\t")
	out.WriteTo(os.Stdout)

	disconnect(s)
}
