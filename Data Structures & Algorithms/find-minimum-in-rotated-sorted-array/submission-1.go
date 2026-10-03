func findMin(nums []int) int {
	l,r := 0,len(nums)-1
	mini := nums[0]

	for l <= r{
		if nums[l]<nums[r] && nums[l]<mini{
			mini = nums[l]
			break
		}

		mid := l+(r-l)/2
		if(nums[mid]< mini){
			mini = nums[mid]
		}

		if nums[mid]>=nums[l]{
			l = mid+1
		}else{
			r = mid-1
		}
	}
	return mini
}
