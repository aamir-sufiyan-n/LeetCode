func smallestNumber(n int, t int) int {
  for{
    prod:=1
    num:=n
    for num != 0 {
        digit:=num%10
        prod=prod*digit
        num/=10
    }
    if prod%t==0{return n}else{n++}
  }
    return 0   
}