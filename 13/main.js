const path = require('path')
const fs = require('fs')

/// INCOMPLETE 

const main = () => {
  const folds = []
  let lenY = 895, lenX = 1311
  const input = []
  let data = fs.readFileSync(path.join(__dirname, "data.txt"), "utf8").split("\n").filter(Boolean);
  
  for(let i = 0; i < data.length; i++) {
    const arr = data[i].split(",")
    if(arr.length < 2) {
      folds.push(arr.join(","))
      continue
    } else {        
      if(parseInt(arr[0]) > lenY) lenY = parseInt(arr[0])
      if(parseInt(arr[1]) > lenX) lenX = parseInt(arr[1])
      input.push({x: parseInt(arr[0]), y: parseInt(arr[1])})
    }
  }

  let grid = []
  for(let y = 0; y <= lenY; y++){
    const row = []
    for(let x = 0; x <= lenX; x++){
      row.push(0)
    }
    grid.push(row)
  }

  let foldIndex = 0
  let yL = grid.length 
  let xL = grid[0].length

  input.forEach((coord) => {grid[coord.y][coord.x] += 1})

  folds.forEach((fold, index) => {
    foldIndex = fold.replace(/foldalong(x|y)\=/g, "")
    let foldDir = fold.match(/[x|y]/g)[0]
    console.log(foldDir, foldIndex)
    if(foldDir === "y"){
      grid = foldY(foldIndex, grid, xL, yL)
      yL = foldIndex
    } else if(foldDir === "x") {
      grid = foldX(foldIndex, grid, yL, xL)
      xL = foldIndex
    }
    if(index == 0){
      console.log(count(grid, foldIndex, grid.length ))
    }
  })

  const result = []
  for(let y = 0; y < yL - 1; y++) {
    const row = []
    for(let x = 0; x < xL; x++){
      if(grid[y][x] > 0){
        row.push("#")
      }else {
        row.push(" ")
      }
    }
    result.push(row)
  }
  result.forEach((row) => console.log(row.join("")))
}

const count = (grid, xMax, yMax) => {
  let count = 0
  for(let y = 0; y < yMax; y++ ){
    for(let x = 0; x < xMax; x++){
      if(grid[y][x] > 0) {
        count += 1
      }
    }	  
  }
  return count
}



const foldX = (foldIndex, grid, yL, xL) => {
  for(let y = 0; y < yL - 1; y++ ){
    for (let x = 0; x < foldIndex - 1; x++) {
      grid[y][x] += grid[y][xL - x - 2]
    }
  }
  return grid
}

const foldY = (foldIndex, grid, xL, yL) => {
  for(let y = 0; y < foldIndex - 1; y++ ){
    for (let x = 0; x < xL - 1; x++) {
      grid[y][x] += grid[yL - y - 2][x]
    }
  }
  return grid
}

main()