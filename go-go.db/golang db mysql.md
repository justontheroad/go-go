### golang db
```
To access databases in Go, you use a sql.DB. You use this type to create statements and transactions, execute queries, and fetch results.

The first thing you should know is that a sql.DB isn’t a database connection. It also doesn’t map to any particular database software’s notion of a “database” or “schema.” It’s an abstraction of the interface and existence of a database, which might be as varied as a local file, accessed through a network connection, or in-memory and in-process.

sql.DB performs some important tasks for you behind the scenes:

It opens and closes connections to the actual underlying database, via the driver.
It manages a pool of connections as needed, which may be a variety of things as mentioned.
The sql.DB abstraction is designed to keep you from worrying about how to manage concurrent access to the underlying datastore. A connection is marked in-use when you use it to perform a task, and then returned to the available pool when it’s not in use anymore. One consequence of this is that if you fail to release connections back to the pool, you can cause sql.DB to open a lot of connections, potentially running out of resources (too many connections, too many open file handles, lack of available network ports, etc). We’ll discuss more about this later.

After creating a sql.DB, you can use it to query the database that it represents, as well as creating statements and transactions.
```
```
在Go中访问数据库，请使用sql.DB。你可以使用这种类型来创建语句和事务，执行查询以及获取结果。

第一件事你应该知道，sql.DB不是数据库连接。它也没有映射到任何特定数据库软件的“数据库”或“模式”的概念。它是数据库的接口和存在的一种抽象，可以是通过网络连接访问的本地文件，也可以是内存中和进程中文件。
sql.DB 在幕后为你执行一些重要的任务：
    它通过驱动程序打开和关闭到实际底层数据库的连接。
    它根据需要管理连接池，如前所述，连接池可以是各种各样的东西。

db抽象的设计目的是让你不必担心如何管理对底层数据存储的并发访问。当你使用一个连接来执行一个任务时，它会被标记为“使用中”，当它不再使用时，它会返回到可用池中。这样做的一个后果是，如果你不能释放连接回池，你可能会导致sql.DB打开大量的连接，潜在地耗尽资源(太多的连接，太多打开的文件句柄，缺少可用的网络端口，等等)。稍后我们将对此进行更多讨论。

在创建sql.DB之后，你可以使用它来查询它所表示的数据库，以及创建语句和事务。
```

1. 导入 mysql 数据库驱动
    ```
    import (
    	"database/sql"
    	_ "github.com/go-sql-driver/mysql"
    )
    ```
2. 访问数据库
    ```
    db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    ```
3. 查询
    ```
    rows, err := db.Query("SELECT `user_id`, `user_name`, `email` FROM `user` WHERE `id` = ?", 1)
    if err != nil {
    	log.Fatal(err)
    	return
    }
    defer rows.Close()
    var id int
        var name, mail string
    for rows.Next() {
    	err := rows.Scan(&id, &name, &mail)
    	if err != nil {
    		log.Fatal(err)
    	}
    	log.Println(id, name, mail)
    }
    ```
    - 预查询
        ```
        stmt, err := db.Prepare("SELECT `user_id`, `user_name`, `email` FROM `user` WHERE `id` = ?")
        if err != nil {
        	log.Fatal(err)
        	return
        }
        defer stmt.Close()
        rows, err := stmt.Query(1)
        if err != nil {
        	log.Fatal(err)
        	return
        }
        defer rows.Close()
        var id int
        var name, mail string
        for rows.Next() {
        	err := rows.Scan(&id, &name, &mail)
        	if err != nil {
        		log.Fatal(err)
        	}
        	log.Println(id, name, mail)
        }
        if err = rows.Err(); err != nil {
        	log.Fatal(err)
        }
        ```
    - 单行查询
        ```
        var id int
        var name, mail string
        err = db.QueryRow("SELECT `user_id`, `user_name`, `email` FROM `user` WHERE `id` = ?", 1).Scan(&id, &name, &mail)
        if err != nil {
        	log.Fatal(err)
        	return
        }
        return
        ```
        ```
        stmt, err := db.Prepare("SELECT `user_id`, `user_name`, `email` FROM `user` WHERE `id` = ?")
        if err != nil {
        	log.Fatal(err)
        	return
        }
        defer stmt.Close()
        var id int
        var name, mail string
        err = stmt.QueryRow(1).Scan(&idm &name, &mail)
        if err != nil {
        	log.Fatal(err)
        }
        return
        ```
