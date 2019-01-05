package SixteenthDay

type instruction [4]int
type machine [4]int

/*
 0 => gtri
 1 => bani
 2 => eqrr
 3 => gtir

 4 => eqir
 5 => bori
 6 => seti
 7 => setr

 8 => addr
 9 => borr
10 => muli
11 => banr

12 => addi
13 => eqri
14 => mulr
15 => gtrr
*/

func (ins instruction) opcode() int {
	return ins[0]
}

func (ins instruction) registerA(m *machine) int {
	return m[ins[1]]
}

func (ins instruction) valueA() int {
	return ins[1]
}

func (ins instruction) registerB(m *machine) int {
	return m[ins[2]]
}

func (ins instruction) valueB() int {
	return ins[2]
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
	m[ins[3]] = value
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
