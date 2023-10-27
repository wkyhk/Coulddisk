package models

import "gorm.io/gorm"

type ShareBasic struct {
	gorm.Model
	Identity           string
	UserIdentity       string
	RepositoryIdentity string
	ExpiredTime        int
	ClickNum           int
}

func (s ShareBasic) TableName() string {
	return "share_basic"
}
