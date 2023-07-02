package _0regularexpressionmatching



func isMatch(s string, p string) bool {
   if s == p {
      return true
   }
   for len(s) != 0 || len(p) != 0 {
      if len(s) != len(p) {
         return false
      }
      
   }
   return true
}