#include <iostream>

#include "qofiaui.h"

extern "C"
void uion_command_c(char* cmd) {
    std::cout<<cmd<<std::endl;
}

int main(int argc, char *argv[])
{
    qofiaui_context uictx;
    uictx.uion_command = &uion_command_c;

    qofiaui_main(&uictx);
    return 0;
}
