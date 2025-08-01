package config

import "fmt"

// BlockType type of block
type BlockType string

// BlockAlignment aligment of a Block
type BlockAlignment string

// Overflow defines how to handle a right block that overflows with the previous block
type Overflow string

const (
	// Prompt writes one or more Segments
	Prompt BlockType = "prompt"
	// RPrompt is a right aligned prompt
	RPrompt BlockType = "rprompt"
	// Left aligns left
	Left BlockAlignment = "left"
	// Right aligns right
	Right BlockAlignment = "right"
	// Break adds a line break
	Break Overflow = "break"
	// Hide hides the block
	Hide Overflow = "hide"
)

// Block defines a part of the prompt with optional segments
type Block struct {
	Type            BlockType      `json:"type,omitempty" toml:"type,omitempty" yaml:"type,omitempty"`
	Alignment       BlockAlignment `json:"alignment,omitempty" toml:"alignment,omitempty" yaml:"alignment,omitempty"`
	Filler          string         `json:"filler,omitempty" toml:"filler,omitempty" yaml:"filler,omitempty"`
	Overflow        Overflow       `json:"overflow,omitempty" toml:"overflow,omitempty" yaml:"overflow,omitempty"`
	LeadingDiamond  string         `json:"leading_diamond,omitempty" toml:"leading_diamond,omitempty" yaml:"leading_diamond,omitempty"`
	TrailingDiamond string         `json:"trailing_diamond,omitempty" toml:"trailing_diamond,omitempty" yaml:"trailing_diamond,omitempty"`
	Segments        []*Segment     `json:"segments,omitempty" toml:"segments,omitempty" yaml:"segments,omitempty"`
	MaxWidth        int            `json:"max_width,omitempty" toml:"max_width,omitempty" yaml:"max_width,omitempty"`
	MinWidth        int            `json:"min_width,omitempty" toml:"min_width,omitempty" yaml:"min_width,omitempty"`
	Newline         bool           `json:"newline,omitempty" toml:"newline,omitempty" yaml:"newline,omitempty"`
	Force           bool           `json:"force,omitempty" toml:"force,omitempty" yaml:"force,omitempty"`
	Index           int            `json:"index,omitempty" toml:"index,omitempty" yaml:"index,omitempty"`
}

func (b *Block) key() any {
	if b.Index > 0 {
		return b.Index - 1
	}

	return fmt.Sprintf("%s-%s", b.Type, b.Alignment)
}
