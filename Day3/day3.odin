package main

import "core:strings"
import "core:os"
import "core:fmt"

check_all_possible_slopes :: proc(grid: [dynamic]string, y: int) -> int {
    trees: int = 0
    x: int = 0
    // Top left
    if grid[y][0] == '#' {
        trees += 1
    }
    // first (part 1) slope
    x = (x + 3) % len(grid[y])
    if grid[y][x] == '#' {
        trees += 1
    }

    return trees
}

main :: proc() {
    lines := read_by_lines("input.txt")
    defer delete(lines)

    // Start traversing
    // Initial values
    trees: int = 0
    x: int = 0
    y: int = 0
    max_y: int = len(lines) - 1

    // loop through rows
    for i := 0; i < len(lines); i+=1 {
        // Check top left if tree
        /* if lines[i][x] == '#' { */
        /*     trees += 1 */
        /*     continue; */
        /* } */

        // check Part 1 slope
        /* x = (x + 3) % len(lines[i]) */
        trees += check_all_possible_slopes(lines, y)
    }

    fmt.println("Trees:", trees)
}

read_by_lines :: proc(filepath: string) -> [dynamic]string {
    data, err := os.read_entire_file(filepath, context.allocator)
    if err != nil {
        fmt.eprintln("Could not read file", filepath)
        return nil
    }

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
