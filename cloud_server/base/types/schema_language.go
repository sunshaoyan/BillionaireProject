package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"unicode/utf8"
)

const (
	INT     = "int"
	TEXT    = "text"
	FLOAT   = "float"
	BOOLEAN = "bool"
	ARRAY   = "array"
	STRUCT  = "struct"
	RAW     = "raw"
)

type TypeSchema interface {
	Refine([]byte) (interface{}, error)
	DefaultValue() (interface{}, bool)
	MarshalBSON() ([]byte, error)
	UnmarshalBSON([]byte) error
}

type DataSchema struct {
	Name       string     `json:"name" bson:"name"  binding:"required"`
	Optional   bool       `json:"optional" bson:"optional"`
	Type       string     `json:"type" bson:"type"  binding:"required"`
	TypeSchema TypeSchema `json:"spec" bson:"spec"`
}

func (ds *DataSchema) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	out["name"] = ds.Name
	if ds.Optional {
		out["optional"] = ds.Optional
	}
	out["type"] = ds.Type
	if ds.TypeSchema != nil {
		out["spec"] = ds.TypeSchema
	}
	return bson.Marshal(out)
}

func (ds *DataSchema) UnmarshalBSON(b []byte) error {
	var ras struct {
		Name        string        `bson:"name"`
		Optional    bool          `bson:"optional"`
		Type        string        `bson:"type"`
		RawTypeSpec bson.RawValue `bson:"spec"`
	}
	if err := bson.Unmarshal(b, &ras); err != nil {
		return err
	}
	ds.Name = ras.Name
	ds.Optional = ras.Optional
	switch ras.Type {
	case INT:
		var is IntSpec
		if ras.RawTypeSpec.Value != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec.Value, &is); err != nil {
				return err
			}
		}
		ds.TypeSchema = &is
	case TEXT:
		var ts TextSpec
		if ras.RawTypeSpec.Value != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec.Value, &ts); err != nil {
				return err
			}
		}
		ds.TypeSchema = &ts
	case FLOAT:
		var fs FloatSpec
		if ras.RawTypeSpec.Value != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec.Value, &fs); err != nil {
				return err
			}
		}
		ds.TypeSchema = &fs
	case BOOLEAN:
		var bs BoolSpec
		if ras.RawTypeSpec.Value != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec.Value, &bs); err != nil {
				return err
			}
		}
		ds.TypeSchema = &bs
	case ARRAY:
		var as ArraySpec
		if ras.RawTypeSpec.Value == nil {
			return errors.New("array spec optional")
		}
		if err := bson.Unmarshal(ras.RawTypeSpec.Value, &as); err != nil {
			return err
		}
		ds.TypeSchema = &as
	case STRUCT:
		var ss StructSpec
		if ras.RawTypeSpec.Value == nil {
			return errors.New("struct spec optional")
		}
		if err := bson.Unmarshal(ras.RawTypeSpec.Value, &ss); err != nil {
			return err
		}
		ds.TypeSchema = &ss
	case RAW:
		ds.TypeSchema = &RawSpec{}
	default:
		return fmt.Errorf("invalid data type %s", ras.Type)
	}
	ds.Type = ras.Type
	return nil
}

