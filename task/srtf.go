package task

import "sort"

type byEstimatedTimeRemaining []*Task

func (a byEstimatedTimeRemaining) Len() int {
	return len(a)
}

func (a byEstimatedTimeRemaining) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (a byEstimatedTimeRemaining) Less(i, j int) bool {
	return a[i].EstimatedTimeRemaining < a[j].EstimatedTimeRemaining
}

func (t Tasks) SRTFSort() error {
	sort.Sort(byEstimatedTimeRemaining(t))
	return nil
}
