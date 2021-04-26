package main

import (
	"log"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	acq := os.Getenv("ACQ_MERCHANT_ID")
	log.Println("====================== ", acq)
}
