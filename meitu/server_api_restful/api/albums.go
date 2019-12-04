package api

import (
	"../../../conf"
	model "../../model/meituri"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

func GetAlbumDetail(c *gin.Context) {
	useridstr := getUserIdStrWithToken(c)
	userid, err := strconv.Atoi(useridstr)
	if err == nil && userid > 0 {
		modelIdStr := c.Query("model_id")
		albumIdStr := c.Query("album_id")
		albumId, _ := strconv.Atoi(albumIdStr)
		album := model.Album{}
		db.Where("id = ?", albumIdStr).First(&album)
		if album.ID > 0 {
			//downloadFile(durl,path,filename)
			var now = time.Now().Format("2006-01-02")
			var record = model.VisitHistroy{
				Albumid:  albumId,
				Userid:   userid,
				Date:     now,
				Relation: albumIdStr + "_" + useridstr + "_" + now,
			}
			var tableName = conf.VisitHistroy + strconv.Itoa(userid/1000)
			if !db.HasTable(tableName) {
				db.Table(tableName).Create(model.VisitHistroy{})
			}
			db.Table(tableName).Create(&record)

			appendImagesForAlbum(modelIdStr, albumIdStr, &album)

			c.JSON(200, gin.H{"data": &album})
		} else {
			c.JSON(404, gin.H{"message": "album not exist"})
		}
	}

}

func appendImagesForAlbum(modelIdStr string, albumIdStr string, album *model.Album) {
	p := "/" + modelIdStr + "/" + albumIdStr + "/"
	path := conf.FSMuri + p
	rd, err := ioutil.ReadDir(path)
	if err == nil {
		var paths []string
		for _, fi := range rd {
			if fi.IsDir() {
				fmt.Printf("[%s]\n", fi.Name())
			} else {
				fmt.Println(fi.Name())
				p := conf.FILE_SERVER + conf.Muri + p + fi.Name()
				paths = append(paths, p)
				fmt.Println(len(paths), cap(paths), paths, p)
			}
		}
		album.Images = paths
	} else {
		println(err.Error())
	}

}
func GetAlbumsList(c *gin.Context) {
	tag, err0 := strconv.Atoi(c.Query("tag"))

	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	var albums []model.Album

	if err0 == nil {
		db.Where("tags LIKE ?", tag).Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&albums) //.Order("created_at desc")
		//c.String(200,)

		c.JSON(200, gin.H{"data": &albums})
	} else {
		if nil == err1 && nil == err2 {
			//if len(search) == 0 {
			//} else {
			//}
			db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Order("id desc").Find(&albums) //.Order("created_at desc")
			//c.String(200,)
			c.JSON(200, gin.H{"data": &albums})
		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}
}
