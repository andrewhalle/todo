package common

import (
	"log"
)

func CheckDie(e error) {
	if e != nil {
		log.Fatalln("error: ", e)
	}
}
