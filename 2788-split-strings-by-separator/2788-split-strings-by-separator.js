/**
 * @param {string[]} words
 * @param {character} separator
 * @return {string[]}
 */
var splitWordsBySeparator = function(words, separator) {
    let result=[]
    for(let i=0;i<words.length;i++)
    {
        let word=words[i].split(separator)
        for(let j=0;j<word.length;j++)
        {
            if(word[j].trim())
             result.push(word[j])
        }
    }
    return result
};