/*
Input:
- List of direction and amount
    - i.e. L5 -> Left rotation by 5
    - Will always be one letter followed immediataly by the number
- Dial starts at 50
- Max dial number is 99 -> You cannot rotate twice in one op
Goal:
- The number of times the dial is left pointing at 0 at the end of any rotation
- Number is going to be between 0 and 99
Naive Approach:
- Use variable to keep track of current number (start at 50)
- Use variable to keep track password count
- Loop through each operation
    - split the operation into a direction and amount
    - If it is L it is a negative, if it is R it is a positive
    - Calculate new number by adding to current number 
    - Take the mod of that number and 100 (number mod 100)
    - If number is 0 or 100 increment main counter
- O(n) time
- O(n) space
Improved Approach:
*/

// Include the fs module
const fs = require('fs');

// Read the file synchronously
const data = fs.readFileSync('./input1.txt', { encoding: 'utf8', flag: 'r' });
const ops = data.split('\n');
let maxOp = 0;

ops.forEach((op: string) => {
    if (op.length > 0) {
        const dir = op[0];
        const clicks = op.slice(1);
        if (parseInt(clicks) > maxOp) {
            maxOp = parseInt(clicks);
        }
    }
});

// Display the file content
console.log(maxOp);
