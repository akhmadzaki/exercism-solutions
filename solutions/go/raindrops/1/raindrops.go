package raindrops

import "strconv"

func Convert(number int) string {
	ret := ""
    if number%3!=0 && number%5!=0 && number%7!=0 {
        return strconv.Itoa(number)
    } 
	if number%3==0 {
        ret += "Pling"
    }
	if number%5==0 {
        ret += "Plang"
    }
	if number%7==0 {
        ret += "Plong"
    }

    return ret
}