func (ds *DataSchema) UnmarshalJSON(b []byte) error {
	var ras struct {
		Name        string          `json:"name"`
		Optional    bool            `json:"optional"`
		Type        string          `json:"type"`
		RawTypeSpec json.RawMessage `json:"spec"`
	}

	if err := json.Unmarshal(b, &ras); err != nil {
		return err
	}
	ds.Name = ras.Name
	ds.Optional = ras.Optional
	switch ras.Type {
	case INT:
		var is IntSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &is); err != nil {
				return err
			}
		}
		ds.TypeSchema = &is
	case TEXT:
		var ts TextSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &ts); err != nil {
				return err
			}
		}
		ds.TypeSchema = &ts
	case FLOAT:
		var fs FloatSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &fs); err != nil {
				return err
			}
		}
		ds.TypeSchema = &fs
	case BOOLEAN:
		var bs BoolSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &bs); err != nil {
				return err
			}
		}
		ds.TypeSchema = &bs
	case ARRAY:
		var as ArraySpec
		if ras.RawTypeSpec == nil {
			return errors.New("array spec optional")
		}
		if err := json.Unmarshal(ras.RawTypeSpec, &as); err != nil {
			return err
		}
		ds.TypeSchema = &as
	case STRUCT:
		var ss StructSpec
		if ras.RawTypeSpec == nil {
			return errors.New("struct spec optional")
		}
		if err := json.Unmarshal(ras.RawTypeSpec, &ss); err != nil {
			return err
		}
		ds.TypeSchema = &ss
	case RAW:
		ds.TypeSchema = &RawSpec{}
	default:
		return fmt.Errorf("invalid data type %s", ras.Type)
	}
	ds.Type = ras.Type
	return nil
}

func (ds *DataSchema) Refine(v []byte) (interface{}, error) {
	return ds.TypeSchema.Refine(v)
}

func (ds *DataSchema) DefaultValue() (interface{}, bool) {
	return ds.TypeSchema.DefaultValue()
}

type RawSpec struct {
}

func (rs *RawSpec) MarshalBSON() ([]byte, error) {
	return []byte("{}"), nil
}
func (rs *RawSpec) UnmarshalBSON(b []byte) error {
	return bson.Unmarshal(b, rs)
}

func (rs *RawSpec) Refine(v []byte) (interface{}, error) {
	var m interface{}
	err := bson.Unmarshal(v, &m)
	if err != nil {
		return nil, errors.New("invalid format of raw message")
	}
	return bson.Raw(v), nil
}

func (rs *RawSpec) DefaultValue() (interface{}, bool) {
	return nil, false
}

type StructSpec struct {
	Fields []DataSchema `bson:"fields"`
}

func (ss *StructSpec) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	out["fields"] = ss.Fields
	return bson.Marshal(out)
}

func (ss *StructSpec) UnmarshalBSON(b []byte) error {
	var (
		rss struct {
			Fields []bson.Raw `bson:"fields"`
		}
		memberSet = make(map[string]struct{})
	)
	if err := bson.Unmarshal(b, &rss); err != nil {
		return err
	}
	if len(rss.Fields) == 0 {
		return errors.New("struct spec empty")
	}
	for i := range rss.Fields {
		var ds DataSchema
		if err := bson.Unmarshal(rss.Fields[i], &ds); err != nil {
			return err
		}
		if _, ok := memberSet[ds.Name]; ok {
			return fmt.Errorf("struct has duplicate field %s", ds.Name)
		}
		memberSet[ds.Name] = struct{}{}
		ss.Fields = append(ss.Fields, ds)
	}
	return nil
}

func (ss *StructSpec) UnmarshalJSON(b []byte) error {
	var (
		rss struct {
			Fields []json.RawMessage `json:"fields"`
		}
		memberSet = make(map[string]struct{})
	)
	if err := json.Unmarshal(b, &rss); err != nil {
		return err
	}
	if len(rss.Fields) == 0 {
		return errors.New("struct spec empty")
	}
	for i := range rss.Fields {
		var ds DataSchema
		if err := json.Unmarshal(rss.Fields[i], &ds); err != nil {
			return err
		}
		if _, ok := memberSet[ds.Name]; ok {
			return fmt.Errorf("struct has duplicate field %s", ds.Name)
		}
		memberSet[ds.Name] = struct{}{}
		ss.Fields = append(ss.Fields, ds)
	}
	return nil
}

func (ss *StructSpec) Refine(v []byte) (interface{}, error) {
	var (
		rtMap     = make(map[string]interface{})
		rawStruct map[string]json.RawMessage
	)
	err := json.Unmarshal(v, &rawStruct)
	if err != nil {
		return nil, errors.New("invalid format of struct")
	}
	for _, schema := range ss.Fields {
		rawMember, ok := rawStruct[schema.Name]
		if !ok {
			if defaultValue, ok := schema.DefaultValue(); ok {
				rtMap[schema.Name] = defaultValue
				continue
			}
			if !schema.Optional {
				return nil, fmt.Errorf("struct required field '%s' not placed", schema.Name)
			}
		} else {
			v, err := schema.TypeSchema.Refine(rawMember)
			if err != nil {
				return nil, fmt.Errorf("struct field %s invalid:%v", schema.Name, err)
			}
			rtMap[schema.Name] = v
		}
	}
	return rtMap, nil
}

