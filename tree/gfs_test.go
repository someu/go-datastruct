package tree_test

import (
	"go-datastruct/tree"
	"testing"
)

func TestGFS(t *testing.T) {
	st := tree.NewSearchTree()

	st.Insert(60)
	st.Insert(50)
	st.Delete(60)
	st.Insert(30)
	st.Insert(80)
	st.Insert(15)
	st.Insert(10)
	st.Insert(20)
	st.Insert(40)
	st.Insert(60)
	st.Insert(55)
	st.Insert(90)
	st.Insert(70)
	st.Insert(85)
	st.Insert(85)

	t.Logf("Origin: \r\n" + tree.FormatTree(st))

	preOrderTargets := []int{50, 30, 15, 10, 20, 40, 80, 60, 55, 70, 90, 85}
	middleOrderTargets := []int{10, 15, 20, 30, 40, 50, 55, 60, 70, 80, 85, 90}
	postOrderTargets := []int{10, 20, 15, 40, 30, 55, 70, 60, 85, 90, 80, 50}

	rPreResults := []int{}
	rMiddleResults := []int{}
	rPostResults := []int{}

	preResults := []int{}
	middleResults := []int{}
	postResults := []int{}

	st.DFSRecursionPreOrder(func(i int) {
		rPreResults = append(rPreResults, i)
	})
	st.DFSRecursionMiddleOrder(func(i int) {
		rMiddleResults = append(rMiddleResults, i)
	})
	st.DFSRecursionPostOrder(func(i int) {
		rPostResults = append(rPostResults, i)
	})
	st.DFSPreOrder(func(i int) {
		preResults = append(preResults, i)
	})
	st.DFSMiddleOrder(func(i int) {
		middleResults = append(middleResults, i)
	})
	st.DFSPostOrder(func(i int) {
		postResults = append(postResults, i)
	})

	validateTargets("DFSRecursionPreOrder", preOrderTargets, rPreResults, t)
	validateTargets("DFSPreOrder", preOrderTargets, preResults, t)
	validateTargets("DFSRecursionMiddleOrder", middleOrderTargets, rMiddleResults, t)
	validateTargets("DFSMiddleOrder", middleOrderTargets, middleResults, t)
	validateTargets("DFSRecursionPostOrder", postOrderTargets, rPostResults, t)
	validateTargets("DFSPostOrder", postOrderTargets, postResults, t)

	t.Log(rPreResults)
	t.Log(middleResults)
	t.Log(rPostResults)

}
