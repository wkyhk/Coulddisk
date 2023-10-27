package models

import "gorm.io/gorm"

type UserRepository struct {
	gorm.Model
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	Size               int64
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
