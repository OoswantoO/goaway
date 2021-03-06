package api

import (
	"../../../conf"
	model "../../../model/meituri"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"strings"
)

//func GetTandomHotTab(c *gin.Context){
//
//}

//伪随机 15 个tag
func GetTandomHotTab(c *gin.Context) {
	gender := c.Query("gender")
	pageNo, err := strconv.Atoi(c.Query("pageNo"))

	if err == nil {
		var companies [] model.Company
		var groups [] model.Group
		var models []model.Model
		var tags []model.Tag
		var tabs [] model.Tab

		if gender == "woman" {
			c.JSON(200, gin.H{"toast":"参数暂不支持" })
		}else{
			db.Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&companies)
			for _, c := range companies {
				c.Type = conf.Company
				tabs = append(tabs, c.Tab)
			}
			db.Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&groups)
			for _, g := range groups {
				g.Type = conf.Group
				tabs = append(tabs, g.Tab)
			}
			db.Table("models").Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&models)
			for _, m := range models {
				m.Type = conf.Model
				tabs = append(tabs, m.Tab)
			}
			db.Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&tags)
			for _, t := range tags {
				t.Type = conf.Tag
				tabs = append(tabs, t.Tab)
			}
			c.JSON(200, gin.H{"data": tabs})
		}
	}else{
		c.JSON(200, gin.H{"toast": err.Error()})
	}
}

func FollowTabs(c *gin.Context) {
	var user_id = getUserIdWithToken(c)
	if user_id > 0 {
		//var tabstr=c.PostForm("tabs")
		//readStringFromBody(c.Request.Body)
		tabstr, err0 := ioutil.ReadAll(c.Request.Body)
		if err0 == nil {
			var tabs [] model.Tab
			err := json.NewDecoder(strings.NewReader(string(tabstr))).Decode(&tabs)
			if err == nil {
				for _, t := range tabs {
					var follow = model.FollowTab{
						Userid:   user_id,
						Resid:    t.ID,
						Type:     t.Type,
						Alias:    t.Alias,
						Relation: strconv.Itoa(user_id) + "_" + strconv.Itoa(t.Type) + "_" + strconv.Itoa(t.ID),
					}
					//var new = db.NewRecord(&follow)
					//if new {
						db.Create(&follow)
					//} else {
					//	db.Model(&follow).Where("id = ?", follow.ID).Update("alias", follow.Alias)
					//	db.Model(&follow).UpdateColumn("alias", follow.Alias)
					//}
				}
				c.JSON(200, gin.H{"toast": "关注成功"})
			}
		}
	} else {

	}
}

//func readStringFromBody(closer io.ReadCloser) {
//	buf := make([]byte, 1024)
//	for n>0{
//		n, _ :=closer.Read(buf)
//	}
//	tabstr := string(buf[0:n])
//}
func FollowedTabs(c *gin.Context) {
	var userId = getUserIdWithToken(c)
	if userId > 0 {
		var followed []model.Tab
		var tabs []model.FollowTab
		db.Where("userid = ", userId).Find(&tabs)
		for _, t := range tabs {
			if t.Type == 0 {
				com := model.Company{}
				db.Where("id = ", t.Resid).First(&com)
				followed = append(followed, com.Tab)
			} else if t.Type == 1 {
				gro := model.Group{}
				db.Where("id = ", t.Resid).First(&gro)
				followed = append(followed, gro.Tab)
			} else if t.Type == 2 {
				mo := model.Model{}
				db.Where("id = ", t.Resid).First(&mo)
				followed = append(followed, mo.Tab)
			} else if t.Type == 3 {
				tag := model.Tag{}
				db.Where("id = ", t.Resid).First(&tag)
				followed = append(followed, tag.Tab)
			}
		}
		c.JSON(200, gin.H{"data": followed})
	} else {
		c.JSON(200, gin.H{"toast": "token 失效"})
	}

}
