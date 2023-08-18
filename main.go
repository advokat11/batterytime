package main

import (
	"fmt"
	"os"
	"time"

	"github.com/distatus/battery"
)

func main() {

	f, err := os.Create(os.Getenv("USERPROFILE") + "\\Desktop\\times.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Запись текущего времени в первую строку
	f.WriteString(time.Now().Format("15:04:05\n"))

	// Ждем 30 минут и записываем время во вторую строку
	time.Sleep(30 * time.Minute)
	f.WriteString(time.Now().Format("15:04:05\n"))

	// Ждем еще 30 минут и записываем время в третью строку
	time.Sleep(30 * time.Minute)
	f.WriteString(time.Now().Format("15:04:05\n"))

	// Ждем пока заряд ноутбука достигнет 100%
	for {
		battery, _ := battery.Get(0)
		if battery.Current >= 100 {
			f.WriteString(time.Now().Format("15:04:05\n"))
			break
		}
		time.Sleep(10 * time.Second)
	}
}
