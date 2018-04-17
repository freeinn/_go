package _go

import (
	"sort"
	"unsafe"
)

// 排序
type (
	Sorter interface {
		float64AZ()
		float32AZ()
		int64AZ()
		int32AZ()
		int16AZ()
		int8AZ()
		intAZ()
		float64ZA()
		float32ZA()
		int64ZA()
		int32ZA()
		int16ZA()
		int8ZA()
		intZA()
		stringAZ()
		stringZA()
	}
	sorter struct{}
)

func (*sorter) float64AZ(values []float64) { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) float32AZ(values []float32) { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) int64AZ(values []int64)     { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) int32AZ(values []int32)     { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) int16AZ(values []int16)     { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) int8AZ(values []int8)       { sort.Ints(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]) }
func (*sorter) intAZ(values []int)         { sort.Ints(values) }
func (*sorter) float64ZA(values []float64) { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) float32ZA(values []float32) { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) int64ZA(values []int64)     { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) int32ZA(values []int32)     { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) int16ZA(values []int16)     { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) int8ZA(values []int)        { sort.Sort(sort.Reverse(sort.IntSlice(((*[1 << 20]int)(unsafe.Pointer(&values[0])))[:len(values)]))) }
func (*sorter) intZA(values []int)         { sort.Sort(sort.Reverse(sort.IntSlice(values))) }
func (*sorter) stringAZ(values []string)   { sort.Strings(values) }
func (*sorter) stringZA(values []string)   { sort.Sort(sort.Reverse(sort.StringSlice(values))) }
