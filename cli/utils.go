package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// inputScanner reads input from the terminal and stores it in the provided interface.
// Type exit to break the scanning process
func inputScanner(input any, text string) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(text)
		scanner.Scan()
		in := scanner.Text()
		if in == "exit" {
			break
		}

		switch v := input.(type) {
		case *string:
			*v = in
			return nil
		case *float64:
			f, err := strconv.ParseFloat(in, 64)
			if err != nil {
				fmt.Println("Input invalid. Silahkan coba lagi.")
				fmt.Println()
				continue
			}
			*v = f
			return nil
		case *int64:
			i, err := strconv.ParseInt(in, 10, 64)
			if err != nil {
				fmt.Println("Input invalid. Silahkan coba lagi.")
				fmt.Println()
				continue
			}
			*v = i
			return nil
		case *int:
			i, err := strconv.Atoi(in)
			if err != nil {
				fmt.Println("Input invalid. Silahkan coba lagi.")
				fmt.Println()
				continue
			}
			*v = i
			return nil
		case *bool:
			b, err := strconv.ParseBool(in)
			if err != nil {
				fmt.Println("Input invalid. Silahkan coba lagi.")
				fmt.Println()
				continue
			}
			*v = b
			return nil
		default:
			return fmt.Errorf("unsupported data type: %T", input)
		}
	}

	return fmt.Errorf("exit trigger")
}
