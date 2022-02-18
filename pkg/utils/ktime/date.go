package ktime

import (
	"time"
)

/**
 * 时间处理类
 */
const (
	GoTimeFormat             = "2006-01-02 15:04:05"
	GoDateFormat             = "2006-01-02"
	GoDateFormatOblique      = "2006/01/02"
	GoHourMinuteSecondFormat = "15:04:05"
)

/**
 * 把时间戳专程字符串类型的时间
 * 输入时间为秒
 * @param timestamp 1546830348
 * @return  2019-01-07 11:05:48
 */
func TimeStampToStringTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(GoTimeFormat)
}

func TimeStampToStringOnlyTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(GoHourMinuteSecondFormat)
}

func TimeStampToStringOnlyDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(GoDateFormat)
}

/**
 * 返回当前时间
 * @return 2019-01-07 11:05:48
 */
func GetCurrentStringTime() string {
	return time.Now().Format(GoTimeFormat)
}

/**
 * 把time.Time 类型转成string类型
 * @timeType timeType
 * @return 2019-01-07 11:05:48
 */
func TimeToStringTime(timeType time.Time) string {
	return timeType.Local().Format(GoTimeFormat)
}

/**
 * 把time.Time 类型转成string类型
 * @timeType timeType
 * @return 2019-01-07
 */
func TimeToStringDate(timeType time.Time) string {
	return timeType.Local().Format(GoDateFormat)
}

/**
 * 把time.Time 类型转成string类型
 * @timeType timeType
 * @return 2019/01/07
 */
func TimeToStringDateOblique(timeType time.Time) string {
	timeType = timeType.Local()
	return timeType.Format(GoDateFormatOblique)
}

/**
 * 把string类型时间转time.Time
 * @timeStr 2019-01-07
 * @return
 */
func StringToDateStamp(timeStr string) int64 {
	t, _ := time.ParseInLocation(GoDateFormat, timeStr, time.Local)
	return t.Unix()
}

/**
 * 把时间戳专程字符串类型的时间
 * 输入时间为秒
 * @param timeStr 2019-01-07 11:05:48
 * @return  2019-01-07 11:05:48
 */
func StringToTimeStamp(timeStr string) int64 {
	t, _ := time.ParseInLocation(GoTimeFormat, timeStr, time.Local)
	return t.Unix()
}

/**
 * 把日期字符串转成时间类型
 * 输入字符串
 * @param timeStr 2019-01-07
 * @return  2019-01-07 11:05:48
 */
func StringToDate(timeStr string) time.Time {
	t, _ := time.ParseInLocation(GoDateFormat, timeStr, time.Local)
	return t
}

/**
 * 把日期字符串转成整时间类型
 * 输入字符串
 * @param timeStr 2019-01-07
 * @return  2019-01-07 11:05:48
 */
func StringToTime(timeStr string) time.Time {
	t, _ := time.ParseInLocation(GoTimeFormat, timeStr, time.Local)
	return t
}

func GetCurentDateTimeStamp(time time.Time) int {
	date := time.Local().Format(GoDateFormat)
	return int(StringToDate(date).Unix())
}
