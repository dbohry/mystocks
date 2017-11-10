package tools

//ValidateFatal check error
//
func ValidateFatal(err error) {
	if err != nil {
		panic(err)
	}
}

//ValidateExecution validate if an execution was successful
//
func ValidateExecution(err error) bool {
	shouldContinue := true
	if err != nil {
		shouldContinue = false
	}
	return shouldContinue
}

//ValidatePanic check error
//
func ValidatePanic(err error) {
	if err != nil {
		panic(err)
	}
}