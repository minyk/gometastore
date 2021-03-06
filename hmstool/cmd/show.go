// Copyright © 2018 Alex Kolbasov
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

package cmd

import "github.com/spf13/cobra"

// dbCmd represents the db command
var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"ls"},
	Short:   "Show objects",
}

var showDbCmd = &cobra.Command{
	Use:     "databases",
	Aliases: []string{"db"},
	Short:   "show databases",
	Run:     listDbs,
}

var showTablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "list tables",
	Run:   listTables,
}

var showTableCmd = &cobra.Command{
	Use:   "table",
	Short: "show table",
	Run:   showTables,
}

var showPartitionsCmd = &cobra.Command{
	Use:   "partitions",
	Short: "show partitions",
	Run:   showPartitions,
}

var showPartitionCmd = &cobra.Command{
	Use:   "partition",
	Short: "show partition",
	Run:   showPartition,
}

func init() {
	showCmd.PersistentFlags().StringP(optDbName, "d", "", "database name")
	showCmd.PersistentFlags().StringP(optTableName, "t", "", "table name")
	showCmd.AddCommand(showDbCmd)
	showCmd.AddCommand(showTableCmd)
	showCmd.AddCommand(showTablesCmd)
	showCmd.AddCommand(showPartitionsCmd)
	showCmd.AddCommand(showPartitionCmd)
	rootCmd.AddCommand(showCmd)
}
