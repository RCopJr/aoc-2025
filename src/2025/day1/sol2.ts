/*
Input:
- List of direction and amount
    - i.e. L5 -> Left rotation by 5
    - Will always be one letter followed immediataly by the number
- Dial starts at 50
- Max dial number is 99 -> You cannot rotate twice in one op
Goal:
- Number of times the dial crosses 0 at any point of the process
- This can even happen during an operation
- Remember that since the max value is 99 it can never cross 0 twice on one operation
Naive Approach:
- When do we know when it will cross 0?
    - If previous value becomes positive to negative
    - [This might cover all cases] When you cross any hundreds value: 0, 100, 200, and the -ves as well
    - We want to make sure operation still happens at constant time -> just need to update our calculation method
- How to know when you cross a hundreds value?
    - [Handles when you just cross 0] Can divide by 100 and parse into an int. As long as long as the value is the same before and after the operation, you did not cross 0
- [Handles case when you reach a hundreds value] When mod is just 0, you increment counter
    - On next operation, you already know that you cannot cross 0 again.
    - Should just skip counter check if previous mod was 0
- Corner Cases?
    - If you are in a hundreds value and reach 0 or 100, does it change your hundreds value?
    - What to do if you are already at a 0 value?
    - Can you get 0? -> No, smallest operation is 1
- Use variable to keep track of current number (start at 50)
- Use variable to keep track password count
- Flag to see if old dial val was a 0 -> skip password check, just update dial number
- Loop through each operation
    - split the operation into a direction and amount
    - If it is L it is a negative, if it is R it is a positive
    - Calculate new number by adding to current number and store it in a new variable
    - If flag is false:
        - Take the mod of that number and 100 (number mod 100)
            - If value is 0, increment counter and set flag to true
            - Else, divide old number and floor by 100 and do the same with new number.
                - If they are different, increment password, else, continue.
    - If flag is true:
        - set flag to false (We know for sure you can never land on a 0 again)
- O(n) time
- O(n) space
Bugs: 
- -10 / 100 floored is 1 butt, -110 / 100 floored is also -1
    - So I think problem was arising whenever I went from 0 to -100 to something even lower than -100 
Solution:
- new edge case: old num and new num have different signs
- For negative old and new nums, abs them, then divide and floor, then compare
- Nevermind, floor goes down, so it should still be fine
Bugs: 
- Biggest number is 999 not 99
- Therefore, you can pass 0 even more than once, andddd can pass 0 even when you are on 0
Solution:
- Amount of times you have passed 0 should just be difference in 100's
- No longer need the flag
- Need to count different in 100's and check if final number is 0 or nah?
- Just cross once -> 100 should be different
- cross multiple times -> 100 should be different and number of crosses is difference in 100;s
Bugs: 
- 50 L 150 should be 2 but becomes 1
- This is because floor -5 = floor -100
- Positive is fine, it is just negative
Solution:
- if negative, Maybe subtract 1 to the currNum before calculating? 
     - this way, -100 will be -2 and -2 will still be -1
     - Doesn't work because arithmetic will be off
- What if you do if cross 0 you add 1 (positive to negative), else make the negative hundredsd match the positive hundreds

NEW THINKING
- if operation is small enough, you no for sure no zeros
- Only when operation is big enough do you know -> Changes depending on R or L operation
- if it is big enough, you know that operation - val is already 1 intersection.
    - Any intersection after than is just (operation - val) % 100
*/

const fs = require('fs');

const data = fs.readFileSync('./input1.txt', { encoding: 'utf8', flag: 'r' });
const ops = data.split('\n');

let dial = 50;
let passCount = 0;

ops.forEach((op: string) => {
    if (!op) {
        return;
    }
    const dir = op[0];
    const clicks = parseInt(op.slice(1));
    const originalDial = dial;

    if (dial === 0) {
        passCount += Math.floor(clicks / 100);
        if (dir === 'L') {
            if (clicks > 100) {
                const extraClicks = clicks - dial;
                dial = (extraClicks % 100) == 0 ? 0 : 100 - (extraClicks % 100);
            } else {
                dial = 100 - clicks;
            }
        } else if (dir === 'R') {
            if (clicks >= 100) {
                const extraClicks = clicks - (100 - dial);
                dial = extraClicks % 100;
            } else {
                dial += clicks;
            }

        }
    } else if (dir === 'L') {
        if (clicks >= dial) {
            passCount += 1;
            const extraClicks = clicks - dial;
            passCount += Math.floor(extraClicks / 100);
            dial = (extraClicks % 100) == 0 ? 0 : 100 - (extraClicks % 100);
        } else {
            dial -= clicks;
        }
    } else if (dir === 'R') {
        if ((100 - dial) <= clicks) {
            passCount += 1;
            const extraClicks = clicks - (100 - dial);
            passCount += Math.floor(extraClicks / 100);
            dial = extraClicks % 100;
        } else {
            dial += clicks;
        }
    }
    if (dial === 100) {
        console.log(originalDial, dial, op)
    }
    // console.log(dial);
});

// Display the file content
console.log(passCount);


    // const oldNum = currNum;
    // if (dir === 'L') {
    //     currNum -= numClicks;
    // } else {
    //     currNum += numClicks;
    // }
    //
    // let oldHundreds;
    // let newHundreds;
    //
    // oldHundreds = oldNum < 0 ? Math.ceil(oldNum/100) : Math.floor(oldNum/100);
    // newHundreds = currNum < 0 ? Math.ceil(currNum/100) : Math.floor(currNum/100);
    //
    // if (oldNum > 0 && currNum <= 0 || currNum >= 0 && oldNum < 0) {
    //     passCount += 1;
    //     if (oldNum % 100 === 0) {
    //         passCount -= 1;
    //     }
    //     console.log(oldNum, op, currNum);
    // }
    //
    // if (oldHundreds !== newHundreds) {
    //     passCount += Math.abs(newHundreds - oldHundreds);
    //     if (oldNum % 100 === 0) {
    //         passCount -= 1;
    //     }
    //     console.log(oldNum, op, currNum);
    // } 
export {};
