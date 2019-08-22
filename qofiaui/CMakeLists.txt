project(qofiaui)
cmake_minimum_required(VERSION 3.0)
set(CMAKE_VERBOSE_MAKEFILE on)
# use -std=gnu++11 to fix pch error: trigraphs was enabled in PCH file but is currently disabled
set(CMAKE_CXX_FLAGS "-g -O0 -std=gnu++11")
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
function(GOBUILD TARGET PKGORFILES)
  set(options OPTIONAL FAST)
  set(oneValueArgs DESTINATION BINNAME)
  set(multiValueArgs DEPENDS ARGS)
  cmake_parse_arguments(GOBUILD "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN} )

  #message("BINNAME=${GOBUILD_BINNAME}")
  #message("ARGN=${GOBUILD_ARGN}")
  #message("UNPARSED=${GOBUILD_UNPARSED_ARGUMENTS},")

  foreach(gofile ${GOBUILD_UNPARSED_ARGUMENTS})
    set(PKGORFILES "${PKGORFILES} ${gofile}")
  endforeach()
  foreach(arg ${GOBUILD_ARGS})
    set(goargs "${goargs} ${arg}")
  endforeach()
  if(NOT "${GOBUILD_BINNAME}")
    set(GOBUILD_BINNAME ${TARGET})
  endif()
  set(depfiles "")
  foreach(dep ${GOBUILD_DEPENDS})
    get_target_property(deptype ${dep} TYPE)
    #message("${dep} ${deptype}")
    get_target_property(propval ${dep} OUTPUT_${BUILD_TYPE})
    #message("${dep} ${propval} ${BUILD_TYPE}" )
    set(depfile "${CMAKE_${deptype}_PREFIX}${dep}${CMAKE_${deptype}_SUFFIX}")
    #message("depfile=${depfile}")
    set(depfiles "${depfile} ${depfiles}")
  endforeach()
  string(REGEX REPLACE " $" "" depfiles "${depfiles}")
  set(depsumsfile "${TARGET}_depsums.go")

  #message("target=${TARGET}, package=${PKGORFILES}, goargs=${goargs}")
  #message("DEPENDS=${GOBUILD_DEPENDS}")

  set(godir ${PKGORFILES})
  if(NOT IS_DIRECTORY ${PKGORFILES})
    #message("not isdir")
    foreach(file ${PKGORFILES})
      get_filename_component(godir ${file} DIRECTORY)
      break()
    endforeach()
  endif()
  #message("godir=${godir}")

  set(gogrepv "")
  if (${CMAKE_VERBOSE_MAKEFILE})
    set(goargs "${goargs} -x")
    set(gogrepv "2>&1 |grep -v packagefile")
    #set(gogrepv "2>err.log |grep -v packagefile")
  endif()

  set(gomkfile "gobuild_${TARGET}.sh")
  set (${TARGET}_TEST_FLAG_FILE "${TARGET}_GOBUILD.PASS0.log")

  file(WRITE ${gomkfile} "set -o pipefail\n")
  file(APPEND ${gomkfile} "go build ${goargs} -o ${GOBUILD_BINNAME} ${PKGORFILES} ${gogrepv}\n")
  file(APPEND ${gomkfile} "maincmdret=\"$?\"\n")
  file(APPEND ${gomkfile} "#echo ${TARGET}_GOBUILD.PASS$maincmdret.log\n")
  file(APPEND ${gomkfile} "touch ${TARGET}_GOBUILD.PASS$maincmdret.log\n")

  add_custom_target(${TARGET} ALL
    COMMENT "Go target ${TARGET} ..."
    COMMAND ${CMAKE_COMMAND} -E remove -f ${${TARGET}_TEST_FLAG_FILE}
    COMMAND echo "package main" > ${depsumsfile}
    COMMAND echo "\\/\\*" >> ${depsumsfile}
    COMMAND md5sum ${depfiles} >> ${depsumsfile}
    COMMAND echo "\\*\\/" >> ${depsumsfile}

    COMMAND ${CMAKE_COMMAND} -E rename ${depsumsfile} ${godir}/${depsumsfile}
    COMMAND sh ${CMAKE_CURRENT_BINARY_DIR}/${gomkfile}
    #COMMAND sh -c "go build ${goargs} -o ${GOBUILD_BINNAME} ${PKGORFILES}"
    DEPENDS ${GOBUILD_DEPENDS}
    )
  add_custom_target(${TARGET}_errchk ALL COMMENT "Check Go build ${TARGET} ..."
    DEPENDS ${TARGET} ${${TARGET}_TEST_FLAG_FILE})
  set_property(TARGET ${TARGET} PROPERTY ADDITIONAL_MAKE_CLEAN_FILES "aaa")
  set(ADDITIONAL_MAKE_CLEAN_FILES "${ADDITIONAL_MAKE_CLEAN_FILES} ${GOBUILD_BINNAME}")
endfunction(GOBUILD)

get_cmake_property(_variableNames VARIABLES)
list (SORT _variableNames)
foreach (_variableName ${_variableNames})
#    message(STATUS "${_variableName}=${${_variableName}}")
endforeach()

gobuild(qofia2 ../qofia2/ # b.go
  ARGS -v -i -p 1
  DEPENDS qofiaui # foobar
  #BINNAME qofia2
  )

