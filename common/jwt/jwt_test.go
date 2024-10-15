package jwt

import "testing"

func TestGenToken(t *testing.T) {
	payLoad := &PayLoad{
		UserId: 1,
		Role:   1,
	}
	token, err := GenToken(*payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(claims)
	}
}
