# version make
cmake_minimum_required(VERSION 3.5)

# name project
project(LoggerService)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -Wall -DHAVE_INTTYPES_H -DHAVE_NETINET_IN_H")

set(THRIFT_DIR "/usr/local/include/thrift")
set(BOOST_DIR "/usr/include/boost/")

# add file header
include_directories(${THRIFT_DIR} ${BOOST_DIR} ${CMAKE_SOURCE_DIR})
link_directories(/usr/local/lib)

set(BASE_SOURCE_FILES
        gen-cpp/Logger.cpp
        gen-cpp/LoggerService_types.cpp
        gen-cpp/LoggerService_constants.cpp)
set(SERVER_FILES
        server/LoggerServer.cpp)
set(CLIENT_FILES
        client/LoggerClient.cpp)

add_executable(LoggerServer ${SERVER_FILES} ${BASE_SOURCE_FILES})
add_executable(LoggerClient ${CLIENT_FILES} ${BASE_SOURCE_FILES})

target_link_libraries(LoggerServer thrift)
target_link_libraries(LoggerClient thrift)