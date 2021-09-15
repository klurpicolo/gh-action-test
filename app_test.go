package main

import (
 "testing"
)

func TestPlus1And2Shouldbe3(t *testing.T) {
     v := Plus(1,2)
     if 3 != v { 
       t.Error("1 + 2 should be 3 but", v)
     }
}

func TestPlus0and4Shouldbe5(t *testing.T) {
  v := Plus(0,4)
  if 5 != v { 
    t.Error("0 + 4 should be 5 but", v)
  }
}