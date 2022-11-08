package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"debug_uart"
)

const qty_cell int = 7

// func button_save(i int, save *bool) *widget.Button {
// 	btn := widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() {
// 		*save = true
// 	})
// 	return btn
// }

func button_balancing(i int, uart *debug_uart.Uart) *widget.Button {
	btn := widget.NewButton(strconv.Itoa(i+1), func() {
		uart.Balance(byte(i + 49))
	})
	return btn
}

func GetBit(v byte, n int) bool {
	return v&(1<<n) != 0
}

func main() {

	uart := debug_uart.Make()
	// defer uart.Port.Close()
	values := make([]uint16, 10)
	var state_answer byte
	var error_answer byte

	a := app.New()
	// r, _ := fyne.LoadResourceFromPath("logo.png")
	// a.SetIcon(r)
	w := a.NewWindow("BMS")
	w.SetMaster()
	a.Settings().SetTheme(theme.DarkTheme())

	title := canvas.NewText("Сell parameters", color.NRGBA{241, 212, 130, 255})
	title.TextSize = 20
	title.Alignment = 1

	var cell [qty_cell]*widget.ProgressBar
	for i, _ := range cell {
		cell[i] = widget.NewProgressBar()
		cell[i].Min = 3.2
		cell[i].Max = 4.2
		cell[i].SetValue(3.2)
		cell[i].TextFormatter = func() string { return fmt.Sprintf("%.2f Volts", 0.0) }
	}

	var entry [qty_cell]*widget.Entry
	for i, _ := range entry {
		entry[i] = widget.NewEntry()
		entry[i].SetPlaceHolder("Correction")
	}

	var btn_save [qty_cell]*widget.Button
	var save [qty_cell]bool
	// for i, _ := range btn_save {
	// 	btn_save[i] = button_save(i, &save[i])
	// }
	btn_save[0] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[0] = true })
	btn_save[1] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[1] = true })
	btn_save[2] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[2] = true })
	btn_save[3] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[3] = true })
	btn_save[4] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[4] = true })
	btn_save[5] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[5] = true })
	btn_save[6] = widget.NewButtonWithIcon("save", theme.DocumentSaveIcon(), func() { save[6] = true })

	var balancing [qty_cell]*widget.Button
	for i, _ := range balancing {
		balancing[i] = button_balancing(i, uart)
	}

	temp_board := widget.NewLabel("Board: stop")
	temp_1 := widget.NewLabel("Sensor 1: stop")
	temp_2 := widget.NewLabel("Sensor 2: stop")

	var circle_green [8]*canvas.Circle
	for i, _ := range circle_green {
		circle_green[i] = canvas.NewCircle(color.NRGBA{147, 189, 158, 50})
		circle_green[i].Resize(fyne.NewSize(16, 16))
		circle_green[i].Move(fyne.Position{X: 10, Y: 10})
	}

	var circle_red [8]*canvas.Circle
	for i, _ := range circle_red {
		circle_red[i] = canvas.NewCircle(color.NRGBA{217, 26, 20, 50})
		circle_red[i].Resize(fyne.NewSize(16, 16))
		circle_red[i].Move(fyne.Position{X: 10, Y: 10})
	}

	state_0 := widget.NewLabel("reserved")
	state_0.Move(fyne.Position{X: 30, Y: 0})
	state_1 := widget.NewLabel("started")
	state_1.Move(fyne.Position{X: 30, Y: 0})
	state_2 := widget.NewLabel("charging")
	state_2.Move(fyne.Position{X: 30, Y: 0})
	state_3 := widget.NewLabel("test")
	state_3.Move(fyne.Position{X: 30, Y: 0})
	state_4 := widget.NewLabel("reserved")
	state_4.Move(fyne.Position{X: 30, Y: 0})
	state_5 := widget.NewLabel("balancing")
	state_5.Move(fyne.Position{X: 30, Y: 0})
	state_6 := widget.NewLabel("DC 15V")
	state_6.Move(fyne.Position{X: 30, Y: 0})
	state_7 := widget.NewLabel("reserved")
	state_7.Move(fyne.Position{X: 30, Y: 0})

	error_0 := widget.NewLabel("DC 15V")
	error_0.Move(fyne.Position{X: 30, Y: 0})
	error_1 := widget.NewLabel("high temp")
	error_1.Move(fyne.Position{X: 30, Y: 0})
	error_2 := widget.NewLabel("temp sensor")
	error_2.Move(fyne.Position{X: 30, Y: 0})
	error_3 := widget.NewLabel("low temp")
	error_3.Move(fyne.Position{X: 30, Y: 0})
	error_4 := widget.NewLabel("board temp")
	error_4.Move(fyne.Position{X: 30, Y: 0})
	error_5 := widget.NewLabel("high volatge")
	error_5.Move(fyne.Position{X: 30, Y: 0})
	error_6 := widget.NewLabel("min voltage")
	error_6.Move(fyne.Position{X: 30, Y: 0})
	error_7 := widget.NewLabel("max voltage")
	error_7.Move(fyne.Position{X: 30, Y: 0})

	state_bit_0 := container.NewWithoutLayout(circle_green[0], state_0)
	state_bit_1 := container.NewWithoutLayout(circle_green[1], state_1)
	state_bit_2 := container.NewWithoutLayout(circle_green[2], state_2)
	state_bit_3 := container.NewWithoutLayout(circle_green[3], state_3)
	state_bit_4 := container.NewWithoutLayout(circle_green[4], state_4)
	state_bit_5 := container.NewWithoutLayout(circle_green[5], state_5)
	state_bit_6 := container.NewWithoutLayout(circle_green[6], state_6)
	state_bit_7 := container.NewWithoutLayout(circle_green[7], state_7)

	error_bit_0 := container.NewWithoutLayout(circle_red[0], error_0)
	error_bit_1 := container.NewWithoutLayout(circle_red[1], error_1)
	error_bit_2 := container.NewWithoutLayout(circle_red[2], error_2)
	error_bit_3 := container.NewWithoutLayout(circle_red[3], error_3)
	error_bit_4 := container.NewWithoutLayout(circle_red[4], error_4)
	error_bit_5 := container.NewWithoutLayout(circle_red[5], error_5)
	error_bit_6 := container.NewWithoutLayout(circle_red[6], error_6)
	error_bit_7 := container.NewWithoutLayout(circle_red[7], error_7)

	label_state := canvas.NewText("States", color.NRGBA{242, 255, 0, 255})
	label_state.Alignment = 1
	label_error := canvas.NewText("Errors", color.NRGBA{242, 255, 0, 255})
	label_error.Alignment = 1
	// &layout.Spacer{},

	// temp_box := container.NewGridWithColumns(3, container.NewGridWithColumns(3, temp_board, temp_1, temp_2), container.NewGridWithColumns(3, &layout.Spacer{}, &layout.Spacer{}, container.New(layout.NewMaxLayout(), test_color, btn_test)), container.NewGridWithColumns(2, label_state, label_error))
	temp_box := container.NewGridWithColumns(3, container.NewGridWithColumns(3, temp_board, temp_1, temp_2), container.NewGridWithColumns(3, &layout.Spacer{}, &layout.Spacer{}, &layout.Spacer{}), container.NewGridWithColumns(2, label_state, label_error))
	cell_box1 := container.NewGridWithColumns(3, cell[0], container.NewGridWithColumns(3, entry[0], btn_save[0], balancing[0]), container.NewGridWithColumns(2, state_bit_0, error_bit_0))
	cell_box2 := container.NewGridWithColumns(3, cell[1], container.NewGridWithColumns(3, entry[1], btn_save[1], balancing[1]), container.NewGridWithColumns(2, state_bit_1, error_bit_1))
	cell_box3 := container.NewGridWithColumns(3, cell[2], container.NewGridWithColumns(3, entry[2], btn_save[2], balancing[2]), container.NewGridWithColumns(2, state_bit_2, error_bit_2))
	cell_box4 := container.NewGridWithColumns(3, cell[3], container.NewGridWithColumns(3, entry[3], btn_save[3], balancing[3]), container.NewGridWithColumns(2, state_bit_3, error_bit_3))
	cell_box5 := container.NewGridWithColumns(3, cell[4], container.NewGridWithColumns(3, entry[4], btn_save[4], balancing[4]), container.NewGridWithColumns(2, state_bit_4, error_bit_4))
	cell_box6 := container.NewGridWithColumns(3, cell[5], container.NewGridWithColumns(3, entry[5], btn_save[5], balancing[5]), container.NewGridWithColumns(2, state_bit_5, error_bit_5))
	cell_box7 := container.NewGridWithColumns(3, cell[6], container.NewGridWithColumns(3, entry[6], btn_save[6], balancing[6]), container.NewGridWithColumns(2, state_bit_6, error_bit_6))

	select_port := widget.NewSelect(debug_uart.GetPort(), func(string) {})
	select_port.SetSelectedIndex(1)

	var connect bool
	btn_connect := widget.NewButton("Connect", func() {
		if connect {
			uart.Stop()
			time.Sleep(10 * time.Millisecond)
			uart.Close()
			connect = false
		} else {
			err := uart.Listen(select_port.Selected)
			if err != nil {
				return
			}
			connect = true
		}
	})

	btn_start := widget.NewButton("Start", func() {

		if uart.Started {
			uart.Stop()

		} else {
			uart.Start()
		}
	})

	var test bool
	btn_test := widget.NewButton("Test", func() {
		if test {
			test = false
			uart.Test_off()
		} else {
			test = true
			uart.Test_on()
		}
	})

	test_color := canvas.NewRectangle(color.NRGBA{45, 45, 45, 255})

	var charging bool
	btn_charger := widget.NewButton("Charger", func() {
		charging = true
	})

	charger_color := canvas.NewRectangle(color.NRGBA{45, 45, 45, 255})

	btn_box := container.NewGridWithColumns(
		3,
		container.NewGridWithColumns(2, select_port, btn_connect),
		container.NewGridWithColumns(3, btn_start, container.New(layout.NewMaxLayout(), test_color, btn_test), container.New(layout.NewMaxLayout(), charger_color, btn_charger)),
		container.NewGridWithColumns(2, state_bit_7, error_bit_7),
	)

	content := container.NewVBox(
		title,
		temp_box,
		cell_box1, cell_box2, cell_box3, cell_box4, cell_box5, cell_box6, cell_box7,
		btn_box,
	)

	go func() {
		for range time.Tick(100 * time.Millisecond) {
			if uart.Started {
				btn_start.SetText("Stop")
				values = uart.GetData()
				for i, _ := range cell {
					cell[i].SetValue(float64(values[i]) / 100)
					cell[i].TextFormatter = func() string { return fmt.Sprintf("%.2f Volts", cell[i].Value) }
					cell[i].Refresh()
				}

				temp_board.SetText("Board: " + strconv.Itoa(int(values[7])-273) + "C°")
				temp_1.SetText("Sensor 1: " + strconv.Itoa(int(values[8])-273) + "C°")
				temp_2.SetText("Sensor 2: " + strconv.Itoa(int(values[9])-273) + "C°")

				state_answer = uart.GetState()
				for i, n := range circle_green {
					if GetBit(state_answer, i) {
						n.FillColor = color.NRGBA{17, 255, 0, 255}
					} else {
						n.FillColor = color.NRGBA{147, 189, 158, 50}
					}
					n.Refresh()
				}
				error_answer = uart.GetError()
				for i, n := range circle_red {
					if GetBit(error_answer, i) {
						n.FillColor = color.NRGBA{255, 3, 15, 255}
					} else {
						n.FillColor = color.NRGBA{217, 26, 20, 50}
					}
					n.Refresh()
				}

				for i, _ := range save {
					if save[i] {
						data, _ := strconv.ParseUint(entry[i].Text, 10, 16)
						uart.Correction(i, uint16(data))
						save[i] = false
					}
				}

			} else if uart.Stopped {
				for i, _ := range cell {
					cell[i].SetValue(3.2)
					cell[i].TextFormatter = func() string { return fmt.Sprintf("%.2f Volts", 0.0) }
				}
				temp_board.SetText("Board: stop")
				temp_1.SetText("Sensor 1: stop")
				temp_2.SetText("Sensor 2: stop")
				btn_start.SetText("Start")

				for i, _ := range circle_green {
					circle_green[i].FillColor = color.NRGBA{147, 189, 158, 50}
					circle_green[i].Refresh()
					circle_red[i].FillColor = color.NRGBA{217, 26, 20, 50}
					circle_red[i].Refresh()
				}
			}

			if connect {
				btn_connect.SetText("Disconnect")
			} else {
				btn_connect.SetText("Connect")
			}

			if GetBit(state_answer, 3) {
				test_color.FillColor = color.NRGBA{28, 138, 6, 255}
				test_color.Refresh()
			} else {
				test_color.FillColor = color.NRGBA{45, 45, 45, 50}
				test_color.Refresh()
			}

			if charging {
				uart.Charging()
				charging = false
			}

			if GetBit(state_answer, 2) {
				charger_color.FillColor = color.NRGBA{28, 138, 6, 255}
				charger_color.Refresh()
			} else {
				charger_color.FillColor = color.NRGBA{45, 45, 45, 50}
				charger_color.Refresh()
			}

		}
	}()

	w.SetContent(content)
	w.Resize(fyne.NewSize(600, 300))
	w.ShowAndRun()

}
