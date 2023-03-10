// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “Grant“ resource creates and manages privileges given to a user for a database schema.
//
// See [PostgreSQL documentation](https://www.postgresql.org/docs/current/sql-grant.html)
//
// > **Note:** This resource needs Postgresql version 9 or above.
//
// ## Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := postgresql.NewGrant(ctx, "readonlyTables", &postgresql.GrantArgs{
//				Database:   pulumi.String("test_db"),
//				ObjectType: pulumi.String("table"),
//				Objects: pulumi.StringArray{
//					pulumi.String("table1"),
//					pulumi.String("table2"),
//				},
//				Privileges: pulumi.StringArray{
//					pulumi.String("SELECT"),
//				},
//				Role:   pulumi.String("test_role"),
//				Schema: pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Examples
//
// Revoke default accesses for public schema:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := postgresql.NewGrant(ctx, "revokePublic", &postgresql.GrantArgs{
//				Database:   pulumi.String("test_db"),
//				ObjectType: pulumi.String("schema"),
//				Privileges: pulumi.StringArray{},
//				Role:       pulumi.String("public"),
//				Schema:     pulumi.String("public"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Grant struct {
	pulumi.CustomResourceState

	// The database to grant privileges on for this role.
	Database pulumi.StringOutput `pulumi:"database"`
	// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`.
	Objects pulumi.StringArrayOutput `pulumi:"objects"`
	// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
	Privileges pulumi.StringArrayOutput `pulumi:"privileges"`
	// The name of the role to grant privileges on, Set it to "public" for all roles.
	Role pulumi.StringOutput `pulumi:"role"`
	// The database schema to grant privileges on for this role (Required except if objectType is "database")
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
	WithGrantOption pulumi.BoolPtrOutput `pulumi:"withGrantOption"`
}

// NewGrant registers a new resource with the given unique name, arguments, and options.
func NewGrant(ctx *pulumi.Context,
	name string, args *GrantArgs, opts ...pulumi.ResourceOption) (*Grant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.ObjectType == nil {
		return nil, errors.New("invalid value for required argument 'ObjectType'")
	}
	if args.Privileges == nil {
		return nil, errors.New("invalid value for required argument 'Privileges'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource Grant
	err := ctx.RegisterResource("postgresql:index/grant:Grant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrant gets an existing Grant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantState, opts ...pulumi.ResourceOption) (*Grant, error) {
	var resource Grant
	err := ctx.ReadResource("postgresql:index/grant:Grant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Grant resources.
type grantState struct {
	// The database to grant privileges on for this role.
	Database *string `pulumi:"database"`
	// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
	ObjectType *string `pulumi:"objectType"`
	// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`.
	Objects []string `pulumi:"objects"`
	// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to grant privileges on, Set it to "public" for all roles.
	Role *string `pulumi:"role"`
	// The database schema to grant privileges on for this role (Required except if objectType is "database")
	Schema *string `pulumi:"schema"`
	// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type GrantState struct {
	// The database to grant privileges on for this role.
	Database pulumi.StringPtrInput
	// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
	ObjectType pulumi.StringPtrInput
	// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`.
	Objects pulumi.StringArrayInput
	// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
	Privileges pulumi.StringArrayInput
	// The name of the role to grant privileges on, Set it to "public" for all roles.
	Role pulumi.StringPtrInput
	// The database schema to grant privileges on for this role (Required except if objectType is "database")
	Schema pulumi.StringPtrInput
	// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
	WithGrantOption pulumi.BoolPtrInput
}

func (GrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantState)(nil)).Elem()
}

type grantArgs struct {
	// The database to grant privileges on for this role.
	Database string `pulumi:"database"`
	// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
	ObjectType string `pulumi:"objectType"`
	// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`.
	Objects []string `pulumi:"objects"`
	// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to grant privileges on, Set it to "public" for all roles.
	Role string `pulumi:"role"`
	// The database schema to grant privileges on for this role (Required except if objectType is "database")
	Schema *string `pulumi:"schema"`
	// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a Grant resource.
type GrantArgs struct {
	// The database to grant privileges on for this role.
	Database pulumi.StringInput
	// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
	ObjectType pulumi.StringInput
	// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`.
	Objects pulumi.StringArrayInput
	// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
	Privileges pulumi.StringArrayInput
	// The name of the role to grant privileges on, Set it to "public" for all roles.
	Role pulumi.StringInput
	// The database schema to grant privileges on for this role (Required except if objectType is "database")
	Schema pulumi.StringPtrInput
	// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
	WithGrantOption pulumi.BoolPtrInput
}

func (GrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantArgs)(nil)).Elem()
}

type GrantInput interface {
	pulumi.Input

	ToGrantOutput() GrantOutput
	ToGrantOutputWithContext(ctx context.Context) GrantOutput
}

func (*Grant) ElementType() reflect.Type {
	return reflect.TypeOf((**Grant)(nil)).Elem()
}

func (i *Grant) ToGrantOutput() GrantOutput {
	return i.ToGrantOutputWithContext(context.Background())
}

func (i *Grant) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantOutput)
}

