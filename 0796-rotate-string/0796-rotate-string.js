/**
 * @param {string} s
 * @param {string} goal
 * @return {boolean}
 */
var rotateString = function(s, goal) {
    let joined=s+s
    if(s.length != goal.length){
        return false
    }
    return joined.includes(goal)
};