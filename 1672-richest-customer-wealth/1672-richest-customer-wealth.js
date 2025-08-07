/**
 * @param {number[][]} accounts
 * @return {number}
 */
var maximumWealth = function(accounts) {
    let wealth=0;
    let arr=[];
    for(let i=0;i<accounts.length;i++)
    {
        wealth=accounts[i].reduce((prev,current)=>prev+current )
        arr.push(wealth);
    }    
    return Math.max(...arr);
};