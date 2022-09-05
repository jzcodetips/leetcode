import java.util.*;
import static java.lang.Math.max;

class LongestSubstringKDistinct {
	public static int findLength(String str, int k) {
		int windowStart = 0, longest = 0;
		HashMap<Character, Integer> letterOccs = new HashMap<>();

		for (int windowEnd = 0; windowEnd < str.length(); windowEnd++) {
			char rightChar = str.charAt(windowEnd);
			
			letterOccs.put(rightChar , letterOccs.getOrDefault(rightChar , 0) + 1);

			while (letterOccs.size() > k) {
				char leftChar = str.charAt(windowStart);
				
				letterOccs.put(leftChar, letterOccs.get(leftChar) - 1);

				if (letterOccs.get(leftChar) == 0) {
					letterOccs.remove(leftChar);
				}

				windowStart++;
			}

			longest = Math.max(longest, windowEnd - windowStart + 1);

		}

		return longest;
	}

	public static void main(String[] args) {
		System.out.println("length is " + LongestSubstringKDistinct.findLength("araaci", 2));
		System.out.println("length is " + LongestSubstringKDistinct.findLength("araaci", 1));
		System.out.println("length is " + LongestSubstringKDistinct.findLength("cbbebi", 3));
	}
}
