/*
Package NAME OF YOUR PACKAGES implements ....

	Build with: go version go1.19.2 darwin/arm64
	Created by Uğur Özyılmazel on 2022-10-29.
	Copyright (c) 2022 VB YAZILIM. All rights reserved.
*/
package main

import (
	"fmt"
	"os"

	"github.com/vigo/ghstars/src/ghstars"
)

func main() {
	cmd := ghstars.New(nil)
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
