/**
 * @param {string} moves
 * @return {number}
 */
var furthestDistanceFromOrigin = function(moves) {
 let position=0;
 let Rcount=0
 let Lcount=0
 for(let i=0;i<moves.length;i++)
 {
    if(moves[i]=="L")
    Lcount++ 
    else if(moves[i]=="R")
    Rcount++ 
 }
 for(let j=0;j<moves.length;j++)
 {
    if(moves[j]=="L")
    position--
    else if(moves[j]=="R")
    position++
    else if(moves[j]=="_"&&Lcount>Rcount)
    position--
    else if(moves[j]=="_"&&Lcount<=Rcount)
    position++
 }
 return Math.abs(position)
};