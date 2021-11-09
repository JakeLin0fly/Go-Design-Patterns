package Builder_Pattern

// 产品
type Computer struct {
	mainBoard string
	cpu string
	ram string
	disk string
	display string
}

func (c *Computer) setMainBoard(mainBoard string) {
	c.mainBoard = mainBoard
}
func (c *Computer) setCPU(cpu string) {
	c.cpu = cpu
}
func (c *Computer) setRam(ram string) {
	c.ram = ram
}
func (c *Computer) setDisk(disk string) {
	c.disk = disk
}
func (c *Computer) setDisplay(display string) {
	c.display = display
}

// 抽象建造者
type Builder interface {
	createMainBoard(mainBoard string)
	createCPU(cpu string)
	createRam(ram string)
	createDisk(disk string)
	createDisplay(display string)

	build() Computer
}

// 具体建造者
type ComputerBuilder struct {
	computer Computer
}

func (c *ComputerBuilder) createMainBoard(mainBoard string) {
	c.computer.setMainBoard(mainBoard)
}
func (c *ComputerBuilder) createCPU(cpu string) {
	c.computer.setCPU(cpu)
}
func (c *ComputerBuilder) createRam(ram string) {
	c.computer.setRam(ram)
}
func (c *ComputerBuilder) createDisk(disk string) {
	c.computer.setDisk(disk)
}
func (c *ComputerBuilder) createDisplay(display string) {
	c.computer.setDisplay(display)
}
func (c *ComputerBuilder) build() Computer {
	return c.computer
}

// 指挥者
type Director struct {
	builder Builder
}

func (d *Director) setBuilder(builder Builder)  {
	d.builder = builder
}

func (d *Director) createComputer() Computer {
	d.builder.createMainBoard("mainBoard")
	d.builder.createCPU("cpu")
	d.builder.createRam("ram")
	d.builder.createDisk("disk")
	d.builder.createDisplay("display")

	return d.builder.build()
}