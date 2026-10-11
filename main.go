package main

import "fmt"

const (
	VATRate = 0.00
	StandardDeliveryPrice = 299.0
	ExpressDeliveryPrice = 599.0
	FreeDeliveryThreshold = 5000.0
)

func calculateTotal(price float64, quantity int, discountRate float64) float64 {
	subtotal := price * float64(quantity)
	discount := subtotal * discountRate
	priceAfterDiscount := subtotal - discount
	vat := priceAfterDiscount * VATRate

	return priceAfterDiscount + vat
}

func validateOrder(price float64, quantity int, discountRate float64) string {
     
	switch {
	case price <= 0:
		return "цена должна быть больше 0"
	case quantity <= 0:
		return "количество должно быть больше 0"
	case discountRate < 0 :
		return "скидка не может быть отрицательной"
	case discountRate > 0.5:
		return "скидка не может быть больше 50%"	
	default:
		return ""
	}
}

func deliveryCost(subtotal float64, isExpressDelivery bool) float64 {

     if subtotal >= FreeDeliveryThreshold && !isExpressDelivery{
		return 0.0
	} else if isExpressDelivery {
		return ExpressDeliveryPrice
	} else {
		return StandardDeliveryPrice
	} 
}

func main() {
	productName := "Колонки"
	price := 1399.0
	quantity := 6
	discountRate := 0.15
	isExpressDelivery := false
    
	validateOrder := validateOrder(price, quantity, discountRate)
	if validateOrder != "" {
       fmt.Printf("Ошибка оформления заказа: %s\n", validateOrder)
	   return
	}
	subtotal := price * float64(quantity)
	deliveryPrice := deliveryCost(subtotal, isExpressDelivery)
	total := calculateTotal(price, quantity, discountRate) + deliveryPrice
	typeDelivery := ""
	if isExpressDelivery {
		typeDelivery = "экспресс"
	} else {
		typeDelivery = "обычная"
	}

	fmt.Printf("Товар: %s\n", productName)
	fmt.Printf("Цена за единицу: %.2f ₽\n", price)
	fmt.Printf("Количество: %d\n", quantity)
	fmt.Printf("Скидка: %.2f%%\n", discountRate*100)
	fmt.Printf("Тип доставка: %s\n", typeDelivery)
	fmt.Printf("Стоимость доставки: %.2f ₽\n", deliveryPrice)
	fmt.Printf("Итого: %.2f ₽\n", total)
}