func (ss *StructSpec) DefaultValue() (interface{}, bool) {
	rtMap := make(map[string]interface{})
	for _, schema := range ss.Fields {
		defaultValue, ok := schema.DefaultValue()
		if ok {
			rtMap[schema.Name] = defaultValue
			continue
		}
		if !schema.Optional {
			return nil, false
		}
	}
	return rtMap, true
}

type ArraySpec struct {
	MinSize    int        `json:"min_size" bson:"min_size"`
	MaxSize    int        `json:"max_size" bson:"max_size"`
	MinSet     bool       `json:"-" bson:"min_placed"`
	MaxSet     bool       `json:"-" bson:"max_placed"`
	Type       string     `json:"type" bson:"type"`
	TypeSchema TypeSchema `json:"spec" bson:"spec"`
}

func (as *ArraySpec) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	if as.MinSet {
		out["min_length"] = as.MinSize
	}
	if as.MaxSet {
		out["max_length"] = as.MaxSize
	}
	out["type"] = as.Type
	if as.TypeSchema != nil {
		out["spec"] = as.TypeSchema
	}
	return bson.Marshal(out)
}

func (as *ArraySpec) UnmarshalBSON(b []byte) error {
	var ras struct {
		MinSize     int      `bson:"min_length"`
		MaxSize     int      `bson:"max_length"`
		MinSet      bool     `bson:"min_placed"`
		MaxSet      bool     `bson:"max_placed"`
		Type        string   `bson:"type"`
		RawTypeSpec bson.Raw `bson:"spec"`
	}

	if err := bson.Unmarshal(b, &ras); err != nil {
		return err
	}
	as.MinSize = ras.MinSize
	as.MinSet = ras.MinSet
	as.MaxSize = ras.MaxSize
	as.MaxSet = ras.MaxSet
	as.Type = ras.Type

	switch ras.Type {
	case INT:
		var is IntSpec
		if ras.RawTypeSpec != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec, &is); err != nil {
				return err
			}
		}
		as.TypeSchema = &is
	case TEXT:
		var ts TextSpec
		if ras.RawTypeSpec != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec, &ts); err != nil {
				return err
			}
		}
		as.TypeSchema = &ts
	case FLOAT:
		var fs FloatSpec
		if ras.RawTypeSpec != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec, &fs); err != nil {
				return err
			}
		}
		as.TypeSchema = &fs
	case BOOLEAN:
		var bs BoolSpec
		if ras.RawTypeSpec != nil {
			if err := bson.Unmarshal(ras.RawTypeSpec, &bs); err != nil {
				return err
			}
		}
		as.TypeSchema = &bs
	case ARRAY:
		var as ArraySpec
		if ras.RawTypeSpec == nil {
			return errors.New("array spec optional")
		}
		if err := bson.Unmarshal(ras.RawTypeSpec, &as); err != nil {
			return err
		}
		as.TypeSchema = &as
	case STRUCT:
		var ss StructSpec
		if ras.RawTypeSpec == nil {
			return errors.New("struct spec optional")
		}
		if err := bson.Unmarshal(ras.RawTypeSpec, &ss); err != nil {
			return err
		}
		as.TypeSchema = &ss
	case RAW:
		as.TypeSchema = &RawSpec{}
	default:
		return fmt.Errorf("invalid type of array spec")
	}
	return nil
}

