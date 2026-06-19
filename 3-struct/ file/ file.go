package file

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func WriteFile(content []byte, name string) {
	err := os.WriteFile(name, content, 0644)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	fmt.Println("Файл сохранён:", name)
}