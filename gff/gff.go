package gff

import (
  "bufio"
  "fmt"
  "io"
  "strconv"
  "strings"
  "github.com/buchanae/pop/pos"
)

// Best known spec: https://github.com/The-Sequence-Ontology/Specifications/blob/master/gff3.md

type Phase int
const (
  PhaseUnknown Phase = iota
  PhaseZero
  PhaseOne
  PhaseTwo
)

type Directive string
type Comment string

type Record struct {
  SeqID string
  Source string
  Type string
  Start pos.OneBased
  End pos.OneBased
  Score float64
  Strand pos.Strand
  Phase Phase
  RawAttributes string
  Attributes Attributes
}

type Attributes map[string]string

// TODO add IsCircular and Target
func (a Attributes) Parents() []string {
  return strings.Split(a["Parent"], ",")
}
func (a Attributes) Aliases() []string {
  return strings.Split(a["Alias"], ",")
}
func (a Attributes) Notes() []string {
  return strings.Split(a["Note"], ",")
}
func (a Attributes) Dbxrefs() []string {
  return strings.Split(a["Dbxref"], ",")
}
func (a Attributes) OntologyTerms() []string {
  return strings.Split(a["Ontology_term"], ",")
}

type Scanner struct {
  scan *bufio.Scanner
  line Line
  lineno int
  err error
  SkipComments bool
  SkipDirectives bool
  SkipInvalid bool
}

func NewScanner(r io.Reader) *Scanner {
  s := bufio.NewScanner(r)
  s.Split(bufio.ScanLines)
  return &Scanner{scan: s}
}

func (s *Scanner) Scan() bool {
  if s.err != nil {
    return false
  }
  b := s.scan.Scan()
  if !b {
    return false
  }

  s.lineno++
  t := s.scan.Text()

  // Skip blank lines
  t = strings.TrimSpace(t)
  if len(t) == 0 {
    return s.Scan()
  }

  // Directive
  if t[:2] == "##" {
    if s.SkipDirectives {
      return s.Scan()
    }
    s.line = Directive(t[2:])
    return true
  }

  // Comment
  if t[0] == '#' {
    if s.SkipComments {
      return s.Scan()
    }
    s.line = Comment(t[1:])
    return true
  }

  col := strings.Split(t, "\t")

  if len(col) != 9 && !s.SkipInvalid {
    s.err = fmt.Errorf("line %d: len(columns) != 9", s.lineno)
    return false
  }

  seqid := col[0]
  source := col[1]
  type_ := col[2]
  start := col[3]
  end := col[4]
  score := col[5]
  strand := col[6]
  phase := col[7]
  attrib := col[8]

  rec := Record{
    SeqID: seqid,
    Source: source,
    Type: type_,
    RawAttributes: attrib,
  }

  if start != "." {
    p, err := strconv.Atoi(start)
    if err != nil {
      s.err = fmt.Errorf("line %d: can't parse 'start' column, %s", s.lineno, err)
      return false
    }
    rec.Start = pos.OneBased(p)
  }

  if end != "." {
    p, err := strconv.Atoi(end)
    if err != nil {
      s.err = fmt.Errorf("line %d: can't parse 'end' column, %s", s.lineno, err)
      return false
    }
    rec.End = pos.OneBased(p)
  }

  if score != "." {
    p, err := strconv.ParseFloat(score, 64)
    if err != nil {
      s.err = fmt.Errorf("line %d: can't parse 'score' column, %s", s.lineno, err)
      return false
    }
    rec.Score = p
  }

  switch strand {
  case ".":
    rec.Strand = pos.NoStrand
  case "?":
    rec.Strand = pos.Unknown
  case "+":
    rec.Strand = pos.Forward
  case "-":
    rec.Strand = pos.Reverse
  default:
    s.err = fmt.Errorf("line %d: unrecognized strand '%s'", s.lineno, strand)
    return false
  }

  switch phase {
  case ".":
    rec.Phase = PhaseUnknown
  case "0":
    rec.Phase = PhaseZero
  case "1":
    rec.Phase = PhaseOne
  case "2":
    rec.Phase = PhaseTwo
  default:
    s.err = fmt.Errorf("line %d: unrecognized phase '%s'", s.lineno, phase)
    return false
  }

  for _, a := range strings.Split(attrib, ";") {
    asp := strings.Split(a, "=")
    if len(asp) != 2 {
      s.err = fmt.Errorf("line %d: can't split attribute %s", s.lineno, a)
      return false
    }
    rec.Attributes[asp[0]] = asp[1]
  }

  s.line = &rec
  return true
}

func (s *Scanner) Err() error {
  if s.err != nil {
    return s.err
  }
  return s.scan.Err()
}

func (s *Scanner) Line() Line {
  return s.line
}


// Line defines the interface of all types
// that may be returned by this reader:
// Record, Directive, and Comment.
type Line interface {
  gffline()
}

// Implementations of type Line
func (Directive) gffline() {}
func (Comment) gffline() {}
func (*Record) gffline() {}
