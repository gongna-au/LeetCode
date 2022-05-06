func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := [][]int{[]int{}}
	//记录去重后的nums的下标
	index := []int{}
	i := 0
	j := i + 1
	for i < len(nums) {
		if j > len(nums)-1 {
			index = append(index, i)
			break
		}
		if nums[i] != nums[j] {
			index = append(index, i)
			i = j
			j++
		} else {
			j++
		}
	}

	//递归调用
	for _, v := range index {

		temp := []int{}
		temp = append(temp, nums[v])
		findsubsetsWithDup(nums, v, &result, temp)
	}
	for i := 0; i < len(result); {
		for j := i + 1; j < len(result); {
			if reflect.DeepEqual(result[i], result[j]) {
				result = append(result[:j], result[j+1:]...)
			} else {
				j++
			}

		}
		i++
	}
	return result

}
func findsubsetsWithDup(nums []int, start int, result *[][]int, temp []int) {
	temp2 := make([]int, len(temp))
	copy(temp2, temp)
	*result = append(*result, temp2)
	if start == len(nums)-1 {
		return
	}

	j := 1

	temp3 := []int{}
	for i := start + 1; i < len(nums); i++ {
		temp3 = append(temp, nums[i])
		findsubsetsWithDup(nums, start+j, result, temp3)
		j++

	}

}