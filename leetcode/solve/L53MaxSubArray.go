package main

/**
* 给一个整数数组nums 找到具有最大和的连续字数组
public int maxSubArray(int[] nums) {
       if (nums == null || nums.length < 1) {
           return 0;
       }

       int result = nums[0];
       int temp = 0;
       for (int i =0 ;i < nums.length ; i++) {
           int t = nums[i];
           result = t > result ? t : result;

           if (t > 0) {
               if (temp > 0) {
                   temp += t;
                  result = temp > result? temp: result;
               } else {
                   temp = t;
               }
           } else {

               temp += t;
           }
       }
       return result;
   }
*/
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	result := nums[0]
	temp := 0

	for i := 0; i < len(nums); i++ {
		t := nums[i]
		// Go 不支持三元运算符 result = t > result ? t : result
		if t > result {
			result = t
		}

		if t > 0 {
			if temp > 0 {
				temp = temp + t
				if temp > result {
					result = temp
				}
			} else {
				temp = t
			}
		} else {
			temp = temp + t
		}
	}
	return result
}

// func main() {

// }
