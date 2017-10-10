package pos

type ZeroBased int
type OneBased int

type Strand int
const (
  Unknown Strand = iota
  Forward
  Reverse
  NoStrand
)
