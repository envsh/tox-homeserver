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
func NewUi_MessageItemView() *Ui_MessageItemView {
	return &Ui_MessageItemView{}
}

type Ui_MessageItemView struct {
	HorizontalLayout_2        *qtwidgets.QHBoxLayout
	VerticalLayout_3          *qtwidgets.QVBoxLayout
	ToolButton_2              *qtwidgets.QToolButton
	VerticalSpacer            *qtwidgets.QSpacerItem
	VerticalLayout            *qtwidgets.QVBoxLayout
	HorizontalLayout          *qtwidgets.QHBoxLayout
	LabelUserName4MessageItem *qtwidgets.QLabel
	HorizontalSpacer          *qtwidgets.QSpacerItem
	LabelMsgTime              *qtwidgets.QLabel
	ToolButton                *qtwidgets.QToolButton
	Label_5                   *qtwidgets.QLabel
	VerticalLayout_2          *qtwidgets.QVBoxLayout
	ToolButton_3              *qtwidgets.QToolButton
	VerticalSpacer_2          *qtwidgets.QSpacerItem
	MessageItemView           *qtwidgets.QWidget
	SizePolicy                *qtwidgets.QSizePolicy
	Icon                      *qtgui.QIcon // 116
	Font                      *qtgui.QFont // 116
	SizePolicy1               *qtwidgets.QSizePolicy
}

//  struct block end

