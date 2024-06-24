/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package notifications

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	CommonNotification = NotificationAggregator{map[string][]Notification{}}
}

const (
	InfoNotification    MessageType = "INFO"
	WarningNotification MessageType = "WARNING"
	ErrorNotification   MessageType = "ERROR"
)

type MessageType string

type Notification struct {
	Type           MessageType
	Message        string
	CallingObjects []client.Object
}

type NotificationAggregator struct {
	Notifications map[string][]Notification
}

var CommonNotification NotificationAggregator

// DispatchNotification is used to send a notification to the NotificationAggregator
func (na *NotificationAggregator) DispatchNotification(notification Notification, ProviderName string) {
	na.Notifications[ProviderName] = append(na.Notifications[ProviderName], notification)
}

// ProcessNotifications takes all generated notifications and displays it in a tabular format based on provider
func (na *NotificationAggregator) ProcessNotifications() {
	for provider, msgs := range na.Notifications {
		t := newTableConfig()

		t.SetTitle(fmt.Sprintf("Notifications from %v", provider))
		t.AppendHeader(table.Row{"Message Type", "Notification", "Calling Object"})

		for _, n := range msgs {
			t.AppendRow(table.Row{n.Type, n.Message, convertObjectsToStr(n.CallingObjects)})
		}

		t.Render()
	}
}

func convertObjectsToStr(ob []client.Object) string {
	var sb strings.Builder

	for i, o := range ob {
		if i > 0 {
			sb.WriteString(", ")
		}
		object := o.GetObjectKind().GroupVersionKind().Kind + ": " + client.ObjectKeyFromObject(o).String()
		sb.WriteString(object)
	}

	return sb.String()
}
