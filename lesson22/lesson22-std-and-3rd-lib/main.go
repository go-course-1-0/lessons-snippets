package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"math"
	"os"
	"slices"
	"strings"
)

// fmt, math, strings, slices, errors, slog, json

func main() {
	str := fmt.Sprint(123)
	fmt.Printf("%v - %T\n", str, str)

	var a int
	//fmt.Scan(&a)
	fmt.Sscan(str, &a)

	fmt.Printf("%v - %T\n", str, str)
	fmt.Printf("%v - %T\n", a, a)

	hw := fmt.Sprintf("Hello world %d", a)
	fmt.Printf("%v - %T\n", hw, hw)

	var first, second string
	var third int

	fmt.Sscan(hw, &first, &second, &third)
	fmt.Printf("%v - %T\n", first, first)
	fmt.Printf("%v - %T\n", second, second)
	fmt.Printf("%v - %T\n", third, third)

	var b, c float64 = 10, 5.5

	fmt.Println(math.Pow(b, c))
	fmt.Println(math.Sqrt(81))
	fmt.Println(math.Pow(81, 1.0/2))

	fmt.Println(strings.Contains(hw, "Siyovush"))
	fmt.Println(strings.Contains(hw, "World"))
	fmt.Println(strings.Contains(hw, "world"))
	userInput := "WorLd"

	fmt.Println(strings.Contains(strings.ToLower(hw), strings.ToLower(userInput)))
	fmt.Println(strings.Count(hw, "o"))

	for _, word := range strings.Split(hw, " ") {
		fmt.Println(word)
	}

	sl := strings.Split(hw, " ")
	fmt.Println(sl)

	names := []string{"Alice", "Bob", "Vera"}
	names = slices.Insert(names, 1, "Bill", "Billie")
	names = slices.Insert(names, len(names), "Zac")
	fmt.Println(names)

	fmt.Println("Такого студента не существует")
	log.Println("This student does not exist")

	var ErrStudentNotFound = errors.New("такого студента не существует")

	fmt.Printf("%v - %T\n", ErrStudentNotFound, ErrStudentNotFound)
	fmt.Printf("%v - %T\n", ErrStudentNotFound.Error(), ErrStudentNotFound.Error())

	log.Println(ErrStudentNotFound)
	log.Println(ErrStudentNotFound)
	log.Println(ErrStudentNotFound)
	log.Println(ErrStudentNotFound)
	//log.Fatalln(ErrStudentNotFound)

	NewError := errors.Join(ErrStudentNotFound, errors.New("hello world"))
	fmt.Printf("%v - %T\n", NewError, NewError)

	NewError2 := fmt.Errorf("прикиньте, ошибка: %w", ErrStudentNotFound)
	fmt.Printf("%v - %T\n", NewError2, NewError2)

	err := errors.Unwrap(NewError2)

	fmt.Printf("%v - %T\n", err, err)
	fmt.Printf("%v - %T\n", NewError2, NewError2)

	slog.Info("Hello")
	slog.Error("This is error")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello", "count", 3)
	logger.Error(ErrStudentNotFound.Error(), "id", 54, "email", "some@email.com", "phone", "999888777666")

	lg := `{"time":"2024-03-25T18:31:55.609880223+05:00","level":"ERROR","msg":"такого студента не существует","id":54,"email":"some@email.com","phone":"999888777666"}`

	var logRecord map[string]interface{}

	json.Unmarshal([]byte(lg), &logRecord)

	fmt.Println(logRecord)
	for key, value := range logRecord {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

	var lr lRecord

	json.Unmarshal([]byte(lg), &lr)
	fmt.Printf("%v - %T\n", lr, lr)

	s, _ := json.Marshal(lr)
	fmt.Println(string(s))

	s, _ = json.MarshalIndent(lr, "", "\t")
	fmt.Println(string(s))
}

type lRecord struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"msg"`
}
