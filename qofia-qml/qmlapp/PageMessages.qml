import QtQuick 2.10
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.0

ColumnLayout {
    Layout.fillWidth: true

    RowLayout{
        Button{
            text:"ICON"
            Layout.maximumWidth: 40
        }
        ColumnLayout{
            RowLayout{
                Label{
                    text: "ctname"
                        Layout.fillWidth: true
                }
                Label{
                    text: "num0"
                }
            }
            Label{
                text: "ctstmsg"
                Layout.fillWidth: true
            }
        }
        ColumnLayout{
            Button{
                text: "mic"
                Layout.maximumWidth: 20
                Layout.maximumHeight: 20
            }
            Button{
                text: "mute"
                Layout.maximumWidth: 20
                Layout.maximumHeight: 20
            }
        }
        Button{
            text: "ACall"
            Layout.maximumWidth: 40
        }
        Button{
            text:"VCall"
            Layout.maximumWidth: 40
        }
        Button{
            text: "Opti"
            Layout.maximumWidth: 40

            Menu {
                id: optimenu
                y: parent.height

                MenuItem{
                    text: "111"
                }
                MenuItem {
                    text: "222"
                }
                MenuItem{
                    text: "333"
                }
            }
        }
    }

    // content
    RowLayout{
        HSpacerItem{}
        Label{
            text: "num1"
        }
        Button{
            text: "load more"
            Layout.maximumWidth: 40
            Layout.maximumHeight: 20
        }
        HSpacerItem{}
    }

        ListModel {
            id: lstmdl1
            objectName: "lstmdl1"

            ListElement { name: "Mercury"; surfaceColor: "gray" }
            ListElement { name: "Venus"; surfaceColor: "yellow" }
            ListElement { name: "Earth"; surfaceColor: "blue" }
            ListElement { name: "Mars"; surfaceColor: "orange" }
            ListElement { name: "Jupiter"; surfaceColor: "orange" }
            ListElement { name: "Saturn"; surfaceColor: "yellow" }
            ListElement { name: "Uranus"; surfaceColor: "lightBlue" }
            ListElement { name: "Neptune"; surfaceColor: "lightBlue" }
        }

      ScrollView   {

          clip: true
          ScrollBar.vertical.minimumSize: 0.03
    ScrollBar.vertical.interactive: true
    ScrollBar.horizontal.interactive: false
    ScrollBar.vertical.policy: ScrollBar.AlwaysOn
    ScrollBar.horizontal.policy: ScrollBar.AlwaysOn
    Layout.fillHeight: true
    Layout.fillWidth: true

            ListView {
            model:lstmdl1
                    //             delegate: MessageItem{}
                                     delegate: ContactItem{seqno:index}
            }
    }

    RowLayout{
        Layout.fillHeight: false
        Rectangle{
            Layout.fillHeight: true
            Layout.fillWidth: true
            color: "blue"
        }
    }

    // footer
    RowLayout{
        Button{
            text:"Attac"
            Layout.maximumWidth: 40
        }
        Button{
            text: "Snap"
            Layout.maximumWidth: 40
        }
        TextEdit{
            text: "aaa"
            selectByMouse: true
            Layout.fillWidth: true
        }
        Button{
            text:"emoji"
            Layout.maximumWidth: 40
        }
        Button{
            text: "send"
            Layout.maximumWidth: 40
        }
    }

    id: _PageMessages1
    objectName: "_PageMessages1"
    function msgmdl_add() {console.log("called...")}
    function msgmdl_delete() {print("called...")}
    function msgmdl_update() {print("called...")}
}
