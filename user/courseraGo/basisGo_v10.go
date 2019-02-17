package main

import "fmt"

func main() {
	fmt.Println("Ready")
	if true {
		fmt.Println("True val")
	}

	mapVal := map[string]string{"name": "John"}
	if keyVal, keyExist := mapVal["name"]; keyExist {
		fmt.Println(keyVal)
	}

	cond := 1
	if cond == 1 {
		fmt.Println("cond is ", cond)
	} else if cond == 2 {
		fmt.Println("cond is ", cond)

	} else {
		fmt.Println("Nahuy idid")
	}

	mapVal["firstName"] = "Evgen"
	mapVal["lastName"] = "Vlasov"

Loop:
	for key, val := range mapVal {
		fmt.Println("switch in loop ", key, val)
		switch {
		case key == "lastName":
			break
		case key == "firstName" && val == "Evgen":
			fmt.Println("switch - break loop here")
			break Loop
		}

	}

	str := "Привет, Мирок, епте!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
