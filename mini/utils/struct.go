package utils

import (
	"encoding/json"
	"reflect"
)

//支持struct，map，json字符串之间互转
func Transfer(beConverted, to interface{}) error {
	t:=reflect.TypeOf(beConverted)
	if t.Name()=="string" {
		err:=json.Unmarshal([]byte(beConverted.(string)),to)
		if err != nil {
			return err
		}
	}else {
		bt,err:=json.Marshal(beConverted)
		if err != nil {
			return err
		}
		err=json.Unmarshal(bt,to)
		if err != nil {
			return err
		}
	}
	return nil
}