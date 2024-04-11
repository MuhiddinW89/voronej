package consoleinput

import (
	"app/api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ConsoleInput() {
	// Получаем аргументы командной строки
	args := os.Args[1:]

	// Проверяем, переданы ли аргументы
	if len(args) == 0 {
		fmt.Println("Пожалуйста, укажите числа через запятую после команды.")
		return
	}

	// Получаем строку аргументов
	numbersStr := args[0]

	// Разбиваем строку на отдельные числа
	numbers := strings.Split(numbersStr, ",")

	// Создаем слайс для хранения чисел
	var parsedNumbers []int

	// Проходим по каждому элементу и преобразуем его в число
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Ошибка при преобразовании числа %s: %v\n", numStr, err)
			return
		}
		parsedNumbers = append(parsedNumbers, num)
	}

	// Выводим результат
	fmt.Println("Введенные числа:")
	for _, num := range parsedNumbers {
		fmt.Println(num)
	}
}

func ConsoleOutput() {
	time.Sleep(1 * time.Second)

	response, err := http.Get("http://localhost:4001/product/6")
	if err != nil {
		log.Printf("Ошибка при выполнении запроса к серверу: %v", err)
	}

	defer response.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Ошибка при чтении ответа сервера: %v", err)
		return
	}

	var product models.Product

	// Распаковываем JSON и помещаем его в структуру модели
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Printf("Ошибка при разборе JSON: %v", err)
		return
	}
	fmt.Println(string(body))
}
