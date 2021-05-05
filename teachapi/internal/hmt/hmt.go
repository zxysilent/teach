package hmt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// HmtAuth 凭证载体
type HmtAuth struct {
	Id    int    `json:"id"`
	Num   string `json:"num"`
	ExpAt int64  `json:"exp"`
}

// Encode 生成 token
func (jc *HmtAuth) Encode(key string) string {
	data, _ := json.Marshal(jc)
	bStr := base64.RawURLEncoding.EncodeToString(data)
	hm := hmac.New(sha1.New, []byte(key))
	hm.Write([]byte(bStr))
	sign := base64.RawURLEncoding.EncodeToString(hm.Sum(nil))
	return bStr + "." + sign
}

// Verify 验证
func Verify(raw string, key string) (*HmtAuth, error) {
	parts := strings.Split(raw, ".")
	if len(parts) != 2 {
		return nil, errors.New("非法的token: " + raw)
	}
	hm := hmac.New(sha1.New, []byte(key))
	hm.Write([]byte(parts[0]))
	sign := base64.RawURLEncoding.EncodeToString(hm.Sum(nil))
	if sign != parts[1] {
		return nil, errors.New("token 非法")
	}
	datab, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, errors.New("base64解码失败,error: " + err.Error())
	}
	ha := &HmtAuth{}
	err = json.Unmarshal(datab, ha)
	if err != nil {
		return nil, errors.New("json解码失败,error: " + err.Error())
	}
	if time.Now().Unix() > ha.ExpAt {
		return nil, errors.New("toekn 已经过期")
	}
	return ha, nil
}
