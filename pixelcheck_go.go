// pixelcheck_go.go — тест экрана на битые пиксели (Go + Fyne)

package main

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

type pixelCheck struct {
	index int
	colors []color.Color
	names  []string
	window fyne.Window
	rect   *canvas.Rectangle
	label  *canvas.Text
}

func (p *pixelCheck) nextColor() {
	p.index = (p.index + 1) % len(p.colors)
	p.rect.FillColor = p.colors[p.index]
	// Обновляем текст
	c := p.colors[p.index]
	r, g, b, _ := c.RGBA()
	brightness := int((r + g + b) / 3 >> 8)
	p.label.Text = p.names[p.index]
	if brightness < 128 {
		p.label.Color = color.White
	} else {
		p.label.Color = color.Black
	}
	p.label.Refresh()
	p.rect.Refresh()
}

func main() {
	a := app.New()
	w := a.NewWindow("🖥️ PixelCheck Pro")
	w.SetFullScreen(true)
	w.SetMaster()

	colors := []color.Color{
		color.Black, color.White, color.RGBA{255,0,0,255},
		color.RGBA{0,255,0,255}, color.RGBA{0,0,255,255},
		color.RGBA{255,255,0,255}, color.RGBA{255,0,255,255},
		color.RGBA{0,255,255,255}, color.RGBA{128,128,128,255},
	}
	names := []string{
		"Чёрный", "Белый", "Красный", "Зелёный", "Синий",
		"Жёлтый", "Пурпурный", "Голубой", "Серый",
	}
	p := &pixelCheck{
		index:  0,
		colors: colors,
		names:  names,
		window: w,
	}

	p.rect = canvas.NewRectangle(colors[0])
	p.rect.Resize(fyne.NewSize(800, 600))

	p.label = canvas.NewText(names[0], color.White)
	p.label.TextSize = 24
	p.label.Alignment = fyne.TextAlignCenter

	content := container.NewWithoutLayout(p.rect, p.label)
	p.rect.Move(fyne.NewPos(0, 0))
	p.label.Move(fyne.NewPos(20, 20))
	w.SetContent(content)

	// Обработка кликов и клавиш
	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeySpace {
			p.nextColor()
		} else if ke.Name == fyne.KeyEscape {
			w.Close()
		}
	})
	w.Canvas().SetOnTypedRune(func(r rune) {
		if r == ' ' {
			p.nextColor()
		}
	})
	if desk, ok := w.Canvas().(desktop.Canvas); ok {
		desk.SetOnMouseDown(func(me *desktop.MouseEvent) {
			p.nextColor()
		})
	}

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
