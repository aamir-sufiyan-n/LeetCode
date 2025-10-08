/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    let array=[...nums1,...nums2]
    let sorted=array.sort()
    if(sorted.length%2==0)
    {
        let median = sorted[sorted.length/2]+sorted[(sorted.length/2)-1]
        return median/2
    }
    else{
        return sorted[Math.floor(sorted.length/2)]
    }
};