package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int64
	Name string
	Age  *int
}

func (p *User) TableName() string {
	return "user"
}

func UseGorm() {

	dsn := "root@tcp(127.0.0.1:3306)/play?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	age := 19
	//db.Create(&User{
	//	Name: "xiaoming",
	//	Age:  &age,
	//})

	fmt.Println(db)

	db.Updates(&User{
		ID:  1,
		Age: &age,
	})

	//db.Update()
}
