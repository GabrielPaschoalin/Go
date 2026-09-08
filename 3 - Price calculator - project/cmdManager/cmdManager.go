package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}
func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every prices with ENTER")

	var stringArray []string

	for {
		var value string
		fmt.Print("Price: ")
		fmt.Scan(&value)

		if value == "0" {
			break
		}

		stringArray = append(stringArray, value)
	}

	return stringArray, nil

}

func (cmd CMDManager) WriteResult(data interface{}) error {

	fmt.Println(data)
	return nil

}
