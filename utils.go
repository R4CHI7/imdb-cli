package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getDisplayableRuntime(runtime string) string {
	runtimeData := strings.Split(runtime, " ")
	if runtimeData[1] == "min" {
		minutes, _ := strconv.Atoi(runtimeData[0])
		return fmt.Sprintf("%d Hours %d Minutes", (minutes / 60), (minutes % 60))
	}
	return runtime
}

func getImdbLink(titleId string) string {
	return fmt.Sprintf("http://www.imdb.com/title/%s", titleId)
}
