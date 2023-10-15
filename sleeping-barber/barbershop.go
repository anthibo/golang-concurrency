package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarbers(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There's nothing todo, so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s waked %s up.", client, barber)
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				//  shop is closed, so send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}

		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s has finished cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarberDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")

	close(shop.ClientsChan)
	shop.Open = false

	for i := 1; i <= shop.NumberOfBarbers; i++ {
		// block until every single barber is done
		<-shop.BarberDoneChan
	}

	close(shop.BarberDoneChan)

	color.Green("===============================")
	color.Green("The barbershop is closed to the day, and everyone is sent home")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s arrives", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default:
			color.Red("The waiting room is full, so %s leaved", client)

		}
	} else {
		color.Red("The shop is already closed, so %s leaved", client)

	}
}
