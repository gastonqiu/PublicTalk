package main

var j uint32
var Eface interface{} // outsmart compiler (avoid static inference)

func assertion() {
	i := uint32(42)
	Eface = i
	j = Eface.(uint32)
}
