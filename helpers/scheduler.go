package helpers

import (
	"regexp"
	"strconv"

	"github.com/fanchann/dbackupdul/model"
)

func Scheduler(param string) string {
	switch param {
	case "hourly":
		return model.EveryHour
	case "daily":
		return model.EveryDay
	case "midnight":
		return model.EveryMidnight
	case "weekly":
		return model.EveryWeek
	default:
		if isValidTime(param) {
			return model.CustomTime + param
		}
		return model.EveryDay
	}
}

func isValidTime(input string) bool {
	//check the pattern
	re := regexp.MustCompile(`^(\d+h)?(\d+m)?(\d+s)?$`)

	match := re.FindStringSubmatch(input)
	if len(match) == 0 {
		return false
	}

	// if match with regex pattern, check again every value is valid
	for _, m := range match[1:] {
		if m == "" {
			continue
		}
		_, err := strconv.Atoi(m[:len(m)-1])
		if err != nil {
			return false
		}
	}

	return true
}
