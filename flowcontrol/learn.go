package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func Learn() {
	//if
	score := 10
	if score == 10 {
		fmt.Println("Very Good!")
	} else {
		fmt.Println("Try Again!")
	}

	//switch
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MacOS")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Unknown")
	}

	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	content := string(b)
	fmt.Println(content)

	//for
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i+1)
	// }

	//break
	for i := 5; i >= 0; i-- {
		if i == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}