package main

import "core:strings"
import "core:os"
import "core:fmt"

main :: proc() {
    fmt.println("Hello there!")

    groups: [dynamic]string = GetGroups("example.txt")
    defer delete(groups)

    count := 0

    for i, group in groups {
        fmt.print("--- Group ", i, "---\n", group, "\n\n")
    }
}

GetGroups ::proc(filepath: string) -> [dynamic]string {
    if filepath == "" {
        fmt.println("Error, filename can't be empty:", filepath)
        return nil
    }

    data, err := os.read_entire_file(filepath, context.temp_allocator)
    if err != nil {
        fmt.println("Error reading file:", filepath, err)
        return nil
    }

    groups_arr: [dynamic]string 
    it := string(data)

    for group in strings.split_iterator(&it, "\n\n") {
        append(&groups_arr, group)
    }

    return groups_arr
}