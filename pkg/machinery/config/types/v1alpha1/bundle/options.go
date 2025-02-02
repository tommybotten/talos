// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bundle

import (
	jsonpatch "github.com/evanphx/json-patch"

	"github.com/talos-systems/talos/pkg/machinery/config/configpatcher"
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1/generate"
)

// Option controls config options specific to config bundle generation.
type Option func(o *Options) error

// InputOptions holds necessary params for generating an input.
type InputOptions struct {
	ClusterName string
	Endpoint    string
	KubeVersion string
	GenOptions  []generate.GenOption
}

// Options describes generate parameters.
type Options struct {
	ExistingConfigs string // path to existing config files
	Verbose         bool   // wheither to write any logs during generate
	InputOptions    *InputOptions

	Patches             []configpatcher.Patch
	PatchesControlPlane []configpatcher.Patch
	PatchesWorker       []configpatcher.Patch
}

// DefaultOptions returns default options.
func DefaultOptions() Options {
	return Options{
		Verbose: true,
	}
}

// WithExistingConfigs sets the path to existing config files.
func WithExistingConfigs(configPath string) Option {
	return func(o *Options) error {
		o.ExistingConfigs = configPath

		return nil
	}
}

// WithInputOptions allows passing in of various params for net-new input generation.
func WithInputOptions(inputOpts *InputOptions) Option {
	return func(o *Options) error {
		o.InputOptions = inputOpts

		return nil
	}
}

// WithVerbose allows setting verbose logging.
func WithVerbose(verbose bool) Option {
	return func(o *Options) error {
		o.Verbose = verbose

		return nil
	}
}

// WithJSONPatch allows patching every config in a bundle with a patch.
//
// Deprecated: use WithPatch instead.
func WithJSONPatch(patch jsonpatch.Patch) Option {
	return func(o *Options) error {
		o.Patches = append(o.Patches, patch)

		return nil
	}
}

// WithPatch allows patching every config in a bundle with a patch.
func WithPatch(patch []configpatcher.Patch) Option {
	return func(o *Options) error {
		o.Patches = append(o.Patches, patch...)

		return nil
	}
}

// WithJSONPatchControlPlane allows patching init and controlplane config in a bundle with a patch.
//
// Deprecated: use WithPatchControlPlane instead.
func WithJSONPatchControlPlane(patch jsonpatch.Patch) Option {
	return func(o *Options) error {
		o.PatchesControlPlane = append(o.PatchesControlPlane, patch)

		return nil
	}
}

// WithPatchControlPlane allows patching init and controlplane config in a bundle with a patch.
func WithPatchControlPlane(patch []configpatcher.Patch) Option {
	return func(o *Options) error {
		o.PatchesControlPlane = append(o.PatchesControlPlane, patch...)

		return nil
	}
}

// WithJSONPatchWorker allows patching worker config in a bundle with a patch.
//
// Deprecated: use WithPatchWorker instead.
func WithJSONPatchWorker(patch jsonpatch.Patch) Option {
	return func(o *Options) error {
		o.PatchesWorker = append(o.PatchesWorker, patch)

		return nil
	}
}

// WithPatchWorker allows patching worker config in a bundle with a patch.
func WithPatchWorker(patch []configpatcher.Patch) Option {
	return func(o *Options) error {
		o.PatchesWorker = append(o.PatchesWorker, patch...)

		return nil
	}
}
