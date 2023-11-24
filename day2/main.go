package main

/*
A , X -> Rock       -> 1
B , Y -> Paper      -> 2
C , Z -> Scissor    -> 3

0 -> Lose
3 -> Draw
6 -> Win
*/

import (
    "fmt"
    "os"
    "bufio"
)

//function to convert ascii value into normal options
func convert_ascii(input uint) uint{
    if input == 65 || input == 88 {
        return 1
    } else if input == 66 || input == 89 {
        return 2
    } else {
        return 3
    }
}

/* partone 
//function to get output of round based on opponents and players player_move
func get_round_output(input_1, input_2 uint) uint{
    if input_1 != input_2 {
        if input_2 > input_1 {
            if input_2 == 3 && input_1 == 1 {
                return 0
            }
            return 6
        } else {
            if input_2 ==1 && input_1 == 3 {
                return 6
            }
        }
        return 0
    } else {
        return 3
    }
}
*/

func get_player_move(opponent_move, desired_output uint) (uint, uint){
    if desired_output == 1 {
        if opponent_move == 1 {
            return 3, 0
        }
        return (opponent_move - 1), 0
    } else if desired_output == 2 {
        return opponent_move, 3
    } else {
        if opponent_move == 3 {
            return 1, 6
        }
        return (opponent_move + 1), 6
    }
}

func main() {
    //variable to store result accumulation of all scores of each round
    var result uint
    //opening input file
    file, err := os.Open("input")
    //if error arises, display error message
    if err != nil {
        fmt.Println(err)
    }
    //deferenced function to close the input file
    defer file.Close()
    //new scanner to read the file
    scanner := bufio.NewScanner(file)
    //reading input line by line
    for scanner.Scan() {
        opponent_move := convert_ascii(uint(scanner.Text()[0]))
        //used in part one
        //player_move := convert_ascii(uint(scanner.Text()[2]))

        desired_output := convert_ascii(uint(scanner.Text()[2]))

        //function used for part one
        //round_output := get_round_output(opponent_move, player_move)
        player_expected_move, output_value := get_player_move(opponent_move, desired_output)
        round_result := player_expected_move + output_value

        //fmt.Printf("%d %d %d\n", opponent_move, player_move, round_result)
        result += round_result
    }
    fmt.Println("result -> ", result)
}
