import QtQuick 2.0
import QtQuick.Window 2.0
import QtQuick.Controls 2.0


ApplicationWindow {
    title: qsTr("Hello World")
    width: 340
    height: 580
    visible: true

    header: ToolBar {
        // ...
        Topnav {}
    }

    StackView {
        id: stack
        initialItem: "PageMessages.qml"
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


