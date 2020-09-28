package main

import "fmt"

func main()  {
	var ( 
		limit int;
		total float64 = 1.0;
		factorial int;
	)

	fmt.Scan(&limit);
	factorial = 1;

	for i:=1; i<=limit; i++ {
		factorial = factorial * i;
		total = total + 1/float64(factorial);
	}
	fmt.Println(total);
}