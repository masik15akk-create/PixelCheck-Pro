// pixelcheck_cpp.cpp — тест экрана на битые пиксели (C++ + SFML)

#include <SFML/Graphics.hpp>
#include <vector>
#include <string>

struct ColorMode {
    std::string name;
    sf::Color color;
};

int main() {
    sf::VideoMode desktop = sf::VideoMode::getDesktopMode();
    sf::RenderWindow window(desktop, "🖥️ PixelCheck Pro", sf::Style::Fullscreen);
    window.setMouseCursorVisible(false);
    window.setVerticalSyncEnabled(true);

    std::vector<ColorMode> colors = {
        {"Чёрный", sf::Color::Black},
        {"Белый", sf::Color::White},
        {"Красный", sf::Color::Red},
        {"Зелёный", sf::Color::Green},
        {"Синий", sf::Color::Blue},
        {"Жёлтый", sf::Color::Yellow},
        {"Пурпурный", sf::Color::Magenta},
        {"Голубой", sf::Color::Cyan},
        {"Серый", sf::Color(128,128,128)}
    };
    size_t current = 0;

    sf::Font font;
    if (!font.loadFromFile("/usr/share/fonts/truetype/liberation/LiberationSans-Regular.ttf")) {
        // fallback: используем встроенный шрифт, но для простоты оставим пустым
    }
    sf::Text info;
    info.setFont(font);
    info.setCharacterSize(24);
    info.setPosition(20, 20);

    bool running = true;
    while (running) {
        sf::Event event;
        while (window.pollEvent(event)) {
            if (event.type == sf::Event::Closed || 
                (event.type == sf::Event::KeyPressed && event.key.code == sf::Keyboard::Escape)) {
                running = false;
            }
            if (event.type == sf::Event::KeyPressed && event.key.code == sf::Keyboard::Space) {
                current = (current + 1) % colors.size();
            }
            if (event.type == sf::Event::MouseButtonPressed) {
                current = (current + 1) % colors.size();
            }
        }

        const auto& mode = colors[current];
        window.clear(mode.color);
        info.setString(mode.name + " (" + std::to_string(mode.color.r) + "," +
                       std::to_string(mode.color.g) + "," + std::to_string(mode.color.b) + ")");
        // Контрастный цвет текста
        int brightness = (mode.color.r + mode.color.g + mode.color.b) / 3;
        info.setFillColor(brightness < 128 ? sf::Color::White : sf::Color::Black);
        window.draw(info);
        window.display();
    }
    return 0;
}
