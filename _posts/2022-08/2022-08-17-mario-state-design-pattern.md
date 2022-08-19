---
layout: post
title: "[Design Pattern] SuperMario"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Design Pattern
---

![](https://i.insider.com/5849abcedd089526558b45fd?width=1000&format=jpeg&auto=webp)

```go
package aug

import (
	"fmt"
	"testing"
)

type IMarioState interface {
	getState() string
	obtainMushRoom(*MarioStateMachine)
	meetMonster(*MarioStateMachine)
}

// Small
type SmallMario struct{}

func (m *SmallMario) getState() string                     { return "Small" }
func (m *SmallMario) obtainMushRoom(sm *MarioStateMachine) { sm.marioState = &SuperMario{} }
func (m *SmallMario) meetMonster(sm *MarioStateMachine)    { sm.marioState = &DeadMario{} }

// Super
type SuperMario struct{}

func (m *SuperMario) getState() string                     { return "Super" }
func (m *SuperMario) obtainMushRoom(sm *MarioStateMachine) {}
func (m *SuperMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = &SmallMario{}
}

// Dead
type DeadMario struct{}

func (m *DeadMario) getState() string                     { return "Dead" }
func (m *DeadMario) obtainMushRoom(sm *MarioStateMachine) {}
func (m *DeadMario) meetMonster(sm *MarioStateMachine)    {}

type MarioStateMachine struct {
	marioState IMarioState
}

func (m *MarioStateMachine) ObtainMushRoom() {
	m.marioState.obtainMushRoom(m)
}

func (m *MarioStateMachine) MeetMonster() {
	m.marioState.meetMonster(m)
}

func (m *MarioStateMachine) state() string {
	return m.marioState.getState()
}

func TestMario(t *testing.T) {
	fsm := &MarioStateMachine{
		marioState: &SmallMario{},
	}
	fsm.MeetMonster()
	fmt.Printf("%v\n", fsm.state())

	fsm = &MarioStateMachine{
		marioState: &SmallMario{},
	}
	fsm.ObtainMushRoom()
	fmt.Printf("%v\n", fsm.state())
	fsm.MeetMonster()
	fmt.Printf("%v\n", fsm.state())
}
```