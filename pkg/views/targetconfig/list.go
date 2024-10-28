// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package targetconfig

import (
	"fmt"
	"sort"

	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/views"
	"github.com/daytonaio/daytona/pkg/views/util"
)

type rowData struct {
	ConfigName string
	Provider   string
	IsDefault  string
	Options    string
}

func ListTargetConfigs(targetConfigs []apiclient.TargetConfig) {
	sortTargetConfigs(&targetConfigs)

	data := [][]string{}

	for _, targetConfig := range targetConfigs {
		data = append(data, getRowFromRowData(&targetConfig))
	}

	table := util.GetTableView(data, []string{
		"Name", "Provider", "Default", "Options",
	}, nil, func() {
		renderUnstyledList(targetConfigs)
	})

	fmt.Println(table)
}

func getRowFromRowData(targetConfig *apiclient.TargetConfig) []string {
	var isDefault string
	var data rowData

	data.ConfigName = targetConfig.Name
	data.Provider = targetConfig.ProviderInfo.Name
	data.Options = targetConfig.Options

	if targetConfig.IsDefault {
		isDefault = views.ActiveStyle.Render("Yes")
	} else {
		isDefault = views.InactiveStyle.Render("/")
	}

	row := []string{
		views.NameStyle.Render(data.ConfigName),
		views.DefaultRowDataStyle.Render(data.Provider),
		isDefault,
		views.DefaultRowDataStyle.Render(data.Options),
	}

	return row
}

func sortTargetConfigs(targetConfigs *[]apiclient.TargetConfig) {
	sort.Slice(*targetConfigs, func(i, j int) bool {
		t1 := (*targetConfigs)[i]
		t2 := (*targetConfigs)[j]
		return t1.ProviderInfo.Name < t2.ProviderInfo.Name
	})
}

func renderUnstyledList(targetConfigs []apiclient.TargetConfig) {
	output := "\n"

	for _, targetConfig := range targetConfigs {
		output += fmt.Sprintf("%s %s", views.GetPropertyKey("Name: "), targetConfig.Name) + "\n\n"

		output += fmt.Sprintf("%s %s", views.GetPropertyKey("Provider: "), targetConfig.ProviderInfo.Name) + "\n\n"

		if targetConfig.IsDefault {
			output += fmt.Sprintf("%s %s", views.GetPropertyKey("Default: "), "Yes") + "\n\n"
		}

		output += fmt.Sprintf("%s %s", views.GetPropertyKey("Options: "), targetConfig.Options) + "\n\n"

		if targetConfig.Name != targetConfigs[len(targetConfigs)-1].Name {
			output += views.SeparatorString + "\n\n"
		}
	}

	fmt.Println(output)
}