package hello

import (
	"fmt"
	"log"
	"math/rand"
    "time"
)

func init()  {
	fmt.Println("hello init")
	rand.Seed(time.Now().UnixNano())
}

func SayHello(name string) (string,error) {
    key,msg:=World(name)
	message :=  fmt.Sprintf(randomFormat(), key)

	if msg != nil {
        log.Fatal(msg)
    }
	return message,msg
}

func SayHellos(names []string)(map[string]string,error){
	messages := make(map[string]string)
	for _,name := range names{
		message, err := SayHello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with 
        // the name.
        messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}