package main

import (
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v2"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {

	fmt.Println()
	fmt.Println("[======database driver======]")
	fmt.Println()

	promptSelect := promptui.Select{
		Label: "choose a driver",
		Items: []string{"mysql", "mssql", "postgres", "sqlite"},
	}

	_, _, err := promptSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("[======database config======]")
	fmt.Println()

	prompt("address ")
	prompt("username ")
	promptPassword("password ")

	fmt.Println()
	fmt.Println("[===========tables===========]")
	fmt.Println()

	tables := []string{"post", "users", "article", "audio", "[finish]"}
	var value string
	for tables[0] != "[finish]" && value != "[finish]" {
		value = selects(tables)
		tables = removeItem(tables, value)
	}

	fmt.Println()
	fmt.Println("[=========generating=========]")
	fmt.Println()

	bar := progressbar.New(100)
	for i := 0; i < 100; i++ {
		_ = bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println()
}

func prompt(label string) {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Username must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func promptPassword(label string) {
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Username must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}

	_, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func selects(tables []string) string {
	promptSelect := promptui.Select{
		Label: "choose table to generate",
		Items: tables,
	}

	_, result, err := promptSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func removeItem(tables []string, table string) []string {
	index := 0
	for i := 0; i < len(tables); i++ {
		if tables[i] == table {
			index = i
		}
	}
	return append(tables[:index], tables[index+1:]...)
}
