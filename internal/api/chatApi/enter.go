package chatApi

type AiModel uint8

const (
	Tongyi AiModel = iota
	ChatGpt35Turbo
)

type ChatApi struct {
}
