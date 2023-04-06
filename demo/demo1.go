package main

import (
	"context"
	"time"
)

func main() {

	ctx1, _ := context.WithCancel(context.TODO())

	_, cancel := context.WithTimeout(ctx1, 1000*time.Second)

	cancel()
	cancel()
}
