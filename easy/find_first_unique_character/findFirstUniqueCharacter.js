// Find the first unique character of a string. 
// 
// findFirstUniqueCharacter('aba') = > b
// findFirstUniqueCharacter('abb') = > a
// findFirstUniqueCharacter('abcb') = > a
// findFirstUniqueCharacter('abcba') = > c
// findFirstUniqueCharacter('abcabc') => null
// findFirstUniqueCharacter('fxrifxr') => i

// Brut force
function BrutForceFindFirstUniqueCharacter(text) {
  for (let i = 0; i < text.length; i++) {
    let isFound = true
    for (let j = 0; j < text.length; j++) {
      if (i !== j && text[i] === text[j]) {
        isFound = false
      }
    }

    if (isFound)
      return text[i]
  }

  return null
}

// Dictionary
function DictionaryFindFirstUniqueCharacter(text) {
  let letters = {}
  for (let i = 0; i < text.length; i++) {
    let letter = text[i]
    letters[letter] = typeof letters[letter] === "undefined" ? 1 : letters[letter] + 1
  }

  for (let i = 0; i < text.length; i++) {
    let letter = text[i]
    if (letters[letter] === 1) {
      return letter
    }
  }

  return null
}

// unittests
var tests = [
  'aba',
  'abb',
  'abcb',
  'abcba',
  'abcabc',
  'fxrifxr',
  'bba',
  'aa',
  'aarb'
];

tests.map(function(txt) {
  console.log(`Brut force ${BrutForceFindFirstUniqueCharacter(txt)}`)
  console.log(`Dictionary ${DictionaryFindFirstUniqueCharacter(txt)}`)
});