package repo

import (
	"testing"

	"github.com/michelo851a1203/golang-testify-mock-example/model"
	"github.com/stretchr/testify/assert"
)

func TestFindEmployeeGreaterThanAge(t *testing.T) {
	employeeRepo := NewEmployeeRepoImp()
	employees, err := employeeRepo.FindEmployeeGreaterThanAge(30)
	expectEmployee := []*model.Employee{
		{
			ID:   3,
			Name: "Doe",
			Age:  40,
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expectEmployee, employees)
}
