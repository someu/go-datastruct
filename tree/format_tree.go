package tree

import (
	"math"
	"strconv"
	"strings"
)

const (
	blank = " "
)

func fillBlanks(mid []string, num int) []string {
	leftFill := int((num - len(mid)) / 2)
	rightFill := num - leftFill - len(mid)

	line := []string{}
	for ; leftFill > 0; leftFill-- {
		line = append(line, blank)
	}
	line = append(line, mid...)
	for ; rightFill > 0; rightFill-- {
		line = append(line, blank)
	}
	return line

}

func formatTreeNode(node *SearchTreeNode) [][]string {
	levels := [][]string{}
	if node == nil {
		return levels
	} else if node.Left == nil && node.Right == nil {
		levels = append(levels, strings.Split(strconv.Itoa(node.Element.(int)), ""))
	} else {
		left := formatTreeNode(node.Left)
		right := formatTreeNode(node.Right)
		leftHeight := len(left)
		rightHeight := len(right)
		leftWidth := 0
		rightWidth := 0
		if leftHeight > 0 {
			leftWidth = len(left[0])
		}
		if rightHeight > 0 {
			rightWidth = len(right[0])
		}
		halfWidth := int(math.Max(float64(rightWidth), float64(leftWidth)))

		eleArr := strings.Split(strconv.Itoa(node.Element.(int)), "")
		levels = append(levels, fillBlanks(eleArr, halfWidth*2+2+len(eleArr)))

		level1 := []string{}
		level1 = append(level1, fillBlanks([]string{}, halfWidth)...)
		if leftHeight > 0 {
			level1 = append(level1, "/")
		} else {
			level1 = append(level1, " ")
		}
		level1 = append(level1, fillBlanks([]string{}, len(eleArr))...)
		if rightHeight > 0 {
			level1 = append(level1, "\\")
		} else {
			level1 = append(level1, " ")
		}
		level1 = append(level1, fillBlanks([]string{}, halfWidth)...)
		levels = append(levels, level1)

		childHight := int(math.Max(float64(rightHeight), float64(leftHeight)))
		for i := 0; i < childHight; i++ {
			rightLevel := []string{}
			if rightHeight > i {
				rightLevel = right[i]
			}
			leftLevel := []string{}
			if leftHeight > i {
				leftLevel = left[i]
			}
			levelI := []string{}
			levelI = append(levelI, fillBlanks(leftLevel, halfWidth)...)
			levelI = append(levelI, fillBlanks([]string{}, len(eleArr)+2)...)
			levelI = append(levelI, fillBlanks(rightLevel, halfWidth)...)
			levels = append(levels, levelI)
		}
	}

	return levels
}

func FormatTree(tree *SearchTree) string {
	lines := formatTreeNode(tree.root)
	format := ""
	for _, line := range lines {
		format += strings.Join(line, "") + "\r\n"
	}
	return format
}

func formatAVLTreeNode(node *AVLTreeNode) [][]string {
	levels := [][]string{}
	if node == nil {
		return levels
	} else if node.Left == nil && node.Right == nil {
		levels = append(levels, strings.Split(strconv.Itoa(node.Element), ""))
	} else {
		left := formatAVLTreeNode(node.Left)
		right := formatAVLTreeNode(node.Right)
		leftHeight := len(left)
		rightHeight := len(right)
		leftWidth := 0
		rightWidth := 0
		if leftHeight > 0 {
			leftWidth = len(left[0])
		}
		if rightHeight > 0 {
			rightWidth = len(right[0])
		}
		halfWidth := int(math.Max(float64(rightWidth), float64(leftWidth)))

		eleArr := strings.Split(strconv.Itoa(node.Element), "")
		levels = append(levels, fillBlanks(eleArr, halfWidth*2+2+len(eleArr)))

		level1 := []string{}
		level1 = append(level1, fillBlanks([]string{}, halfWidth)...)
		if leftHeight > 0 {
			level1 = append(level1, "/")
		} else {
			level1 = append(level1, " ")
		}
		level1 = append(level1, fillBlanks([]string{}, len(eleArr))...)
		if rightHeight > 0 {
			level1 = append(level1, "\\")
		} else {
			level1 = append(level1, " ")
		}
		level1 = append(level1, fillBlanks([]string{}, halfWidth)...)
		levels = append(levels, level1)

		childHight := int(math.Max(float64(rightHeight), float64(leftHeight)))
		for i := 0; i < childHight; i++ {
			rightLevel := []string{}
			if rightHeight > i {
				rightLevel = right[i]
			}
			leftLevel := []string{}
			if leftHeight > i {
				leftLevel = left[i]
			}
			levelI := []string{}
			levelI = append(levelI, fillBlanks(leftLevel, halfWidth)...)
			levelI = append(levelI, fillBlanks([]string{}, len(eleArr)+2)...)
			levelI = append(levelI, fillBlanks(rightLevel, halfWidth)...)
			levels = append(levels, levelI)
		}
	}

	return levels
}

func FormatAVLTree(tree *AVLTree) string {
	lines := formatAVLTreeNode(tree.root)
	format := ""
	for _, line := range lines {
		format += strings.Join(line, "") + "\r\n"
	}
	return format
}
