package main
import "fmt"
func Solution(A []int) int {
	n := len(A)
    count := 0
    for i := 0; i < n-1; i++ {
        if A[i] > A[i+1] {
            if  A[i+1] >= A[i-1] { //Kiểm tra pt sau A[i] và trước A[i]có lớn hơn hoặc bằng nếu nhỏ hơn thì return 0
                for j := i + 1; j < n-1; j++ { 
                    if A[j] > A[j+1] {//Kiểm tra phần tử sau A[j] > A[j + 1] - mất chuỗi lớn hơn return 0
                        return 0
                    }else if A[j] == A[j+1]{
                        count--
                    }else {
                        return 0
                    }
                }
                count++
            } else {
                return 0
            }	 
            count++
        }
    }
    if count != 0 {
        return count
    }
    return n
}

func main() {
	var N int
    for N <= 0 || N > 200{
        fmt.Println("Nhập số cây từ [1..200]: ")
	    fmt.Scan(&N)
        continue
    }
	A := make([]int, N)
	fmt.Println("Nhập chiều cao của cây trong khoảng[1..1000]: ")
	for i := 0; i < N; i++ {
        for A[i] <= 0 || A[i] > 1000{
            fmt.Printf("Tree %d: ", i+1 )
		    fmt.Scan(&A[i])     
            continue
        }
		
	}
	fmt.Println("So cay co the cat: ", Solution(A))
}