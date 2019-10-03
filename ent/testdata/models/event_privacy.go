// Code generated by github.com/lolopinto/ent/ent, DO NOT edit.

package models

import (
	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/privacy"
	"github.com/lolopinto/ent/ent/viewer"
)

// EventPrivacyPolicy is the privacy policy for the Event ent which helps decides if it's
// visible to the viewer
type EventPrivacyPolicy struct {
	Event *Event
}

// Rules is the list of rules that decides the visibility of the Event ent to the viewer
func (policy EventPrivacyPolicy) Rules() []ent.PrivacyPolicyRule {
	return []ent.PrivacyPolicyRule{
		privacy.AllowIfOmniscientRule{},
		// BEGIN MANUAL SECTION: Add custom privacy rules below
		// END MANUAL SECTION of privacy rules
		privacy.AlwaysDenyRule{},
	}
}

// AllowIfViewerCanSeeEventRule is a reusable rule that can be called by different ents to see if the contact can be visible
type AllowIfViewerCanSeeEventRule struct {
	EventID string
}

// GenEval evaluates that the ent is visible to the user
func (rule AllowIfViewerCanSeeEventRule) GenEval(viewer viewer.ViewerContext, entity interface{}, privacyResultChan chan<- ent.PrivacyResult) {
	entResultChan := make(chan EventResult)
	go GenLoadEvent(viewer, rule.EventID, entResultChan)
	entResult := <-entResultChan

	if entResult.Error != nil {
		privacyResultChan <- ent.SkipPrivacyResult
	} else {
		privacyResultChan <- ent.AllowPrivacyResult
	}
}
