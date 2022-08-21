package main
type Run struct {
    Time int // in milliseconds
    Results string
    Failed bool
}
// Get average runtime of successful runs in seconds
func averageRuntimeInSeconds(runs []Run) float64 {
    var totalTime int
    var failedRuns int
    for _, run := range runs {
        if run.Failed {
            failedRuns++
        } else {
            totalTime += run.Time
        }
    }
    averageRuntime := float64(totalTime) / float64(len(runs) - failedRuns) / 1000
    return averageRuntime
}
