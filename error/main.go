package main

import (
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func loadConfig() error {
	_, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	return nil
}
func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
}
