package main

import "core:strconv"
import "core:strings"
import "core:os"
import "core:log"
import "core:fmt"

Instruction :: struct {
    inst: string,
    num: int,
}

main :: proc() {
    lines: [dynamic]string = get_lines("example.txt")
    defer delete(lines)

    for line in lines {
        current := parse_line_into_instruction(line)
        fmt.println(current)
    }
}

parse_line_into_instruction :: proc(line: string) -> Instruction {
    splitted, err := strings.split(line, " ", context.temp_allocator)
    if err != nil {
        log.fatal("Could not split:", err)
    }

    value, ok := strconv.parse_int(splitted[1])
    if ok != true {
        log.fatal("Could not convert str to int:", err)
    }
    return Instruction{splitted[0], value}
}

get_lines :: proc(filepath: string) -> [dynamic]string {
    if filepath == "" {
        log.fatal("Filepath can't be empty")
    }
    data, err := os.read_entire_file(filepath, context.temp_allocator)
    if err != nil {
        log.fatalf("Error reading file %q: %d\n", filepath, err)
    }
    result := [dynamic]string{}
    it := string(data)
    for line in strings.split_lines_iterator(&it) {
        append(&result, line)
    }
    return result
}
