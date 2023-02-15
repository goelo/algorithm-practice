package sort

import "fmt"

func binarySearch1(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}

func binarySearch2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	//当我们更新区间时，如果左边界l更新为l=mid，此时mid的取值就应为mid=(l+r+1)/2。
	//因为当右边界r=l+1时，此时mid=(l+l+1)/2，下取整，mid仍为l，左边界再次更新为l=mid，相当于没有变化，
	//for循环就会陷入死循环。因此，我们总结出来一个小技巧，
	//当左边界要更新为l=mid时，我们就令 mid=(l+r+1)/2，上取整，此时就不会因为r取特殊值r=l+1而陷入死循环了。
	for l < r {
		mid := l + (r-l+1)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 7, 9, 23}
	target := binarySearch1(nums, 3)
	fmt.Println(target)
	fmt.Println(binarySearch2(nums, 3))
}
