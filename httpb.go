package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// DB
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
	col, s := connect()

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		var entry bson.M
		w.Header().Set("Content-Type", "application/json")
		err := nextEntry(col, &entry)

		if err != nil {
			panic(err)
		}

		bs, err := json.Marshal(entry)
		w.Write(bs)
	})

	http.ListenAndServe(":9000", nil)
	disconnect(s)
}
