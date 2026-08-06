func smallestNumber(n int, t int) int {
  for{
    prod:=1
    num:=n
    for num != 0 {
        prod*=num%10
        num/=10
    }
    if prod%t==0{return n}else{n++}
  }
}