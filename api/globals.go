package api

import "sync"

var ThresholdProfiles []*ThresholdProfile
var ThresholdProfilesLock sync.Mutex // guards ThresholdProfiles

var NotificationProfiles []*NotificationProfile
var NotificationProfilesLock sync.Mutex // guards NotificationProfiles

var LocationProfiles []*LocationProfile
var LocationProfilesLock sync.Mutex // guards LocationProfiles

var MonitorGroups []*MonitorGroup
var MonitorGroupsLock sync.Mutex // guards MonitorGroups

var UserGroups []*UserGroup
var UserGroupsLock sync.Mutex // guards UserGroups

var TagsList []*Tag
var TagsListLock sync.Mutex // guards Tags
