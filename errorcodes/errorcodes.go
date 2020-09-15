package errorcodes

// ErrorCodes -> Defing error code with type string
type ErrorCodes string

const (
	// DivideByZero -> code : 400 & description : Cannot divide any number by zero
	DivideByZero ErrorCodes = "DIVIDE_BY_ZERO"  
)