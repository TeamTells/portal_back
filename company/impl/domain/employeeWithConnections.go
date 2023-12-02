package domain

import "time"

type EmployeeWithConnections struct {
	Id              int
	FirstName       string
	SecondName      string
	Surname         string
	DateOfBirth     time.Time
	Email           string
	Icon            string
	TelephoneNumber string
	Company         struct {
		Id   int
		Name string
	}
	Departments []struct {
		Id   int
		Name string
	}
	Roles []struct {
		Id   int
		Name string
	}
}
