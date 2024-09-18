package main

import (
	api "api/internal/commit-tracker"
)

func main() {

	api.NewAPIService().Start()
}
