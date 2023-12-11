package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 从wsl访问主机数据库
const connectStrWSL = "wsl_root:@tcp(172.24.192.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

type Image struct {
	Id      int64  `gorm:"column:id"`
	Content string `gorm:"column:content"`
}

func main() {

	db, err := gorm.Open("mysql", connectStrWSL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// db.LogMode(true)

	// Read the entire file into a byte slice
	imgBytes, err := os.ReadFile("./../assets/snow.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Append the base64 encoded output
	base64Encoding = toBase64(imgBytes)

	// Create
	var img Image
	img = Image{Content: base64Encoding}
	db.Create(&img)

	// Read
	var img2 Image
	db.Table("images").Where("id = ?", img.Id).Scan(&img2)

	imgBytes, err = fromBase64(img2.Content)
	if err != nil {
		log.Fatal(err)
	}

	imgDec, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./../assets/snow2.jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100

	err = jpeg.Encode(out, imgDec, &opts)
	if err != nil {
		log.Println(err)
	}
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func fromBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
