package main

import (
	"fmt"
	"main/pkg/cdeklib"
)

func main() {
	account := "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	securePassword := "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	apiURL := "https://api.edu.cdek.ru/"

	// Set up sending and delivery locations
	sendingAddress := "Россия, г. Москва, Cлавянский бульвар д.1"
	deliveryAddress := "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43"

	// Sending address
	fromLocation := cdeklib.LocationCalc{
		Address: sendingAddress,
	}

	// Delivery address
	toLocation := cdeklib.LocationCalc{
		Address: deliveryAddress,
	}

	// Parcel size
	size := cdeklib.Size{
		Weight: 1000,
		Length: 20,
		Width:  20,
		Height: 20,
	}

	var client = cdeklib.NewClient(true, apiURL, account, securePassword)

	tariffs, err := client.Calculate(fromLocation, toLocation, size)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(tariffs)

	//CreateOrder
	recipient := cdeklib.Recipient{
		Name: "Name",
		Phones: []cdeklib.Phone{
			{
				Number: "+79991112233",
			},
		},
	}
	packages := []cdeklib.Package{
		{
			Length: 15,
			Width:  25,
			Height: 30,
			Weight: 1000,
			Number: "TestNumber",
			Items: []cdeklib.Item{
				{
					Name:    "TestItem",
					WareKey: "TestWareKey",
					Payment: cdeklib.Money{Value: 0},
					Value:   0,
					Cost:    1000,
					Weight:  1000,
					Amount:  1,
				},
			},
		},
	}

	//time.Sleep(12 * time.Minute)

	orderID, err := client.CreateOrder(sendingAddress, deliveryAddress, recipient, packages, 233)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orderID)

	//CheckOrder
	status, err := client.GetStatus(orderID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(status)
}
