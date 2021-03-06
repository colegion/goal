// Package strconv implements conversions from string representation.
// All functions here require url.Values, a key, and an optional index;
// and return their appropriate types.
package strconv

import (
	"net/url"
	"strconv"
	"strings"
)

/*
	Below are functions for parsing booleans.
*/

// Bool returns the boolean value represented by the string.
// It treats "", "0", "f", "F", "false", and "FALSE" as false
// and any other value as true.
func Bool(vs url.Values, k string, is ...int) bool {
	switch strings.ToLower(get(vs, k, is)) {
	case "", "0", "f", "false":
		return false
	}
	return true
}

// Bools returns a slice of booleans represented by the strings.
func Bools(vs url.Values, k string, is ...int) (as []bool) {
	for i := range vs[k] {
		as = append(as, Bool(vs, k, i))
	}
	return
}

/*
	Below are functions for parsing strings.
*/

// String gets a string and returns it as is.
func String(vs url.Values, k string, is ...int) string {
	return get(vs, k, is)
}

// Strings returns a slice of strings.
func Strings(vs url.Values, k string, is ...int) []string {
	return vs[k]
}

/*
	Below are functions for parsing integers.
*/

// Int interprets a string s in the automatically determined base.
// The base is implied by the string's prefix: base 16 for "0x",
// base 8 for "0", and base 10 otherwise.
// If string cannot be parsed 0 will be returned anyway.
func Int(vs url.Values, k string, is ...int) int {
	n, _ := strconv.ParseInt(get(vs, k, is), 0, 0)
	return int(n)
}

// Ints returns a slice of integers represented by a slice of strings.
func Ints(vs url.Values, k string, is ...int) (as []int) {
	for i := range vs[k] {
		as = append(as, Int(vs, k, i))
	}
	return
}

// Int8 is similar to Int but is used for int8 type.
func Int8(vs url.Values, k string, is ...int) int8 {
	n, _ := strconv.ParseInt(get(vs, k, is), 0, 8)
	return int8(n)
}

// Int8s is similar to Ints but for []int8 type.
func Int8s(vs url.Values, k string, is ...int) (as []int8) {
	for i := range vs[k] {
		as = append(as, Int8(vs, k, i))
	}
	return
}

// Int16 is similar to Int but is used for int16 type.
func Int16(vs url.Values, k string, is ...int) int16 {
	n, _ := strconv.ParseInt(get(vs, k, is), 0, 16)
	return int16(n)
}

// Int16s is similar to Ints but for []int16 type.
func Int16s(vs url.Values, k string, is ...int) (as []int16) {
	for i := range vs[k] {
		as = append(as, Int16(vs, k, i))
	}
	return
}

// Int32 is similar to Int but is used for int32 type.
func Int32(vs url.Values, k string, is ...int) int32 {
	n, _ := strconv.ParseInt(get(vs, k, is), 0, 16)
	return int32(n)
}

// Int32s is similar to Ints but for []int32 type.
func Int32s(vs url.Values, k string, is ...int) (as []int32) {
	for i := range vs[k] {
		as = append(as, Int32(vs, k, i))
	}
	return
}

// Int64 is similar to Int but is used for int64 type.
func Int64(vs url.Values, k string, is ...int) int64 {
	n, _ := strconv.ParseInt(get(vs, k, is), 0, 16)
	return n
}

// Int64s is similar to Ints but for []int64 type.
func Int64s(vs url.Values, k string, is ...int) (as []int64) {
	for i := range vs[k] {
		as = append(as, Int64(vs, k, i))
	}
	return
}

/*
	Below are functions for parsing unsigned integers.
*/

// Uint is similar to Int but is used for uint type.
func Uint(vs url.Values, k string, is ...int) uint {
	n, _ := strconv.ParseUint(get(vs, k, is), 0, 0)
	return uint(n)
}

// Uints returns a slice of unsigned integers represented by a slice of strings.
func Uints(vs url.Values, k string, is ...int) (as []uint) {
	for i := range vs[k] {
		as = append(as, Uint(vs, k, i))
	}
	return
}

