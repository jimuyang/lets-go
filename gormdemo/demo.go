package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) String() string {
	return fmt.Sprintf("id: %d, code: %s, price: %d", p.ID, p.Code, p.Price)
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo4go?charset=utf8&parseTime=true")
	if err != nil {
		panic("cannot connect to mysql")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})
	// 插入
	db.Create(&Product{Code: "111", Price: 10,})

	// 查询
	var product Product
	db.First(&product)
	fmt.Println(product.String())

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println(product.String())
	db.First(&product, "code = ?", "111")
	fmt.Println(product.String())

	// 删除
	//db.Delete(&product)
}
