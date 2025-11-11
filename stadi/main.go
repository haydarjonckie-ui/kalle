package main

import "fmt"

func main() {
	fmt.Println("дарс таер кунед ")
} 

package main

import (
  "fmt"
  "time"
)

type Car struct {
  Model       string
  Color       string
  Year        int
  Hp          int
  HasTuning   bool
  TuningLevel int
}

func main() {
  car := Car{
    Model:       "06",
    Color:       "Black",
    Year:        1971,
    Hp:          300,
    HasTuning:   true,
    TuningLevel: 3,
  }
  // car1 := Car{Model: "06", Color: "Black", Year: 1971, Hp: 300}

  // ptr := &car

  fmt.Println("\n-------MENU-------")
  fmt.Println("1. Перекрасить машину")
  fmt.Println("2. Увеличить мощность")
  fmt.Println("3. Добавить тюнинг")
  fmt.Println("4. Показать текущее состояние машины")
  fmt.Println("0. Выход")
  fmt.Print("Выберите действие: ")

  var choice int
  fmt.Scan(&choice)

  switch choice {
  case 1:
    RepaintCar(&car)
  case 2:
    PowerUp(&car)
  case 3:
    Tuning(&car)
  case 4:
    ShowCarStatus(&car)
  case 0:
    fmt.Println("Выход из программы")
    return
  default:
    fmt.Println("Неверный выбор. Повторите попытку.")
  }
}

func RepaintCar(c *Car) {
  fmt.Println()
  var newColor string
  fmt.Print("Введите в какой цвет перекрасить твой аппарат: ")
  fmt.Scan(&newColor)

  if newColor == c.Color {
    fmt.Println("Зачем перекрашивать в тот же цвет скажи да")
    return
  }
  for i := 1; i <= 3; i++ {
    fmt.Printf("Покраска... шаг %v из 3\n", i)
    time.Sleep(3 * time.Second)
  }
  c.Color = newColor
  fmt.Printf("дедушка твоя перекрашена в %v цвет\n", c.Color)
  fmt.Println()
}

func PowerUp(c *Car) {
  fmt.Println()
  var extraHp int

  fmt.Println("насколько лошадиных сил хотите увеличить мощь?")
  fmt.Scan(&extraHp)

  if extraHp <= 0 {
    fmt.Println("Нельзя увеличить мощность на отрицательное либо нулевое значение")
    return
  }

  hp := c.Hp + extraHp
  if hp >= 1000 {
    fmt.Println("Мужик, такой возможности нет, максимум поставим 1000 л.с.")
    c.Hp = 1000
  } else {
    c.Hp = hp
  }
  fmt.Println("Мощность увеличена, мощность твоего аппарата:", c.Hp)
  fmt.Println()
}

func Tuning(t *Car) {
  fmt.Println()
  var level int
  if t.HasTuning == true {
    fmt.Println("Тюнинг уже установлен")
    fmt.Printf("Текущий уруэн тюнинга %v\n", t.TuningLevel)
    fmt.Print("Хотите повысить уруэн? (1 - да || 0 - нет): ")
    var choice int
    fmt.Scan(&choice)
    if choice == 0 {
      return
    }
  }
  fmt.Print("Введите уруэн тюнинга (1-5): ")
  fmt.Scan(&level)
  if level < 1 || level > 5 {
    fmt.Println("Недопустимый уруэн тюнинга")
    return
  }

  t.TuningLevel = level
  fmt.Printf("Тюнинг установлен, текущий уруэн тюнинга %v\n", t.TuningLevel)
  fmt.Println()
}

func ShowCarStatus(s *Car) {
  fmt.Println()
  fmt.Println("Текущее состояние машины:")
  fmt.Printf("модель: %v\n", s.Model)
  fmt.Printf("цвет машины: %v\n", s.Color)
  fmt.Printf("Год выпуска: %v\n", s.Year)
  fmt.Printf("Мощность двигателя %v лошадиных сил\n", s.Hp)
  if s.HasTuning {
    fmt.Printf("  Тюнинг: установлен (уровень %v)\n", s.TuningLevel)
  } else {
    fmt.Println("  Тюнинг: отсутствует")
  }
  fmt.Println()
}