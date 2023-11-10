package service

import "dependncy-injection/business"

// Service interface
type Service interface {
	HandleRequest() string
}

// ServiceImpl struct
type ServiceImpl struct {
	logic business.BusinessLogic
}

// NewService Constructor
func NewService(logic business.BusinessLogic) *ServiceImpl {
	return &ServiceImpl{
		logic: logic,
	}
}

// HandleRequest Implement HandleRequest()
func (s ServiceImpl) HandleRequest() string {
	return "Handled request: " + s.logic.ProcessData()
}
