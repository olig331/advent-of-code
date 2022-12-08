use std::fs;

fn main() {
    let mut grid: Vec<Vec<i32>> = Vec::new();

    fs::read_to_string("input.txt")
        .expect("ok")
        .split("\n")
        .for_each(|f| {
            grid.push(
                f.split("")
                    .filter(|s| !s.is_empty())
                    .map(|i| i.parse::<i32>().unwrap())
                    .collect::<Vec<i32>>(),
            )
        });

    let mut part_1_count = 0;
    let mut part_2_res = 0;

    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            let mut x = j;
            let curr = grid[i][j];
            let mut is_done = false;

            if j == 0 || j == grid[i].len() - 1 {
                part_1_count += 1;
                continue;
            }

            if i == 0 || i == grid.len() - 1 {
                part_1_count += 1;
                continue;
            }

            while x > 0 {
                x -= 1;
                if grid[i][x] >= curr {
                    break;
                }
                if x == 0 && grid[i][x] < curr {
                    is_done = true;
                    part_1_count += 1;
                }
            }
            x = j;
            while x < grid[i].len() - 1 {
                x += 1;
                if is_done {
                    break;
                }
                if grid[i][x] >= curr {
                    break;
                }
                if x == grid[i].len() - 1 && grid[i][x] < curr {
                    is_done = true;
                    part_1_count += 1;
                }
            }

            let mut y = i;
            while y > 0 {
                if is_done {
                    break;
                }
                y -= 1;
                if grid[y][j] >= curr {
                    break;
                }
                if y == 0 && grid[y][j] < curr {
                    is_done = true;
                    part_1_count += 1;
                }
            }
            y = i;
            while y < grid.len() - 1 {
                if is_done {
                    break;
                }
                y += 1;
                if grid[y][j] >= curr {
                    break;
                }
                if y == grid.len() - 1 && grid[y][j] < curr {
                    part_1_count += 1;
                }
            }
        }
    }

    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            let mut x = j;
            let curr = grid[i][j];
            let mut score: Vec<i32> = Vec::new();

            if j == 0 || j == grid[i].len() - 1 {
                score.push(0);
            }

            if i == 0 || i == grid.len() - 1 {
                score.push(0);
            }

            while x > 0 {
                x -= 1;
                if x == 0 || grid[i][x] >= curr {
                    score.push(abs(j as i32 - x as i32));
                    break;
                }
            }
            x = j;

            while x < grid[i].len() - 1 {
                x += 1;
                if x == grid[i].len() - 1 || grid[i][x] >= curr {
                    score.push(abs(j as i32 - x as i32));

                    break;
                }
            }
            let mut y = i;
            while y > 0 {
                y -= 1;

                if y == 0 || grid[y][j] >= curr {
                    score.push(abs(i as i32 - y as i32));
                    break;
                }
            }

            y = i;
            while y < grid.len() - 1 {
                y += 1;
                if y == grid.len() - 1 || grid[y][j] >= curr {
                    score.push(abs(i as i32 - y as i32));

                    break;
                }
            }

            let res = score
                .iter()
                .reduce(|a, b| Box::leak(Box::new(a * b)))
                .unwrap();

            if res > &part_2_res {
                part_2_res = *res;
            }
        }
    }

    println!("Part 1: {}", part_1_count);
    println!("Part 2: {}", part_2_res);
}

fn abs(x: i32) -> i32 {
    let y = if x >= 0 { x } else { -x };
    y
}
