package go_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // using defer to make sure it always executed

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Geralt Of Rivia" // mengirim data ke channel
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel // menyimpan data ke variable dari channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// receive channel as parameter and set the value
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Jokowi"
	// ibaratnya kita menunggu logic yang akan dijalankan oleh function ini, lalu mengirim result nya ke channel
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel) // send channel as parameter

	data := <-channel // receive data from channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Joko Widodo"
	// data := <-channel // will error because can't receive / menerima
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	// channel <- "Jokowi" // will error can't send
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4) // buffer ada 4, bisa menampung tiga data, jika mengirim lebih dari 5 maka akan error deadlock
	defer close(channel)

	channel <- "Arthur"
	channel <- "Shelby"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	go func() { // menggunakan goroutine
		channel <- "Thomas"
		channel <- "Shelby"
	}()

	go func() { // menggunakan go routine
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func loopNumber(i int) string {
	return "Perulangan ke " + strconv.Itoa(i)
}

func TestRangeChannel(t *testing.T) {
	channels := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channels <- loopNumber(i)
		}
		close(channels)
	}()

	for channel := range channels {
		fmt.Println("Menerima data", channel)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		// select which channel we want to receive
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		// select which channel we want to receive
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default: // default value jika belum ada channel
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
