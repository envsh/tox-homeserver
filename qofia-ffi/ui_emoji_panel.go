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
func NewUi_EmojiPanel() *Ui_EmojiPanel {
	return &Ui_EmojiPanel{}
}

type Ui_EmojiPanel struct {
	VerticalLayout           *qtwidgets.QVBoxLayout
	MainWidget               *qtwidgets.QWidget
	VerticalLayout_2         *qtwidgets.QVBoxLayout
	ScrollArea               *qtwidgets.QScrollArea
	ScrollAreaWidgetContents *qtwidgets.QWidget
	VerticalLayout_3         *qtwidgets.QVBoxLayout
	GridLayout               *qtwidgets.QGridLayout
	ToolButton_33            *qtwidgets.QToolButton
	ToolButton_34            *qtwidgets.QToolButton
	ToolButton_35            *qtwidgets.QToolButton
	ToolButton_36            *qtwidgets.QToolButton
	ToolButton_37            *qtwidgets.QToolButton
	ToolButton_38            *qtwidgets.QToolButton
	ToolButton_39            *qtwidgets.QToolButton
	ToolButton_40            *qtwidgets.QToolButton
	EmojiPanel               *qtwidgets.QWidget
	Icon                     *qtgui.QIcon // 116
	Icon1                    *qtgui.QIcon // 116
	Icon2                    *qtgui.QIcon // 116
	Icon3                    *qtgui.QIcon // 116
	Icon4                    *qtgui.QIcon // 116
	Icon5                    *qtgui.QIcon // 116
	Icon6                    *qtgui.QIcon // 116
	Icon7                    *qtgui.QIcon // 116
}

//  struct block end

