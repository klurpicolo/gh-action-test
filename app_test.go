package main

import (
 "testing"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {
     v := Plus(1,2)
     if 3 != v { 
       t.Error("fizzbuzz of 1 should be '1' but have", v)
     }
}