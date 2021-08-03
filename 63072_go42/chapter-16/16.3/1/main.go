package main

import (
	"fmt"
	"strconv"

	"crypto/md5"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var md = md5.New()

// 测试专用
func Read(db *leveldb.DB, num int) {
	var kStr string
	var haskKey string
	kStr = strconv.Itoa(num)
	md.Write([]byte(kStr))
	haskKey = fmt.Sprintf("%x", md.Sum(nil))
	md.Reset()

	db.Get([]byte(haskKey), nil)
}

// 测试专用
func Write(db *leveldb.DB, num int) {
	var kStr string
	var haskKey string
	kStr = strconv.Itoa(num)
	md.Write([]byte(kStr))
	haskKey = fmt.Sprintf("%x", md.Sum(nil))
	md.Reset()

	db.Put([]byte(haskKey), []byte(kStr), nil)
}

func main() {
	// 打开数据库文件 /path/to/db ,第一个参数为存放数据的目录，不是具体文件，
	// o := &opt.Options{	Filter: filter.NewBloomFilter(10),}
	// OpenFile第2个参数这里指定为nil，在数据集大时可设置如布隆过滤器。
	// *opt.Options 为nil默认为false ，true为只读模式ReadOnly
	db, _ := leveldb.OpenFile("levdb", nil)

	defer db.Close()

	// 读数据库:Get(key,nil)，写数据库:Put(key,value,nil)
	// Put第三个参数为nil，默认设置即可，默认时写的时候如果机器崩溃了数据会丢失。
	// key和value都是字节切片
	_ = db.Put([]byte("key1"), []byte("好好检查"), nil)
	_ = db.Put([]byte("key2"), []byte("天天向上"), nil)
	_ = db.Put([]byte("key:3"), []byte("就会一个本事"), nil)
	_ = db.Put([]byte("uname"), []byte("Jim"), nil)
	_ = db.Put([]byte("time"), []byte("1450932202"), nil)

	// 读数据库:Get(key,nil)，返回字节切片。
	data, _ := db.Get([]byte("key1"), nil)
	fmt.Println("key1=>", string(data))

	// 删除某个key(key,nil)，key不存在时并不返回错误。
	_ = db.Delete([]byte("key"), nil)

	// 迭代数据库内容。
	iter := db.NewIterator(nil, nil)
	fmt.Println("迭代所有key/value")
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), "=>", string(value))

	}
	iter.Release()
	iter.Error()

	// Seek()定位到比给定key值(字节值)大的第一个key，可迭代所有筛选出的key-value。
	iter = db.NewIterator(nil, nil)
	fmt.Println("\nSeek()按值筛选查找key")
	for ok := iter.Seek([]byte("t")); ok; ok = iter.Next() {
		// Use key/value.
		fmt.Println("Seek-then-Iterate:")
		fmt.Println(string(iter.Key()), "=>", string(iter.Value()))
	}
	iter.Release()

	// 迭代内容子集:start表示key中包含有的字符串，Limit表示key不能包含有字符串。
	fmt.Println("\n 按照指定（排除）条件筛选key")
	iter = db.NewIterator(&util.Range{Start: []byte("key"), Limit: []byte("no")}, nil)
	for iter.Next() {
		// Use key/value.
		fmt.Println("Iterate over subset of database content:")
		fmt.Println(string(iter.Key()), "=>", string(iter.Value()))
	}
	iter.Release()

	// 迭代子集内容，key的前缀是指定字符串。
	fmt.Println("\n 查找指定前缀key")
	iter = db.NewIterator(util.BytesPrefix([]byte("key")), nil)
	for iter.Next() {
		// Use key/value.
		fmt.Println("Iterate over subset of database content with a particular prefix:")
		fmt.Println(string(iter.Key()), "=>", string(iter.Value()))
	}
	iter.Release()

	_ = iter.Error()

	// 批量写:
	batch := new(leveldb.Batch)
	var kStr string
	var batchkey string
	for i := 0; i < 10; i++ {
		kStr = strconv.Itoa(i)
		md.Write([]byte(kStr))
		batchkey = fmt.Sprintf("%x", md.Sum(nil))
		batch.Put([]byte(batchkey), []byte(kStr))
	}
	md.Reset()
	batch.Delete([]byte("lazy"))
	_ = db.Write(batch, nil)
}
