package _go

import (
	"sort"
	"unsafe"
)

// 排序
type sorter struct{}

var (
	Sorter sorter
)

func (*sorter) Float64AZ(values []float64) { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) Float32AZ(values []float32) { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) Int64AZ(values []int64)     { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) Int32AZ(values []int32)     { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) Int16AZ(values []int16)     { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) Int8AZ(values []int8)       { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) IntAZ(values []int)         { sort.Ints(values) }
func (*sorter) Float64ZA(values []float64) { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) Float32ZA(values []float32) { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) Int64ZA(values []int64)     { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) Int32ZA(values []int32)     { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) Int16ZA(values []int16)     { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) Int8ZA(values []int)        { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) IntZA(values []int)         { sort.Sort(sort.Reverse(sort.IntSlice(values))) }
func (*sorter) StringAZ(values []string)   { sort.Strings(values) }
func (*sorter) StringZA(values []string)   { sort.Sort(sort.Reverse(sort.StringSlice(values))) }
