
### example
# gobuild(qofia2 ../qofia2/ # b.go
#  ARGS -v -i -p 1
#  DEPENDS qofiaui # foobar
#  #BINNAME qofia2
#  )

### go build area
function(GOBUILD TARGET PKGORFILES)
  set(options RACE LINKSHARED)
  set(oneValueArgs BINNAME BUILDMODE)
  set(multiValueArgs DEPENDS ARGS TAGS)
  cmake_parse_arguments(GOBUILD "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN} )

  # Full cmake varname: GOBUILD_BINAME

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
    set(depfile "${CMAKE_${deptype}_PREFIX}${dep}${CMAKE_${deptype}_SUFFIX}")
    set(depfiles "${depfile} ${depfiles}")
  endforeach()
  string(REGEX REPLACE " $" "" depfiles "${depfiles}")
  set(depsumsfile "${TARGET}_depsums.go")

  set(godir ${PKGORFILES})
  if(NOT IS_DIRECTORY ${PKGORFILES})
    foreach(file ${PKGORFILES})
      get_filename_component(godir ${file} DIRECTORY)
      break()
    endforeach()
  endif()

  set(gogrepv "")
  if (${CMAKE_VERBOSE_MAKEFILE})
    set(goargs "${goargs} -x")
    set(gogrepv "2>&1 |grep -v packagefile")
  endif()
  if("${CMAKE_BUILD_TYPE}" STREQUAL "Release" OR "${CMAKE_BUILD_TYPE}" STREQUAL "MinSizeRel")
    set(goargs "${goargs} -ldflags \"-w -s\"")
  endif()

  set(gomkfile "gobuild_${TARGET}.sh")
  set (${TARGET}_TEST_FLAG_FILE "${TARGET}_GOBUILD.PASS0.log")

  file(WRITE ${gomkfile} "set -o pipefail\n")
  file(APPEND ${gomkfile} "set -x\n")
  file(APPEND ${gomkfile} "go build ${goargs} -o ${GOBUILD_BINNAME} ${PKGORFILES} ${gogrepv}\n")
  file(APPEND ${gomkfile} "maincmdret=\"$?\"\n")
  file(APPEND ${gomkfile} "#echo ${TARGET}_GOBUILD.PASS$maincmdret.log\n")
  file(APPEND ${gomkfile} "touch ${TARGET}_GOBUILD.PASS$maincmdret.log\n")

  add_custom_target(${TARGET} ALL
    COMMENT "Go target ${TARGET} ..."
    COMMAND ${CMAKE_COMMAND} -E remove -f ${${TARGET}_TEST_FLAG_FILE}
    COMMAND echo "package main" > ${depsumsfile} # TODO package name not always main
    COMMAND echo "\\/\\*" >> ${depsumsfile}
    COMMAND md5sum ${depfiles} >> ${depsumsfile}
    COMMAND echo "\\*\\/" >> ${depsumsfile}

    COMMAND ${CMAKE_COMMAND} -E rename ${depsumsfile} ${godir}/${depsumsfile}
    COMMAND sh ${CMAKE_BINARY_DIR}/${gomkfile}
    #COMMAND sh -c "go build ${goargs} -o ${GOBUILD_BINNAME} ${PKGORFILES}"
    DEPENDS ${GOBUILD_DEPENDS}
    )
  add_custom_target(${TARGET}_errchk ALL COMMENT "Check Go build ${TARGET} ..."
    DEPENDS ${TARGET} ${${TARGET}_TEST_FLAG_FILE})
  set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES
    ${CMAKE_BINARY_DIR}/${GOBUILD_BINNAME}
    ${CMAKE_BINARY_DIR}/${gomkfile}
    ${CMAKE_BINARY_DIR}/${${TARGET}_TEST_FLAG_FILE}
    )
endfunction(GOBUILD)

# util func
function(DUMPVARS)
  get_cmake_property(_variableNames VARIABLES)
  list (SORT _variableNames)
  foreach (_variableName ${_variableNames})
    message(STATUS "${_variableName}=${${_variableName}}")
  endforeach()
endfunction(DUMPVARS)

