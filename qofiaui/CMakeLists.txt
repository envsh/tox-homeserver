project(qofiaui)
cmake_minimum_required(VERSION 3.0)
set(CMAKE_VERBOSE_MAKEFILE on)
set(CMAKE_CXX_FLAGS "-g -O0 -std=c++11")
#set(CMAKE_CXX_FLAGS "-O0 -std=c++11")
set(CMAKE_CXX_COMPILER "clang++")
set(CMAKE_BUILD_TYPE "Debug")

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)

find_package(Qt5 COMPONENTS Core REQUIRED)
find_package(Qt5 COMPONENTS Gui REQUIRED)
find_package(Qt5 COMPONENTS Widgets REQUIRED)

qt5_wrap_ui(uisrcs
  add_friend.ui
  contact_item_view.ui
  create_room.ui
  emoji_category.ui
  emoji_panel.ui
  message_item_view.ui
# scroll_widget.ui
  untitled.ui
  )
#qt5_add_resources(uircs rcc.qrc)

add_library(qofiaui SHARED qofiaui.cpp mainwin.cpp contact_item.cpp message_item.cpp
  frontend.cpp ${uisrcs} ${uircs})
target_link_libraries(qofiaui Qt5::Widgets)

add_executable(demo demo.cpp)
target_link_libraries(demo qofiaui)
