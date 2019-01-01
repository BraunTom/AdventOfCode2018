package FourthDay

import (
	"regexp"
	"strconv"
	"time"
)

type Entry struct {
	month       int
	day         int
	hour        int
	minute      int
	description string

	id int
}

type Entries []Entry

func (e Entries) Len() int {
	return len(e)
}

func (e Entries) Less(i, j int) bool {
	if e[i].month < e[j].month {
		return true
	}
	if e[i].month == e[j].month {
		if e[i].day < e[j].day {
			return true
		}
		if e[i].day == e[j].day {
			if e[i].hour < e[j].hour {
				return true
			}
			if e[i].hour == e[j].hour {
				if e[i].minute < e[j].minute {
					return true
				}
			}
		}
	}

	return false
}

func (e Entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Entry) Type() string {
	switch e.description[0] {
	case 'G':
		return "shift"
	case 'f':
		return "asleep"
	case 'w':
		return "wakeup"
	}

	return "unknown"
}

func (e Entry) unixTime() {
	time.Date(2000, time.November, e.day, e.hour, e.minute, 0, 0, time.UTC).Unix()
}

func (e Entry) ID() int {
	if e.id != 0 {
		return e.id
	}

	idRegex := regexp.MustCompile("#\\d* ")
	id := idRegex.FindString(e.description)

	result, _ := strconv.Atoi(id[1 : len(id)-1])
	e.id = result
	return result
}

func (e Entry) TimeDifference(secondEntry Entry) int {
	return 0
}
