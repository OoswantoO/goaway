package api

import (
	model "../../../model/meituri"
	"github.com/gin-gonic/gin"
)

func GetHotTag(c *gin.Context) {
	var tags = []model.Tag{}
	db.Select("id,name,des,hot").Order("hot desc").Limit(10).Find(&tags) //.Order("created_at desc")
	//c.String(200,)
	c.JSON(200, gin.H{"data": tags})
}
func GetAllTag(c *gin.Context) {
	var tags = []model.Tag{}
	db.Select("id,name,des,hot").Find(&tags) //.Order("created_at desc")
	//c.String(200,)
	c.JSON(200, gin.H{"data": tags})
}
