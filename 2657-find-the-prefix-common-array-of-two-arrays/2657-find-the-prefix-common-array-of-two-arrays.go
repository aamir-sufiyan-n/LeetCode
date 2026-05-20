func findThePrefixCommonArray(A []int, B []int) []int {
    n:=len(A)
    var res=make([]int,n)
    SeenA:=make([]bool,n+1)
    SeenB:=make([]bool,n+1)
    count:=0

    for i:=0;i<len(A);i++{
        SeenA[A[i]]=true
        SeenB[B[i]]=true

        if SeenA[A[i]] && SeenB[A[i]]{
            count++
        }
        if A[i]!=B[i] && SeenA[B[i]] && SeenB[B[i]]{
            count++
        }
        res[i]=count
    }
    return res
}