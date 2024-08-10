package entity

type StringOfficeIDStruct struct {
	// Fail
	OfficeID   string
	DriverName string `security:"personal_data"`
}

type UnmatchOfficeIDStruct struct {
	// Fail
	OfficeI    uint
	DriverName string `security:"personal_data"`
}

type UintOfficeIDStruct struct {
	OfficeID   uint
	DriverName string `security:"personal_data"`
}

type Uint32OfficeIDStruct struct {
	OfficeID   uint32
	DriverName string `security:"personal_data"`
}

type Uint64OfficeIDStruct struct {
	OfficeID   uint64
	DriverName string `security:"personal_data"`
}

type IntOfficeIDStruct struct {
	OfficeID   int
	DriverName string `security:"personal_data"`
}

type Int32OfficeIDStruct struct {
	OfficeID   int32
	DriverName string `security:"personal_data"`
}

type Int64OfficeIDStruct struct {
	OfficeID   int64
	DriverName string `security:"personal_data"`
}
