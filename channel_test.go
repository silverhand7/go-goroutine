package go_goroutine

import (
	"fmt"
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

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
