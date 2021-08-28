package main

import "fmt"

// Фасад предоставляет простой интерфейс для сложной логики одной или
// нескольких подсистем. Фасад делегирует запросы клиентов соответствующим
// объектам внутри подсистемы. Фасад также отвечает за управление их жизненным
// циклом. Все это защищает клиента от нежелательной сложности подсистемы.

//+ Фасад изолирует клиентов от компонентов сложной системы и дает простой интерфейс для взаимодействия
//- Фасад имеет риск стать божественным объектом, скапливая в себе слишком большой функционал

type CPU struct {
	name string
}

func newCPU(cpuName string) *CPU {
	return &CPU{
		name: cpuName,
	}
}

func (c *CPU) freeze() {
	fmt.Println("Freezing CPU")
}

func (c *CPU) jump(pos string) {
	fmt.Println("CPU jumping to", pos)
}

func (c *CPU) execute() {
	fmt.Println("Executing")
}

type Memory struct {
	name string
}

func newMemory(memName string) *Memory {
	return &Memory{
		name: memName,
	}
}

func (m *Memory) store(pos string, data string) {
	fmt.Printf("Storing %s at address %s\n", data, pos)
}

type HardDrive struct {
	name string
}

func (h *HardDrive) read(sector string, size string) string {
	fmt.Printf("Read some data from disk sector %s with size %s\n", sector, size)
	return "some_data"
}

func newHDD(hddName string) *HardDrive {
	return &HardDrive{
		name: hddName,
	}
}

type ComputerFacade struct {
	CPU       *CPU
	Memory    *Memory
	HardDrive *HardDrive
}

func (pc *ComputerFacade) Start() {
	fmt.Println("Starting PC.")

	pc.CPU = newCPU("CPU1")
	pc.Memory = newMemory("Memory1")
	pc.HardDrive = newHDD("HDD1")

	pc.CPU.freeze()
	data := pc.HardDrive.read("100", "2048")
	pc.Memory.store("0x00", data)
	pc.CPU.jump("0x00")
	pc.CPU.execute()
}

func main() {
	pc := ComputerFacade{}
	pc.Start()
}
