cmake_minimum_required(VERSION 3.25)
project(testMyLib C)

set(CMAKE_C_STANDARD 99)

#设置头文件路径
set(INC_DIR ./include)

#设置链接库路径
set(LINK_DIR ./lib)

#引入头文件
include_directories(${INC_DIR})

#引入库文件
link_directories(${LINK_DIR})

add_executable(testMyLib main.c)

#将第三方库连接在一起
target_link_libraries(testMyLib libmyLib.dll)

