package main

import "core:strings"
import "core:os"
import "core:fmt"

main :: proc() {
    groups: [dynamic]string
    defer delete(groups)
    GetGroups("input.txt", &groups)

    for group in groups {
        fmt.println(group)
    }

    count := 0
    for group in groups {
        numPeople := strings.count(group, "\n") + 1
        count += AnswersByAllInGroup(numPeople, group)
    }
    fmt.println("Unique answers:", count)
}

AnswersByAllInGroup :: proc(people: int, text: string) -> int {
    answers:= make(map[rune]int)
    defer delete(answers)

    for char in text {
        if char >= 'a' && char <= 'z' {
            answers[char] += 1
        }
    }

    count := 0
    for _, times in answers {
        if times == people {
            count += 1
        }
    }

    return count
}

CountUniqueAnswers :: proc(text: string) -> int {
    answers := make(map[rune]bool)
    defer delete(answers)

    for char in text {
        if char >= 'a' && char <= 'z' {
            answers[char] = true
        }
    }

    return len(answers)
}

GetGroups :: proc(filepath: string, array: ^[dynamic]string) {
    if filepath == "" {
        fmt.println("Error, filename can't be empty:", filepath)
        return 
    }

    data, err := os.read_entire_file(filepath, context.temp_allocator)
    if err != nil {
        fmt.println("Error reading file:", filepath, err)
        return
    }

    // groups_arr: [dynamic]string 
    it := string(data)

    for group in strings.split_iterator(&it, "\n\n") {
        append(array, group)
    }
}
