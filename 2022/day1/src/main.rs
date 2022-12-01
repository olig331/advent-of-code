use std::fs;

fn main() {
    let file = fs::read_to_string("input.txt").expect("");
    let mut values:Vec<u64> = file
        .split("\n\n")
        .map(|elf_calories| {
            elf_calories.split("\n")
            .map(|items| items.parse::<u64>().unwrap()).sum()
        }).collect::<Vec<u64>>();
    values.sort();

    
    let last3 = values.as_slice()[values.len()-3..].to_vec();
    let last = values.last();

    println!("{:?}", last);
    println!("{:?}", last3.iter().sum::<u64>())
}
