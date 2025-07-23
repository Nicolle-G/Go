from main import Is_Palindromo

def test_IsPalindromo():
    assert Is_Palindromo (word ="oso") == True
def test_Not_IsPalindromo():
    assert Is_Palindromo (word ="casa" ) == False