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
    test := fuelCalculator(scanner.Text())
    temp1 := test
    var temp int
    for (fuelCalculator(strconv.Itoa(test)) >= 0){
        test = fuelCalculator(strconv.Itoa(test))
        temp += test
    }
    test = temp1 + temp
    arr = append(arr, test)
}
total := totalCalculator(arr)
fmt.Println(total)
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
        //fmt.Println(total)
    }
    return total
}