func (as *ArraySpec) UnmarshalJSON(b []byte) error {
	var ras struct {
		MinSize     int             `json:"min_length"`
		MaxSize     int             `json:"max_length"`
		MinSet      bool            `json:"min_placed"`
		MaxSet      bool            `json:"max_placed"`
		Type        string          `json:"type"`
		RawTypeSpec json.RawMessage `json:"spec"`
	}

	if err := json.Unmarshal(b, &ras); err != nil {
		return err
	}
	as.MinSize = ras.MinSize
	as.MinSet = ras.MinSet
	as.MaxSize = ras.MaxSize
	as.MaxSet = ras.MaxSet
	as.Type = ras.Type

	switch ras.Type {
	case INT:
		var is IntSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &is); err != nil {
				return err
			}
		}
		as.TypeSchema = &is
	case TEXT:
		var ts TextSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &ts); err != nil {
				return err
			}
		}
		as.TypeSchema = &ts
	case FLOAT:
		var fs FloatSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &fs); err != nil {
				return err
			}
		}
		as.TypeSchema = &fs
	case BOOLEAN:
		var bs BoolSpec
		if ras.RawTypeSpec != nil {
			if err := json.Unmarshal(ras.RawTypeSpec, &bs); err != nil {
				return err
			}
		}
		as.TypeSchema = &bs
	case ARRAY:
		var as ArraySpec
		if ras.RawTypeSpec == nil {
			return errors.New("array spec optional")
		}
		if err := json.Unmarshal(ras.RawTypeSpec, &as); err != nil {
			return err
		}
		as.TypeSchema = &as
	case STRUCT:
		var ss StructSpec
		if ras.RawTypeSpec == nil {
			return errors.New("struct spec optional")
		}
		if err := json.Unmarshal(ras.RawTypeSpec, &ss); err != nil {
			return err
		}
		as.TypeSchema = &ss
	case RAW:
		as.TypeSchema = &RawSpec{}
	default:
		return fmt.Errorf("invalid type of array spec")
	}
	return nil
}

func (as *ArraySpec) Refine(v []byte) (interface{}, error) {
	rtArray := make([]interface{}, 0)

	var rawArray []json.RawMessage
	err := json.Unmarshal(v, &rawArray)
	if err != nil {
		return nil, errors.New("invalid format of array")
	}
	if (as.MinSet && len(rawArray) < as.MinSize) || (as.MaxSet && len(rawArray) > as.MaxSize) {
		return nil, errors.New("array length out of range")
	}
	for i := range rawArray {
		v, err := as.TypeSchema.Refine(rawArray[i])
		if err != nil {
			return nil, fmt.Errorf("member of integer array invalid:%v", err)
		}
		rtArray = append(rtArray, v)
	}
	return rtArray, nil
}

func (as *ArraySpec) DefaultValue() (interface{}, bool) {
	return nil, false //array dont support default value
}

type IntSpec struct {
	Min         int64   `json:"min" bson:"min"`
	Max         int64   `json:"max" bson:"max"`
	Default     int64   `json:"default" bson:"default"`
	Validate    []int64 `json:"validate" bson:"validate"`
	ValidateSet bool    `json:"-" bson:"validate_placed"`
	MinSet      bool    `json:"-" bson:"min_placed"`
	MaxSet      bool    `json:"-" bson:"max_placed"`
	DefaultSet  bool    `json:"-" bson:"default_placed"`
}

func (is *IntSpec) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	if is.MinSet {
		out["min"] = is.Min
	}
	if is.MaxSet {
		out["max"] = is.Max
	}
	if is.DefaultSet {
		out["default"] = is.Default
	}
	if is.ValidateSet {
		out["validate"] = is.Validate
	}
	return bson.Marshal(out)
}

