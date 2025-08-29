/**
 * @param {number[]} arr
 * @return {boolean}
 */
var uniqueOccurrences = function(arr) {
    let count=0;
    let array=[];
    let countArray=[];
    let result=true;
    for(let i=0;i<arr.length;i++)
    {
        if(!array.includes(arr[i]))
        {
            for(let j=0;j<arr.length;j++)
            {
                if(arr[j]==arr[i])
                count++
            }
            array.push(arr[i])
            countArray.push(count)
            count=0
        }
    }
    result=countArray.length===new Set(countArray).size;
    return result
};