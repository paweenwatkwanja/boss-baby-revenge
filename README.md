# boss-baby-revenge

# instructions
1. choices of running the program
    1. run with an input from the terminal
        - This is for a short string of sequence with less than a thousand characters

    2. run with an input from the sequence.txt file for a very long sequence
        - This is for a very long string of sequence with more than a thousand characters
        - if you choose this way, please insert the sequence inside sequence.txt in the root directory
    3. The commands
        - make run
            - run the prorgam
        - make test 
            - run the test
        - make benchmark
            - run the benchmark test

# explanation
1. time complexity
    - O(1)
        - It is O(1) for the best case scenario because there are conditions evaluating the input from the very beginning. From the requirements, there are two conditions: the boss must not shoot first; the shot must be revenged. These could be translated into two conditions that I can use first. For the first condition, the sequence must not start with 'R' and for the second one, the sequence must not end with 'S'. 
        
        I then use slicing method to achieve these and return the output right after either of these conditions is met. This process only use one action regardless of how large the size of the input is.
        
        Therefore, this results in O(1).
    - O(n)
        - For the average and worst case scenatios, this program has O(n) of time complexity. Even though I try to figure how to use a solution with less time than O(n), I could not think of any solution due to the nature of how this sequence should be processed. That is, there is no clear separation which I can use to process it separately. 
        
        That's why my solution is to traverse through the sequence and evaluating the conditions.

        Therefore, the time complexity is O(n) for both average and worst cases.

2. space complexity
    - O(1)
        - This solution has O(1) space complexity. As it has only a few variables only with simple data types such as an integer, this kind of variable does not use many resources in each iteration. The space resource does not increase if the input size changes. That's why its space complexity is O(1)
        - e.g.
            - there are multiple variables in the program
                - var length int (4 bytes)
                - bool of isRevengeFirstWithNoShot() (1 byte)
                - bool of isLastShotWithNoRevenge() (1 byte)
                - var unrevengedCount int (4 bytes)
                - var char string (1 byte)
            - total memory usage is 4 + 1 + 1 + 4 + 1 = 11 bytes for every iteration so it is a constant space complexity.

# input constraints
- I write a logic to check if the input has lower case letters or white spaces at the beginning and at the end of the input. These will be taken care and the result will be all upper case letters with no white spaces on both ends.

- The code also has a logic to check the range of the input. It must be between 1 and 1,000,000. If the input is outside of these numbers, the prgram will throw an error.

- For the last logic, I add the logic checking the incorrect elements inside a sequence (apart from 'S' and 'R'). If the code finds any character other than 'S' and 'R', it will throw an error.