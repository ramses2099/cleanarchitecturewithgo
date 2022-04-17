package model

import "time"

type User struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Age      int64     `json:"age"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at"`
}
