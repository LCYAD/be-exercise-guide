//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Student struct {
	ID           int32 `sql:"primary_key"`
	FirstName    string
	LastName     string
	Email        string
	Dob          time.Time
	DepartmentID *int32
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}