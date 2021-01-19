package core

import "log"

type Command interface{}

func Unexpected(c Command) {
	log.Printf("Unexpected command received: %v", c)
}
