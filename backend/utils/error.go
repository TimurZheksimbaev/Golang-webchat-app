package utils

import (
	"fmt"
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Log(message string) {
	log.Println(message)
}

func LogExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DatabaseError(message string, err error) error {
	return fmt.Errorf("[DATABASE] %s: %s", message, err)
}

func ConfigError(message string, err error) error {
	return fmt.Errorf("[CONFIG] %s: %s", message, err)
}

func AuthError(message string, err error) error {
	return fmt.Errorf("[AUTHENTICATION] %s: %s", message, err)
}

func ServiceError(message string, err error) error {
	return fmt.Errorf("[SERVICE] %s: %s", message, err)
}

func HandlerError(message string, err error) error {
	return fmt.Errorf("[HANDLER] %s: %s", message, err)
}

func WebsocketError(message string, err error) error {
	return fmt.Errorf("[WEBSOCKET] %s: %s", message, err)
}


