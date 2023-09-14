package service

import (
	"testing"

	"github.com/michelo851a1203/golang-testify-mock-example/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepoMock struct {
	mock.Mock
}

func (m *EmployeeRepoMock) FindEmployeeGreaterThanAge(age uint) ([]*model.Employee, error) {
	args := m.Called(age)
	return args.Get(0).([]*model.Employee), nil
}

func TestFindCountEmployeeGreaterThanAge(t *testing.T) {
	mockRepo := new(EmployeeRepoMock)
	mockRepo.On("FindEmployeeGreaterThanAge", uint(25)).Return([]*model.Employee{
		{
			ID:   3,
			Name: "Doe",
			Age:  40,
		},
	}, nil)

	employeeService := NewEmployeeService(mockRepo)
	expect := 1
	actual := employeeService.FindCountEmployeeGreaterThanAge(25)

	assert.Equal(t, expect, actual)
}
