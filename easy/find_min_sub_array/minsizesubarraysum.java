import static java.lang.Math.min;

class MinSizeSubArraySum {
	public static int findMinSubArraySum(int s, int[] arr) {
		int windowStart = 0, smallest = Integer.MAX_VALUE, sum = 0;

		for (int windowEnd = 0; windowEnd < arr.length; windowEnd++) {
			sum += arr[windowEnd];

			while (sum >= s) {
				smallest = Math.min(smallest, windowEnd - windowStart + 1);
				sum -= arr[windowStart];
				windowStart++;
			}
		}

		return smallest == Integer.MAX_VALUE ? 0 : smallest;	
	}

	public static void main(String[] args) {
		System.out.println("min sub array sum: " +
			MinSizeSubArraySum.findMinSubArraySum(7, new int[]{3, 4, 1, 1, 6}));
		System.out.println("min sub array sum: " +
			MinSizeSubArraySum.findMinSubArraySum(7, new int[]{2, 1, 5, 2, 8}));
	}
}
