const path = require('path');
const fs = require('fs');

const directions = [
  { x: -1, y: 1 },
  { x: -1, y: -1 },
  { x: 1, y: 1 },
  { x: 1, y: -1 },
  { x: 0, y: 1 },
  { x: 0, y: -1 },
  { x: -1, y: 0 },
  { x: 1, y: 0 },
]

let part2Result = 0
let count = 0;
let data = fs.readFileSync(path.join(__dirname, "data.txt"), "utf8")
  .split("\n")
  .filter(Boolean)
  .map((line) => {
    const arr = line.split("")
    return arr.map((char) => parseInt(char))
  })

const main = () => {
  const cLen = data[0].length - 1
  const rLen = data.length - 1
  let stop = false
  while (!stop) { // replace with 100 range loop for part 1 

    for (let y = 0; y < data.length; y++) {
      for (let x = 0; x < data[y].length; x++) {
        data[y][x] += 1
      }
    }

    for (let y = 0; y < data.length; y++) {
      for (let x = 0; x < data[y].length; x++) {
        if (data[y][x] >= 10) {
          flash(y, x, cLen, rLen)
        }
      }
    }

    for (let y = 0; y < data.length; y++) {
      for (let x = 0; x < data[y].length; x++) {
        if (data[y][x] === -1) {
          data[y][x] = 0
        }
      }
    }
    if (sumBoard() === 0) {
      stop = true
    }
    part2Result++
  }
  console.log(part2Result)
  console.log(count)
}

const flash = (y, x, cLen, rLen) => {
  data[y][x] = -1
  count += 1

  for (let i = 0; i < directions.length; i++) {
    const newY = y + directions[i].y
    const newX = x + directions[i].x

    if (newY >= 0 && newY <= rLen && newX >= 0 && newX <= cLen && data[newY][newX] !== -1) {
      data[newY][newX] += 1
      if (data[newY][newX] >= 10) {
        flash(newY, newX, cLen, rLen)
      }
    }
  }
}

const sumBoard = () => {
  let sum = 0;
  for (let i = 0; i < data.length; i++) {
    for (let j = 0; j < data[i].length; j++) {
      sum += data[i][j]
    }
  }
  return sum
}

main()