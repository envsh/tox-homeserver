package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"

func init() { qtcore.KeepMe() }
func init() { qtgui.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }

//  header block end

//  struct block begin
func NewUi_MainWindow() *Ui_MainWindow {
	return &Ui_MainWindow{}
}

type Ui_MainWindow struct {
	Centralwidget      *qtwidgets.QWidget
	VerticalLayout     *qtwidgets.QVBoxLayout
	HorizontalLayout   *qtwidgets.QHBoxLayout
	Label              *qtwidgets.QLabel
	Label_2            *qtwidgets.QLabel
	Label_3            *qtwidgets.QLabel
	Label_4            *qtwidgets.QLabel
	Label_7            *qtwidgets.QLabel
	Label_8            *qtwidgets.QLabel
	PlainTextEdit      *qtwidgets.QPlainTextEdit
	HorizontalLayout_2 *qtwidgets.QHBoxLayout
	Label_5            *qtwidgets.QLabel
	ComboBox           *qtwidgets.QComboBox
	LineEdit           *qtwidgets.QLineEdit
	Label_6            *qtwidgets.QLabel
	HorizontalLayout_3 *qtwidgets.QHBoxLayout
	LineEdit_2         *qtwidgets.QLineEdit
	PushButton         *qtwidgets.QPushButton
	Menubar            *qtwidgets.QMenuBar
	Statusbar          *qtwidgets.QStatusBar
	MainWindow         *qtwidgets.QMainWindow
	SizePolicy         *qtwidgets.QSizePolicy
	SizePolicy1        *qtwidgets.QSizePolicy
}

//  struct block end

