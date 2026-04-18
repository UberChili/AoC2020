package main

import "core:strings"
import "core:os"
import "core:log"
import "core:fmt"

main :: proc() {
    lines := get_lines("example.txt")
    defer delete(lines)

    for line in lines {
        fmt.println(line)
    }
}

get_lines :: proc(filepath: string) -> [dynamic]string {
    if filepath == "" {
        log.fatal("Can't open file. Filename can't be empty.")
    }

    data, err := os.read_entire_file(filepath, context.temp_allocator)
    if err != nil {
        fmt.eprintln("Error: Could not open file", filepath, err)
        os.exit(1)
    }
    /* defer delete(data, context.allocator) */

    it := string(data)
    result: [dynamic]string
    for line in strings.split_lines_iterator(&it) {
        append(&result, line)
    }

    return result
}
