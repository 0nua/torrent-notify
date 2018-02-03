package error

import (
	"log"
)

func Catch(err error) {
	if err == nil {
		return
	}

	log.Panic(err)
}
