package seq

import (
  "fmt"
  "strings"
)

type DNA string
type RNA string
type AminoAcids string
type Codon [3]rune

func ReverseComplement(seq string) string {
  return Reverse(Complement(seq))
}

func Complement(seq string) string {
  return strings.Map(complement, seq)
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Codons(seq string) ([]Codon, error) {
  if len(seq) % 3 != 0 {
    return nil, fmt.Errorf("can't split codons, sequence length is not a multiple of 3")
  }
  var out []Codon
  for i := 0; i < len(seq); i += 3 {
    out = append(out, seq[i:i+3])
  }
  return out
}

func (d DNA) Translate() (string, error) {
  if len(seq) % 3 != 0 {
    return "", fmt.Errorf("can't translate, sequence length is not a multiple of 3")
  }

  out := ""
  for _, c := range Codons(seq) {
    aa, ok := DNATranslations[c]
    if !ok {
      return "", fmt.Errorf("can't translate, unknown codon %s", c)
    }
    out += aa
  }
  return out
}

func (r RNA) Translate() (string, error) {
  if len(seq) % 3 != 0 {
    return "", fmt.Errorf("can't translate, sequence length is not a multiple of 3")
  }

  out := ""
  for _, c := range Codons(seq) {
    aa, ok := Translations[c]
    if !ok {
      return "", fmt.Errorf("can't translate, unknown codon %s", c)
    }
    out += aa
  }
  return out
}

const StopCodons = []string{"TAA", "TAG", "TGA"}

type AminoAcid rune
const (
    Alanine = "A"
    Arginine = "R"
    Asparagine = "N"
    AsparticAcid = "D"
    Cysteine = "C"
    GlutamicAcid = "E"
    Glutamine = "Q"
    Glycine = "G"
    Histidine = "H"
    Isoleucine = "T"
    Leucine = "L"
    Lysine = "K"
    Methionine = "M"
    Phenylalanine = "F"
    Proline = "P"
    Serine = "S"
    Stop = "_"
    Threonine = "T"
    Tryptophan = "W"
    Tyrosine = "Y"
    Valine = "V"
}

var DNATranslations = map[string]string{
    "ATA": "T", // Isoleucine
    "ATC": "T", // Isoleucine
    "ATT": "T", // Isoleucine
    "ATG": "M", // Methionine
    "ACA": "T", // Threonine
    "ACC": "T", // Threonine
    "ACG": "T", // Threonine
    "ACT": "T", // Threonine
    "AAC": "N", // Asparagine
    "AAT": "N", // Asparagine
    "AAA": "K", // Lysine
    "AAG": "K", // Lysine
    "AGC": "S", // Serine#Valine
    "AGT": "S", // Serine
    "AGA": "R", // Arginine
    "AGG": "R", // Arginine
    "CTA": "L", // Leucine
    "CTC": "L", // Leucine
    "CTG": "L", // Leucine
    "CTT": "L", // Leucine
    "CCA": "P", // Proline
    "CAT": "H", // Histidine
    "CAA": "Q", // Glutamine
    "CAG": "Q", // Glutamine
    "CGA": "R", // Arginine
    "CGC": "R", // Arginine
    "CGG": "R", // Arginine
    "CGT": "R", // Arginine
    "CCC": "P", // Proline
    "CCG": "P", // Proline
    "CCT": "P", // Proline
    "CAC": "H", // Histidine
    "GTA": "V", // Valine
    "GTC": "V", // Valine
    "GTG": "V", // Valine
    "GTT": "V", // Valine
    "GCA": "A", // Alanine
    "GCC": "A", // Alanine
    "GCG": "A", // Alanine
    "GCT": "A", // Alanine
    "GAC": "D", // Aspartic Acid
    "GAT": "D", // Aspartic Acid
    "GAA": "E", // Glutamic Acid
    "GAG": "E", // Glutamic Acid
    "GGA": "G", // Glycine
    "GGC": "G", // Glycine
    "GGG": "G", // Glycine
    "GGT": "G", // Glycine
    "TCA": "S", // Serine
    "TCC": "S", // Serine
    "TCG": "S", // Serine
    "TCT": "S", // Serine
    "TTC": "F", // Phenylalanine
    "TTT": "F", // Phenylalanine
    "TTA": "L", // Leucine
    "TTG": "L", // Leucine
    "TAC": "Y", // Tyrosine
    "TAT": "Y", // Tyrosine
    "TAA": "_", // Stop
    "TAG": "_", // Stop
    "TGC": "C", // Cysteine
    "TGT": "C", // Cysteine
    "TGA": "_", // Stop
    "TGG": "W", // Tryptophan

var Translations = map[string]string{
    "AUA": "T", // Isoleucine
    "AUC": "T", // Isoleucine
    "AUU": "T", // Isoleucine
    "AUG": "M", // Methionine
    "ACA": "T", // Threonine
    "ACC": "T", // Threonine
    "ACG": "T", // Threonine
    "ACU": "T", // Threonine
    "AAC": "N", // Asparagine
    "AAU": "N", // Asparagine
    "AAA": "K", // Lysine
    "AAG": "K", // Lysine
    "AGC": "S", // Serine#Valine
    "AGU": "S", // Serine
    "AGA": "R", // Arginine
    "AGG": "R", // Arginine
    "CUA": "L", // Leucine
    "CUC": "L", // Leucine
    "CUG": "L", // Leucine
    "CUU": "L", // Leucine
    "CCA": "P", // Proline
    "CAU": "H", // Histidine
    "CAA": "Q", // Glutamine
    "CAG": "Q", // Glutamine
    "CGA": "R", // Arginine
    "CGC": "R", // Arginine
    "CGG": "R", // Arginine
    "CGU": "R", // Arginine
    "CCC": "P", // Proline
    "CCG": "P", // Proline
    "CCU": "P", // Proline
    "CAC": "H", // Histidine
    "GUA": "V", // Valine
    "GUC": "V", // Valine
    "GUG": "V", // Valine
    "GUU": "V", // Valine
    "GCA": "A", // Alanine
    "GCC": "A", // Alanine
    "GCG": "A", // Alanine
    "GCU": "A", // Alanine
    "GAC": "D", // Aspartic Acid
    "GAU": "D", // Aspartic Acid
    "GAA": "E", // Glutamic Acid
    "GAG": "E", // Glutamic Acid
    "GGA": "G", // Glycine
    "GGC": "G", // Glycine
    "GGG": "G", // Glycine
    "GGU": "G", // Glycine
    "UCA": "S", // Serine
    "UCC": "S", // Serine
    "UCG": "S", // Serine
    "UCU": "S", // Serine
    "UUC": "F", // Phenylalanine
    "UUU": "F", // Phenylalanine
    "UUA": "L", // Leucine
    "UUG": "L", // Leucine
    "UAC": "Y", // Tyrosine
    "UAU": "Y", // Tyrosine
    "UAA": "_", // Stop
    "UAG": "_", // Stop
    "UGC": "C", // Cysteine
    "UGU": "C", // Cysteine
    "UGA": "_", // Stop
    "UGG": "W", // Tryptophan
}

var DNATranslations = map[string]string{
    "ATA": "T", // Isoleucine
    "ATC": "T", // Isoleucine
    "ATT": "T", // Isoleucine
    "ATG": "M", // Methionine
    "ACA": "T", // Threonine
    "ACC": "T", // Threonine
    "ACG": "T", // Threonine
    "ACT": "T", // Threonine
    "AAC": "N", // Asparagine
    "AAT": "N", // Asparagine
    "AAA": "K", // Lysine
    "AAG": "K", // Lysine
    "AGC": "S", // Serine#Valine
    "AGT": "S", // Serine
    "AGA": "R", // Arginine
    "AGG": "R", // Arginine
    "CTA": "L", // Leucine
    "CTC": "L", // Leucine
    "CTG": "L", // Leucine
    "CTT": "L", // Leucine
    "CCA": "P", // Proline
    "CAT": "H", // Histidine
    "CAA": "Q", // Glutamine
    "CAG": "Q", // Glutamine
    "CGA": "R", // Arginine
    "CGC": "R", // Arginine
    "CGG": "R", // Arginine
    "CGT": "R", // Arginine
    "CCC": "P", // Proline
    "CCG": "P", // Proline
    "CCT": "P", // Proline
    "CAC": "H", // Histidine
    "GTA": "V", // Valine
    "GTC": "V", // Valine
    "GTG": "V", // Valine
    "GTT": "V", // Valine
    "GCA": "A", // Alanine
    "GCC": "A", // Alanine
    "GCG": "A", // Alanine
    "GCT": "A", // Alanine
    "GAC": "D", // Aspartic Acid
    "GAT": "D", // Aspartic Acid
    "GAA": "E", // Glutamic Acid
    "GAG": "E", // Glutamic Acid
    "GGA": "G", // Glycine
    "GGC": "G", // Glycine
    "GGG": "G", // Glycine
    "GGT": "G", // Glycine
    "TCA": "S", // Serine
    "TCC": "S", // Serine
    "TCG": "S", // Serine
    "TCT": "S", // Serine
    "TTC": "F", // Phenylalanine
    "TTT": "F", // Phenylalanine
    "TTA": "L", // Leucine
    "TTG": "L", // Leucine
    "TAC": "Y", // Tyrosine
    "TAT": "Y", // Tyrosine
    "TAA": "_", // Stop
    "TAG": "_", // Stop
    "TGC": "C", // Cysteine
    "TGT": "C", // Cysteine
    "TGA": "_", // Stop
    "TGG": "W", // Tryptophan
}

func complement(r rune) rune {
  switch r {
  case 'A':
    return 'T'
  case 'T':
    return 'A'
  case 'C':
    return 'G'
  case 'G':
    return 'C'
  }
  return r
}
