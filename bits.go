package bitart

type Bits []byte

func (b Bits) String() string {
	return join(
		b...,
	)
}

type Bits8 uint8

func (b Bits8) String() string {
	return join(
		uint8(b),
	)
}

type Bits16 uint16

func (b Bits16) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
	)
}

type Bits24 uint32

func (b Bits24) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
		uint8(b>>16),
	)
}

type Bits32 uint32

func (b Bits32) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
		uint8(b>>16),
		uint8(b>>24),
	)
}

type Bits40 uint64

func (b Bits40) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
		uint8(b>>16),
		uint8(b>>24),
		uint8(b>>32),
	)
}

type Bits48 uint64

func (b Bits48) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
		uint8(b>>16),
		uint8(b>>24),
		uint8(b>>32),
		uint8(b>>40),
	)
}

type Bits56 uint64

func (b Bits56) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
		uint8(b>>16),
		uint8(b>>24),
		uint8(b>>32),
		uint8(b>>40),
		uint8(b>>48),
	)
}

type Bits64 uint64

func (b Bits64) String() string {
	return join(
		uint8(b),
		uint8(b>>8),
		uint8(b>>16),
		uint8(b>>24),
		uint8(b>>32),
		uint8(b>>40),
		uint8(b>>48),
		uint8(b>>56),
	)
}
