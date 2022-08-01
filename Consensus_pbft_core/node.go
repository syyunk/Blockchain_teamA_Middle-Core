package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var count int = 0

type Node struct {
	NodeID           string
	NodeAddressTable map[string]string // key=nodeID, value=url
	View             *View
	CurrentState     *State
	CommittedMsgs    []*RequestMsg // kinda block.
	MsgBuffer        *MsgBuffer
	MsgEntrance      chan interface{}
	MsgDelivery      chan interface{}
	Alarm            chan bool
	Active           bool
}

type MsgBuffer struct {
	ReqMsgs        []*RequestMsg
	PrePrepareMsgs []*PrePrepareMsg
	PrepareMsgs    []*VoteMsg
	CommitMsgs     []*VoteMsg
}

type View struct {
	ID      int64
	Primary string
}

//합의 진행 속도
const ResolvingTimeDuration = time.Millisecond / 100 // 1 second.

//
// cmd> exe 5000 [enter] <---- 5000 = nodeID
func NewNode(nodeID string) *Node {
	const viewID = 10000000000 // temporary.

	node := &Node{
		// Hard-coded for test.
		NodeID: nodeID,
		NodeAddressTable: map[string]string{
			"P1":  "localhost:5000",
			"P2":  "localhost:4000",
			"P3":  "localhost:3000",
			"P4":  "localhost:2000",
			// "P5":  "localhost:1000",
			// "P6":  "localhost:1100",
			// "P7":  "localhost:1200",
			// "P8":  "localhost:1300",
			// "P50": "localhost:1400",
			// "P9":  "localhost:1500",
			// "P10": "localhost:1600",
			// "P11": "localhost:1700",
			// "P12": "localhost:1800",
			// "P13": "localhost:1900",
			// "P14": "localhost:2100",
			// "P15": "localhost:2200",
			// "P16": "localhost:2300",
			// "P17": "localhost:2400",
			// "P18": "localhost:2500",
			// "P19": "localhost:2600",
			// "P20": "localhost:2700",
			// "P21": "localhost:2800",
			// "P22": "localhost:2900",
			// "P23": "localhost:3100",
			// "P24": "localhost:3200",
			// "P25": "localhost:3300",
			// "P26": "localhost:3400",
			// "P27": "localhost:3500",
			// "P28": "localhost:3600",
			// "P29": "localhost:3700",
			// "P30": "localhost:3800",
			// "P31": "localhost:3900",
			// "P32": "localhost:4100",
			// "P33": "localhost:4200",
			// "P34": "localhost:4300",
			// "P35": "localhost:4400",
			// "P36": "localhost:4500",
			// "P37": "localhost:4600",
			// "P38": "localhost:4700",
			// "P39": "localhost:4800",
			// "P40": "localhost:4900",
			// "P41": "localhost:5100",
			// "P42": "localhost:5200",
			// "P43": "localhost:5300",
			// "P44": "localhost:5400",
			// "P45": "localhost:5500",
			// "P46": "localhost:5600",
			// "P47": "localhost:5700",
			// "P48": "localhost:5800",
			// "P49": "localhost:5900",
		},
		View: &View{
			ID:      viewID,
			Primary: "P1",
		},

		// Consensus-related struct
		CurrentState:  nil,
		CommittedMsgs: make([]*RequestMsg, 0),
		MsgBuffer: &MsgBuffer{
			ReqMsgs:        make([]*RequestMsg, 0),
			PrePrepareMsgs: make([]*PrePrepareMsg, 0),
			PrepareMsgs:    make([]*VoteMsg, 0),
			CommitMsgs:     make([]*VoteMsg, 0),
		},

		// Channels
		MsgEntrance: make(chan interface{}, 10000),
		MsgDelivery: make(chan interface{}, 10000),
		Alarm:       make(chan bool),
	}

	// Start message dispatcher
	go node.dispatchMsg()

	// Start alarm trigger
	go node.alarmToDispatcher()
	// Start message resolver
	go node.resolveMsg()

	return node
}

func (node *Node) Broadcast(msg interface{}, path string) map[string]error {
	errorMap := make(map[string]error)

	for nodeID, url := range node.NodeAddressTable {
		if nodeID == node.NodeID {
			continue
		}

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			errorMap[nodeID] = err
			continue
		}

		send(url+path, jsonMsg)
	}

	if len(errorMap) == 0 {
		return nil
	} else {
		return errorMap
	}
}

