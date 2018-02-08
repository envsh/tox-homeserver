import QtQuick 2.1
import StratifyLabs.UI 3.0


Rectangle {
    id: root
    color: "green"
    width: 200
    height: 300

    // 发送给 Qt Widgets 的信号
    signal qmlSignal
    // 从 Qt Widgets 接收到的信号
    signal cSignal

        SRow {
                SButton{
                    span: 5;
            text: "111";
            }
            
            SButton{
                id: btn233
                span: 7;
                text: "222333444";
            }
           }

    SRow {
        y: btn233.height
  SLabel {
    span: 2;
    text: "4 Options";
  }
  SDropup {
    style: "block";
    span: 10;
    model:
        ["First",
      "Second",
      "Third",
      "Fourth"];
  }
}
    Text {
        id: myText
        text: "Click me"
        anchors.verticalCenterOffset: -110
        anchors.horizontalCenterOffset: -20
        font.pointSize: 14
        anchors.centerIn: parent
    }

    MouseArea {
        anchors.bottomMargin: 231
        anchors.fill: parent
        onClicked: qmlSignal()
    }

    TextInput {
        id: textInput
        x: 19
        y: 90
        width: 80
        height: 20
        text: qsTr("Text Input")
        font.pixelSize: 12
    }

    ListView {
        id: listView
        x: 45
        y: 132
        width: 110
        height: 160
        delegate: Item {
            x: 5
            width: 80
            height: 40
            Row {
                id: row1
                Rectangle {
                    width: 40
                    height: 40
                    color: colorCode
                }

                Text {
                    text: name
                    font.bold: true
                    anchors.verticalCenter: parent.verticalCenter
                }
                spacing: 10
            }
        }
        model: ListModel {
            ListElement {
                name: "Grey"
                colorCode: "grey"
            }

            ListElement {
                name: "Red"
                colorCode: "red"
            }

            ListElement {
                name: "Blue"
                colorCode: "blue"
            }

            ListElement {
                name: "Green"
                colorCode: "green"
            }
        }
    }

 }