//  setupUi block begin
// void setupUi(QWidget *MessageItemView)
func (this *Ui_MessageItemView) SetupUi(MessageItemView *qtwidgets.QWidget) {
	this.MessageItemView = MessageItemView
	// { // 126
	if MessageItemView.ObjectName() == "" {
		MessageItemView.SetObjectName("MessageItemView")
	}
	MessageItemView.Resize(399, 62)
	this.SizePolicy = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Fixed, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                  // 114
	this.SizePolicy.SetVerticalStretch(0)                                                    // 114
	this.SizePolicy.SetHeightForWidth(this.MessageItemView.SizePolicy().HasHeightForWidth()) // 114
	this.MessageItemView.SetSizePolicy(this.SizePolicy)                                      // 114
	this.MessageItemView.SetFocusPolicy(qtcore.Qt__ClickFocus)                               // 114
	this.HorizontalLayout_2 = qtwidgets.NewQHBoxLayout_1(this.MessageItemView)               // 111
	this.HorizontalLayout_2.SetSpacing(0)                                                    // 114
	this.HorizontalLayout_2.SetObjectName("HorizontalLayout_2")                              // 112
	this.HorizontalLayout_2.SetContentsMargins(0, 0, 0, 0)                                   // 114
	this.VerticalLayout_3 = qtwidgets.NewQVBoxLayout()                                       // 111
	this.VerticalLayout_3.SetObjectName("VerticalLayout_3")                                  // 112
	this.ToolButton_2 = qtwidgets.NewQToolButton(this.MessageItemView)                       // 111
	this.ToolButton_2.SetObjectName("ToolButton_2")                                          // 112
	this.ToolButton_2.SetFocusPolicy(qtcore.Qt__NoFocus)                                     // 114
	this.Icon = qtgui.NewQIcon()
	this.Icon.AddFile(":/icons/icon_avatar_40.png", qtcore.NewQSize(), qtgui.QIcon__Normal, qtgui.QIcon__Off) // 115
	this.ToolButton_2.SetIcon(this.Icon)                                                                      // 114
	this.ToolButton_2.SetIconSize(qtcore.NewQSize_1(26, 26))                                                  // 113
	this.ToolButton_2.SetAutoRaise(true)                                                                      // 114

	this.VerticalLayout_3.Layout().AddWidget(this.ToolButton_2) // 115

	this.VerticalSpacer = qtwidgets.NewQSpacerItem(36, 10, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_3.AddItem(this.VerticalSpacer) // 115

	this.HorizontalLayout_2.AddLayout(this.VerticalLayout_3, 0) // 115

	this.VerticalLayout = qtwidgets.NewQVBoxLayout()                              // 111
	this.VerticalLayout.SetObjectName("VerticalLayout")                           // 112
	this.HorizontalLayout = qtwidgets.NewQHBoxLayout()                            // 111
	this.HorizontalLayout.SetObjectName("HorizontalLayout")                       // 112
	this.LabelUserName4MessageItem = qtwidgets.NewQLabel(this.MessageItemView, 0) // 111
	this.LabelUserName4MessageItem.SetObjectName("LabelUserName4MessageItem")     // 112
	this.Font = qtgui.NewQFont()
	this.Font.SetBold(false)                                                                  // 114
	this.Font.SetWeight(50)                                                                   // 114
	this.LabelUserName4MessageItem.SetFont(this.Font)                                         // 114
	this.LabelUserName4MessageItem.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.HorizontalLayout.Layout().AddWidget(this.LabelUserName4MessageItem) // 115

	this.HorizontalSpacer = qtwidgets.NewQSpacerItem(10, 20, qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Minimum)

	this.HorizontalLayout.AddItem(this.HorizontalSpacer) // 115

	this.LabelMsgTime = qtwidgets.NewQLabel(this.MessageItemView, 0)                                           // 111
	this.LabelMsgTime.SetObjectName("LabelMsgTime")                                                            // 112
	this.LabelMsgTime.SetAlignment(qtcore.Qt__AlignRight | qtcore.Qt__AlignTrailing | qtcore.Qt__AlignVCenter) // 114
	this.LabelMsgTime.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction)                               // 114

	this.HorizontalLayout.Layout().AddWidget(this.LabelMsgTime) // 115

	this.ToolButton = qtwidgets.NewQToolButton(this.MessageItemView) // 111
	this.ToolButton.SetObjectName("ToolButton")                      // 112
	this.ToolButton.SetFocusPolicy(qtcore.Qt__NoFocus)               // 114
	this.ToolButton.SetAutoRaise(true)                               // 114

	this.HorizontalLayout.Layout().AddWidget(this.ToolButton) // 115

	this.VerticalLayout.AddLayout(this.HorizontalLayout, 0) // 115

	this.Label_5 = qtwidgets.NewQLabel(this.MessageItemView, 0) // 111
	this.Label_5.SetObjectName("Label_5")                       // 112
	this.SizePolicy1 = qtwidgets.NewQSizePolicy_1(qtwidgets.QSizePolicy__Expanding, qtwidgets.QSizePolicy__Preferred, 1)
	this.SizePolicy1.SetHorizontalStretch(0)                                          // 114
	this.SizePolicy1.SetVerticalStretch(0)                                            // 114
	this.SizePolicy1.SetHeightForWidth(this.Label_5.SizePolicy().HasHeightForWidth()) // 114
	this.Label_5.SetSizePolicy(this.SizePolicy1)                                      // 114
	this.Label_5.SetContextMenuPolicy(qtcore.Qt__DefaultContextMenu)                  // 114
	this.Label_5.SetWordWrap(true)                                                    // 114
	this.Label_5.SetOpenExternalLinks(true)                                           // 114
	this.Label_5.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction)           // 114

	this.VerticalLayout.Layout().AddWidget(this.Label_5) // 115

	this.HorizontalLayout_2.AddLayout(this.VerticalLayout, 0) // 115

	this.VerticalLayout_2 = qtwidgets.NewQVBoxLayout()                 // 111
	this.VerticalLayout_2.SetObjectName("VerticalLayout_2")            // 112
	this.ToolButton_3 = qtwidgets.NewQToolButton(this.MessageItemView) // 111
	this.ToolButton_3.SetObjectName("ToolButton_3")                    // 112
	this.ToolButton_3.SetFocusPolicy(qtcore.Qt__NoFocus)               // 114
	this.ToolButton_3.SetIcon(this.Icon)                               // 114
	this.ToolButton_3.SetIconSize(qtcore.NewQSize_1(28, 28))           // 113
	this.ToolButton_3.SetAutoRaise(true)                               // 114

	this.VerticalLayout_2.Layout().AddWidget(this.ToolButton_3) // 115

	this.VerticalSpacer_2 = qtwidgets.NewQSpacerItem(36, 10, qtwidgets.QSizePolicy__Minimum, qtwidgets.QSizePolicy__Expanding)

	this.VerticalLayout_2.AddItem(this.VerticalSpacer_2) // 115

	this.HorizontalLayout_2.AddLayout(this.VerticalLayout_2, 0) // 115

	this.RetranslateUi(MessageItemView)

	qtcore.QMetaObject_ConnectSlotsByName(MessageItemView) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *MessageItemView)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_MessageItemView) RetranslateUi(MessageItemView *qtwidgets.QWidget) {
	// noimpl: {
	this.MessageItemView.SetWindowTitle(qtcore.QCoreApplication_Translate("MessageItemView", "Form", "dummy123", 0))
	this.ToolButton_2.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "...", "dummy123", 0))
	this.LabelUserName4MessageItem.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "TextLabel", "dummy123", 0))
	this.LabelMsgTime.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "TextLabel", "dummy123", 0))
	this.ToolButton.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "...", "dummy123", 0))
	// noimpl: label_5->setText(QString());
	this.ToolButton_3.SetText(qtcore.QCoreApplication_Translate("MessageItemView", "...", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_MessageItemView2() *Ui_MessageItemView {
	this := &Ui_MessageItemView{}
	w := qtwidgets.NewQWidget(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_MessageItemView) QWidget_PTR() *qtwidgets.QWidget {
	return this.MessageItemView.QWidget_PTR()
}

//  done block end
