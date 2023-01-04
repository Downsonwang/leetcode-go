package qianxin


func work(difficulty []int,step []int,ability []int) int {
	         data := make([]int,0)
             for i,d := range difficulty{
				 if step[i]  > data[d]{
					 data[d] = step[i]
				 }
			 }
			 for i:=1; ;i++{
				 if data[i-1] > data[i]{
					 data[i] = data[i-1]
				 }
			 }
			 res := 0
			 for _,w :=range ability {
				 res +=data[w]
			 }
			 return res
}