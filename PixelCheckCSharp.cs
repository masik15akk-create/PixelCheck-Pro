// PixelCheckCSharp.cs — тест экрана на битые пиксели (C# + Windows Forms)

using System;
using System.Drawing;
using System.Windows.Forms;

class PixelCheckForm : Form
{
    private int index = 0;
    private Color[] colors = {
        Color.Black, Color.White, Color.Red, Color.Green, Color.Blue,
        Color.Yellow, Color.Magenta, Color.Cyan, Color.Gray
    };
    private string[] names = {
        "Чёрный", "Белый", "Красный", "Зелёный", "Синий",
        "Жёлтый", "Пурпурный", "Голубой", "Серый"
    };
    private Label infoLabel;

    public PixelCheckForm()
    {
        this.Text = "🖥️ PixelCheck Pro";
        this.WindowState = FormWindowState.Maximized;
        this.FormBorderStyle = FormBorderStyle.None;
        this.TopMost = true;
        this.KeyPreview = true;
        this.BackColor = colors[0];

        infoLabel = new Label();
        infoLabel.Text = names[0];
        infoLabel.Font = new Font("Arial", 24, FontStyle.Bold);
        infoLabel.ForeColor = Color.White;
        infoLabel.Dock = DockStyle.Bottom;
        infoLabel.TextAlign = ContentAlignment.MiddleCenter;
        this.Controls.Add(infoLabel);

        this.MouseClick += (s, e) => NextColor();
        this.KeyDown += (s, e) => {
            if (e.KeyCode == Keys.Space) NextColor();
            else if (e.KeyCode == Keys.Escape) Application.Exit();
        };
    }

    private void NextColor()
    {
        index = (index + 1) % colors.Length;
        this.BackColor = colors[index];
        infoLabel.Text = names[index];
        Color c = colors[index];
        int brightness = (c.R + c.G + c.B) / 3;
        infoLabel.ForeColor = brightness < 128 ? Color.White : Color.Black;
    }

    [STAThread]
    static void Main()
    {
        Application.EnableVisualStyles();
        Application.Run(new PixelCheckForm());
    }
}
