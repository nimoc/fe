package go_trap

import (
	gtest "github.com/og/x/test"
	"log"
	"testing"
	"time"
)


func TestSliceEmpty(t *testing.T) {
	as := gtest.NewAS(t)
	var nilSlice []string
	emptySlice := []string{}
	as.Equal(len(nilSlice), 0)
	as.Equal(len(emptySlice), 0)
	as.NotEqual(nilSlice, emptySlice)
	// Not equal:
	// expected: []string(nil)
	// actual  : []string{}
	as.Equal(nilSlice, []string(nil))
}

func TestTime(t *testing.T) {
	time.Sleep(1)
	log.Print("1 not one second, but no errors will be reported at compile time")
	time.Sleep(1*time.Second)
	log.Println("You should use 1 * time.Second")
}