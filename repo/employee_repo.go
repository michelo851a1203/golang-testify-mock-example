package repo

import (
	"sort"

	"github.com/michelo851a1203/golang-testify-mock-example/model"
)

type EmployeeRepo interface {
	FindEmployeeGreaterThanAge(age uint) ([]*model.Employee, error)
}

type EmployeeRepoImp struct {
}

func NewEmployeeRepoImp() *EmployeeRepoImp {
	return &EmployeeRepoImp{}
}

func (e *EmployeeRepoImp) FindEmployeeGreaterThanAge(age uint) ([]*model.Employee, error) {
	fakeEmployee := []*model.Employee{
		{
			ID:   1,
			Name: "John",
			Age:  30,
		},
		{
			ID:   2,
			Name: "Jane",
			Age:  25,
		},
		{
			ID:   3,
			Name: "Doe",
			Age:  40,
		},
	}

	result := []*model.Employee{}
	for _, v := range fakeEmployee {
		if v.Age > age {
			result = append(result, v)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Age < result[j].Age
	})

	return result, nil

}
