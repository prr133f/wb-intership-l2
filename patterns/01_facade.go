package patterns

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("Freezing CPU")
}

func (c *CPU) Jump(pos int) {
	fmt.Printf("Jumpong on %d address\n", pos)
}

func (c *CPU) Execute() {
	fmt.Println("Executing task on CPU")
}

type Memory struct{}

func (m *Memory) Load(pos int, data string) {
	fmt.Printf("Loading data (%s) on %d address\n", data, pos)
}

type HardDrive struct{}

func (h *HardDrive) Read(pos int, size int) string {
	fmt.Printf("Reading data (sizeof=%d) on %d address\n", size, pos)

	return "Hello"
}

type ComputerFacade struct {
	c *CPU
	m *Memory
	h *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		c: &CPU{},
		m: &Memory{},
		h: &HardDrive{},
	}
}

func (cf *ComputerFacade) Start() {
	cf.c.Freeze()
	cf.m.Load(1, cf.h.Read(1, 5))
	cf.c.Jump(1)
	cf.c.Execute()
}
