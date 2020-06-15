// Copyright 2017 Google Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "fmt"
import "github.com/fatih/color"
import "github.com/InVisionApp/tabular"

var tab tabular.Table

func init() {
	tab = tabular.New()
	tab.Col("env", "Environment", 14)
	tab.Col("cls", "Cluster", 10)
	tab.Col("svc", "Service", 15)
	tab.Col("hst", "Database Host", 20)
	tab.ColRJ("pct", "%CPU", 7)
}

var data = []struct {
	e, c, s, d string
	v          float64
}{
	{
		e: "production",
		c: "cluster-1",
		s: "service-a",
		d: "database-host-1",
		v: 70.01,
	},
	{
		e: "production",
		c: "cluster-1",
		s: "service-b",
		d: "database-host-2",
		v: 99.51,
	},
	{
		e: "production",
		c: "cluster-2",
		s: "service-a",
		d: "database-host-1",
		v: 70.01,
	},
	{
		e: "production",
		c: "cluster-2",
		s: "service-b",
		d: "database-host-2",
		v: 99.51,
	},
}

func main() {
	fmt.Println("Hello, world!")
	color.Cyan("Prints text in cyan.")

	// Print a subset of columns (Environments and Clusters)
	format := tab.Print("env", "cls")
	for _, x := range data {
		fmt.Printf(format, x.e, x.c)
	}	
}

