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

type Exam struct {
	ID         int32 `sql:"primary_key"`
	Name       string
	Type       int16
	StartedAt  *time.Time
	FinishedAt *time.Time
	CourseID   *int32
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}
