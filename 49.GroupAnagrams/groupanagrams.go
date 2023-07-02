package _9_GroupAnagrams

//给你一个字符串数组,请将字母异位词组合在一起
// strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
func groupAnagrams(strs []string) [][]string{
   mp := map[[26]int][]string{}
   for _,str := range strs{
       cnt := [26]int{}
       for _,b := range str{
           cnt[b-'a']++
       }
       mp[cnt] = append(mp[cnt],str)
   }
   ans := make([][]string,0,len(mp))
   for _,v :=range mp{
       ans = append(ans,v)
   }
   return ans
}