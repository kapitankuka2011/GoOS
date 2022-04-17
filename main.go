package main

import (
	"fmt"
	"os/exec"
	"os"
	"runtime"
	"time"
)

func print(msg string) {
	fmt.Printf(msg + "\n")	
}

func err(msg string) {
	print("Error>: " + msg)
}

func execute(cmd string) {
	c := exec.Command(cmd)
	c.Stdout = os.Stdout
	c.Run()
}

func clear() {
	if runtime.GOOS == "windows" {
		execute("cls")
	} else {
		execute("clear")
	}
}

func startos(authorized int, password string) {
	if authorized != 1 {
		err("Startup declined.")
	} else {
		if password != "authorized" {
			err("Startup declined.")
		} else {
			clear()
			var option string
			print("1. Notepad")
			print("2. Calculator")
			print("3. Halt")
			fmt.Printf(">")
			fmt.Scanln(&option)
			if option == "1" {
				clear()
				print("Notepad 1.0\nUse cls to clear screen.\nUse exit to exit.\n")
				for {
					var pad string
					fmt.Scanln(&pad)
					if pad == "cls" {
						clear()
					} else if pad == "exit" {
						clear()
						startos(1, "authorized")
					}
				}
			} else if option == " " || option == "" {
				startos(1, "authorized")
			} else if option == "3" {
				clear()
				print("Halting...")
				time.Sleep(2 * time.Second)
				os.Exit(0)
			} else if option == "2" {
				clear()
				print("Calculator 1.0\n\n")
				var num1 float64
				var num2 float64
				var op string
				fmt.Printf("Enter first number: ")
				fmt.Scanln(&num1)
				fmt.Printf("Enter second number: ")
				fmt.Scanln(&num2)
				fmt.Printf("Enter operation: ")
				fmt.Scanln(&op)
				if op == "+" {
					fmt.Printf("%f + %f = %f\n", num1, num2, num1 + num2)
				} else if op == "-" {
					fmt.Printf("%f - %f = %f\n", num1, num2, num1 - num2)
				} else if op == "*" {
					fmt.Printf("%f * %f = %f\n", num1, num2, num1 * num2)
				} else if op == "/" {
					fmt.Printf("%f / %f = %f\n", num1, num2, num1 / num2)
				}
				time.Sleep(4 * time.Second)
				startos(1, "authorized")
			} else {
				print("Invalid option!")
				time.Sleep(1 * time.Second)
				startos(1, "authorized")
			}
		}
	}
}

func main() {
	fmt.Printf("Starting...")
	time.Sleep(2 * time.Second)
	print("Booted")
	time.Sleep(1 * time.Second)
	startos(1, "authorized")
}