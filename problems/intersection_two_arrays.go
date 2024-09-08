//https://leetcode.com/problems/intersection-of-two-arrays/
package problems

func intersection(nums1 []int, nums2 []int) []int {
    // Создаем map вместо HashSet
    set := make(map[int]bool)

    // Добавляем элементы из nums1 в set
    for _, num := range nums1 {
        set[num] = true
    }

    // Создаем временный слайс для хранения результатов
    var temp []int

    // Проверяем элементы из nums2
    for _, num := range nums2 {
        if set[num] {
            temp = append(temp, num)
            delete(set, num) 
        }
    }

    return temp
}