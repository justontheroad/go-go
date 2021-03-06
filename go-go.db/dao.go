package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//MysqlConnect mysql连接
type MysqlConnect struct {
	Host   string
	Port   int16
	User   string
	Pwd    string
	DbName string

	db *sql.DB
}

//Open 打开连接
func (c *MysqlConnect) Open() (db *sql.DB, err error) {
	log.Println(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName))
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName))

	if nil == err {
		c.db = db
	} else {
		c.db = nil
	}
	return
}

//SafeClose 安全关闭
func (c *MysqlConnect) SafeClose() {
	if nil != c.checkDb() {
		c.db.Close()
		// defer c.db.Close()
	}
}

func (c *MysqlConnect) checkDb() (err error) {
	if nil == c.db {
		err = errors.New("db 连接异常")
		return
	}
	// 检查数据库是否可用和可访问
	if err = c.db.Ping(); nil != err {
		return
	}
	err = nil
	return
}

//User 用户
type User struct {
	UserID   int64
	UserName string
	Email    string
}

//Add 添加数据
func (user *User) Add(userName string, email string) (id int64) {
	db, err := GetDb()
	if err != nil {
		log.Fatalln(err)
		id = -1
		return
	}

	stmt, err := db.Prepare("INSERT INTO `user` (`user_name`, `email`) VALUES (?, ?)")
	if err != nil {
		log.Fatalln(err)
		id = -1
		return
	}
	defer stmt.Close()

	rs, err := stmt.Exec(userName, email)
	if err != nil {
		log.Fatalln(err)
		id = -1
		return
	}

	id, err = rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
		id = -1
		return
	}

	log.Println(id)
	return
}

//Update 更新
func (user *User) Update(id int64, userName string, email string) (ret int64) {
	db, err := GetDb()
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}

	stmt, err := db.Prepare("UPDATE `user` SET `user_name` = ?, `email` = ? WHERE `user_id` = ?")
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}
	defer stmt.Close()
	rs, err := stmt.Exec(userName, email, id)
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}

	ret, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}
	log.Println(ret)
	return
}

//List 列表
func (user *User) List() (users []*User) {
	db, err := GetDb()
	if err != nil {
		log.Fatalln(err)
		users = nil
		return
	}

	row, err := db.Query("SELECT `user_id`, `user_name`, `email` FROM `user`")
	if err != nil {
		log.Fatal(err.Error())
		users = nil
		return
	}

	for row.Next() {
		var user User
		row.Scan(&user.UserID, &user.UserName, &user.Email)
		users = append(users, &user)
	}
	return
}

//Info 用户详情
func (user *User) Info(id int64) (userInfo *User) {
	db, err := GetDb()
	if err != nil {
		log.Fatalln(err)
		userInfo = nil
		return
	}

	stmt, err := db.Prepare("SELECT `user_id`, `user_name`, `email` FROM `user` WHERE `user_id` = ?")
	if err != nil {
		log.Fatal(err)
		userInfo = nil
		return
	}
	defer stmt.Close()

	userInfo = &User{}
	err = stmt.QueryRow(id).Scan(&userInfo.UserID, &userInfo.UserName, &userInfo.Email)
	if err != nil {
		log.Fatal(err)
	}
	return
}

//BatchDel 批量删除
func (user *User) BatchDel(idArr []int64) (ret int64) {
	db, err := GetDb()
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}

	var placeholder string
	idLen := len(idArr)
	for i := 0; i < idLen; i++ {
		if 0 == i && 1 < idLen {
			placeholder += "?,"
		} else {
			placeholder += "?"
		}
	}
	// log.Println(placeholder)

	stmt, err := db.Prepare(fmt.Sprintf("DELETE FROM `user` WHERE `user_id` IN (%s)", placeholder))
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}
	defer stmt.Close()

	// idSlice := idArr[0:len(idArr)]
	idLen = len(idArr)
	idSlice := make([]interface{}, idLen, idLen)
	for i := 0; i < idLen; i++ {
		idSlice[i] = idArr[i]
	}
	// 使用...运算符以变量...的形式进行参数传递；变量必须是与可变参数类型相同的slice
	rs, err := stmt.Exec(idSlice...)
	// rsd, err := stmtd.Exec(strings.Join(delUserID, ","))
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}

	ret, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}
	log.Println(ret)
	return
}

//DelTx 事务删除
func (user *User) DelTx(id int64) (ret int64) {
	db, err := GetDb()
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}

	defer tx.Rollback() // tx.Commit()执行之后，事务已提交，再调用tx.Rollback()时已无法回滚事务
	stmt, err := tx.Prepare("DELETE FROM `user` WHERE `user_id` = ?")
	if err != nil {
		log.Fatal(err)
		ret = -1
		return
	}
	defer stmt.Close() // danger!

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
		ret = -1
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		ret = -1
	}

	ret = 1
	return
}

var con *MysqlConnect

//GetDb 获取db
func GetDb() (db *sql.DB, err error) {
	if nil != con.checkDb() {
		db, err = con.Open()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(db)
	} else {
		db = con.db
		log.Println(con)
	}
	return
}
