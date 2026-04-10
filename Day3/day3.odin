package main

import "core:strings"
import "core:os"
import "core:fmt"

main :: proc() {
    lines := read_by_lines("input.txt")

    /* for line in lines { */
    /*     fmt.println(line) */
    /* } */

    // Start traversing
    // Initial values
    trees: int = 0
    x: int = 0
    y: int = 0

    // loop through rows
    for i := 0; i < len(lines); i+=1 {
        fmt.print(lines[i])
        fmt.print(" length:", len(lines[i]))
        x = (x + 3) % len(lines[i])
        fmt.println("col:", x)
    }
}

read_by_lines :: proc(filepath: string) -> [dynamic]string {
    data, err := os.read_entire_file(filepath, context.temp_allocator)
    if err != nil {
        fmt.eprintln("Could not read file", filepath)
        return nil
    }
    /* defer delete(data, context.allocator) */

    result: [dynamic]string
    it := string(data)
    /* if strings.has_prefix(it, "\xef\xbb\xbf") { */
    /*     it = it[3:] */
    /* } */
    for line in strings.split_lines_iterator(&it) {
        append(&result, line)
    }
    
    return result
}
