cmake_minimum_required(VERSION 3.25)
project(myLib C)

set(CMAKE_C_STANDARD 99)

add_library(myLib SHARED library.c)
