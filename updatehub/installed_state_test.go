/*
 * UpdateHub
 * Copyright (C) 2017
 * O.S. Systems Sofware LTDA: contato@ossystems.com.br
 *
 * SPDX-License-Identifier:     GPL-2.0
 */

package updatehub

import (
	"testing"

	"github.com/UpdateHub/updatehub/metadata"
	"github.com/UpdateHub/updatehub/testsmocks/activeinactivemock"
	"github.com/stretchr/testify/assert"
)

func TestStateInstalled(t *testing.T) {
	m := &metadata.UpdateMetadata{}
	s := NewInstalledState(m)

	aim := &activeinactivemock.ActiveInactiveMock{}

	uh, err := newTestUpdateHub(s, aim)
	assert.NoError(t, err)

	nextState, _ := s.Handle(uh)
	expectedState := NewRebootState()

	assert.Equal(t, expectedState, nextState)
	assert.Equal(t, m, s.UpdateMetadata())

	aim.AssertExpectations(t)
}

func TestStateInstalledToMap(t *testing.T) {
	state := NewInstalledState(nil)

	expectedMap := map[string]interface{}{}
	expectedMap["status"] = "installed"

	assert.Equal(t, expectedMap, state.ToMap())
}
