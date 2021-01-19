package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	con = &MysqlConnect{User: "root", Pwd: "12345678", Host: "10.211.55.3", Port: 3306, DbName: "db_test"}
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test")
	// db, err := GetDb()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	defer con.SafeClose()

	user := &User{}
	// 添加数据
	id := user.Add("test", "test@test.com")

	// 修改数据
	user.Update(id, fmt.Sprintf("test%d", id), fmt.Sprintf("test%d@test.com", id))

	// 查询一行数据
	userInfo := user.Info(id)
	log.Println("user info", userInfo)

	// 查询数据
	var users []*User = user.List()
	usersLen := len(users)
	var delUser []int64 = make([]int64, 2, 2)
	for i := 0; i < usersLen; i++ {
		if i < 2 {
			delUser[i] = users[i].UserID
		} else {
			break
		}
	}
	// var delUser *list.List = list.New()
	// for _, user := range users {
	// 	if 2 > delUser.Len() {
	// 		delUser.PushBack(strconv.FormatInt(user.UserID, 10))
	// 	}
	// }
	// log.Println(users)
	// log.Println(delUser)
	// var ind int = 0
	// var delUserID []string = make([]string, 2, 2)
	// for i := delUser.Front(); i != nil; i = i.Next() {
	// 	delUserID[ind] = i.Value.(string)
	// 	ind++
	// }

	// 删除数据
	user.BatchDel(delUser)

	// 事务删除
	user.DelTx(id)

	// Redis操作
	rc := &RedisClient{RedisConfig: RedisConfig{Host: "10.211.55.3", Port: 6379, Db: 0}}
	GetRdb(rc)

	key := "test:redis:host"
	rc.Set(key, "10.211.55.3", 30*time.Second)

	host, _ := rc.Get(key)
	log.Println("redis key:value", key, host)

	rc.IncrAndExpire("test:redis:increment", 30*time.Second)
}
