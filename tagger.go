package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type tokens []struct {
	Surface    string `json:"surface"`
	Reading    string `json:"reading"`
	Lem        string `json:"lem"`
	Pos        string `json:"pos"`
	Difficulty string `json:"difficulty"`
}

func tagInput(inputText string) (*tokens, error) {
	cmd := exec.Command("docker", "run", "neilhemm9:nihongo-tokenizer",
		"--pickle_name", "jlpt_all.pkl",
		"--inputText", inputText)
	japanese_tokens, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running command: %w", err)
	}
	fmt.Println(string(japanese_tokens))

	var toks tokens
	err = json.Unmarshal(japanese_tokens, &toks)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal tokens: %w", err)
	}
	return &toks, nil

}
