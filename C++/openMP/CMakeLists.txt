cmake_minimum_required(VERSION 3.25)
project(openMP C)

set(CMAKE_C_STANDARD 11)

add_executable(openMP main.c openMP.h)
# openMP 配置
FIND_PACKAGE(OpenMP REQUIRED)
if (OPENMP_FOUND)
    message("OPENMP FOUND")
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} ${OpenMP_C_FLAGS}")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} ${OpenMP_CXX_FLAGS}")
endif ()
