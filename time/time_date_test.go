package array

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTimeMonth(t *testing.T) {
	// TODO 不要被IDE骗了， 这里输出是string
	fmt.Println(time.January)
}

func TestNow(t *testing.T) {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month()) // MONTH比较古怪 time.August
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
}

// 自定义输出格式日期
func TestNowFormat(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("06/1/2 15:4:5"))
}

// 获取当前unix时间戳，精确到秒级
func TestNowUnix(t *testing.T) {
	log.Println(time.Now().Unix())
}

// 将时间戳转换为指定格式的日期
func TestDateUnix(t *testing.T) {
	log.Println(time.Unix(1705490028, 0).Format("2006-01-02 15:04:05"))
}

// 将标准格式时间转换为时间戳
func TestStrtotimeUnix(t *testing.T) {
	timeParseString, err := time.Parse("2006-01-02 15:04:05", "2024-01-17 19:13:48")
	if err == nil {
		fmt.Println(timeParseString.Unix()) // 1705518828
	} else {
		fmt.Println(err.Error())
	}
}
