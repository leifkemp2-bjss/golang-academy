package main

import(
	"github.com/bearbin/go-age"
	"time"
)

func assignment6() {
	date := time.Date(2000, 3, 10, 0, 0, 0, 0, time.UTC)

	res := age.Age(date)

	pl(res)
}
