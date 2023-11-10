package database

// DatabaseAccess interface
type DatabaseAccess interface {
	GetData() string
}

// Database struct
type Database struct{}

// NewDatabase Constructor
func NewDatabase() *Database {
	return &Database{}
}

// GetData Implement GetData
func (d Database) GetData() string {
	return "Data from database"
}
