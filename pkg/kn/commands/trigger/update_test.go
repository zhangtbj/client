// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trigger

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	eventing_client "knative.dev/client/pkg/eventing/v1alpha1"
	knserving_client "knative.dev/client/pkg/serving/v1alpha1"
	"knative.dev/client/pkg/util"
	serving_v1alpha1 "knative.dev/serving/pkg/apis/serving/v1alpha1"
)

func TestTriggerUpdate(t *testing.T) {
	eventingClient := eventing_client.NewMockKnEventingClient(t)
	servingClient := knserving_client.NewMockKnServiceClient(t)

	servingRecorder := servingClient.Recorder()
	servingRecorder.GetService("mysvc", &serving_v1alpha1.Service{
		TypeMeta:   metav1.TypeMeta{Kind: "Service"},
		ObjectMeta: metav1.ObjectMeta{Name: "mysvc"},
	}, nil)

	eventingRecorder := eventingClient.Recorder()
	present := createTrigger("default", triggerName, map[string]string{"type": "dev.knative.foo"}, "mybroker", "mysvc")
	updated := createTrigger("default", triggerName, map[string]string{"type": "dev.knative.new"}, "mybroker", "mysvc")
	eventingRecorder.GetTrigger(triggerName, present, nil)
	eventingRecorder.UpdateTrigger(updated, nil)

	out, err := executeTriggerCommand(eventingClient, servingClient, "update", triggerName,
		"--filter", "type=dev.knative.new", "--sink", "svc:mysvc")
	assert.NilError(t, err, "Trigger should be updated")
	util.ContainsAll(out, "Trigger", triggerName, "updated", "namespace", "default")

	eventingRecorder.Validate()
	servingRecorder.Validate()
}

func TestTriggerUpdateWithError(t *testing.T) {
	eventingClient := eventing_client.NewMockKnEventingClient(t)
	eventingRecorder := eventingClient.Recorder()
	eventingRecorder.GetTrigger(triggerName, nil, fmt.Errorf("trigger not found"))

	out, err := executeTriggerCommand(eventingClient, nil, "update", triggerName,
		"--filter", "type=dev.knative.new", "--sink", "svc:newsvc")
	assert.ErrorContains(t, err, "trigger not found")
	util.ContainsAll(out, "Usage", triggerName)

	eventingRecorder.Validate()
}

func TestTriggerUpdateInvalidBroker(t *testing.T) {
	eventingClient := eventing_client.NewMockKnEventingClient(t)
	eventingRecorder := eventingClient.Recorder()
	present := createTrigger("default", triggerName, map[string]string{"type": "dev.knative.new"}, "mybroker", "newsvc")
	eventingRecorder.GetTrigger(triggerName, present, nil)

	out, err := executeTriggerCommand(eventingClient, nil, "update", triggerName,
		"--broker", "newbroker")
	assert.ErrorContains(t, err, "broker is immutable")
	util.ContainsAll(out, "Usage", triggerName)

	eventingRecorder.Validate()
}
