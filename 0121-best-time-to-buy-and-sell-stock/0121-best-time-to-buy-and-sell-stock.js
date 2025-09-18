/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let profits=[]
    let bought=prices[0]
    for(let i=0;i<prices.length;i++)
    {
        if(prices[i]>bought)
        profits.push(prices[i]-bought)
        else if(prices[i]<=bought)
        bought=prices[i]
    }
    profits.length==0?profits.push(0):profits
    return (Math.max(...profits))
};