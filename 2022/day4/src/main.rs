use std::fs;

fn main() {
    let mut part_1_count = 0;
    let mut part_2_count = 0;
    let file = fs::read_to_string("input.txt").expect("ok");
    file.split("\n").for_each(|line|{
      let strings:Vec<&str> = line.split(|c:char| c.to_string() == "," || c.to_string() == "-").collect();
      let ints:Vec<i32> = strings.iter().map(|s| s.parse::<i32>().unwrap()).collect();

      part_1_count += fully_overlaps(&ints);
      part_2_count += partial_overlap(&ints);

    });
    println!("Part 1: {}", part_1_count);
    println!("Part 2: {}", part_2_count)
}

fn fully_overlaps(ints: &Vec<i32>) -> i32 {
  if ints[0] >= ints[2] && ints[1] <= ints[3] {
    return 1;
  } else if ints[2] >= ints[0] && ints[3] <= ints[1] {
    return 1;
  }
  return 0
}

fn partial_overlap(ints: &Vec<i32>) -> i32 {
  if ints[0] <= ints[2] && ints[2] <= ints[1] {
    return 1
  }
  if ints[2] <= ints[0] && ints[0] <= ints[3] {
    return 1
  }  
  return 0
}