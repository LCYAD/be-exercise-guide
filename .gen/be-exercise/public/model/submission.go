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

type Submission struct {
	ID           int32 `sql:"primary_key"`
	StudentID    *int32
	AssignmentID *int32
	ExamID       *int32
	SubmittedAt  time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}
