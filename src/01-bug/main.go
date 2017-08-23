package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io"
	"text/template"
)

func main() {
	usingGORM()
}

const status = 10

func templater() {

	b := Bug(400)
	b.startServer()
}

type Bug int

func (b *Bug) startServer() {
	i := int(*b)
	fmt.Println(i)
}

// do some http stuff

func learningTemplating() {

	type Choord map[string]int
	type Email []*map[string]string
	type Human struct {
		Name string
		Age  uint8
	}

	render("file.html")

}

func render(view string, data ...interface{}) (func(writer io.Writer), error) {
	const base_dir = "templates/"
	tmp, err := template.ParseGlob(base_dir + view)
	if err == nil {
		tmp.ParseFiles()
		return func(w io.Writer) {
			tmp.Execute(w, data)
		}, nil
	} else {
		fmt.Errorf(err.Error())
		return nil, err
	}
}

func usingGORM() {
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)



	// Delete - delete product
	//db.Delete(&product)
}
