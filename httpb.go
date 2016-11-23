package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func main() {
	var entry bson.M

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(err)
	}

	c := session.DB("local").C("startup_log")

	err = c.Find(nil).One(&entry)

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

	session.Close()
}
