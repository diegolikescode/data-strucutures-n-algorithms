set( executable main )

cmake_minimum_required(VERSION 3.10)

project(DSA)

set(SRC_DIR "${CMAKE_SOURCE_DIR}/src")

file(GLOB_RECURSE SOURCES "${SRC_DIR}/*.c")

add_executable(${executable} ${SOURCES})

include_directories(${SRC_DIR})

add_custom_target(build_and_run
    COMMAND ${CMAKE_MAKE_PROGRAM}
    COMMAND ./main
    WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
    COMMENT "Building with Makefile and running the program..."
    VERBATIM
)
