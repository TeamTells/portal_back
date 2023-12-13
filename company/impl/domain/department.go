package domain

type Department struct {
	Id               int
	Name             string
	ParentDepartment *parentDepartment
	Supervisor       *supervisor
}

type parentDepartment struct {
	Id   int
	Name string
}

type supervisor struct {
	Id   int
	Name string
}
