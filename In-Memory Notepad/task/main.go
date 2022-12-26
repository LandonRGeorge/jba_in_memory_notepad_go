package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxNotes int

type notes []string

func (n *notes) create(s string) {
	if s == "" {
		fmt.Fprint(os.Stdout, "[Error] Missing note argument\n")
		return
	}
	if len(*n) >= maxNotes {
		fmt.Println("[Error] Notepad is full")
		return
	}
	*n = append(*n, s)
	fmt.Println("[OK] The note was successfully created")
}

func (n *notes) list() {
	_n := *n
	if len(_n) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}
	for i := range *n {
		fmt.Printf("[Info] %d: %s\n", i+1, _n[i])
	}
}

func (n *notes) clear() {
	*n = nil
	fmt.Println("[OK] All notes were successfully deleted")
}

func (n *notes) update(text string) {
	_n := *n
	if text == "" {
		fmt.Fprint(os.Stdout, "[Error] Missing position argument\n")
		return
	}

	split := strings.SplitN(text, " ", 2)

	if len(split) == 1 {
		fmt.Fprint(os.Stdout, "[Error] Missing note argument\n")
		return
	}

	idx, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Fprintf(os.Stdout, "[Error] Invalid position: %v\n", split[0])
		return
	}
	if idx < 0 || idx > maxNotes {
		fmt.Fprintf(os.Stdout, "[Error] Position %d is out of the boundary [1, %d]\n", idx, maxNotes)
		return
	}
	if idx > len(_n) {
		fmt.Fprint(os.Stdout, "[Error] There is nothing to update\n")
		return
	}
	_n[idx-1] = split[1]
	fmt.Printf("[OK] The note at position %d was successfully updated\n", idx)
}

func (n *notes) delete(text string) {
	_n := *n
	if text == "" {
		fmt.Fprint(os.Stdout, "[Error] Missing position argument\n")
		return
	}
	idx, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stdout, "[Error] Invalid position: %v\n", text)
		return
	}
	if idx < 0 || idx > maxNotes {
		fmt.Fprintf(os.Stdout, "[Error] Position %d is out of the boundary [1, %d]\n", idx, maxNotes)
		return
	}
	if idx > len(_n) {
		fmt.Fprint(os.Stdout, "[Error] There is nothing to delete\n")
		return
	}
	*n = append(_n[:idx-1], _n[idx:]...)
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", idx)
}

func parseInput(text string) (string, string) {
	texts := strings.SplitN(text, " ", 2)
	command := texts[0]
	if len(texts) == 1 {
		return command, ""
	}
	data := texts[1]
	return command, data
}

func main() {
	var n notes
	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&maxNotes)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command and data: ")
		scanner.Scan()
		text := scanner.Text()
		command, data := parseInput(text)
		switch command {
		case "create":
			n.create(data)
		case "list":
			n.list()
		case "clear":
			n.clear()
		case "exit":
			fmt.Println("[Info] bye!")
			os.Exit(0)
		case "update":
			n.update(data)
		case "delete":
			n.delete(data)
		default:
			fmt.Fprint(os.Stdout, "[Error] Unknown command\n")
			continue
		}
	}
}
