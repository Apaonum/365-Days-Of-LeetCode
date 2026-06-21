/*
1833. Maximum Ice Cream Bars | Medium
It is a sweltering summer day, and a boy wants to buy some ice cream bars.

At the store, there are n ice cream bars. You are given an array costs of length n, where costs[i] is the price of the ith ice cream bar in coins. The boy initially has coins coins to spend, and he wants to buy as many ice cream bars as possible. 

Note: The boy can buy the ice cream bars in any order.

Return the maximum number of ice cream bars the boy can buy with coins coins.

You must solve the problem by counting sort.

 

Example 1:

Input: costs = [1,3,2,4,1], coins = 7
Output: 4
Explanation: The boy can buy ice cream bars at indices 0,1,2,4 for a total price of 1 + 3 + 2 + 1 = 7.
Example 2:

Input: costs = [10,6,8,7,7,8], coins = 5
Output: 0
Explanation: The boy cannot afford any of the ice cream bars.
Example 3:

Input: costs = [1,6,3,1,2,5], coins = 20
Output: 6
Explanation: The boy can buy all the ice cream bars for a total price of 1 + 6 + 3 + 1 + 2 + 5 = 18.
 

Constraints:

costs.length == n
1 <= n <= 10^5
1 <= costs[i] <= 10^5
1 <= coins <= 10^8
*/ 

package main

func maxIceCream(costs []int, coins int) int {
	// 1. หาว่าราคาไอศกรีมที่แพงที่สุดคือเท่าไหร่ เพื่อสร้างกระดานขนาดพอดี
	maxCost := 0
	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}

	// 2. สร้างกระดานเพื่อนับความถี่ของแต่ละราคา (Counting Sort Array)
	freq := make([]int, maxCost+1)
	for _, cost := range costs {
		freq[cost]++
	}

	totalBars := 0

	// 3. เริ่มไล่ซื้อจากราคาถูกที่สุด (1 บาท) ไปจนถึงราคาแพงที่สุด
	for cost := 1; cost <= maxCost; cost++ {
		count := freq[cost]

		// ถ้าราคานี้ไม่มีขาย ก็ข้ามไป
		if count == 0 {
			continue
		}

		// ถ้าเงินที่มี ซื้อไอศกรีมราคานี้ไม่ได้แม้แต่แท่งเดียว ก็เลิกซื้อ (เพราะราคาถัดๆ ไปจะยิ่งแพงกว่านี้)
		if coins < cost {
			break
		}

		// คำนวณว่าเงินที่มี สามารถเหมาไอศกรีมราคานี้ได้มากที่สุดกี่แท่ง
		// โดยเลือกค่าน้อยที่สุดระหว่าง "จำนวนแท่งที่มีขาย" กับ "จำนวนแท่งที่เงินเราซื้อไหว"
		canBuy := coins / cost
		if count < canBuy {
			canBuy = count
		}

		// จ่ายเงิน และหยิบไอศกรีมใส่กระเป๋า
		coins -= canBuy * cost
		totalBars += canBuy
	}

	return totalBars
}