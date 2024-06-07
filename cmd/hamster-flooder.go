package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var letterRunes = []rune("abcdefg1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	var wg sync.WaitGroup
	client := &http.Client{}
	client.Timeout = 2000 * time.Millisecond

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(*client)
		}()
	}

	wg.Wait()
	log.Print("Done")
}

func Worker(client http.Client) {
	req, err := http.NewRequest("GET", "https://stream.nknikolay.ru/st.php", nil)
	if err != nil {
		log.Print(err)
	}

	cookie := http.Cookie{
		Name:  "st",
		Value: "66631" + RandStringRunes(8),
	}
	req.AddCookie(&cookie)

	for {
		resp, err := client.Do(req)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(resp.StatusCode)
	}
}
