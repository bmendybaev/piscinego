#!/bin/bash

# Print the result of ls -l skipping 1 line out of 2, starting with the first one
ls -l | awk 'NR % 2 == 0'
