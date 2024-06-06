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
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const (
	bgColor = text.BgBlack

	infoColor    = text.FgHiWhite
	warningColor = text.FgHiMagenta
	errorColor   = text.FgRed
	defaultColor = text.FgWhite
)

func newTableConfig() table.Writer {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetRowPainter(func(row table.Row) text.Colors {
		switch row[1] {
		case InfoNotification:
			return text.Colors{bgColor, infoColor}
		case WarningNotification:
			return text.Colors{bgColor, warningColor}
		case ErrorNotification:
			return text.Colors{bgColor, errorColor}
		default:
			return text.Colors{bgColor, defaultColor}
		}
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, WidthMax: 100},
	})

	style := table.StyleRounded
	style.Options.SeparateRows = true
	style.Title.Align = text.AlignCenter

	t.SetStyle(style)

	return t
}