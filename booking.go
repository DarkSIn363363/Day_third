package main

func bookingStatus(
	startHour int,
	durationHours int,
	participants int,
	isAdmin bool,
) string {

	if startHour < 9 || startHour > 17 {
		return "Некорректное время бронирования"
	}
	if durationHours < 1 {
		return "Некорректная длительность бронирования"
	}
	if !isAdmin && durationHours > 3 {
		return "Бронирование не должно превышать 3 часов"
	}
	if isAdmin && durationHours > 5 {
		return "Бронирование не должно превышать 5 часов"
	}
	if (startHour + durationHours) > 18 {
		return "Бронирование не может заканчиваться после 18:00"
	}
	if !isAdmin && participants > 8 || !isAdmin && participants < 1 {
		return "Бронирование доступно от 1-8 человек"
	} 
	if isAdmin && participants > 20 || isAdmin && participants < 1 {
		return "Бронирование доступно от 1-20 человек"
	} 
	
	return "бронирование подтверждено"
}

