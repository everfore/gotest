package OSMsg

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"testing"
)

var login = OSMsg_Login
var osmsg *OSMsg
var inner_loginMsg *OSMsg_LoginMsg
var loginMsg *OSMsg
var loginMsgBytes []byte

func init() {
	osmsg = &OSMsg{Pattern: &login}
	inner_loginMsg = &OSMsg_LoginMsg{
		From: proto.Int32(111),
		To:   proto.Int32(112),
		Msg:  proto.String("hello, world."),
	}
	loginMsg = &OSMsg{
		Pattern: &login,
		Login:   inner_loginMsg,
	}
	loginMsgBytes, _ = proto.Marshal(loginMsg)
}

func checkerr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		proto.Marshal(loginMsg)
	}
}

func BenchmarkUnMarshal(b *testing.B) {
	var new_loginMsg OSMsg
	for i := 0; i < b.N; i++ {
		_ = proto.Unmarshal(loginMsgBytes, &new_loginMsg)
	}
}

// func TestBytes(t *testing.T) {
// 	b, err := proto.Marshal(loginMsg)
// 	checkerr(err)
// 	fmt.Println(b)
// 	fmt.Printf("loginMsg:\n%v\n%v\n\n", loginMsg, loginMsg.String())

// 	b2, err2 := proto.Marshal(inner_loginMsg)
// 	checkerr(err2)
// 	fmt.Println(b2)
// 	fmt.Printf("loginMsg:\n%v\n%v\n\n", inner_loginMsg, inner_loginMsg.String())

// }
