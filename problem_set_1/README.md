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
            