import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0

ColumnLayout {
    Layout.fillWidth: true

    RowLayout {
        Button{
            text: "Icon"
            Layout.maximumWidth: 30
        }
        ColumnLayout{

            Label{
                Layout.fillWidth: true
                text: "user name"
            }
            Label{
                Layout.fillWidth: true
                text: "user stmsg"
            }
        }
        Button{
            text: "Onlst"
            Layout.maximumWidth: 30
        }
        Button{
            text: "quit"
            Layout.maximumWidth: 30
        }
    }

    // content
    RowLayout{

        TextInput{
            text: "hhhh"
            selectByMouse: true
            Layout.fillWidth: true
        }

        Button{
            text: "..."
            Layout.maximumWidth: 30
        }
        Button{
            text: "X"
            Layout.maximumWidth: 30
        }
    }

    RowLayout{
        Layout.fillHeight: true
        Rectangle{
            Layout.fillHeight: true
            Layout.fillWidth: true
            color: "blue"
        }
    }


    // footer
    RowLayout{
        HSpacerItem{}
         Button{
             text: "addfrnd"
             Layout.maximumWidth: 40
         }
         HSpacerItem{}
         Button{
             text:"addgrp"
             Layout.maximumWidth: 40
         }
         HSpacerItem{}
         Button{
             text: "filetr"
             Layout.maximumWidth: 40
         }
         HSpacerItem{}
         Button{
             text: "setts"
             Layout.maximumWidth: 40
         }
         HSpacerItem{}
    }
}

