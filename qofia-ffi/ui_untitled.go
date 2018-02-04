package main

//  header block begin
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtwidgets"
import "qt.go/qtmock"

func init() { qtcore.KeepMe() }
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
	Splitter                   *qtwidgets.QSplitter
	LayoutWidget               *qtwidgets.QWidget
	VerticalLayout_2           *qtwidgets.QVBoxLayout
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
	ScrollArea                 *qtwidgets.QScrollArea
	ScrollAreaWidgetContents   *qtwidgets.QWidget
	VerticalLayout_9           *qtwidgets.QVBoxLayout
	VerticalLayout_10          *qtwidgets.QVBoxLayout
	VerticalSpacer_2           *qtwidgets.QSpacerItem
	HorizontalLayout_3         *qtwidgets.QHBoxLayout
	ToolButton_4               *qtwidgets.QToolButton
	ToolButton_5               *qtwidgets.QToolButton
	ToolButton_6               *qtwidgets.QToolButton
	ToolButton_7               *qtwidgets.QToolButton
	LayoutWidget1              *qtwidgets.QWidget
	VerticalLayout_5           *qtwidgets.QVBoxLayout
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
	ListWidget                 *qtwidgets.QListWidget
	ScrollArea_2               *qtwidgets.QScrollArea
	ScrollAreaWidgetContents_2 *qtwidgets.QWidget
	VerticalLayout_3           *qtwidgets.QVBoxLayout
	VerticalLayout_8           *qtwidgets.QVBoxLayout
	VerticalSpacer             *qtwidgets.QSpacerItem
	HorizontalLayout_6         *qtwidgets.QHBoxLayout
	ToolButton_8               *qtwidgets.QToolButton
	ToolButton_9               *qtwidgets.QToolButton
	ToolButton_10              *qtwidgets.QToolButton
	LineEdit_2                 *qtwidgets.QLineEdit
	ToolButton_18              *qtwidgets.QToolButton
	MainWindow                 *qtwidgets.QMainWindow
	Icon                       *qtgui.QIcon // 116
	Font                       *qtgui.QFont // 116
	SizePolicy                 *qtwidgets.QSizePolicy
	Icon1                      *qtgui.QIcon // 116
	Icon2                      *qtgui.QIcon // 116
	Icon3                      *qtgui.QIcon // 116
	Icon4                      *qtgui.QIcon // 116
	Icon5                      *qtgui.QIcon // 116
	Icon6                      *qtgui.QIcon // 116
	Font1                      *qtgui.QFont // 116
	Icon7                      *qtgui.QIcon // 116
	Icon8                      *qtgui.QIcon // 116
	Icon9                      *qtgui.QIcon // 116
	Icon10                     *qtgui.QIcon // 116
	Icon11                     *qtgui.QIcon // 116
	Icon12                     *qtgui.QIcon // 116
	Icon13                     *qtgui.QIcon // 116
	Icon14                     *qtgui.QIcon // 116
}

//  struct block end

