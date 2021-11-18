package main

import (
	"fmt"

	"company/engineer"
)

func main() {
	// something like bruce = new Employee(Bruce, ...) in Java, JS
	bruce := engineer.NewEngineer("Bruce", "Lee")
	fmt.Println(bruce)

	Tony := engineer.NewEngineerWithOptions("Tony", "Ja", engineer.Option{
		Salary:   100,
		Position: "Backend",
	})
	fmt.Println(Tony)

	bSalary := 300.0
	bPosition := "midfield"
	Beckham := engineer.NewEngineerWithOptionPointers("David", "Beckham", engineer.OptionPointers{
		&bSalary, &bPosition,
	})
	fmt.Println(Beckham)

	Messi := engineer.NewEngineerWithOptFn("Messi", "Lionel", engineer.WithPosition("Attacker"), engineer.WithSalary(200000.0))
	fmt.Println(Messi)
}
