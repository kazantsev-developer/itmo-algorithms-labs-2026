package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name       string
	IsMale     bool
	IsSmoker   bool
	IsBrunette bool
	FromITMO   bool
	IsSports   bool
	HasPets    bool
	HasDog     bool
}

func main() {
	groupData := []Student{
		{Name: "Егорова Мария", IsMale: false, IsSmoker: false, IsBrunette: true, FromITMO: false, IsSports: true, HasPets: true, HasDog: true},
		{Name: "Фролова Кристина", IsMale: false, IsSmoker: false, IsBrunette: false, FromITMO: true, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Костюченко Тимофей", IsMale: true, IsSmoker: false, IsBrunette: false, FromITMO: false, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Портнова Ксения", IsMale: false, IsSmoker: true, IsBrunette: false, FromITMO: true, IsSports: false, HasPets: true, HasDog: true},
		{Name: "Кенжаев Рахим", IsMale: true, IsSmoker: false, IsBrunette: true, FromITMO: true, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Даньшин Семён", IsMale: true, IsSmoker: true, IsBrunette: true, FromITMO: true, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Шихайло Сергей", IsMale: true, IsSmoker: true, IsBrunette: false, FromITMO: true, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Бабушкин Александр", IsMale: true, IsSmoker: true, IsBrunette: false, FromITMO: true, IsSports: false, HasPets: false, HasDog: false},
		{Name: "Якунин Андрей", IsMale: true, IsSmoker: false, IsBrunette: false, FromITMO: true, IsSports: true, HasPets: false, HasDog: true},
		{Name: "Казанцев Александр", IsMale: true, IsSmoker: false, IsBrunette: false, FromITMO: false, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Юркин Александр", IsMale: true, IsSmoker: true, IsBrunette: true, FromITMO: true, IsSports: true, HasPets: false, HasDog: false},
		{Name: "Панкратова Анна", IsMale: false, IsSmoker: true, IsBrunette: false, FromITMO: false, IsSports: true, HasPets: true, HasDog: false},
		{Name: "Колесникова Лариса", IsMale: false, IsSmoker: false, IsBrunette: false, FromITMO: false, IsSports: true, HasPets: false, HasDog: false},
	}

	reader := bufio.NewReader(os.Stdin)
	ask := func(q string) bool {
		fmt.Print(q + " (да/нет): ")
		text, _ := reader.ReadString('\n')
		return strings.TrimSpace(strings.ToLower(text)) == "да"
	}

	fmt.Println("Система определения студентов")

	male := ask("Студент мужского пола?")
	smoke := ask("Студент курит?")
	brunette := ask("Студент является брюнетом?")
	itmo := ask("Студент учился в ИТМО в бакалавриате?")
	sports := ask("Студент занимается спортом?")
	pets := ask("У студента есть домашние животные?")

	dog := false
	if pets {
		dog = ask("Это собака?")
	}

	var matched []string
	for _, s := range groupData {
		if s.IsMale == male &&
			s.IsSmoker == smoke &&
			s.IsBrunette == brunette &&
			s.FromITMO == itmo &&
			s.IsSports == sports &&
			s.HasPets == pets &&
			s.HasDog == dog {
			matched = append(matched, s.Name)
		}
	}

	fmt.Println("\nРезультат поиска:")
	if len(matched) == 1 {
		fmt.Printf("Вы загадали студента: %s\n", matched[0])
	} else if len(matched) > 1 {
		fmt.Printf("Под описание подходит несколько студентов: %s\n", strings.Join(matched, ", "))
	} else {
		fmt.Println("Студент с такими характеристиками в базе не найден.")
	}
}
