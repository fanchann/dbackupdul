package helpers

import "time"

func DateFormatter() string {
	return time.Now().Format("20060102150405") // year,month,date,hour,second
}
