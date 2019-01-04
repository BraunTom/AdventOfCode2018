package FourthDay

import (
	"regexp"
	"strconv"
)

type entryType int

const (
	shift entryType = iota
	asleep
	wakeup
)

type entry struct {
	month       int
	day         int
	hour        int
	minute      int
	description string
}

type entries []entry

func (e entries) Len() int {
	return len(e)
}

func (e entries) Less(i, j int) bool {
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

func (e entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e entry) Type() entryType {
	switch e.description[0] {
	case 'G':
		return shift
	case 'f':
		return asleep
	case 'w':
		return wakeup
	}

	panic("should not happen")
}

func (e entry) ID() int {
	idRegex := regexp.MustCompile("#\\d* ")
	id := idRegex.FindString(e.description)

	result, _ := strconv.Atoi(id[1 : len(id)-1])
	return result
}

func (e entry) TimeDifference(secondEntry entry) int {
	return 0
}
