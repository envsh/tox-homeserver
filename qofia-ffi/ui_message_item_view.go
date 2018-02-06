package main

//  header block begin
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtwidgets"
import "qt.go/qtquickwidgets"
import "qt.go/qtmock"

func init() { qtcore.KeepMe() }
func init() { qtwidgets.KeepMe() }
func init() { qtquickwidgets.KeepMe() }
func init() { qtmock.KeepMe() }

//  header block end

//  struct block begin
func NewUi_MessageItemView() *Ui_MessageItemView {
	return &Ui_MessageItemView{}
}

type Ui_MessageItemView struct {
	HorizontalLayout_2 *qtwidgets.QHBoxLayout
	VerticalLayout_3   *qtwidgets.QVBoxLayout
	Label              *qtwidgets.QLabel
	VerticalSpacer     *qtwidgets.QSpacerItem
	VerticalLayout     *qtwidgets.QVBoxLayout
	HorizontalLayout   *qtwidgets.QHBoxLayout
	Label_3            *qtwidgets.QLabel
	Label_4            *qtwidgets.QLabel
	ToolButton         *qtwidgets.QToolButton
	TextBrowser        *qtwidgets.QTextBrowser
	VerticalLayout_2   *qtwidgets.QVBoxLayout
	Label_2            *qtwidgets.QLabel
	VerticalSpacer_2   *qtwidgets.QSpacerItem
	MessageItemView    *qtwidgets.QWidget
	SizePolicy         *qtwidgets.QSizePolicy
}

//  struct block end

//  setupUi block begin
// void setupUi(QWidget *MessageItemView)
func (this *Ui_MessageItemView) SetupUi(MessageItemView *qtwidgets.QWidget) {
	this.MessageItemView = MessageItemView
	// { // 126
	if MessageItemView.ObjectName().IsEmpty() {
		MessageItemView.SetObjectName(qtcore.NewQString_5("MessageItemView"))
	}
	MessageItemView.Resize(399, 133)
	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout_1(qtwidgets.NewQWidgetFromPointer(this.MessageItemView.GetCthis())) // 111
	this.HorizontalLayout_2.SetObjectName(qtcore.NewQString_5("HorizontalLayout_2"))                                       // 112
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout()                                                                     // 111
	this.VerticalLayout_3.SetObjectName(qtcore.NewQString_5("VerticalLayout_3"))                                           // 112
	this.Label = qtwidgets.NewQLabel(this.MessageItemView, 0)                                                              // 111
	this.Label.SetObjectName(qtcore.NewQString_5("Label"))                                                                 // 112
	this.Label.SetPixmap(qtgui.NewQPixmap_3(qtcore.NewQString_5(":/Icons/Icon_avatar_40.Png"), "dummy123", 0))             // 114

	this.VerticalLayout_3.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label.GetCthis())) // 115

	this.VerticalSpacer = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_3.AddItem(qtwidgets.NewQLayoutItemFromPointer(this.VerticalSpacer.GetCthis())) // 115

	this.HorizontalLayout_2.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_3.GetCthis()), 0) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()                                                           // 111
	this.VerticalLayout.SetObjectName(qtcore.NewQString_5("VerticalLayout"))                                   // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()                                                         // 111
	this.HorizontalLayout.SetObjectName(qtcore.NewQString_5("HorizontalLayout"))                               // 112
	this.Label_3 = qtwidgets.NewQLabel(this.MessageItemView, 0)                                                // 111
	this.Label_3.SetObjectName(qtcore.NewQString_5("Label_3"))                                                 // 112
	this.Label_3.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_3.GetCthis())) // 115

	this.Label_4 = qtwidgets.NewQLabel(this.MessageItemView, 0)                                                // 111
	this.Label_4.SetObjectName(qtcore.NewQString_5("Label_4"))                                                 // 112
	this.Label_4.SetTextInteractionFlags(qtcore.Qt__LinksAccessibleByMouse | qtcore.Qt__TextSelectableByMouse) // 114

	this.HorizontalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_4.GetCthis())) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.MessageItemView) // 111
	this.ToolButton.SetObjectName(qtcore.NewQString_5("ToolButton")) // 112
	this.ToolButton.SetAutoRaise(true)                               // 114

	this.HorizontalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.ToolButton.GetCthis())) // 115

	this.VerticalLayout.AddLayout(qtwidgets.NewQLayoutFromPointer(this.HorizontalLayout.GetCthis()), 0) // 115

	this.TextBrowser = qtwidgets.NewQTextBrowser(this.MessageItemView) // 111
	this.TextBrowser.SetObjectName(qtcore.NewQString_5("TextBrowser")) // 112
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                               // 114
	this.SizePolicy.SetVerticalStretch(0)                                                 // 114
	this.SizePolicy.SetHeightForWidth(this.TextBrowser.SizePolicy().HasHeightForWidth())  // 114
	this.TextBrowser.SetSizePolicy(this.SizePolicy)                                       // 114
	this.TextBrowser.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)          // 114
	this.TextBrowser.SetSizeAdjustPolicy(qtwidgets.QAbstractScrollArea__AdjustToContents) // 114
	this.TextBrowser.SetOpenExternalLinks(true)                                           // 114

	this.VerticalLayout.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.TextBrowser.GetCthis())) // 115

	this.HorizontalLayout_2.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout.GetCthis()), 0) // 115

	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout()                                                           // 111
	this.VerticalLayout_2.SetObjectName(qtcore.NewQString_5("VerticalLayout_2"))                                 // 112
	this.Label_2 = qtwidgets.NewQLabel(this.MessageItemView, 0)                                                  // 111
	this.Label_2.SetObjectName(qtcore.NewQString_5("Label_2"))                                                   // 112
	this.Label_2.SetPixmap(qtgui.NewQPixmap_3(qtcore.NewQString_5(":/Icons/Icon_avatar_40.Png"), "dummy123", 0)) // 114

	this.VerticalLayout_2.Layout().AddWidget(qtwidgets.NewQWidgetFromPointer(this.Label_2.GetCthis())) // 115

	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(20, 40, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_2.AddItem(qtwidgets.NewQLayoutItemFromPointer(this.VerticalSpacer_2.GetCthis())) // 115

	this.HorizontalLayout_2.AddLayout(qtwidgets.NewQLayoutFromPointer(this.VerticalLayout_2.GetCthis()), 0) // 115

	this.RetranslateUi(MessageItemView)

	qtcore.QMetaObject_ConnectSlotsByName(qtcore.NewQObjectFromPointer(MessageItemView.GetCthis())) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *MessageItemView)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MessageItemView) RetranslateUi(MessageItemView *qtwidgets.QWidget) {
	this.MessageItemView.SetWindowTitle(qtcore.QCoreApplication_Translate("MessageItemView", "Form", "dummy123", 0))
	this.Label_3.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "TextLabel", "dummy123", 0))
	this.Label_4.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "TextLabel", "dummy123", 0))
	this.ToolButton.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "...", "dummy123", 0))
}

//  retranslateUi block end
