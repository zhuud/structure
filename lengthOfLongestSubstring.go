package structure

func LengthOfLongestSubstring(s string) int {
    n := len(s)
    m := make(map[byte]int)
    maxScroll := 0

    for i := 0; i < n; i++ {
        for j := 0; j < maxScroll + 1; j++ {
            if i + j > n - 1 {
                break
            }
            if _ , ok := m[s[i+j]]; ok {
                m = map[byte]int{}
                break
            }
            m[s[i+j]] = 0
            maxScroll = max(maxScroll, len(m))
        }
    }

    return maxScroll
}

func LengthOfLongestSubstringV2(s string) int {
    n := len(s)
    m := make(map[byte]byte)
    maxScroll := 0

    for i := 0; i < n; i++ {
        for {
            // 判断map有无重复数据
            if _ , ok := m[s[i]]; !ok {
                break
            }
            // 有记录窗口大小
            l := len(m)
            // 窗口左边界右移1
            delete(m, s[i - l])
        }
        // 无窗口数据放进map 窗口右边界右移1 结束
        m[s[i]] = 0
        maxScroll = max(maxScroll, len(m))
    }

    return maxScroll
}

func LengthOfLongestSubstringV3(s string) int {
    // 来记录一个字母如果后面出现重复时，应该调整到的新位置 i + 1 ，即字母后面的位置
    m := [128]int{0}
    longest := 0
    left := 0

    for i := 0; i < len(s); i++ {
        // 存不存在， 存在则把左边界移到存在字母后面的位置
        left = max(left, m[s[i]])
        // 记录当前字母后一位位置
        m[s[i]] = i + 1
        longest = max(longest, i-left+1)
    }

    return longest
}


func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}