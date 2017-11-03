package tools

import "log"

//ValidateFatal check error
//
func ValidateFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//ValidatePanic check error
//
func ValidatePanic(err error) {
	if err != nil {
		panic(err)
	}
}