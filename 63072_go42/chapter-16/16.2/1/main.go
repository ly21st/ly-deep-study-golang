package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

type Cate struct {
	cid     int
	cname   string
	addtime int
	scope   int
}

func main() {
	// 注意修改mydb为自己的数据库名
	dbw := DbWorker{Dsn: "root:123456@tcp(localhost:3306)/mydb?charset=utf8mb4"}
	// 支持下面几种DSN写法，具体看mysql服务端配置，常见为第2种
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	dbtemp, err := sql.Open("mysql", dbw.Dsn)
	dbw.Db = dbtemp

	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	// 插入数据测试
	dbw.insertData()

	// 删除数据测试
	dbw.deleteData()

	// 修改数据测试
	dbw.editData()

	// 查询数据测试
	dbw.queryData()

	// 事务操作测试
	dbw.transaction()
}

// 插入数据，sql预编译
func (dbw *DbWorker) insertData() {
	stmt, _ := dbw.Db.Prepare(`INSERT INTO t_article_cate (cname, addtime, scope) VALUES (?, ?, ?)`)
	defer stmt.Close()

	ret, err := stmt.Exec("栏目1", time.Now().Unix(), 10)
	// 通过返回的ret可以进一步查询本次插入数据影响的行数
	// RowsAffected和最后插入的Id(如果数据库支持查询最后插入Id)
	if err != nil {
		fmt.Printf("插入数据出错： %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("最后插入记录的ID：", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("插入有效的行数：", RowsAffected)
	}
}

// 删除数据，预编译
func (dbw *DbWorker) deleteData() {
	stmt, err := dbw.Db.Prepare(`DELETE FROM t_article_cate WHERE cid=?`)
	ret, err := stmt.Exec(122)
	// 通过返回的ret可以进一步查询本次删除数据影响的行数RowsAffected
	if err != nil {
		fmt.Printf("删除数据出错： %v\n", err)
		return
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("删除有效的行数：:", RowsAffected)
	}
}

// 修改数据，预编译
func (dbw *DbWorker) editData() {
	stmt, err := dbw.Db.Prepare(`UPDATE t_article_cate SET scope=? WHERE cid=?`)
	ret, err := stmt.Exec(111, 123)
	// 通过返回的ret可以进一步查询本次修改数据影响的行数RowsAffected
	if err != nil {
		fmt.Printf("修改数据出错： %v\n", err)
		return
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("修改有效的行数：:", RowsAffected)
	}
}

// 查询数据，预编译
func (dbw *DbWorker) queryData() {
	// 如果方法包含Query，那么这个方法是用于查询并返回rows的。其他用Exec()
	// 另外一种写法
	// rows, err := db.Query("select id, name from users where id = ?", 1)
	stmt, _ := dbw.Db.Prepare(`SELECT cid, cname, addtime, scope From t_article_cate where status=?`)
	// err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	// 单行查询，直接处理
	defer stmt.Close()

	rows, err := stmt.Query(0)
	defer rows.Close()
	if err != nil {
		fmt.Printf("查询数据出错： %v\n", err)
		return
	}

	columns, _ := rows.Columns() // 得到字段名
	fmt.Println("字段名：", columns)

	// 下面两种取数据方式都比较原生，实际中一般都选择ORM框架
	// 1.数据保存到结构体
	cate := Cate{}
	for rows.Next() {
		//将行数据保存到结构体
		err = rows.Scan(&cate.cid, &cate.cname, &cate.addtime, &cate.scope)
		fmt.Println(cate)
	}
	/*
		// 2.通过slice取数据
		rowMaps := make([]map[string]string, 4) // 数据key:value的slice
		values := make([]sql.RawBytes, len(columns))
		scans := make([]interface{}, len(columns))
		for i := range values {
			scans[i] = &values[i]
		}
		i := 0
		for rows.Next() {
			//将行数据保存到字典
			err = rows.Scan(scans...)          // 保存当前数据行row的值到slice
			each := make(map[string]string, 4) // 定义map(key:value)

			for i, col := range values {
				each[columns[i]] = string(col) // 写k/v 数据
			}

			// rowMaps切片追加数据，索引位置有意思。不这样写就不是希望的样子。
			rowMaps = append(rowMaps[:i], each)
			fmt.Println(each)
			i++
		}

		for i, col := range rowMaps {
			fmt.Println(i, col)
		}
	*/
	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// 事务处理，预编译
func (dbw *DbWorker) transaction() {
	tx, err := dbw.Db.Begin()
	if err != nil {
		fmt.Printf("事务处理数据出错： %v\n", err)
		return
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(`INSERT INTO t_article_cate (cname, addtime, scope) VALUES (?, ?, ?)`)
	if err != nil {
		fmt.Printf("事务处理数据出错： %v\n", err)
		return
	}

	for i := 100; i < 110; i++ {
		cname := strings.Join([]string{"栏目-", string(i)}, "-")
		_, err = stmt.Exec(cname, time.Now().Unix(), i+20)
		if err != nil {
			fmt.Printf("事务处理插入数据出错： %v\n", err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("事务处理提交出错： %v\n", err)
		return
	}
	stmt.Close()
	fmt.Println("事务处理成功完成\n")
}
