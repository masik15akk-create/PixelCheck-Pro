// PixelCheckJava.java — тест экрана на битые пиксели (Java + Swing)

import javax.swing.*;
import java.awt.*;
import java.awt.event.*;

public class PixelCheckJava extends JFrame {
    private static final Color[] COLORS = {
        Color.BLACK, Color.WHITE, Color.RED, Color.GREEN, Color.BLUE,
        Color.YELLOW, Color.MAGENTA, Color.CYAN, Color.GRAY
    };
    private static final String[] NAMES = {
        "Чёрный", "Белый", "Красный", "Зелёный", "Синий",
        "Жёлтый", "Пурпурный", "Голубой", "Серый"
    };
    private int index = 0;
    private JLabel infoLabel;

    public PixelCheckJava() {
        setTitle("🖥️ PixelCheck Pro");
        setUndecorated(true);
        setExtendedState(JFrame.MAXIMIZED_BOTH);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setBackground(COLORS[0]);

        infoLabel = new JLabel(NAMES[0], SwingConstants.CENTER);
        infoLabel.setFont(new Font("Arial", Font.BOLD, 24));
        infoLabel.setForeground(Color.WHITE);
        add(infoLabel, BorderLayout.SOUTH);

        addMouseListener(new MouseAdapter() {
            public void mouseClicked(MouseEvent e) { nextColor(); }
        });

        addKeyListener(new KeyAdapter() {
            public void keyPressed(KeyEvent e) {
                if (e.getKeyCode() == KeyEvent.VK_SPACE) nextColor();
                else if (e.getKeyCode() == KeyEvent.VK_ESCAPE) System.exit(0);
            }
        });

        setVisible(true);
    }

    private void nextColor() {
        index = (index + 1) % COLORS.length;
        getContentPane().setBackground(COLORS[index]);
        infoLabel.setText(NAMES[index]);
        // Адаптивный цвет текста
        Color c = COLORS[index];
        int brightness = (c.getRed() + c.getGreen() + c.getBlue()) / 3;
        infoLabel.setForeground(brightness < 128 ? Color.WHITE : Color.BLACK);
    }

    public static void main(String[] args) {
        SwingUtilities.invokeLater(PixelCheckJava::new);
    }
}
