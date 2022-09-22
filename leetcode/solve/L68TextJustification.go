package main

import (
	"container/list"
	"strings"
)

// 给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
// 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。

// 挺有意思
func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	width := 0
	li := list.New()
	for _, str := range words {
		sl := len(str)
		if width == 0 {
			// 还能放下
			li.PushBack(str)
			width += sl
		} else if width+sl+1 <= maxWidth {
			// 还能放下
			li.PushBack(str)
			width += sl + 1
		} else {
			// 放不下的时候 做一波处理
			left := maxWidth - width
			liL := li.Len()
			if liL == 1 {
				// 1个就放不下了
				var sb strings.Builder
				sb.WriteString(li.Front().Value.(string))
				for j := 0; j < maxWidth-width; j++ {
					sb.WriteString(" ")
				}
				result = append(result, sb.String())
				li = list.New()
				width = 0
				li.PushBack(str)
				width += sl
				continue
			}

			var shang int
			shang = left / (liL - 1)
			yu := left % (liL - 1)

			var sb strings.Builder
			e, i := li.Front(), 0
			sb.WriteString(e.Value.(string))
			for e = e.Next(); e != nil; e = e.Next() {
				spaceNum := shang + 1
				if i < yu {
					spaceNum++
				}
				for j := 0; j < spaceNum; j++ {
					sb.WriteString(" ")
				}
				sb.WriteString(e.Value.(string))
				i++
			}
			result = append(result, sb.String())

			li = list.New()
			width = 0
			li.PushBack(str)
			width += sl
		}
	}
	if width > 0 {
		var sb strings.Builder
		e := li.Front()
		sb.WriteString(e.Value.(string))
		for e = e.Next(); e != nil; e = e.Next() {
			sb.WriteString(" ")
			sb.WriteString(e.Value.(string))
		}
		for j := 0; j < maxWidth-width; j++ {
			sb.WriteString(" ")
		}
		result = append(result, sb.String())
	}
	return result
}

// func main() {
// 	// fmt.Println(len("你好")) // 6
// 	res := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
// 	for _, e := range res {
// 		fmt.Printf("%q", e)
// 	}
// }
