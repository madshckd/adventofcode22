package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//final check
func contain_check_confirmation(containee, container uint) bool {
    if containee <= container {
        return true 
    }
    return false
}

//first check investigating the possibility, to know about which part is container
func possibility_to_contain(range_1, range_2 uint) uint {
    if range_1 < range_2 {
        return 1
    } else if range_2 < range_1 {
        return 2 
    } else {
        return 0
    }
}

//function to convert string to uint
func convert_U(string_input string) uint {
    parsed_Uint, err := strconv.ParseUint(string_input, 10, 32)

    if err != nil {
        fmt.Println(err)
        return 0
    }
    return uint(parsed_Uint)
}

func check_overlap(pair_one, pair_two []string) bool{
    start_F, start_S := convert_U(pair_one[0]), convert_U(pair_two[0])
    end_F, end_S := convert_U(pair_one[1]), convert_U(pair_two[1])

    if start_F < start_S {
        if start_S <= end_F {
            if end_S < end_F || end_S >= end_F {
                return true
            }
        }
        return false
    } else if start_S < start_F {
        if start_F <= end_S {
            if end_F < end_S || end_F >= end_S {
                return true
            }
        }
        return false
    } else {
        if end_F < end_S || end_F >= end_S {
            return true
        }
        return false
    }

}

//function to further breakdowm to find if it one pair fully overlappin the other
func find_range(part_one, part_two string) (bool, bool) {
    part_one_range := strings.Split(part_one, "-")
    part_two_range := strings.Split(part_two, "-")

    check_F := possibility_to_contain(convert_U(part_one_range[0]), convert_U(part_two_range[0]))

    if check_F == 1 {
        return contain_check_confirmation(convert_U(part_two_range[1]), 
        convert_U(part_one_range[1])), check_overlap(part_one_range,
        part_two_range)
    } else if check_F == 2 {
        return contain_check_confirmation(convert_U(part_one_range[1]), 
        convert_U(part_two_range[1])), check_overlap(part_one_range, 
        part_two_range)
    } else {
        if convert_U(part_one_range[1]) < convert_U(part_two_range[1]) {
            return contain_check_confirmation(convert_U(part_one_range[1]), 
            convert_U(part_two_range[1])), 
            check_overlap(part_one_range, part_two_range)
        } else { 
            return contain_check_confirmation(convert_U(part_two_range[1]), 
            convert_U(part_one_range[1])), 
            check_overlap(part_one_range, part_two_range)
        }
    }
}

//function to split the pair (comma)
func split_C(input_pair string) (bool, bool) {
    splitted_pair := strings.Split(input_pair, ",")
    return find_range(splitted_pair[0], splitted_pair[1])
}

func main() {
    var fully_overlapped, overlapped uint
    //opening file
    file, err := os.Open("input")
    //if error print error message
    if err != nil {
        fmt.Println(err)
    }
    //deferenced function to close file
    defer file.Close()

    //scanner
    scanner := bufio.NewScanner(file)

    //reading line by line
    for scanner.Scan() {
        result_F, result_O := split_C(scanner.Text())

        //incrementing if fully overlapped
        if result_F {
            fully_overlapped++
        }
        if result_O {
            overlapped++
        }

    }

    //part one : result
    fmt.Println("Fully Overlapping pairs => ", fully_overlapped)
    fmt.Println("Overlapped pairs => ", overlapped)
}
