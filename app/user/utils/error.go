package utils

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
)

// MustHandleError log the error info and then exit
func MustHandleError(err error) {
	if err != nil {
		fmt.Print("MERROE!", err)
		klog.Fatal(err)
	}
}

// ShouldHandleError log the error info
func ShouldHandleError(err error) {
	if err != nil {
		fmt.Print("SERROE!", err)
		klog.Error(err)
	}
}
