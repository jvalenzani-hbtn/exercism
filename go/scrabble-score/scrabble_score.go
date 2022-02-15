package scrabble

import (
"errors"
"unicode"
    )

func letterScore(letter rune) (int, error){
    switch letter{
        case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
    		return 1, nil
		case 'D', 'G':
    		return 2, nil
		case 'B', 'C', 'M', 'P':
    		return 3, nil
		case 'F', 'H', 'V', 'W', 'Y':
    		return 4, nil
		case 'K':
    		return 5, nil
		case 'J', 'X':
    		return 8, nil
		case 'Q', 'Z':
    		return 10, nil
        default:
    		return 0, errors.New("Invalid letter")
    }
}

func Score(word string) int {
  
  var points int = 0
  for _, letter := range word{
    v, err := letterScore(unicode.ToUpper(letter))
    if err != nil {
        break
    }
    points += v
  }
  return points
}
