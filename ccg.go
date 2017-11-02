package main

import "fmt"
import "os"
import "strconv"

func computeCollatz(n int) []int {
    var steps []int 
    steps = append(steps, n)
    for n != 1 && n != 0 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = 3*n + 1
        }
		steps = append(steps, n)
    }
   	return steps
}


func main() {
    args := os.Args
    if len(args) < 2 {
        panic("You need to provide a number ./ccg <number>")
    }
    num, err := strconv.Atoi(args[1])
    if(err != nil) {
        panic(err)
    }
    output := computeCollatz(num)
    fmt.Print(output)
}
