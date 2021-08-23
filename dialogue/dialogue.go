package dialogue

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Speaker interface {
	IsCalled() string
	IsPlayable() bool
	IsDead() bool
}

func Ask(s Speaker, question string, handle func(answer string) error) error {
	var answer string
	if s.IsPlayable() {
		color.Cyan("%-8s %q", "?", question)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return err
			}
			if scanner.Text() != "\n" {
				return handle(scanner.Text())
			}
		}
		return handle(answer)
	}

	return nil
}

func Say(s Speaker, words string) {
	defer color.Unset()
	msgFmt := "%-8s "
	prefix := fmt.Sprintf(msgFmt, s.IsCalled())
	if s.IsPlayable() {
		prefix = color.CyanString(prefix)
	}

	if s.IsDead() {
		color.Set(color.FgRed)
		prefix = prefix + "[DEAD] "
	}

	fmt.Println(prefix + words)
}
