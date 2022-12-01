const path = require('path');
const fs = require('fs');

const type = {
  "0": [0, 1, 2, 4, 5, 6],
  "1": [2, 5],
  "2": [0, 2, 3, 4, 6],
  "3": [0, 2, 3, 5, 6],
  "4": [1, 2, 3, 5],
  "5": [0, 1, 3, 5, 6],
  "6": [0, 1, 3, 4, 5, 6],
  "7": [0, 2, 5],
  "8": [0, 1, 2, 3, 4, 5, 6],
  "9": [0, 1, 2, 3, 5, 6]
}

const main =  () => {  
  const input = fs
    .readFileSync(path.join(__dirname, 'data.txt'), 'utf8')
    .split('\n')
    .filter(Boolean)
    .map((line) => {
      const arr = line.split(" | ")
      return {input: arr[0].trim().split(" "), output: arr[1].trim().split(" ")}
    });

  const result = []
  let current = []
  input.forEach((obj) => {
    const board = calcBoard(obj.input)   
    let num = ""
    for(let i = 0; i < obj.output.length; i++) {
      num = ""
      let total = []
      obj.output[i].split("").forEach((char) => total.push(board.indexOf(char)))
      const sorted = total.sort((a,b) =>  a - b)
      for(const key in type) {        
        if(JSON.stringify(type[key]) === JSON.stringify(sorted)) {num += key}
      }
      current.push(num)
      total = []
    }
    result.push(parseInt(current.join("")))
    current = []
  })
  console.log(result.reduce((a, b) => a + b))
};


const calcBoard = (input) => {
  const allChars = ["a", "b", "c", "d", "e", "f", "g"]
  const board = new Array(7).fill("")
  const sorted = input.sort((a,b) => a.length - b.length)
  const top = sorted[1].split("").filter((char) => char !== sorted[0][0] && char !== sorted[0][1]);
  board[0] = top.join("")

  for(let i = 3; i < 6; i++) {    
    const curr = sorted[i].split("");
    if(curr.indexOf(sorted[0][0]) !== -1 && curr.indexOf(sorted[0][1]) !== -1) { // Bottom And middle 
      let three = sorted[i].split("")
      three.splice(three.indexOf(sorted[1][0]), 1)
      three.splice(three.indexOf(sorted[1][1]), 1)
      three.splice(three.indexOf(sorted[1][2]), 1)
      three.forEach((char) => sorted[2].split("").indexOf(char) !== -1 ? board[3] = char : board[6] = char)       
      break
    }
  }

  const topLeft = sorted[2].split("")  
    topLeft.splice(topLeft.indexOf(sorted[0][0]), 1)
    topLeft.splice(topLeft.indexOf(board[3]), 1)
    topLeft.splice(topLeft.indexOf(sorted[0][1]), 1)
    board[1] = topLeft.join("")
    
  for( let i = 3; i < 6; i++) { // check for the one with top left
    let curr = sorted[i].split("")
    if(curr.indexOf(board[1]) != -1 ){ // this is 5
      if(curr.indexOf(sorted[0][0]) !== -1){
        board[5] = sorted[0][0]
        board[2] = sorted[0][1]
        break
      }
      board[5] = sorted[0][1]
      board[2] = sorted[0][0]
      break      
    }
  } 

  const unused = allChars.filter((obj) => { return board.indexOf(obj) == -1; });
  for(let i = 0; i < board.length; i++) {
    if(board[i] === "") {
      board[i] = unused.join("")
    }
  }  
  return board
}

main()