package hmt

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	auth := HmtAuth{
		Id:    1,
		Num:   "zxysilent",
		ExpAt: time.Now().Add(time.Hour * 2).Unix(),
	}
	t.Log(auth.Encode("key"))
}
func TestVerify(t *testing.T) {
	raw := "eyJpZCI6MSwibnVtIjoienh5c2lsZW50IiwiZXhwIjoxNjE0MDc1NjQ1fQ.GgaM0B6S7qF8dYLa0aOlV6ti6Kc"
	hmtAuth, err := Verify(raw, "key")
	t.Log(hmtAuth, err)
}
func BenchmarkEncode(b *testing.B) {
	auth := HmtAuth{
		Id:    1,
		Num:   "zxysilent",
		ExpAt: time.Now().Add(time.Hour * 2).Unix(),
	}
	for i := 0; i < b.N; i++ {
		auth.Encode("key")
	}
}
func BenchmarkVerify(b *testing.B) {
	raw := "eyJpZCI6MSwibnVtIjoienh5c2lsZW50IiwiZXhwIjoxNjE0MDc1NjQ1fQ.GgaM0B6S7qF8dYLa0aOlV6ti6Kc"
	for i := 0; i < b.N; i++ {
		Verify(raw, "key")
	}
}
