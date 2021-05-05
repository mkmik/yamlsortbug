package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

func genBad() (string, error) {
	m := map[string]string{
		"11": "",
		"1z": "",
		"2z": "",
	}

	d, err := yaml.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func genOk() (string, error) {
	// almost any sequence of strings preserves the canonical ordering
	// except some special patterns as the one in genBad.
	// This is an example of a minimal deviation from the bad pattern.
	// I've run this for hours and it didn't find a single flip.
	m := map[string]string{
		"11": "",
		"1z": "",
		"0z": "",
	}

	d, err := yaml.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func probe(gen func() (string, error)) error {
	orig, err := gen()
	if err != nil {
		return err
	}

	i := 0
	for ; i < 100000; i++ {
		s, err := gen()
		if err != nil {
			return err
		}

		if s != orig {
			return fmt.Errorf("found different ordering after %d iterations:\n%s\nvs:\n\n%s\n", i, orig, s)
		}
	}
	fmt.Printf("Found no sort order differences after %d iterations\n", i)
	return nil
}

func mainE() error {
	fmt.Println("Probing good pattern:")
	if err := probe(genOk); err != nil {
		return err
	}
	fmt.Println("\nProbing bad pattern:")
	if err := probe(genBad); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := mainE(); err != nil {
		log.Fatal(err)
	}
}