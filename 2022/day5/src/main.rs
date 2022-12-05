use std::fs;

fn main() {
    let mut state:Vec<Vec<&str>> = vec![
        vec!["Z", "T", "F", "R", "W", "J", "G"],
        vec!["G", "W", "M"], 
        vec!["J", "N", "H", "G"], 
        vec!["J", "R", "C", "N", "W"], 
        vec!["W", "F", "S", "B", "G", "Q", "V", "M"], 
        vec!["S", "R", "T", "D", "V", "W", "C"], 
        vec!["H", "B", "N", "C", "D", "Z", "G", "V"], 
        vec!["S", "J", "N", "M", "G", "C"], 
        vec!["G", "P", "N", "W", "C", "J", "D", "L"]
    ];
    
    let mut result: String = "".to_owned();
    let file = fs::read_to_string("input.txt").expect("ok");
    file.split("\n")
        .for_each(|line| {
            let mut ints:Vec<_> = line.split(" ").filter_map(|a| a.parse::<i32>().ok()).collect();
            ints[1] = ints[1] - 1;
            ints[2] = ints[2] - 1;
            
            let mut i = 1;
            while i <= ints[0] {   
                if state[ints[1] as usize].len() > 0 {            
                    let to_move = state[ints[1] as usize].last().unwrap().to_owned();
                    state[ints[2] as usize].push(to_move);
                    state[ints[1] as usize].pop();
                }
                i += 1;
            };    
        });

    let size = state.len() - 1;
    for i in 0..size + 1 {
        result.push_str(state[i].last().unwrap());
    }
    println!("{}", result);
    println!("{}", part2(file));    
}

fn part2(file: String) -> String {
    let mut state:Vec<Vec<&str>> = vec![
        vec!["Z", "T", "F", "R", "W", "J", "G"],
        vec!["G", "W", "M"], 
        vec!["J", "N", "H", "G"], 
        vec!["J", "R", "C", "N", "W"], 
        vec!["W", "F", "S", "B", "G", "Q", "V", "M"], 
        vec!["S", "R", "T", "D", "V", "W", "C"], 
        vec!["H", "B", "N", "C", "D", "Z", "G", "V"], 
        vec!["S", "J", "N", "M", "G", "C"], 
        vec!["G", "P", "N", "W", "C", "J", "D", "L"]
    ];

    let mut result: String = "".to_owned();
    file.split("\n")
    .for_each(|line| {
        let mut ints:Vec<_> = line.split(" ").filter_map(|a| a.parse::<i32>().ok()).collect();
        ints[1] = ints[1] - 1;
        ints[2] = ints[2] - 1;

        let len = state[ints[1] as usize].len();
        let index_to_slice = len.saturating_sub(ints[0] as usize);
        let mut slice = state[ints[1] as usize][index_to_slice..].to_vec();

        state[ints[2] as usize].append(&mut slice);
        state[ints[1] as usize].truncate(index_to_slice);
    });
    let size = state.len() - 1;
    for i in 0..size + 1 {
        result.push_str(state[i].last().unwrap());
    }
    return result;
}