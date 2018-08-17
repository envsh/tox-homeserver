import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.0


RowLayout {
    width: parent.width
    Layout.fillWidth: true

    // back button
    Button {
        text: "◁ Back"
        Layout.maximumWidth: 60

        ToolTip.delay: 1000
        ToolTip.timeout: 5000
        ToolTip.visible: hovered
        ToolTip.text: qsTr("This tool tip is shown after hovering the button for a second.")
    }

    // back arrow
    Button {
        text: ""
        Layout.maximumWidth: 30
        icon.source: "../icons/barbuttonicon_back_gray64.png"
        onClicked: print(parent.height, parent.width, parent, parent.parent, parent.parent.width,  width)
    }

    HSpacerItem{}

    // label, TODO
    Label {
        text: "curwin:"
        y : (parent.height-height)/2
    }

    // combobox
    ComboBox {
        model: ["First", "Second", "Third"]
        Layout.maximumWidth: 90
    }


    // menu
    Button {
        text: "Tool"
        Layout.maximumWidth: 40
        onClicked: toolMenu.open()

        Menu {
            id: toolMenu
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

    // next arrow
    Button{
        text: ""
        Layout.maximumWidth: 30
        icon.source: "../icons/barbuttonicon_forward_gray64.png"
    }
}

