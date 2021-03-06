package api

import (
	model "../../../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
)

//func likeModel(c *gin.Context) {
//	userId, err1 := strconv.Atoi(c.PostForm("userId"))
//	modelId, err2 := strconv.Atoi(c.PostForm("modelId"))
//	if nil == err1 && nil == err2 {
//		like := model.Like{
//			Userid:   userId,
//			Modelid:  modelId,
//			Relation: strconv.Itoa(userId) + "_" + strconv.Itoa(modelId),
//		}
//		db.Save(&like)
//		model := model.Model{
//			ID: modelId,
//		}
//		db.Find(&model)
//		model.Hot = model.Hot + 100
//		fmt.Print(model.Name, model.Hot)
//		db.Update(&model)
//		c.JSON(200, gin.H{"toast": "收藏成功"})
//	} else {
//		c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
//	}
//}
func LikeModelList(c *gin.Context) {
	//user_id, err := strconv.Atoi(c.Query(USERID))
	user_id := getUserIdWithToken(c)
	if user_id == -1 {
		return
	} else if user_id > 0 {
		var likes []model.LikeModel
		db.Table("likes").Preload("Model").Where("userid = ?", user_id).Find(&likes)
		c.JSON(200, gin.H{"data": likes})
		return
	} else {
		//println(err.Error())
	}
	c.JSON(400, gin.H{"msg": "未知错误"})
}
func LikeAlbumList(c *gin.Context) {

	//user_id, err := strconv.Atoi(c.Query(USERID))
	user_id := getUserIdWithToken(c)

	var tableNmae = "like_album" + strconv.Itoa(user_id/1000)
	if user_id == -1 {

	} else if user_id > 0 {
		var likes []model.LikeAlbum
		db.Table(tableNmae).Preload("Album").Where("userid = ?", user_id).Find(&likes)
		c.JSON(200, gin.H{"data": likes})
		return
	} else {
		//println(err.Error())
	}
	c.JSON(400, gin.H{"msg": "未知错误"})

}
func Like(c *gin.Context) {
	user_id := getUserIdWithToken(c)
	if user_id == -1 {
		return
	} else if user_id > 0 {
		//user_id, err0 := strconv.Atoi(c.Query("user_id"))
		model_id, _ := strconv.Atoi(c.Query("model_id"))
		album_id, err1 := strconv.Atoi(c.Query("album_id"))
		//dislike,err2:=strconv.Atoi(c.Query("dis"))
		if nil == err1 {
			likeColum(user_id, model_id, album_id, c)
			return
		} else {
			followModel(user_id, model_id, c)
			return
		}
	}
	c.JSON(404, gin.H{"msg": "登录超时"})
}

func likeColum(user_id int, model_id int, album_id int, c *gin.Context) {
	var tableNmae = "like_album" + strconv.Itoa(user_id/1000)
	var like = model.LikeAlbum{
		Userid:   user_id,
		Modelid:  model_id,
		Albumid:  album_id,
		Relation: strconv.Itoa(user_id) + "_" + strconv.Itoa(model_id) + "_" + strconv.Itoa(album_id),
	}
	db.Table(tableNmae).Where("relation = ?", like.Relation).First(&like)
	newrec := db.Table(tableNmae).NewRecord(&like)
	var m = model.Model{}
	m.ID = model_id

	db.First(&m)
	var album = model.Album{
		ID: album_id,
	}
	db.First(&album)

	if newrec {
		db.Table(tableNmae).Save(&like)
		m.Hot += like_album_hot
		//println(model.Hot)
		db.Table("models_cn").Save(&m)
		album.Hot += 10
		db.Save(&album)
		c.JSON(200, gin.H{"toast": "收藏成功", "data": album.Hot})
	} else {
		db.Table(tableNmae).Delete(&like)
		m.Hot -= like_album_hot
		//println(model.Hot)
		db.Table("models_cn").Save(&m)
		album.Hot -= 10
		db.Save(&album)
		c.JSON(200, gin.H{"toast": "取消成功", "data": album.Hot})
	}
}

func followModel(user_id int, model_id int, c *gin.Context) {
	var like = model.LikeModel{
		Userid:   user_id,
		Modelid:  model_id,
		Relation: strconv.Itoa(user_id) + "_" + strconv.Itoa(model_id),
	}
	db.Where("relation = ?", like.Relation).First(&like)
	newrec := db.NewRecord(&like)
	var model = model.Model{}
	model.ID = model_id

	db.Table("models_cn").First(&model)
	if newrec {
		db.Save(&like)
		model.Hot += follow_hot
		//println(model.Hot)
		db.Table("models_cn").Save(&model)
		c.JSON(200, gin.H{"toast": "收藏成功", "data": model.Hot})
	} else {
		db.Delete(&like)
		model.Hot -= follow_hot
		//println(model.Hot)
		db.Table("models_cn").Save(&model)
		c.JSON(200, gin.H{"toast": "取消成功", "data": model.Hot})
	}
}

const follow_hot = 100
const like_album_hot = 10
const view_item_hot = 1
