package main
import "fmt"

func cipher(message int64, e int64, n int64) int64{
    return (message^e)%n
}

func mdc(a int64, b int64) int64{
    if b==0{
        return a
    }
    return mdc(b, a%b)
}

func findE(n int64) int64{
    var ret, mdct int64
    mdct = 0
    ret = 3
    for i:= 3; mdct != 1; i++{
        ret = int64(i)
	mdct = mdc(int64(i),n)
    }
    return ret
}

func findD(e int64,n int64) int64{
    var i,result int64
    result = 0
    for i = 2; result != 1; i++{
        if (e-i)%n == 0 
           result = i	
    }
    return i
}

func main(){
    var n, e, d, message, phiN, xMsg, msg int64
    n = 2*7 //p*q
    phiN = 1*6 // 
    e = findE(phiN) //Euler phi n = (p-1)*(q-1)
    d = findD(e,phiN)
    message = 33
    fmt.Println(message)
    xMsg = cipher(message,e,n)
    fmt.Println(xMsg)
    msg = cipher(xMsg, d, n)
    fmt.Println(msg)

}
