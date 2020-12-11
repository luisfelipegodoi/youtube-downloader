package main

import (
	"fmt"
	"os/exec"
)

func main() {

	output, err := exec.Command("./youtubedr", "download", "Sv6dMFF_yts").Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(output))
}
