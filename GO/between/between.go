//complete the sentence that takes two integers and return an array of all integers between the input parameters, including them.

package kata

func Between(a, b int) []int {
	result := make([]int, b-a +1)
	for i := 0; i <= b-a; i++ {
		result[i] = a + i
	}
	return result
}

func main(){
	Between(1,4)
}
