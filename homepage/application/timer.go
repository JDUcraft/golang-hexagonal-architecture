package application

import "main/homepage/domain"

func ReadTime(timeRepository domain.TimeRepository) string {
	return "The time is: " + timeRepository.GetTime()
}