func (is *IntSpec) UnmarshalBSON(b []byte) error {
	var raw struct {
		Min      bson.RawValue `bson:"min"`
		Max      bson.RawValue `bson:"max"`
		Default  bson.RawValue `bson:"default"`
		Validate bson.RawValue `bson:"validate"`
	}
	if err := bson.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.Min.Value != nil {
		if i, ok := raw.Min.Int64OK(); ok {
			is.Min, is.MinSet = i, true
		} else if i, ok := raw.Min.Int32OK(); ok {
			is.Min, is.MinSet = int64(i), true
		} else {
			return fmt.Errorf("integer spec min invalid: %s", raw.Min.String())

		}
	}
	if raw.Max.Value != nil {
		if i, ok := raw.Max.Int64OK(); ok {
			is.Max, is.MaxSet = i, true
		} else if i, ok := raw.Max.Int32OK(); ok {
			is.Max, is.MaxSet = int64(i), true
		} else {
			return fmt.Errorf("integer spec max invalid: %s", raw.Max.String())
		}
	}
	if raw.Default.Value != nil {
		var d int64
		if i, ok := raw.Default.Int64OK(); ok {
			d = i
		} else if i, ok := raw.Default.Int32OK(); ok {
			d = int64(i)
		} else {
			return fmt.Errorf("integer spec Default invalid: %s", raw.Default.String())
		}
		if (is.MinSet && d < is.Min) || (is.MaxSet && d > is.Max) {
			return fmt.Errorf("integer spec default value out of valid range:%d", d)
		}
		is.Default, is.DefaultSet = d, true
	}
	if raw.Validate.Value != nil {
		if a, ok := raw.Validate.ArrayOK(); ok {
			v := primitive.M{}
			err := bson.Unmarshal(a, &v)
			if err != nil {
				return fmt.Errorf("integer spec validate invalid %s", raw.Validate)
			}
			va := make([]int64, 0)
			for _, i := range v {
				ii, ok := i.(int64)
				if !ok {
					return fmt.Errorf("integer spec validate element invalid %s", i)
				}
				va = append(va, ii)
			}
			is.Validate, is.ValidateSet = va, true
		} else {
			return fmt.Errorf("integer spec validate not array %s", raw.Validate)
		}

	}
	return nil
}

