package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"

func init() { qtcore.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }

//  header block end

//  struct block begin
func NewUi_MainWindow() *Ui_MainWindow {
	return &Ui_MainWindow{}
}

type Ui_MainWindow struct {
	Actionooo                  *qtwidgets.QAction
	ActionQuit                 *qtwidgets.QAction
	Action_About               *qtwidgets.QAction
	Centralwidget              *qtwidgets.QWidget
	VerticalLayout_6           *qtwidgets.QVBoxLayout
	HorizontalLayout_7         *qtwidgets.QHBoxLayout
	ToolButton_11              *qtwidgets.QToolButton
	Label                      *qtwidgets.QLabel
	ToolButton_12              *qtwidgets.QToolButton
	StackedWidget              *qtwidgets.QStackedWidget
	Page_4                     *qtwidgets.QWidget
	VerticalLayout_13          *qtwidgets.QVBoxLayout
	QuickWidget_2              *qtquickwidgets.QQuickWidget
	Page_3                     *qtwidgets.QWidget
	VerticalLayout_11          *qtwidgets.QVBoxLayout
	QuickWidget                *qtquickwidgets.QQuickWidget
	Page                       *qtwidgets.QWidget
	VerticalLayout_5           *qtwidgets.QVBoxLayout
	HorizontalLayout           *qtwidgets.QHBoxLayout
	ToolButton_17              *qtwidgets.QToolButton
	VerticalLayout             *qtwidgets.QVBoxLayout
	Label_2                    *qtwidgets.QLabel
	Label_3                    *qtwidgets.QLabel
	ToolButton                 *qtwidgets.QToolButton
	HorizontalLayout_2         *qtwidgets.QHBoxLayout
	LineEdit                   *qtwidgets.QLineEdit
	ToolButton_2               *qtwidgets.QToolButton
	ToolButton_3               *qtwidgets.QToolButton
	ListWidget_2               *qtwidgets.QListWidget
	ListWidget                 *qtwidgets.QListWidget
	HorizontalLayout_3         *qtwidgets.QHBoxLayout
	ToolButton_4               *qtwidgets.QToolButton
	ToolButton_5               *qtwidgets.QToolButton
	ToolButton_6               *qtwidgets.QToolButton
	ToolButton_7               *qtwidgets.QToolButton
	Page_2                     *qtwidgets.QWidget
	VerticalLayout_2           *qtwidgets.QVBoxLayout
	HorizontalLayout_5         *qtwidgets.QHBoxLayout
	Label_4                    *qtwidgets.QLabel
	VerticalLayout_4           *qtwidgets.QVBoxLayout
	Label_5                    *qtwidgets.QLabel
	HorizontalLayout_4         *qtwidgets.QHBoxLayout
	Label_6                    *qtwidgets.QLabel
	Label_7                    *qtwidgets.QLabel
	VerticalLayout_7           *qtwidgets.QVBoxLayout
	ToolButton_15              *qtwidgets.QToolButton
	ToolButton_16              *qtwidgets.QToolButton
	ToolButton_13              *qtwidgets.QToolButton
	ToolButton_14              *qtwidgets.QToolButton
	ScrollArea_2               *qtwidgets.QScrollArea
	ScrollAreaWidgetContents_2 *qtwidgets.QWidget
	VerticalLayout_3           *qtwidgets.QVBoxLayout
	VerticalLayout_8           *qtwidgets.QVBoxLayout
	VerticalSpacer             *qtwidgets.QSpacerItem
	ScrollArea                 *qtwidgets.QScrollArea
	ScrollAreaWidgetContents   *qtwidgets.QWidget
	VerticalLayout_9           *qtwidgets.QVBoxLayout
	VerticalLayout_10          *qtwidgets.QVBoxLayout
	VerticalSpacer_2           *qtwidgets.QSpacerItem
	HorizontalLayout_6         *qtwidgets.QHBoxLayout
	ToolButton_8               *qtwidgets.QToolButton
	ToolButton_9               *qtwidgets.QToolButton
	ToolButton_10              *qtwidgets.QToolButton
	LineEdit_2                 *qtwidgets.QLineEdit
	ToolButton_18              *qtwidgets.QToolButton
	MainWindow                 *qtwidgets.QMainWindow
	Icon                       *qtgui.QIcon // 116
	Icon1                      *qtgui.QIcon // 116
	Icon2                      *qtgui.QIcon // 116
	Font                       *qtgui.QFont // 116
	SizePolicy                 *qtwidgets.QSizePolicy
	Icon3                      *qtgui.QIcon // 116
	Icon4                      *qtgui.QIcon // 116
	Icon5                      *qtgui.QIcon // 116
	Icon6                      *qtgui.QIcon // 116
	Icon7                      *qtgui.QIcon // 116
	Icon8                      *qtgui.QIcon // 116
	Font1                      *qtgui.QFont // 116
	Icon9                      *qtgui.QIcon // 116
	Icon10                     *qtgui.QIcon // 116
	Icon11                     *qtgui.QIcon // 116
	Icon12                     *qtgui.QIcon // 116
	Icon13                     *qtgui.QIcon // 116
	Icon14                     *qtgui.QIcon // 116
	Icon15                     *qtgui.QIcon // 116
	Icon16                     *qtgui.QIcon // 116
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
	MainWindow.Resize(396, 598)
	this.MainWindow.SetDocumentMode(false)                                 // 114
	this.Actionooo = qtwidgets.NewQAction(MainWindow)                      // 111
	this.Actionooo.SetObjectName("Actionooo")                              // 112
	this.ActionQuit = qtwidgets.NewQAction(MainWindow)                     // 111
	this.ActionQuit.SetObjectName("ActionQuit")                            // 112
	this.Action_About = qtwidgets.NewQAction(MainWindow)                   // 111
	this.Action_About.SetObjectName("Action_About")                        // 112
	this.Centralwidget = qtwidgets.NewQWidget(this.MainWindow, 0)          // 111
	this.Centralwidget.SetObjectName("Centralwidget")                      // 112
	this.VerticalLayout_6 = qtwidgets.NewQVBoxLayout_1(this.Centralwidget) // 111
	this.VerticalLayout_6.SetSpacing(0)                                    // 114
	this.VerticalLayout_6.SetObjectName("VerticalLayout_6")                // 112
	this.VerticalLayout_6.SetContentsMargins(0, 0, 0, 0)                   // 114
	this.HorizontalLayout_7 = qtwidgets.NewQHBoxLayout()                   // 111
	this.HorizontalLayout_7.SetSpacing(0)                                  // 114
	this.HorizontalLayout_7.SetObjectName("HorizontalLayout_7")            // 112
	this.ToolButton_11 = qtwidgets.NewQToolButton(this.Centralwidget)      // 111
	this.ToolButton_11.SetObjectName("ToolButton_11")                      // 112
	this.ToolButton_11.SetFocusPolicy(qtcore.Qt__NoFocus)                  // 114
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(":/icons/barbuttonicon_back_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_11.SetIcon(this.Icon)                                                                                // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.ToolButton_11) // 115

	this.Label = qtwidgets.NewQLabel(this.Centralwidget, 0)                                                  // 111
	this.Label.SetObjectName("Label")                                                                        // 112
	this.Label.SetAlignment(qtcore.Qt__AlignCenter)                                                          // 114
	this.Label.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.Label) // 115

	this.ToolButton_12 = qtwidgets.NewQToolButton(this.Centralwidget) // 111
	this.ToolButton_12.SetObjectName("ToolButton_12")                 // 112
	this.ToolButton_12.SetFocusPolicy(qtcore.Qt__NoFocus)             // 114
	this.Icon1 = qtgui.NewQIcon()
	this.Icon1.AddFile(":/icons/barbuttonicon_forward_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_12.SetIcon(this.Icon1)                                                                                   // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.ToolButton_12) // 115

	this.VerticalLayout_6.AddLayout(this.HorizontalLayout_7, 0) // 115

	this.StackedWidget = qtwidgets.NewQStackedWidget(this.Centralwidget)                // 111
	this.StackedWidget.SetObjectName("StackedWidget")                                   // 112
	this.StackedWidget.SetLineWidth(1)                                                  // 114
	this.Page_4 = qtwidgets.NewQWidget(nil, 0)                                          // 111
	this.Page_4.SetObjectName("Page_4")                                                 // 112
	this.VerticalLayout_13 = qtwidgets.NewQVBoxLayout_1(this.Page_4)                    // 111
	this.VerticalLayout_13.SetSpacing(0)                                                // 114
	this.VerticalLayout_13.SetObjectName("VerticalLayout_13")                           // 112
	this.VerticalLayout_13.SetContentsMargins(0, 0, 0, 0)                               // 114
	this.QuickWidget_2 = qtquickwidgets.NewQQuickWidget(this.Page_4)                    // 111
	this.QuickWidget_2.SetObjectName("QuickWidget_2")                                   // 112
	this.QuickWidget_2.SetResizeMode(qtquickwidgets.QQuickWidget__SizeRootObjectToView) // 114

	this.VerticalLayout_13.Layout().AddWidget(this.QuickWidget_2) // 115

	this.StackedWidget.AddWidget(this.Page_4)                                         // 115
	this.Page_3 = qtwidgets.NewQWidget(nil, 0)                                        // 111
	this.Page_3.SetObjectName("Page_3")                                               // 112
	this.VerticalLayout_11 = qtwidgets.NewQVBoxLayout_1(this.Page_3)                  // 111
	this.VerticalLayout_11.SetSpacing(0)                                              // 114
	this.VerticalLayout_11.SetObjectName("VerticalLayout_11")                         // 112
	this.VerticalLayout_11.SetContentsMargins(0, 0, 0, 0)                             // 114
	this.QuickWidget = qtquickwidgets.NewQQuickWidget(this.Page_3)                    // 111
	this.QuickWidget.SetObjectName("QuickWidget")                                     // 112
	this.QuickWidget.SetResizeMode(qtquickwidgets.QQuickWidget__SizeRootObjectToView) // 114
	this.QuickWidget.SetSource(qtcore.NewQUrl_1("qrc:/qml/area.qml", 0))              // 114

	this.VerticalLayout_11.Layout().AddWidget(this.QuickWidget) // 115

	this.StackedWidget.AddWidget(this.Page_3)                     // 115
	this.Page = qtwidgets.NewQWidget(nil, 0)                      // 111
	this.Page.SetObjectName("Page")                               // 112
	this.VerticalLayout_5 = qtwidgets.NewQVBoxLayout_1(this.Page) // 111
	this.VerticalLayout_5.SetObjectName("VerticalLayout_5")       // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()            // 111
	this.HorizontalLayout.SetSpacing(0)                           // 114
	this.HorizontalLayout.SetObjectName("HorizontalLayout")       // 112
	this.ToolButton_17 = qtwidgets.NewQToolButton(this.Page)      // 111
	this.ToolButton_17.SetObjectName("ToolButton_17")             // 112
	this.ToolButton_17.SetMaximumSize_1(32, 32)                   // 113
	this.ToolButton_17.SetFocusPolicy(qtcore.Qt__NoFocus)         // 114
	this.Icon2 = qtgui.NewQIcon()
	this.Icon2.AddFile(":/icons/icon_avatar_40.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_17.SetIcon(this.Icon2)                                                                     // 114
	this.ToolButton_17.SetIconSize(qtcore.NewQSize_1(32, 32))                                                  // 113

	this.HorizontalLayout.Layout().AddWidget(this.ToolButton_17) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()    // 111
	this.VerticalLayout.SetSpacing(0)                   // 114
	this.VerticalLayout.SetObjectName("VerticalLayout") // 112
	this.Label_2 = qtwidgets.NewQLabel(this.Page, 0)    // 111
	this.Label_2.SetObjectName("Label_2")               // 112
	this.Font = qtgui.NewQFont()
	this.Font.SetPointSize(12)                                                                       // 114
	this.Font.SetBold(true)                                                                          // 114
	this.Font.SetWeight(75)                                                                          // 114
	this.Label_2.SetFont(this.Font)                                                                  // 114
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__TextEditable | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_2) // 115

	this.Label_3 = qtwidgets.NewQLabel(this.Page, 0) // 111
	this.Label_3.SetObjectName("Label_3")            // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                          // 114
	this.SizePolicy.SetVerticalStretch(0)                                                            // 114
	this.SizePolicy.SetHeightForWidth(this.Label_3.SizePolicy().HasHeightForWidth())                 // 114
	this.Label_3.SetSizePolicy(this.SizePolicy)                                                      // 114
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__TextEditable | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_3) // 115

	this.HorizontalLayout.AddLayout(this.VerticalLayout, 0) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton.SetObjectName("ToolButton")           // 112
	this.ToolButton.SetFocusPolicy(qtcore.Qt__NoFocus)    // 114
	this.Icon3 = qtgui.NewQIcon()
	this.Icon3.AddFile(":/icons/online_30.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton.SetIcon(this.Icon3)                                                                   // 114
	this.ToolButton.SetToolButtonStyle(qtcore.Qt__ToolButtonIconOnly)                                     // 114
	this.ToolButton.SetAutoRaise(true)                                                                    // 114

	this.HorizontalLayout.Layout().AddWidget(this.ToolButton) // 115

	this.VerticalLayout_5.AddLayout(this.HorizontalLayout, 0) // 115

	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_2.SetSpacing(0)                       // 114
	this.HorizontalLayout_2.SetObjectName("HorizontalLayout_2") // 112
	this.LineEdit = qtwidgets.NewQLineEdit(this.Page)           // 111
	this.LineEdit.SetObjectName("LineEdit")                     // 112

	this.HorizontalLayout_2.Layout().AddWidget(this.LineEdit) // 115

	this.ToolButton_2 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_2.SetObjectName("ToolButton_2")         // 112
	this.ToolButton_2.SetMinimumSize_1(60, 0)               // 113
	this.ToolButton_2.SetFocusPolicy(qtcore.Qt__NoFocus)    // 114
	this.ToolButton_2.SetAutoRaise(true)                    // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.ToolButton_2) // 115

	this.ToolButton_3 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_3.SetObjectName("ToolButton_3")         // 112
	this.ToolButton_3.SetFocusPolicy(qtcore.Qt__NoFocus)    // 114
	this.Icon4 = qtgui.NewQIcon()
	this.Icon4.AddFile(":/icons/remove-symbol_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_3.SetIcon(this.Icon4)                                                                            // 114
	this.ToolButton_3.SetAutoRaise(true)                                                                             // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.ToolButton_3) // 115

	this.VerticalLayout_5.AddLayout(this.HorizontalLayout_2, 0) // 115

	this.ListWidget_2 = qtwidgets.NewQListWidget(this.Page) // 111
	this.ListWidget_2.SetObjectName("ListWidget_2")         // 112

	this.VerticalLayout_5.Layout().AddWidget(this.ListWidget_2) // 115

	this.ListWidget = qtwidgets.NewQListWidget(this.Page) // 111
	this.ListWidget.SetObjectName("ListWidget")           // 112
	this.ListWidget.SetAlternatingRowColors(false)        // 114

	this.VerticalLayout_5.Layout().AddWidget(this.ListWidget) // 115

	this.HorizontalLayout_3 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_3.SetSpacing(0)                       // 114
	this.HorizontalLayout_3.SetObjectName("HorizontalLayout_3") // 112
	this.ToolButton_4 = qtwidgets.NewQToolButton(this.Page)     // 111
	this.ToolButton_4.SetObjectName("ToolButton_4")             // 112
	this.Icon5 = qtgui.NewQIcon()
	this.Icon5.AddFile(":/icons/add-square-button-gray.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_4.SetIcon(this.Icon5)                                                                              // 114
	this.ToolButton_4.SetIconSize(qtcore.NewQSize_1(22, 22))                                                           // 113
	this.ToolButton_4.SetAutoRaise(true)                                                                               // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_4) // 115

	this.ToolButton_5 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_5.SetObjectName("ToolButton_5")         // 112
	this.Icon6 = qtgui.NewQIcon()
	this.Icon6.AddFile(":/icons/groupgray.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_5.SetIcon(this.Icon6)                                                                 // 114
	this.ToolButton_5.SetIconSize(qtcore.NewQSize_1(22, 22))                                              // 113
	this.ToolButton_5.SetAutoRaise(true)                                                                  // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_5) // 115

	this.ToolButton_6 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_6.SetObjectName("ToolButton_6")         // 112
	this.Icon7 = qtgui.NewQIcon()
	this.Icon7.AddFile(":/icons/transfer_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_6.SetIcon(this.Icon7)                                                                       // 114
	this.ToolButton_6.SetIconSize(qtcore.NewQSize_1(22, 22))                                                    // 113
	this.ToolButton_6.SetAutoRaise(true)                                                                        // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_6) // 115

	this.ToolButton_7 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_7.SetObjectName("ToolButton_7")         // 112
	this.Icon8 = qtgui.NewQIcon()
	this.Icon8.AddFile(":/icons/settings_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_7.SetIcon(this.Icon8)                                                                       // 114
	this.ToolButton_7.SetIconSize(qtcore.NewQSize_1(22, 22))                                                    // 113
	this.ToolButton_7.SetAutoRaise(true)                                                                        // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_7) // 115

	this.VerticalLayout_5.AddLayout(this.HorizontalLayout_3, 0) // 115

	this.StackedWidget.AddWidget(this.Page)                                                 // 115
	this.Page_2 = qtwidgets.NewQWidget(nil, 0)                                              // 111
	this.Page_2.SetObjectName("Page_2")                                                     // 112
	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout_1(this.Page_2)                         // 111
	this.VerticalLayout_2.SetSpacing(0)                                                     // 114
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")                                 // 112
	this.VerticalLayout_2.SetContentsMargins(0, 0, 0, 0)                                    // 114
	this.HorizontalLayout_5 = qtwidgets.NewQHBoxLayout()                                    // 111
	this.HorizontalLayout_5.SetSpacing(0)                                                   // 114
	this.HorizontalLayout_5.SetObjectName("HorizontalLayout_5")                             // 112
	this.Label_4 = qtwidgets.NewQLabel(this.Page_2, 0)                                      // 111
	this.Label_4.SetObjectName("Label_4")                                                   // 112
	this.Label_4.SetMaximumSize_1(32, 32)                                                   // 113
	this.Label_4.SetPixmap(qtgui.NewQPixmap_3(":/Icons/Icon_avatar_40.Png", "dummy123", 0)) // 114

	this.HorizontalLayout_5.Layout().AddWidget(this.Label_4) // 115

	this.VerticalLayout_4 = qtwidgets.NewQVBoxLayout()                               // 111
	this.VerticalLayout_4.SetSpacing(0)                                              // 114
	this.VerticalLayout_4.SetObjectName("VerticalLayout_4")                          // 112
	this.Label_5 = qtwidgets.NewQLabel(this.Page_2, 0)                               // 111
	this.Label_5.SetObjectName("Label_5")                                            // 112
	this.SizePolicy.SetHeightForWidth(this.Label_5.SizePolicy().HasHeightForWidth()) // 114
	this.Label_5.SetSizePolicy(this.SizePolicy)                                      // 114
	this.Font1 = qtgui.NewQFont()
	this.Font1.SetBold(true)                                                                                   // 114
	this.Font1.SetWeight(75)                                                                                   // 114
	this.Label_5.SetFont(this.Font1)                                                                           // 114
	this.Label_5.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout_4.Layout().AddWidget(this.Label_5) // 115

	this.HorizontalLayout_4 = qtwidgets.NewQHBoxLayout()                                                       // 111
	this.HorizontalLayout_4.SetSpacing(0)                                                                      // 114
	this.HorizontalLayout_4.SetObjectName("HorizontalLayout_4")                                                // 112
	this.Label_6 = qtwidgets.NewQLabel(this.Page_2, 0)                                                         // 111
	this.Label_6.SetObjectName("Label_6")                                                                      // 112
	this.Label_6.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_4.Layout().AddWidget(this.Label_6) // 115

	this.Label_7 = qtwidgets.NewQLabel(this.Page_2, 0)                                                         // 111
	this.Label_7.SetObjectName("Label_7")                                                                      // 112
	this.SizePolicy.SetHeightForWidth(this.Label_7.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_7.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_7.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_4.Layout().AddWidget(this.Label_7) // 115

	this.VerticalLayout_4.AddLayout(this.HorizontalLayout_4, 0) // 115

	this.HorizontalLayout_5.AddLayout(this.VerticalLayout_4, 0) // 115

	this.VerticalLayout_7 = qtwidgets.NewQVBoxLayout()         // 111
	this.VerticalLayout_7.SetSpacing(0)                        // 114
	this.VerticalLayout_7.SetObjectName("VerticalLayout_7")    // 112
	this.ToolButton_15 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_15.SetObjectName("ToolButton_15")          // 112
	this.ToolButton_15.SetMaximumSize_1(16, 16)                // 113
	this.ToolButton_15.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon9 = qtgui.NewQIcon()
	this.Icon9.AddFile(":/icons/phone_mic_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_15.SetIcon(this.Icon9)                                                                       // 114
	this.ToolButton_15.SetIconSize(qtcore.NewQSize_1(12, 16))                                                    // 113
	this.ToolButton_15.SetAutoRaise(true)                                                                        // 114

	this.VerticalLayout_7.Layout().AddWidget(this.ToolButton_15) // 115

	this.ToolButton_16 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_16.SetObjectName("ToolButton_16")          // 112
	this.ToolButton_16.SetMaximumSize_1(16, 16)                // 113
	this.ToolButton_16.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon10 = qtgui.NewQIcon()
	this.Icon10.AddFile(":/icons/speaker_volume_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_16.SetIcon(this.Icon10)                                                                            // 114
	this.ToolButton_16.SetIconSize(qtcore.NewQSize_1(12, 12))                                                          // 113
	this.ToolButton_16.SetAutoRaise(true)                                                                              // 114

	this.VerticalLayout_7.Layout().AddWidget(this.ToolButton_16) // 115

	this.HorizontalLayout_5.AddLayout(this.VerticalLayout_7, 0) // 115

	this.ToolButton_13 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_13.SetObjectName("ToolButton_13")          // 112
	this.ToolButton_13.SetMinimumSize_1(0, 0)                  // 113
	this.ToolButton_13.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon11 = qtgui.NewQIcon()
	this.Icon11.AddFile(":/icons/phone_call_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_13.SetIcon(this.Icon11)                                                                        // 114
	this.ToolButton_13.SetIconSize(qtcore.NewQSize_1(32, 32))                                                      // 113
	this.ToolButton_13.SetAutoRaise(true)                                                                          // 114

	this.HorizontalLayout_5.Layout().AddWidget(this.ToolButton_13) // 115

	this.ToolButton_14 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_14.SetObjectName("ToolButton_14")          // 112
	this.ToolButton_14.SetMinimumSize_1(0, 0)                  // 113
	this.ToolButton_14.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon12 = qtgui.NewQIcon()
	this.Icon12.AddFile(":/icons/video_recorder_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_14.SetIcon(this.Icon12)                                                                            // 114
	this.ToolButton_14.SetIconSize(qtcore.NewQSize_1(32, 32))                                                          // 113
	this.ToolButton_14.SetAutoRaise(true)                                                                              // 114

	this.HorizontalLayout_5.Layout().AddWidget(this.ToolButton_14) // 115

	this.VerticalLayout_2.AddLayout(this.HorizontalLayout_5, 0) // 115

	this.ScrollArea_2 = qtwidgets.NewQScrollArea(this.Page_2)                           // 111
	this.ScrollArea_2.SetObjectName("ScrollArea_2")                                     // 112
	this.ScrollArea_2.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)       // 114
	this.ScrollArea_2.SetWidgetResizable(true)                                          // 114
	this.ScrollAreaWidgetContents_2 = qtwidgets.NewQWidget(nil, 0)                      // 111
	this.ScrollAreaWidgetContents_2.SetObjectName("ScrollAreaWidgetContents_2")         // 112
	this.ScrollAreaWidgetContents_2.SetGeometry(0, 0, 392, 237)                         // 114
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents_2) // 111
	this.VerticalLayout_3.SetSpacing(0)                                                 // 114
	this.VerticalLayout_3.SetObjectName("VerticalLayout_3")                             // 112
	this.VerticalLayout_3.SetContentsMargins(0, 0, 0, 0)                                // 114
	this.VerticalLayout_8 = qtwidgets.NewQVBoxLayout()                                  // 111
	this.VerticalLayout_8.SetSpacing(0)                                                 // 114
	this.VerticalLayout_8.SetObjectName("VerticalLayout_8")                             // 112
	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 0, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_8.AddItem(this.VerticalSpacer) // 115

	this.VerticalLayout_3.AddLayout(this.VerticalLayout_8, 0) // 115

	this.ScrollArea_2.SetWidget(this.ScrollAreaWidgetContents_2) // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ScrollArea_2) // 115

	this.ScrollArea = qtwidgets.NewQScrollArea(this.Page_2)                           // 111
	this.ScrollArea.SetObjectName("ScrollArea")                                       // 112
	this.ScrollArea.SetWidgetResizable(true)                                          // 114
	this.ScrollAreaWidgetContents = qtwidgets.NewQWidget(nil, 0)                      // 111
	this.ScrollAreaWidgetContents.SetObjectName("ScrollAreaWidgetContents")           // 112
	this.ScrollAreaWidgetContents.SetGeometry(0, 0, 392, 237)                         // 114
	this.VerticalLayout_9 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents) // 111
	this.VerticalLayout_9.SetSpacing(0)                                               // 114
	this.VerticalLayout_9.SetObjectName("VerticalLayout_9")                           // 112
	this.VerticalLayout_9.SetContentsMargins(0, 0, 0, 0)                              // 114
	this.VerticalLayout_10 = qtwidgets.NewQVBoxLayout()                               // 111
	this.VerticalLayout_10.SetSpacing(0)                                              // 114
	this.VerticalLayout_10.SetObjectName("VerticalLayout_10")                         // 112
	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_10.AddItem(this.VerticalSpacer_2) // 115

	this.VerticalLayout_9.AddLayout(this.VerticalLayout_10, 0) // 115

	this.ScrollArea.SetWidget(this.ScrollAreaWidgetContents) // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ScrollArea) // 115

	this.HorizontalLayout_6 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_6.SetSpacing(0)                       // 114
	this.HorizontalLayout_6.SetObjectName("HorizontalLayout_6") // 112
	this.ToolButton_8 = qtwidgets.NewQToolButton(this.Page_2)   // 111
	this.ToolButton_8.SetObjectName("ToolButton_8")             // 112
	this.Icon13 = qtgui.NewQIcon()
	this.Icon13.AddFile(":/icons/paper-clip-outline_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_8.SetIcon(this.Icon13)                                                                                 // 114
	this.ToolButton_8.SetIconSize(qtcore.NewQSize_1(22, 22))                                                               // 113
	this.ToolButton_8.SetAutoRaise(true)                                                                                   // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_8) // 115

	this.ToolButton_9 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_9.SetObjectName("ToolButton_9")           // 112
	this.Icon14 = qtgui.NewQIcon()
	this.Icon14.AddFile(":/icons/question-mark-gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_9.SetIcon(this.Icon14)                                                                            // 114
	this.ToolButton_9.SetIconSize(qtcore.NewQSize_1(22, 22))                                                          // 113
	this.ToolButton_9.SetAutoRaise(true)                                                                              // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_9) // 115

	this.ToolButton_10 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_10.SetObjectName("ToolButton_10")          // 112
	this.Icon15 = qtgui.NewQIcon()
	this.Icon15.AddFile(":/icons/smile_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_10.SetIcon(this.Icon15)                                                                   // 114
	this.ToolButton_10.SetIconSize(qtcore.NewQSize_1(22, 22))                                                 // 113
	this.ToolButton_10.SetAutoRaise(true)                                                                     // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_10) // 115

	this.LineEdit_2 = qtwidgets.NewQLineEdit(this.Page_2) // 111
	this.LineEdit_2.SetObjectName("LineEdit_2")           // 112

	this.HorizontalLayout_6.Layout().AddWidget(this.LineEdit_2) // 115

	this.ToolButton_18 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_18.SetObjectName("ToolButton_18")          // 112
	this.Icon16 = qtgui.NewQIcon()
	this.Icon16.AddFile(":/icons/cursor_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_18.SetIcon(this.Icon16)                                                                    // 114
	this.ToolButton_18.SetIconSize(qtcore.NewQSize_1(22, 22))                                                  // 113
	this.ToolButton_18.SetToolButtonStyle(qtcore.Qt__ToolButtonTextBesideIcon)                                 // 114
	this.ToolButton_18.SetAutoRaise(true)                                                                      // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_18) // 115

	this.VerticalLayout_2.AddLayout(this.HorizontalLayout_6, 0) // 115

	this.StackedWidget.AddWidget(this.Page_2) // 115

	this.VerticalLayout_6.Layout().AddWidget(this.StackedWidget) // 115

	this.MainWindow.SetCentralWidget(this.Centralwidget) // 114

	this.RetranslateUi(MainWindow)

	this.StackedWidget.SetCurrentIndex(2) // 114

	qtcore.QMetaObject_ConnectSlotsByName(MainWindow) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QMainWindow *MainWindow)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MainWindow) RetranslateUi(MainWindow *qtwidgets.QMainWindow) {
	this.MainWindow.SetWindowTitle(qtcore.QCoreApplication_Translate("MainWindow", "MainWindow", "dummy123", 0))
	this.Actionooo.SetText(qtcore.QCoreApplication_Translate("MainWindow", "ooo", "dummy123", 0))
	this.ActionQuit.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Quit", "dummy123", 0))
	this.Action_About.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&About", "dummy123", 0))
	this.ToolButton_11.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.ToolButton_12.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_17.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.ToolButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.LineEdit.SetPlaceholderText(qtcore.QCoreApplication_Translate("MainWindow", "Filter...", "dummy123", 0))
	this.ToolButton_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.Label_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "0 people", "dummy123", 0))
	this.Label_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.ToolButton_15.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_16.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_13.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_14.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_8.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_9.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_10.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.LineEdit_2.SetPlaceholderText(qtcore.QCoreApplication_Translate("MainWindow", "Write a message...", "dummy123", 0))
	this.ToolButton_18.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Send", "dummy123", 0))
}

//  retranslateUi block end
