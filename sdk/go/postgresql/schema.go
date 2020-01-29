// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package postgresql

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Schema`` resource creates and manages [schema
// objects](https://www.postgresql.org/docs/current/static/ddl-schemas.html) within
// a PostgreSQL database.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/schema.html.markdown.
type Schema struct {
	pulumi.CustomResourceState

	// The database name to alter schema
	Database pulumi.StringOutput `pulumi:"database"`
	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	DropCascade pulumi.BoolPtrOutput `pulumi:"dropCascade"`
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists pulumi.BoolPtrOutput `pulumi:"ifNotExists"`
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ROLE who owns the schema.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies SchemaPolicyArrayOutput `pulumi:"policies"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		args = &SchemaArgs{}
	}
	var resource Schema
	err := ctx.RegisterResource("postgresql:index/schema:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("postgresql:index/schema:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
	// The database name to alter schema
	Database *string `pulumi:"database"`
	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	DropCascade *bool `pulumi:"dropCascade"`
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists *bool `pulumi:"ifNotExists"`
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name *string `pulumi:"name"`
	// The ROLE who owns the schema.
	Owner *string `pulumi:"owner"`
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies []SchemaPolicy `pulumi:"policies"`
}

type SchemaState struct {
	// The database name to alter schema
	Database pulumi.StringPtrInput
	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	DropCascade pulumi.BoolPtrInput
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists pulumi.BoolPtrInput
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name pulumi.StringPtrInput
	// The ROLE who owns the schema.
	Owner pulumi.StringPtrInput
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies SchemaPolicyArrayInput
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	// The database name to alter schema
	Database *string `pulumi:"database"`
	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	DropCascade *bool `pulumi:"dropCascade"`
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists *bool `pulumi:"ifNotExists"`
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name *string `pulumi:"name"`
	// The ROLE who owns the schema.
	Owner *string `pulumi:"owner"`
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies []SchemaPolicy `pulumi:"policies"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// The database name to alter schema
	Database pulumi.StringPtrInput
	// When true, will also drop all the objects that are contained in the schema. (Default: false)
	DropCascade pulumi.BoolPtrInput
	// When true, use the existing schema if it exists. (Default: true)
	IfNotExists pulumi.BoolPtrInput
	// The name of the schema. Must be unique in the PostgreSQL
	// database instance where it is configured.
	Name pulumi.StringPtrInput
	// The ROLE who owns the schema.
	Owner pulumi.StringPtrInput
	// Can be specified multiple times for each policy.  Each
	// policy block supports fields documented below.
	Policies SchemaPolicyArrayInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

