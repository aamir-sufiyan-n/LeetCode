/**
 * @param {string} title
 * @return {string}
 */
var capitalizeTitle = function(title) {
    let words=title.split(" ")
    let result=[]
    words.forEach((w)=>
    {
        if(w.length<=2)
        {
            let s=w.toLowerCase()
            return result.push(s)
        }
        else{
            let s=w[0].toUpperCase()+w.slice(1).toLowerCase()
            return result.push(s)
        }
    })

    return result.join(' ')
};