/**
 * @param {string} address
 * @return {string}
 */
var defangIPaddr = function(address) {
    // let addressnew="";
    // for(let i=0;i<address.length;i++){
    //     if(address[i]===".")
    //         addressnew+="[.]"
    //     else{
    //         addressnew+=address[i]
    //     }
    // }
    // return addressnew
    return address.replaceAll(".","[.]")
};