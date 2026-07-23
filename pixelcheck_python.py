# pixelcheck_python.py — тест экрана на битые пиксели (Python + Tkinter)

import tkinter as tk
import sys

class PixelCheck:
    def __init__(self):
        self.root = tk.Tk()
        self.root.title("🖥️ PixelCheck Pro")
        self.root.attributes('-fullscreen', True)
        self.root.bind("<Escape>", self.quit)
        self.root.bind("<space>", self.next_color)
        self.root.bind("<Button-1>", self.next_color)
        self.root.bind("<Configure>", self.on_resize)

        self.colors = [
            ("Чёрный", "#000000"),
            ("Белый", "#FFFFFF"),
            ("Красный", "#FF0000"),
            ("Зелёный", "#00FF00"),
            ("Синий", "#0000FF"),
            ("Жёлтый", "#FFFF00"),
            ("Пурпурный", "#FF00FF"),
            ("Голубой", "#00FFFF"),
            ("Серый", "#808080")
        ]
        self.current = 0
        self.canvas = tk.Canvas(self.root, highlightthickness=0)
        self.canvas.pack(fill=tk.BOTH, expand=True)
        self.update_color()
        self.info_label = tk.Label(self.root, text="", font=("Arial", 18),
                                   fg="white", bg="black")
        self.info_label.place(relx=0.02, rely=0.02)

        self.root.mainloop()

    def on_resize(self, event):
        self.canvas.config(width=event.width, height=event.height)
        self.update_color()

    def update_color(self):
        color_name, hex_color = self.colors[self.current]
        self.canvas.config(bg=hex_color)
        # Определяем цвет текста для контраста
        brightness = sum(int(hex_color[i:i+2], 16) for i in (1,3,5)) / 3
        fg = "white" if brightness < 128 else "black"
        self.info_label.config(text=f"{color_name} ({hex_color})", fg=fg)
        # Изменяем цвет фона метки на противоположный для читаемости
        self.info_label.config(bg=hex_color)

    def next_color(self, event=None):
        self.current = (self.current + 1) % len(self.colors)
        self.update_color()

    def quit(self, event=None):
        self.root.destroy()
        sys.exit()

if __name__ == "__main__":
    PixelCheck()
