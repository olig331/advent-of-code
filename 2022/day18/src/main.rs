use ::std::fs;

struct Coords {
    x: i32,
    y: i32,
    z: i32,
}

fn main() {
    let mut cubes: Vec<Coords> = Vec::new();
    let mut part_1_result = 0;

    let binding = fs::read_to_string("test.txt").expect("Err");
    let file: Vec<&str> = binding.split("\n").collect();

    for i in 0..file.len() {
        let ints = file[i]
            .split(",")
            .map(|i| i.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();

        cubes.push(Coords {
            x: ints[0],
            y: ints[1],
            z: ints[2],
        });
    }

    for i in 0..file.len() {
        let ints = file[i]
            .split(",")
            .map(|i| i.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();

        let mut count = 0;
        for j in 0..cubes.len() {
            if i == j {
                continue;
            }

            if ints[0] == cubes[j].x && ints[1] == cubes[j].y && abs(ints[2] - cubes[j].z) == 1 {
                count += 1;
            }

            if ints[0] == cubes[j].x && ints[2] == cubes[j].z && abs(ints[1] - cubes[j].y) == 1 {
                count += 1;
            }

            if ints[1] == cubes[j].y && ints[2] == cubes[j].z && abs(ints[0] - cubes[j].x) == 1 {
                count += 1;
            }
        }
        if count <= 6 {
            part_1_result += 6 - count;
        }
    }
    println!("{}", part_1_result);

    // part 2
}

fn abs(x: i32) -> i32 {
    let y = if x >= 0 { x } else { -x };
    y
}
