package numswithconsecutivesamediff

/**
* Each level contains the numbers that are of the same amount of digits. Also, each level corresponds to the solutions
* with a specific number of digits.
* level 1: [1, 2, 3, 4, 5, 6, 7, 8, 9]
* level 2: [18, 29, 70, 81, 92]
* level 3: [181, 292, 707, 818, 929]
* For example, given N=3 and K=7, at the first level, we would have potentially 9 candidates
* (i.e. [1, 2, 3, 4, 5, 7, 8, 9]). When we move on to the second level, the candidates are reduced down to
* [18, 29, 70, 81, 92]. Finally, at the last level, we would have the solutions as [181, 292, 707, 818, 929].
* Algorithm
* Here are a few steps to implement the BFS algorithm for this problem.
* - We could implement the algorithm with nested two-levels loops, where the outer loop iterates through levels and
* the inner loop handles the elements within each level.
* - We could use a list data structure to keep the numbers for a single level, i.e. here we name the variable as queue
* - For each number in the queue, we could apply the same logics as in the DFS approach, except the last step,
* rather than making a recursive call for the next number we simply append the number to the queue for the next level.
* Complexity Analysis
* Let N be the number of digits for a valid combination, and K be the difference between digits.
* - Time Complexity: O(2^N)
*     Essentially with the BFS approach, all the intermeidate candidates form a binary tree, same as the execution
*     tree as in the DFS approach. Only this time, we traverse in a breadth-first manner, rather than the depth-first.
*     Therefore, the overall time complexity of the algorithm would be O(2^{N})
* - Space Complexity: O(2^{N})
*     We use two queues to maintain the intermediate solutions, which contain no more than two levels of elements.
*     The number of elements at the level of i is up to 9 * 2^{i-1}
*     To sum up, the space complexity of the algorithm would be {O}(9 * 2^{N-1} + 9 * 2^{N-2}) = O(2^N).
 */
func numsSameConsecDiff(n int, k int) []int {
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	queue := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for level := 0; level < n-1; level++ {
		nextQueue := []int{}

		for _, num := range queue {
			tailDigit := num % 10

			// using set() to avoid duplicates when K == 0
			nextDigits := map[int]bool{tailDigit + k: true, tailDigit - k: true}

			for nextDigit := range nextDigits {
				if 0 <= nextDigit && nextDigit < 10 {
					newNum := num*10 + nextDigit
					nextQueue = append(nextQueue, newNum)
				}
			}
		}

		queue = nextQueue
	}

	return queue
}
