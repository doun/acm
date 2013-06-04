package main

import "testing"
import "fmt"

func TestCombine(t *testing.T){
	for i:=1;i<9; i++{
		k :=0
		fmt.Println("trying:",i)
		Combine(9,i,func(g []int)(bool){
			fmt.Println(g)
			k++
			return false
		})
	}
}
