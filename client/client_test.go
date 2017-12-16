package client

import "testing"

func TestFnoName(t *testing.T) {
	ltox := NewLigTox()
	ltox.FriendSendMessage(1, "aa")
}
