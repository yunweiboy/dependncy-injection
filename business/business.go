package business

import "dependncy-injection/database"

// BusinessLogic interface
type BusinessLogic interface {
	ProcessData() string
}

// Business struct
type Business struct {
	db database.DatabaseAccess
}

// NewBusiness Constructor
func NewBusiness(db database.DatabaseAccess) *Business {
	return &Business{
		db: db,
	}
}

// ProcessData Implement ProcessData
func (b Business) ProcessData() string {
	return "Business logic processed " + b.db.GetData()
}
