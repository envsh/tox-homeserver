project(qofiaui)
cmake_minimum_required(VERSION 3.0)
set(CMAKE_VERBOSE_MAKEFILE on)
set(CMAKE_CXX_FLAGS "-g -O0 -std=c++11")
set(CMAKE_CXX_FLAGS "-O0 -std=c++11")

set(CMAKE_AUTOUIC ON)

find_package(Qt5 COMPONENTS Core REQUIRED)
find_package(Qt5 COMPONENTS Gui REQUIRED)
find_package(Qt5 COMPONENTS Widgets REQUIRED)

qt5_wrap_ui(uisrcs
  ../qofia-ffi/add_friend.ui
  ../qofia-ffi/contact_item_view.ui
  ../qofia-ffi/create_room.ui
  ../qofia-ffi/emoji_category.ui
  ../qofia-ffi/emoji_panel.ui
  ../qofia-ffi/message_item_view.ui
#  ../qofia-ffi/scroll_widget.ui
  ../qofia-ffi/untitled.ui
  )

add_library(qofiaui SHARED qofiaui.cpp mainwin.cpp contact_item.cpp message_item.cpp ${uisrcs})
target_link_libraries(qofiaui Qt5::Widgets)

add_executable(demo demo.cpp)
target_link_libraries(demo qofiaui)
