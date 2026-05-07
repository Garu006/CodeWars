package kata
package main

func CountPositivesSumNegatives(numbers []int) []int {
  var res []int
  return res
  
  // your code here
  count:= 0
  sum:= 0
  for i:=0; i<len(numbers); i++ {
	if numbers[i] > 0 {
		count++
	} else if numbers[i] < 0 {
		sum += numbers[i]
		res = append(res, sum, count)
	}
  }
}

func main() {
	CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, -1, -2, -3, -4, -5})
	CountPositivesSumNegatives([]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	CountPositivesSumNegatives([]int{1, 2, 3, 4, 5})
	CountPositivesSumNegatives([]int{-1, -2, -3, -4, -5})
	CountPositivesSumNegatives([]int{})
	CountPositivesSumNegatives(nil)
	CountPositivesSumNegatives([]int{0})
	CountPositivesSumNegatives([]int{-1})
	CountPositivesSumNegatives([]int{1})
	CountPositivesSumNegatives([]int{-1, 2})
	CountPositivesSumNegatives([]int{1, -2})
	CountPositivesSumNegatives([]int{1, -2, 3})
	CountPositivesSumNegatives([]int{-1, 2, 3})
	CountPositivesSumNegatives([]int{1, -2, 3, -4})
	CountPositivesSumNegatives([]int{-1, 2, 3, -4})
	CountPositivesSumNegatives([]int{1, -2, 3, -4, 5})
	CountPositivesSumNegatives([]int{-1, 2, 3, -4, 5})
}
