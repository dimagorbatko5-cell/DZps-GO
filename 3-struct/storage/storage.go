package storage

import (
	"encoding/json"
	"fmt"

	"dimagorbatko5-cell/DZps-GO/3-struct/bins"
	"dimagorbatko5-cell/DZps-GO/3-struct/files"
)

func Save(list bins.BinList, fileName string) {
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	err = files.WriteFile(data, fileName)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
	}
}

func Load(fileName string) (bins.BinList, error) {
	var list bins.BinList
	data, err := files.ReadFile(fileName)
	if err != nil {
		return list, err
	}
	err = json.Unmarshal(data, &list)
	return list, err
}