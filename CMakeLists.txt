set( executable main )

# Minimum required CMake version
cmake_minimum_required(VERSION 3.10)

# Project name and language
project(DSA)

# Set the source directory
set(SRC_DIR "${CMAKE_SOURCE_DIR}/src")

# Collect all .c files in the src directory
file(GLOB_RECURSE SOURCES "${SRC_DIR}/*.c")

# Add the executable
add_executable(${executable} ${SOURCES})

# Include directories if headers are in ./src or elsewhere
include_directories(${SRC_DIR})

# add_custom_command(TARGET ${executable} POST_BUILD
#                    COMMAND make
#                    COMMAND ./${executable}
#                    VERBATIM
# )

add_custom_target(build_and_run
    COMMAND ${CMAKE_MAKE_PROGRAM} # This triggers the Makefile
    COMMAND ./main                # Run the resulting binary
    WORKING_DIRECTORY ${CMAKE_BINARY_DIR} # Set the working directory to the build directory
    COMMENT "Building with Makefile and running the program..."
    VERBATIM
)
