package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strconv"
    "math"

)

func main() {
file, err := os.Open("input.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
arr := []int{}
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    arr = append(arr,fuelCalculator(scanner.Text()))
 }
fmt.Println(totalCalculator(arr))
}


func fuelCalculator(input string) int {
    inputf, err := strconv.ParseFloat(input, 64)
    if err != nil{
        log.Println(err)
    }
    inputf = (inputf/3)
    inputf = math.Floor(inputf)
    inputf = inputf -2
    return int(inputf)
}

func totalCalculator(arr []int) int{
    var total int
    for _ , i := range arr {
        total += i
    }
    return total
}
