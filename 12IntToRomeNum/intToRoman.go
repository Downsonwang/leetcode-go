/*
 * @Descripttion:
 * @Author:
 * @Date: 2022-05-08 18:27:04
 * @LastEditTime: 2022-06-01 09:52:49
 */
package inttoromenum

func intToRoman(num int) string {
	res := ""
	if num < 1 && num > 3999 {
		return ""
	}
	option := map[int]string{
		1:    "I",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	indexs := []int{1000, 900, 500, 400, 100, 90, 40, 9, 4, 1}
	//num = 58
	for index, _ := range indexs {
		if index <= num {
			ts := num / index
			num %= index

			for i := 0; i < ts; i++ {
				res += option[i]
			}

		}
	}
	return res
}
