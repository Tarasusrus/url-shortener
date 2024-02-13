package helpers

import (
	"fmt"
	"log"
	"runtime"
)

// LogError логирует ошибку и информацию о месте ее возникновения.
func LogError(err error) {
	if err != nil {
		log.Println(CallerInfo(), "Error: ", err)
	}
}

// CallerInfo возвращает информацию о вызывающей функции.
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
