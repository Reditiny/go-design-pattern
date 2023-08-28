package main

import (
	responsibility_chan "design_pattern/3_3_responsibility_chan/chan"
)

func main() {
	// 建立各级领导的责任链
	boss := &responsibility_chan.Boss{Max: 7}
	manager := &responsibility_chan.Manager{Max: 3, NextHandler: boss}
	leader := &responsibility_chan.GroupLeader{Max: 1, NextHandler: manager}
	// 员工提出请假
	breakDays := 9
	leader.Handle(breakDays)
}
