package service

import "github.com/michelo851a1203/golang-testify-mock-example/repo"

type EmployeeService interface {
}

type EmployeeServiceImp struct {
	EmployeeRepo repo.EmployeeRepo
}

func NewEmployeeService(repo repo.EmployeeRepo) *EmployeeServiceImp {
	return &EmployeeServiceImp{
		EmployeeRepo: repo,
	}
}

func (e *EmployeeServiceImp) FindCountEmployeeGreaterThanAge(age uint) int {
	result, _ := e.EmployeeRepo.FindEmployeeGreaterThanAge(age)
	return len(result)
}
