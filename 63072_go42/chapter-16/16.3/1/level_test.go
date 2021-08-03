package main

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestWrite(t *testing.T) {
	db, _ := leveldb.OpenFile("levdb", nil)
	defer db.Close()
	Write(db, 6)
	t.Log("TestWrite: 完成")
}

func TestRead(t *testing.T) {
	db, _ := leveldb.OpenFile("levdb", nil)
	defer db.Close()
	Read(db, 7)
	t.Log("TestRead：完成")
}

func BenchmarkWrite(b *testing.B) {
	db, _ := leveldb.OpenFile("levdb", nil)
	defer db.Close()
	b.N = 500000
	for i := 0; i < b.N; i++ {
		Write(db, i)
	}
}

func BenchmarkRead(b *testing.B) {
	db, _ := leveldb.OpenFile("levdb", nil)
	defer db.Close()
	b.N = 500000
	for i := 0; i < b.N; i++ {
		Read(db, i)
	}
}
