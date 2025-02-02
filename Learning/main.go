package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill:", reader)

	b := newBill(name)
	fmt.Println("Created the Bill -", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose a option: (a - add an item, s - save bill, t - add tip):", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("Item price:", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item Added:", name, price)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)
	case "t":
		tip, _ := getInput("Enter Tip Amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip Added:", tip)
		promptOptions(b)

	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
