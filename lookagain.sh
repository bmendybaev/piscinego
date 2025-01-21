#!/bin/bash

# Find all files ending with .sh, remove the .sh extension, and sort them in descending order.
find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r
