package OSMsg2

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"testing"
)

var login = MsgType_Login
var osmsg *OSMsg
var loginMsg *LoginMsg
var loginMsgBytes []byte

func init() {
	osmsg = &OSMsg{Pattern: &login}
	loginMsg = &LoginMsg{
		Pattern: &login,
		From:    proto.Int32(111),
		To:      proto.Int32(112),
		Msg:     proto.String("hello, world."),
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
		// bs, err := proto.Marshal(loginMsg)
		// if checkerr(err) {
		// 	b.Error(err)
		// } else if !reflect.DeepEqual(bs, loginMsgBytes) {
		// 	b.Error("not equal")
		// }
	}
}

func BenchmarkUnMarshal(b *testing.B) {
	var new_loginMsg OSMsg
	var loginMsg_ LoginMsg
	for i := 0; i < b.N; i++ {
		_ = proto.Unmarshal(loginMsgBytes, &new_loginMsg)
		switch new_loginMsg.GetPattern() {
		case MsgType_Group:
			fmt.Println("group")
		case MsgType_Individual:
			fmt.Println("Individual")
		case MsgType_Logup:
			fmt.Println("Logup")
		case MsgType_Logout:
			fmt.Println("Logout")
		case MsgType_Login:
			_ = proto.Unmarshal(loginMsgBytes, &loginMsg_)
		}
	}
}

// func TestBytes(t *testing.T) {
// 	b, err := proto.Marshal(loginMsg)
// 	checkerr(err)
// 	fmt.Println(b)
// 	fmt.Printf("loginMsg:\n%v\n%v\n\n", loginMsg, loginMsg.String())

// 	b2, err2 := proto.Marshal(loginMsg)
// 	checkerr(err2)
// 	fmt.Println(b2)
// 	fmt.Printf("loginMsg:\n%v\n%v\n\n", loginMsg, loginMsg.String())

// }