//  setupUi block begin
// void setupUi(QWidget *EmojiPanel)
func (this *Ui_EmojiPanel) SetupUi(EmojiPanel *qtwidgets.QWidget) {
	this.EmojiPanel = EmojiPanel
	// { // 126
	if EmojiPanel.ObjectName() == "" {
		EmojiPanel.SetObjectName("EmojiPanel")
	}
	EmojiPanel.Resize(371, 355)
	this.VerticalLayout = qtwidgets.NewQVBoxLayout_1(this.EmojiPanel)                 // 111
	this.VerticalLayout.SetSpacing(0)                                                 // 114
	this.VerticalLayout.SetObjectName("VerticalLayout")                               // 112
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)                                // 114
	this.MainWidget = qtwidgets.NewQWidget(this.EmojiPanel, 0)                        // 111
	this.MainWidget.SetObjectName("MainWidget")                                       // 112
	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout_1(this.MainWidget)               // 111
	this.VerticalLayout_2.SetSpacing(3)                                               // 114
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")                           // 112
	this.VerticalLayout_2.SetContentsMargins(0, 3, 0, 3)                              // 114
	this.ScrollArea = qtwidgets.NewQScrollArea(this.MainWidget)                       // 111
	this.ScrollArea.SetObjectName("ScrollArea")                                       // 112
	this.ScrollArea.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)       // 114
	this.ScrollArea.SetWidgetResizable(true)                                          // 114
	this.ScrollAreaWidgetContents = qtwidgets.NewQWidget(nil, 0)                      // 111
	this.ScrollAreaWidgetContents.SetObjectName("ScrollAreaWidgetContents")           // 112
	this.ScrollAreaWidgetContents.SetGeometry(0, 0, 369, 315)                         // 114
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout_1(this.ScrollAreaWidgetContents) // 111
	this.VerticalLayout_3.SetObjectName("VerticalLayout_3")                           // 112
	this.ScrollArea.SetWidget(this.ScrollAreaWidgetContents)                          // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ScrollArea) // 115

	this.GridLayout = qtwidgets.NewQGridLayout(nil)                // 111
	this.GridLayout.SetObjectName("GridLayout")                    // 112
	this.ToolButton_33 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_33.SetObjectName("ToolButton_33")              // 112
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(":/icons/emoji-categories/people.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_33.SetIcon(this.Icon)                                                                              // 114
	this.ToolButton_33.SetIconSize(qtcore.NewQSize_1(20, 20))                                                          // 113
	this.ToolButton_33.SetAutoRaise(true)                                                                              // 114

	this.GridLayout.AddWidget_2(this.ToolButton_33, 0, 0, 1, 1, 0) // 115

	this.ToolButton_34 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_34.SetObjectName("ToolButton_34")              // 112
	this.Icon1 = qtgui.NewQIcon()
	this.Icon1.AddFile(":/icons/emoji-categories/nature.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_34.SetIcon(this.Icon1)                                                                              // 114
	this.ToolButton_34.SetIconSize(qtcore.NewQSize_1(20, 20))                                                           // 113
	this.ToolButton_34.SetAutoRaise(true)                                                                               // 114

	this.GridLayout.AddWidget_2(this.ToolButton_34, 0, 1, 1, 1, 0) // 115

	this.ToolButton_35 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_35.SetObjectName("ToolButton_35")              // 112
	this.Icon2 = qtgui.NewQIcon()
	this.Icon2.AddFile(":/icons/emoji-categories/foods.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_35.SetIcon(this.Icon2)                                                                             // 114
	this.ToolButton_35.SetIconSize(qtcore.NewQSize_1(20, 20))                                                          // 113
	this.ToolButton_35.SetAutoRaise(true)                                                                              // 114

	this.GridLayout.AddWidget_2(this.ToolButton_35, 0, 2, 1, 1, 0) // 115

	this.ToolButton_36 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_36.SetObjectName("ToolButton_36")              // 112
	this.Icon3 = qtgui.NewQIcon()
	this.Icon3.AddFile(":/icons/emoji-categories/activity.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_36.SetIcon(this.Icon3)                                                                                // 114
	this.ToolButton_36.SetIconSize(qtcore.NewQSize_1(20, 20))                                                             // 113
	this.ToolButton_36.SetAutoRaise(true)                                                                                 // 114

	this.GridLayout.AddWidget_2(this.ToolButton_36, 0, 3, 1, 1, 0) // 115

	this.ToolButton_37 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_37.SetObjectName("ToolButton_37")              // 112
	this.Icon4 = qtgui.NewQIcon()
	this.Icon4.AddFile(":/icons/emoji-categories/travel.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_37.SetIcon(this.Icon4)                                                                              // 114
	this.ToolButton_37.SetAutoRaise(true)                                                                               // 114

	this.GridLayout.AddWidget_2(this.ToolButton_37, 0, 4, 1, 1, 0) // 115

	this.ToolButton_38 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_38.SetObjectName("ToolButton_38")              // 112
	this.Icon5 = qtgui.NewQIcon()
	this.Icon5.AddFile(":/icons/emoji-categories/objects.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_38.SetIcon(this.Icon5)                                                                               // 114
	this.ToolButton_38.SetAutoRaise(true)                                                                                // 114

	this.GridLayout.AddWidget_2(this.ToolButton_38, 0, 5, 1, 1, 0) // 115

	this.ToolButton_39 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_39.SetObjectName("ToolButton_39")              // 112
	this.Icon6 = qtgui.NewQIcon()
	this.Icon6.AddFile(":/icons/emoji-categories/symbols.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_39.SetIcon(this.Icon6)                                                                               // 114
	this.ToolButton_39.SetAutoRaise(true)                                                                                // 114

	this.GridLayout.AddWidget_2(this.ToolButton_39, 0, 6, 1, 1, 0) // 115

	this.ToolButton_40 = qtwidgets.NewQToolButton(this.MainWidget) // 111
	this.ToolButton_40.SetObjectName("ToolButton_40")              // 112
	this.Icon7 = qtgui.NewQIcon()
	this.Icon7.AddFile(":/icons/emoji-categories/flags.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_40.SetIcon(this.Icon7)                                                                             // 114
	this.ToolButton_40.SetAutoRaise(true)                                                                              // 114

	this.GridLayout.AddWidget_2(this.ToolButton_40, 0, 7, 1, 1, 0) // 115

	this.VerticalLayout_2.AddLayout(this.GridLayout, 0) // 115

	this.VerticalLayout.Layout().AddWidget(this.MainWidget) // 115

	this.RetranslateUi(EmojiPanel)

	qtcore.QMetaObject_ConnectSlotsByName(EmojiPanel) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *EmojiPanel)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_EmojiPanel) RetranslateUi(EmojiPanel *qtwidgets.QWidget) {
	// noimpl: {
	this.EmojiPanel.SetWindowTitle(qtcore.QCoreApplication_Translate("EmojiPanel", "EmojiPanel", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_33.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "people", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_33.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_33.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_34.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "nature", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_34.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_34.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_35.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "food", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_35.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_35.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_36.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "activity", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_36.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_36.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_37.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "travel", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_37.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_37.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_38.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "objects", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_38.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_38.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_39.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "symbols", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_39.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_39.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: #ifndef QT_NO_TOOLTIP
	this.ToolButton_40.SetToolTip(qtcore.QCoreApplication_Translate("EmojiPanel", "flags", "dummy123", 0))
	// noimpl: #endif // QT_NO_TOOLTIP
	// noimpl: #ifndef QT_NO_STATUSTIP
	this.ToolButton_40.SetStatusTip(qtcore.QCoreApplication_Translate("EmojiPanel", "Back by logic order.(Android Back)", "dummy123", 0))
	// noimpl: #endif // QT_NO_STATUSTIP
	this.ToolButton_40.SetText(qtcore.QCoreApplication_Translate("EmojiPanel", "\342\227\201 ", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_EmojiPanel2() *Ui_EmojiPanel {
	this := &Ui_EmojiPanel{}
	w := qtwidgets.NewQWidget(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_EmojiPanel) QWidget_PTR() *qtwidgets.QWidget {
	return this.EmojiPanel.QWidget_PTR()
}

//  done block end
