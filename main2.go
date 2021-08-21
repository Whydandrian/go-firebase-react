package main

import (
	"fmt"

	"github.com/Whydandrian/go-firebase-react/firebase"
)

func mains()  {
	bucket := firebase.CloudStorage()
	fmt.Println(bucket)

}