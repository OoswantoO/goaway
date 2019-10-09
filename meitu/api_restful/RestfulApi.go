package api_restful

import (
	model "../model/meituri"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func checkTokenEnable(token string) bool {
	return strings.EqualFold(token, "token")
}

func GetModelList(c *gin.Context) {
	//search := c.PostForm("search")
	if tokenEnable(c) {
		pageNo, err1 := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, err2 := strconv.Atoi(c.PostForm("pageSize"))
		if nil == err1 && nil == err2 {
			var models = []model.Models{}
			//if len(search) == 0 {
			//} else {
			//}
			db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models) //.Order("created_at desc")
			//c.String(200,)
			c.JSON(200, gin.H{"data": models})
		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}
}

func tokenEnable(c *gin.Context) bool {
	token := c.GetHeader("token")
	if !checkTokenEnable(token) {
		c.JSON(401, gin.H{"status": -1, "msg": "token已失效"})
		return false
	} else {
		return true
	}
}

func GetHotTag(c *gin.Context) {
	var tags = []model.Tags{}
	db.Select("id,shortname,des,hot").Order("hot desc").Limit(10).Find(&tags) //.Order("created_at desc")
	//c.String(200,)
	c.JSON(200, gin.H{"data": tags})
}
func GetAllTag(c *gin.Context) {
	var tags = []model.Tags{}
	db.Select("id,shortname,des,hot").Find(&tags) //.Order("created_at desc")
	//c.String(200,)
	c.JSON(200, gin.H{"data": tags})
}

func GetColumsList(c *gin.Context) {
	tag, err0 := strconv.Atoi(c.Query("tag"))

	pageNo, err1 := strconv.Atoi(c.PostForm("pageNo"))
	pageSize, err2 := strconv.Atoi(c.PostForm("pageSize"))
	var colums = []model.Colums{}

	if err0 == nil {
		db.Where("tags LIKE ?", tag).Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&colums) //.Order("created_at desc")
		//c.String(200,)
		c.JSON(200, gin.H{"data": colums})
	} else {
		if nil == err1 && nil == err2 {
			//if len(search) == 0 {
			//} else {
			//}
			db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Order("id desc").Find(&colums) //.Order("created_at desc")
			//c.String(200,)
			c.JSON(200, gin.H{"data": colums})
		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}

}

func resetPass(c *gin.Context) {

}

//todo
func tokenLogin(c *gin.Context) {

}
func Login(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("pwd")

	//pwd:=c.PostForm("pwd")
	user := model.Users{}
	db.First(&user, "account = ?", account)
	fmt.Println(user.Info())
	if user.ID > 0 {
		if strings.EqualFold(user.Pwd, pwd) {
			user.Token = "token"
			c.JSON(200, gin.H{"msg": "密码正确",
				"data": &user})
		} else {
			c.JSON(200, gin.H{"status": -1, "msg": "确认密码不符"})
		}
	} else {
		c.JSON(200, gin.H{"msg": "用户不存在"})
	}
}
func RegistAccount(c *gin.Context) {
	//c.Header("tel","")
	tel := c.PostForm("tel")
	pwd := c.PostForm("pwd")
	repwd := c.PostForm("repwd")
	//fmt.Println(pwd + repwd)
	if strings.EqualFold(pwd, repwd) {
		email := c.PostForm("email")
		if len(tel) > 0 {
			if len(pwd) > 6 {
				user := model.Users{
					Account: tel,
					Tel:     tel,
					Pwd:     pwd,
				}
				createSuccess := db.NewRecord(&user)
				if createSuccess {
					if err := db.Create(&user).Error; err != nil {
						//return -3
						println(err.Error())
					}
				} else {
					c.JSON(200, gin.H{"status": -1, "msg": "创建失败"})
				}

			} else {
				c.JSON(200, gin.H{"status": -1, "msg": "密码过短"})
			}
		} else if len(email) > 4 {
			if len(pwd) > 6 {
				user := model.Users{
					Account: email,
					Email:   email,
					Pwd:     pwd,
				}
				createSuccess := db.NewRecord(&user)
				if createSuccess {
					if err := db.Create(&user).Error; err != nil {
						//return -3
						println(err.Error())
					}
				} else {
					c.JSON(200, gin.H{"status": -1, "msg": "创建失败"})
				}
			} else {
				c.JSON(200, gin.H{"status": -1, "msg": "密码过短"})
			}
		}
	} else {
		c.JSON(200, gin.H{"status": -1, "msg": "确认密码不符"})
	}
	c.JSON(200, gin.H{"status": 1, "msg": "创建成功"})
}

//todo
func InsertUser(user *model.Users) error {

	return nil
}
func ChangeUserName(userId uint64, userNaem string) error {

	return nil
}
func DeleteUser(userId uint64, statusType model.UserStatusType) error {

	return nil
}
func GetUser(userId uint64) (model.Users, error) {
	return model.Users{}, nil
}
func GetUsers() ([]model.Users, error) {
	return []model.Users{}, nil
}