//  setupUi block begin
// void setupUi(QMainWindow *MainWindow)
func (this *Ui_MainWindow) SetupUi(MainWindow *qtwidgets.QMainWindow) {
	this.MainWindow = MainWindow
	// { // 126
	if MainWindow.ObjectName() == "" {
		MainWindow.SetObjectName("MainWindow")
	}
	MainWindow.Resize(644, 600)
	this.Centralwidget = qtwidgets.NewQWidget(this.MainWindow, 0)        // 111
	this.Centralwidget.SetObjectName("Centralwidget")                    // 112
	this.VerticalLayout = qtwidgets.NewQVBoxLayout_1(this.Centralwidget) // 111
	this.VerticalLayout.SetObjectName("VerticalLayout")                  // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()                   // 111
	this.HorizontalLayout.SetObjectName("HorizontalLayout")              // 112
	this.Label = qtwidgets.NewQLabel(this.Centralwidget, 0)              // 111
	this.Label.SetObjectName("Label")                                    // 112

	this.HorizontalLayout.Layout().AddWidget(this.Label) // 115

	this.Label_2 = qtwidgets.NewQLabel(this.Centralwidget, 0)                                                  // 111
	this.Label_2.SetObjectName("Label_2")                                                                      // 112
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(this.Label_2) // 115

	this.Label_3 = qtwidgets.NewQLabel(this.Centralwidget, 0) // 111
	this.Label_3.SetObjectName("Label_3")                     // 112

	this.HorizontalLayout.Layout().AddWidget(this.Label_3) // 115

	this.Label_4 = qtwidgets.NewQLabel(this.Centralwidget, 0)                                                  // 111
	this.Label_4.SetObjectName("Label_4")                                                                      // 112
	this.Label_4.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(this.Label_4) // 115

	this.Label_7 = qtwidgets.NewQLabel(this.Centralwidget, 0)                                                  // 111
	this.Label_7.SetObjectName("Label_7")                                                                      // 112
	this.Label_7.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(this.Label_7) // 115

	this.Label_8 = qtwidgets.NewQLabel(this.Centralwidget, 0)                                                  // 111
	this.Label_8.SetObjectName("Label_8")                                                                      // 112
	this.Label_8.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(this.Label_8) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout, 0) // 115

	this.PlainTextEdit = qtwidgets.NewQPlainTextEdit(this.Centralwidget) // 111
	this.PlainTextEdit.SetObjectName("PlainTextEdit")                    // 112

	this.VerticalLayout.Layout().AddWidget(this.PlainTextEdit) // 115

	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_2.SetObjectName("HorizontalLayout_2") // 112
	this.Label_5 = qtwidgets.NewQLabel(this.Centralwidget, 0)   // 111
	this.Label_5.SetObjectName("Label_5")                       // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Fixed, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                          // 114
	this.SizePolicy.SetVerticalStretch(0)                                            // 114
	this.SizePolicy.SetHeightForWidth(this.Label_5.SizePolicy().HasHeightForWidth()) // 114
	this.Label_5.SetSizePolicy(this.SizePolicy)                                      // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.Label_5) // 115

	this.ComboBox = qtwidgets.NewQComboBox(this.Centralwidget) // 111
	this.ComboBox.SetObjectName("ComboBox")                    // 112

	this.HorizontalLayout_2.Layout().AddWidget(this.ComboBox) // 115

	this.LineEdit = qtwidgets.NewQLineEdit(this.Centralwidget) // 111
	this.LineEdit.SetObjectName("LineEdit")                    // 112
	this.SizePolicy1 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy1.SetHorizontalStretch(0)                                           // 114
	this.SizePolicy1.SetVerticalStretch(0)                                             // 114
	this.SizePolicy1.SetHeightForWidth(this.LineEdit.SizePolicy().HasHeightForWidth()) // 114
	this.LineEdit.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.LineEdit) // 115

	this.Label_6 = qtwidgets.NewQLabel(this.Centralwidget, 0) // 111
	this.Label_6.SetObjectName("Label_6")                     // 112

	this.HorizontalLayout_2.Layout().AddWidget(this.Label_6) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout_2, 0) // 115

	this.HorizontalLayout_3 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_3.SetObjectName("HorizontalLayout_3")  // 112
	this.LineEdit_2 = qtwidgets.NewQLineEdit(this.Centralwidget) // 111
	this.LineEdit_2.SetObjectName("LineEdit_2")                  // 112

	this.HorizontalLayout_3.Layout().AddWidget(this.LineEdit_2) // 115

	this.PushButton = qtwidgets.NewQPushButton(this.Centralwidget) // 111
	this.PushButton.SetObjectName("PushButton")                    // 112

	this.HorizontalLayout_3.Layout().AddWidget(this.PushButton) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout_3, 0) // 115

	this.MainWindow.SetCentralWidget(this.Centralwidget)      // 114
	this.Menubar = qtwidgets.NewQMenuBar(this.MainWindow)     // 111
	this.Menubar.SetObjectName("Menubar")                     // 112
	this.Menubar.SetGeometry(0, 0, 644, 29)                   // 114
	this.MainWindow.SetMenuBar(this.Menubar)                  // 114
	this.Statusbar = qtwidgets.NewQStatusBar(this.MainWindow) // 111
	this.Statusbar.SetObjectName("Statusbar")                 // 112
	this.MainWindow.SetStatusBar(this.Statusbar)              // 114

	this.RetranslateUi(MainWindow)

	qtcore.QMetaObject_ConnectSlotsByName(MainWindow) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QMainWindow *MainWindow)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MainWindow) RetranslateUi(MainWindow *qtwidgets.QMainWindow) {
	// noimpl: {
	this.MainWindow.SetWindowTitle(qtcore.QCoreApplication_Translate("MainWindow", "MainWindow", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("MainWindow", "name:", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", ",..", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "status:", "dummy123", 0))
	this.Label_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "127.0.0.1:8089", "dummy123", 0))
	this.Label_8.SetText(qtcore.QCoreApplication_Translate("MainWindow", "---", "dummy123", 0))
	this.Label_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Contact:", "dummy123", 0))
	this.Label_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "select one usable contact for send message", "dummy123", 0))
	this.PushButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "send", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_MainWindow2() *Ui_MainWindow {
	this := &Ui_MainWindow{}
	w := qtwidgets.NewQMainWindow(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_MainWindow) QWidget_PTR() *qtwidgets.QWidget {
	return this.MainWindow.QWidget_PTR()
}

//  done block end
