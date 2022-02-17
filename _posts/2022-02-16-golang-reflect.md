---
layout: post
title: "reflect.ValueOf() & reflect.TypeOf()"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Golang
---


isRTMPPullWithAudio := false
playSpecRef := reflect.ValueOf(&remoteSpec.PlaySpec).Elem()
for i := 0; i < playSpecRef.NumField(); i++ {
	sf := playSpecRef.Type().FieldByIndex([]int{i})
	if sf.Tag.Get("default") != fmt.Sprint(playSpecRef.FieldByIndex([]int{i}).Interface()) {
		isRTMPPullWithAudio = true
		break
	}
}