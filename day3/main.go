package main

import (
	"bufio"
	"fmt"
	"os"
)

//function to find common item for part one
func find_common_item(compartment_F, compartment_S string) uint{

    var common_item byte

    //iterating through two compartments linearly
    for i:=0; i<len(compartment_F); i++{
        for j:=0; j<len(compartment_S); j++{
            if compartment_F[i] == compartment_S[j] {
                common_item = compartment_F[i]
                break
            }
        }
    }

    //after calculating the priority value for common element, priority value is returned
    return get_priority_value(common_item)
}

//function to find priority value for common item that is passed
func get_priority_value(common_item byte) uint{

    var priority uint

    /*
    priority value
    a-z := 1 - 26
    A-Z := 27 - 52
    */
    if uint(common_item) >= 97 && uint(common_item) <= 172 {
        priority = uint(common_item) - 96
    } else if uint(common_item) >= 65 && uint(common_item) <= 132 {
        priority = uint(common_item) - 38
    } else {
        priority = 0
    }
    return priority
}

func part_one(input string) uint{

    //compartmentalising single line of input based on lenght of the line
    compartment_F := input[:(len(input)/2)]
    compartment_S := input[(len(input)/2) : ]

    //after finding common element between two compartments, priority value is returned
    return find_common_item(compartment_F, compartment_S)
}

func part_two(input []string) uint{
    
    var common_item byte

    //same thing as done in the part one except, now it has three input to find common item
    //then priority values will be found for that common item.
    for i:=0; i<len(input[0]); i++{
        for j:=0; j<len(input[1]); j++{
            if input[0][i] == input[1][j] {
                for k:=0; k<len(input[2]); k++ {
                    if input[0][i] == input[2][k] {
                        common_item = input[0][i]
                        break
                    }
                }
            }
        }
    }
    return get_priority_value(common_item)
}

func main() {
    //variables to store, priority value and slice for part two input storage
    var priority_total_one, priority_total_two uint
    var part_two_input []string

    //opening input file
    file, err := os.Open("input")

    //if error arises, printing the error 
    if err != nil {
        fmt.Println(err)
    }

    //deferenced function call to close the input file
    defer file.Close()

    //scanner created to read opened input file line by line
    scanner := bufio.NewScanner(file)

    //using the scanner created, reading file line by line
    for scanner.Scan() {

        //for part one, passing each line as an input
        priority_value_one := part_one(scanner.Text())
        priority_total_one += priority_value_one

        // for part two computation, appending or accepting only 3 lines of input as a group iteratively 
        part_two_input = append(part_two_input, scanner.Text())
        if len(part_two_input) == 3 {
            priority_value_two := part_two(part_two_input)
            priority_total_two += priority_value_two
            part_two_input = nil
        }
    }
    //displaying total priority values (result)
    fmt.Println("Priority Total (part_one) -> ", priority_total_one)
    fmt.Println("Priority Total (part_two) -> ", priority_total_two)
}
