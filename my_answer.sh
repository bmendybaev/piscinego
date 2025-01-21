#!/bin/bash

# Clone the repository
git clone https://github.com/01-edu/the-final-cl-test.git the-final-cl-test >/dev/null 2>&1

# Extract the required line and print it
head -2 the-final-cl-test/final.txt | tail -1
