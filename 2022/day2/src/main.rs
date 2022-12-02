use std::{fs, ops::Index};

// a = Rock
// b = Paper 
// c = scissors
// x = 1  Rock
// y = 2  Paper
// z = 3  Scissors

fn main() {
    let file = fs::read_to_string("input.txt").expect("ok");        
    let first_get_results = GameResults {
        x: Part1Wdl { a: 4, b: 1, c: 7 },
        y: Part1Wdl { a: 8, b: 5, c: 2 },
        z: Part1Wdl { a: 3, b: 9, c: 6 },        
    };    

    let second_get_results = Strategy {
        a: Part2Wdl { w: 8, d: 4, l: 3 },
        b: Part2Wdl { w: 9, d: 5, l: 1 },
        c: Part2Wdl { w: 7, d: 6, l: 2 }
    };

    let mut part_1_tally = 0;
    let mut part_2_tally = 0;
    file
        .split("\n")
        .for_each(|line| {
            let m = line.as_bytes()[2].to_ascii_lowercase() as char;
            let o = line.as_bytes()[0].to_ascii_lowercase() as char;

            let mut to_do = "l";
            if m.to_string() == "y" {
                to_do = "d";
            }
            if m.to_string() == "z" {
                to_do = "w";
            }

            part_1_tally += first_get_results[&m.to_string()][&o.to_string()];
            part_2_tally += second_get_results[&o.to_string()][to_do];
        });

    println!("Part 1: {:?}", part_1_tally);
    println!("Part 2: {:?}", part_2_tally);
        
}

impl Index<&'_ str> for Part1Wdl {
    type Output = i32;
    fn index(&self, s: &str) -> &i32 {
        match s {
            "a" => &self.a,
            "b" => &self.b,
            "c" => &self.c,
            _ => panic!("unknown field: {}", s),
        }
    }
}

impl Index<&'_ str> for GameResults {
    type Output = Part1Wdl;
    fn index(&self, s: &str) -> &Part1Wdl {
        match s {
            "x" => &self.x,
            "y" => &self.y,
            "z" => &self.z,
            _ => panic!("unknown field: {}", s),
        }
    }
}

impl Index<&'_ str> for Part2Wdl {
    type Output = i32;
    fn index(&self, s: &str) -> &i32 {
        match s {
            "w" => &self.w,
            "d" => &self.d,
            "l" => &self.l,
            _ => panic!("unknown field: {}", s),
        }
    }
}

impl Index<&'_ str> for Strategy {
    type Output = Part2Wdl;
    fn index(&self, s: &str) -> &Part2Wdl {
        match s {
            "a" => &self.a,
            "b" => &self.b,
            "c" => &self.c,
            _ => panic!("unknown field: {}", s),
        }
    }
}

struct GameResults {
    x: Part1Wdl,
    y: Part1Wdl,
    z: Part1Wdl,
}

struct Part1Wdl {
    a: i32,
    b: i32,
    c: i32,
}

struct Strategy {
    a: Part2Wdl,
    b: Part2Wdl,
    c: Part2Wdl
}

struct Part2Wdl {
    w: i32,
    d: i32,
    l: i32,
}