//  setupUi block begin
// void setupUi(QMainWindow *MainWindow)
func (this *Ui_MainWindow) SetupUi(MainWindow *qtwidgets.QMainWindow) {
	this.MainWindow = MainWindow
	// { // 126
	if MainWindow.ObjectName().IsEmpty() {
		MainWindow.SetObjectName(qtcore.NewQString_5("MainWindow"))
	}
	MainWindow.Resize(689, 531)
	this.Actionooo = qtwidgets.NewQAction(qtcore.NewQObjectFromPointer(MainWindow.GetCthis()))    // 111
	this.Actionooo.SetObjectName(qtcore.NewQString_5("Actionooo"))                                // 112
	this.ActionQuit = qtwidgets.NewQAction(qtcore.NewQObjectFromPointer(MainWindow.GetCthis()))   // 111
	this.ActionQuit.SetObjectName(qtcore.NewQString_5("ActionQuit"))                              // 112
	this.Action_About = qtwidgets.NewQAction(qtcore.NewQObjectFromPointer(MainWindow.GetCthis())) // 111
	this.Action_About.SetObjectName(qtcore.NewQString_5("Action_About"))                          // 112

	this.Centralwidget = qtwidgets.NewQWidget(qtwidgets.NewQWidgetFromPointer(this.MainWindow.GetCthis()), 0) // 111
	this.Centralwidget.SetObjectName(qtcore.NewQString_5("Centralwidget"))                                    // 112

	this.VerticalLayout_6 = qtwidgets.NewQVBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.Centralwidget.GetCthis())) // 111
	this.VerticalLayout_6.SetSpacing(1)                                                                                // 114
	this.VerticalLayout_6.SetObjectName(qtcore.NewQString_5("VerticalLayout_6"))                                       // 112
	this.VerticalLayout_6.SetContentsMargins(1, 1, 1, 1)                                                               // 114

	this.Splitter = qtwidgets.NewQSplitter(this.Centralwidget)   // 111
	this.Splitter.SetObjectName(qtcore.NewQString_5("Splitter")) // 112
	this.Splitter.SetOrientation(qtcore.Qt__Horizontal)          // 114

	this.LayoutWidget = qtwidgets.NewQWidget(qtwidgets.NewQWidgetFromPointer(this.Splitter.GetCthis()), 0) // 111
	this.LayoutWidget.SetObjectName(qtcore.NewQString_5("LayoutWidget"))                                   // 112

	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.LayoutWidget.GetCthis())) // 111
	this.VerticalLayout_2.SetSpacing(3)                                                                               // 114
	this.VerticalLayout_2.SetObjectName(qtcore.NewQString_5("VerticalLayout_2"))                                      // 112
	this.VerticalLayout_2.SetContentsMargins(0, 0, 0, 0)                                                              // 114
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()                                                                // 111
	this.HorizontalLayout.SetSpacing(3)                                                                               // 114
	this.HorizontalLayout.SetObjectName(qtcore.NewQString_5("HorizontalLayout"))                                      // 112
	this.ToolButton_17 = qtwidgets.NewQToolButton(this.LayoutWidget)                                                  // 111
	this.ToolButton_17.SetObjectName(qtcore.NewQString_5("ToolButton_17"))                                            // 112
	this.ToolButton_17.SetMaximumSize_1(32, 32)                                                                       // 113
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(qtcore.NewQString_5(":/icons/icon_avatar_40.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_17.SetIcon(this.Icon)                                                                                          // 114
	this.ToolButton_17.SetIconSize(qtcore.NewQSize_1(32, 32))                                                                      // 113

	this.HorizontalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_17.GetCthis())) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()                         // 111
	this.VerticalLayout.SetObjectName(qtcore.NewQString_5("VerticalLayout")) // 112
	this.Label_2 = qtwidgets.NewQLabel(this.LayoutWidget, 0)                 // 111
	this.Label_2.SetObjectName(qtcore.NewQString_5("Label_2"))               // 112
	this.Font = qtgui.NewQFont()
	this.Font.SetPointSize(12)                                                                       // 114
	this.Font.SetBold(true)                                                                          // 114
	this.Font.SetWeight(75)                                                                          // 114
	this.Label_2.SetFont(this.Font)                                                                  // 114
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__TextEditable | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_2.GetCthis())) // 115

	this.Label_3 = qtwidgets.NewQLabel(this.LayoutWidget, 0)   // 111
	this.Label_3.SetObjectName(qtcore.NewQString_5("Label_3")) // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                          // 114
	this.SizePolicy.SetVerticalStretch(0)                                                            // 114
	this.SizePolicy.SetHeightForWidth(this.Label_3.SizePolicy().HasHeightForWidth())                 // 114
	this.Label_3.SetSizePolicy(this.SizePolicy)                                                      // 114
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__TextEditable | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_3.GetCthis())) // 115

	this.HorizontalLayout.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout.GetCthis()), 0) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.LayoutWidget)    // 111
	this.ToolButton.SetObjectName(qtcore.NewQString_5("ToolButton")) // 112
	this.Icon1 = qtgui.NewQIcon()
	this.Icon1.AddFile(qtcore.NewQString_5(":/icons/online_30.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton.SetIcon(this.Icon1)                                                                                        // 114
	this.ToolButton.SetToolButtonStyle(qtcore.Qt__ToolButtonIconOnly)                                                          // 114
	this.ToolButton.SetAutoRaise(true)                                                                                         // 114

	this.HorizontalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton.GetCthis())) // 115

	this.VerticalLayout_2.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout.GetCthis()), 0) // 115

	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout()                             // 111
	this.HorizontalLayout_2.SetObjectName(qtcore.NewQString_5("HorizontalLayout_2")) // 112
	this.LineEdit = qtwidgets.NewQLineEdit(this.LayoutWidget)                        // 111
	this.LineEdit.SetObjectName(qtcore.NewQString_5("LineEdit"))                     // 112

	this.HorizontalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.LineEdit.GetCthis())) // 115

	this.ToolButton_2 = qtwidgets.NewQToolButton(this.LayoutWidget)      // 111
	this.ToolButton_2.SetObjectName(qtcore.NewQString_5("ToolButton_2")) // 112
	this.ToolButton_2.SetMinimumSize_1(60, 0)                            // 113
	this.ToolButton_2.SetAutoRaise(true)                                 // 114

	this.HorizontalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_2.GetCthis())) // 115

	this.ToolButton_3 = qtwidgets.NewQToolButton(this.LayoutWidget)      // 111
	this.ToolButton_3.SetObjectName(qtcore.NewQString_5("ToolButton_3")) // 112
	this.Icon2 = qtgui.NewQIcon()
	this.Icon2.AddFile(qtcore.NewQString_5(":/icons/remove-symbol_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_3.SetIcon(this.Icon2)                                                                                                 // 114
	this.ToolButton_3.SetAutoRaise(true)                                                                                                  // 114

	this.HorizontalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_3.GetCthis())) // 115

	this.VerticalLayout_2.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout_2.GetCthis()), 0) // 115

	this.ListWidget_2 = qtwidgets.NewQListWidget(this.LayoutWidget)      // 111
	this.ListWidget_2.SetObjectName(qtcore.NewQString_5("ListWidget_2")) // 112

	this.VerticalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ListWidget_2.GetCthis())) // 115

	this.ScrollArea = qtwidgets.NewQScrollArea(this.LayoutWidget)                                                                 // 111
	this.ScrollArea.SetObjectName(qtcore.NewQString_5("ScrollArea"))                                                              // 112
	this.ScrollArea.SetWidgetResizable(true)                                                                                      // 114
	this.ScrollAreaWidgetContents = qtwidgets.NewQWidget(nil, 0)                                                                  // 111
	this.ScrollAreaWidgetContents.SetObjectName(qtcore.NewQString_5("ScrollAreaWidgetContents"))                                  // 112
	this.ScrollAreaWidgetContents.SetGeometry(0, 0, 310, 201)                                                                     // 114
	this.VerticalLayout_9 = qtwidgets.NewQVBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.ScrollAreaWidgetContents.GetCthis())) // 111
	this.VerticalLayout_9.SetObjectName(qtcore.NewQString_5("VerticalLayout_9"))                                                  // 112
	this.VerticalLayout_10 = qtwidgets.NewQVBoxLayout()                                                                           // 111
	this.VerticalLayout_10.SetObjectName(qtcore.NewQString_5("VerticalLayout_10"))                                                // 112
	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_10.AddItem(qtwidgets.NewQLayoutItemFromPointer(this.VerticalSpacer_2.GetCthis())) // 115

	this.VerticalLayout_9.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_10.GetCthis()), 0) // 115

	this.ScrollArea.SetWidget(this.ScrollAreaWidgetContents) // 114

	this.VerticalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ScrollArea.GetCthis())) // 115

	this.HorizontalLayout_3 = qtwidgets.NewQHBoxLayout()                             // 111
	this.HorizontalLayout_3.SetSpacing(3)                                            // 114
	this.HorizontalLayout_3.SetObjectName(qtcore.NewQString_5("HorizontalLayout_3")) // 112
	this.ToolButton_4 = qtwidgets.NewQToolButton(this.LayoutWidget)                  // 111
	this.ToolButton_4.SetObjectName(qtcore.NewQString_5("ToolButton_4"))             // 112
	this.Icon3 = qtgui.NewQIcon()
	this.Icon3.AddFile(qtcore.NewQString_5(":/icons/add-square-button-gray.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_4.SetIcon(this.Icon3)                                                                                                   // 114
	this.ToolButton_4.SetAutoRaise(true)                                                                                                    // 114

	this.HorizontalLayout_3.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_4.GetCthis())) // 115

	this.ToolButton_5 = qtwidgets.NewQToolButton(this.LayoutWidget)      // 111
	this.ToolButton_5.SetObjectName(qtcore.NewQString_5("ToolButton_5")) // 112
	this.Icon4 = qtgui.NewQIcon()
	this.Icon4.AddFile(qtcore.NewQString_5(":/icons/groupgray.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_5.SetIcon(this.Icon4)                                                                                      // 114
	this.ToolButton_5.SetAutoRaise(true)                                                                                       // 114

	this.HorizontalLayout_3.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_5.GetCthis())) // 115

	this.ToolButton_6 = qtwidgets.NewQToolButton(this.LayoutWidget)      // 111
	this.ToolButton_6.SetObjectName(qtcore.NewQString_5("ToolButton_6")) // 112
	this.Icon5 = qtgui.NewQIcon()
	this.Icon5.AddFile(qtcore.NewQString_5(":/icons/transfer_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_6.SetIcon(this.Icon5)                                                                                            // 114
	this.ToolButton_6.SetAutoRaise(true)                                                                                             // 114

	this.HorizontalLayout_3.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_6.GetCthis())) // 115

	this.ToolButton_7 = qtwidgets.NewQToolButton(this.LayoutWidget)      // 111
	this.ToolButton_7.SetObjectName(qtcore.NewQString_5("ToolButton_7")) // 112
	this.Icon6 = qtgui.NewQIcon()
	this.Icon6.AddFile(qtcore.NewQString_5(":/icons/settings_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_7.SetIcon(this.Icon6)                                                                                            // 114
	this.ToolButton_7.SetAutoRaise(true)                                                                                             // 114

	this.HorizontalLayout_3.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_7.GetCthis())) // 115

	this.VerticalLayout_2.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout_3.GetCthis()), 0) // 115

	this.Splitter.AddWidget(qtwidgets.NewQWidgetFromPointer(this.LayoutWidget.GetCthis()))                             // 115
	this.LayoutWidget1 = qtwidgets.NewQWidget(qtwidgets.NewQWidgetFromPointer(this.Splitter.GetCthis()), 0)            // 111
	this.LayoutWidget1.SetObjectName(qtcore.NewQString_5("LayoutWidget1"))                                             // 112
	this.VerticalLayout_5 = qtwidgets.NewQVBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.LayoutWidget1.GetCthis())) // 111
	this.VerticalLayout_5.SetSpacing(3)                                                                                // 114
	this.VerticalLayout_5.SetObjectName(qtcore.NewQString_5("VerticalLayout_5"))                                       // 112
	this.VerticalLayout_5.SetContentsMargins(0, 0, 0, 0)                                                               // 114
	this.HorizontalLayout_5 = qtwidgets.NewQHBoxLayout()                                                               // 111
	this.HorizontalLayout_5.SetSpacing(3)                                                                              // 114
	this.HorizontalLayout_5.SetObjectName(qtcore.NewQString_5("HorizontalLayout_5"))                                   // 112
	this.Label_4 = qtwidgets.NewQLabel(this.LayoutWidget1, 0)                                                          // 111
	this.Label_4.SetObjectName(qtcore.NewQString_5("Label_4"))                                                         // 112
	this.Label_4.SetMaximumSize_1(32, 32)                                                                              // 113
	this.Label_4.SetPixmap(qtgui.NewQPixmap_3(qtcore.NewQString_5(":/Icons/Icon_avatar_40.Png"), "dummy123", 0))       // 114

	this.HorizontalLayout_5.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_4.GetCthis())) // 115

	this.VerticalLayout_4 = qtwidgets.NewQVBoxLayout()                               // 111
	this.VerticalLayout_4.SetObjectName(qtcore.NewQString_5("VerticalLayout_4"))     // 112
	this.Label_5 = qtwidgets.NewQLabel(this.LayoutWidget1, 0)                        // 111
	this.Label_5.SetObjectName(qtcore.NewQString_5("Label_5"))                       // 112
	this.SizePolicy.SetHeightForWidth(this.Label_5.SizePolicy().HasHeightForWidth()) // 114
	this.Label_5.SetSizePolicy(this.SizePolicy)                                      // 114
	this.Font1 = qtgui.NewQFont()
	this.Font1.SetBold(true)                                                                                   // 114
	this.Font1.SetWeight(75)                                                                                   // 114
	this.Label_5.SetFont(this.Font1)                                                                           // 114
	this.Label_5.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout_4.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_5.GetCthis())) // 115

	this.HorizontalLayout_4 = qtwidgets.NewQHBoxLayout()                                                       // 111
	this.HorizontalLayout_4.SetObjectName(qtcore.NewQString_5("HorizontalLayout_4"))                           // 112
	this.Label_6 = qtwidgets.NewQLabel(this.LayoutWidget1, 0)                                                  // 111
	this.Label_6.SetObjectName(qtcore.NewQString_5("Label_6"))                                                 // 112
	this.Label_6.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_4.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_6.GetCthis())) // 115

	this.Label_7 = qtwidgets.NewQLabel(this.LayoutWidget1, 0)                                                  // 111
	this.Label_7.SetObjectName(qtcore.NewQString_5("Label_7"))                                                 // 112
	this.SizePolicy.SetHeightForWidth(this.Label_7.SizePolicy().HasHeightForWidth())                           // 114
	this.Label_7.SetSizePolicy(this.SizePolicy)                                                                // 114
	this.Label_7.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_4.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_7.GetCthis())) // 115

	this.VerticalLayout_4.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout_4.GetCthis()), 0) // 115

	this.HorizontalLayout_5.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_4.GetCthis()), 0) // 115

	this.VerticalLayout_7 = qtwidgets.NewQVBoxLayout()                           // 111
	this.VerticalLayout_7.SetObjectName(qtcore.NewQString_5("VerticalLayout_7")) // 112
	this.ToolButton_15 = qtwidgets.NewQToolButton(this.LayoutWidget1)            // 111
	this.ToolButton_15.SetObjectName(qtcore.NewQString_5("ToolButton_15"))       // 112
	this.ToolButton_15.SetMaximumSize_1(16, 16)                                  // 113
	this.Icon7 = qtgui.NewQIcon()
	this.Icon7.AddFile(qtcore.NewQString_5(":/icons/phone_mic_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_15.SetIcon(this.Icon7)                                                                                            // 114
	this.ToolButton_15.SetIconSize(qtcore.NewQSize_1(12, 16))                                                                         // 113
	this.ToolButton_15.SetAutoRaise(true)                                                                                             // 114

	this.VerticalLayout_7.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_15.GetCthis())) // 115

	this.ToolButton_16 = qtwidgets.NewQToolButton(this.LayoutWidget1)      // 111
	this.ToolButton_16.SetObjectName(qtcore.NewQString_5("ToolButton_16")) // 112
	this.ToolButton_16.SetMaximumSize_1(16, 16)                            // 113
	this.Icon8 = qtgui.NewQIcon()
	this.Icon8.AddFile(qtcore.NewQString_5(":/icons/speaker_volume_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_16.SetIcon(this.Icon8)                                                                                                 // 114
	this.ToolButton_16.SetIconSize(qtcore.NewQSize_1(12, 12))                                                                              // 113
	this.ToolButton_16.SetAutoRaise(true)                                                                                                  // 114

	this.VerticalLayout_7.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_16.GetCthis())) // 115

	this.HorizontalLayout_5.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_7.GetCthis()), 0) // 115

	this.ToolButton_13 = qtwidgets.NewQToolButton(this.LayoutWidget1)      // 111
	this.ToolButton_13.SetObjectName(qtcore.NewQString_5("ToolButton_13")) // 112
	this.ToolButton_13.SetMinimumSize_1(0, 0)                              // 113
	this.Icon9 = qtgui.NewQIcon()
	this.Icon9.AddFile(qtcore.NewQString_5(":/icons/phone_call_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_13.SetIcon(this.Icon9)                                                                                             // 114
	this.ToolButton_13.SetIconSize(qtcore.NewQSize_1(32, 32))                                                                          // 113
	this.ToolButton_13.SetAutoRaise(true)                                                                                              // 114

	this.HorizontalLayout_5.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_13.GetCthis())) // 115

	this.ToolButton_14 = qtwidgets.NewQToolButton(this.LayoutWidget1)      // 111
	this.ToolButton_14.SetObjectName(qtcore.NewQString_5("ToolButton_14")) // 112
	this.ToolButton_14.SetMinimumSize_1(0, 0)                              // 113
	this.Icon10 = qtgui.NewQIcon()
	this.Icon10.AddFile(qtcore.NewQString_5(":/icons/video_recorder_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_14.SetIcon(this.Icon10)                                                                                                 // 114
	this.ToolButton_14.SetIconSize(qtcore.NewQSize_1(32, 32))                                                                               // 113
	this.ToolButton_14.SetAutoRaise(true)                                                                                                   // 114

	this.HorizontalLayout_5.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_14.GetCthis())) // 115

	this.VerticalLayout_5.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout_5.GetCthis()), 0) // 115

	this.ListWidget = qtwidgets.NewQListWidget(this.LayoutWidget1)   // 111
	this.ListWidget.SetObjectName(qtcore.NewQString_5("ListWidget")) // 112
	this.ListWidget.SetAlternatingRowColors(false)                   // 114

	this.VerticalLayout_5.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ListWidget.GetCthis())) // 115

	this.ScrollArea_2 = qtwidgets.NewQScrollArea(this.LayoutWidget1)                                                                // 111
	this.ScrollArea_2.SetObjectName(qtcore.NewQString_5("ScrollArea_2"))                                                            // 112
	this.ScrollArea_2.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)                                                   // 114
	this.ScrollArea_2.SetWidgetResizable(true)                                                                                      // 114
	this.ScrollAreaWidgetContents_2 = qtwidgets.NewQWidget(nil, 0)                                                                  // 111
	this.ScrollAreaWidgetContents_2.SetObjectName(qtcore.NewQString_5("ScrollAreaWidgetContents_2"))                                // 112
	this.ScrollAreaWidgetContents_2.SetGeometry(0, 0, 364, 216)                                                                     // 114
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.ScrollAreaWidgetContents_2.GetCthis())) // 111
	this.VerticalLayout_3.SetObjectName(qtcore.NewQString_5("VerticalLayout_3"))                                                    // 112
	this.VerticalLayout_8 = qtwidgets.NewQVBoxLayout()                                                                              // 111
	this.VerticalLayout_8.SetObjectName(qtcore.NewQString_5("VerticalLayout_8"))                                                    // 112
	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_8.AddItem(qtwidgets.NewQLayoutItemFromPointer(this.VerticalSpacer.GetCthis())) // 115

	this.VerticalLayout_3.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_8.GetCthis()), 0) // 115

	this.ScrollArea_2.SetWidget(this.ScrollAreaWidgetContents_2) // 114

	this.VerticalLayout_5.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ScrollArea_2.GetCthis())) // 115

	this.HorizontalLayout_6 = qtwidgets.NewQHBoxLayout()                             // 111
	this.HorizontalLayout_6.SetSpacing(3)                                            // 114
	this.HorizontalLayout_6.SetObjectName(qtcore.NewQString_5("HorizontalLayout_6")) // 112
	this.ToolButton_8 = qtwidgets.NewQToolButton(this.LayoutWidget1)                 // 111
	this.ToolButton_8.SetObjectName(qtcore.NewQString_5("ToolButton_8"))             // 112
	this.Icon11 = qtgui.NewQIcon()
	this.Icon11.AddFile(qtcore.NewQString_5(":/icons/paper-clip-outline_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_8.SetIcon(this.Icon11)                                                                                                      // 114
	this.ToolButton_8.SetAutoRaise(true)                                                                                                        // 114

	this.HorizontalLayout_6.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_8.GetCthis())) // 115

	this.ToolButton_9 = qtwidgets.NewQToolButton(this.LayoutWidget1)     // 111
	this.ToolButton_9.SetObjectName(qtcore.NewQString_5("ToolButton_9")) // 112
	this.Icon12 = qtgui.NewQIcon()
	this.Icon12.AddFile(qtcore.NewQString_5(":/icons/question-mark-gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_9.SetIcon(this.Icon12)                                                                                                 // 114
	this.ToolButton_9.SetAutoRaise(true)                                                                                                   // 114

	this.HorizontalLayout_6.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_9.GetCthis())) // 115

	this.ToolButton_10 = qtwidgets.NewQToolButton(this.LayoutWidget1)      // 111
	this.ToolButton_10.SetObjectName(qtcore.NewQString_5("ToolButton_10")) // 112
	this.Icon13 = qtgui.NewQIcon()
	this.Icon13.AddFile(qtcore.NewQString_5(":/icons/smile_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_10.SetIcon(this.Icon13)                                                                                        // 114
	this.ToolButton_10.SetAutoRaise(true)                                                                                          // 114

	this.HorizontalLayout_6.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_10.GetCthis())) // 115

	this.LineEdit_2 = qtwidgets.NewQLineEdit(this.LayoutWidget1)     // 111
	this.LineEdit_2.SetObjectName(qtcore.NewQString_5("LineEdit_2")) // 112

	this.HorizontalLayout_6.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.LineEdit_2.GetCthis())) // 115

	this.ToolButton_18 = qtwidgets.NewQToolButton(this.LayoutWidget1)      // 111
	this.ToolButton_18.SetObjectName(qtcore.NewQString_5("ToolButton_18")) // 112
	this.Icon14 = qtgui.NewQIcon()
	this.Icon14.AddFile(qtcore.NewQString_5(":/icons/cursor_gray64.png"), qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_18.SetIcon(this.Icon14)                                                                                         // 114
	this.ToolButton_18.SetToolButtonStyle(qtcore.Qt__ToolButtonTextBesideIcon)                                                      // 114
	this.ToolButton_18.SetAutoRaise(true)                                                                                           // 114

	this.HorizontalLayout_6.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton_18.GetCthis())) // 115

	this.VerticalLayout_5.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout_6.GetCthis()), 0) // 115

	this.Splitter.AddWidget(qtwidgets.NewQWidgetFromPointer(this.LayoutWidget1.GetCthis())) // 115

	this.VerticalLayout_6.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Splitter.GetCthis())) // 115

	this.MainWindow.SetCentralWidget(this.Centralwidget) // 114

	this.RetranslateUi(MainWindow)

	qtcore.QMetaObject_ConnectSlotsByName(qtcore.NewQObjectFromPointer(MainWindow.GetCthis())) // 100111
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