func (is *IntSpec) UnmarshalJSON(b []byte) error {
	var raw struct {
		Min      json.RawMessage `json:"min"`
		Max      json.RawMessage `json:"max"`
		Default  json.RawMessage `json:"default"`
		Validate json.RawMessage `json:"validate"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.Min != nil {
		if i, err := strconv.ParseInt(string(raw.Min), 10, 64); err == nil {
			is.Min, is.MinSet = i, true
		} else {
			return fmt.Errorf("integer spec min invalid: %s", raw.Min)

		}
	}
	if raw.Max != nil {
		if i, err := strconv.ParseInt(string(raw.Max), 10, 64); err == nil {
			is.Max, is.MaxSet = i, true
		} else {
			return fmt.Errorf("integer spec max invalid: %s", raw.Max)
		}
	}
	if raw.Default != nil {
		var d int64
		if i, err := strconv.ParseInt(string(raw.Default), 10, 64); err == nil {
			d = i
		} else {
			return fmt.Errorf("integer spec Default invalid: %s", raw.Default)
		}
		if (is.MinSet && d < is.Min) || (is.MaxSet && d > is.Max) {
			return fmt.Errorf("integer spec default value out of valid range:%d", d)
		}
		is.Default, is.DefaultSet = d, true
	}
	if raw.Validate != nil {
		v := make([]int64, 0)
		err := json.Unmarshal(raw.Validate, &v)
		if err != nil {
			return fmt.Errorf("integer spec validate invalid %s", raw.Validate)
		}
		is.Validate, is.ValidateSet = v, true
	}
	return nil
}

func (is *IntSpec) Refine(v []byte) (interface{}, error) {
	i, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		return nil, errors.New("invalid integer")
	}
	if (is.MinSet && i < is.Min) || (is.MaxSet && i > is.Max) {
		return nil, errors.New("integer out of range")
	}
	if is.ValidateSet {
		for _, validate := range is.Validate {
			if validate == i {
				return i, nil
			}
		}
		return nil, errors.New("integer out of validate")
	} else {
		return i, nil
	}
}

func (is *IntSpec) DefaultValue() (interface{}, bool) {
	if is.DefaultSet {
		return is.Default, true
	}
	return nil, false
}

type TextSpec struct {
	MaxSize     int64    `json:"max_length" bson:"max_length"`
	MinSize     int64    `json:"min_length" bson:"min_length"`
	Default     string   `json:"default" bson:"default"`
	Validate    []string `json:"validate" bson:"validate"`
	ValidateSet bool     `json:"-" bson:"validate_placed"`
	MaxSet      bool     `json:"-" bson:"max_placed"`
	MinSet      bool     `json:"-" bson:"min_placed"`
	DefaultSet  bool     `json:"-" bson:"default_placed"`
}

func (ts *TextSpec) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	if ts.MinSet {
		out["min_size"] = ts.MinSize
	}
	if ts.MaxSet {
		out["max_size"] = ts.MaxSize
	}
	if ts.DefaultSet {
		out["default"] = ts.Default
	}
	if ts.ValidateSet {
		out["validate"] = ts.Validate
	}
	return bson.Marshal(out)
}

func (ts *TextSpec) UnmarshalBSON(b []byte) error {
	var raw struct {
		MinSize  bson.RawValue `bson:"min_size"`
		MaxSize  bson.RawValue `bson:"max_size"`
		Default  bson.RawValue `bson:"default"`
		Validate bson.RawValue `bson:"validate"`
	}
	if err := bson.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.MinSize.Value != nil {
		if i, ok := raw.MinSize.Int64OK(); ok {
			ts.MinSize, ts.MinSet = int64(i), true
		} else if i, ok := raw.MinSize.Int32OK(); ok {
			ts.MinSize, ts.MinSet = int64(i), true
		} else {
			return fmt.Errorf("text spec MinSize invalid: %s", raw.MinSize.String())
		}

	}
	if raw.MaxSize.Value != nil {
		if i, ok := raw.MaxSize.Int64OK(); ok {
			ts.MaxSize, ts.MaxSet = int64(i), true
		} else if i, ok := raw.MaxSize.Int32OK(); ok {
			ts.MaxSize, ts.MaxSet = int64(i), true
		} else {
			return fmt.Errorf("text spec MaxSize invalid: %s", raw.MinSize.String())
		}
	}
	if raw.Default.Value != nil {
		defaultLen := int64(utf8.RuneCount(raw.Default.Value))
		if ts.MinSet && defaultLen < ts.MinSize || ts.MaxSet && defaultLen > ts.MaxSize {
			return errors.New("text spec default length out of valid range")
		}
		ts.Default, ts.DefaultSet = raw.Default.String(), true
	}
	if raw.Validate.Value != nil {
		v := primitive.M{}
		err := bson.Unmarshal(raw.Validate.Value, &v)
		if err != nil {
			return fmt.Errorf("text spec validate invalid: %s", raw.Validate.String())
		}
		ss := make([]string, 0)
		for _, vali := range v {
			if s, ok := vali.(string); ok {
				ss = append(ss, s)
			} else {
				return fmt.Errorf("text spec validate not array: %s", raw.Validate.String())
			}
		}
		ts.Validate, ts.ValidateSet = ss, true
	}
	return nil
}

func (ts *TextSpec) UnmarshalJSON(b []byte) error {
	var raw struct {
		MinSize  json.RawMessage `json:"min_size"`
		MaxSize  json.RawMessage `json:"max_size"`
		Default  json.RawMessage `json:"default"`
		Validate json.RawMessage `json:"validate"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.MinSize != nil {
		if i, err := strconv.ParseInt(string(raw.MinSize), 10, 64); err == nil {
			ts.MinSize, ts.MinSet = int64(i), true
		} else {
			return fmt.Errorf("text spec MinSize invalid: %s", raw.MinSize)
		}
	}
	if raw.MaxSize != nil {
		if i, err := strconv.ParseInt(string(raw.MaxSize), 10, 64); err == nil {
			ts.MaxSize, ts.MaxSet = int64(i), true
		} else {
			return fmt.Errorf("text spec MaxSize invalid: %s", raw.MinSize)
		}
	}
	if raw.Default != nil {
		defaultLen := int64(utf8.RuneCount(raw.Default))
		if ts.MinSet && defaultLen < ts.MinSize || ts.MaxSet && defaultLen > ts.MaxSize {
			return errors.New("text spec default length out of valid range")
		}
		ts.Default, ts.DefaultSet = string(raw.Default), true
	}
	if raw.Validate != nil {
		v := make([]string, 0)
		err := json.Unmarshal(raw.Validate, &v)
		if err != nil {
			return fmt.Errorf("text spec validate invalid: %s", raw.Validate)
		}
		ts.Validate, ts.ValidateSet = v, true
	}
	return nil
}

func (ts *TextSpec) Refine(v []byte) (interface{}, error) {
	if len(v) < 2 || v[0] != '"' || v[len(v)-1] != '"' {
		return nil, errors.New("invalid text")
	}
	if (ts.MinSet && int64(utf8.RuneCount(v[1:len(v)-1])) < ts.MinSize) || (ts.MaxSet && int64(utf8.RuneCount(v[1:len(v)-1])) > ts.MaxSize) {
		return nil, errors.New("text out of range")
	}
	if ts.ValidateSet {
		for _, validate := range ts.Validate {
			if validate == string(v[1:len(v)-1]) {
				return string(v[1 : len(v)-1]), nil
			}
		}
		return nil, errors.New("text out of validate")
	} else {
		return string(v[1 : len(v)-1]), nil
	}
}

func (ts *TextSpec) DefaultValue() (interface{}, bool) {
	if ts.DefaultSet {
		return ts.Default, true
	}
	return nil, false
}

type FloatSpec struct {
	Min         float64   `json:"min" bson:"min"`
	Max         float64   `json:"max" bson:"max"`
	Default     float64   `json:"default" bson:"default"`
	Validate    []float64 `json:"validate" bson:"validate"`
	MinSet      bool      `json:"-" bson:"min_placed"`
	MaxSet      bool      `json:"-" bson:"max_placed"`
	DefaultSet  bool      `json:"-" bson:"default_placed"`
	ValidateSet bool      `json:"-" bson:"validate_placed"`
}

func (fs *FloatSpec) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	if fs.MinSet {
		out["min"] = fs.Min
	}
	if fs.MaxSet {
		out["max"] = fs.Max
	}
	if fs.DefaultSet {
		out["default"] = fs.Default
	}
	if fs.ValidateSet {
		out["validate"] = fs.Validate
	}
	return bson.Marshal(out)
}

