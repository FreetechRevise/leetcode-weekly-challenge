package week208

func max(a, b int)int{
    if a > b {
        return a
    } else {
        return b 
    }
}

func minOperations(logs []string) int {
    var level int 
    for _, log := range logs{
        switch log{
        case "../":
            level -= 1
            level = max(0, level)
        case "./":
            
        default:
            level += 1
        }
    }
    return level
}
