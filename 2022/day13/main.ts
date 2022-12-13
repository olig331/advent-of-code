import fs from "fs";

const pairs = fs
  .readFileSync("./input.txt", "utf8")
  .split("\n\n")
  .map((pair) => pair.split("\n").map((x) => JSON.parse(x)));

const part1Indexes: number[] = [];

const checkItems = (item1: any, item2: any): any => {
  if (typeof item1 === "number" && typeof item2 === "number") {
    if (item1 === item2) {
      return 0;
    }
    if (item1 < item2) {
      return 1;
    }
    return -1;
  }

  if (typeof item1 === "number" && typeof item2 === "object") {
    item1 = [item1];
  }

  if (typeof item2 === "number" && typeof item1 === "object") {
    item2 = [item2];
  }

  for (let i = 0; i < item1.length; i++) {
    if (item2[i] === undefined) {
      return -1;
    }

    const newItem = checkItems(item1[i], item2[i]);
    if (newItem !== 0) {
      return newItem;
    }
  }

  if (item1.length === item2.length) {
    return 0;
  } else if (item1.length > item2.length) {
    return -1;
  } else {
    return 1;
  }
};

for (let i = 0; i < pairs.length; i++) {
  const pair1 = pairs[i][0];
  const pair2 = pairs[i][1];

  if (checkItems(pair1, pair2) > 0) {
    part1Indexes.push(i + 1);
  }
}

const dividers = [[[2]], [[6]]];

const part2 = pairs
  .flat()
  .concat(dividers)
  .sort((a, b) => checkItems(a, b))
  .reverse();

const dividerIndexes: number[] = [];

for (let i = 0; i < part2.length; i++) {
  const toCompare = JSON.stringify(part2[i]);
  if (toCompare === "[[2]]" || toCompare === "[[6]]") {
    dividerIndexes.push(i + 1);
  }
}

console.log(
  "Part 1::",
  part1Indexes.reduce((a, b) => a + b)
);
console.log("Part 2::", dividerIndexes[0] * dividerIndexes[1]);
