Doing a DATA STRUCTURES book for fun!!! (Cory Althoff's Self-Taught Computer Scientist)

    \\XXXXXX//
      XXXXXXXX
     //XXXXXX\\                      OOOOOOOOOOOOOOOOOOOO
    ////XXXX\\\\                     OOOOOOOOOOOOOOOOOOOO
   //////XX\\\\\\     |||||||||||||||OOOOOOOOOOOOOOOVVVVVVVVVVVVV
  ////////\\\\\\\\    |!!!|||||||||||OOOOOOOOOOOOOOOOVVVVVVVVVVV'
 ////////  \\\\\\\\ .d88888b|||||||||OOOOOOOOOOOOOOOOOVVVVVVVVV'
////////    \\\\\\\d888888888b||||||||||||            'VVVVVVV'
///////      \\\\\\88888888888||||||||||||             'VVVVV'
//////        \\\\\Y888888888Y||||||||||||              'VVV'
/////          \\\\\\Y88888Y|||||||||||||| .             'V'
////            \\\\\\|iii|||||||||||||||!:::.            '
///              \\\\\\||||||||||||||||!:::::::.
//                \\\\\\\\           .:::::::::::.
/                  \\\\\\\\        .:::::::::::::::.
                    \\\\\\\\     .:::::::::::::::::::.
                     \\\\\\\\

The book is in Python but I will use Golang because I can

...Anywho here are my notes enjoy!

(If notes are specific to an algo/snippet, they'll be in the corresponding src file!)

=== ASCII ================================================================================

7 bit number codes, 2^7 = 128 chars
  sometimes extended to 8 to cover 2^8 = 256 chars
  
UTF8 is just 32 bit number codes
  2^32 = a LOT of char codes!
  it is a superset of ASCII, meaning that an ASCII code is always the same UTF8 code

=== SORTING (General) ====================================================================

DO NOT WRITE YOUR OWN SORT

a STABLE SORT doesn't disturb sequences other than the sort key
  e.g. if we sort ["akita", "bob", "albatross"] by first letter,
       we will ALWAYS get ["akita", "albatross", "bob"]

       akita will remain BEFORE albatross even if they're now together

an UNSTABLE sort would not guarantee that akita ALWAYS remains before albatross

!!! PYTHON uses TIMSORT which is an INSERTION SORT + MERGESORT hybrid
       
=== INTERESTING QUESTIONS... =============================================================

> why are Golang's random numbers so repeated?
  you'll see many of the same if you randn a bunch of stuff into an array?

=== COMMON STRING QUESTIONS ==============================================================

ANAGRAM -> str a has the same letters as str b
  DETECTION: sort the strings then compare
  RUNTIME: o(n log n)
    
PALINDROME -> str reads the same both ways
  DETECTION: reverse the string then check if it's same as original
  RUNTIME: o(n) since reversing is slowest operation and is o(n)

LAST DIGIT: get the rightmost digit in a string
  SOLUTION: filter out characters then take last element in digits
            ^ using PYTHON LIST COMPREHENSION

  -- GOLANG IMPLEMENTATION --
  
  [x] ANAGRAM
  [x] PALINDROME
  [ ] LAST DIGIT
