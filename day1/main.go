package main 

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

type calorie_T3 struct {
    calorie_H1, calorie_H2, calorie_H3 uint64
}


func push_calorie(calorie_value uint64, top_calories *calorie_T3) {
    if calorie_value > top_calories.calorie_H1 {
        top_calories.calorie_H3 = top_calories.calorie_H2
        top_calories.calorie_H2 = top_calories.calorie_H1
        top_calories.calorie_H1 = calorie_value
    } else {
        if calorie_value > top_calories.calorie_H2 {
            top_calories.calorie_H3 = top_calories.calorie_H2
            top_calories.calorie_H2 = calorie_value
        } else {
            if calorie_value > top_calories.calorie_H3 {
                top_calories.calorie_H3 = calorie_value
            }
        }
    }
}

func main() {
    // variable inventory
    var total_calorie, high_calorie uint64 
    //var elves_count uint

    //pointer to struct for part2 of the problem
    top_calories := new(calorie_T3)

    // opening a file
    file, err := os.Open("input")

    // if error exists
    if err != nil { 
        fmt.Println(err)
    }

    // closes file once all the instruction in main function executed completely
    defer file.Close()

    // new scanner
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        // adding calorie values for each elves
        if scanner.Text() != "" {
            value, err := strconv.ParseUint(scanner.Text(), 10, 32)
            if err != nil {
                fmt.Println("conversion error : ", err)
            }
            total_calorie += value
        } else {
            //added for part 2 of the problem
            push_calorie(total_calorie, top_calories)
            // changing high_calorie values
            if total_calorie > high_calorie {
                high_calorie = total_calorie
            }

            // resetting total_calorie values
            total_calorie = 0
        }

    } 
    
    // displaying overall result
    fmt.Println("High calorie value : ", high_calorie)

    // total of top3 calorie values
    calorie_result_T3 := top_calories.calorie_H1 + top_calories.calorie_H2 + top_calories.calorie_H3

    // all top 3 calorie values
    fmt.Printf("%d\n%d\n%d\n", top_calories.calorie_H1, top_calories.calorie_H2, top_calories.calorie_H3)
    // part 2 result
    fmt.Println("total of top3 calories : ", calorie_result_T3)
}