// Uint8 is similar to Int but is used for uint8 type.
func Uint8(vs url.Values, k string, is ...int) uint8 {
	n, _ := strconv.ParseUint(get(vs, k, is), 0, 8)
	return uint8(n)
}

// Uint8s is similar to Uints but for []uint8 type.
func Uint8s(vs url.Values, k string, is ...int) (as []uint8) {
	for i := range vs[k] {
		as = append(as, Uint8(vs, k, i))
	}
	return
}

// Uint16 is similar to Int but is used for uint16 type.
func Uint16(vs url.Values, k string, is ...int) uint16 {
	n, _ := strconv.ParseUint(get(vs, k, is), 0, 16)
	return uint16(n)
}

// Uint16s is similar to Uints but for []uint16 type.
func Uint16s(vs url.Values, k string, is ...int) (as []uint16) {
	for i := range vs[k] {
		as = append(as, Uint16(vs, k, i))
	}
	return
}

// Uint32 is similar to Int but is used for uint32 type.
func Uint32(vs url.Values, k string, is ...int) uint32 {
	n, _ := strconv.ParseUint(get(vs, k, is), 0, 32)
	return uint32(n)
}

// Uint32s is similar to Uints but for []uint32 type.
func Uint32s(vs url.Values, k string, is ...int) (as []uint32) {
	for i := range vs[k] {
		as = append(as, Uint32(vs, k, i))
	}
	return
}

// Uint64 is similar to Int but is used for uint64 type.
func Uint64(vs url.Values, k string, is ...int) uint64 {
	n, _ := strconv.ParseUint(get(vs, k, is), 0, 64)
	return n
}

// Uint64s is similar to Uints but for []uint64 type.
func Uint64s(vs url.Values, k string, is ...int) (as []uint64) {
	for i := range vs[k] {
		as = append(as, Uint64(vs, k, i))
	}
	return
}

/*
	Below are functions for parsing floating-point numbers.
*/

// Float32 converts the string s to a floating-point number.
// If it is impossible to parse the string, it still will return 0.0.
func Float32(vs url.Values, k string, is ...int) float32 {
	f, _ := strconv.ParseFloat(get(vs, k, is), 32)
	return float32(f)
}

// Float32s returns a slice of floating-point numbers represented by a slice of strings.
func Float32s(vs url.Values, k string, is ...int) (as []float32) {
	for i := range vs[k] {
		as = append(as, Float32(vs, k, i))
	}
	return
}

// Float64 is similar to Float32 but is used for float64 type.
func Float64(vs url.Values, k string, is ...int) float64 {
	f, _ := strconv.ParseFloat(get(vs, k, is), 64)
	return f
}

// Float64s is similar to Float32s but for []float64 type.
func Float64s(vs url.Values, k string, is ...int) (as []float64) {
	for i := range vs[k] {
		as = append(as, Float64(vs, k, i))
	}
	return
}

/*
	Below are various helper functions.
*/

// get receives url.Values, a key, and an index (as a slice, but only
// the first element is used). Then returns vs[k][is[0]].
// If the url.Values or slice of indexes is empty or key is not found
// empty string will be returned.
func get(vs url.Values, k string, is []int) string {
	// Get index from the first element of indexes slice
	// or use the last element's index as a default.
	i := len(vs[k])
	if len(is) > 0 {
		// If a specific index requested, use it.
		i = is[0]
	} else if i > 0 {
		// Otherwise, make sure the index is in range.
		// it should be len(values) - 1, but only if
		// len > 0. If it't not, use just 0.
		i--
	}

	// Make sure such key does exist and index is in range.
	// Return "" otherwise.
	ss := vs[k]
	if len(ss) <= i {
		return ""
	}

	// Return the final result.
	return vs[k][i]
}

// MeaningOfLife is the answer to the Ultimate Question of Life,
// the Universe, and Everything.
var MeaningOfLife = 42
