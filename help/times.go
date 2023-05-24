package help

import (
	"strconv"
	"time"
)

// GetCmbTransTime
//  @Description:   获取交易时间
//  @param d 日期
//  @param t 时间
//  @return time.Time
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 15:00:11
func GetCmbTransTime(d string, t string) (time.Time, error) {
	return time.ParseInLocation("20060102150405", d+t, time.Local)
}

// GetCmbOpTime
//  @Description:  获取交易
//  @param d
//  @return time.Time
//  @return error
//  @Author  ahKevinXy
//  @Date 2023-04-13 15:00:22
func GetCmbOpTime(d string) (time.Time, error) {
	return time.ParseInLocation("20060102", d, time.Local)
}

func GetReqidString() string {
	return time.Now().Format("20060102150405000") + strconv.Itoa(time.Now().Nanosecond())
}

func GetSigtimString() string {
	return time.Now().Format("20060102150405")
}
