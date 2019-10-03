package utils

import (
	"fmt"
)

func LogHeading(shown string, hidden string) {
	hidden = "URL"
	fmt.Println("\n\n#==================================================#")
	fmt.Printf("# [%s] %s\n", hidden, shown)
	fmt.Println("#==================================================#")
}

func Log(shown string, hidden string) {
	hidden = "URL"
	fmt.Printf("# [%s] %s\n", hidden, shown)
}

func LogTree(shown string, hidden string, indent int) {
	hidden = "URL"

	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("\\_ [%s] %s\n", hidden, shown)

}
