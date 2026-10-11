package main

import (
	"math"
	"testing"
)

func TestValidateOrder(t *testing.T) {
	tests := []struct {
		name         string
		price        float64
		quantity     int
		discountRate float64
		want         string
	}{
		{
			name:         "корректный заказ",
			price:        1000,
			quantity:     1,
			discountRate: 0,
			want:         "",
		},
		{
			name:         "нулевая цена",
			price:        0,
			quantity:     1,
			discountRate: 0,
			want:         "цена должна быть больше 0",
		},
		{
			name:         "отрицательная цена",
			price:        -10,
			quantity:     1,
			discountRate: 0,
			want:         "цена должна быть больше 0",
		},
		{
			name:         "нулевое количество",
			price:        1000,
			quantity:     0,
			discountRate: 0,
			want:         "количество должно быть больше 0",
		},
		{
			name:         "отрицательная скидкка",
			price:        1000,
			quantity:     1,
			discountRate: -0.01,
			want:         "скидка не может быть отрицательной",
		},
		{
			name:         "скидка выше максимума",
			price:        1000,
			quantity:     1,
			discountRate: 0.51,
			want:         "скидка не может быть больше 50%",
		},
		{
			name:         "граница допустимой скидки",
			price:        1000,
			quantity:     1,
			discountRate: 0.50,
			want:         "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := validateOrder(test.price, test.quantity, test.discountRate)

			if got != test.want {
				t.Errorf(
					"validateOrder(%.2f, %d, %.2f) = %s, want %s",
					test.price,
					test.quantity,
					test.discountRate,
					got,
					test.want,
				)
			}
		})
	}
}

func TestDeliveryCost(t *testing.T) {
	tests := []struct {
		name              string
		subtotal          float64
		isExpressDelivery bool
		want              float64
	}{
		{
			name:         "обычная, ниже порога",
			subtotal:        4999.99,
			isExpressDelivery: false,
			want:         299.0,
		},
		{
			name:         "обычная, ровно на пороге",
			subtotal:        5000,
			isExpressDelivery: false,
			want:         0,
		},
		{
			name:         "обычная, выше порога",
			subtotal:        7000,
			isExpressDelivery: false,
			want:         0,
		},
		{
			name:         "экспресс, ниже порога",
			subtotal:        1000,
			isExpressDelivery: true,
			want:         599.0,
		},
		{
			name:         "экспресс, выше порога",
			subtotal:        7000,
			isExpressDelivery: true,
			want:         599.0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := deliveryCost(test.subtotal, test.isExpressDelivery)

			if got != test.want {
				t.Errorf(
					"deliveryCost(%.2f, %t) = %.2f, want %.2f",
					test.subtotal,
					test.isExpressDelivery,
					got,
					test.want,
				)
			}
		})
	}
}

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name          string
		price         float64
		quantity      int
		discountRate  float64
		deliveryPrice float64
		want          float64
	}{
		{
			name:         "без скидки",
			price:        1000,
			quantity:     2,
			discountRate: 0,
			want:         2000,
		},
		{
			name:         "со скидкой 10 процентов",
			price:        1000,
			quantity:     2,
			discountRate: 0.10,
			want:         1800,
		},
		{
			name:         "один товар со скидкой 50 процентов",
			price:        1000,
			quantity:     1,
			discountRate: 0.50,
			want:         500,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculateTotal(test.price, test.quantity, test.discountRate)

			const epsilon = 0.000001

			if math.Abs(got-test.want) > epsilon {
				t.Errorf(
					"calculateTotal(%.2f, %d, %.2f) = %.2f, want %.2f",
					test.price,
					test.quantity,
					test.discountRate,
					got,
					test.want,
				)
			}
		})
	}
}
