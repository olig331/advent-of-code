use std::fs;

fn main() {

   let file = fs::read_to_string("input.txt").expect("ok");
   let mut result = 0;
   file.split("\n").for_each(|line| {
        let length = line.chars().count() / 2;
        let first = &line[..length];
        let second = &line[length..];
        first.chars().any(|c| {
            if second.contains(c) == true {
                if c.is_lowercase() {
                    result += c as u32 - 96;
                } else {
                    result += c as u32 - 38;
                }
                return true        
            }            
            return false
        });        
   });
   println!("Part 1: {:?}", result);
   part2(file);
}


fn part2(file:String){
    let mut result = 0;
    let lines:Vec<&str> = file.split("\n").collect();
    for i in 0..lines.len() {
        if i == 0 || i % 3 == 0 {            
            lines[i].chars().any(|c| {
                if lines[i + 1].contains(c) == true { 
                    if lines[i + 2].contains(c) == true {
                        if c.is_lowercase() {
                            result += c as u32 - 96;
                        } else {
                            result += c as u32 - 38;
                        }
                        return true
                    }                   
                    return false
                } 
                return false;
            });        
        }
    }
    println!("Part 2: {:?}", result)
}