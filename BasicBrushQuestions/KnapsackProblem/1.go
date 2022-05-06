package main

import ("fmt")



func knapsackProblem(w[]int ,v[]int, capacity int ){
	for i:=0;i<len(v);i++{
		map[w[i]]=v[i]
	}
	dp:=[][2]int{}
	if w[0]>capacity{
		dp[0][0]=0 
		dp[0][1]=0
	}else{
		dp[0][0]=v[0]
		dp[0][1]=w[0]
	}
	for i=1;i<len(v);i++{
		if dp[i][1]< capacity{
			dp[i][0]=dp[i-1]+v[i]
			dp[i][1]=dp[i-1]+w[i]
		}

		


	}

}

func main(){
	w:=[]int{ }
	v:=[]int{ }





}
	

