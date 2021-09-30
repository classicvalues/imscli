// Copyright 2021 Adobe. All rights reserved.
// This file is licensed to you under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License. You may obtain a copy
// of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under
// the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
// OF ANY KIND, either express or implied. See the License for the specific language
// governing permissions and limitations under the License.

package cmd

import (
	"fmt"

	"github.com/adobe/imscli/ims"
	"github.com/spf13/cobra"
)

func invalidateServiceTokenCmd(imsConfig *ims.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serviceToken",
		Aliases: []string{"svc"},
		Short:   "Invalidate a service token.",
		Long:    "Invalidate a service token.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true

			err := imsConfig.InvalidateToken()
			if err != nil {
				return fmt.Errorf("error invalidating the service token: %v", err)
			}
			fmt.Println("Service token successfully invalidated.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&imsConfig.RefreshToken, "serviceToken", "t", "", "Service token.")
	cmd.Flags().StringVarP(&imsConfig.ClientID, "clientID", "c", "", "IMS Client ID.")
	cmd.Flags().StringVarP(&imsConfig.ClientSecret, "clientSecret", "s", "", "IMS Client Secret.")
	cmd.Flags().BoolVarP(&imsConfig.Cascading, "cascading", "a", false,
		"Also invalidate all tokens obtained with the refresh token.")

	return cmd
}
