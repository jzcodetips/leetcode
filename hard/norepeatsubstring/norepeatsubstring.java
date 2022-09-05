import java.util.*;
import static java.lang.Math.max;

class NoRepeatSubstring {
	public static int findLength(String str) {
		int windowStart = 0, longest = 0;
		HashMap<Character, Integer> hm = new HashMap();

		for (int windowEnd = 0; windowEnd < str.length(); windowEnd++) {
			char rightChar = str.charAt(windowEnd);

			if (hm.containsKey(rightChar)) {
				windowStart = Math.max(windowStart, hm.get(rightChar) + 1);
			}

			hm.put(rightChar, windowEnd);

			longest = Math.max(longest, windowEnd - windowStart + 1);
		}
		
		return longest;
	}

	public static void main(String[] args) {
		System.out.println("length is: " + NoRepeatSubstring.findLength("aabccbb"));
		System.out.println("length is: " + NoRepeatSubstring.findLength("abbbb"));
		System.out.println("length is: " + NoRepeatSubstring.findLength("abccde"));
	}
}
