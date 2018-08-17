import QtQuick 2.0
import QtQuick.Window 2.0
import QtQuick.Controls 2.3


ApplicationWindow {
    title: qsTr("Hello World")
    objectName: "approotwin"
    id: approotwin
    width: 340
    height: 580
    visible: true

    signal aaaaa(int v)
    function testb123(v) {print("called...", v)}

    header: ToolBar {
        // ...
        Topnav {}
    }

    Button {
        objectName: "tstbtn1"
        id: tstbtn1
        text: "hehhe"
        onClicked: {
            approotwin.aaaaa(123)
        }
    }

    StackView {
        objectName: "rootsv1"
        id: stack
        //         initialItem: "PageMessages.qml"
        initialItem: "gomdltst.qml"
        anchors.fill: parent
    }

    Component {
        id: itemComponent

        Item {
            Component.onDestruction: print("Destroying second item")
        }
    }

    Component {
        id: mainView


        Row {
            spacing: 10

            Button {
                text: "Push"
                onClicked: {

                    stack.push(mainView)

                }
            }
            Button {
                text: "Pop"
                enabled: stack.depth > 1
                onClicked: stack.pop(mainView)

            }
            Text {
                text: stack.depth
            }
        }
    }
}


