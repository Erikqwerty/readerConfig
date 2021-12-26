package rc

import (
	"fmt"
	"testing"
)

func TestReader(t *testing.T) {
	//Arrange
	testTable := []struct {
		configFile string
		Expected   map[string]string
	}{
		{
			configFile: "./config/app.conf",
			Expected: map[string]string{
				"app":    "newApp",
				"addr":   "127.0.0.1:2324",
				"passwd": "admin123",
			},
		},
		{
			configFile: "./config/test_1.conf",
			Expected: map[string]string{
				"addr": "smtp.gmail.com",
				"port": "587",
			},
		},
	}
	//Act
	for _, testCase := range testTable {
		rezult, err := readerConfig(testCase.configFile)
		if err != nil {
			fmt.Println(err)
		}

		expect := mapToString(testCase.Expected)
		rez := mapToString(rezult)

		//Assert
		if expect != rez {
			t.Error("\n Ожидаемый вывод:", testCase.Expected, "\n Полученый вывод:", rezult)
		}
	}
}

func mapToString(m map[string]string) string {
	var s string
	for key, val := range m {
		s = fmt.Sprintf("%s=\"%s\"", key, val)
	}
	return s
}
