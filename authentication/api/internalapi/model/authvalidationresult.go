package model

type AuthValidationResult int

const (
	SUCCESS  = 0
	NOTFOUND = 1
	EXPIRED  = 2
)
