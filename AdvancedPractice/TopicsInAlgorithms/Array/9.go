func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	result := [][]int{}
	//排序很重要
	sort.Ints(candidates)

	findcombinationSum(candidates, target, 0, &result, []int{})
	for i := 0; i < len(result); i++ {
		sort.Ints(result[i])
	}
	for i := 0; i < len(result)-1; {

		tag := reflect.DeepEqual(result[i], result[i+1])

		if tag == true {

			result = append(result[0:i+1], result[i+2:]...)

		} else {
			i++
		}

	}
	return result

}

func findcombinationSum(candidates []int, target int, index int, result *[][]int, temp []int) {

	if target <= 0 {
		if target == 0 {
			*result = append(*result, temp)
		}

		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			return
		} else if candidates[i] <= target {
			temp = append(temp, candidates[i])
			findcombinationSum(candidates, target-candidates[i], index, result, temp)
			temp = temp[:len(temp)-1]

		}

	}

}