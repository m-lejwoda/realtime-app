package main

import( 
	"os"
	"fmt"
	"bufio"
	"github.com/gorilla/websocket"
	)

func main() {
	fmt.Println("Add user")
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil{
		fmt.Println("Error during connection appeared", err)
	}	
	defer conn.Close()
	userInput := bufio.NewScanner(os.Stdin)
	go func() {
		_, message, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("Error", err)
			return 
		}
		fmt.Println("message", message)
	}()

	for{
		if !userInput.Scan(){
			break
		}
		text := userInput.Text()
		fmt.Println("text", text)
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		fmt.Println("Wyslane")
		if err != nil{
			fmt.Println("Sending Error", err)
		}
		xd,test,xd2 := conn.ReadMessage()
		fmt.Println("xd",xd)
		fmt.Println("test",test)
		fmt.Println(xd2)

	}
}