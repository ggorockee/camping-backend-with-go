package util

import "log"

func HandleFunc(err error) {
	if err != nil {
		log.Printf("Error occurred: %v", err)
	}
}
