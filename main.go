package main

import (
	"fmt"
	"time"
)

// define struct pesan ke pacar
type PesanCintaKePacar struct {
	From  string
	Pesan string
}

type Server struct {
	ChannelMessage chan PesanCintaKePacar
	ChannelQuit    chan struct{}
}

func (s *Server) StartAndListen() {
free:
	for {
		select {
		case msg := <-s.ChannelMessage:
			fmt.Printf("looping msg channel From:%s dan Pesan %s\n", msg.From, msg.Pesan)
		case <-s.ChannelQuit:
			fmt.Println("server is shutwon")
			break free
		default:
		}
	}
}

func ServerChannelQuit(ChannelQuit chan struct{}) {
	close(ChannelQuit)
}

func SeedMessageToServer(ChannelMessage chan PesanCintaKePacar, pesan string) {
	msg := PesanCintaKePacar{
		From:  "akramulfata",
		Pesan: pesan,
	}
	ChannelMessage <- msg
}

func main() {
	s := &Server{
		ChannelMessage: make(chan PesanCintaKePacar),
		ChannelQuit:    make(chan struct{}),
	}
	// run server
	go s.StartAndListen()
	// call gorutine send message
	go func() {
		time.Sleep(2 * time.Second)
		SeedMessageToServer(s.ChannelMessage, "server is running")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ServerChannelQuit(s.ChannelQuit)
	}()

	select {}
}

// func main() {
// new := time.Now()
// userId := 10
// sendChannel := make(chan string, 128)
// berikan jumlah seluruh gorutine yang dijalankan agar tidak terjadi deedlock
// wg := &sync.WaitGroup{}
// go getUserId(userId, sendChannel, wg)
// wg.Add(1)
// go getUserIdRecomendasion(userId, sendChannel, wg)
// wg.Add(1)
// go getUserIdRecomendasion(userId, sendChannel, wg)
// wg.Add(1)
// wg.Wait()
// setelah berhasil kita matikan
// close(sendChannel)
// walupun kita sudah close channel nya dia tetap deadlock

// 281.480372ms tanpa go gorutine
// fetUserId := getUserId(userId)
// fetUserRecomendasi := getUserIdRecomendasion(userId)
// fetGeUserLike := getUserIdLike(userId)

// call func tanpa gorutine
// fmt.Println(fetUserId)
// fmt.Println(fetUserRecomendasi)
// fmt.Println(fetGeUserLike)

// loop to call channel
// for sendChannels := range sendChannel {
// 	fmt.Println(sendChannels)
// }
// ini akan deedlocck error karena kita tidak mengatur jumlah channel yang berjalan
// fmt.Println(time.Since(new))

// }

// func getUserId(userId int, sendChannel chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	// return "return dara user"
// 	sendChannel <- "return dara user"

// 	wg.Done()
// }

// func getUserIdRecomendasion(userId int, sendChannel chan string, wg *sync.WaitGroup) {
// 	time.Sleep(150 * time.Millisecond)
// 	// return "get user recomendasi"
// 	sendChannel <- "get user recomendasi"

// 	wg.Done()
// }

// func getUserIdLike(userId int, sendChannel chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)
// 	// return "get like user"
// 	sendChannel <- "get like user"
// 	wg.Done()
// }
