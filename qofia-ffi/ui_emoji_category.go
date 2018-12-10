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
func NewUi_EmojiCategory() *Ui_EmojiCategory {
	return &Ui_EmojiCategory{}
}

type Ui_EmojiCategory struct {
	VerticalLayout      *qtwidgets.QVBoxLayout
	Label               *qtwidgets.QLabel
	TableWidget         *qtwidgets.QTableWidget
	EmojiCategory       *qtwidgets.QWidget
	SizePolicy          *qtwidgets.QSizePolicy
	Font                *qtgui.QFont                // 116
	__qtablewidgetitem  *qtwidgets.QTableWidgetItem // 215
	__qtablewidgetitem1 *qtwidgets.QTableWidgetItem // 215
	__qtablewidgetitem2 *qtwidgets.QTableWidgetItem // 215
	__qtablewidgetitem3 *qtwidgets.QTableWidgetItem // 215
	__qtablewidgetitem4 *qtwidgets.QTableWidgetItem // 215
	__qtablewidgetitem5 *qtwidgets.QTableWidgetItem // 215
}

//  struct block end

//  setupUi block begin
// void setupUi(QWidget *EmojiCategory)
func (this *Ui_EmojiCategory) SetupUi(EmojiCategory *qtwidgets.QWidget) {
	this.EmojiCategory = EmojiCategory
	// { // 126
	if EmojiCategory.ObjectName() == "" {
		EmojiCategory.SetObjectName("EmojiCategory")
	}
	EmojiCategory.Resize(400, 300)
	this.SizePolicy = qtwidgets.NewQSizePolicy1(qtwidgets.QSizePolicy__Preferred, qtwidgets.QSizePolicy__Expanding, 1)
	this.SizePolicy.SetHorizontalStretch(0)                                                // 114
	this.SizePolicy.SetVerticalStretch(0)                                                  // 114
	this.SizePolicy.SetHeightForWidth(this.EmojiCategory.SizePolicy().HasHeightForWidth()) // 114
	this.EmojiCategory.SetSizePolicy(this.SizePolicy)                                      // 114
	this.VerticalLayout = qtwidgets.NewQVBoxLayout1(this.EmojiCategory)                    // 111
	this.VerticalLayout.SetSpacing(6)                                                      // 114
	this.VerticalLayout.SetObjectName("VerticalLayout")                                    // 112
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)                                     // 114
	this.Label = qtwidgets.NewQLabel(this.EmojiCategory, 0)                                // 111
	this.Label.SetObjectName("Label")                                                      // 112
	this.Font = qtgui.NewQFont()
	this.Font.SetBold(true)                                               // 114
	this.Font.SetWeight(75)                                               // 114
	this.Label.SetFont(this.Font)                                         // 114
	this.Label.SetTextInteractionFlags(qtcore.Qt__TextBrowserInteraction) // 114

	this.VerticalLayout.Layout().AddWidget(this.Label) // 115

	this.TableWidget = qtwidgets.NewQTableWidget(this.EmojiCategory) // 111
	// if this.TableWidget.ColumnCount() > 3 // 119
	this.TableWidget.SetColumnCount(3)                                    // 114
	this.__qtablewidgetitem = qtwidgets.NewQTableWidgetItem(0)            // 111
	this.TableWidget.SetHorizontalHeaderItem(0, this.__qtablewidgetitem)  // 114
	this.__qtablewidgetitem1 = qtwidgets.NewQTableWidgetItem(0)           // 111
	this.TableWidget.SetHorizontalHeaderItem(1, this.__qtablewidgetitem1) // 114
	this.__qtablewidgetitem2 = qtwidgets.NewQTableWidgetItem(0)           // 111
	this.TableWidget.SetHorizontalHeaderItem(2, this.__qtablewidgetitem2) // 114
	// if this.TableWidget.RowCount() > 3 // 119
	this.TableWidget.SetRowCount(3)                                               // 114
	this.__qtablewidgetitem3 = qtwidgets.NewQTableWidgetItem(0)                   // 111
	this.TableWidget.SetVerticalHeaderItem(0, this.__qtablewidgetitem3)           // 114
	this.__qtablewidgetitem4 = qtwidgets.NewQTableWidgetItem(0)                   // 111
	this.TableWidget.SetVerticalHeaderItem(1, this.__qtablewidgetitem4)           // 114
	this.__qtablewidgetitem5 = qtwidgets.NewQTableWidgetItem(0)                   // 111
	this.TableWidget.SetVerticalHeaderItem(2, this.__qtablewidgetitem5)           // 114
	this.TableWidget.SetObjectName("TableWidget")                                 // 112
	this.TableWidget.SetVerticalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)    // 114
	this.TableWidget.SetHorizontalScrollBarPolicy(qtcore.Qt__ScrollBarAlwaysOff)  // 114
	this.TableWidget.SetEditTriggers(qtwidgets.QAbstractItemView__NoEditTriggers) // 114
	this.TableWidget.SetProperty("showDropIndicator", qtcore.NewQVariant9(false)) // 114
	this.TableWidget.SetDragDropOverwriteMode(false)                              // 114
	this.TableWidget.SetAlternatingRowColors(true)                                // 114
	this.TableWidget.SetSelectionMode(qtwidgets.QAbstractItemView__NoSelection)   // 114

	this.VerticalLayout.Layout().AddWidget(this.TableWidget) // 115

	this.RetranslateUi(EmojiCategory)

	qtcore.QMetaObject_ConnectSlotsByName(EmojiCategory) // 100111
	// } // setupUi // 126

}

