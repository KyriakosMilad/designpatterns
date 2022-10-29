package main

const (
	WinOSType = "Win"
	MacOSType = "Mac"
)

type MacFactory struct{}
type MacBtn struct{}

func (mf *MacFactory) createBtn() *MacBtn {
	return &MacBtn{}
}

type WinFactory struct{}
type WinBtn struct{}

func (wf *WinFactory) createBtn() *WinBtn {
	return &WinBtn{}
}

type GenFactory struct{}

type Button interface{}

func (g *GenFactory) createBtn(os string) Button {
	if os == WinOSType {
		wf := &WinFactory{}
		return wf.createBtn()
	} else if os == MacOSType {
		mf := &MacFactory{}
		return mf.createBtn()
	}

	return nil
}
