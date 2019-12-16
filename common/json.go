package common

import "encoding/json"

//序列化
func Bytes(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err == nil {
		return data
	} else {
		Error.Println("Bytes error", v, err)
		return nil
	}
}
