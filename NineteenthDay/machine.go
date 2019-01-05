package NineteenthDay

import (
	"fmt"
)

type instruction [4]int
type machine struct {
	register [6]int
	ipIndex  int
}

func instructionToOpcode(ins string) int {
	switch ins {
	case "gtri":
		return 0
	case "bani":
		return 1
	case "eqrr":
		return 2
	case "gtir":
		return 3
	case "eqir":
		return 4
	case "bori":
		return 5
	case "seti":
		return 6
	case "setr":
		return 7
	case "addr":
		return 8
	case "borr":
		return 9
	case "muli":
		return 10
	case "banr":
		return 11
	case "addi":
		return 12
	case "eqri":
		return 14
	case "mulr":
		return 14
	case "gtrr":
		return 15
	default:
		fmt.Println(ins)
		panic("wtf")
	}
}

func (ins instruction) opcode() int {
	return ins[0]
}

func (ins instruction) registerA(m *machine) int {
	return m.register[ins[1]]
}

func (ins instruction) valueA() int {
	return ins[1]
}

func (ins instruction) registerB(m *machine) int {
	return m.register[ins[2]]
}

func (ins instruction) valueB() int {
	return ins[2]
}

func (m *machine) bindIp(index int) {
	m.ipIndex = index
}

func (m *machine) execute(instructions []instruction) {
	for m.ip() >= 0 && m.ip() < len(instructions) {
		m.executeInstruction(instructions[m.ip()])
		m.incrementIp()
	}
}

func (m *machine) ip() int {
	return m.register[m.ipIndex]
}

func (m *machine) incrementIp() {
	m.register[m.ipIndex]++
}

func (m *machine) executeInstruction(ins instruction) {
	switch ins.opcode() {
	case 0:
		m.gtri(ins)
	case 1:
		m.bani(ins)
	case 2:
		m.eqrr(ins)
	case 3:
		m.gtir(ins)
	case 4:
		m.eqir(ins)
	case 5:
		m.bori(ins)
	case 6:
		m.seti(ins)
	case 7:
		m.setr(ins)
	case 8:
		m.addr(ins)
	case 9:
		m.borr(ins)
	case 10:
		m.muli(ins)
	case 11:
		m.banr(ins)
	case 12:
		m.addi(ins)
	case 13:
		m.eqri(ins)
	case 14:
		m.mulr(ins)
	case 15:
		m.gtrr(ins)
	}
}

func (m *machine) write(ins instruction, value int) {
	m.register[ins[3]] = value
}

func (m *machine) addr(ins instruction) {
	m.write(ins, ins.registerA(m)+ins.registerB(m))
}

func (m *machine) mulr(ins instruction) {
	m.write(ins, ins.registerA(m)*ins.registerB(m))
}

func (m *machine) banr(ins instruction) {
	m.write(ins, ins.registerA(m)&ins.registerB(m))
}

func (m *machine) borr(ins instruction) {
	m.write(ins, ins.registerA(m)|ins.registerB(m))
}

/*--------------------------------------------------------*/

func (m *machine) addi(ins instruction) {
	m.write(ins, ins.registerA(m)+ins.valueB())
}
func (m *machine) muli(ins instruction) {
	m.write(ins, ins.registerA(m)*ins.valueB())
}

func (m *machine) bani(ins instruction) {
	m.write(ins, ins.registerA(m)&ins.valueB())
}

func (m *machine) bori(ins instruction) {
	m.write(ins, ins.registerA(m)|ins.valueB())
}

/*--------------------------------------------------------*/

func (m *machine) setr(ins instruction) {
	m.write(ins, ins.registerA(m))
}

func (m *machine) seti(ins instruction) {
	m.write(ins, ins.valueA())
}

/*--------------------------------------------------------*/

func (m *machine) gtir(ins instruction) {
	if ins.valueA() > ins.registerB(m) {
		m.write(ins, 1)
	} else {
		m.write(ins, 0)
	}
}

func (m *machine) gtri(ins instruction) {
	if ins.registerA(m) > ins.valueB() {
		m.write(ins, 1)
	} else {
		m.write(ins, 0)
	}
}

func (m *machine) gtrr(ins instruction) {
	if ins.registerA(m) > ins.registerB(m) {
		m.write(ins, 1)
	} else {
		m.write(ins, 0)
	}
}

/*--------------------------------------------------------*/

func (m *machine) eqir(ins instruction) {
	if ins.valueA() == ins.registerB(m) {
		m.write(ins, 1)
	} else {
		m.write(ins, 0)
	}
}

func (m *machine) eqri(ins instruction) {
	if ins.registerA(m) == ins.valueB() {
		m.write(ins, 1)
	} else {
		m.write(ins, 0)
	}
}

func (m *machine) eqrr(ins instruction) {
	if ins.registerA(m) == ins.registerB(m) {
		m.write(ins, 1)
	} else {
		m.write(ins, 0)
	}
}
