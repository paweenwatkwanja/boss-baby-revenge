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
        - From my solution, I think of the best case scenario after I read the conditions: the boss must not shoot first; the shot must be revenged. These could be translated into two conditions that I can evaluate first. For the first condition, the sequence must not start with 'R' and for the second one, the sequence must not end with 'S'. 
        
        Therefore, I use slicing method to achieve these and return the output right after either of these conditions is met. This results in O(1).
    - O(n)
        - Even though I try to figure how to use a solution with less time than O(n), I could not think of any solution due to the nature of how this sequence should be processed. That is, there is no clear separation which I can use to process it separately. 
        
        That's why my solution is to traverse through the sequence. Using loop uses computation resources and I do not want to waste it. Therefore, I come up with an idea of evaluating the conditions while traversing is in process. 

        Therefore, the time complexity is O(n) for both average and worst cases.

2. space complexity
    - O(n)
        - From my solution, the input is a string and its size and memory usage depend upon the size of the input. Therefore, the space complexiy of the input is O(n)

    - O(1)
        - This complexity is for the logic while processing the input. As it uses simple data structure such as a variable with a data type of int for the counter, this kind of variable does not use many resources for each character of the input. The resources do not increase if the input size changes. That's why its space complexity is O(1)

# input constraints
- I write a logic to check if the input has lower case letters and white spaces at the beginning and at the end of the input. These will be taken care and the result will be all upper case letters with no white spaces on both ends.

- The code also has a logic to check the range of the input. It must be between 1 and 1,000,000

- For the last logic for validating the input, I intentionally add the logic checking the incorrect elements inside a sequence (apart from 'S' and 'R') to reduce the usage of computation resources. Therefore, while the code is in process checking the shots, if the code finds any character other than 'S' and 'R', it will throw an error because this could be from some user error who accidentally compromises the data.