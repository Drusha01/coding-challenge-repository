
Problem Set 1
Problem: Palindrome Pairs
Problem Description:
Given a list of unique words, your task is to find all pairs of distinct indices (i, j) in the list so that
the concatenation of the two words, i.e., words[i] + words[j], forms a palindrome.

I loop through the list of strings
    I loop second time through the list of strings
        if the index i and j is not the same
            if it is palindrome
                append the i and j to the list of indexes
            
            


Problem Set 2 (Moderately Challenging)
Problem: Valid Parentheses
Problem Description:
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is
valid. An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

I parse through the given string giving it O(N)
    checking if it is an open parenthesis or close parenthesis
        if it is open it will just append to the b string,
        else if it is an closing parenthesis 
            it will check if the last one match the closing parenthesis then removing
            else it will return as invalid parenthesis

once done parsing, it will check if the len is zero hence it is a valid parenthesis