package main

import (
	"fmt"
	"time"
)


func main(){

	fmt.Println("%b", time.Now().Unix() % 2 == 0)

	client_ch := make(chan string)

	go func(){
		for{
			var b bool
			b = time.Now().Unix() % 2 == 0
			time.Sleep(1 * time.Second)

			fmt.Println("%b", b)
			if b{
				client_ch <- "TC0001"
			}else{
				client_ch <- "TC0002"
			}
		}
	}()

	server_process(client_ch)
}

func server_process(ch chan string){

	server1_ch := make(chan string)
	server2_ch := make(chan string)

	go func(){
		for{
			time.Sleep(1 * time.Second)
			msg := <- ch
			switch msg {
			case "TC0001":
				server1_ch <- msg
				break
			case "TC0002":
				server2_ch <- msg
				break
			}
		}
	}()

	for {
		select {
		case body := <-server1_ch:
			//fmt.Println("channal 1 [%s]", body)
			tc0001_process(body)

		case body2 := <-server2_ch:
			//fmt.Println("channal 2 [%s]", body2)
			tc0002_process(body2)
			/*default:
			fmt.Println("default..")
		}*/
		}
	}
}


func tc0001_process(message string){
	fmt.Println("tc0001 process %s", message)
}

func tc0002_process(message string){
	fmt.Println("tc0002_process %s", message)
}




