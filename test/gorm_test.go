package test

import (
	models "CLOUDDISK/core/Models"
	"encoding/json"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	db, err := gorm.Open(mysql.Open("whk:123456@tcp(192.168.137.100:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	db.Find(&data)
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))

}
