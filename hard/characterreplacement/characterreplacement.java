import java.util.*;

class CharacterReplacement {
	public static int findLength(String str, int k) {
		int windowStart = 0, maxLength = 0, maxRepeatLetterCount = 0;
		Map<Character, Integer> letterFrequencyMap = new HashMap<>();

		for (int windowEnd = 0; windowEnd < str.length(); windowEnd++) {
			char rightChar = str.charAt(windowEnd);

			letterFrequencyMap.put(rightChar, letterFrequencyMap.getOrDefault(rightChar, 0) + 1);
			maxRepeatLetterCount = Math.max(maxRepeatLetterCount, letterFrequencyMap.get(rightChar));

			if (windowEnd - windowStart + 1 - maxRepeatLetterCount > k) {
				char leftChar = str.charAt(windowStart);

				letterFrequencyMap.put(leftChar, letterFrequencyMap.get(leftChar) - 1);
				windowStart++;
			}

			maxLength = Math.max(maxLength, windowEnd - windowStart + 1);
		}

		return maxLength;
	}

	public static void main(String[] args) {
		System.out.println("length is: " + CharacterReplacement.findLength("aabccbb", 2));
		System.out.println("length is: " + CharacterReplacement.findLength("abbcb", 1));
		System.out.println("length is: " + CharacterReplacement.findLength("abccde", 1));
	}
}
