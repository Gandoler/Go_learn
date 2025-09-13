# Анонимные функции (Лямбды) в Go

В Go можно создавать **анонимные функции** — функции без имени.
Они объявляются прямо на месте и часто используются как вспомогательные: для краткости, конфигурации поведения или передачи логики в другие функции.

---

## Объявление анонимной функции

```go
func(x int, y int) int {
	return x + y
}
````

> Эта конструкция **определяет функцию**, но не вызывает её.

Чтобы вызвать такую функцию сразу, добавляют круглые скобки:

```go
result := func(x int, y int) int {
	return x + y
}(3, 4)

fmt.Println(result) // 7
```

> Такая запись называется **немедленным вызовом анонимной функции** (IIFE — immediately invoked function expression).

---

## Захват внешних переменных (Замыкание)

Анонимные функции могут использовать переменные, определённые вне их тела:

```go
func main() {
	suffix := "!"

	exclaim := func(text string) string {
		return text + suffix
	}

	fmt.Println(exclaim("Hello")) // Hello!
}
```

> Функция `exclaim()` получила доступ к переменной `suffix`, даже не принимая её как параметр. Это называется **замыканием**.

---

## Пример на будущее: сортировка с анонимной функцией

```go
import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"go", "hexlet", "code"}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	fmt.Println(words) // [go code hexlet]
}
```

> Здесь `sort.Slice()` использует анонимную функцию-компаратор, чтобы определить порядок элементов.
> Скоро вы узнаете, что такое срезы и как с ними работать.

---

## Функции, возвращающие другие функции

В Go функция может **возвращать другую функцию**, что часто используется для создания замыканий:

```go
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	fmt.Println(double(4)) // => 8
	fmt.Println(triple(4)) // => 12
}
```

> Функции `double` и `triple` «запомнили» своё окружение и используют соответствующий множитель.
