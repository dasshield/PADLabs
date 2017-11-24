package common

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Panic("err", + err)
	}
}
