package main

import (
	"fmt"
	"github.com/amir79esmaeili/sms-gateway/cmd"
	"os"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
