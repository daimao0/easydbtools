package model

// Column is the domain model for a column
type Column struct {
	Name       string
	Type       string
	IsPrimary  bool
	IsNullable bool
	IsUnique   bool
	IsAuto     bool
	Default    string
	Comment    string
}