func (fs *FloatSpec) UnmarshalBSON(b []byte) error {
	var raw struct {
		Min      bson.RawValue `bson:"min"`
		Max      bson.RawValue `bson:"max"`
		Default  bson.RawValue `bson:"default"`
		Validate bson.RawValue `bson:"validate"`
	}
	if err := bson.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.Min.Value != nil {
		if f, ok := raw.Min.DoubleOK(); ok {
			fs.Min, fs.MinSet = f, true
		} else {
			return fmt.Errorf("float spec min invalid: %s", raw.Min.String())
		}
	}
	if raw.Max.Value != nil {
		if f, ok := raw.Max.DoubleOK(); ok {
			fs.Max, fs.MaxSet = f, true
		} else {
			return fmt.Errorf("float spec max invalid: %s", raw.Max.String())
		}
	}
	if raw.Default.Value != nil {
		d, ok := raw.Default.DoubleOK()
		if !ok {
			return fmt.Errorf("float spec default invalid: %s", raw.Default.String())
		}
		if (fs.MinSet && d < fs.Min) || (fs.MaxSet && d > fs.Max) {
			return fmt.Errorf("float spec default value out of valid range:%s", raw.Default.String())
		}
		fs.Default, fs.DefaultSet = d, true
	}
	if raw.Validate.Value != nil {
		v := primitive.M{}
		err := bson.Unmarshal(raw.Validate.Value, &v)
		if err != nil {
			return fmt.Errorf("float spec validate invalid: %s", raw.Validate.String())
		}
		vv := make([]float64, 0)
		for _, f := range v {
			if ff, ok := f.(float64); ok {
				vv = append(vv, ff)
			}
		}
		fs.Validate, fs.ValidateSet = vv, true
	}
	return nil
}

