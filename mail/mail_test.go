package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	SendMail("golang test")
}
