package main

import (
	"github.com/TelitsynNikita/go_core_rest_api"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := go_core_rest_api.NewApiService()

	srv.InitApiService()

	if err := srv.RunService(); err != nil {
		logrus.Fatal(err)
	}
}
