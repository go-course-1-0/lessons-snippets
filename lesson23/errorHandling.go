package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	//fmt.Println("Введите сторону квадрата:")
	//fmt.Print("a = ")
	//var a float64
	//_, err := fmt.Scan(&a)
	//if err != nil {
	//	fmt.Println("Некорректный ввод данных:", err)
	//	fmt.Println("Пожалуйста, введите правильное число с плавающей точкой")
	//
	//	fmt.Printf("%v - %T\n", err, err)
	//
	//	err2 := errors.Unwrap(err)
	//	fmt.Printf("%v - %T\n", err2, err2)
	//	return
	//}
	//
	//fmt.Printf("Периметр квадрата со стороной %.2f равна:\n%.2f\n", a, 4*a)
	//
	//var b float64
	//fmt.Println(divide(4, 8))
	//
	//if _, err := fmt.Scan(&a, &b); err != nil {
	//	fmt.Println("Некорректный ввод данных:", err)
	//}
	//
	//res, err := divide(a, b)
	//fmt.Println(res)
	//fmt.Println(err)

	//if err := createStudent("Siyovush", "987654321"); err != nil {
	//	log.Println(err)
	//}
	//if err := createStudent("", "987654321"); err != nil {
	//	log.Println(err)
	//}
	//if err := createStudent("Siyovush", ""); err != nil {
	//	log.Println(err)
	//}
	//if err := createStudent("Siyovush", "987"); err != nil {
	//	log.Println(err)
	//}
	//if err := createStudent("Siyovush", "987987987987987987"); err != nil {
	//	log.Println(err)
	//}
	//if err := createStudent("", ""); err != nil {
	//	log.Println(err)
	//}

	if err := createStudent2("Siyovush", "987654321", 1999); err != nil {
		log.Println(err)
	}
	if err := createStudent2("Siyovush", "987654321", 240); err != nil {
		log.Println(err)
		fmt.Println("err is ErrRequired:", errors.Is(err, ErrMinMax))
		fmt.Println("err as ErrRequired:", errors.As(err, &ErrLen))
	}
	if err := createStudent2("Siyovush", "987654321", 2400); err != nil {
		log.Println(err)
	}
	if err := createStudent2("Siyovush", "987654321", -2400); err != nil {
		log.Println(err)
	}
	if err := createStudent2("", "987654321", 0); err != nil {
		log.Println(err)
	}
	if err := createStudent2("Siyovush", "", 2025); err != nil {
		log.Println(err)
	}
	if err := createStudent2("Siyovush", "987", 2001); err != nil {
		log.Println(err)
	}
	if err := createStudent2("Siyovush", "987987987987987987", -10); err != nil {
		log.Println(err)
	}

	//err := gorm.First()
	//
	//if errors.Is(err, gorm.ErrRecordNotfound) {
	//	log.Println()
	//	404
	//}
	//
	//if errors.Is(err, gorm.ErrDBConn) {
	//	log.Println()
	//	500
	//}
}

// errors.New
// fmt.Errorf

// errors.Is(err1, err2)
// errors.As()

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func createStudent(name, phone string) error {
	// required
	if name == "" {
		return fmt.Errorf("%s студента не может быть пустым", "имя")
	}

	if phone == "" {
		return fmt.Errorf("%s студента не может быть пустым", "телефон")
	}

	// len
	if len(phone) != 9 {
		return fmt.Errorf("длина поля %q должна быть ровно %d символов", "телефон", 9)
	}

	fmt.Println("\nConnecting to DB...")
	fmt.Printf("Saving %s, %s to students table...\n", name, phone)
	fmt.Println("Student record created successfully.\n")

	return nil
}

func createStudent2(name, phone string, yearOfBirth int) error {
	// required
	if name == "" {
		return fmt.Errorf("name: %w", ErrRequired)
	}

	if phone == "" {
		return fmt.Errorf("phone: %w", ErrRequired)
	}

	// len
	if len(phone) != 9 {
		return fmt.Errorf("phone: %w", ErrLen)
	}

	// minMax
	if yearOfBirth < 1900 || yearOfBirth > time.Now().Year() {
		return fmt.Errorf("yearOfBirth: %w", ErrMinMax)
	}

	fmt.Println("\nConnecting to DB...")
	fmt.Printf("Saving %s, %s to students table...\n", name, phone)
	fmt.Println("Student record created successfully.\n")

	return nil
}

var (
	ErrRequired = fmt.Errorf("данное поле обязательно к заполнению")
	ErrLen      = fmt.Errorf("некорректная длина поля")
	ErrMinMax   = fmt.Errorf("некорректное значение")
)
