package main

import (
	"fmt"
	"time"

	"github.com/agb1986/stock-ticker/internal"
)

func main() {
	now := time.Now()
	fmt.Printf("\nFetching stocks @ %s\n", now.Format(time.DateTime))
	internal.TableRender()
	fmt.Printf("\nExecution Time: %s\n", time.Since(now))
}
