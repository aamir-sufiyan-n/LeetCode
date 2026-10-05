func processStr(s string) string {
    res:=[]byte{}
    for _,c:=range s {
        if c>='a' && c<='z'{
            res=append(res,byte(c))
        } else if c=='*' {
            if len(res)>0{
                res=res[:len(res)-1]
            }
        } else if c=='#' {
            res=append(res,res...)
        } else {
            for i,j:=0,len(res)-1; i<j; i,j=i+1, j-1 {
                res[i],res[j] = res[j],res[i]
            }
        }
    }
    return string(res)
}