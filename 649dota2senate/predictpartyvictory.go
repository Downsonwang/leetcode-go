package _49dota2senate


func predictPartyVictory(senate string) string  {

	  var radiant,dire []int
	  for index,_ :=range senate{
		  if senate[index] == 'R'{
			  radiant = append(radiant,index)
		  }else {
			  dire = append(dire,index)
		  }


	  }

	  for len(radiant) >0 && len(dire) > 0 {
		  if radiant[0] < dire[0]{
			  radiant = append(radiant,radiant[0]+len(senate))
		  }else{
			  dire = append(dire,dire[0]+len(senate))
		  }
		  dire = dire[1:]
		  radiant = radiant[1:]
	  }
	  if len(dire) > 0 {
		  return "Dire"
	  }
	   return "Radiant"

}