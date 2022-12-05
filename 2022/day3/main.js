const arr = [
  "vJrwpWtwJgWrhcsFMMfFFhFp",
"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
"PmmdzqPrVvPwwTWBwg",
"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
"ttgJtRGJQctTZtZT",
"CrZsJsPPZsGzwwsLwLmpwMDw",
]

let total = 0
// fs.readFileSync("./test.txt", "utf8").split("\n")
arr.forEach((line) => {
  let lowerCaseCount = 0; //  a = 97 - 96
  let upperCaseCount = 0; // -64  A = 65
  let count = 0
  line.split("").forEach((char) => {
    if(char === char.toUpperCase()){
      upperCaseCount++;
    }else {
      lowerCaseCount++
    }
    count += char.charCodeAt(0);
  })
  
  const lc = lowerCaseCount * 96;
  const uc = upperCaseCount * 38;

  total += count - (lc + uc)
})

console.log(total)