package main

import (
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// アプリケーションのインスタンスを作成
	a := app.New()

	// ウィンドウを作成
	w := a.NewWindow("GoSurf")

	// ウィンドウのサイズを設定
	w.Resize(fyne.NewSize(800, 600))

	// URL入力フィールド
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL and press Enter")

	// Webコンテンツ表示用ラベル
	contentLabel := widget.NewLabel("")
	contentLabel.Wrapping = fyne.TextWrapWord

	//URLを読み込む関数
	loadURL := func(url string) {
		resp, err := http.Get(url)
		if err != nil {
			contentLabel.SetText(fmt.Sprintf("Failed to load URL: %s", err))
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			contentLabel.SetText(fmt.Sprintf("Failed to read content: %s", err))
			return
		}
		contentLabel.SetText(string(body))
	}

	// Enterキーを押したときにURLを読み込む
	urlEntry.OnSubmitted = func(url string) {
		loadURL(url)
	}

	// UIを設定
	w.SetContent(container.NewVBox(
		urlEntry,
		widget.NewButton("Go Home", func() {
			loadURL("http://google.com") // ホームページ（例）に戻る
		}),
		contentLabel,
	))

	w.ShowAndRun()
}
