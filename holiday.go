package japan_holiday

import "time"

var jstLoc = time.FixedZone("Asia/Tokyo", 9*60*60)

func IsJapanHoliday(date time.Time) bool {
	t := date.In(jstLoc)
	key := t.Format("2006-1-2")
	_, ok := holidays[key]
	return ok
}

func GetJapanHolidayName(date time.Time) (name string, ok bool) {
	t := date.In(jstLoc)
	key := t.Format("2006-1-2")
	name, ok = holidays[key]
	return name, ok
}
