package helpers

import (
	"fmt"
	"log"
	"runtime"
)

// LogError получает ошибку в качестве аргумента.
// Если ошибка не равна nil, функция логирует информацию об ошибке и местонахождении вызывающей функции.
func LogError(err error) {
	if err != nil {
		log.Println(CallerInfo(), "Error: ", err)
	}
}

// CallerInfo возвращает имя вызывающей функции, имя файла и номер строки вызыва.
// Эта информация получается с помощью функции runtime.Caller(2),
// где передаваемый номер относится к глубине стека вызовов, отсчитываемый от самой CallerInfo функции.
// Эта функция помогает определить местоположение, откуда была вызвана функция LogError.
func CallerInfo() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		// Если функция runtime. Caller не смогла получить информацию, возвращает строку "unknown caller"
		return "unknown caller"
	}

	fn := runtime.FuncForPC(pc)

	// Возвращает строку, содержащую имя вызывающей функции, имя файла и номер строки.
	return fmt.Sprintf("%s() - %s:%d", fn.Name(), file, line)
}
