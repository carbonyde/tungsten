package tungsten

import (
	"strconv"
	"time"
)

func NumToString(num int) string {
	return strconv.Itoa(num)
}

func FormatDate(date string) string {
	t, err := time.Parse(time.DateOnly, date)

	if err != nil {
		panic(err)
	}

	return t.Format("02 Jan 2006")
}