4. 添加数据
    ```
    _, err = db.Exec("INSERT INTO `user` (`user_name`, `email`) VALUES (?, ?)", "test", "test@test.com")
    if err != nil {
    	log.Fatal(err)
    	return
    }
    id, err = rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
		return
	}
	rowCnt, err := res.RowsAffected()
    if err != nil {
    	log.Fatal(err)
    }
    return
    ```
5. 修改数据
    ```
    _, err = db.Exec(UPDATE `user` SET `user_name` = ?, `email` = ? WHERE `user_id` = ?, "test1", "test1@test.com", 1)
    if err != nil {
    	log.Fatal(err)
    	return
    }
	rowCnt, err := res.RowsAffected()
    if err != nil {
    	log.Fatal(err)
    }
    return
    ```
6. 删除数据
    ```
    _, err = db.Exec(DELETE FROM `user` WHERE `user_id` = ?, 1)
    if err != nil {
    	log.Fatal(err)
    	return
    }
	rowCnt, err := res.RowsAffected()
    if err != nil {
    	log.Fatal(err)
    }
    return
    ```
7. 使用事务
    ```
    tx, err := db.Begin()
    if err != nil {
    	log.Fatal(err)
    	return
    }
    
    _, err = db.Exec(DELETE FROM `user` WHERE `user_id` = ?, 1)
    if err != nil {
    	log.Fatal(err)
    	return
    }
	rowCnt, err := res.RowsAffected()
    if err != nil {
    	log.Fatal(err)
    	tx.Rollback()
    	return
    }

    err = tx.Commit()
    if err != nil {
    	log.Fatal(err)
    }
    ```
8. 常用方法
    1. SetMaxOpenConns()，默认无限制，
设置最大打开的连接数，若到达则会阻塞操作，直到其他连接释放；
    2. SetMaxIdleConns()，默认为2，设置连接池最大闲置连接数，若到达则归还的连接会直接关闭；
    3. SetConnMaxLifetime()，如果小于0，则永不过期，闲置连接的最大生命周期，应小于数据库连接的超时时间；
        - 从连接创建时计时或连接池取连接时重新计时，连接只有在连接池中超时才会被清理掉
    4. Conn()，从连接池取出或者创建一个连接返回。在返回的sql.Conn上Ping,Exec,Query,QueryRow,Begin都是在当前连接上操作
        - sql.Open() 它没有建立与数据库的任何连接，也没有验证驱动程序的连接参数。取而代之的是，它只是准备数据库抽象以供以后使用；
        - sql.Conn连接关闭，会阻塞等连接上的事务完成，直到Tx.Commit()或者Tx.Rollback()
    5. Ping()，检查数据库是否可用和可访问；
    6. Close()，关闭连接。
9. 进阶，预查询时使用可变参数
    1. 拼接占位符
    ```
    var placeholder string
	idLen := len(idArr)
	for i := 0; i < idLen; i++ {
		if 0 == i && 1 < idLen {
			placeholder += "?,"
		} else {
			placeholder += "?"
		}
	}
    ```
    2. 预查询语句使用动态生成的占位符文本
    ```
    stmt, err := db.Prepare(fmt.Sprintf("DELETE FROM `user` WHERE `user_id` IN (%s)", placeholder))
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}
	defer stmt.Close()
    ```
    3. 使用可变参数，传递执行参数
    ```
    idLen = len(idArr)
	idSlice := make([]interface{}, idLen, idLen)
	for i := 0; i < idLen; i++ {
		idSlice[i] = idArr[i]
	}
	// 使用...运算符以变量...的形式进行参数传递；变量必须是与可变参数类型相同的slice
	rs, err := stmt.Exec(idSlice...)
	if err != nil {
		log.Fatalln(err)
		ret = -1
		return
	}
    ```

### 完整例子
```
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

func main() {
	con = &MysqlConnect{User: "root", Pwd: "12345678", Host: "10.211.55.3", Port: 3306, DbName: "db_test"}
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

	// 删除数据
	user.BatchDel(delUser)

	// 事务删除
	user.DelTx(id)
}
```

> [Go database/sql tutorial](http://go-database-sql.org/index.html)