func (node *Node) Reply(msg *ReplyMsg) error {
	// Print all committed messages.
	//for _, value := range node.CommittedMsgs {
	//	fmt.Printf("Committed value: %s, %d, %s, %d", value.ClientID, value.Timestamp, value.Operation, value.SequenceID)
	//}
	//fmt.Print("\n")

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Client가 없으므로, 일단 Primary에게 보내는 걸로 처리
	send(node.NodeAddressTable[node.View.Primary]+"/reply", jsonMsg)

	return nil
}

// GetReq can be called when the node's CurrentState is nil.
// Consensus start procedure for the Primary.
func (node *Node) GetReq(reqMsg *RequestMsg) error {
	LogMsg(reqMsg)

	// Create a new state for the new
	err := node.createStateForNewConsensus()
	if err != nil {
		return err
	}

	// Start the consensus process.
	prePrepareMsg, err := node.CurrentState.StartConsensus(reqMsg)
	if err != nil {
		return err
	}

	LogStage(fmt.Sprintf("Consensus Process (ViewID:%d)", node.CurrentState.ViewID), false)

	// Send getPrePrepare message
	if prePrepareMsg != nil {
		node.Broadcast(prePrepareMsg, "/preprepare")
		LogStage("Pre-prepare", true)
	}

	return nil
}

// GetPrePrepare can be called when the node's CurrentState is nil.
// Consensus start procedure for normal participants.
func (node *Node) GetPrePrepare(prePrepareMsg *PrePrepareMsg) error {
	LogMsg(prePrepareMsg)

	//=======================================================================
	for i := 0; i < len(node.CommittedMsgs); i++ {
		if prePrepareMsg.SequenceID == node.CommittedMsgs[i].SequenceID {
			return nil
		}
	}
	//=======================================================================

	// Create a new state for the new
	err := node.createStateForNewConsensus()
	if err != nil {
		return err
	}

	prePareMsg, err := node.CurrentState.PrePrepare(prePrepareMsg)
	if err != nil {
		return err
	}

	if prePareMsg != nil {
		// Attach node ID to the message
		prePareMsg.NodeID = node.NodeID

		LogStage("Pre-prepare", true)
		node.Broadcast(prePareMsg, "/prepare")
		LogStage("Prepare", false)
	}

	return nil
}

func (node *Node) GetPrepare(prepareMsg *VoteMsg) error {
	LogMsg(prepareMsg)

	//=======================================================================
	for i := 0; i < len(node.CommittedMsgs); i++ {
		if prepareMsg.SequenceID == node.CommittedMsgs[i].SequenceID {
			return nil
		}
	}
	//=======================================================================

	commitMsg, err := node.CurrentState.Prepare(prepareMsg)
	if err != nil {
		return err
	}

	if commitMsg != nil {
		// Attach node ID to the message
		commitMsg.NodeID = node.NodeID

		LogStage("Prepare", true)
		node.Broadcast(commitMsg, "/commit")
		LogStage("Commit", false)
	}

	return nil
}

func (node *Node) GetCommit(commitMsg *VoteMsg) error {
	LogMsg(commitMsg)

	//=======================================================================
	for i := 0; i < len(node.CommittedMsgs); i++ {
		if commitMsg.SequenceID == node.CommittedMsgs[i].SequenceID {
			return nil
		}
	}
	//=======================================================================

	replyMsg, committedMsg, err := node.CurrentState.Commit(commitMsg)
	if err != nil {
		return err
	}

	if replyMsg != nil {
		if committedMsg == nil {
			return errors.New("committed message is nil, even though the reply message is not nil")
		}

		// Attach node ID to the message
		replyMsg.NodeID = node.NodeID

		// Save the last version of committed messages to node.
		node.CommittedMsgs = append(node.CommittedMsgs, committedMsg)

		LogStage("Commit", true)
		node.Reply(replyMsg)
		LogStage("Reply", true)

		count += 1
		fmt.Println("============Concensus", count, "Complete============")
		node.StateInitial()
	}

	return nil
}

func (node *Node) GetReply(msg *ReplyMsg) {
	fmt.Printf("Result: %s by %s\n", msg.Result, msg.NodeID)
}

