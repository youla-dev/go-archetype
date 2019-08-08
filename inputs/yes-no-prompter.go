package inputs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pkg/errors"
)

type yesNoPrompter struct {
	PromptResponse
}

func newYesNoPrompter(spec InputSpec) *yesNoPrompter {
	return &yesNoPrompter{PromptResponse: PromptResponse{InputSpec: spec}}
}

func (p *yesNoPrompter) Prompt() (PromptResponse, error) {
	if p.Answered {
		return p.PromptResponse, nil
	}
	var (
		yes    bool
		answer string
	)
	prompt := &survey.Confirm{
		Message: p.Text,
	}
	err := survey.AskOne(prompt, &yes)
	if err != nil {
		return PromptResponse{}, errors.Wrap(err, "prompt error")
	}
	if yes {
		answer = "true"
	} else {
		answer = "false"
	}
	return p.SetStringResponse(answer), nil
}

func (p *yesNoPrompter) GetID() string {
	return p.ID
}

func (p *yesNoPrompter) SetStringResponse(answer string) PromptResponse {
	answer = p.beNiceAndTryToConvert(answer)
	b, err := strconv.ParseBool(answer)
	if err != nil {
		panic(fmt.Sprintf("Unknown input to yes/no boolean input (use true/false): %+v", err))
	}
	p.Truthy = b
	p.Answer = answer
	p.Answered = true
	return p.PromptResponse
}

// Tries to find a suitable conversion b/w the input string and a true/false value.
// If not found, just returns the original value itself
func (p *yesNoPrompter) beNiceAndTryToConvert(str string) string {
	switch strings.ToLower(str) {
	case "yes", "ok", "sure", "why not":
		return "true"
	case "no", "hell no", "as if":
		return "false"
	default:
		return str
	}
}

var _ Prompter = &yesNoPrompter{}
