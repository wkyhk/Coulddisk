package models

import "gorm.io/gorm"

type RepositoryPool struct {
	gorm.Model
	Identity string `json:"identity"`
	Hash     string `json:"hash"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	Size     int64  `json:"size"`
	Path     string `json:"path"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
