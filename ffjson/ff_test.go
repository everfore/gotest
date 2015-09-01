package ffjson

import (
	"testing"

	"encoding/json"

	"github.com/pquerna/ffjson/ffjson"
)

type User1 struct {
	Name string
	Age  byte
	Buf  [1000]byte
}

type User2 struct {
	Name string
	Age  byte
	Buf  [1000]byte
}

var user1 = User2{Name: "jack", Age: 20}
var user2 = User1{Name: "jack", Age: 20, Buf: [1000]byte{}}
var std_user1, _ = json.Marshal(user1)
var ffj_user1, _ = ffjson.Marshal(user2)

func init() {
	// runtime.GOMAXPROCS(1)
}

func BenchmarkMarshalStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(user1)
	}
}

func BenchmarkMarshalFfjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ffjson.Marshal(user1)
	}
}

func BenchmarkMarshalStd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(user2)
	}
}

func BenchmarkMarshalFfjson2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ffjson.Marshal(user2)
	}
}

func BenchmarkUnMarshalStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(std_user1, &user1)
	}
}

func BenchmarkUnMarshalFfjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(std_user1, &user1)
	}
}

func BenchmarkUnMarshalStd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(ffj_user1, &user2)
	}
}

func BenchmarkUnMarshalFfjson2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(ffj_user1, &user2)
	}
}
