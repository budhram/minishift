/*
Copyright (C) 2017 Red Hat, Inc.

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

package image

import (
	"github.com/spf13/cobra"
)

const (
	noCachedImagesSpecified = "You need to either specify a list of images on the command line or configure the list of cached images via 'image config [add|remove]'."
)

var ImageCmd = &cobra.Command{
	Use:     "image SUBCOMMAND [flags]",
	Aliases: []string{"images"},
	Short:   "Exports and imports container images.",
	Long:    "Exports and imports container images.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
