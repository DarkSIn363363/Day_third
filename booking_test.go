package main

import (
	"testing"
)

func TestBookingStatus(t *testing.T) {
	tests := []struct {
		name          string
		startHour     int
	    durationHours int
	    participants  int
	    isAdmin       bool
		want          string
	}{
		{
			name:          "корректное бронирование",
			startHour:     9,
	        durationHours: 3,
	        participants:  3,
	        isAdmin:       false,
		    want:          "бронирование подтверждено",
		},
		{
			name:          "корректное бронирование администратора",
			startHour:     10,
	        durationHours: 4,
	        participants:  15,
	        isAdmin:       true,
		    want:          "бронирование подтверждено",
		},
		{
			name:          "начало в 8:00",
			startHour:     8,
	        durationHours: 2,
	        participants:  5,
	        isAdmin:       false,
		    want:          "Некорректное время бронирования",
		},
		{
			name:          "начало в 18:00",
			startHour:     18,
	        durationHours: 2,
	        participants:  5,
	        isAdmin:       false,
		    want:          "Некорректное время бронирования",
		},
		{
			name:          "длительность 0",
			startHour:     9,
	        durationHours: 0,
	        participants:  2,
	        isAdmin:       false,
		    want:          "Некорректная длительность бронирования",
		},
		{
			name:          "обычный пользователь бронирует на 4 часа",
			startHour:     9,
	        durationHours: 4,
	        participants:  2,
	        isAdmin:       false,
		    want:          "Бронирование не должно превышать 3 часов",
		},
		{
			name:          "администратор бронирует на 6 часов",
			startHour:     9,
	        durationHours: 6,
	        participants:  2,
	        isAdmin:       true,
		    want:          "Бронирование не должно превышать 5 часов",
		},
		{
			name:          "бронирование заканчивается позже 18:00",
			startHour:     17,
	        durationHours: 2,
	        participants:  2,
	        isAdmin:       false,
		    want:          "Бронирование не может заканчиваться после 18:00",
		},
		{
			name:          "0 участников",
			startHour:     9,
	        durationHours: 2,
	        participants:  0,
	        isAdmin:       false,
		    want:          "Бронирование доступно от 1-8 человек",
		},
		{
			name:          "9 участников у обычного пользователя",
			startHour:     9,
	        durationHours: 2,
	        participants:  9,
	        isAdmin:       false,
		    want:          "Бронирование доступно от 1-8 человек",
		},
		{
			name:          "21 участник у администратора",
			startHour:     9,
	        durationHours: 2,
	        participants:  21,
	        isAdmin:       true,
		    want:          "Бронирование доступно от 1-20 человек",
		},
		{
			name:          "несколько ошибок одновременно 1",
			startHour:     8,
	        durationHours: 15,
	        participants:  21,
	        isAdmin:       true,
		    want:          "Некорректное время бронирования",
		},
		{
			name:          "несколько ошибок одновременно 2",
			startHour:     17,
	        durationHours: 4,
	        participants:  0,
	        isAdmin:       false,
		    want:          "Бронирование не должно превышать 3 часов",
		},
		{
			name:          "несколько ошибок одновременно 3",
			startHour:     17,
	        durationHours: 2,
	        participants:  0,
	        isAdmin:       false,
		    want:          "Бронирование не может заканчиваться после 18:00",
		},
		}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := bookingStatus(test.startHour, test.durationHours, test.participants, test.isAdmin)

			if got != test.want {
				t.Errorf(
					"bookingStatus(%d, %d, %d, %t) = %s, want %s",
					test.startHour,
					test.durationHours,
					test.participants,
					test.isAdmin,
					got,
					test.want,
				)
			}
		})
	}
}