package main

import (
	"fmt"
	"net/http"
	"os"
	//"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"

)
func main() {
	fmt.Println("hello")
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mimi World from path: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":" + PORT, nil)


	//r:= gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()

	//db, err := gorm.Open("sqlite3", "test.db")

	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

}

