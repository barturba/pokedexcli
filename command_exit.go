package main

import "os"

func commandExit(cfg *config, param string) error {
	os.Exit(0)
	return nil
}
