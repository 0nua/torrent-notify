package error

import "log"

func Check(err error) {
	if err == nil {
		return
	}

	log.Panic(err)
}
