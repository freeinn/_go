package _go

import (
	"encoding/json"
	"fmt"
	"time"
	"math/rand"
)

// =============================================================================================================================

// 打印JSON
func PrintJSON(data interface{}) {
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(js))
}

// =============================================================================================================================

// 字典键值对调
func MapKVReversal(m map[string]string) map[string]string {
	temp := make(map[string]string)
	for k, v := range m {
		temp[v] = k
	}
	return temp
}

// =============================================================================================================================

// 切片反转
func SliceA2Z(slice []string) (temp []string) {
	for i := len(slice) - 1; i >= 0; i-- {
		temp = append(temp, slice[i])
	}
	return
}

// 切片去重
func SliceUnique(intSlice []string) []string {
	keys, list := make(map[string]bool), make([]string, 0)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// =============================================================================================================================

// 任务定时器 func runtimer( 开始时间, 周期时间, 任务方法 ) 开始时间格式 "00:00" (零点)
func RunTimer(startTime string, cycle time.Duration, task func()) {
	for {
		time.Sleep(time.Second) // 每秒监测一次
		if time.Now().Format("15:04") == startTime {
			for {
				task()
				time.Sleep(cycle)
			}
		}
	}
}

// 随机字符种子
func RandomChar(l int) string {
	//str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := "0123456789"
	strs := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, strs[r.Intn(len(strs))])
	}
	return string(result)
}

// 冒泡排序
func BubbleAsort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			// 控制排序顺序
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}
