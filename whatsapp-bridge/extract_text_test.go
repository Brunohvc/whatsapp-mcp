package main

import (
	"testing"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

// Promo groups post each deal as a photo with the link in the caption; the caption must be stored as the content.
func TestExtractTextContentKeepsMediaCaptions(t *testing.T) {
	casos := map[string]*waProto.Message{
		"imagem":    {ImageMessage: &waProto.ImageMessage{Caption: proto.String("Box por R$ 99 https://amzn.to/x")}},
		"vídeo":     {VideoMessage: &waProto.VideoMessage{Caption: proto.String("Box por R$ 99 https://amzn.to/x")}},
		"documento": {DocumentMessage: &waProto.DocumentMessage{Caption: proto.String("Box por R$ 99 https://amzn.to/x")}},
	}
	for nome, msg := range casos {
		if got := extractTextContent(msg); got != "Box por R$ 99 https://amzn.to/x" {
			t.Errorf("%s: legenda perdida, veio %q", nome, got)
		}
	}
	if got := extractTextContent(&waProto.Message{Conversation: proto.String("oi")}); got != "oi" {
		t.Errorf("texto simples mudou: %q", got)
	}
}
