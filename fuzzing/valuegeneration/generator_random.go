package valuegeneration

import (
	"math/big"
	"math/rand"

	"github.com/crytic/medusa-geth/common"
	"github.com/crytic/medusa/utils"
)

// RandomValueGenerator represents a ValueGenerator used to generate transaction fields and call arguments with values
// provided by a random number generator.
type RandomValueGenerator struct {
	// config describes the configuration defining value generation parameters.
	config *RandomValueGeneratorConfig

	// randomProvider offers a source of random data.
	randomProvider *rand.Rand
}

// RandomValueGeneratorConfig defines the parameters for a RandomValueGenerator.
type RandomValueGeneratorConfig struct {
	// GenerateRandomArrayMinSize defines the minimum size which a generated array should be.
	GenerateRandomArrayMinSize int
	// GenerateRandomArrayMaxSize defines the maximum size which a generated array should be.
	GenerateRandomArrayMaxSize int
	// GenerateRandomBytesMinSize defines the minimum size which a generated byte slice should be.
	GenerateRandomBytesMinSize int
	// GenerateRandomBytesMaxSize defines the maximum size which a generated byte slice should be.
	GenerateRandomBytesMaxSize int
	// GenerateRandomStringMinSize defines the minimum size which a generated string should be.
	GenerateRandomStringMinSize int
	// GenerateRandomStringMaxSize defines the maximum size which a generated string should be.
	GenerateRandomStringMaxSize int
}

// NewRandomValueGenerator creates a new RandomValueGenerator.
func NewRandomValueGenerator(config *RandomValueGeneratorConfig, randomProvider *rand.Rand) *RandomValueGenerator {
	// Create and return our generator
	generator := &RandomValueGenerator{
		config:         config,
		randomProvider: randomProvider,
	}
	return generator
}

// GenerateAddress generates a random address to use when populating inputs.
func (g *RandomValueGenerator) GenerateAddress() common.Address {
	// Generate random bytes of the address length, then convert it to an address.
	addressBytes := make([]byte, common.AddressLength)
	g.randomProvider.Read(addressBytes)
	return common.BytesToAddress(addressBytes)
}

// MutateAddress takes an address input and returns a mutated value based off the input.
func (g *RandomValueGenerator) MutateAddress(addr common.Address) common.Address {
	// This value generator does not apply mutations.
	return addr
}

// GenerateArrayOfLength generates a random array length to use when populating inputs. This is used to determine how
// many elements a non-byte, non-string array should have.
func (g *RandomValueGenerator) GenerateArrayOfLength() int {
	rangeSize := uint64(g.config.GenerateRandomArrayMaxSize-g.config.GenerateRandomArrayMinSize) + 1
	return int(g.GenerateInteger(false, 16).Uint64()%rangeSize) + g.config.GenerateRandomArrayMinSize
}

// MutateArray takes a dynamic or fixed sized array as input, and returns a mutated value based off of the input.
// Returns the mutated value. If any element of the returned array is nil, the value generator will be called upon
// to generate it new.
func (g *RandomValueGenerator) MutateArray(value []any, fixedLength bool) []any {
	// This value generator does not apply mutations.
	return value
}

// GenerateBool generates a random bool to use when populating inputs.
func (g *RandomValueGenerator) GenerateBool() bool {
	return g.randomProvider.Uint32()%2 == 0
}

// MutateBool takes a boolean input and returns a mutated value based off the input.
func (g *RandomValueGenerator) MutateBool(bl bool) bool {
	// This value generator does not apply mutations.
	return bl
}

// GenerateBytes generates a random dynamic-sized byte array to use when populating inputs.
func (g *RandomValueGenerator) GenerateBytes() []byte {
	rangeSize := uint64(g.config.GenerateRandomBytesMaxSize-g.config.GenerateRandomBytesMinSize) + 1
	b := make([]byte, int(g.randomProvider.Uint64()%rangeSize)+g.config.GenerateRandomBytesMinSize)
	g.randomProvider.Read(b)
	return b
}

// MutateBytes takes a dynamic-sized byte array input and returns a mutated value based off the input.
func (g *RandomValueGenerator) MutateBytes(b []byte) []byte {
	// This value generator does not apply mutations.
	return b
}

// GenerateFixedBytes generates a random fixed-sized byte array to use when populating inputs.
func (g *RandomValueGenerator) GenerateFixedBytes(length int) []byte {
	b := make([]byte, length)
	g.randomProvider.Read(b)
	return b
}

// MutateFixedBytes takes a fixed-sized byte array input and returns a mutated value based off the input.
func (g *RandomValueGenerator) MutateFixedBytes(b []byte) []byte {
	// This value generator does not apply mutations.
	return b
}

// GenerateString generates a random dynamic-sized string to use when populating inputs.
func (g *RandomValueGenerator) GenerateString() string {
	rangeSize := uint64(g.config.GenerateRandomStringMaxSize-g.config.GenerateRandomStringMinSize) + 1
	b := make([]byte, int(g.randomProvider.Uint64()%rangeSize)+g.config.GenerateRandomStringMinSize)
	g.randomProvider.Read(b)
	return string(b)
}

// MutateString takes a string input and returns a mutated value based off the input.
func (g *RandomValueGenerator) MutateString(s string) string {
	// This value generator does not apply mutations.
	return s
}

// GenerateInteger generates a random integer to use when populating inputs.
func (g *RandomValueGenerator) GenerateInteger(signed bool, bitLength int) *big.Int {
	// Fill a byte array of the appropriate size with random bytes
	b := make([]byte, bitLength/8)
	g.randomProvider.Read(b)

	// Create an unsigned integer.
	res := big.NewInt(0).SetBytes(b)

	// Constrain our integer bounds
	return utils.ConstrainIntegerToBitLength(res, signed, bitLength)
}

// MutateInteger takes an integer input and returns a mutated value based off the input.
func (g *RandomValueGenerator) MutateInteger(i *big.Int, signed bool, bitLength int) *big.Int {
	// This value generator does not apply mutations.
	return i
}
