// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Mutation struct {
}

type MutationProject struct {
	Name       string     `json:"name"`
	UserID     string     `json:"userId"`
	TotalTime  int        `json:"totalTime"`
	GoalTime   int        `json:"goalTime"`
	Deadline   time.Time  `json:"deadline"`
	IsEngaging bool       `json:"isEngaging"`
	StartTime  *time.Time `json:"startTime,omitempty"`
}

type MutationUser struct {
	Name string `json:"name"`
}

type Project struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	User       *User      `json:"user"`
	TotalTime  int        `json:"totalTime"`
	GoalTime   int        `json:"goalTime"`
	Deadline   time.Time  `json:"deadline"`
	IsEngaging bool       `json:"isEngaging"`
	StartTime  *time.Time `json:"startTime,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
}

type Query struct {
}

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Projects  []*Project `json:"projects,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}