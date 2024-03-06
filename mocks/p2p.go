package mocks

import (
	"encoding/json"

	"github.com/0xsequence/bundler/p2p"
	"github.com/0xsequence/bundler/proto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/mock"
)

type MockP2p struct {
	mock.Mock

	Broadcasted []proto.Message
	Handlers    map[proto.MessageType][]p2p.MsgHandler
}

func (p *MockP2p) Broadcast(payload proto.Message) error {
	p.Broadcasted = append(p.Broadcasted, payload)
	return p.Called(payload).Error(0)
}

func (p *MockP2p) HandleMessageType(messageType proto.MessageType, handler p2p.MsgHandler) {
	if p.Handlers == nil {
		p.Handlers = make(map[proto.MessageType][]p2p.MsgHandler)
	}
	p.Handlers[messageType] = append(p.Handlers[messageType], handler)
	p.Called(messageType, handler)
}

func (p *MockP2p) ExtBroadcast(from peer.ID, messageType proto.MessageType, payload interface{}) {
	if p.Handlers == nil {
		p.Handlers = make(map[proto.MessageType][]p2p.MsgHandler)
	}

	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	for _, handler := range p.Handlers[messageType] {
		handler(from, data)
	}
}

func (p *MockP2p) ClearBroadcasted() {
	p.Broadcasted = []proto.Message{}
}

func (m *MockP2p) Address() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockP2p) Sign(data []byte) ([]byte, error) {
	args := m.Called(data)
	return args.Get(0).([]byte), args.Error(1)
}

var _ p2p.Interface = &MockP2p{}