// GrantArrayInput is an input type that accepts GrantArray and GrantArrayOutput values.
// You can construct a concrete instance of `GrantArrayInput` via:
//
//	GrantArray{ GrantArgs{...} }
type GrantArrayInput interface {
	pulumi.Input

	ToGrantArrayOutput() GrantArrayOutput
	ToGrantArrayOutputWithContext(context.Context) GrantArrayOutput
}

type GrantArray []GrantInput

func (GrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Grant)(nil)).Elem()
}

func (i GrantArray) ToGrantArrayOutput() GrantArrayOutput {
	return i.ToGrantArrayOutputWithContext(context.Background())
}

func (i GrantArray) ToGrantArrayOutputWithContext(ctx context.Context) GrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantArrayOutput)
}

// GrantMapInput is an input type that accepts GrantMap and GrantMapOutput values.
// You can construct a concrete instance of `GrantMapInput` via:
//
//	GrantMap{ "key": GrantArgs{...} }
type GrantMapInput interface {
	pulumi.Input

	ToGrantMapOutput() GrantMapOutput
	ToGrantMapOutputWithContext(context.Context) GrantMapOutput
}

type GrantMap map[string]GrantInput

func (GrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Grant)(nil)).Elem()
}

func (i GrantMap) ToGrantMapOutput() GrantMapOutput {
	return i.ToGrantMapOutputWithContext(context.Background())
}

func (i GrantMap) ToGrantMapOutputWithContext(ctx context.Context) GrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantMapOutput)
}

type GrantOutput struct{ *pulumi.OutputState }

func (GrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Grant)(nil)).Elem()
}

func (o GrantOutput) ToGrantOutput() GrantOutput {
	return o
}

func (o GrantOutput) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return o
}

// The database to grant privileges on for this role.
func (o GrantOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// The PostgreSQL object type to grant the privileges on (one of: database, schema, table, sequence, function, procedure, routine, foreign_data_wrapper, foreign_server).
func (o GrantOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringOutput { return v.ObjectType }).(pulumi.StringOutput)
}

// The objects upon which to grant the privileges. An empty list (the default) means to grant permissions on *all* objects of the specified type. You cannot specify this option if the `objectType` is `database` or `schema`.
func (o GrantOutput) Objects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringArrayOutput { return v.Objects }).(pulumi.StringArrayOutput)
}

// The list of privileges to grant. There are different kinds of privileges: SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER, CREATE, CONNECT, TEMPORARY, EXECUTE, and USAGE. An empty list could be provided to revoke all privileges for this role.
func (o GrantOutput) Privileges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringArrayOutput { return v.Privileges }).(pulumi.StringArrayOutput)
}

// The name of the role to grant privileges on, Set it to "public" for all roles.
func (o GrantOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The database schema to grant privileges on for this role (Required except if objectType is "database")
func (o GrantOutput) Schema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Grant) pulumi.StringPtrOutput { return v.Schema }).(pulumi.StringPtrOutput)
}

// Whether the recipient of these privileges can grant the same privileges to others. Defaults to false.
func (o GrantOutput) WithGrantOption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Grant) pulumi.BoolPtrOutput { return v.WithGrantOption }).(pulumi.BoolPtrOutput)
}

type GrantArrayOutput struct{ *pulumi.OutputState }

func (GrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Grant)(nil)).Elem()
}

func (o GrantArrayOutput) ToGrantArrayOutput() GrantArrayOutput {
	return o
}

func (o GrantArrayOutput) ToGrantArrayOutputWithContext(ctx context.Context) GrantArrayOutput {
	return o
}

func (o GrantArrayOutput) Index(i pulumi.IntInput) GrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Grant {
		return vs[0].([]*Grant)[vs[1].(int)]
	}).(GrantOutput)
}

type GrantMapOutput struct{ *pulumi.OutputState }

func (GrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Grant)(nil)).Elem()
}

func (o GrantMapOutput) ToGrantMapOutput() GrantMapOutput {
	return o
}

func (o GrantMapOutput) ToGrantMapOutputWithContext(ctx context.Context) GrantMapOutput {
	return o
}

func (o GrantMapOutput) MapIndex(k pulumi.StringInput) GrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Grant {
		return vs[0].(map[string]*Grant)[vs[1].(string)]
	}).(GrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrantInput)(nil)).Elem(), &Grant{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantArrayInput)(nil)).Elem(), GrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantMapInput)(nil)).Elem(), GrantMap{})
	pulumi.RegisterOutputType(GrantOutput{})
	pulumi.RegisterOutputType(GrantArrayOutput{})
	pulumi.RegisterOutputType(GrantMapOutput{})
}
