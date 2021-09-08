// Code generated by entc, DO NOT EDIT.

package logsearch

import (
	"time"
)

const (
	// Label holds the string label denoting the logsearch type in the database.
	Label = "log_search"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKeyword holds the string denoting the keyword field in the database.
	FieldKeyword = "keyword"
	// FieldPage holds the string denoting the page field in the database.
	FieldPage = "page"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the logsearch in the database.
	Table = "log_searches"
)

// Columns holds all SQL columns for logsearch fields.
var Columns = []string{
	FieldID,
	FieldKeyword,
	FieldPage,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)