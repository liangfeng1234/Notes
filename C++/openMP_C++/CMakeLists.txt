cmake_minimum_required(VERSION 3.25)
project(openMP_C__)

set(CMAKE_CXX_STANDARD 17)

add_executable(openMP_C__ main.cpp)
FIND_PACKAGE(OpenMP REQUIRED)
if (OPENMP_FOUND)
    message("OPENMP FOUND")
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} ${OpenMP_C_FLAGS}")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} ${OpenMP_CXX_FLAGS}")
endif ()