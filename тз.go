package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	for strings.Contains(t.Content, "  ") {
		t.Content = strings.ReplaceAll(t.Content, "  ", " ")
	}

	slice := []rune(t.Content)
	result := []rune{}
	for i := 0; i < len(slice); i++ {
		switch slice[i] {
		case '-':
			if i > 0 && i < len(slice)-1 {
				LastChar := result[len(result)-1]
				result = result[:len(result)-1]
				result = append(result, slice[i+1], LastChar)
				i++
			}
		default:
			result = append(result, slice[i])
		}
	}
	t.Content = string(result)
	t.Content = strings.ReplaceAll(t.Content, "+", "!")
	var sum int
	var newcontent []rune
	for _, r := range t.Content {
		if unicode.IsDigit(r) {
			sum += int(r - '0')

		} else {
			newcontent = append(newcontent, r)
		}
	}
	t.Content = string(newcontent)
	if sum > 0 {
		t.Content = fmt.Sprintf("%s%d", t.Content, sum)
	}
}
func main() {
	text := &Text{}
	// Создаем новый сканер для чтения  из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// просим пользователя ввести строку
	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
		fmt.Println("Результат:", text.Content)

	}
}
