package main

import (
	"fmt"
	"time"
)

func main() {
	var year, month, day int
	now := time.Now()
	location := now.Location()
	currentYear, currentMonth, currentDay := now.Date()
	currentMonthConverted := int(currentMonth)

Year:
	fmt.Printf("请输入出生年份：")
	fmt.Scanf("%d\n", &year)
	if year > currentYear {
		fmt.Println("请输入合法年份！")
		goto Year
	}

Month:
	fmt.Printf("请输入出生月份：")
	fmt.Scanf("%d\n", &month)
	monthConverted := time.Month(month)
	if month < 1 || month > 12 || year == currentYear && month > currentMonthConverted {
		fmt.Println("请输入合法月份！")
		goto Month
	}

Day:
	fmt.Printf("请输入出生日期：")
	fmt.Scanf("%d\n", &day)
	lastDayOfMonth := time.Date(year, monthConverted, 1, 0, 0, 0, 0, location).AddDate(0, 1, -1).Day()
	if day < 1 || day > lastDayOfMonth || year == currentYear && month == currentMonthConverted && day > currentDay {
		fmt.Println("请输入合法日期！")
		goto Day
	}

	fmt.Printf("你的出生时间：%02d / %02d / %04d\n", day, month, year)
	birthday := time.Date(currentYear, monthConverted, day, 0, 0, 0, 0, location)

	dayDuration := time.Hour * 24
	if currentMonth == birthday.Month() && currentDay == birthday.Day() {
		fmt.Println("今天是你的生日，生日快乐！")
	} else if now.Before(birthday) {
		// The plus one is not always correct, of course
		daysLeft := (birthday.Sub(now) / dayDuration) + 1
		fmt.Printf("%d 天后是你的生日！", daysLeft)
	} else {
		daysAgo := now.Sub(birthday) / dayDuration
		fmt.Printf("%d 天前是你的生日！", daysAgo)
	}
}
