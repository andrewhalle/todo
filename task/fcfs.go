package task

import "sort"

type byArrivalTime []*Task

func (a byArrivalTime) Len() int {
	return len(a)
}

func (a byArrivalTime) Swap(i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func (a byArrivalTime) Less(i, j int) bool {
	return a[i].ArrivalTime.Before(a[j].ArrivalTime)
}

func (t Tasks) FCFSSort() error {
	sort.Sort(byArrivalTime(t))
	return nil
}
