// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// XODB is the common interface for database operations that can be used with
// types from public.
//
// This should work with database/sql.DB and database/sql.Tx.
type XODB interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

type QueryOpt interface {
}
