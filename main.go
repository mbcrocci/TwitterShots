package main

import (
	"fmt"
	"os"

	"labix.org/v2/mgo"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

type SearchWord struct {
	Word string `form:"word"`
}

func DB() martini.Handler {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(os.Getenv("MONGO_DB")))
		defer s.Close()
		c.Next()
	}
}

func GetAll(db *mgo.Database) []SearchWord {
	var wordlist []SearchWord
	db.C("words").Find(nil).All(&wordlist)
	return wordlist
}

func main() {

	if connectToMongo() {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())

	m.Get("/", func(r render.Render, db *mgo.Database) {
		r.HTML(200, "index", GetAll(db))
	})

	m.Post("/", binding.Form(SearchWord{}), func(word SearchWord, r render.Render, db *mgo.Database) {
		db.C("words").Insert(word)
		r.HTML(200, "index", GetAll(db))
	})

	m.NotFound(func() string {
		return "Someone stole this page :o"
	})

	m.Run()
}
