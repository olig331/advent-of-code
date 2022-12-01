const path = require('path')
const fs = require('fs')

let data = fs
  .readFileSync(path.join(__dirname, "data.txt"), "utf8")
  .split("\n")
  .filter(Boolean)
  .map((line) => {
    let arr = line.split("")
    return arr.map((char) => parseInt(char))
  })

let reuslts = []
let count = 0

const directions = [
	{x: 0, y: 1},
	{x: 0, y: -1},
	{x: -1, y: 0},
	{x: 1, y: 0},
];

const main = () => {

  for(let i = 0; i < data.length; i++) {
    for(let j = 0;j < data[i].length; j++) {
      if(data[i][j] !== 9)  {
        data[i][j] = null
      }
    }
  }

  for(let i = 0; i < data.length; i++) {
    for(let j = 0; j < data[i].length; j++) {
      if(data[i][j] === null) {
        count++
        data[i][j] = -1
        check(i, j)
        reuslts.push(count)
        count = 0
      }
    }
  }

  results = reuslts.sort((a,b) => b - a)
  console.log(results[0] * results[1] * reuslts[2])
}

function check(y, x) {
  for(let i = 0; i < directions.length; i++) {    
    let newY = y + directions[i].y
    let newX = x + directions[i].x
    if(newY >= 0 && newY < data.length && newX >=0 && newX < data[0].length) {
      if(data[newY][newX] === null){
        count++
        data[newY][newX] = - 1
        check(newY, newX)
      }
    }
  } 
}

main()

