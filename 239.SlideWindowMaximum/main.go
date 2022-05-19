package main

// 滑动窗口最大值
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sliding-window-maximum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//结果数组
	result := []int{}
	//  double-ended queue，双端队列
	queue := []int{}

	//遍历参数数组
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
			//一旦新的元素比队列中最的元素大,则把所有元素删除
			queue = queue[:len(queue)-1]
		}
		//将当前元素放入queue中
		queue = append(queue, nums[i])
		//维护队列,保证其头元素(queue[0])为当前窗口最大值
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		if i >= k-1 {
			//放入结果数组,queue[0]永远是最大的
			result = append(result, queue[0])
		}

	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7, 1, 3, 4, 5, 6, 1, 3, 3, 4}
	k := 3
	result := maxSlidingWindow(nums, k)
	println(result)

}