func (fs *FloatSpec) UnmarshalJSON(b []byte) error {
	var raw struct {
		Min      json.RawMessage `json:"min"`
		Max      json.RawMessage `json:"max"`
		Default  json.RawMessage `json:"default"`
		Validate json.RawMessage `json:"validate"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.Min != nil {
		if f, err := strconv.ParseFloat(string(raw.Min), 64); nil == err {
			fs.Min, fs.MinSet = f, true
		} else {
			return fmt.Errorf("float spec min invalid: %s", raw.Min)
		}
	}
	if raw.Max != nil {
		if f, err := strconv.ParseFloat(string(raw.Max), 64); nil == err {
			fs.Max, fs.MaxSet = f, true
		} else {
			return fmt.Errorf("float spec max invalid: %s", raw.Max)
		}
	}
	if raw.Default != nil {
		if d, err := strconv.ParseFloat(string(raw.Min), 64); nil == err {
			if (fs.MinSet && d < fs.Min) || (fs.MaxSet && d > fs.Max) {
				return fmt.Errorf("float spec default value out of valid range:%s", raw.Default)
			}
			fs.Default, fs.DefaultSet = d, true
		} else {
			return fmt.Errorf("float spec default invalid: %s", raw.Default)
		}
	}
	if raw.Validate != nil {
		v := make([]float64, 0)
		err := json.Unmarshal(raw.Validate, &v)
		if err != nil {
			return fmt.Errorf("float spec validate invalid: %s", raw.Validate)
		}
		fs.Validate, fs.ValidateSet = v, true
	}
	return nil
}

func (fs *FloatSpec) Refine(v []byte) (interface{}, error) {
	i, err := strconv.ParseFloat(string(v), 64)
	if err != nil {
		return nil, errors.New("invalid float")
	}
	if (fs.MinSet && i < fs.Min) || (fs.MaxSet && i > fs.Max) {
		return nil, errors.New("float out of range")
	}
	if fs.ValidateSet {
		for _, validate := range fs.Validate {
			if validate == i {
				return i, nil
			}
		}
		return i, errors.New("float out of validate")
	} else {
		return i, nil
	}
}

func (fs *FloatSpec) DefaultValue() (interface{}, bool) {
	if fs.DefaultSet {
		return fs.Default, true
	}
	return nil, false
}

type BoolSpec struct {
	Default    bool `json:"default" bson:"default"`
	DefaultSet bool `json:"-" bson:"default_placed"`
}

func (bs *BoolSpec) MarshalBSON() ([]byte, error) {
	var (
		out = make(map[string]interface{})
	)
	if bs.DefaultSet {
		out["default"] = bs.Default
	}
	return bson.Marshal(out)
}

func (bs *BoolSpec) UnmarshalBSON(b []byte) error {
	var raw struct {
		Default bson.RawValue `bson:"default"`
	}
	if err := bson.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.Default.Value != nil {
		d, ok := raw.Default.BooleanOK()
		if !ok {
			return fmt.Errorf("boolean spec default invalid: %s", raw.Default.String())
		}
		bs.Default, bs.DefaultSet = d, true
	}
	return nil
}

func (bs *BoolSpec) UnmarshalJSON(b []byte) error {
	var raw struct {
		Default json.RawMessage `json:"default"`
	}
	if err := bson.Unmarshal(b, &raw); err != nil {
		return err
	}
	if raw.Default != nil {
		if d, err := strconv.ParseBool(string(raw.Default)); nil == err {
			bs.Default, bs.DefaultSet = d, true
		} else {
			return fmt.Errorf("boolean spec default invalid: %s", raw.Default)
		}

	}
	return nil
}

func (bs *BoolSpec) Refine(v []byte) (interface{}, error) {
	b, err := strconv.ParseBool(string(v))
	if err != nil {
		return nil, errors.New("invalid boolean")
	}
	return b, nil
}

func (bs *BoolSpec) DefaultValue() (interface{}, bool) {
	if bs.DefaultSet {
		return bs.Default, true
	}
	return nil, false
}
