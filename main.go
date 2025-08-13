package main

import (
	"fmt"
	"time"

	"github.com/agb1986/stock-ticker/internal"
	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		now := time.Now()
		fmt.Printf("\nFetching stocks @ %s\n", now.Format(time.DateTime))
		internal.TableRender()
		fmt.Printf("\nExecution Time: %s\n", time.Since(now))
		robotgo.KeyTap("f10")
		robotgo.Sleep(840)
	}
}