//=========================================================
func (node *Node) StateInitial() {
	//clean the buffer
	node.Active = false
	node.MsgBuffer = &MsgBuffer{
		ReqMsgs:        make([]*RequestMsg, 0),
		PrePrepareMsgs: make([]*PrePrepareMsg, 0),
		PrepareMsgs:    make([]*VoteMsg, 0),
		CommitMsgs:     make([]*VoteMsg, 0),
	}
}

//=========================================================

func (node *Node) createStateForNewConsensus() error {
	// Check if there is an ongoing consensus process.
	if node.Active == true {
		return errors.New("another consensus is ongoing")
	}

	node.Active = true

	// Get the last sequence ID
	var lastSequenceID int64
	if len(node.CommittedMsgs) == 0 {
		lastSequenceID = -1
	} else {
		lastSequenceID = node.CommittedMsgs[len(node.CommittedMsgs)-1].SequenceID
	}

	// Create a new state for this new consensus process in the Primary
	node.CurrentState = CreateState(node.View.ID, lastSequenceID)

	LogStage("Create the replica status", true)

	return nil
}

func (node *Node) dispatchMsg() {
	for {
		select {
		case msg := <-node.MsgEntrance:
			err := node.routeMsg(msg)
			if err != nil {
				fmt.Println(err)
				// TODO: send err to ErrorChannel
			}
		case <-node.Alarm:
			err := node.routeMsgWhenAlarmed()
			if err != nil {
				fmt.Println(err)
				// TODO: send err to ErrorChannel
			}
		}
	}
}

func (node *Node) routeMsg(msg interface{}) []error {
	switch msg.(type) {
	case *RequestMsg:
		if node.Active == false {
			// Copy buffered messages first.
			msgs := make([]*RequestMsg, len(node.MsgBuffer.ReqMsgs))
			copy(msgs, node.MsgBuffer.ReqMsgs)

			// Append a newly arrived message.
			msgs = append(msgs, msg.(*RequestMsg))

			// Empty the buffer.
			node.MsgBuffer.ReqMsgs = make([]*RequestMsg, 0)

			// Send messages.
			node.MsgDelivery <- msgs
		} else {
			node.MsgBuffer.ReqMsgs = append(node.MsgBuffer.ReqMsgs, msg.(*RequestMsg))
		}
	case *PrePrepareMsg:
		if node.Active == false {
			// Copy buffered messages first.
			msgs := make([]*PrePrepareMsg, len(node.MsgBuffer.PrePrepareMsgs))
			copy(msgs, node.MsgBuffer.PrePrepareMsgs)

			// Append a newly arrived message.
			msgs = append(msgs, msg.(*PrePrepareMsg))

			// Empty the buffer.
			node.MsgBuffer.PrePrepareMsgs = make([]*PrePrepareMsg, 0)

			// Send messages.
			node.MsgDelivery <- msgs
		} else {
			node.MsgBuffer.PrePrepareMsgs = append(node.MsgBuffer.PrePrepareMsgs, msg.(*PrePrepareMsg))
		}
	case *VoteMsg:
		if msg.(*VoteMsg).MsgType == PrepareMsg {
			if node.CurrentState == nil || node.CurrentState.CurrentStage != PrePrepared {
				node.MsgBuffer.PrepareMsgs = append(node.MsgBuffer.PrepareMsgs, msg.(*VoteMsg))
			} else {
				// Copy buffered messages first.
				msgs := make([]*VoteMsg, len(node.MsgBuffer.PrepareMsgs))
				copy(msgs, node.MsgBuffer.PrepareMsgs)

				// Append a newly arrived message.
				msgs = append(msgs, msg.(*VoteMsg))

				// Empty the buffer.
				node.MsgBuffer.PrepareMsgs = make([]*VoteMsg, 0)

				// Send messages.
				node.MsgDelivery <- msgs
			}
		} else if msg.(*VoteMsg).MsgType == CommitMsg {
			if node.CurrentState == nil || node.CurrentState.CurrentStage != Prepared {
				node.MsgBuffer.CommitMsgs = append(node.MsgBuffer.CommitMsgs, msg.(*VoteMsg))
			} else {
				// Copy buffered messages first.
				msgs := make([]*VoteMsg, len(node.MsgBuffer.CommitMsgs))
				copy(msgs, node.MsgBuffer.CommitMsgs)

				// Append a newly arrived message.
				msgs = append(msgs, msg.(*VoteMsg))

				// Empty the buffer.
				node.MsgBuffer.CommitMsgs = make([]*VoteMsg, 0)

				// Send messages.
				node.MsgDelivery <- msgs
			}
		}
	}

	return nil
}

