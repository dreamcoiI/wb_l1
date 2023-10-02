package main

import "os"

func createHugeString(size int, f *os.File) error {
	buffer := make([]byte, size)
	_, err := f.Write(buffer)
	return err
}

func someFunc() {
	f, err := os.Create("justString.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	err = createHugeString(1<<30, f) // лучше и быстрее держать большую строку в файле а не в глобальной переменной
	if err != nil {
		panic(err)
	}
}

func main() {
	someFunc()
}
