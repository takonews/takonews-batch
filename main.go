package main

import (
	"fmt"
	"time"

	_ "github.com/takonews/takonews-api/db/migrations"
	"github.com/takonews/takonews-batch/cron"
)

func main() {
	// cron set
	cron := cron.Cron
	fmt.Println("=========cron start=========")
	cron.Start()
	defer cron.Stop()

	for {
		time.Sleep(1000000000000)
	}
}
