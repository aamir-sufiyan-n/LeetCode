/**
 * @param {string} moves
 * @return {boolean}
 */
var judgeCircle = function(moves) {
    let x=0;
    let y=0;
    let splitted=moves.split("");

    splitted.forEach(Element=>{
        switch(Element)
        {
        case "U":
        y++
        break;
        case "D":
        y--
        break;
        case "L":
        x--
        break;
        case "R":
        x++
        break;
        }
        

    })
    if(x==0 && y==0)
        return true;
    else 
        return false;
};