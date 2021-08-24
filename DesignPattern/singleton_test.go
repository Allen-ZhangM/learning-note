package DesignPattern

import "testing"

func TestGetInstance(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
}

func TestGetInstanceOnce(t *testing.T) {
	for i := 0; i < 30; i++ {
		go getInstanceOnce()
	}
}
