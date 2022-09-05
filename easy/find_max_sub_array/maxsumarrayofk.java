import static java.lang.Math.max;

class MaxSumSubArrayOfSizeK {
	public static int findMaxSumSubArray(int k, int[] arr) {
		int windowSum = 0, maxSum = 0, windowStart = 0;

		for (int windowEnd = 0; windowEnd < arr.length; windowEnd++) {
			windowSum += arr[windowEnd];

			if (windowEnd >= k - 1) {
				maxSum = Math.max(maxSum, windowSum);
				windowSum -= arr[windowStart];
				windowStart++;
			}
		}


		return maxSum;
	}

	public static void main(String[] args) {
		System.out.println("max sum of 3: " + MaxSumSubArrayOfSizeK.findMaxSumSubArray(3, new int[]{2, 1, 5, 1, 3, 2}));
	}
}
