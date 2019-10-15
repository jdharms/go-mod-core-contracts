/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"
)

// CallbackAlert indicates an action to take when a callback fires.
type CallbackAlert struct {
	ActionType ActionType `json:"type,omitempty"`
	Id         string     `json:"id,omitempty"`
}

func (ca CallbackAlert) MarshalJSON() ([]byte, error) {
	test := struct {
		ActionType ActionType `json:"type,omitempty"`
		Id         string     `json:"id,omitempty"`
	}{
		ActionType: ca.ActionType,
		Id: ca.Id,
	}

	return json.Marshal(test)
}

/*
 * String function for representing a CallbackAlert
 */
func (ca CallbackAlert) String() string {
	out, err := json.Marshal(ca)
	if err != nil {
		return err.Error()
	}

	return string(out)
}
