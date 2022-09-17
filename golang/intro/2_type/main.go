package main

import (
	"fmt"
)

func main() {
	// Подробнее о типах данных в Go
	// Вы можете прочитать в документации
	// В целом, для знакомства с Go в первую очередь вы должны пройти
	// golang tour. Он проведёт по всем синтаксическим конструкциям языка

	// В Go есть константы. Константы не могут быть изменены, также
	// не может быть получена ссылка на константу. Это сделано для того, что
	// константа была собственно константой, не могла быть изменена
	// ни при каких условиях
	const t = 123

	// Если вы раскомментируете строку, то увидите ошибку линтера (если он настроен)
	// или увидите ошибку при запуске программы
	// fmt.Println(&t)

	// Такая задача решается следующим образом (если нужна правда ссылка)
	v := t
	fmt.Println(&v)

	// Важная часть языка - структуры
	// В Go нет так называемых классов
	// Но есть структуры, у которых могут быть методы
	type T struct {
		A string
	}

	type MyType struct {
		A string
		B []string
		C map[string]string
		D *string
		E *T
	}

	str := MyType{
		A: "str",
		// E: &T{A: "1"},
	}

	fmt.Println(str.A)
	fmt.Println(str.B)
	fmt.Println(str.C)
	// проверка str.C на nil даст true
	fmt.Println(str.C == nil)
	fmt.Println(str.D)
	fmt.Println(str.E)
	// fmt.Println(str.E.A)

	// Обратите внимание на работу с указателями
	// E является указателем на тип T
	// Если E - nil, то обращение E.A приведёт к panic
	// и падению приложения

	if str.E == nil {
		fmt.Println("Если бы обратились к str.E.A была бы panic")
	}

	// Поэтому не забывайте проверять переменные, являющиеся указателями на nil
	if str.E != nil {
		fmt.Println("Если вы укажите значение E, то сработает этот код")
	}

	// Если map не объявлена, то этот код приведёт к panic:
	// (раскомментируйте и запустите)
	// str.C["1"] = "1"

	// Чтобы не было паники необходимо аллоцировать map:
	str.C = make(map[string]string)
	fmt.Println(str.C)
	// а в данном случае str.C на nil даст false
	fmt.Println(str.C == nil)

	str.C["1"] = "1" // без паники!

	// Ниже (вне функции main)
	// была объявлена структура с методами
	mstr := M{
		A: "a",
	}

	fmt.Println(mstr.GetA())
	fmt.Println(mstr.GetAFromPointerM())

	// Используется, когда функция возвращает только одну
	// ошибку или нам нужно из функции использовать только ошибку
	if err := mstr.IsBAllocate(); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := mstr.IsBAllocateWithBool(); err != nil {
		fmt.Println(err.Error())
	}

	// Используется, когда функция возвращает не только ошибку
	err := mstr.IsBAllocate()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Обратите внимание, что в таком случае используем просто =
	// потому что err объявлена выше
	_, err = mstr.IsBAllocateWithBool()
	if err != nil {
		fmt.Println(err.Error())
	}

	sl := make([]string, 0, 10)
	sl = append(sl, "10")
}

type M struct {
	A string
	B map[string]string
}

func (m M) GetA() string {
	return m.A
}

func (m *M) GetAFromPointerM() string {
	return m.A
}

// Например мы хотим написать функцию, которая проверяет
// аллоцировано ли B, но например если нет, мы хотим
// вернуть ошибку, а не просто bool
func (m *M) IsBAllocate() error {
	if m.B == nil {
		return fmt.Errorf("B is not allocate")
	}

	return nil
}

func (m *M) IsBAllocateWithBool() (bool, error) {
	if m.B == nil {
		return false, fmt.Errorf("B is not allocate")
	}

	return true, nil
}
