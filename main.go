/*
Copyright 2020 The Flux authors

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

package main

import (
	"context"
	"flag"
	"log"

	"github.com/go-logr/logr"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/shencan/terraform-provider-aliyun/pkg"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	version string = "dev"
)

func main() {
	ctrllog.SetLogger(logr.New(ctrllog.NullLogSink{}))
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "github.com/shencan/aliyun",
		Debug:   *debugFlag,
	}
	err := providerserver.Serve(context.Background(), pkg.New(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
