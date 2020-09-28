package week208

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a 
    }
}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
    if boardingCost * 4 < runningCost{
        return -1
    }
    var income int 
    var cost int 
    var waitPeople int 
    var maxProfit int 
    var lastMaxProfit int 
    var runPeople = 4 
    var round int 
    var maxRound int 
    for waitPeople != 0 || round < len(customers) {
        if round < len(customers) {
            waitPeople += customers[round]
        }
        upCount := min(runPeople, waitPeople)
        waitPeople -= upCount
        income += upCount * boardingCost
        cost += runningCost
        maxProfit = income- cost 
        if maxProfit > lastMaxProfit{
            maxRound = round + 1
            lastMaxProfit = maxProfit
        }
        round += 1
    }
    if lastMaxProfit == 0 {
        maxRound = -1
    }
    return maxRound
}
