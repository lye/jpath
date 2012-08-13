package jpath

import (
	"io"
	"io/ioutil"
	"fmt"
	"math"
	"strconv"
	"encoding/json"
)

// JPath is a wrapper around an arbitrary interface{}
//
// JPath provides an error-checking-free interface to accessing the underlying
// value. The goal is not to completely eschew error checking, but to remove the
// validation from the encoding step. It allows you to make assumptions about
// the underlying structure without heaps of control structures, and simply returns
// zero-values when your assumptions are wrong.
//
// A zero-value JPath is valid, and will simply return the zero-values of whatever
// you ask of it.
type JPath struct {
	// I is the underlying value this JPath wraps. It could be anything.
	I interface{}
}

// ParseBytes parses the bytes as JSON and overwrites the underlying value with the result.
func (jp *JPath) ParseBytes(bytes []byte) error {
	jp.I = nil
	return json.Unmarshal(bytes, &jp.I)
}

// ParseString parses the passed string as JSON and overwrites the underlying value with the result.
func (jp *JPath) ParseString(str string) error {
	return jp.ParseBytes([]byte(str))
}

// ParseReader buffers the contents of the reader in-memory, then passes it to ParseBytes.
func (jp *JPath) ParseReader(r io.Reader) error {
	bytes, er := ioutil.ReadAll(r)
	if er != nil {
		return er
	}

	return jp.ParseBytes(bytes)
}

// Length returns the length of the underlying array, or 0 if the underlying object is not an array.
func (jp JPath) Length() int {
	if jp.I == nil {
		return 0
	}

	if ary, ok := jp.I.([]interface{}) ; ok {
		return len(ary)
	}

	return 0
}

// Index returns a new JPath wrapping the ith value of the underlying array. If the underlying value
// is not an array, it returns a zero-value JPath.
func (jp JPath) Index(i int) JPath {
	if jp.I == nil {
		return jp
	}

	ary, ok := jp.I.([]interface{})

	if !ok || len(ary) <= i {
		return JPath{nil}
	}

	return JPath{ary[i]}
}

// Field returns a new JPath wrapping the specified field if the underlying value is an object. Otherwise,
// it returns a zero-value JPath.
func (jp JPath) Field(s string) JPath {
	if jp.I == nil {
		return jp
	}

	obj, ok := jp.I.(map[string]interface{})

	if !ok {
		return JPath{nil}
	}

	return JPath{obj[s]}
}

// Fields returns a slice of strings containing the field names of the underlying object. If the underlying
// value is not an object, returns an empty slice.
func (jp JPath) Fields() []string {
	if jp.I == nil {
		return []string{}
	}

	ret := []string{}

	if obj, ok := jp.I.(map[string]interface{}) ; ok {
		for k, _ := range obj {
			ret = append(ret, k)
		}
	}

	return ret
}

// String does its best to convert the underlying value to a string. For zero-value JPath objects
// and JPath objects which wrap an object or array, String returns an empty string.
func (jp JPath) String() string {
	if jp.I == nil {
		return ""
	}

	if str, ok := jp.I.(string) ; ok {
		return str
	}

	if num, ok := jp.I.(int) ; ok {
		return fmt.Sprintf("%d", num)
	}

	if num, ok := jp.I.(int32) ; ok {
		return fmt.Sprintf("%d", num)
	}

	if num, ok := jp.I.(uint32) ; ok {
		return fmt.Sprintf("%d", num)
	}

	if num, ok := jp.I.(float64) ; ok {
		return fmt.Sprintf("%f", num)
	}

	return ""
}

// StringMap returns a map[string]string. If the underlying value is an object, the returned map
// consists of any fields that are strings. Otherwise, an empty map is returned. Any non-string
// values are coerced to strings.
func (jp JPath) StringMap() map[string]string {
	ret := map[string]string{}

	if jp.I == nil {
		return ret
	}

	for _, fieldName := range jp.Fields() {
		fieldValue := jp.Field(fieldName)
		ret[fieldName] = fieldValue.String()
	}

	return ret
}

// Float64 returns a float64 representation of the underlying value. Strings are coerced to
// numerics, if possible. Arrays, objects and non-number strings are decoded as 0.
func (jp JPath) Float64() float64 {
	if jp.I == nil {
		return 0
	}

	if num, ok := jp.I.(float64) ; ok {
		return num
	}

	if num, ok := jp.I.(int) ; ok {
		return float64(num)
	}

	if num, ok := jp.I.(int32) ; ok {
		return float64(num)
	}

	if num, ok := jp.I.(uint32) ; ok {
		return float64(num)
	}

	if str, ok := jp.I.(string) ; ok {
		num, _ := strconv.ParseFloat(str, 64)
		return num
	}

	return 0
}

// Float32 casts the return value of Float64.
func (jp JPath) Float32() float32 {
	return float32(jp.Float64())
}

// Int64 casts the return value of Float64 (all JSON numerics are encoded as doubles). 
// NaN float values are considered 0.
func (jp JPath) Int64() int64 {
	fval := jp.Float64()

	if math.IsNaN(fval) {
		return 0
	}

	return int64(fval)
}

// Int32 casts the return value of Int64
func (jp JPath) Int32() int32 {
	return int32(jp.Int64())
}

// Int16 casts the return value of Int64
func (jp JPath) Int16() int16 {
	return int16(jp.Int64())
}

// Int8 casts the return value of Int64
func (jp JPath) Int8() int8 {
	return int8(jp.Int64())
}

// Int casts the return value of Int64
func (jp JPath) Int() int {
	return int(jp.Int64())
}

// Uint64 casts the return value of Float64 (all JSON numerics are encoded as doubles).
func (jp JPath) Uint64() uint64 {
	fval := jp.Float64()

	if math.IsNaN(fval) {
		return 0
	}

	return uint64(fval)
}

// Uint32 casts the return value of Uint64
func (jp JPath) Uint32() uint32 {
	return uint32(jp.Uint64())
}

// Uint16 casts the return value of Uint64
func (jp JPath) Uint16() uint16 {
	return uint16(jp.Uint64())
}

// Uint8 casts the return value of Uint64
func (jp JPath) Uint8() uint8 {
	return uint8(jp.Uint64())
}

// Uint casts the return value of Uint64
func (jp JPath) Uint() uint {
	return uint(jp.Uint64())
}

// IsNull returns true if the underlying value was a JSON null, or if this JPath is a zero-value.
func (jp JPath) IsNull() bool {
	return jp.I == nil
}

// IsUndefined returns true if the underlying value was a JSON null, or if this JPath is a zero-value.
// A distinction is not made between JSON null and undefined -- the underlying json library doesn't
// allow us to make one.
func (jp JPath) IsUndefined() bool {
	return jp.I == nil
}
