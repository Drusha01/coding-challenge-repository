Problem Set 3 (Challenging)
Problem: Longest Increasing Subsequence
Given an unsorted array of integers, find the length of the longest increasing subsequence.
For context. The "Longest Increasing Subsequence" is a common problem in computer science and
dynamic programming. In the context of an array of integers, the goal is to find the length of the longest
subsequence of a given array such that all elements of the subsequence are sorted in increasing order.

check if the length of array is zero 
    return as 0 as maximum number of subseries
else 
    set the create temp list to get the largest series as 1
    loop the number list
        as a note that list[i] will only come after list[j], we append list[i] after every subsequence then update our temp[i]
    loop through the temp list
        get max in temp list
    return max