package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID       int32  `bun:"id,pk,autoincrement"`
	Email    string `bun:"email,notnull,unique,type:varchar(255)"`
	Password string `bun:"password,notnull,type:varchar(255)"`
	Username string `bun:"username,notnull,type:varchar(255)"`
}
