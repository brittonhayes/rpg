package dialogue

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/goutils"
	"github.com/brittonhayes/rpg/color"
)

type Speaker interface {
	IsCalled() string
	IsPlayable() bool
	IsDead() bool
	HasColor() color.Color
}

func Ask(s Speaker, question string, handle func(answer string) error) error {

	// name, _ := goutils.Abbreviate(s.IsCalled(), 13)
	msg := message(s, question)

	fmt.Println(msg)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		if scanner.Text() != "\n" {
			return handle(scanner.Text())
		}
	}

	return nil
}

func Say(s Speaker, words string) {
	msg := message(s, words)

	if s.IsDead() {
		msg = message(s, "[DEAD] "+words)
	}

	fmt.Println(msg)
}

func message(s Speaker, content string) string {
	name, _ := goutils.Abbreviate(s.IsCalled(), 13)
	name = strings.ToUpper(name)

	return fmt.Sprintf("%-15s %25s", name, s.HasColor().Sprintf("%q", content))
}
