package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// 1. Перебираем все значения слайса
	for _, data := range dataset {
		// 2. Парсим значение методом Parse()
		err := dp.Parse(data)
		if err != nil {
			// 3. Ошибку парсинга логируем и переходим к следующей итерации
			log.Println(err)
			continue
		}

		// 4. Формируем строку методом ActionInfo()
		info, err := dp.ActionInfo()
		if err != nil {
			// при ошибке логируем её
			log.Println(err)
			continue
		}

		// выводим строку на экран
		fmt.Println(info)
	}
}
