package types

import (
	"bytes"
	"encoding/json"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"unicode/utf8"
)

type M map[string]interface{}

func ConvertToM(res map[string]interface{}) M {
	m := M{}
	for key, val := range res {
		if newV, ok := val.(map[string]interface{}); ok {
			m[key] = ConvertToM(newV)
		} else {
			m[key] = val
		}
	}
	return m
}

type UnRequire struct {
	Value  string
	IsNull bool
	Chose  bool
}

func (u *UnRequire) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	tmp := json.RawMessage{}
	err := dec.Decode(&tmp)
	if err == nil && string(tmp) == "null" {
		*u = UnRequire{IsNull: true}
		return nil
	}
	*u = UnRequire{Chose: true}
	dec2 := json.NewDecoder(bytes.NewReader(data))
	return dec2.Decode(&(u.Value))
}

func (u *UnRequire) HasValue() bool {
	if u.IsNull {
		return true
	}
	if u.Chose {
		return true
	}
	return false
}

func (u UnRequire) MarshalJSON() (data []byte, err error) {
	return json.Marshal(u.Value)
}

func (u UnRequire) String() string {
	return u.Value
}

func UnRequireStructLevelValidation(v *validator.Validate, structLevel *validator.StructLevel) {
	res := structLevel.CurrentStruct.Interface().(UnRequire)
	if res.IsNull {
		return
	}
	if len(res.Value) == 0 && res.Chose {
		structLevel.ReportError(reflect.ValueOf(res.Value), "Value", "value", "value")
	}
	if utf8.RuneCountInString(res.Value) > 512 {
		structLevel.ReportError(reflect.ValueOf(res.Value), "Value", "value", "value")
	}
	if res.Chose && len(res.Value) == 0 {
		structLevel.ReportError(reflect.ValueOf(res.Value), "Value", "value", "value")
	}
}

// -----------------------------------------------------------------------------

type TrimString string

func (ts TrimString) MarshalJSON() ([]byte, error) {
	str := string(ts)
	// 去掉开头和结尾的空格
	//tsd := strings.TrimSpace(str)
	return json.Marshal(str)
}

func (ts *TrimString) UnmarshalJSON(byte_ []byte) error {
	var s string
	// 先转换成 string
	err := json.Unmarshal(byte_, &s)
	if err != nil {
		return err
	}
	// 去掉空格
	//str := strings.TrimSpace(s)
	*ts = TrimString(s)
	return nil
}

func (ts TrimString) String() string {
	return string(ts)
}
