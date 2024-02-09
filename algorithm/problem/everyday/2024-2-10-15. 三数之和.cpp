/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/

// 暴力法 时间复杂度 O(n^3)
func threeSum2(nums []int) [][]int {
    result := make([][]int, 0)
    mp := make(map[string]struct{})
    n := len(nums)
    for i := 0;i<n;i++ {
        for j:=i+1;j<n;j++ {
            for k:=j+1;k<n;k++ {
                if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
                    if _, ok := mp[getkey(nums[i],nums[j],nums[k])]; !ok {
                        result = append(result, []int{nums[i],nums[j],nums[k]})
                        mp[getkey(nums[i],nums[j],nums[k])] = struct{}{}
                    }
                }
            }
        }
    }
    return result
}

func getkey(a,b,c int) string {
    d := []int{a, b, c}
    sort.Ints(d)
    return fmt.Sprintf("%d%d%d", d[0], d[1], d[2])
}

// O(n^2) 先排序
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r := i+1, len(nums)-1
        for l < r {
            s := nums[i] + nums[l] + nums[r]
            if s < 0 {
                l++
            } else if s > 0 {
                r--
            } else {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            }
        }
    }
    return res
}