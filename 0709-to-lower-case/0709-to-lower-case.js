/**
 * @param {string} s
 * @return {string}
 */
var toLowerCase = function(s) 
{
    splitted=s.split("");
    for(let i=0;i<splitted.length;i++)
    {
        if(splitted[i]===splitted[i].toUpperCase())
           splitted[i]=splitted[i].toLowerCase();
    }
    return splitted.join("");
};