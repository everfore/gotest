package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name     string
	Nickname []string
	Profile  Profile
	PProfile *Profile
	IProfile interface{}
	MProfile map[string]interface{}
	Msg      chan bool
	Do       func(name string) bool
}

func DoFunc(name string) bool {
	fmt.Println(name)
	return true
}
func NewUser() User {
	return User{
		Name:     "foo",
		Nickname: []string{"nick1", "nick2"},
		Profile:  Profile{},
		PProfile: new(Profile),
		IProfile: new(Profile),
		MProfile: make(map[string]interface{}),
		Msg:      make(chan bool, 10),
		Do:       DoFunc,
	}
}

type Profile struct {
	Id int32
}

func main() {
	user := NewUser()
	rf_ptr(&user)
}

func rf_ptr(user interface{}) {
	vu := reflect.ValueOf(user).Elem()
	fmt.Println(vu.Field(0).CanSet(), vu.Field(0).Type().Name())
	vu.Field(0).SetString("refsetstring")
	vu.Field(1).SetLen(1)
	vu.Field(2).Set(reflect.ValueOf(Profile{Id: 10}))
	vu.Field(3).Set(reflect.ValueOf(new(Profile)))
	vu.Field(4).Set(reflect.ValueOf(make([]byte, 10)))
	vu.Field(5).SetMapIndex(reflect.ValueOf("key-11"), reflect.ValueOf("val-11"))
	vu.Field(6).Send(reflect.ValueOf(true))

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf("mee")
	out := vu.Field(7).Call(in)
	fmt.Println(out)
	fmt.Println(user)
	fmt.Println(<-user.(*User).Msg)

	fmt.Println(reflect.New(vu.Type()))
}
