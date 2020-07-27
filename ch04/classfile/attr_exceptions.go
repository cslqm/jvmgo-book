package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionsAttribute() []uint16 {
	return self.exceptionIndexTable
}
