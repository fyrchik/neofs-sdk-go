package transformer

type EpochSource interface {
	CurrentEpoch() uint64
}
