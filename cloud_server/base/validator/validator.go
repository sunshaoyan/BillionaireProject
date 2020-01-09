package validator

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"strconv"
	"sync"
	"unicode/utf8"
	"xdms/base/errs"
	"xdms/base/types"
)

func Init() {
	binding.Validator = new(defaultValidator)
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	if valueType == reflect.Struct {
		tmp := reflect.TypeOf(obj).Elem()
		tmpVal := reflect.ValueOf(obj).Elem()
		for idx := 0; idx < tmp.NumField(); idx++ {
			nameLen := tmp.Field(idx).Tag.Get("strlen")
			jsonTag := tmp.Field(idx).Tag.Get("json")
			if jsonTag == "" {
				jsonTag = tmp.Field(idx).Name
			}
			if nameLen != "" {
				strVal := ""
				i, err := strconv.Atoi(nameLen)
				if err != nil {
					return err
				}
				i2 := tmpVal.Field(idx).Interface()
				switch i2.(type) {
				case string:
					strVal = i2.(string)
				case types.UnRequire:
					strVal = i2.(types.UnRequire).Value
				case types.TrimString:
					strVal = i2.(types.TrimString).String()
				}
				if utf8.RuneCountInString(strVal) > i {
					return errs.ParameterError.AddMsgf(jsonTag + " length validate error")
				}
			}
		}
	}

	return nil
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://godoc.org/gopkg.in/go-playground/validator.v8
func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		config := &validator.Config{TagName: "binding"}
		v.validate = validator.New(config)
		v.validate.RegisterStructValidation(types.UnRequireStructLevelValidation, types.UnRequire{})
	})
}
