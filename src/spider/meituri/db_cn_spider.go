package main

import (
	"./gorm"
	"../../../conf"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	runtime.GOMAXPROCS(100)
	//WG.Add(1)
	//891 8245 8225
	//downloadModelColums([]int{8245}) //795,1289,954,3175,467,1558,429, 3239, 2008, 893,919
	//getModelColums()
}
func changdir(id int) {
	err := os.Rename("../meituri/"+strconv.Itoa(id), conf.FSRoot+"../meituri_cn/"+strconv.Itoa(id))
	if err != nil {
		fmt.Println(err)
		return
	}
}