func (node *Node) routeMsgWhenAlarmed() []error {
	if node.CurrentState == nil {
		// Check ReqMsgs, send them.
		if len(node.MsgBuffer.ReqMsgs) != 0 {
			msgs := make([]*RequestMsg, len(node.MsgBuffer.ReqMsgs))
			copy(msgs, node.MsgBuffer.ReqMsgs)

			node.MsgDelivery <- msgs
		}

		// Check PrePrepareMsgs, send them.
		if len(node.MsgBuffer.PrePrepareMsgs) != 0 {
			msgs := make([]*PrePrepareMsg, len(node.MsgBuffer.PrePrepareMsgs))
			copy(msgs, node.MsgBuffer.PrePrepareMsgs)

			node.MsgDelivery <- msgs
		}
	} else {
		switch node.CurrentState.CurrentStage {
		case PrePrepared:
			// Check PrepareMsgs, send them.
			if len(node.MsgBuffer.PrepareMsgs) != 0 {
				msgs := make([]*VoteMsg, len(node.MsgBuffer.PrepareMsgs))
				copy(msgs, node.MsgBuffer.PrepareMsgs)

				node.MsgDelivery <- msgs
			}
		case Prepared:
			// Check CommitMsgs, send them.
			if len(node.MsgBuffer.CommitMsgs) != 0 {
				msgs := make([]*VoteMsg, len(node.MsgBuffer.CommitMsgs))
				copy(msgs, node.MsgBuffer.CommitMsgs)

				node.MsgDelivery <- msgs
			}
		}
	}

	return nil
}

func (node *Node) resolveMsg() {
	for {
		// Get buffered messages from the dispatcher.
		msgs := <-node.MsgDelivery
		switch msgs.(type) {
		case []*RequestMsg:
			errs := node.resolveRequestMsg(msgs.([]*RequestMsg))
			if len(errs) != 0 {
				for _, err := range errs {
					fmt.Println(err)
				}
				// TODO: send err to ErrorChannel
			}
		case []*PrePrepareMsg:
			errs := node.resolvePrePrepareMsg(msgs.([]*PrePrepareMsg))
			if len(errs) != 0 {
				for _, err := range errs {
					fmt.Println(err)
				}
				// TODO: send err to ErrorChannel
			}
		case []*VoteMsg:
			voteMsgs := msgs.([]*VoteMsg)
			if len(voteMsgs) == 0 {
				break
			}

			if voteMsgs[0].MsgType == PrepareMsg {
				errs := node.resolvePrepareMsg(voteMsgs)
				if len(errs) != 0 {
					for _, err := range errs {
						fmt.Println(err)
					}
					// TODO: send err to ErrorChannel
				}
			} else if voteMsgs[0].MsgType == CommitMsg {
				errs := node.resolveCommitMsg(voteMsgs)
				if len(errs) != 0 {
					for _, err := range errs {
						fmt.Println(err)
					}
					// TODO: send err to ErrorChannel
				}
			}
		}
	}
}

func (node *Node) alarmToDispatcher() {
	for {
		time.Sleep(ResolvingTimeDuration)
		node.Alarm <- true
	}
}

func (node *Node) resolveRequestMsg(msgs []*RequestMsg) []error {
	errs := make([]error, 0)

	// Resolve messages
	for _, reqMsg := range msgs {
		err := node.GetReq(reqMsg)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

func (node *Node) resolvePrePrepareMsg(msgs []*PrePrepareMsg) []error {
	errs := make([]error, 0)

	// Resolve messages
	for _, prePrepareMsg := range msgs {
		err := node.GetPrePrepare(prePrepareMsg)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

func (node *Node) resolvePrepareMsg(msgs []*VoteMsg) []error {
	errs := make([]error, 0)

	// Resolve messages
	for _, prepareMsg := range msgs {
		err := node.GetPrepare(prepareMsg)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

func (node *Node) resolveCommitMsg(msgs []*VoteMsg) []error {
	errs := make([]error, 0)

	// Resolve messages
	for _, commitMsg := range msgs {
		err := node.GetCommit(commitMsg)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}
