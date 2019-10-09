package orm

import (
	model "../../../model/meituba"
	"database/sql"
	"fmt"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "Qunsi003"
	NETWORK  = "tcp"
	SERVER   = "rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com"
	PORT     = 3306
	DATABASE = "meituba"
)

var db *sql.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)                  //设置最大连接数
	db.SetMaxIdleConns(16)                   //设置闲置连接数
}
func UpdateColumTime(c *model.Colums, tableName string) {
	db.Exec("UPDATE "+tableName+" set time=? where id=?", c.Vtimes, c.ID)
}

func SaveColum(c *model.Colums, tableName string) {
	db.Exec("insert INTO "+tableName+"(id,title,vtimes,count,tags) values(?,?,?,?,?)",
		c.ID, c.Title, c.Vtimes, c.Count, c.Tags)
	//if err != nil {
	//	fmt.Println("Insert failed,err:%v", err)
	//}
	//lastInsertID,err := result.LastInsertId()  //插入数据的主键id
	//if err != nil {
	//	fmt.Printf("Get lastInsertID failed,err:%v",err)
	//	return
	//}
	//fmt.Println("LastInsertID:",lastInsertID)
	//rowsaffected,err := result.RowsAffected()  //影响行数
	//if err != nil {
	//	fmt.Printf("Get RowsAffected failed,err:%v",err)
	//	return
	//}
	//fmt.Println("RowsAffected:",rowsaffected)
}
func SaveTag(c *model.Tags) {
	db.Exec("insert INTO tags"+"(cname,ename) values(?,?)",
		c.Cname, c.Ename)
	//if err != nil {
	//	fmt.Println("Insert failed,err:%v", err)
	//}
}
