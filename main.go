package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	//startWindow()
	printMenu()

	for {
		menu := getSelectedMenu()
		switch menu {
		case 1:
			fmt.Println("Monitoring...")
		case 2:
			fmt.Println("Helping...")
		case 3:
			fmt.Println(getVersion())
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Error!")
			os.Exit(-1)
		}
	}
}

/*func startWindow() {

	fmt.Println("╔══════════════════════════════════════════════════════════════╗") //ALT+201, ALT+205, ALT+187
	i := 1
	for i <= 3 {
		fmt.Println("║                                                              ║") //ALT+186
		i++
	}
	fmt.Println("╚══════════════════════════════════════════════════════════════╝") //ALT+200, ALT+205, ALT+188
}*/

func printMenu() {
	fmt.Println()
	fmt.Println("[1] Monitor")
	fmt.Println("[2] Version")
	fmt.Println("[3] Help")
	fmt.Println()
	fmt.Println("[0] Exit")
	fmt.Println()

}

func getSelectedMenu() int {
	var selectedMenu int
	fmt.Scan(&selectedMenu)
	fmt.Println("Menu selecionado:", selectedMenu)
	fmt.Println()

	return selectedMenu
}

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if ok {
		return "Version: " + info.Main.Version
	} else {
		return "Could not retrieve version information"
	}
}
