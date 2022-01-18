package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cheggaaa/pb/v3"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	numPersons := flag.Int("persons", 1000, "number of persons to generate")
	flag.Parse()

	bar := pb.StartNew(*numPersons)

	f, err := os.Create("data/dump.sql")
	defer f.Close()

	CheckError("failed to create file", err)

	Write(f, "INSERT INTO profiles (id, email, password, first_name, last_name, city, age, interests) VALUES \n")

	for i := 1; i <= *numPersons; i++ {
		delim := ","

		if *numPersons == i {
			delim = ";"
		}

		uniqueEmail := fmt.Sprintf(
			"%s_%.2x@%s",
			gofakeit.Username(),
			md5.Sum([]byte(gofakeit.Word())),
			gofakeit.DomainName(),
		)

		Write(f, fmt.Sprintf("(NULL, %s, %s, %s, %s, %s, %d, %s)%s",
			strconv.Quote(strings.ToLower(uniqueEmail)),
			strconv.Quote("$2a$08$GcINgmM.ynSC9iSFGnFgReNJe5BXjgQIUyRhKMqip4F32VYclhowW"), // test
			strconv.Quote(gofakeit.FirstName()),
			strconv.Quote(gofakeit.LastName()),
			strconv.Quote(gofakeit.City()),
			rand.Intn(gofakeit.Number(10, 100)),
			strconv.Quote(gofakeit.Hobby()),
			delim,
		))

		bar.Increment()
	}

	bar.Finish()
}

func CheckError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}

func Write(f *os.File, s string) {
	_, err := f.WriteString(s)
	CheckError("cannot write to file", err)
}
