package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	Boltdb()
}
func Boltdb() error {
	// Bolt 会在数据文件上获得一个文件锁，所以多个进程不能同时打开同一个数据库。
	// 打开一个已经打开的 Bolt 数据库将导致它挂起，直到另一个进程关闭它。
	// 为防止无限期等待，可以将超时选项传递给Open()函数：
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	//	两种处理方式：读-写和只读操作，读-写方式开始于db.Update方法：
	//	Bolt 一次只允许一个读-写事务，但是一次允许多个只读事务。
	// 每个事务处理都有一个始终如一的数据视图
	err = db.Update(func(tx *bolt.Tx) error {
		// 这里还有另外一层：k-v存储在bucket中，
		// 可以将bucket当做一个key的集合或者是数据库中的表。
		//（顺便提一句，buckets中可以包含其他的buckets，这将会相当有用）
		// Buckets 是键值对在数据库中的集合。所有在bucket中的key必须唯一。
		// 使用DB.CreateBucket() 函数建立buket
		// Tx.DeleteBucket() 删除bucket
		// b := tx.Bucket([]byte("MyBucket"))
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))

		//要将 key-value 对保存到 bucket，请使用 Bucket.Put() 函数
		//这将在 MyBucket 的 bucket 中将 "answer" key的值设置为"42"。
		err = b.Put([]byte("answer"), []byte("42"))
		err = b.Put([]byte("why"), []byte("101010"))
		return err
	})

	// 可以看到，传入db.update函数一个参数，在函数内部你可以get/set数据和处理error。
	// 如返回为nil，事务会从数据库得到一个commit，但如果返回一个实际错误，则会做回滚，
	// 在函数中做的事情都不会commit。这很自然，因为不需要人为地去关心事务的回滚，
	// 只需要返回一个错误，其他的由Bolt去完成。
	// 只读事务和读写事务不应该相互依赖，一般不应该在同一个例程中同时打开。
	// 这可能会导致死锁，因为读写事务需要定期重新映射数据文件，
	// 但只有在只读事务处于打开状态时才能这样做。

	// 批量读写事务。每一次新的事务都需要等待上一次事务的结束，
	// 可以通过DB.Batch()批处理来完成
	err = db.Batch(func(tx *bolt.Tx) error {
		return nil
	})

	//只读事务在db.View函数数中可以读取，但是不能做修改。
	db.View(func(tx *bolt.Tx) error {
		//要检索这个value，可以使用 Bucket.Get() 函数。
		//由于Get是有安全保障的，所以不会返回错误，不存在的key返回nil
		b := tx.Bucket([]byte("MyBucket"))
		//tx.Bucket([]byte("MyBucket")).Cursor() 可这样写
		v := b.Get([]byte("answer"))
		id, _ := b.NextSequence()
		fmt.Printf("The answer is: %s %d \n", v, id)

		//游标遍历key
		c := b.Cursor()
		fmt.Println("\n游标遍历key")
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		//游标上有以下函数：
		//First()  移动到第一个键。
		//Last()   移动到最后一个键。
		//Seek()   移动到特定的一个键。
		//Next()   移动到下一个键。
		//Prev()   移动到上一个键。

		//Prefix 前缀扫描
		fmt.Println("\nPrefix 前缀扫描")
		prefix := []byte("a")
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})

	// 如果知道所在桶中拥有键，也可以使用ForEach()来迭代：
	db.View(func(tx *bolt.Tx) error {
		fmt.Println("\nForEach()来迭代")
		b := tx.Bucket([]byte("MyBucket"))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})

	// 事务处理
	// 开始事务
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 使用事务
	_, err = tx.CreateBucket([]byte("MyBucket"))
	if err != nil {
		return err
	}

	// 事务提交
	if err = tx.Commit(); err != nil {
		return err
	}
	return err

	// 还可以在一个键中存储一个桶，以创建嵌套的桶：
	// func (*Bucket) CreateBucket(key []byte) (*Bucket, error)
	// func (*Bucket) CreateBucketIfNotExists(key []byte) (*Bucket, error)
	// func (*Bucket) DeleteBucket(key []byte) error
}
