package main

//Команда позволяет заворачивать запросы или простые операции в отдельные объекты.

//+
//1) Удаление прямой зависимости между объектами, вызывающими операции и между их исполнителями.
//2) Команда позволяет реализовать простую отмену и повтор операций.
//3) Команда позволяет реализовать отложенный запуск операций.
//4) Команда позволяет собирать сложные команды из простых.

//-
//Усложнение кода программы

func main() {
	app := &app{"Important Data"}

	saveCommand := &saveCommand{
		invoker: app,
	}

	saveButton := &button{
		command: saveCommand,
	}

	hotkeyCombo := &hotkey{
		command: saveCommand,
	}

	saveButton.press()
	hotkeyCombo.pressHotkey()
}
