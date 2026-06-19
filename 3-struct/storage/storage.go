package storage

import (
	"encoding/json"
	"fmt"
	"dimagorbatko5-cell/DZps-GO/3-struct/bins"   // ← замени demo на имя своего модуля (см. ниже)
	"dimagorbatko5-cell/DZps-GO/3-struct/file"
)

func Save(list bins.BinList, fileName string) {
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	file.WriteFile(data, fileName)
}

func Load(fileName string) (bins.BinList, error) {
	var list bins.BinList
	data, err := file.ReadFile(fileName)
	if err != nil {
		return list, err
	}
	err = json.Unmarshal(data, &list)
	return list, err
}