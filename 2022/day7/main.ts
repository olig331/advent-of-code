import fs from "fs";

class TreeNode {
  public children: TreeNode[];
  public type: string;
  public parent: TreeNode | null;
  public size: number;
  public name: string;

  constructor(name: string, type: string, parent?: TreeNode, size?: number) {
    this.name = name;
    this.children = [];
    this.type = type || "";
    this.parent = parent || null;
    this.size = size || 0;
  }
}

const tree = new TreeNode("/", "dir");

let current = tree;

const l = fs.readFileSync("./input.txt", "utf8").split("\n");

for (let i = 1; i < l.length; i++) {
  if (l[i].startsWith("$ cd ..")) {
    if (current.parent) {
      current = current.parent;
    }
    continue;
  }

  if (l[i].startsWith("$ cd ")) {
    const dirName = l[i].replace("$ cd ", "");
    const isNewDir = current.children.findIndex((n) => n.name === dirName);
    if (isNewDir > 0) {
      current = current.children[isNewDir];
      continue;
    } else {
      current.children.push(new TreeNode(dirName, "dir", current));
      current = current.children[current.children.length - 1];
    }
    continue;
  }

  if (l[i].startsWith("$ ls")) {
    continue;
  }

  if (!l[i].startsWith("dir")) {
    const fileStats = l[i].split(" ");
    current.children.push(
      new TreeNode(fileStats[1], "file", current, parseInt(fileStats[0], 10))
    );
  }
}

const solve = () => {
  const allTotals: number[] = [];

  const used = 42476859;
  const sizeNeeded = 30000000;
  const totalSize = 70000000;
  const freeSpace = totalSize - used;

  let part1Tally = 0;
  const recurse = (folder: TreeNode) => {
    let total = 0;
    for (let i = 0; i < folder.children.length; i++) {
      const file = folder.children[i];
      if (file.type === "dir") {
        total += recurse(file);
      } else {
        total += file.size!;
      }
    }
    if (total <= 100000) {
      part1Tally += total;
    }

    if (freeSpace + total >= sizeNeeded) {
      allTotals.push(total);
    }

    return total;
  };
  recurse(tree);
  allTotals.sort((a, b) => a - b);

  console.log("PART 1:", part1Tally);
  console.log("PART 2: ", allTotals[0]);
};

solve();
