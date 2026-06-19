func angleClock(hour int, minutes int) float64 {
    hourAngle:=30.0 * float64(hour) + 0.5 * float64(minutes)
    minuteAngle:= 6.0 * float64(minutes)

    res:=math.Abs(hourAngle-minuteAngle)
    if res>180{ return 360 - res}
    return res
}