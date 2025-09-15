package main

/*
original code from task

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

// корректный вариант
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//скопируем строку нужной длины для использования
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}
