package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//constant containing stack counts
const (
    STACK_COUNT int = 9
)

//global variables stacks and stack pointer pointing last index of each stack
var stacks [STACK_COUNT][]string
var stack_P [STACK_COUNT] int

//initial stacks 
func init_stack() {
    stacks[0] = append(stacks[0], "N", "D", "M", "Q", "B", "P", "Z")
    stacks[1] = append(stacks[1], "C", "L", "Z", "Q", "M", "D", "H", "V")
    stacks[2] = append(stacks[2], "Q", "H", "R", "D", "V", "F", "Z", "G")
    stacks[3] = append(stacks[3], "H", "G", "D", "F", "N")
    stacks[4] = append(stacks[4], "N", "F", "Q")
    stacks[5] = append(stacks[5], "D", "Q", "V", "Z", "F", "B", "T")
    stacks[6] = append(stacks[6], "Q", "M", "T", "Z", "D", "V", "S", "H")
    stacks[7] = append(stacks[7], "M", "G", "F", "P", "N", "Q")
    stacks[8] = append(stacks[8], "B", "W", "R", "M")

    init_stack_length()
}

//initial length of stacks
func init_stack_length() {
    for i := 0; i< STACK_COUNT; i++ {
        stack_P[i] = (len(stacks[i])) - 1
    }
}

//function to convert string to uint
func to_Uint(number_string string) uint{
    number, err := strconv.ParseUint(number_string, 10, 32)

    //printing errors
    if err != nil {
        fmt.Println(err)
    }

    return uint(number)
}

//function to move elements linearly in LIFO manner
//one by one
func move_F(count_M, stack_F, stack_T uint) {
    var movement uint

    for movement < count_M {
        //element to move
        elem_T := (stack_P[stack_F - 1])
        //calling move_T to move element to the destined stack
        move_T(stacks[stack_F - 1][elem_T], stack_T)
        //calling move_A to process after-movement procedure
        move_A(stack_F, stack_T)
        movement++
    }

}

//function to move desired elements to stack mentioned by move_F 
func move_T(move_E string, stack_T uint) {
    stacks[stack_T - 1] = append(stacks[stack_T - 1], move_E)
}

//function organise stack after movement includes
/*
> removing element from stack_F (from stack)
> changing stack_P values to point to the new top index of stacks that is changed
*/
func move_A(stack_F, stack_T uint) {
    stack_P[stack_T - 1] = (len(stacks[stack_T - 1])) - 1
    stacks[stack_F - 1] = stacks[stack_F - 1][ : stack_P[stack_F - 1]]
    stack_P[stack_F - 1] = (len(stacks[stack_F - 1])) - 1
}

//function to split the input text to pass values like 
//number of elements to move, from which stack to which stack 
func split_input(input string){ 
    splitted_input := strings.Split(input, " ")
    //calling move_F function
    move_F(to_Uint(splitted_input[1]),
        to_Uint(splitted_input[3]),
        to_Uint(splitted_input[5]))
}

func main() {
    //initialising stacks
    init_stack()

    //opening file
    file, err := os.Open("input")

    //error in opening file
    if err != nil {
        fmt.Println(err)
    }

    //closing the file
    defer file.Close()

    //new scanner that read a file
    scanner := bufio.NewScanner(file)

    //passing input text
    for scanner.Scan() {
        split_input(scanner.Text())
    }

    //displaying output, top element in each stack
    for i := 0; i < STACK_COUNT; i++ {
        fmt.Printf("%s ", stacks[i][stack_P[i]])
    }

    fmt.Println()
}
