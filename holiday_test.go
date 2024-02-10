package jpholiday

import (
	"testing"
	"time"
)

func TestIsJapanHoliday(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	// 1954-12-31 is out of service
	if IsJapanHoliday(time.Date(1954, 12, 31, 0, 0, 0, 0, jst)) {
		t.Errorf("1954-12-31 is not a holiday")
	}

	// 1955-01-01 is a holiday
	if !IsJapanHoliday(time.Date(1955, 1, 1, 0, 0, 0, 0, jst)) {
		t.Errorf("1955-01-01 is a holiday")
	}

	// 1972-5-5 is a holiday
	if !IsJapanHoliday(time.Date(1972, 5, 5, 0, 0, 0, 0, jst)) {
		t.Errorf("1972-5-5 is a holiday")
	}

	// 1991-7-4 is not a holiday
	if IsJapanHoliday(time.Date(1991, 7, 4, 0, 0, 0, 0, jst)) {
		t.Errorf("1991-7-4 is not a holiday")
	}

	// 2025-11-24 is a holiday
	if !IsJapanHoliday(time.Date(2025, 11, 24, 0, 0, 0, 0, jst)) {
		t.Errorf("2025-11-24 is a holiday")
	}

	// 2025-11-25 is not a holiday
	if IsJapanHoliday(time.Date(2025, 11, 25, 0, 0, 0, 0, jst)) {
		t.Errorf("2025-11-25 is not a holiday")
	}

	utc, _ := time.LoadLocation("UTC")

	// 1954-12-31 14:59:59:999 (UTC) is out of service
	if IsJapanHoliday(time.Date(1954, 12, 31, 15, 0, 0, 0, utc).Add(-1)) {
		t.Errorf("1954-12-31 14:59:59 (UTC) is not a holiday")
	}

	// 1954-12-31 15:00:00:000 (UTC) is a holiday
	if !IsJapanHoliday(time.Date(1954, 12, 31, 15, 0, 0, 0, utc)) {
		t.Errorf("1954-12-31 15:00:00 (UTC) is a holiday")
	}

	// 1955-01-01 14:59:59:999 (UTC) is a holiday
	if !IsJapanHoliday(time.Date(1955, 1, 1, 15, 0, 0, 0, utc).Add(-1)) {
		t.Errorf("1955-01-01 14:59:59 (UTC) is a holiday")
	}

	// 1955-01-01 15:00:00:000 (UTC) is not a holiday
	if IsJapanHoliday(time.Date(1955, 1, 1, 15, 0, 0, 0, utc)) {
		t.Errorf("1955-01-01 15:00:00 (UTC) is not a holiday")
	}
}

func Test_GetJapanHolidayName(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	// 1954-12-31 is out of service
	_, ok := GetJapanHolidayName(time.Date(1954, 12, 31, 0, 0, 0, 0, jst))
	if ok {
		t.Errorf("1954-12-31 is not a holiday")
	}

	// 1955-01-01 is a holiday
	name, ok := GetJapanHolidayName(time.Date(1955, 1, 1, 0, 0, 0, 0, jst))
	if !ok {
		t.Errorf("1955-01-01 is a holiday")
	}
	if name != "元日" {
		t.Errorf("1955-01-01 is 元日")
	}

	// 1972-5-5 is a holiday
	name, ok = GetJapanHolidayName(time.Date(1972, 5, 5, 0, 0, 0, 0, jst))
	if !ok {
		t.Errorf("1972-5-5 is a holiday")
	}
	if name != "こどもの日" {
		t.Errorf("1972-5-5 is こどもの日")
	}

	// 1991-7-4 is not a holiday
	_, ok = GetJapanHolidayName(time.Date(1991, 7, 4, 0, 0, 0, 0, jst))
	if ok {
		t.Errorf("1991-7-4 is not a holiday")
	}

	// 2025-11-24 is a holiday
	name, ok = GetJapanHolidayName(time.Date(2025, 11, 24, 0, 0, 0, 0, jst))
	if !ok {
		t.Errorf("2025-11-24 is a holiday")
	}
	if name != "休日" {
		t.Errorf("2025-11-24 is 休日")
	}

	// 2025-11-25 is not a holiday
	_, ok = GetJapanHolidayName(time.Date(2025, 11, 25, 0, 0, 0, 0, jst))
	if ok {
		t.Errorf("2025-11-25 is not a holiday")
	}

	utc, _ := time.LoadLocation("UTC")

	// 1954-12-31 14:59:59:999 (UTC) is out of service
	_, ok = GetJapanHolidayName(time.Date(1954, 12, 31, 15, 0, 0, 0, utc).Add(-1))
	if ok {
		t.Errorf("1954-12-31 14:59:59 (UTC) is not a holiday")
	}

	// 1954-12-31 15:00:00:000 (UTC) is a holiday
	name, ok = GetJapanHolidayName(time.Date(1954, 12, 31, 15, 0, 0, 0, utc))
	if !ok {
		t.Errorf("1954-12-31 15:00:00 (UTC) is a holiday")
	}

	// 1955-01-01 14:59:59:999 (UTC) is a holiday
	name, ok = GetJapanHolidayName(time.Date(1955, 1, 1, 15, 0, 0, 0, utc).Add(-1))
	if !ok {
		t.Errorf("1955-01-01 14:59:59 (UTC) is a holiday")
	}

	// 1955-01-01 15:00:00:000 (UTC) is not a holiday
	_, ok = GetJapanHolidayName(time.Date(1955, 1, 1, 15, 0, 0, 0, utc))
	if ok {
		t.Errorf("1955-01-01 15:00:00 (UTC) is not a holiday")
	}
}
