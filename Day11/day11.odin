package main

import "core:strings"
import "core:log"
import "core:os"
import "core:fmt"

main :: proc() {
    rows := get_lines("example.txt")
    defer delete(rows)

    /* for row in rows { */
    /*     fmt.println(row) */
    /* } */
    // Going slowly
    min_y := 0
    min_x := 0
    max_y := len(rows) - 1
    max_x := len(rows[0]) - 1
    fmt.printfln("Max x: %d. Max y: %d", max_x, max_y)

    // Grid traversal
    for row := 0; row < len(rows); row += 1 {
        for c := 0; c < len(rows[0]); c += 1 {
            // if empty
            if is_empty(rows[row][c]) {
                // if no adjacent
            }
        }
    }
}

no_adjacent :: proc(c: rune, grid: []int) -> bool {
    // top
    // bottom
    // left
    // right
    // top left
    // bottom left
    // top right
    // bottom right
}

check_if_taken :: proc(c: rune) -> bool {
    if c == '#' {
        return true
    }
    return false
}

is_empty :: proc(c: u8) -> bool {
    if c == 'L' {
        return true
    }
    return false
}

is_floor :: proc(c: rune) -> bool {
    if c == '.' {
        return true
    }
    return false
}

get_lines :: proc(filename: string) -> [dynamic]string {
    // Skipping checking if file opens correctly and all that
    data, err := os.read_entire_file(filename, context.temp_allocator)
    if err != nil {
        log.fatal("Could not open file:", err)
    }

    it := string(data)

    result := [dynamic]string{}
    for line in strings.split_lines_iterator(&it) {
        append(&result, line)
    }
    return result
}
