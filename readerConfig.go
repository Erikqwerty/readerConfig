package rc

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

// Читает конфигурацию из указанного файла и создает карту map[параметр]значение.
// Формат конфигурационного файла.
// параметр = значение;
// параметр2=значение; и тд.
// ";" - обязательно.
func readerConfig(dir string) (map[string]string, error) {
	// читаем файл в буфер
	fileText, err := readeFile(dir)
	if err != nil {
		return nil, nil
	}

	// Удаление пробелова из конфигурационного файла для упращения алгоритма.
	File := strings.Replace(fileText, " ", "", -1)

	// Разделяем перечень параметров в конфигурационном файле через ;
	// Получаем срез строк в формате параметр=значение
	slFile := breakString(File, ';')

	// Инициализируем карту в которую будем складывать map[параметр]значение.
	config := make(map[string]string)
	// Цикл для переборки среза строк slFile (в формате параметр=значение) и создания
	// карты. Разделяем параметр и значение на 2 строки записываем в срез для того чтобы
	// создать карту map[параметр]значение. arg+" " - костыль алгоритм не совершенен.
	for _, arg := range slFile {
		slArg := breakString(arg+" ", '=')
		config[slArg[0]] = slArg[1]
	}

	return config, nil
}

// Читает файл возвращает строку файла.
func readeFile(dir string) (string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Чтение файла с ридером
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	return wr.String(), nil
}

// Функция для создания среза из строки. Разделяет строку (текст|char|текст).
// Новые элементы среза ([0]текст(|char|[1]текст... и т.д.).
func breakString(str string, char rune) []string {
	slice := []string{}
	for len(str) > 0 {
		for i, ch := range str {
			if ch == char || len(str) == i+1 {
				slice = append(slice, str[:i])
				str = str[i+1:]
				break
			}
		}
	}
	return slice
}
