# 原题
https://leetcode.com/problems/two-sum/

# 算法
## 暴力解决
O(n^2)

```go
func twoSum(nums []int, target int) []int {
    cnt := len(nums)
    if cnt <= 1 {
        return nil
    }
    
    for i := 0; i < cnt - 1; i++ {
        for j := i+1; j < cnt; j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return nil
}
```

## 优化

```go
func twoSum(nums []int, target int) []int {
    cnt := len(nums)
    if cnt <= 1 {
        return nil
    }
    
    numsMap := make(map[int]int, cnt)
    for i := 0; i < cnt; i++ {
        numsMap[nums[i]] = i
    } 
    for i := 0; i < cnt; i++ {
        j := target - nums[i]
        if v, ok := numsMap[j]; ok && v != i {
            return []int{i, v}
        }
    }
    return nil
}
```