// void retranslateUi(QWidget *EmojiCategory)
//  setupUi block end

//  retranslateUi block begin
func (this *Ui_EmojiCategory) RetranslateUi(EmojiCategory *qtwidgets.QWidget) {
	// noimpl: {
	this.EmojiCategory.SetWindowTitle(qtcore.QCoreApplication_Translate("EmojiCategory", "EmojiCategory", "dummy123", 0))
	this.Label.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "TextLabel", "dummy123", 0))
	// noimpl: QTableWidgetItem *___qtablewidgetitem = tableWidget->horizontalHeaderItem(0);
	this.__qtablewidgetitem.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "c1", "dummy123", 0))
	// noimpl: QTableWidgetItem *___qtablewidgetitem1 = tableWidget->horizontalHeaderItem(1);
	this.__qtablewidgetitem1.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "c2", "dummy123", 0))
	// noimpl: QTableWidgetItem *___qtablewidgetitem2 = tableWidget->horizontalHeaderItem(2);
	this.__qtablewidgetitem2.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "c3", "dummy123", 0))
	// noimpl: QTableWidgetItem *___qtablewidgetitem3 = tableWidget->verticalHeaderItem(0);
	this.__qtablewidgetitem3.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "r1", "dummy123", 0))
	// noimpl: QTableWidgetItem *___qtablewidgetitem4 = tableWidget->verticalHeaderItem(1);
	this.__qtablewidgetitem4.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "r2", "dummy123", 0))
	// noimpl: QTableWidgetItem *___qtablewidgetitem5 = tableWidget->verticalHeaderItem(2);
	this.__qtablewidgetitem5.SetText(qtcore.QCoreApplication_Translate("EmojiCategory", "r3", "dummy123", 0))
	// noimpl: } // retranslateUi
	// noimpl:
	// noimpl: };
	// noimpl:
}

//  retranslateUi block end

//  new2 block begin
func NewUi_EmojiCategory2() *Ui_EmojiCategory {
	this := &Ui_EmojiCategory{}
	w := qtwidgets.NewQWidget(nil, 0)
	this.SetupUi(w)
	return this
}

//  new2 block end

//  done block begin

func (this *Ui_EmojiCategory) QWidget_PTR() *qtwidgets.QWidget {
	return this.EmojiCategory.QWidget_PTR()
}

//  done block end
