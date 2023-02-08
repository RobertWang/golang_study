package main

func main() {
	var a int8 = -1
	var b int8 = -128 / a

	println(b)
}

/*
package main

func main1() {
    const a int8 = -1
    var b int8 = -128 / a
	// -128 / a (constant 128 of type int8) overflows int8compilerNumericOverflow

    println(b)
}
*/
