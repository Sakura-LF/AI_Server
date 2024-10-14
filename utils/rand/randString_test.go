package rand

import "testing"

func TestGetRandomString(t *testing.T) {

	for i := 0; i < 10; i++ {
		name, err := GetRandomUserName()
		if err != nil {
			t.Log(err)
			return
		}
		t.Log("UserName", name)
		nickName := GetRandomNickName(name)
		t.Log("NickName", nickName)
	}
}
