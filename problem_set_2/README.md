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