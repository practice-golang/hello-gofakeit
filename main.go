package main // import "hello-gofakeit"

import (
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Foo struct {
	Str           string
	Int           int
	Pointer       *int
	Name          string         `fake:"{firstname}"`  // Any available function all lowercase
	Sentence      string         `fake:"{sentence:3}"` // Can call with parameters
	RandStr       string         `fake:"{randomstring:[hello,world]}"`
	Number        string         `fake:"{number:1,10}"`       // Comma separated for multiple values
	Regex         string         `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
	Map           map[string]int `fakesize:"2"`
	Array         []string       `fakesize:"2"`
	Bar           Bar
	Skip          *string   `fake:"skip"` // Set to "skip" to not generate data for
	Created       time.Time // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

type Bar struct {
	Name   string
	Number int
	Float  float32
}

func main() {
	log.Println(gofakeit.Name())
	log.Println(gofakeit.Email())
	log.Println(gofakeit.Phone())
	log.Println(gofakeit.BS())
	log.Println(gofakeit.BeerName())
	log.Println(gofakeit.Color())
	log.Println(gofakeit.Company())
	log.Println(gofakeit.CreditCardNumber(&gofakeit.CreditCardOptions{}))
	log.Println(gofakeit.HackerPhrase())
	log.Println(gofakeit.JobTitle())
	log.Println(gofakeit.CurrencyShort())

	log.Println("----------------------------------------")

	log.Println(gofakeit.PetName())
	log.Println(gofakeit.Animal())
	log.Println(gofakeit.AnimalType())
	log.Println(gofakeit.FarmAnimal())
	log.Println(gofakeit.Cat())
	log.Println(gofakeit.Dog())

	log.Println("----------------------------------------")

	var f Foo
	gofakeit.Struct(&f)

	log.Println(f.Str)
	log.Println(f.Int)
	log.Println(*f.Pointer)
	log.Println(f.Name)
	log.Println(f.Sentence)
	log.Println(f.RandStr)
	log.Println(f.Number)
	log.Println(f.Regex)
	log.Println(f.Map)
	log.Println(f.Array)
	log.Printf("%+v", f.Bar)
	log.Println(f.Skip)
	log.Println(f.Created.String())
}
