package main

import (
	"fmt"

	"dimagorbatko5-cell/DZps-GO/3-struct/bins"
	"dimagorbatko5-cell/DZps-GO/3-struct/storage"
)

func main() {
	list := bins.NewBinList()

	bin := bins.NewBin("1", "Мой первый bin", false)
	list.AddBin(bin)

	storage.Save(list, "bins.json")

	loaded, err := storage.Load("bins.json")
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
	fmt.Println(loaded)
}