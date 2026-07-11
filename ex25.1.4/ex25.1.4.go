package main

import (
	"fmt"
	"sync"
	"time"
)

type car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() { //producer, consumer pattern
	tireCh := make(chan *car)
	paintCh := make(chan *car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")

}
func MakeBody(tireCh chan *car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			//make car body
			car := &car{}
			car.Body = "Sport car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}
func PaintCar(paintCh chan *car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "red"
		duration := time.Since(startTime)
		fmt.Printf("%.2f Complete Car : %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
