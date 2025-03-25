#!/bin/bash
g++ -std=c++11 -Wall main.c++ -o myfs -lfuse -D_FILE_OFFSET_BITS=64