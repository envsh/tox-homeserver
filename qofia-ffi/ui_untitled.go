package main

//  header block begin
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtmock"
import "github.com/kitech/qt.go/qtrt"

func init() { qtcore.KeepMe() }
func init() { qtgui.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }
func init() { qtrt.KeepMe() }

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
	ToolButton_33              *qtwidgets.QToolButton
	ToolButton_11              *qtwidgets.QToolButton
	HorizontalSpacer           *qtwidgets.QSpacerItem
	Label                      *qtwidgets.QLabel
	ComboBox                   *qtwidgets.QComboBox
	ToolButton_19              *qtwidgets.QToolButton
	ToolButton_12              *qtwidgets.QToolButton
	StackedWidget              *qtwidgets.QStackedWidget
	Page_4                     *qtwidgets.QWidget
	VerticalLayout_13          *qtwidgets.QVBoxLayout
	QuickWidget_2              *qtquickwidgets.QQuickWidget
	Page_3                     *qtwidgets.QWidget
	VerticalLayout_11          *qtwidgets.QVBoxLayout
	QuickWidget                *qtquickwidgets.QQuickWidget
	Page_7                     *qtwidgets.QWidget
	VerticalLayout_15          *qtwidgets.QVBoxLayout
	GroupBox                   *qtwidgets.QGroupBox
	GridLayout                 *qtwidgets.QGridLayout
	Label_13                   *qtwidgets.QLabel
	Label_8                    *qtwidgets.QLabel
	ComboBox_2                 *qtwidgets.QComboBox
	Label_9                    *qtwidgets.QLabel
	Label_10                   *qtwidgets.QLabel
	CheckBox_2                 *qtwidgets.QCheckBox
	Label_11                   *qtwidgets.QLabel
	CheckBox_3                 *qtwidgets.QCheckBox
	ComboBox_3                 *qtwidgets.QComboBox
	CheckBox                   *qtwidgets.QCheckBox
	GroupBox_2                 *qtwidgets.QGroupBox
	VerticalLayout_16          *qtwidgets.QVBoxLayout
	Label_12                   *qtwidgets.QLabel
	VerticalSpacer_3           *qtwidgets.QSpacerItem
	Page_8                     *qtwidgets.QWidget
	VerticalLayout_8           *qtwidgets.QVBoxLayout
	VerticalSpacer             *qtwidgets.QSpacerItem
	HorizontalLayout_15        *qtwidgets.QHBoxLayout
	Label_22                   *qtwidgets.QLabel
	HorizontalSpacer_3         *qtwidgets.QSpacerItem
	Label_23                   *qtwidgets.QLabel
	ComboBox_6                 *qtwidgets.QComboBox
	HorizontalLayout_20        *qtwidgets.QHBoxLayout
	HorizontalSpacer_4         *qtwidgets.QSpacerItem
	RadioButton_3              *qtwidgets.QRadioButton
	HorizontalSpacer_6         *qtwidgets.QSpacerItem
	RadioButton_4              *qtwidgets.QRadioButton
	HorizontalSpacer_5         *qtwidgets.QSpacerItem
	PushButton_7               *qtwidgets.QPushButton
	Label_24                   *qtwidgets.QLabel
	VerticalSpacer_2           *qtwidgets.QSpacerItem
	Page                       *qtwidgets.QWidget
	VerticalLayout_2           *qtwidgets.QVBoxLayout
	Widget                     *qtwidgets.QWidget
	HorizontalLayout_17        *qtwidgets.QHBoxLayout
	ToolButton_17              *qtwidgets.QToolButton
	VerticalLayout             *qtwidgets.QVBoxLayout
	HorizontalLayout           *qtwidgets.QHBoxLayout
	Label_2                    *qtwidgets.QLabel
	LineEdit_5                 *qtwidgets.QLineEdit
	HorizontalLayout_16        *qtwidgets.QHBoxLayout
	Label_3                    *qtwidgets.QLabel
	LineEdit_6                 *qtwidgets.QLineEdit
	ToolButton                 *qtwidgets.QToolButton
	ToolButton_32              *qtwidgets.QToolButton
	HorizontalLayout_2         *qtwidgets.QHBoxLayout
	LineEdit                   *qtwidgets.QLineEdit
	ToolButton_2               *qtwidgets.QToolButton
	ToolButton_3               *qtwidgets.QToolButton
	ScrollArea                 *qtwidgets.QScrollArea
	ScrollAreaWidgetContents   *qtwidgets.QWidget
	VerticalLayout_9           *qtwidgets.QVBoxLayout
	HorizontalLayout_3         *qtwidgets.QHBoxLayout
	ToolButton_4               *qtwidgets.QToolButton
	ToolButton_5               *qtwidgets.QToolButton
	ToolButton_6               *qtwidgets.QToolButton
	ToolButton_7               *qtwidgets.QToolButton
	Page_2                     *qtwidgets.QWidget
	VerticalLayout_14          *qtwidgets.QVBoxLayout
	Widget_2                   *qtwidgets.QWidget
	HorizontalLayout_8         *qtwidgets.QHBoxLayout
	Label_4                    *qtwidgets.QLabel
	VerticalLayout_4           *qtwidgets.QVBoxLayout
	HorizontalLayout_5         *qtwidgets.QHBoxLayout
	Label_5                    *qtwidgets.QLabel
	LabelMsgCount              *qtwidgets.QLabel
	HorizontalLayout_4         *qtwidgets.QHBoxLayout
	Label_6                    *qtwidgets.QLabel
	Label_7                    *qtwidgets.QLabel
	VerticalLayout_7           *qtwidgets.QVBoxLayout
	ToolButton_15              *qtwidgets.QToolButton
	ToolButton_16              *qtwidgets.QToolButton
	ToolButton_13              *qtwidgets.QToolButton
	ToolButton_14              *qtwidgets.QToolButton
	ToolButton_22              *qtwidgets.QToolButton
	HorizontalLayout_9         *qtwidgets.QHBoxLayout
	LineHeadSepLeft            *qtwidgets.QFrame
	ToolButton_24              *qtwidgets.QToolButton
	ToolButton_25              *qtwidgets.QToolButton
	ToolButton_26              *qtwidgets.QToolButton
	ToolButton_27              *qtwidgets.QToolButton
	ToolButton_28              *qtwidgets.QToolButton
	ToolButton_29              *qtwidgets.QToolButton
	LabelMsgCount2             *qtwidgets.QLabel
	ToolButton_23              *qtwidgets.QToolButton
	LineHeadSepRight           *qtwidgets.QFrame
	ScrollArea_2               *qtwidgets.QScrollArea
	ScrollAreaWidgetContents_2 *qtwidgets.QWidget
	VerticalLayout_3           *qtwidgets.QVBoxLayout
	HorizontalLayout_6         *qtwidgets.QHBoxLayout
	ToolButton_8               *qtwidgets.QToolButton
	ToolButton_9               *qtwidgets.QToolButton
	ToolButton_10              *qtwidgets.QToolButton
	LineEdit_2                 *qtwidgets.QLineEdit
	ToolButton_18              *qtwidgets.QToolButton
	Page_10                    *qtwidgets.QWidget
	VerticalLayout_22          *qtwidgets.QVBoxLayout
	HorizontalLayout_21        *qtwidgets.QHBoxLayout
	HorizontalSpacer_2         *qtwidgets.QSpacerItem
	Label_27                   *qtwidgets.QLabel
	Label_28                   *qtwidgets.QLabel
	Label_29                   *qtwidgets.QLabel
	HorizontalSpacer_7         *qtwidgets.QSpacerItem
	Widget_3                   *qtwidgets.QWidget
	HorizontalLayout_22        *qtwidgets.QHBoxLayout
	HorizontalSpacer_8         *qtwidgets.QSpacerItem
	ToolButton_20              *qtwidgets.QToolButton
	HorizontalSpacer_10        *qtwidgets.QSpacerItem
	ToolButton_21              *qtwidgets.QToolButton
	HorizontalSpacer_11        *qtwidgets.QSpacerItem
	ToolButton_30              *qtwidgets.QToolButton
	HorizontalSpacer_9         *qtwidgets.QSpacerItem
	Page_add_group             *qtwidgets.QWidget
	VerticalLayout_21          *qtwidgets.QVBoxLayout
	HorizontalLayout_13        *qtwidgets.QHBoxLayout
	PushButton_5               *qtwidgets.QPushButton
	Label_19                   *qtwidgets.QLabel
	PushButton_6               *qtwidgets.QPushButton
	Label_20                   *qtwidgets.QLabel
	ComboBox_5                 *qtwidgets.QComboBox
	RadioButton                *qtwidgets.QRadioButton
	RadioButton_2              *qtwidgets.QRadioButton
	HorizontalLayout_14        *qtwidgets.QHBoxLayout
	Label_21                   *qtwidgets.QLabel
	ComboBox_4                 *qtwidgets.QComboBox
	VerticalSpacer_6           *qtwidgets.QSpacerItem
	Page_add_friend            *qtwidgets.QWidget
	VerticalLayout_20          *qtwidgets.QVBoxLayout
	HorizontalLayout_12        *qtwidgets.QHBoxLayout
	PushButton_3               *qtwidgets.QPushButton
	Label_16                   *qtwidgets.QLabel
	PushButton_4               *qtwidgets.QPushButton
	VerticalLayout_19          *qtwidgets.QVBoxLayout
	Label_17                   *qtwidgets.QLabel
	LineEdit_4                 *qtwidgets.QLineEdit
	Label_18                   *qtwidgets.QLabel
	TextEdit                   *qtwidgets.QTextEdit
	VerticalSpacer_5           *qtwidgets.QSpacerItem
	Page_invite_friend         *qtwidgets.QWidget
	VerticalLayout_18          *qtwidgets.QVBoxLayout
	HorizontalLayout_10        *qtwidgets.QHBoxLayout
	PushButton                 *qtwidgets.QPushButton
	Label_14                   *qtwidgets.QLabel
	PushButton_2               *qtwidgets.QPushButton
	HorizontalLayout_11        *qtwidgets.QHBoxLayout
	Label_15                   *qtwidgets.QLabel
	LineEdit_3                 *qtwidgets.QLineEdit
	ScrollArea_3               *qtwidgets.QScrollArea
	ScrollAreaWidgetContents_3 *qtwidgets.QWidget
	VerticalLayout_17          *qtwidgets.QVBoxLayout
	TableWidget                *qtwidgets.QTableWidget
	Page_9                     *qtwidgets.QWidget
	VerticalLayout_10          *qtwidgets.QVBoxLayout
	HorizontalLayout_19        *qtwidgets.QHBoxLayout
	PushButton_8               *qtwidgets.QPushButton
	Label_26                   *qtwidgets.QLabel
	PushButton_9               *qtwidgets.QPushButton
	HorizontalLayout_18        *qtwidgets.QHBoxLayout
	Label_25                   *qtwidgets.QLabel
	LineEdit_7                 *qtwidgets.QLineEdit
	TableWidget_2              *qtwidgets.QTableWidget
	Page_11                    *qtwidgets.QWidget
	VerticalLayout_26          *qtwidgets.QVBoxLayout
	HorizontalLayout_23        *qtwidgets.QHBoxLayout
	PushButton_10              *qtwidgets.QPushButton
	Label_30                   *qtwidgets.QLabel
	PushButton_11              *qtwidgets.QPushButton
	HorizontalLayout_24        *qtwidgets.QHBoxLayout
	ToolButton_31              *qtwidgets.QToolButton
	VerticalLayout_23          *qtwidgets.QVBoxLayout
	Label_32                   *qtwidgets.QLabel
	LineEdit_9                 *qtwidgets.QLineEdit
	CheckBox_4                 *qtwidgets.QCheckBox
	GroupBox_3                 *qtwidgets.QGroupBox
	HorizontalLayout_25        *qtwidgets.QHBoxLayout
	RadioButton_5              *qtwidgets.QRadioButton
	RadioButton_6              *qtwidgets.QRadioButton
	RadioButton_7              *qtwidgets.QRadioButton
	GroupBox_4                 *qtwidgets.QGroupBox
	VerticalLayout_24          *qtwidgets.QVBoxLayout
	CheckBox_5                 *qtwidgets.QCheckBox
	HorizontalLayout_26        *qtwidgets.QHBoxLayout
	Label_34                   *qtwidgets.QLabel
	LineEdit_8                 *qtwidgets.QLineEdit
	PushButton_12              *qtwidgets.QPushButton
	GroupBox_5                 *qtwidgets.QGroupBox
	VerticalLayout_25          *qtwidgets.QVBoxLayout
	TextEdit_2                 *qtwidgets.QTextEdit
	Page_6                     *qtwidgets.QWidget
	VerticalLayout_5           *qtwidgets.QVBoxLayout
	ListWidget_2               *qtwidgets.QListWidget
	ListWidget                 *qtwidgets.QListWidget
	Page_5                     *qtwidgets.QWidget
	VerticalLayout_12          *qtwidgets.QVBoxLayout
	TextBrowser                *qtwidgets.QTextBrowser
	Page_12                    *qtwidgets.QWidget
	VerticalLayout_28          *qtwidgets.QVBoxLayout
	GroupBox_6                 *qtwidgets.QGroupBox
	GridLayout_2               *qtwidgets.QGridLayout
	Label_31                   *qtwidgets.QLabel
	Label_37                   *qtwidgets.QLabel
	Label_33                   *qtwidgets.QLabel
	Label_38                   *qtwidgets.QLabel
	Label_35                   *qtwidgets.QLabel
	Label_39                   *qtwidgets.QLabel
	Label_36                   *qtwidgets.QLabel
	Label_40                   *qtwidgets.QLabel
	Label_41                   *qtwidgets.QLabel
	Label_42                   *qtwidgets.QLabel
	GroupBox_7                 *qtwidgets.QGroupBox
	VerticalLayout_27          *qtwidgets.QVBoxLayout
	HorizontalLayout_28        *qtwidgets.QHBoxLayout
	Label_46                   *qtwidgets.QLabel
	Label_45                   *qtwidgets.QLabel
	Label_53                   *qtwidgets.QLabel
	Label_54                   *qtwidgets.QLabel
	Label_52                   *qtwidgets.QLabel
	Label_50                   *qtwidgets.QLabel
	Label_44                   *qtwidgets.QLabel
	Label_47                   *qtwidgets.QLabel
	HorizontalLayout_27        *qtwidgets.QHBoxLayout
	Label_43                   *qtwidgets.QLabel
	Label_48                   *qtwidgets.QLabel
	VerticalSpacer_4           *qtwidgets.QSpacerItem
	MainWindow                 *qtwidgets.QMainWindow
	Icon                       *qtgui.QIcon // 116
	SizePolicy                 *qtwidgets.QSizePolicy
	Icon1                      *qtgui.QIcon // 116
	Icon2                      *qtgui.QIcon // 116
	SizePolicy1                *qtwidgets.QSizePolicy
	Font                       *qtgui.QFont // 116
	Icon3                      *qtgui.QIcon // 116
	Icon4                      *qtgui.QIcon // 116
	Icon5                      *qtgui.QIcon // 116
	SizePolicy2                *qtwidgets.QSizePolicy
	Icon6                      *qtgui.QIcon // 116
	Icon7                      *qtgui.QIcon // 116
	Icon8                      *qtgui.QIcon // 116
	Icon9                      *qtgui.QIcon // 116
	Font1                      *qtgui.QFont // 116
	Icon10                     *qtgui.QIcon // 116
	Icon11                     *qtgui.QIcon // 116
	Icon12                     *qtgui.QIcon // 116
	Icon13                     *qtgui.QIcon // 116
	Icon14                     *qtgui.QIcon // 116
	SizePolicy3                *qtwidgets.QSizePolicy
	SizePolicy4                *qtwidgets.QSizePolicy
	SizePolicy5                *qtwidgets.QSizePolicy
	Icon15                     *qtgui.QIcon // 116
	Icon16                     *qtgui.QIcon // 116
	Icon17                     *qtgui.QIcon // 116
	Icon18                     *qtgui.QIcon // 116
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
	MainWindow.Resize(368, 599)
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
	this.VerticalLayout_6.SetSpacing(1)                                    // 114
	this.VerticalLayout_6.SetObjectName("VerticalLayout_6")                // 112
	this.VerticalLayout_6.SetContentsMargins(1, 1, 1, 1)                   // 114
	this.HorizontalLayout_7 = qtwidgets.NewQHBoxLayout()                   // 111
	this.HorizontalLayout_7.SetSpacing(0)                                  // 114
	this.HorizontalLayout_7.SetObjectName("HorizontalLayout_7")            // 112
	this.ToolButton_33 = qtwidgets.NewQToolButton(this.Centralwidget)      // 111
	this.ToolButton_33.SetObjectName("ToolButton_33")                      // 112
	this.ToolButton_33.SetAutoRaise(true)                                  // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.ToolButton_33) // 115

	this.ToolButton_11 = qtwidgets.NewQToolButton(this.Centralwidget) // 111
	this.ToolButton_11.SetObjectName("ToolButton_11")                 // 112
	this.ToolButton_11.SetFocusPolicy(qtcore.Qt__NoFocus)             // 114
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(":/icons/barbuttonicon_back_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_11.SetIcon(this.Icon)                                                                                // 114
	this.ToolButton_11.SetAutoRaise(true)                                                                                // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.ToolButton_11) // 115

	this.HorizontalSpacer = qtwidgets.NewQSpacerItem(3, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer)

	this.HorizontalLayout_7.AddItem(this.HorizontalSpacer) // 115

	this.Label = qtwidgets.NewQLabel(this.Centralwidget, 0)                                                  // 111
	this.Label.SetObjectName("Label")                                                                        // 112
	this.Label.SetAlignment(qtcore.Qt__AlignCenter)                                                          // 114
	this.Label.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.Label) // 115

	this.ComboBox = qtwidgets.NewQComboBox(this.Centralwidget) // 111
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.AddItem("", qtcore.NewQVariant_12("wtf"))    // 115
	this.ComboBox.SetObjectName("ComboBox")                    // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                           // 114
	this.SizePolicy.SetVerticalStretch(0)                                             // 114
	this.SizePolicy.SetHeightForWidth(this.ComboBox.SizePolicy().HasHeightForWidth()) // 114
	this.ComboBox.SetSizePolicy(this.SizePolicy)                                      // 114
	this.ComboBox.SetFocusPolicy(qtcore.Qt__NoFocus)                                  // 114
	this.ComboBox.SetMaxVisibleItems(12)                                              // 114
	this.ComboBox.SetSizeAdjustPolicy(qtwidgets.QComboBox__AdjustToContents)          // 114
	this.ComboBox.SetFrame(false)                                                     // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.ComboBox) // 115

	this.ToolButton_19 = qtwidgets.NewQToolButton(this.Centralwidget) // 111
	this.ToolButton_19.SetObjectName("ToolButton_19")                 // 112
	this.ToolButton_19.SetFocusPolicy(qtcore.Qt__NoFocus)             // 114
	this.ToolButton_19.SetAutoRaise(true)                             // 114
	this.ToolButton_19.SetArrowType(qtcore.Qt__NoArrow)               // 114

	this.HorizontalLayout_7.Layout().AddWidget(this.ToolButton_19) // 115

	this.ToolButton_12 = qtwidgets.NewQToolButton(this.Centralwidget) // 111
	this.ToolButton_12.SetObjectName("ToolButton_12")                 // 112
	this.ToolButton_12.SetFocusPolicy(qtcore.Qt__NoFocus)             // 114
	this.Icon1 = qtgui.NewQIcon()
	this.Icon1.AddFile(":/icons/barbuttonicon_forward_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_12.SetIcon(this.Icon1)                                                                                   // 114
	this.ToolButton_12.SetAutoRaise(true)                                                                                    // 114

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

	this.StackedWidget.AddWidget(this.Page_3)                        // 115
	this.Page_7 = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_7.SetObjectName("Page_7")                              // 112
	this.VerticalLayout_15 = qtwidgets.NewQVBoxLayout_1(this.Page_7) // 111
	this.VerticalLayout_15.SetObjectName("VerticalLayout_15")        // 112
	this.GroupBox = qtwidgets.NewQGroupBox(this.Page_7)              // 111
	this.GroupBox.SetObjectName("GroupBox")                          // 112
	this.GroupBox.SetFlat(true)                                      // 114
	this.GridLayout = qtwidgets.NewQGridLayout(this.GroupBox)        // 111
	this.GridLayout.SetSpacing(12)                                   // 114
	this.GridLayout.SetObjectName("GridLayout")                      // 112
	this.GridLayout.SetContentsMargins(12, 12, 12, 12)               // 114
	this.Label_13 = qtwidgets.NewQLabel(this.GroupBox, 0)            // 111
	this.Label_13.SetObjectName("Label_13")                          // 112

	this.GridLayout.AddWidget_2(this.Label_13, 4, 0, 1, 1, 0) // 115

	this.Label_8 = qtwidgets.NewQLabel(this.GroupBox, 0) // 111
	this.Label_8.SetObjectName("Label_8")                // 112

	this.GridLayout.AddWidget_2(this.Label_8, 0, 0, 1, 1, 0) // 115

	this.ComboBox_2 = qtwidgets.NewQComboBox(this.GroupBox)   // 111
	this.ComboBox_2.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_2.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_2.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_2.SetObjectName("ComboBox_2")               // 112
	this.ComboBox_2.SetFrame(false)                           // 114

	this.GridLayout.AddWidget_2(this.ComboBox_2, 0, 1, 1, 1, 0) // 115

	this.Label_9 = qtwidgets.NewQLabel(this.GroupBox, 0) // 111
	this.Label_9.SetObjectName("Label_9")                // 112

	this.GridLayout.AddWidget_2(this.Label_9, 1, 0, 1, 1, 0) // 115

	this.Label_10 = qtwidgets.NewQLabel(this.GroupBox, 0) // 111
	this.Label_10.SetObjectName("Label_10")               // 112

	this.GridLayout.AddWidget_2(this.Label_10, 2, 0, 1, 1, 0) // 115

	this.CheckBox_2 = qtwidgets.NewQCheckBox(this.GroupBox) // 111
	this.CheckBox_2.SetObjectName("CheckBox_2")             // 112
	this.CheckBox_2.SetChecked(true)                        // 114

	this.GridLayout.AddWidget_2(this.CheckBox_2, 2, 1, 1, 1, 0) // 115

	this.Label_11 = qtwidgets.NewQLabel(this.GroupBox, 0) // 111
	this.Label_11.SetObjectName("Label_11")               // 112

	this.GridLayout.AddWidget_2(this.Label_11, 3, 0, 1, 1, 0) // 115

	this.CheckBox_3 = qtwidgets.NewQCheckBox(this.GroupBox) // 111
	this.CheckBox_3.SetObjectName("CheckBox_3")             // 112
	this.CheckBox_3.SetChecked(true)                        // 114

	this.GridLayout.AddWidget_2(this.CheckBox_3, 3, 1, 1, 1, 0) // 115

	this.ComboBox_3 = qtwidgets.NewQComboBox(this.GroupBox)   // 111
	this.ComboBox_3.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_3.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_3.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_3.SetObjectName("ComboBox_3")               // 112
	this.ComboBox_3.SetEditable(true)                         // 114

	this.GridLayout.AddWidget_2(this.ComboBox_3, 1, 1, 1, 1, 0) // 115

	this.CheckBox = qtwidgets.NewQCheckBox(this.GroupBox) // 111
	this.CheckBox.SetObjectName("CheckBox")               // 112
	this.CheckBox.SetChecked(true)                        // 114

	this.GridLayout.AddWidget_2(this.CheckBox, 4, 1, 1, 1, 0) // 115

	this.VerticalLayout_15.Layout().AddWidget(this.GroupBox) // 115

	this.GroupBox_2 = qtwidgets.NewQGroupBox(this.Page_7)                // 111
	this.GroupBox_2.SetObjectName("GroupBox_2")                          // 112
	this.GroupBox_2.SetFlat(true)                                        // 114
	this.VerticalLayout_16 = qtwidgets.NewQVBoxLayout_1(this.GroupBox_2) // 111
	this.VerticalLayout_16.SetSpacing(12)                                // 114
	this.VerticalLayout_16.SetObjectName("VerticalLayout_16")            // 112
	this.VerticalLayout_16.SetContentsMargins(12, 12, 12, 12)            // 114
	this.Label_12 = qtwidgets.NewQLabel(this.GroupBox_2, 0)              // 111
	this.Label_12.SetObjectName("Label_12")                              // 112

	this.VerticalLayout_16.Layout().AddWidget(this.Label_12) // 115

	this.VerticalLayout_15.Layout().AddWidget(this.GroupBox_2) // 115

	this.VerticalSpacer_3 = qtwidgets.NewQSpacerItem(20, 287, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_3)

	this.VerticalLayout_15.AddItem(this.VerticalSpacer_3) // 115

	this.StackedWidget.AddWidget(this.Page_7)                       // 115
	this.Page_8 = qtwidgets.NewQWidget(nil, 0)                      // 111
	this.Page_8.SetObjectName("Page_8")                             // 112
	this.VerticalLayout_8 = qtwidgets.NewQVBoxLayout_1(this.Page_8) // 111
	this.VerticalLayout_8.SetSpacing(22)                            // 114
	this.VerticalLayout_8.SetObjectName("VerticalLayout_8")         // 112
	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 168, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer)

	this.VerticalLayout_8.AddItem(this.VerticalSpacer) // 115

	this.HorizontalLayout_15 = qtwidgets.NewQHBoxLayout()                                                       // 111
	this.HorizontalLayout_15.SetObjectName("HorizontalLayout_15")                                               // 112
	this.Label_22 = qtwidgets.NewQLabel(this.Page_8, 0)                                                         // 111
	this.Label_22.SetObjectName("Label_22")                                                                     // 112
	this.Label_22.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_15.Layout().AddWidget(this.Label_22) // 115

	this.HorizontalSpacer_3 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_3)

	this.HorizontalLayout_15.AddItem(this.HorizontalSpacer_3) // 115

	this.Label_23 = qtwidgets.NewQLabel(this.Page_8, 0) // 111
	this.Label_23.SetObjectName("Label_23")             // 112

	this.HorizontalLayout_15.Layout().AddWidget(this.Label_23) // 115

	this.VerticalLayout_8.AddLayout(this.HorizontalLayout_15, 0) // 115

	this.ComboBox_6 = qtwidgets.NewQComboBox(this.Page_8)     // 111
	this.ComboBox_6.AddItem("", qtcore.NewQVariant_12("wtf")) // 115
	this.ComboBox_6.SetObjectName("ComboBox_6")               // 112
	this.ComboBox_6.SetEditable(true)                         // 114
	this.ComboBox_6.SetFrame(true)                            // 114
	this.ComboBox_6.SetModelColumn(0)                         // 114

	this.VerticalLayout_8.Layout().AddWidget(this.ComboBox_6) // 115

	this.HorizontalLayout_20 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_20.SetObjectName("HorizontalLayout_20") // 112
	this.HorizontalSpacer_4 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_4)

	this.HorizontalLayout_20.AddItem(this.HorizontalSpacer_4) // 115

	this.RadioButton_3 = qtwidgets.NewQRadioButton(this.Page_8) // 111
	this.RadioButton_3.SetObjectName("RadioButton_3")           // 112
	this.RadioButton_3.SetChecked(true)                         // 114

	this.HorizontalLayout_20.Layout().AddWidget(this.RadioButton_3) // 115

	this.HorizontalSpacer_6 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_6)

	this.HorizontalLayout_20.AddItem(this.HorizontalSpacer_6) // 115

	this.RadioButton_4 = qtwidgets.NewQRadioButton(this.Page_8) // 111
	this.RadioButton_4.SetObjectName("RadioButton_4")           // 112
	this.RadioButton_4.SetEnabled(false)                        // 114

	this.HorizontalLayout_20.Layout().AddWidget(this.RadioButton_4) // 115

	this.HorizontalSpacer_5 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_5)

	this.HorizontalLayout_20.AddItem(this.HorizontalSpacer_5) // 115

	this.VerticalLayout_8.AddLayout(this.HorizontalLayout_20, 0) // 115

	this.PushButton_7 = qtwidgets.NewQPushButton(this.Page_8) // 111
	this.PushButton_7.SetObjectName("PushButton_7")           // 112
	this.PushButton_7.SetFlat(true)                           // 114

	this.VerticalLayout_8.Layout().AddWidget(this.PushButton_7) // 115

	this.Label_24 = qtwidgets.NewQLabel(this.Page_8, 0)                                                         // 111
	this.Label_24.SetObjectName("Label_24")                                                                     // 112
	this.Label_24.SetWordWrap(true)                                                                             // 114
	this.Label_24.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout_8.Layout().AddWidget(this.Label_24) // 115

	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 238, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_2)

	this.VerticalLayout_8.AddItem(this.VerticalSpacer_2) // 115

	this.StackedWidget.AddWidget(this.Page_8)                          // 115
	this.Page = qtwidgets.NewQWidget(nil, 0)                           // 111
	this.Page.SetObjectName("Page")                                    // 112
	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout_1(this.Page)      // 111
	this.VerticalLayout_2.SetSpacing(0)                                // 114
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")            // 112
	this.VerticalLayout_2.SetContentsMargins(0, 0, 0, 0)               // 114
	this.Widget = qtwidgets.NewQWidget(this.Page, 0)                   // 111
	this.Widget.SetObjectName("Widget")                                // 112
	this.HorizontalLayout_17 = qtwidgets.NewQHBoxLayout_1(this.Widget) // 111
	this.HorizontalLayout_17.SetSpacing(3)                             // 114
	this.HorizontalLayout_17.SetObjectName("HorizontalLayout_17")      // 112
	this.HorizontalLayout_17.SetContentsMargins(0, 3, 0, 3)            // 114
	this.ToolButton_17 = qtwidgets.NewQToolButton(this.Widget)         // 111
	this.ToolButton_17.SetObjectName("ToolButton_17")                  // 112
	this.ToolButton_17.SetMaximumSize_1(32, 32)                        // 113
	this.ToolButton_17.SetFocusPolicy(qtcore.Qt__NoFocus)              // 114
	this.Icon2 = qtgui.NewQIcon()
	this.Icon2.AddFile(":/icons/icon_avatar_40.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_17.SetIcon(this.Icon2)                                                                     // 114
	this.ToolButton_17.SetIconSize(qtcore.NewQSize_1(32, 32))                                                  // 113
	this.ToolButton_17.SetAutoRaise(true)                                                                      // 114

	this.HorizontalLayout_17.Layout().AddWidget(this.ToolButton_17) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()        // 111
	this.VerticalLayout.SetSpacing(0)                       // 114
	this.VerticalLayout.SetObjectName("VerticalLayout")     // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()      // 111
	this.HorizontalLayout.SetSpacing(0)                     // 114
	this.HorizontalLayout.SetObjectName("HorizontalLayout") // 112
	this.Label_2 = qtwidgets.NewQLabel(this.Widget, 0)      // 111
	this.Label_2.SetObjectName("Label_2")                   // 112
	this.SizePolicy1 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy1.SetHorizontalStretch(0)                                          // 114
	this.SizePolicy1.SetVerticalStretch(0)                                            // 114
	this.SizePolicy1.SetHeightForWidth(this.Label_2.SizePolicy().HasHeightForWidth()) // 114
	this.Label_2.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Font = qtgui.NewQFont()
	this.Font.SetPointSize(12)                                             // 114
	this.Font.SetBold(true)                                                // 114
	this.Font.SetWeight(75)                                                // 114
	this.Label_2.SetFont(this.Font)                                        // 114
	this.Label_2.SetTextInteractionFlags(qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(this.Label_2) // 115

	this.LineEdit_5 = qtwidgets.NewQLineEdit(this.Widget) // 111
	this.LineEdit_5.SetObjectName("LineEdit_5")           // 112

	this.HorizontalLayout.Layout().AddWidget(this.LineEdit_5) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout, 0) // 115

	this.HorizontalLayout_16 = qtwidgets.NewQHBoxLayout()                             // 111
	this.HorizontalLayout_16.SetSpacing(0)                                            // 114
	this.HorizontalLayout_16.SetObjectName("HorizontalLayout_16")                     // 112
	this.Label_3 = qtwidgets.NewQLabel(this.Widget, 0)                                // 111
	this.Label_3.SetObjectName("Label_3")                                             // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_3.SizePolicy().HasHeightForWidth()) // 114
	this.Label_3.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__TextSelectableByMouse)            // 114

	this.HorizontalLayout_16.Layout().AddWidget(this.Label_3) // 115

	this.LineEdit_6 = qtwidgets.NewQLineEdit(this.Widget) // 111
	this.LineEdit_6.SetObjectName("LineEdit_6")           // 112

	this.HorizontalLayout_16.Layout().AddWidget(this.LineEdit_6) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout_16, 0) // 115

	this.HorizontalLayout_17.AddLayout(this.VerticalLayout, 0) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.Widget) // 111
	this.ToolButton.SetObjectName("ToolButton")             // 112
	this.ToolButton.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon3 = qtgui.NewQIcon()
	this.Icon3.AddFile(":/icons/online_30.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton.SetIcon(this.Icon3)                                                                   // 114
	this.ToolButton.SetToolButtonStyle(qtcore.Qt__ToolButtonIconOnly)                                     // 114
	this.ToolButton.SetAutoRaise(true)                                                                    // 114

	this.HorizontalLayout_17.Layout().AddWidget(this.ToolButton) // 115

	this.ToolButton_32 = qtwidgets.NewQToolButton(this.Widget) // 111
	this.ToolButton_32.SetObjectName("ToolButton_32")          // 112
	this.ToolButton_32.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon4 = qtgui.NewQIcon()
	this.Icon4.AddFile(":/icons/power-button-off.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_32.SetIcon(this.Icon4)                                                                       // 114
	this.ToolButton_32.SetIconSize(qtcore.NewQSize_1(28, 28))                                                    // 113
	this.ToolButton_32.SetToolButtonStyle(qtcore.Qt__ToolButtonIconOnly)                                         // 114
	this.ToolButton_32.SetAutoRaise(true)                                                                        // 114

	this.HorizontalLayout_17.Layout().AddWidget(this.ToolButton_32) // 115

	this.VerticalLayout_2.Layout().AddWidget(this.Widget) // 115

	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_2.SetSpacing(0)                       // 114
	this.HorizontalLayout_2.SetObjectName("HorizontalLayout_2") // 112
	this.HorizontalLayout_2.SetContentsMargins(6, -1, -1, -1)   // 114
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
	this.Icon5 = qtgui.NewQIcon()
	this.Icon5.AddFile(":/icons/remove-symbol_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_3.SetIcon(this.Icon5)                                                                            // 114
	this.ToolButton_3.SetAutoRaise(true)                                                                             // 114

	this.HorizontalLayout_2.Layout().AddWidget(this.ToolButton_3) // 115

	this.VerticalLayout_2.AddLayout(this.HorizontalLayout_2, 0) // 115

	this.ScrollArea = qtwidgets.NewQScrollArea(this.Page) // 111
	this.ScrollArea.SetObjectName("ScrollArea")           // 112
	this.SizePolicy2 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Expanding, 1)
	this.SizePolicy2.SetHorizontalStretch(0)                                                           // 114
	this.SizePolicy2.SetVerticalStretch(0)                                                             // 114
	this.SizePolicy2.SetHeightForWidth(this.ScrollArea.SizePolicy().HasHeightForWidth())               // 114
	this.ScrollArea.SetSizePolicy(this.SizePolicy2)                                                    // 114
	this.ScrollArea.SetFocusPolicy(qtcore.Qt__NoFocus)                                                 // 114
	this.ScrollArea.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)                        // 114
	this.ScrollArea.SetWidgetResizable(true)                                                           // 114
	this.ScrollArea.SetAlignment(qtcore.Qt__AlignLeading | qtcore.Qt__AlignLeft | qtcore.Qt__AlignTop) // 114
	this.ScrollAreaWidgetContents = qtwidgets.NewQWidget(nil, 0)                                       // 111
	this.ScrollAreaWidgetContents.SetObjectName("ScrollAreaWidgetContents")                            // 112
	this.ScrollAreaWidgetContents.SetGeometry(0, 0, 364, 16)                                           // 114
	this.SizePolicy.SetHeightForWidth(this.ScrollAreaWidgetContents.SizePolicy().HasHeightForWidth())  // 114
	this.ScrollAreaWidgetContents.SetSizePolicy(this.SizePolicy)                                       // 114
	this.VerticalLayout_9 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents)                  // 111
	this.VerticalLayout_9.SetSpacing(0)                                                                // 114
	this.VerticalLayout_9.SetObjectName("VerticalLayout_9")                                            // 112
	this.VerticalLayout_9.SetContentsMargins(0, 0, 0, 0)                                               // 114
	this.ScrollArea.SetWidget(this.ScrollAreaWidgetContents)                                           // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ScrollArea) // 115

	this.HorizontalLayout_3 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_3.SetSpacing(0)                       // 114
	this.HorizontalLayout_3.SetObjectName("HorizontalLayout_3") // 112
	this.ToolButton_4 = qtwidgets.NewQToolButton(this.Page)     // 111
	this.ToolButton_4.SetObjectName("ToolButton_4")             // 112
	this.Icon6 = qtgui.NewQIcon()
	this.Icon6.AddFile(":/icons/add-square-button-gray.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_4.SetIcon(this.Icon6)                                                                              // 114
	this.ToolButton_4.SetIconSize(qtcore.NewQSize_1(22, 22))                                                           // 113
	this.ToolButton_4.SetAutoRaise(true)                                                                               // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_4) // 115

	this.ToolButton_5 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_5.SetObjectName("ToolButton_5")         // 112
	this.Icon7 = qtgui.NewQIcon()
	this.Icon7.AddFile(":/icons/groupgray.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_5.SetIcon(this.Icon7)                                                                 // 114
	this.ToolButton_5.SetIconSize(qtcore.NewQSize_1(22, 22))                                              // 113
	this.ToolButton_5.SetAutoRaise(true)                                                                  // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_5) // 115

	this.ToolButton_6 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_6.SetObjectName("ToolButton_6")         // 112
	this.Icon8 = qtgui.NewQIcon()
	this.Icon8.AddFile(":/icons/transfer_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_6.SetIcon(this.Icon8)                                                                       // 114
	this.ToolButton_6.SetIconSize(qtcore.NewQSize_1(22, 22))                                                    // 113
	this.ToolButton_6.SetAutoRaise(true)                                                                        // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_6) // 115

	this.ToolButton_7 = qtwidgets.NewQToolButton(this.Page) // 111
	this.ToolButton_7.SetObjectName("ToolButton_7")         // 112
	this.Icon9 = qtgui.NewQIcon()
	this.Icon9.AddFile(":/icons/settings_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_7.SetIcon(this.Icon9)                                                                       // 114
	this.ToolButton_7.SetIconSize(qtcore.NewQSize_1(22, 22))                                                    // 113
	this.ToolButton_7.SetAutoRaise(true)                                                                        // 114

	this.HorizontalLayout_3.Layout().AddWidget(this.ToolButton_7) // 115

	this.VerticalLayout_2.AddLayout(this.HorizontalLayout_3, 0) // 115

	this.StackedWidget.AddWidget(this.Page)                                                 // 115
	this.Page_2 = qtwidgets.NewQWidget(nil, 0)                                              // 111
	this.Page_2.SetObjectName("Page_2")                                                     // 112
	this.VerticalLayout_14 = qtwidgets.NewQVBoxLayout_1(this.Page_2)                        // 111
	this.VerticalLayout_14.SetObjectName("VerticalLayout_14")                               // 112
	this.Widget_2 = qtwidgets.NewQWidget(this.Page_2, 0)                                    // 111
	this.Widget_2.SetObjectName("Widget_2")                                                 // 112
	this.HorizontalLayout_8 = qtwidgets.NewQHBoxLayout_1(this.Widget_2)                     // 111
	this.HorizontalLayout_8.SetObjectName("HorizontalLayout_8")                             // 112
	this.HorizontalLayout_8.SetContentsMargins(-1, -1, -1, 0)                               // 114
	this.Label_4 = qtwidgets.NewQLabel(this.Widget_2, 0)                                    // 111
	this.Label_4.SetObjectName("Label_4")                                                   // 112
	this.Label_4.SetMaximumSize_1(32, 32)                                                   // 113
	this.Label_4.SetPixmap(qtgui.NewQPixmap_3(":/icons/icon_avatar_40.png", "dummy123", 0)) // 114

	this.HorizontalLayout_8.Layout().AddWidget(this.Label_4) // 115

	this.VerticalLayout_4 = qtwidgets.NewQVBoxLayout()                                // 111
	this.VerticalLayout_4.SetObjectName("VerticalLayout_4")                           // 112
	this.HorizontalLayout_5 = qtwidgets.NewQHBoxLayout()                              // 111
	this.HorizontalLayout_5.SetObjectName("HorizontalLayout_5")                       // 112
	this.Label_5 = qtwidgets.NewQLabel(this.Widget_2, 0)                              // 111
	this.Label_5.SetObjectName("Label_5")                                             // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_5.SizePolicy().HasHeightForWidth()) // 114
	this.Label_5.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Font1 = qtgui.NewQFont()
	this.Font1.SetBold(true)                                                                                   // 114
	this.Font1.SetWeight(75)                                                                                   // 114
	this.Label_5.SetFont(this.Font1)                                                                           // 114
	this.Label_5.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_5.Layout().AddWidget(this.Label_5) // 115

	this.LabelMsgCount = qtwidgets.NewQLabel(this.Widget_2, 0) // 111
	this.LabelMsgCount.SetObjectName("LabelMsgCount")          // 112

	this.HorizontalLayout_5.Layout().AddWidget(this.LabelMsgCount) // 115

	this.VerticalLayout_4.AddLayout(this.HorizontalLayout_5, 0) // 115

	this.HorizontalLayout_4 = qtwidgets.NewQHBoxLayout()                                                       // 111
	this.HorizontalLayout_4.SetSpacing(0)                                                                      // 114
	this.HorizontalLayout_4.SetObjectName("HorizontalLayout_4")                                                // 112
	this.Label_6 = qtwidgets.NewQLabel(this.Widget_2, 0)                                                       // 111
	this.Label_6.SetObjectName("Label_6")                                                                      // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_6.SizePolicy().HasHeightForWidth())                          // 114
	this.Label_6.SetSizePolicy(this.SizePolicy1)                                                               // 114
	this.Label_6.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_4.Layout().AddWidget(this.Label_6) // 115

	this.Label_7 = qtwidgets.NewQLabel(this.Widget_2, 0)                                                       // 111
	this.Label_7.SetObjectName("Label_7")                                                                      // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_7.SizePolicy().HasHeightForWidth())                          // 114
	this.Label_7.SetSizePolicy(this.SizePolicy1)                                                               // 114
	this.Label_7.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout_4.Layout().AddWidget(this.Label_7) // 115

	this.VerticalLayout_4.AddLayout(this.HorizontalLayout_4, 0) // 115

	this.HorizontalLayout_8.AddLayout(this.VerticalLayout_4, 0) // 115

	this.VerticalLayout_7 = qtwidgets.NewQVBoxLayout()           // 111
	this.VerticalLayout_7.SetSpacing(0)                          // 114
	this.VerticalLayout_7.SetObjectName("VerticalLayout_7")      // 112
	this.ToolButton_15 = qtwidgets.NewQToolButton(this.Widget_2) // 111
	this.ToolButton_15.SetObjectName("ToolButton_15")            // 112
	this.ToolButton_15.SetMaximumSize_1(16, 16)                  // 113
	this.ToolButton_15.SetFocusPolicy(qtcore.Qt__NoFocus)        // 114
	this.Icon10 = qtgui.NewQIcon()
	this.Icon10.AddFile(":/icons/phone_mic_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_15.SetIcon(this.Icon10)                                                                       // 114
	this.ToolButton_15.SetIconSize(qtcore.NewQSize_1(12, 16))                                                     // 113
	this.ToolButton_15.SetAutoRaise(true)                                                                         // 114

	this.VerticalLayout_7.Layout().AddWidget(this.ToolButton_15) // 115

	this.ToolButton_16 = qtwidgets.NewQToolButton(this.Widget_2) // 111
	this.ToolButton_16.SetObjectName("ToolButton_16")            // 112
	this.ToolButton_16.SetMaximumSize_1(16, 16)                  // 113
	this.ToolButton_16.SetFocusPolicy(qtcore.Qt__NoFocus)        // 114
	this.Icon11 = qtgui.NewQIcon()
	this.Icon11.AddFile(":/icons/speaker_volume_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_16.SetIcon(this.Icon11)                                                                            // 114
	this.ToolButton_16.SetIconSize(qtcore.NewQSize_1(12, 12))                                                          // 113
	this.ToolButton_16.SetAutoRaise(true)                                                                              // 114

	this.VerticalLayout_7.Layout().AddWidget(this.ToolButton_16) // 115

	this.HorizontalLayout_8.AddLayout(this.VerticalLayout_7, 0) // 115

	this.ToolButton_13 = qtwidgets.NewQToolButton(this.Widget_2) // 111
	this.ToolButton_13.SetObjectName("ToolButton_13")            // 112
	this.ToolButton_13.SetMinimumSize_1(0, 0)                    // 113
	this.ToolButton_13.SetFocusPolicy(qtcore.Qt__NoFocus)        // 114
	this.Icon12 = qtgui.NewQIcon()
	this.Icon12.AddFile(":/icons/phone_call_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_13.SetIcon(this.Icon12)                                                                        // 114
	this.ToolButton_13.SetIconSize(qtcore.NewQSize_1(32, 32))                                                      // 113
	this.ToolButton_13.SetAutoRaise(true)                                                                          // 114

	this.HorizontalLayout_8.Layout().AddWidget(this.ToolButton_13) // 115

	this.ToolButton_14 = qtwidgets.NewQToolButton(this.Widget_2) // 111
	this.ToolButton_14.SetObjectName("ToolButton_14")            // 112
	this.ToolButton_14.SetMinimumSize_1(0, 0)                    // 113
	this.ToolButton_14.SetFocusPolicy(qtcore.Qt__NoFocus)        // 114
	this.Icon13 = qtgui.NewQIcon()
	this.Icon13.AddFile(":/icons/video_recorder_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_14.SetIcon(this.Icon13)                                                                            // 114
	this.ToolButton_14.SetIconSize(qtcore.NewQSize_1(32, 32))                                                          // 113
	this.ToolButton_14.SetAutoRaise(true)                                                                              // 114

	this.HorizontalLayout_8.Layout().AddWidget(this.ToolButton_14) // 115

	this.ToolButton_22 = qtwidgets.NewQToolButton(this.Widget_2) // 111
	this.ToolButton_22.SetObjectName("ToolButton_22")            // 112
	this.Icon14 = qtgui.NewQIcon()
	this.Icon14.AddFile(":/icons/vertical-ellipsis_gray32.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_22.SetIcon(this.Icon14)                                                                               // 114
	this.ToolButton_22.SetIconSize(qtcore.NewQSize_1(32, 32))                                                             // 113
	this.ToolButton_22.SetPopupMode(qtwidgets.QToolButton__InstantPopup)                                                  // 114
	this.ToolButton_22.SetAutoRaise(true)                                                                                 // 114

	this.HorizontalLayout_8.Layout().AddWidget(this.ToolButton_22) // 115

	this.VerticalLayout_14.Layout().AddWidget(this.Widget_2) // 115

	this.HorizontalLayout_9 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_9.SetObjectName("HorizontalLayout_9") // 112
	this.LineHeadSepLeft = qtwidgets.NewQFrame(this.Page_2, 0)  // 111
	this.LineHeadSepLeft.SetObjectName("LineHeadSepLeft")       // 112
	this.SizePolicy3 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy3.SetHorizontalStretch(0)                                                  // 114
	this.SizePolicy3.SetVerticalStretch(0)                                                    // 114
	this.SizePolicy3.SetHeightForWidth(this.LineHeadSepLeft.SizePolicy().HasHeightForWidth()) // 114
	this.LineHeadSepLeft.SetSizePolicy(this.SizePolicy3)                                      // 114
	this.LineHeadSepLeft.SetFrameShape(qtwidgets.QFrame__HLine)                               // 114
	this.LineHeadSepLeft.SetFrameShadow(qtwidgets.QFrame__Sunken)                             // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.LineHeadSepLeft) // 115

	this.ToolButton_24 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_24.SetObjectName("ToolButton_24")          // 112
	this.ToolButton_24.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_24) // 115

	this.ToolButton_25 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_25.SetObjectName("ToolButton_25")          // 112
	this.ToolButton_25.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_25) // 115

	this.ToolButton_26 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_26.SetObjectName("ToolButton_26")          // 112
	this.ToolButton_26.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_26) // 115

	this.ToolButton_27 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_27.SetObjectName("ToolButton_27")          // 112
	this.ToolButton_27.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_27) // 115

	this.ToolButton_28 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_28.SetObjectName("ToolButton_28")          // 112
	this.ToolButton_28.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_28) // 115

	this.ToolButton_29 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_29.SetObjectName("ToolButton_29")          // 112
	this.ToolButton_29.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_29) // 115

	this.LabelMsgCount2 = qtwidgets.NewQLabel(this.Page_2, 0) // 111
	this.LabelMsgCount2.SetObjectName("LabelMsgCount2")       // 112
	this.SizePolicy4 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy4.SetHorizontalStretch(0)                                                 // 114
	this.SizePolicy4.SetVerticalStretch(0)                                                   // 114
	this.SizePolicy4.SetHeightForWidth(this.LabelMsgCount2.SizePolicy().HasHeightForWidth()) // 114
	this.LabelMsgCount2.SetSizePolicy(this.SizePolicy4)                                      // 114
	this.LabelMsgCount2.SetAlignment(qtcore.Qt__AlignCenter)                                 // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.LabelMsgCount2) // 115

	this.ToolButton_23 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_23.SetObjectName("ToolButton_23")          // 112
	this.ToolButton_23.SetAutoRaise(true)                      // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.ToolButton_23) // 115

	this.LineHeadSepRight = qtwidgets.NewQFrame(this.Page_2, 0) // 111
	this.LineHeadSepRight.SetObjectName("LineHeadSepRight")     // 112
	this.SizePolicy5 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Fixed, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy5.SetHorizontalStretch(0)                                                   // 114
	this.SizePolicy5.SetVerticalStretch(0)                                                     // 114
	this.SizePolicy5.SetHeightForWidth(this.LineHeadSepRight.SizePolicy().HasHeightForWidth()) // 114
	this.LineHeadSepRight.SetSizePolicy(this.SizePolicy5)                                      // 114
	this.LineHeadSepRight.SetMinimumSize_1(30, 0)                                              // 113
	this.LineHeadSepRight.SetFrameShape(qtwidgets.QFrame__HLine)                               // 114
	this.LineHeadSepRight.SetFrameShadow(qtwidgets.QFrame__Sunken)                             // 114

	this.HorizontalLayout_9.Layout().AddWidget(this.LineHeadSepRight) // 115

	this.VerticalLayout_14.AddLayout(this.HorizontalLayout_9, 0) // 115

	this.ScrollArea_2 = qtwidgets.NewQScrollArea(this.Page_2)                                               // 111
	this.ScrollArea_2.SetObjectName("ScrollArea_2")                                                         // 112
	this.SizePolicy2.SetHeightForWidth(this.ScrollArea_2.SizePolicy().HasHeightForWidth())                  // 114
	this.ScrollArea_2.SetSizePolicy(this.SizePolicy2)                                                       // 114
	this.ScrollArea_2.SetFocusPolicy(qtcore.Qt__NoFocus)                                                    // 114
	this.ScrollArea_2.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)                           // 114
	this.ScrollArea_2.SetWidgetResizable(true)                                                              // 114
	this.ScrollArea_2.SetAlignment(qtcore.Qt__AlignBottom | qtcore.Qt__AlignLeading | qtcore.Qt__AlignLeft) // 114
	this.ScrollAreaWidgetContents_2 = qtwidgets.NewQWidget(nil, 0)                                          // 111
	this.ScrollAreaWidgetContents_2.SetObjectName("ScrollAreaWidgetContents_2")                             // 112
	this.ScrollAreaWidgetContents_2.SetGeometry(0, 0, 80, 16)                                               // 114
	this.SizePolicy.SetHeightForWidth(this.ScrollAreaWidgetContents_2.SizePolicy().HasHeightForWidth())     // 114
	this.ScrollAreaWidgetContents_2.SetSizePolicy(this.SizePolicy)                                          // 114
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents_2)                     // 111
	this.VerticalLayout_3.SetSpacing(0)                                                                     // 114
	this.VerticalLayout_3.SetObjectName("VerticalLayout_3")                                                 // 112
	this.VerticalLayout_3.SetContentsMargins(0, 0, 0, 0)                                                    // 114
	this.ScrollArea_2.SetWidget(this.ScrollAreaWidgetContents_2)                                            // 114

	this.VerticalLayout_14.Layout().AddWidget(this.ScrollArea_2) // 115

	this.HorizontalLayout_6 = qtwidgets.NewQHBoxLayout()        // 111
	this.HorizontalLayout_6.SetSpacing(0)                       // 114
	this.HorizontalLayout_6.SetObjectName("HorizontalLayout_6") // 112
	this.ToolButton_8 = qtwidgets.NewQToolButton(this.Page_2)   // 111
	this.ToolButton_8.SetObjectName("ToolButton_8")             // 112
	this.ToolButton_8.SetFocusPolicy(qtcore.Qt__NoFocus)        // 114
	this.Icon15 = qtgui.NewQIcon()
	this.Icon15.AddFile(":/icons/paper-clip-outline_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_8.SetIcon(this.Icon15)                                                                                 // 114
	this.ToolButton_8.SetIconSize(qtcore.NewQSize_1(22, 22))                                                               // 113
	this.ToolButton_8.SetAutoRaise(true)                                                                                   // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_8) // 115

	this.ToolButton_9 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_9.SetObjectName("ToolButton_9")           // 112
	this.ToolButton_9.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon16 = qtgui.NewQIcon()
	this.Icon16.AddFile(":/icons/snapshot@2x.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_9.SetIcon(this.Icon16)                                                                   // 114
	this.ToolButton_9.SetIconSize(qtcore.NewQSize_1(22, 22))                                                 // 113
	this.ToolButton_9.SetAutoRaise(true)                                                                     // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_9) // 115

	this.ToolButton_10 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_10.SetObjectName("ToolButton_10")          // 112
	this.ToolButton_10.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon17 = qtgui.NewQIcon()
	this.Icon17.AddFile(":/icons/smile_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_10.SetIcon(this.Icon17)                                                                   // 114
	this.ToolButton_10.SetIconSize(qtcore.NewQSize_1(22, 22))                                                 // 113
	this.ToolButton_10.SetAutoRaise(true)                                                                     // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_10) // 115

	this.LineEdit_2 = qtwidgets.NewQLineEdit(this.Page_2) // 111
	this.LineEdit_2.SetObjectName("LineEdit_2")           // 112

	this.HorizontalLayout_6.Layout().AddWidget(this.LineEdit_2) // 115

	this.ToolButton_18 = qtwidgets.NewQToolButton(this.Page_2) // 111
	this.ToolButton_18.SetObjectName("ToolButton_18")          // 112
	this.ToolButton_18.SetFocusPolicy(qtcore.Qt__NoFocus)      // 114
	this.Icon18 = qtgui.NewQIcon()
	this.Icon18.AddFile(":/icons/cursor_gray64.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_18.SetIcon(this.Icon18)                                                                    // 114
	this.ToolButton_18.SetIconSize(qtcore.NewQSize_1(22, 22))                                                  // 113
	this.ToolButton_18.SetToolButtonStyle(qtcore.Qt__ToolButtonTextBesideIcon)                                 // 114
	this.ToolButton_18.SetAutoRaise(true)                                                                      // 114

	this.HorizontalLayout_6.Layout().AddWidget(this.ToolButton_18) // 115

	this.VerticalLayout_14.AddLayout(this.HorizontalLayout_6, 0) // 115

	this.StackedWidget.AddWidget(this.Page_2)                         // 115
	this.Page_10 = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_10.SetObjectName("Page_10")                             // 112
	this.VerticalLayout_22 = qtwidgets.NewQVBoxLayout_1(this.Page_10) // 111
	this.VerticalLayout_22.SetObjectName("VerticalLayout_22")         // 112
	this.HorizontalLayout_21 = qtwidgets.NewQHBoxLayout()             // 111
	this.HorizontalLayout_21.SetObjectName("HorizontalLayout_21")     // 112
	this.HorizontalSpacer_2 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_2)

	this.HorizontalLayout_21.AddItem(this.HorizontalSpacer_2) // 115

	this.Label_27 = qtwidgets.NewQLabel(this.Page_10, 0)                     // 111
	this.Label_27.SetObjectName("Label_27")                                  // 112
	this.Label_27.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.HorizontalLayout_21.Layout().AddWidget(this.Label_27) // 115

	this.Label_28 = qtwidgets.NewQLabel(this.Page_10, 0)                     // 111
	this.Label_28.SetObjectName("Label_28")                                  // 112
	this.Label_28.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.HorizontalLayout_21.Layout().AddWidget(this.Label_28) // 115

	this.Label_29 = qtwidgets.NewQLabel(this.Page_10, 0)                     // 111
	this.Label_29.SetObjectName("Label_29")                                  // 112
	this.Label_29.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.HorizontalLayout_21.Layout().AddWidget(this.Label_29) // 115

	this.HorizontalSpacer_7 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_7)

	this.HorizontalLayout_21.AddItem(this.HorizontalSpacer_7) // 115

	this.VerticalLayout_22.AddLayout(this.HorizontalLayout_21, 0) // 115

	this.Widget_3 = qtwidgets.NewQWidget(this.Page_10, 0)                              // 111
	this.Widget_3.SetObjectName("Widget_3")                                            // 112
	this.SizePolicy2.SetHeightForWidth(this.Widget_3.SizePolicy().HasHeightForWidth()) // 114
	this.Widget_3.SetSizePolicy(this.SizePolicy2)                                      // 114

	this.VerticalLayout_22.Layout().AddWidget(this.Widget_3) // 115

	this.HorizontalLayout_22 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_22.SetObjectName("HorizontalLayout_22") // 112
	this.HorizontalSpacer_8 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_8)

	this.HorizontalLayout_22.AddItem(this.HorizontalSpacer_8) // 115

	this.ToolButton_20 = qtwidgets.NewQToolButton(this.Page_10) // 111
	this.ToolButton_20.SetObjectName("ToolButton_20")           // 112
	this.ToolButton_20.SetCheckable(true)                       // 114
	this.ToolButton_20.SetAutoRaise(true)                       // 114

	this.HorizontalLayout_22.Layout().AddWidget(this.ToolButton_20) // 115

	this.HorizontalSpacer_10 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_10)

	this.HorizontalLayout_22.AddItem(this.HorizontalSpacer_10) // 115

	this.ToolButton_21 = qtwidgets.NewQToolButton(this.Page_10) // 111
	this.ToolButton_21.SetObjectName("ToolButton_21")           // 112
	this.ToolButton_21.SetAutoRaise(true)                       // 114

	this.HorizontalLayout_22.Layout().AddWidget(this.ToolButton_21) // 115

	this.HorizontalSpacer_11 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_11)

	this.HorizontalLayout_22.AddItem(this.HorizontalSpacer_11) // 115

	this.ToolButton_30 = qtwidgets.NewQToolButton(this.Page_10) // 111
	this.ToolButton_30.SetObjectName("ToolButton_30")           // 112
	this.ToolButton_30.SetCheckable(true)                       // 114
	this.ToolButton_30.SetAutoRaise(true)                       // 114

	this.HorizontalLayout_22.Layout().AddWidget(this.ToolButton_30) // 115

	this.HorizontalSpacer_9 = qtwidgets.NewQSpacerItem(40, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)
	qtrt.ReleaseOwnerToQt(this.HorizontalSpacer_9)

	this.HorizontalLayout_22.AddItem(this.HorizontalSpacer_9) // 115

	this.VerticalLayout_22.AddLayout(this.HorizontalLayout_22, 0) // 115

	this.StackedWidget.AddWidget(this.Page_10)                               // 115
	this.Page_add_group = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_add_group.SetObjectName("Page_add_group")                      // 112
	this.VerticalLayout_21 = qtwidgets.NewQVBoxLayout_1(this.Page_add_group) // 111
	this.VerticalLayout_21.SetObjectName("VerticalLayout_21")                // 112
	this.HorizontalLayout_13 = qtwidgets.NewQHBoxLayout()                    // 111
	this.HorizontalLayout_13.SetObjectName("HorizontalLayout_13")            // 112
	this.PushButton_5 = qtwidgets.NewQPushButton(this.Page_add_group)        // 111
	this.PushButton_5.SetObjectName("PushButton_5")                          // 112
	this.PushButton_5.SetFlat(true)                                          // 114

	this.HorizontalLayout_13.Layout().AddWidget(this.PushButton_5) // 115

	this.Label_19 = qtwidgets.NewQLabel(this.Page_add_group, 0) // 111
	this.Label_19.SetObjectName("Label_19")                     // 112
	this.Label_19.SetAlignment(qtcore.Qt__AlignCenter)          // 114

	this.HorizontalLayout_13.Layout().AddWidget(this.Label_19) // 115

	this.PushButton_6 = qtwidgets.NewQPushButton(this.Page_add_group) // 111
	this.PushButton_6.SetObjectName("PushButton_6")                   // 112
	this.PushButton_6.SetFlat(true)                                   // 114

	this.HorizontalLayout_13.Layout().AddWidget(this.PushButton_6) // 115

	this.VerticalLayout_21.AddLayout(this.HorizontalLayout_13, 0) // 115

	this.Label_20 = qtwidgets.NewQLabel(this.Page_add_group, 0) // 111
	this.Label_20.SetObjectName("Label_20")                     // 112

	this.VerticalLayout_21.Layout().AddWidget(this.Label_20) // 115

	this.ComboBox_5 = qtwidgets.NewQComboBox(this.Page_add_group) // 111
	this.ComboBox_5.SetObjectName("ComboBox_5")                   // 112
	this.ComboBox_5.SetEditable(true)                             // 114

	this.VerticalLayout_21.Layout().AddWidget(this.ComboBox_5) // 115

	this.RadioButton = qtwidgets.NewQRadioButton(this.Page_add_group) // 111
	this.RadioButton.SetObjectName("RadioButton")                     // 112
	this.RadioButton.SetChecked(true)                                 // 114

	this.VerticalLayout_21.Layout().AddWidget(this.RadioButton) // 115

	this.RadioButton_2 = qtwidgets.NewQRadioButton(this.Page_add_group) // 111
	this.RadioButton_2.SetObjectName("RadioButton_2")                   // 112

	this.VerticalLayout_21.Layout().AddWidget(this.RadioButton_2) // 115

	this.HorizontalLayout_14 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_14.SetObjectName("HorizontalLayout_14") // 112
	this.Label_21 = qtwidgets.NewQLabel(this.Page_add_group, 0)   // 111
	this.Label_21.SetObjectName("Label_21")                       // 112

	this.HorizontalLayout_14.Layout().AddWidget(this.Label_21) // 115

	this.ComboBox_4 = qtwidgets.NewQComboBox(this.Page_add_group) // 111
	this.ComboBox_4.AddItem("", qtcore.NewQVariant_12("wtf"))     // 115
	this.ComboBox_4.AddItem("", qtcore.NewQVariant_12("wtf"))     // 115
	this.ComboBox_4.SetObjectName("ComboBox_4")                   // 112
	this.ComboBox_4.SetFrame(false)                               // 114

	this.HorizontalLayout_14.Layout().AddWidget(this.ComboBox_4) // 115

	this.VerticalLayout_21.AddLayout(this.HorizontalLayout_14, 0) // 115

	this.VerticalSpacer_6 = qtwidgets.NewQSpacerItem(20, 369, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_6)

	this.VerticalLayout_21.AddItem(this.VerticalSpacer_6) // 115

	this.StackedWidget.AddWidget(this.Page_add_group)                         // 115
	this.Page_add_friend = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_add_friend.SetObjectName("Page_add_friend")                     // 112
	this.VerticalLayout_20 = qtwidgets.NewQVBoxLayout_1(this.Page_add_friend) // 111
	this.VerticalLayout_20.SetObjectName("VerticalLayout_20")                 // 112
	this.HorizontalLayout_12 = qtwidgets.NewQHBoxLayout()                     // 111
	this.HorizontalLayout_12.SetObjectName("HorizontalLayout_12")             // 112
	this.PushButton_3 = qtwidgets.NewQPushButton(this.Page_add_friend)        // 111
	this.PushButton_3.SetObjectName("PushButton_3")                           // 112
	this.PushButton_3.SetFlat(true)                                           // 114

	this.HorizontalLayout_12.Layout().AddWidget(this.PushButton_3) // 115

	this.Label_16 = qtwidgets.NewQLabel(this.Page_add_friend, 0) // 111
	this.Label_16.SetObjectName("Label_16")                      // 112
	this.Label_16.SetAlignment(qtcore.Qt__AlignCenter)           // 114

	this.HorizontalLayout_12.Layout().AddWidget(this.Label_16) // 115

	this.PushButton_4 = qtwidgets.NewQPushButton(this.Page_add_friend) // 111
	this.PushButton_4.SetObjectName("PushButton_4")                    // 112
	this.PushButton_4.SetFlat(true)                                    // 114

	this.HorizontalLayout_12.Layout().AddWidget(this.PushButton_4) // 115

	this.VerticalLayout_20.AddLayout(this.HorizontalLayout_12, 0) // 115

	this.VerticalLayout_19 = qtwidgets.NewQVBoxLayout()                                                         // 111
	this.VerticalLayout_19.SetObjectName("VerticalLayout_19")                                                   // 112
	this.Label_17 = qtwidgets.NewQLabel(this.Page_add_friend, 0)                                                // 111
	this.Label_17.SetObjectName("Label_17")                                                                     // 112
	this.Label_17.SetWordWrap(true)                                                                             // 114
	this.Label_17.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.VerticalLayout_19.Layout().AddWidget(this.Label_17) // 115

	this.LineEdit_4 = qtwidgets.NewQLineEdit(this.Page_add_friend) // 111
	this.LineEdit_4.SetObjectName("LineEdit_4")                    // 112

	this.VerticalLayout_19.Layout().AddWidget(this.LineEdit_4) // 115

	this.Label_18 = qtwidgets.NewQLabel(this.Page_add_friend, 0) // 111
	this.Label_18.SetObjectName("Label_18")                      // 112

	this.VerticalLayout_19.Layout().AddWidget(this.Label_18) // 115

	this.TextEdit = qtwidgets.NewQTextEdit(this.Page_add_friend) // 111
	this.TextEdit.SetObjectName("TextEdit")                      // 112
	this.TextEdit.SetAcceptRichText(false)                       // 114

	this.VerticalLayout_19.Layout().AddWidget(this.TextEdit) // 115

	this.VerticalLayout_20.AddLayout(this.VerticalLayout_19, 0) // 115

	this.VerticalSpacer_5 = qtwidgets.NewQSpacerItem(20, 209, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_5)

	this.VerticalLayout_20.AddItem(this.VerticalSpacer_5) // 115

	this.StackedWidget.AddWidget(this.Page_add_friend)                           // 115
	this.Page_invite_friend = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_invite_friend.SetObjectName("Page_invite_friend")                  // 112
	this.VerticalLayout_18 = qtwidgets.NewQVBoxLayout_1(this.Page_invite_friend) // 111
	this.VerticalLayout_18.SetObjectName("VerticalLayout_18")                    // 112
	this.HorizontalLayout_10 = qtwidgets.NewQHBoxLayout()                        // 111
	this.HorizontalLayout_10.SetObjectName("HorizontalLayout_10")                // 112
	this.PushButton = qtwidgets.NewQPushButton(this.Page_invite_friend)          // 111
	this.PushButton.SetObjectName("PushButton")                                  // 112
	this.PushButton.SetFlat(true)                                                // 114

	this.HorizontalLayout_10.Layout().AddWidget(this.PushButton) // 115

	this.Label_14 = qtwidgets.NewQLabel(this.Page_invite_friend, 0) // 111
	this.Label_14.SetObjectName("Label_14")                         // 112
	this.Label_14.SetAlignment(qtcore.Qt__AlignCenter)              // 114

	this.HorizontalLayout_10.Layout().AddWidget(this.Label_14) // 115

	this.PushButton_2 = qtwidgets.NewQPushButton(this.Page_invite_friend) // 111
	this.PushButton_2.SetObjectName("PushButton_2")                       // 112
	this.PushButton_2.SetFlat(true)                                       // 114

	this.HorizontalLayout_10.Layout().AddWidget(this.PushButton_2) // 115

	this.VerticalLayout_18.AddLayout(this.HorizontalLayout_10, 0) // 115

	this.HorizontalLayout_11 = qtwidgets.NewQHBoxLayout()           // 111
	this.HorizontalLayout_11.SetObjectName("HorizontalLayout_11")   // 112
	this.Label_15 = qtwidgets.NewQLabel(this.Page_invite_friend, 0) // 111
	this.Label_15.SetObjectName("Label_15")                         // 112

	this.HorizontalLayout_11.Layout().AddWidget(this.Label_15) // 115

	this.LineEdit_3 = qtwidgets.NewQLineEdit(this.Page_invite_friend) // 111
	this.LineEdit_3.SetObjectName("LineEdit_3")                       // 112

	this.HorizontalLayout_11.Layout().AddWidget(this.LineEdit_3) // 115

	this.VerticalLayout_18.AddLayout(this.HorizontalLayout_11, 0) // 115

	this.ScrollArea_3 = qtwidgets.NewQScrollArea(this.Page_invite_friend)                                // 111
	this.ScrollArea_3.SetObjectName("ScrollArea_3")                                                      // 112
	this.ScrollArea_3.SetWidgetResizable(true)                                                           // 114
	this.ScrollArea_3.SetAlignment(qtcore.Qt__AlignLeading | qtcore.Qt__AlignLeft | qtcore.Qt__AlignTop) // 114
	this.ScrollAreaWidgetContents_3 = qtwidgets.NewQWidget(nil, 0)                                       // 111
	this.ScrollAreaWidgetContents_3.SetObjectName("ScrollAreaWidgetContents_3")                          // 112
	this.ScrollAreaWidgetContents_3.SetGeometry(0, 0, 80, 18)                                            // 114
	this.SizePolicy.SetHeightForWidth(this.ScrollAreaWidgetContents_3.SizePolicy().HasHeightForWidth())  // 114
	this.ScrollAreaWidgetContents_3.SetSizePolicy(this.SizePolicy)                                       // 114
	this.VerticalLayout_17 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents_3)                 // 111
	this.VerticalLayout_17.SetObjectName("VerticalLayout_17")                                            // 112
	this.ScrollArea_3.SetWidget(this.ScrollAreaWidgetContents_3)                                         // 114

	this.VerticalLayout_18.Layout().AddWidget(this.ScrollArea_3) // 115

	this.TableWidget = qtwidgets.NewQTableWidget(this.Page_invite_friend)          // 111
	this.TableWidget.SetObjectName("TableWidget")                                  // 112
	this.TableWidget.SetEditTriggers(qtwidgets.QAbstractItemView__NoEditTriggers)  // 114
	this.TableWidget.SetProperty("showDropIndicator", qtcore.NewQVariant_9(false)) // 114
	this.TableWidget.SetDragDropOverwriteMode(false)                               // 114
	this.TableWidget.SetAlternatingRowColors(true)                                 // 114

	this.VerticalLayout_18.Layout().AddWidget(this.TableWidget) // 115

	this.StackedWidget.AddWidget(this.Page_invite_friend)            // 115
	this.Page_9 = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_9.SetObjectName("Page_9")                              // 112
	this.VerticalLayout_10 = qtwidgets.NewQVBoxLayout_1(this.Page_9) // 111
	this.VerticalLayout_10.SetObjectName("VerticalLayout_10")        // 112
	this.HorizontalLayout_19 = qtwidgets.NewQHBoxLayout()            // 111
	this.HorizontalLayout_19.SetObjectName("HorizontalLayout_19")    // 112
	this.PushButton_8 = qtwidgets.NewQPushButton(this.Page_9)        // 111
	this.PushButton_8.SetObjectName("PushButton_8")                  // 112
	this.PushButton_8.SetFlat(true)                                  // 114

	this.HorizontalLayout_19.Layout().AddWidget(this.PushButton_8) // 115

	this.Label_26 = qtwidgets.NewQLabel(this.Page_9, 0) // 111
	this.Label_26.SetObjectName("Label_26")             // 112
	this.Label_26.SetAlignment(qtcore.Qt__AlignCenter)  // 114

	this.HorizontalLayout_19.Layout().AddWidget(this.Label_26) // 115

	this.PushButton_9 = qtwidgets.NewQPushButton(this.Page_9) // 111
	this.PushButton_9.SetObjectName("PushButton_9")           // 112
	this.PushButton_9.SetFlat(true)                           // 114

	this.HorizontalLayout_19.Layout().AddWidget(this.PushButton_9) // 115

	this.VerticalLayout_10.AddLayout(this.HorizontalLayout_19, 0) // 115

	this.HorizontalLayout_18 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_18.SetObjectName("HorizontalLayout_18") // 112
	this.Label_25 = qtwidgets.NewQLabel(this.Page_9, 0)           // 111
	this.Label_25.SetObjectName("Label_25")                       // 112

	this.HorizontalLayout_18.Layout().AddWidget(this.Label_25) // 115

	this.LineEdit_7 = qtwidgets.NewQLineEdit(this.Page_9) // 111
	this.LineEdit_7.SetObjectName("LineEdit_7")           // 112

	this.HorizontalLayout_18.Layout().AddWidget(this.LineEdit_7) // 115

	this.VerticalLayout_10.AddLayout(this.HorizontalLayout_18, 0) // 115

	this.TableWidget_2 = qtwidgets.NewQTableWidget(this.Page_9) // 111
	this.TableWidget_2.SetObjectName("TableWidget_2")           // 112
	this.TableWidget_2.SetAlternatingRowColors(true)            // 114

	this.VerticalLayout_10.Layout().AddWidget(this.TableWidget_2) // 115

	this.StackedWidget.AddWidget(this.Page_9)                         // 115
	this.Page_11 = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_11.SetObjectName("Page_11")                             // 112
	this.VerticalLayout_26 = qtwidgets.NewQVBoxLayout_1(this.Page_11) // 111
	this.VerticalLayout_26.SetObjectName("VerticalLayout_26")         // 112
	this.HorizontalLayout_23 = qtwidgets.NewQHBoxLayout()             // 111
	this.HorizontalLayout_23.SetObjectName("HorizontalLayout_23")     // 112
	this.PushButton_10 = qtwidgets.NewQPushButton(this.Page_11)       // 111
	this.PushButton_10.SetObjectName("PushButton_10")                 // 112
	this.PushButton_10.SetFlat(true)                                  // 114

	this.HorizontalLayout_23.Layout().AddWidget(this.PushButton_10) // 115

	this.Label_30 = qtwidgets.NewQLabel(this.Page_11, 0)                     // 111
	this.Label_30.SetObjectName("Label_30")                                  // 112
	this.Label_30.SetAlignment(qtcore.Qt__AlignCenter)                       // 114
	this.Label_30.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.HorizontalLayout_23.Layout().AddWidget(this.Label_30) // 115

	this.PushButton_11 = qtwidgets.NewQPushButton(this.Page_11) // 111
	this.PushButton_11.SetObjectName("PushButton_11")           // 112
	this.PushButton_11.SetFlat(true)                            // 114

	this.HorizontalLayout_23.Layout().AddWidget(this.PushButton_11) // 115

	this.VerticalLayout_26.AddLayout(this.HorizontalLayout_23, 0) // 115

	this.HorizontalLayout_24 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_24.SetObjectName("HorizontalLayout_24") // 112
	this.ToolButton_31 = qtwidgets.NewQToolButton(this.Page_11)   // 111
	this.ToolButton_31.SetObjectName("ToolButton_31")             // 112
	this.ToolButton_31.SetMinimumSize_1(64, 64)                   // 113
	this.ToolButton_31.SetFocusPolicy(qtcore.Qt__NoFocus)         // 114
	this.ToolButton_31.SetIconSize(qtcore.NewQSize_1(64, 64))     // 113
	this.ToolButton_31.SetAutoRaise(true)                         // 114

	this.HorizontalLayout_24.Layout().AddWidget(this.ToolButton_31) // 115

	this.VerticalLayout_23 = qtwidgets.NewQVBoxLayout()                      // 111
	this.VerticalLayout_23.SetObjectName("VerticalLayout_23")                // 112
	this.Label_32 = qtwidgets.NewQLabel(this.Page_11, 0)                     // 111
	this.Label_32.SetObjectName("Label_32")                                  // 112
	this.Label_32.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.VerticalLayout_23.Layout().AddWidget(this.Label_32) // 115

	this.LineEdit_9 = qtwidgets.NewQLineEdit(this.Page_11) // 111
	this.LineEdit_9.SetObjectName("LineEdit_9")            // 112
	this.LineEdit_9.SetFocusPolicy(qtcore.Qt__NoFocus)     // 114
	this.LineEdit_9.SetReadOnly(true)                      // 114

	this.VerticalLayout_23.Layout().AddWidget(this.LineEdit_9) // 115

	this.HorizontalLayout_24.AddLayout(this.VerticalLayout_23, 0) // 115

	this.VerticalLayout_26.AddLayout(this.HorizontalLayout_24, 0) // 115

	this.CheckBox_4 = qtwidgets.NewQCheckBox(this.Page_11) // 111
	this.CheckBox_4.SetObjectName("CheckBox_4")            // 112

	this.VerticalLayout_26.Layout().AddWidget(this.CheckBox_4) // 115

	this.GroupBox_3 = qtwidgets.NewQGroupBox(this.Page_11)                 // 111
	this.GroupBox_3.SetObjectName("GroupBox_3")                            // 112
	this.HorizontalLayout_25 = qtwidgets.NewQHBoxLayout_1(this.GroupBox_3) // 111
	this.HorizontalLayout_25.SetObjectName("HorizontalLayout_25")          // 112
	this.RadioButton_5 = qtwidgets.NewQRadioButton(this.GroupBox_3)        // 111
	this.RadioButton_5.SetObjectName("RadioButton_5")                      // 112
	this.RadioButton_5.SetChecked(true)                                    // 114

	this.HorizontalLayout_25.Layout().AddWidget(this.RadioButton_5) // 115

	this.RadioButton_6 = qtwidgets.NewQRadioButton(this.GroupBox_3) // 111
	this.RadioButton_6.SetObjectName("RadioButton_6")               // 112

	this.HorizontalLayout_25.Layout().AddWidget(this.RadioButton_6) // 115

	this.RadioButton_7 = qtwidgets.NewQRadioButton(this.GroupBox_3) // 111
	this.RadioButton_7.SetObjectName("RadioButton_7")               // 112

	this.HorizontalLayout_25.Layout().AddWidget(this.RadioButton_7) // 115

	this.VerticalLayout_26.Layout().AddWidget(this.GroupBox_3) // 115

	this.GroupBox_4 = qtwidgets.NewQGroupBox(this.Page_11)               // 111
	this.GroupBox_4.SetObjectName("GroupBox_4")                          // 112
	this.VerticalLayout_24 = qtwidgets.NewQVBoxLayout_1(this.GroupBox_4) // 111
	this.VerticalLayout_24.SetObjectName("VerticalLayout_24")            // 112
	this.CheckBox_5 = qtwidgets.NewQCheckBox(this.GroupBox_4)            // 111
	this.CheckBox_5.SetObjectName("CheckBox_5")                          // 112

	this.VerticalLayout_24.Layout().AddWidget(this.CheckBox_5) // 115

	this.HorizontalLayout_26 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_26.SetObjectName("HorizontalLayout_26") // 112
	this.Label_34 = qtwidgets.NewQLabel(this.GroupBox_4, 0)       // 111
	this.Label_34.SetObjectName("Label_34")                       // 112

	this.HorizontalLayout_26.Layout().AddWidget(this.Label_34) // 115

	this.LineEdit_8 = qtwidgets.NewQLineEdit(this.GroupBox_4) // 111
	this.LineEdit_8.SetObjectName("LineEdit_8")               // 112

	this.HorizontalLayout_26.Layout().AddWidget(this.LineEdit_8) // 115

	this.VerticalLayout_24.AddLayout(this.HorizontalLayout_26, 0) // 115

	this.VerticalLayout_26.Layout().AddWidget(this.GroupBox_4) // 115

	this.PushButton_12 = qtwidgets.NewQPushButton(this.Page_11) // 111
	this.PushButton_12.SetObjectName("PushButton_12")           // 112

	this.VerticalLayout_26.Layout().AddWidget(this.PushButton_12) // 115

	this.GroupBox_5 = qtwidgets.NewQGroupBox(this.Page_11)               // 111
	this.GroupBox_5.SetObjectName("GroupBox_5")                          // 112
	this.VerticalLayout_25 = qtwidgets.NewQVBoxLayout_1(this.GroupBox_5) // 111
	this.VerticalLayout_25.SetObjectName("VerticalLayout_25")            // 112
	this.TextEdit_2 = qtwidgets.NewQTextEdit(this.GroupBox_5)            // 111
	this.TextEdit_2.SetObjectName("TextEdit_2")                          // 112

	this.VerticalLayout_25.Layout().AddWidget(this.TextEdit_2) // 115

	this.VerticalLayout_26.Layout().AddWidget(this.GroupBox_5) // 115

	this.StackedWidget.AddWidget(this.Page_11)                      // 115
	this.Page_6 = qtwidgets.NewQWidget(nil, 0)                      // 111
	this.Page_6.SetObjectName("Page_6")                             // 112
	this.VerticalLayout_5 = qtwidgets.NewQVBoxLayout_1(this.Page_6) // 111
	this.VerticalLayout_5.SetSpacing(0)                             // 114
	this.VerticalLayout_5.SetObjectName("VerticalLayout_5")         // 112
	this.VerticalLayout_5.SetContentsMargins(0, 0, 0, 0)            // 114
	this.ListWidget_2 = qtwidgets.NewQListWidget(this.Page_6)       // 111
	this.ListWidget_2.SetObjectName("ListWidget_2")                 // 112

	this.VerticalLayout_5.Layout().AddWidget(this.ListWidget_2) // 115

	this.ListWidget = qtwidgets.NewQListWidget(this.Page_6) // 111
	this.ListWidget.SetObjectName("ListWidget")             // 112
	this.ListWidget.SetAlternatingRowColors(false)          // 114

	this.VerticalLayout_5.Layout().AddWidget(this.ListWidget) // 115

	this.StackedWidget.AddWidget(this.Page_6)                        // 115
	this.Page_5 = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_5.SetObjectName("Page_5")                              // 112
	this.VerticalLayout_12 = qtwidgets.NewQVBoxLayout_1(this.Page_5) // 111
	this.VerticalLayout_12.SetSpacing(0)                             // 114
	this.VerticalLayout_12.SetObjectName("VerticalLayout_12")        // 112
	this.VerticalLayout_12.SetContentsMargins(0, 0, 0, 0)            // 114
	this.TextBrowser = qtwidgets.NewQTextBrowser(this.Page_5)        // 111
	this.TextBrowser.SetObjectName("TextBrowser")                    // 112

	this.VerticalLayout_12.Layout().AddWidget(this.TextBrowser) // 115

	this.StackedWidget.AddWidget(this.Page_5)                         // 115
	this.Page_12 = qtwidgets.NewQWidget(nil, 0)                       // 111
	this.Page_12.SetObjectName("Page_12")                             // 112
	this.VerticalLayout_28 = qtwidgets.NewQVBoxLayout_1(this.Page_12) // 111
	this.VerticalLayout_28.SetObjectName("VerticalLayout_28")         // 112
	this.GroupBox_6 = qtwidgets.NewQGroupBox(this.Page_12)            // 111
	this.GroupBox_6.SetObjectName("GroupBox_6")                       // 112
	this.GridLayout_2 = qtwidgets.NewQGridLayout(this.GroupBox_6)     // 111
	this.GridLayout_2.SetObjectName("GridLayout_2")                   // 112
	this.Label_31 = qtwidgets.NewQLabel(this.GroupBox_6, 0)           // 111
	this.Label_31.SetObjectName("Label_31")                           // 112

	this.GridLayout_2.AddWidget_2(this.Label_31, 0, 0, 1, 1, 0) // 115

	this.Label_37 = qtwidgets.NewQLabel(this.GroupBox_6, 0)                            // 111
	this.Label_37.SetObjectName("Label_37")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_37.SizePolicy().HasHeightForWidth()) // 114
	this.Label_37.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.GridLayout_2.AddWidget_2(this.Label_37, 0, 1, 1, 1, 0) // 115

	this.Label_33 = qtwidgets.NewQLabel(this.GroupBox_6, 0) // 111
	this.Label_33.SetObjectName("Label_33")                 // 112

	this.GridLayout_2.AddWidget_2(this.Label_33, 1, 0, 1, 1, 0) // 115

	this.Label_38 = qtwidgets.NewQLabel(this.GroupBox_6, 0)                            // 111
	this.Label_38.SetObjectName("Label_38")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_38.SizePolicy().HasHeightForWidth()) // 114
	this.Label_38.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.GridLayout_2.AddWidget_2(this.Label_38, 1, 1, 1, 1, 0) // 115

	this.Label_35 = qtwidgets.NewQLabel(this.GroupBox_6, 0) // 111
	this.Label_35.SetObjectName("Label_35")                 // 112

	this.GridLayout_2.AddWidget_2(this.Label_35, 2, 0, 1, 1, 0) // 115

	this.Label_39 = qtwidgets.NewQLabel(this.GroupBox_6, 0)                            // 111
	this.Label_39.SetObjectName("Label_39")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_39.SizePolicy().HasHeightForWidth()) // 114
	this.Label_39.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.GridLayout_2.AddWidget_2(this.Label_39, 2, 1, 1, 1, 0) // 115

	this.Label_36 = qtwidgets.NewQLabel(this.GroupBox_6, 0) // 111
	this.Label_36.SetObjectName("Label_36")                 // 112

	this.GridLayout_2.AddWidget_2(this.Label_36, 3, 0, 1, 1, 0) // 115

	this.Label_40 = qtwidgets.NewQLabel(this.GroupBox_6, 0)                            // 111
	this.Label_40.SetObjectName("Label_40")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_40.SizePolicy().HasHeightForWidth()) // 114
	this.Label_40.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.GridLayout_2.AddWidget_2(this.Label_40, 3, 1, 1, 1, 0) // 115

	this.Label_41 = qtwidgets.NewQLabel(this.GroupBox_6, 0) // 111
	this.Label_41.SetObjectName("Label_41")                 // 112

	this.GridLayout_2.AddWidget_2(this.Label_41, 4, 0, 1, 1, 0) // 115

	this.Label_42 = qtwidgets.NewQLabel(this.GroupBox_6, 0)                            // 111
	this.Label_42.SetObjectName("Label_42")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_42.SizePolicy().HasHeightForWidth()) // 114
	this.Label_42.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.GridLayout_2.AddWidget_2(this.Label_42, 4, 1, 1, 1, 0) // 115

	this.VerticalLayout_28.Layout().AddWidget(this.GroupBox_6) // 115

	this.GroupBox_7 = qtwidgets.NewQGroupBox(this.Page_12)               // 111
	this.GroupBox_7.SetObjectName("GroupBox_7")                          // 112
	this.VerticalLayout_27 = qtwidgets.NewQVBoxLayout_1(this.GroupBox_7) // 111
	this.VerticalLayout_27.SetObjectName("VerticalLayout_27")            // 112
	this.HorizontalLayout_28 = qtwidgets.NewQHBoxLayout()                // 111
	this.HorizontalLayout_28.SetObjectName("HorizontalLayout_28")        // 112
	this.Label_46 = qtwidgets.NewQLabel(this.GroupBox_7, 0)              // 111
	this.Label_46.SetObjectName("Label_46")                              // 112

	this.HorizontalLayout_28.Layout().AddWidget(this.Label_46) // 115

	this.Label_45 = qtwidgets.NewQLabel(this.GroupBox_7, 0)                            // 111
	this.Label_45.SetObjectName("Label_45")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_45.SizePolicy().HasHeightForWidth()) // 114
	this.Label_45.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.HorizontalLayout_28.Layout().AddWidget(this.Label_45) // 115

	this.VerticalLayout_27.AddLayout(this.HorizontalLayout_28, 0) // 115

	this.Label_53 = qtwidgets.NewQLabel(this.GroupBox_7, 0) // 111
	this.Label_53.SetObjectName("Label_53")                 // 112

	this.VerticalLayout_27.Layout().AddWidget(this.Label_53) // 115

	this.Label_54 = qtwidgets.NewQLabel(this.GroupBox_7, 0) // 111
	this.Label_54.SetObjectName("Label_54")                 // 112
	this.Label_54.SetWordWrap(true)                         // 114

	this.VerticalLayout_27.Layout().AddWidget(this.Label_54) // 115

	this.Label_52 = qtwidgets.NewQLabel(this.GroupBox_7, 0) // 111
	this.Label_52.SetObjectName("Label_52")                 // 112

	this.VerticalLayout_27.Layout().AddWidget(this.Label_52) // 115

	this.Label_50 = qtwidgets.NewQLabel(this.GroupBox_7, 0)                            // 111
	this.Label_50.SetObjectName("Label_50")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_50.SizePolicy().HasHeightForWidth()) // 114
	this.Label_50.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Label_50.SetWordWrap(true)                                                    // 114

	this.VerticalLayout_27.Layout().AddWidget(this.Label_50) // 115

	this.Label_44 = qtwidgets.NewQLabel(this.GroupBox_7, 0) // 111
	this.Label_44.SetObjectName("Label_44")                 // 112

	this.VerticalLayout_27.Layout().AddWidget(this.Label_44) // 115

	this.Label_47 = qtwidgets.NewQLabel(this.GroupBox_7, 0)                            // 111
	this.Label_47.SetObjectName("Label_47")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_47.SizePolicy().HasHeightForWidth()) // 114
	this.Label_47.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Label_47.SetWordWrap(true)                                                    // 114

	this.VerticalLayout_27.Layout().AddWidget(this.Label_47) // 115

	this.HorizontalLayout_27 = qtwidgets.NewQHBoxLayout()         // 111
	this.HorizontalLayout_27.SetObjectName("HorizontalLayout_27") // 112
	this.Label_43 = qtwidgets.NewQLabel(this.GroupBox_7, 0)       // 111
	this.Label_43.SetObjectName("Label_43")                       // 112

	this.HorizontalLayout_27.Layout().AddWidget(this.Label_43) // 115

	this.Label_48 = qtwidgets.NewQLabel(this.GroupBox_7, 0)                            // 111
	this.Label_48.SetObjectName("Label_48")                                            // 112
	this.SizePolicy1.SetHeightForWidth(this.Label_48.SizePolicy().HasHeightForWidth()) // 114
	this.Label_48.SetSizePolicy(this.SizePolicy1)                                      // 114

	this.HorizontalLayout_27.Layout().AddWidget(this.Label_48) // 115

	this.VerticalLayout_27.AddLayout(this.HorizontalLayout_27, 0) // 115

	this.VerticalLayout_28.Layout().AddWidget(this.GroupBox_7) // 115

	this.VerticalSpacer_4 = qtwidgets.NewQSpacerItem(20, 158, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)
	qtrt.ReleaseOwnerToQt(this.VerticalSpacer_4)

	this.VerticalLayout_28.AddItem(this.VerticalSpacer_4) // 115

	this.StackedWidget.AddWidget(this.Page_12) // 115

	this.VerticalLayout_6.Layout().AddWidget(this.StackedWidget) // 115

	this.MainWindow.SetCentralWidget(this.Centralwidget) // 114

	this.RetranslateUi(MainWindow)

	this.StackedWidget.SetCurrentIndex(3) // 114
	this.ComboBox_6.SetCurrentIndex(0)    // 114

	qtcore.QMetaObject_ConnectSlotsByName(MainWindow) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QMainWindow *MainWindow)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MainWindow) RetranslateUi(MainWindow *qtwidgets.QMainWindow) {
	// noimpl: {
	this.MainWindow.SetWindowTitle(qtcore.QCoreApplication_Translate("MainWindow", "go-toxhsui", "dummy123", 0))
	this.Actionooo.SetText(qtcore.QCoreApplication_Translate("MainWindow", "ooo", "dummy123", 0))
	this.ActionQuit.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Quit", "dummy123", 0))
	this.Action_About.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&About", "dummy123", 0))
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_33.SetStatusTip(qtcore.QCoreApplication_Translate("MainWindow", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_33.SetText(qtcore.QCoreApplication_Translate("MainWindow", "\342\227\201 Back", "dummy123", 0))
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_11.SetStatusTip(qtcore.QCoreApplication_Translate("MainWindow", "Back by fixed order.", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_11.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("MainWindow", "curwin: ", "dummy123", 0))
	this.ComboBox.SetItemText(0, qtcore.QCoreApplication_Translate("MainWindow", "1 qmlmctrl", "dummy123", 0))
	this.ComboBox.SetItemText(1, qtcore.QCoreApplication_Translate("MainWindow", "2 qmlorgin", "dummy123", 0))
	this.ComboBox.SetItemText(2, qtcore.QCoreApplication_Translate("MainWindow", "3 settings", "dummy123", 0))
	this.ComboBox.SetItemText(3, qtcore.QCoreApplication_Translate("MainWindow", "4 loginui", "dummy123", 0))
	this.ComboBox.SetItemText(4, qtcore.QCoreApplication_Translate("MainWindow", "5 contactui", "dummy123", 0))
	this.ComboBox.SetItemText(5, qtcore.QCoreApplication_Translate("MainWindow", "6 messageui", "dummy123", 0))
	this.ComboBox.SetItemText(6, qtcore.QCoreApplication_Translate("MainWindow", "7 videoui", "dummy123", 0))
	this.ComboBox.SetItemText(7, qtcore.QCoreApplication_Translate("MainWindow", "8 add group", "dummy123", 0))
	this.ComboBox.SetItemText(8, qtcore.QCoreApplication_Translate("MainWindow", "9 add friend", "dummy123", 0))
	this.ComboBox.SetItemText(9, qtcore.QCoreApplication_Translate("MainWindow", "10 invite friend", "dummy123", 0))
	this.ComboBox.SetItemText(10, qtcore.QCoreApplication_Translate("MainWindow", "11 memberui", "dummy123", 0))
	this.ComboBox.SetItemText(11, qtcore.QCoreApplication_Translate("MainWindow", "12 contactinfoui", "dummy123", 0))
	this.ComboBox.SetItemText(12, qtcore.QCoreApplication_Translate("MainWindow", "13 testui", "dummy123", 0))
	this.ComboBox.SetItemText(13, qtcore.QCoreApplication_Translate("MainWindow", "14 logui", "dummy123", 0))
	this.ComboBox.SetItemText(14, qtcore.QCoreApplication_Translate("MainWindow", "15 aboutui", "dummy123", 0))
	// noimpl:
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_19.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Tools menu", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_19.SetText(qtcore.QCoreApplication_Translate("MainWindow", " &Tool ", "dummy123", 0))
	this.ToolButton_12.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.GroupBox.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "General", "dummy123", 0))
	this.Label_13.SetText(qtcore.QCoreApplication_Translate("MainWindow", "PlaceHolder...", "dummy123", 0))
	this.Label_8.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Theme", "dummy123", 0))
	this.ComboBox_2.SetItemText(0, qtcore.QCoreApplication_Translate("MainWindow", "Light", "dummy123", 0))
	this.ComboBox_2.SetItemText(1, qtcore.QCoreApplication_Translate("MainWindow", "Dark", "dummy123", 0))
	this.ComboBox_2.SetItemText(2, qtcore.QCoreApplication_Translate("MainWindow", "System", "dummy123", 0))
	// noimpl:
	this.Label_9.SetText(qtcore.QCoreApplication_Translate("MainWindow", "ToxHS IP:", "dummy123", 0))
	this.Label_10.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Use HS", "dummy123", 0))
	// noimpl: checkBox_2->setText(QString());
	this.Label_11.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Theme", "dummy123", 0))
	// noimpl: checkBox_3->setText(QString());
	this.ComboBox_3.SetItemText(0, qtcore.QCoreApplication_Translate("MainWindow", "txhs.duckdns.org", "dummy123", 0))
	this.ComboBox_3.SetItemText(1, qtcore.QCoreApplication_Translate("MainWindow", "10.0.0.31", "dummy123", 0))
	this.ComboBox_3.SetItemText(2, qtcore.QCoreApplication_Translate("MainWindow", "127.0.0.1", "dummy123", 0))
	// noimpl:
	// noimpl: checkBox->setText(QString());
	this.GroupBox_2.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "PlaceHolder...", "dummy123", 0))
	this.Label_12.SetText(qtcore.QCoreApplication_Translate("MainWindow", "PlaceHolder...", "dummy123", 0))
	this.Label_22.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Home server URL", "dummy123", 0))
	this.Label_23.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ComboBox_6.SetItemText(0, qtcore.QCoreApplication_Translate("MainWindow", "txhs.duckdns.org:2080", "dummy123", 0))
	// noimpl:
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.RadioButton_3.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "For this one, need a seperate running server. Recommanded.", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.RadioButton_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Remote Server", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.RadioButton_4.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "For this one, not need a seperate server, this program has embeded one, and auto start it. Preview purpose.", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.RadioButton_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Self &Contains", "dummy123", 0))
	this.PushButton_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Sign in", "dummy123", 0))
	this.Label_24.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_17.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Loading ...", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Loading ...", "dummy123", 0))
	this.ToolButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_32.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "logout", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_32.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.LineEdit.SetPlaceholderText(qtcore.QCoreApplication_Translate("MainWindow", "Filter...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_2.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "filter and order", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_3.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "remove only me groups", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: label_4->setText(QString());
	this.Label_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.LabelMsgCount.SetText(qtcore.QCoreApplication_Translate("MainWindow", "0", "dummy123", 0))
	this.Label_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "0 people", "dummy123", 0))
	this.Label_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.ToolButton_15.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_16.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_13.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_14.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.ToolButton_22.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_24.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Scroll content Fixed", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_24.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_25.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Scroll content Max", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_25.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_26.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Scroll content Prefer", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_26.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_27.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Scroll content Min", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_27.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_28.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Scroll content Expand", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_28.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_29.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Scroll content Min Expand", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_29.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.LabelMsgCount2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "000", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_23.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Load older messages", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_23.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_8.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Send File", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_8.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_9.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Send snapshot", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_9.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_10.SetToolTip(qtcore.QCoreApplication_Translate("MainWindow", "Emoji", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	this.ToolButton_10.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.LineEdit_2.SetPlaceholderText(qtcore.QCoreApplication_Translate("MainWindow", "Write a message...", "dummy123", 0))
	this.ToolButton_18.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Send", "dummy123", 0))
	this.Label_27.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Video call with", "dummy123", 0))
	this.Label_28.SetText(qtcore.QCoreApplication_Translate("MainWindow", "fname", "dummy123", 0))
	this.Label_29.SetText(qtcore.QCoreApplication_Translate("MainWindow", "time", "dummy123", 0))
	this.ToolButton_20.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Mic", "dummy123", 0))
	this.ToolButton_21.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Hangup", "dummy123", 0))
	this.ToolButton_30.SetText(qtcore.QCoreApplication_Translate("MainWindow", "M&ute", "dummy123", 0))
	this.PushButton_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Cancel", "dummy123", 0))
	this.Label_19.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Add Group", "dummy123", 0))
	this.PushButton_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Done", "dummy123", 0))
	this.Label_20.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Room Name:", "dummy123", 0))
	this.RadioButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Text", "dummy123", 0))
	this.RadioButton_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Audio", "dummy123", 0))
	this.Label_21.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Room Type:", "dummy123", 0))
	this.ComboBox_4.SetItemText(0, qtcore.QCoreApplication_Translate("MainWindow", "Text", "dummy123", 0))
	this.ComboBox_4.SetItemText(1, qtcore.QCoreApplication_Translate("MainWindow", "Audio", "dummy123", 0))
	// noimpl:
	this.PushButton_3.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Cancel", "dummy123", 0))
	this.Label_16.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Add Friend", "dummy123", 0))
	this.PushButton_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Done", "dummy123", 0))
	this.Label_17.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Tox ID (either 76 hexadecimal characters or name@example.com )", "dummy123", 0))
	this.Label_18.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Message", "dummy123", 0))
	this.TextEdit.SetPlaceholderText(qtcore.QCoreApplication_Translate("MainWindow", "I am %1 ! Let's chat with Tox?", "dummy123", 0))
	this.PushButton.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Cancel", "dummy123", 0))
	this.Label_14.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Inivte Friend", "dummy123", 0))
	this.PushButton_2.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Done", "dummy123", 0))
	this.Label_15.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Search", "dummy123", 0))
	this.PushButton_8.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Cancel", "dummy123", 0))
	this.Label_26.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Group Members", "dummy123", 0))
	this.PushButton_9.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Done", "dummy123", 0))
	this.Label_25.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Search", "dummy123", 0))
	this.PushButton_10.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Cancel", "dummy123", 0))
	this.Label_30.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Contact Info", "dummy123", 0))
	this.PushButton_11.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Done", "dummy123", 0))
	this.ToolButton_31.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Contact icon", "dummy123", 0))
	this.Label_32.SetText(qtcore.QCoreApplication_Translate("MainWindow", "TextLabel", "dummy123", 0))
	this.CheckBox_4.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Auto Accept Group Invite", "dummy123", 0))
	this.GroupBox_3.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "Auto Pickup Call:", "dummy123", 0))
	this.RadioButton_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Manually", "dummy123", 0))
	this.RadioButton_6.SetText(qtcore.QCoreApplication_Translate("MainWindow", "&Audio", "dummy123", 0))
	this.RadioButton_7.SetText(qtcore.QCoreApplication_Translate("MainWindow", "A&udio+Video", "dummy123", 0))
	this.GroupBox_4.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "Auto Accept Files:", "dummy123", 0))
	this.CheckBox_5.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Auto Accept File And Save", "dummy123", 0))
	this.Label_34.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Path:", "dummy123", 0))
	this.PushButton_12.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Delete History Chat Messages!", "dummy123", 0))
	this.GroupBox_5.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "Note:", "dummy123", 0))
	this.GroupBox_6.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "Version:", "dummy123", 0))
	this.Label_31.SetText(qtcore.QCoreApplication_Translate("MainWindow", "golem version:", "dummy123", 0))
	this.Label_37.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_33.SetText(qtcore.QCoreApplication_Translate("MainWindow", "golem hash:", "dummy123", 0))
	this.Label_38.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_35.SetText(qtcore.QCoreApplication_Translate("MainWindow", "toxcore version:", "dummy123", 0))
	this.Label_39.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_36.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Qt Runtime Version:", "dummy123", 0))
	this.Label_40.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.Label_41.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Qt Compile Version:", "dummy123", 0))
	this.Label_42.SetText(qtcore.QCoreApplication_Translate("MainWindow", "...", "dummy123", 0))
	this.GroupBox_7.SetTitle(qtcore.QCoreApplication_Translate("MainWindow", "Infomation:", "dummy123", 0))
	this.Label_46.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Author:", "dummy123", 0))
	this.Label_45.SetText(qtcore.QCoreApplication_Translate("MainWindow", "envoy", "dummy123", 0))
	this.Label_53.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Contributors:", "dummy123", 0))
	this.Label_54.SetText(qtcore.QCoreApplication_Translate("MainWindow", "<html><head/><body><p><a href=\"https://github.com/envsh/tox-homeserver/graphs/contributors\"><span style=\" text-decoration: underline; color:#0000ff;\">https://github.com/envsh/tox-homeserver/graphs/contributors</span></a></p></body></html>", "dummy123", 0))
	this.Label_52.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Home URL:", "dummy123", 0))
	this.Label_50.SetText(qtcore.QCoreApplication_Translate("MainWindow", "<html><head/><body><p><a href=\"https://github.com/envsh/tox-homeserver\"><span style=\" text-decoration: underline; color:#0000ff;\">https://github.com/envsh/tox-homeserver</span></a></p></body></html>", "dummy123", 0))
	this.Label_44.SetText(qtcore.QCoreApplication_Translate("MainWindow", "Feedback:", "dummy123", 0))
	this.Label_47.SetText(qtcore.QCoreApplication_Translate("MainWindow", "<html><head/><body><p><a href=\"https://github.com/envsh/tox-homeserver/issues\"><span style=\" text-decoration: underline; color:#0000ff;\">https://github.com/envsh/tox-homeserver/issues</span></a></p></body></html>", "dummy123", 0))
	this.Label_43.SetText(qtcore.QCoreApplication_Translate("MainWindow", "License:", "dummy123", 0))
	this.Label_48.SetText(qtcore.QCoreApplication_Translate("MainWindow", "GPL3", "dummy123", 0))
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
