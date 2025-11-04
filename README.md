# boss-baby-revenge

# instructions
1. Choices of running the program
    1. Run with an input from the terminal
        - This is for a short string of sequence with less than a thousand characters

    2. Run with an input from the sequence.txt file for a very long sequence
        - This is for a very long string of sequence with more than a thousand characters
        - If you choose this way, please insert the sequence inside sequence.txt in the root directory
    3. The commands
        - make run
            - run the prorgam
        - make test 
            - run the test

# solution approach
- When choosing an approach to solve the problem, there are only two options in my mind. Iteration and recursion. 

After analyzing the sequence which is too random to choose recursive way due to no clear separation where to divide the sequence into smaller chunks. I choose iteration with processing the logic along the way with one level of loop.

As I do not want to waste the usage of this loop, I come up with a solution which allows me to validate the sequence in a single loop without adding more logic or another loop to process. The solution is to use the counter. This solution will solve the problem even the shots are fired for a long period of time before revenge shots are fired.

# solution explanation
- There are two main blocks
- The first one is checking the first and last characters in the sequence.
    - The sequence must not start with 'R' because it breaks the condition which does not allow the boss to take a revenge first.
    - Also, the sequence must not end with 'S', meaning that the shot is not revenged. This breaks one of the rule as well.
    - If either of the condiitions is met, return 'Bad boy' and exit the program
- The second part is a loop checking each character in each iteration.
    - In each iteration, the program starts with getting the current character by using slicing method.
    
    Then the validation begins.

    The first condition checks if the current character is 'S' or not, 
        if yes, add one to unrevengedCount. 
        if no, then check if it is 'R' and unrevengedCount is not 0.
        then the last condition when two conditions above are not met. This mean there is a special character in the sequence. Then program exits

# time & space explanation
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