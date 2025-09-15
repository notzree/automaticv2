package types

type Owner string

const (
	OwnerUser   Owner = "user"
	OwnerLLM    Owner = "llm"
	OwnerSystem Owner = "system"
)
