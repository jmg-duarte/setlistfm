package api

import (
	"testing"
	"fmt"
)

func TestOne(t *testing.T){
	res , _ := GetArtistSetlists("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", 0)
	fmt.Println(res)
}