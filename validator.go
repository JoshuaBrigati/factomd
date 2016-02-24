// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	ss "github.com/FactomProject/factomd/state"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/log"
)

var _ = fmt.Print
var _ = log.Printf
var _ = os.Exit
func Validator(state interfaces.IState) {
	s := state.(*ss.State)
	var _ = s
	
	for {
		fmt.Println(".")
		
		select {
			case i := <- s.ShutdownChan:
				state.GetDB().(interfaces.IDatabase).Close()
				os.Exit(i)
			default :
		}
		
		msg := <-state.InMsgQueue() // Get message from the input queue
		state.UpdateState()
		
		if state.PrintType(msg.Type()) {
			state.Print(fmt.Sprintf("%20s %s", "Validator:", msg.String()))
		}

		switch msg.Validate(state.GetDBHeight(), state) { // Validate the message.
		case 1: // Process if valid
			state.NetworkOutMsgQueue() <- msg
			if state.PrintType(msg.Type()) {
				state.Print(" Valid\n")
			}
			if msg.Leader(state) {
				if state.PrintType(msg.Type()) {
					state.Print(fmt.Sprintf("%20s %s\n", "Leader:", msg.String()))
				}
				msg.LeaderExecute(state)
				state.UpdateState()
			} else if msg.Follower(state) {
				if state.PrintType(msg.Type()) {
					state.Print(fmt.Sprintf("%20s %s\n", "Follower:", msg.String()))
				}
				msg.FollowerExecute(state)
				state.UpdateState()
			} else {
				state.Print(" Message ignored\n")
			}
		case 0: // Hold for later if unknown.
			if state.PrintType(msg.Type()) {
				state.Print(" Hold\n")
			}
		default:
			if state.PrintType(msg.Type()) {
				state.Print(" Invalid\n")
			}
			state.NetworkInvalidMsgQueue() <- msg
		}
	}

}
