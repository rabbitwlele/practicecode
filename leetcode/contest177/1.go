package contest177

import "time"

func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02 15:04:05", date1+" 00:00:00")
	t2, _ := time.Parse("2006-01-02 15:04:05", date1+" 00:00:00")
	hours := t1.Sub(t2).Hours()
	return int(hours / 24)
}
