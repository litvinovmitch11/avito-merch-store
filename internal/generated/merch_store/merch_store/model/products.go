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

type Products struct {
	ID        string `sql:"primary_key"`
	Title     string
	Price     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
