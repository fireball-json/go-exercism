package proteintranslation

import (
	"errors"
)

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop codon")
)

func FromRNA(rna string) ([]string, error) {
	var res []string
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			return nil, ErrInvalidBase
		}
		name, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				break
			}
			return nil, err
		}
		res = append(res, name)
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	codonMap := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	name, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if name == "STOP" {
		return "", ErrStop
	}
	return name, nil
}
