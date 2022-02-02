package application

import "main/homepage/domain"

func ReadTime(clock domain.TimeRepository) string {
	return "The time is: " + clock.GetTime()
}
