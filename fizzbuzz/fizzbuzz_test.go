package fizzbuzz

import "testing"
import (
	. "github.com/smartystreets/goconvey/convey"
)


func TestNonMultiples(t *testing.T){
	Convey("If number is not a multiple of 3, 5 or both just return the number", t, func() {
		So(Fizzbuzz(11), ShouldEqual, "11")
	} )
}

func TestFizz(t *testing.T){
	Convey("If n is a multiple of 3, return fizz", t, func() {
		So(Fizzbuzz(3), ShouldEqual, "fizz" )
	})
}

func TestBuzz(t *testing.T){
	Convey("If n is a multiple of 5, return buzz", t, func() {
		So(Fizzbuzz(10), ShouldEqual, "buzz")
	})
}

func TestFizzbuzz(t *testing.T){
	Convey("If n is a multiple of both 3 and 5, return Fizzbuzz", t, func() {
		So(Fizzbuzz(60), ShouldEqual, "fizzbuzz")
	})
}