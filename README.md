[![Go Test](https://github.com/GarupanOjisan/japan-holiday/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/GarupanOjisan/japan-holiday/actions/workflows/go.yml)
[![Update Holidays](https://github.com/GarupanOjisan/japan-holiday/actions/workflows/update_holidays.yml/badge.svg?branch=main)](https://github.com/GarupanOjisan/japan-holiday/actions/workflows/update_holidays.yml)

# japan-holiday

このパッケージは、日本の祝日に関する機能を提供します。

## インストール

```bash
go get github.com/garupanojisan/japan-holiday
```

## 制限

このパッケージは、1955年から2024年までの祝日に対応しています。

## 使い方

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/GarupanOjisan/japan-holiday"
)

func main() {
	// 2023年1月1日は祝日かどうか
	jst, _ := time.LoadLocation("Asia/Tokyo")
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, jst)
	isHoliday := japan_holiday.IsJapanHoliday(date)
	fmt.Println("2023年1月1日は祝日ですか？", isHoliday)

	// 2023年1月1日の祝日名
	name, ok := japan_holiday.GetJapanHolidayName(date)
	if ok {
		fmt.Println("2023年1月1日の祝日名は", name, "です")
	} else {
		fmt.Println("2023年1月1日は祝日ではありません")
	}
}
```

## 参考情報

このパッケージの祝日情報は、内閣府のウェブサイトに掲載されている祝日のリストを元にしています。

https://www8.cao.go.jp/chosei/shukujitsu/gaiyou.html
