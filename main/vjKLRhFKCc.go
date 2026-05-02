package main
import "fmt"
func main() {
n := 10
fib := make([]int, n)
fib[0] = 0
if n > 1 {
fib[1] = 1
for i := 2; i < n; i++ {
fib[i] = fib[i-1] + fib[i-2]
}
}
for _, v := range fib {
fmt.Println(v)
}
}
func getFibonacci(n int) []int {
fib := make([]int, n)
fib[0] = 0
if n > 1 {
fib[1] = 1
for i := 2; i < n; i++ {
fib[i] = fib[i-1] + fib[i-2]
}
}
return fib
}
func printFibonacci(n int) {
fib := getFibonacci(n)
for _, v := range fib {
fmt.Println(v)
}
}
func isFibonacci(num int) bool {
a, b := 0, 1
for b < num {
a, b = b, a+b
}
return b == num
}
func main() {
printFibonacci(10)
fmt.Println(isFibonacci(21))
fmt.Println(isFibonacci(22))
}
