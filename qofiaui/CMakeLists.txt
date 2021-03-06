project(qofiaui)
cmake_minimum_required(VERSION 3.0)
set(CMAKE_VERBOSE_MAKEFILE on)
# use -std=gnu++11 to fix pch error: trigraphs was enabled in PCH file but is currently disabled
set(CMAKE_CXX_FLAGS "-g -O0 -std=gnu++11")
#set(CMAKE_CXX_FLAGS "-O0 -std=c++11")
set(CMAKE_CXX_COMPILER "clang++")
#set(CMAKE_BUILD_TYPE "Debug")

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
# seperated uis
mainform.ui
loginform.ui
contactform.ui
chatform.ui
memberform.ui
settingform.ui
)
qt5_add_resources(uircs
  rcc.qrc
  # thefont.qrc
  )
set(cppsrcs
  qofiaui.cpp mainwin.cpp contact_item.cpp message_item.cpp
  frontend.cpp uiutils.cpp event.cpp
  mainform.cpp settingform.cpp loginform.cpp
  contactform.cpp chatform.cpp memberform.cpp
  )
add_library(qofiaui SHARED ${cppsrcs} ${uisrcs} ${uircs})
target_link_libraries(qofiaui Qt5::Widgets)

add_executable(demo demo.cpp)
target_link_libraries(demo qofiaui)

# need set -std=gnu++11
set (CMAKE_MODULE_PATH "${CMAKE_SOURCE_DIR}/../CMake")
include(PrecompiledHeader)
add_precompiled_header(qofiaui qthdrsrc.h FORCEINCLUDE)

### go build area
include(gobuild)
gobuild(qofia2 ../qofia2/ # b.go
  ARGS -v -i -p 1
  DEPENDS qofiaui # foobar
  #BINNAME qofia2
  )

get_cmake_property(_variableNames VARIABLES)
list (SORT _variableNames)
foreach (_variableName ${_variableNames})
#    message(STATUS "${_variableName}=${${_variableName}}")
endforeach()

