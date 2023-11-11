func twoSum(nums []int, target int) []int {
    mymap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        j, ok := mymap[target-nums[i]]
        if ok {
            result := []int{j, i}
            return result
        }
        mymap[nums[i]] = i
    }
    result := []int{-1, -1}
    return result
}
// func merge(a []int, b []int) []int{
//     p1 := 0
//     p2 := 0

//     c := make([]int, len(a) + len(b))
//     for (p1 < len(a) && p2 < len(b)) {
//         if (a[p1] < b[p2]){
//             c = append(c, a[p1])
//             p1 = p1 + 1 
//         } else {
//             c = append(c, a[p2])
//             p2 = p2 + 1 
//         }
//     }

//     for p1 < len(a) {
//         c = append(c, a[p1])
//         p1 = p1 + 1
//     }

//     for p2 < len(a) {
//        c =  append(c, a[p2])
//         p2 = p2 + 1 
//     }

//     return c
// }

// func sort(nums []int) []int{
//     if (len(nums) == 0 || len(nums) == 1) {
//         return nums
//     } else if len(nums) == 2 {
//         if (nums[0] > nums[1]) {
//             res := make([]int, 2)
//             res[0] = nums[1]
//             res[1] = nums[0]
//             return res
//         }
//         return nums
//     }

//     mid := len(nums)/2
//     a := sort(nums[:mid])
//     b := sort(nums[mid:])
//     return merge(a, b)
// }

