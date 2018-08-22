package gtkbuilder

import (
	"unsafe"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//Event 事件结构体
type Event struct {
	id   int
	work bool

	s string
	f interface{}
}

//Builder 自建的builder结构体，继承自gtk builder
type Builder struct {
	*gtk.Builder
	Events []*Event
}

//NewBuilder 从gtk build 生成builder结构体
func NewBuilder(gtkBuilder *gtk.Builder) *Builder {
	b := new(Builder)
	b.Builder = gtkBuilder

	return b
}

//Connect 设置 builder 的动作监听
func (b *Builder) Connect(s string, f interface{}) int {
	e := new(Event)
	e.id = len(b.Events)
	e.work = true
	b.Events = append(b.Events, e)

	e.s = s
	e.f = f
	return e.id
}

//Listening builder的动作监听函数
func (b *Builder) Listening() {
	b.ConnectSignalsFull(func(builder *gtk.Builder, obj *glib.GObject, sig, handler string, conn *glib.GObject, flags glib.ConnectFlags, user_data interface{}) {
		for i := range b.Events {
			e := b.Events[i]
			if e.work && e.s == handler {
				obj.SignalConnect(sig, e.f, user_data, flags)
			}
		}
	}, nil)
}

////////////////////////

func (b *Builder) GetWindow(name string) *gtk.Window {
	widget := new(gtk.Window)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetDialog(name string) *gtk.Dialog {
	widget := new(gtk.Dialog)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetAboutDialog(name string) *gtk.AboutDialog {
	widget := new(gtk.AboutDialog)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetContainer(name string) *gtk.Container {
	widget := new(gtk.Container)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetWidget(name string) *gtk.Widget {
	widget := new(gtk.Widget)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetFontSelectionDialog(name string) *gtk.FontSelectionDialog {
	widget := new(gtk.FontSelectionDialog)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMisc(name string) *gtk.Misc {
	widget := new(gtk.Misc)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetLabel(name string) *gtk.Label {
	widget := new(gtk.Label)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetButton(name string) *gtk.Button {
	widget := new(gtk.Button)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetSpinButton(name string) *gtk.SpinButton {
	widget := new(gtk.SpinButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetRadioButton(name string) *gtk.RadioButton {
	widget := new(gtk.RadioButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetFontButton(name string) *gtk.FontButton {
	widget := new(gtk.FontButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetLinkButton(name string) *gtk.LinkButton {
	widget := new(gtk.LinkButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetComboBox(name string) *gtk.ComboBox {
	widget := new(gtk.ComboBox)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetComboBoxEntry(name string) *gtk.ComboBoxEntry {
	widget := new(gtk.ComboBoxEntry)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMessageDialog(name string) *gtk.MessageDialog {
	widget := new(gtk.MessageDialog)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetComboBoxText(name string) *gtk.ComboBoxText {
	widget := new(gtk.ComboBoxText)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetBin(name string) *gtk.Bin {
	widget := new(gtk.Bin)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetStatusbar(name string) *gtk.Statusbar {
	widget := new(gtk.Statusbar)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetInfoBar(name string) *gtk.InfoBar {
	widget := new(gtk.InfoBar)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetFrame(name string) *gtk.Frame {
	widget := new(gtk.Frame)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetBox(name string) *gtk.Box {
	widget := new(gtk.Box)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetPaned(name string) *gtk.Paned {
	widget := new(gtk.Paned)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToggleButton(name string) *gtk.ToggleButton {
	widget := new(gtk.ToggleButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetAccelLabel(name string) *gtk.AccelLabel {
	widget := new(gtk.AccelLabel)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetScaleButton(name string) *gtk.ScaleButton {
	widget := new(gtk.ScaleButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetEntry(name string) *gtk.Entry {
	widget := new(gtk.Entry)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetArrow(name string) *gtk.Arrow {
	widget := new(gtk.Arrow)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetTextView(name string) *gtk.TextView {
	widget := new(gtk.TextView)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMenu(name string) *gtk.Menu {
	widget := new(gtk.Menu)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMenuBar(name string) *gtk.MenuBar {
	widget := new(gtk.MenuBar)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMenuShell(name string) *gtk.MenuShell {
	widget := new(gtk.MenuShell)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMenuItem(name string) *gtk.MenuItem {
	widget := new(gtk.MenuItem)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetItem(name string) *gtk.Item {
	widget := new(gtk.Item)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToolbar(name string) *gtk.Toolbar {
	widget := new(gtk.Toolbar)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToolItem(name string) *gtk.ToolItem {
	widget := new(gtk.ToolItem)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetSeparatorToolItem(name string) *gtk.SeparatorToolItem {
	widget := new(gtk.SeparatorToolItem)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToolButton(name string) *gtk.ToolButton {
	widget := new(gtk.ToolButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToolPalette(name string) *gtk.ToolPalette {
	widget := new(gtk.ToolPalette)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToolItemGroup(name string) *gtk.ToolItemGroup {
	widget := new(gtk.ToolItemGroup)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetMenuToolButton(name string) *gtk.MenuToolButton {
	widget := new(gtk.MenuToolButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetToggleToolButton(name string) *gtk.ToggleToolButton {
	widget := new(gtk.ToggleToolButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetScrolledWindow(name string) *gtk.ScrolledWindow {
	widget := new(gtk.ScrolledWindow)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetViewport(name string) *gtk.Viewport {
	widget := new(gtk.Viewport)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetTreeView(name string) *gtk.TreeView {
	widget := new(gtk.TreeView)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetIconView(name string) *gtk.IconView {
	widget := new(gtk.IconView)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetScale(name string) *gtk.Scale {
	widget := new(gtk.Scale)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetRange(name string) *gtk.Range {
	widget := new(gtk.Range)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetImage(name string) *gtk.Image {
	widget := new(gtk.Image)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetNotebook(name string) *gtk.Notebook {
	widget := new(gtk.Notebook)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetTable(name string) *gtk.Table {
	widget := new(gtk.Table)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetDrawingArea(name string) *gtk.DrawingArea {
	widget := new(gtk.DrawingArea)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetSpinner(name string) *gtk.Spinner {
	widget := new(gtk.Spinner)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetAssistant(name string) *gtk.Assistant {
	widget := new(gtk.Assistant)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetExpander(name string) *gtk.Expander {
	widget := new(gtk.Expander)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetAlignment(name string) *gtk.Alignment {
	widget := new(gtk.Alignment)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetProgressBar(name string) *gtk.ProgressBar {
	widget := new(gtk.ProgressBar)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetFixed(name string) *gtk.Fixed {
	widget := new(gtk.Fixed)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetCheckMenuItem(name string) *gtk.CheckMenuItem {
	widget := new(gtk.CheckMenuItem)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetRadioMenuItem(name string) *gtk.RadioMenuItem {
	widget := new(gtk.RadioMenuItem)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetLayout(name string) *gtk.Layout {
	widget := new(gtk.Layout)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetColorButton(name string) *gtk.ColorButton {
	widget := new(gtk.ColorButton)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetImageMenuItem(name string) *gtk.ImageMenuItem {
	widget := new(gtk.ImageMenuItem)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetActivatable(name string) *gtk.Activatable {
	widget := new(gtk.Activatable)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}

func (b *Builder) GetFontSelection(name string) *gtk.FontSelection {
	widget := new(gtk.FontSelection)
	widget.GWidget = gtk.AS_GWIDGET(unsafe.Pointer(b.GetObject(name).Object))
	return widget
}
