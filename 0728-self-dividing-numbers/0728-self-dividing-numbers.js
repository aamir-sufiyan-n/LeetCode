/**
 * @param {number} left
 * @param {number} right
 * @return {number[]}
 */
var selfDividingNumbers = function(left, right) {
    let divisible=false
    let array=[]
    for(let i=left;i<=right;i++)
    {
        let num=String(i).split("")
        for(j=0;j<num.length;j++)
        {
            if((num[j]!==0)&&(i%Number(num[j])==0))
            divisible=true
            else 
            {
                divisible=false
                break;
            }
        }
        if(divisible==true){
            array.push(i)
            divisible=false}
    }
    return array
};