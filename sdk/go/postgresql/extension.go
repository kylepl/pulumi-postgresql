// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Extension`` resource creates and manages an extension on a PostgreSQL
// server.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/extension.html.markdown.
type Extension struct {
	s *pulumi.ResourceState
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOpt) (*Extension, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["database"] = nil
		inputs["name"] = nil
		inputs["schema"] = nil
		inputs["version"] = nil
	} else {
		inputs["database"] = args.Database
		inputs["name"] = args.Name
		inputs["schema"] = args.Schema
		inputs["version"] = args.Version
	}
	s, err := ctx.RegisterResource("postgresql:index/extension:Extension", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Extension{s: s}, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExtensionState, opts ...pulumi.ResourceOpt) (*Extension, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["database"] = state.Database
		inputs["name"] = state.Name
		inputs["schema"] = state.Schema
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("postgresql:index/extension:Extension", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Extension{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Extension) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Extension) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Which database to create the extension on. Defaults to provider database.
func (r *Extension) Database() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["database"])
}

// The name of the extension.
func (r *Extension) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Sets the schema of an extension.
func (r *Extension) Schema() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["schema"])
}

// Sets the version number of the extension.
func (r *Extension) Version() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Extension resources.
type ExtensionState struct {
	// Which database to create the extension on. Defaults to provider database.
	Database interface{}
	// The name of the extension.
	Name interface{}
	// Sets the schema of an extension.
	Schema interface{}
	// Sets the version number of the extension.
	Version interface{}
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// Which database to create the extension on. Defaults to provider database.
	Database interface{}
	// The name of the extension.
	Name interface{}
	// Sets the schema of an extension.
	Schema interface{}
	// Sets the version number of the extension.
	Version interface{}
}
