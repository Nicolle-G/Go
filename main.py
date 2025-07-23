word = input("Give me a word:\n")

def Is_Palindromo(word):

   invest = ""
   for i in reversed(word):
        invest += i
   return word == invest

if Is_Palindromo(word):
   print("The word is Palindrome")
else:
   print("The word Not is a Palindrome")
