// pixelcheck_rs.rs — тест экрана на битые пиксели (Rust + minifb)

use minifb::{Key, Window, WindowOptions};

const WIDTH: usize = 800;
const HEIGHT: usize = 600;

struct ColorMode {
    name: &'static str,
    r: u8,
    g: u8,
    b: u8,
}

fn main() {
    let mut window = Window::new(
        "🖥️ PixelCheck Pro",
        WIDTH,
        HEIGHT,
        WindowOptions {
            fullscreen: true,
            ..WindowOptions::default()
        },
    )
    .expect("Не удалось создать окно");

    window.limit_update_rate(Some(std::time::Duration::from_micros(16600)));

    let colors = vec![
        ColorMode { name: "Чёрный", r: 0, g: 0, b: 0 },
        ColorMode { name: "Белый", r: 255, g: 255, b: 255 },
        ColorMode { name: "Красный", r: 255, g: 0, b: 0 },
        ColorMode { name: "Зелёный", r: 0, g: 255, b: 0 },
        ColorMode { name: "Синий", r: 0, g: 0, b: 255 },
        ColorMode { name: "Жёлтый", r: 255, g: 255, b: 0 },
        ColorMode { name: "Пурпурный", r: 255, g: 0, b: 255 },
        ColorMode { name: "Голубой", r: 0, g: 255, b: 255 },
        ColorMode { name: "Серый", r: 128, g: 128, b: 128 },
    ];
    let mut current = 0;

    while window.is_open() && !window.is_key_down(Key::Escape) {
        if window.is_key_down(Key::Space) || window.is_key_down(Key::Space) {
            current = (current + 1) % colors.len();
            // Небольшая задержка, чтобы избежать множественных переключений
            std::thread::sleep(std::time::Duration::from_millis(150));
        }

        let cm = &colors[current];
        let color = (cm.r, cm.g, cm.b);

        // Создаём буфер
        let mut buffer: Vec<u32> = vec![0; WIDTH * HEIGHT];
        let pixel = ((color.0 as u32) << 16) | ((color.1 as u32) << 8) | (color.2 as u32);
        for p in buffer.iter_mut() {
            *p = pixel;
        }

        // Рисуем текст (упрощённо — просто выводим в заголовок)
        window.set_title(&format!("🖥️ PixelCheck Pro - {} (R:{},G:{},B:{})",
                                  cm.name, color.0, color.1, color.2));

        window.update_with_buffer(&buffer, WIDTH, HEIGHT).expect("Ошибка обновления");
    }
}
