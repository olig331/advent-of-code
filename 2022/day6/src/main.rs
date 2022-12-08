use std::fs;

fn main() {
    let file = fs::read_to_string("input.txt").expect("");
    let bytes = file.as_bytes();
    let mut part_1_result = 0;
    let mut part_2_result = 0;

    for (i, x) in bytes.windows(4).enumerate() {
        if unique_chars(x) {
            part_1_result = i + 4;
            break;
        }
    }

    for (i, x) in bytes.windows(14).enumerate() {
        if unique_chars(x) {
            part_2_result = i + 14;
            break;
        }
    }

    println!("Part 1: {}", part_1_result);
    println!("Part 2: {}", part_2_result);
}

fn unique_chars(mut bytes: &[u8]) -> bool {
    while let Some((x, rest)) = bytes.split_first() {
        if rest.contains(x) {
            return false
        }
        bytes = rest;
    }
    return true;
}