// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The “DefaultPrivileges“ resource creates and manages default privileges given to a user for a database schema.
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
//			_, err := postgresql.NewDefaultPrivileges(ctx, "readOnlyTables", &postgresql.DefaultPrivilegesArgs{
//				Database:   pulumi.String("test_db"),
//				ObjectType: pulumi.String("table"),
//				Owner:      pulumi.String("db_owner"),
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
// Revoke default privileges for functions for "public" role:
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
//			_, err := postgresql.NewDefaultPrivileges(ctx, "revokePublic", &postgresql.DefaultPrivilegesArgs{
//				Database:   pulumi.Any(postgresql_database.Example_db.Name),
//				Role:       pulumi.String("public"),
//				Owner:      pulumi.String("object_owner"),
//				ObjectType: pulumi.String("function"),
//				Privileges: pulumi.StringArray{},
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
// Deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges
type DefaultPrivileg struct {
	pulumi.CustomResourceState

	// The database to grant default privileges for this role.
	Database pulumi.StringOutput `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges pulumi.StringArrayOutput `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role pulumi.StringOutput `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// Permit the grant recipient to grant it to others
	WithGrantOption pulumi.BoolPtrOutput `pulumi:"withGrantOption"`
}

// NewDefaultPrivileg registers a new resource with the given unique name, arguments, and options.
func NewDefaultPrivileg(ctx *pulumi.Context,
	name string, args *DefaultPrivilegArgs, opts ...pulumi.ResourceOption) (*DefaultPrivileg, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.ObjectType == nil {
		return nil, errors.New("invalid value for required argument 'ObjectType'")
	}
	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	if args.Privileges == nil {
		return nil, errors.New("invalid value for required argument 'Privileges'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultPrivileg
	err := ctx.RegisterResource("postgresql:index/defaultPrivileg:DefaultPrivileg", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultPrivileg gets an existing DefaultPrivileg resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultPrivileg(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultPrivilegState, opts ...pulumi.ResourceOption) (*DefaultPrivileg, error) {
	var resource DefaultPrivileg
	err := ctx.ReadResource("postgresql:index/defaultPrivileg:DefaultPrivileg", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultPrivileg resources.
type defaultPrivilegState struct {
	// The database to grant default privileges for this role.
	Database *string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType *string `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner *string `pulumi:"owner"`
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role *string `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema *string `pulumi:"schema"`
	// Permit the grant recipient to grant it to others
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type DefaultPrivilegState struct {
	// The database to grant default privileges for this role.
	Database pulumi.StringPtrInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType pulumi.StringPtrInput
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringPtrInput
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on.
	Role pulumi.StringPtrInput
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrInput
	// Permit the grant recipient to grant it to others
	WithGrantOption pulumi.BoolPtrInput
}

func (DefaultPrivilegState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegState)(nil)).Elem()
}

type defaultPrivilegArgs struct {
	// The database to grant default privileges for this role.
	Database string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType string `pulumi:"objectType"`
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner string `pulumi:"owner"`
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on.
	Role string `pulumi:"role"`
	// The database schema to set default privileges for this role.
	Schema *string `pulumi:"schema"`
	// Permit the grant recipient to grant it to others
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a DefaultPrivileg resource.
type DefaultPrivilegArgs struct {
	// The database to grant default privileges for this role.
	Database pulumi.StringInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
	ObjectType pulumi.StringInput
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner pulumi.StringInput
	// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on.
	Role pulumi.StringInput
	// The database schema to set default privileges for this role.
	Schema pulumi.StringPtrInput
	// Permit the grant recipient to grant it to others
	WithGrantOption pulumi.BoolPtrInput
}

func (DefaultPrivilegArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegArgs)(nil)).Elem()
}

type DefaultPrivilegInput interface {
	pulumi.Input

	ToDefaultPrivilegOutput() DefaultPrivilegOutput
	ToDefaultPrivilegOutputWithContext(ctx context.Context) DefaultPrivilegOutput
}

func (*DefaultPrivileg) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPrivileg)(nil)).Elem()
}

func (i *DefaultPrivileg) ToDefaultPrivilegOutput() DefaultPrivilegOutput {
	return i.ToDefaultPrivilegOutputWithContext(context.Background())
}

func (i *DefaultPrivileg) ToDefaultPrivilegOutputWithContext(ctx context.Context) DefaultPrivilegOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegOutput)
}

func (i *DefaultPrivileg) ToOutput(ctx context.Context) pulumix.Output[*DefaultPrivileg] {
	return pulumix.Output[*DefaultPrivileg]{
		OutputState: i.ToDefaultPrivilegOutputWithContext(ctx).OutputState,
	}
}

// DefaultPrivilegArrayInput is an input type that accepts DefaultPrivilegArray and DefaultPrivilegArrayOutput values.
// You can construct a concrete instance of `DefaultPrivilegArrayInput` via:
//
//	DefaultPrivilegArray{ DefaultPrivilegArgs{...} }
type DefaultPrivilegArrayInput interface {
	pulumi.Input

	ToDefaultPrivilegArrayOutput() DefaultPrivilegArrayOutput
	ToDefaultPrivilegArrayOutputWithContext(context.Context) DefaultPrivilegArrayOutput
}

type DefaultPrivilegArray []DefaultPrivilegInput

func (DefaultPrivilegArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPrivileg)(nil)).Elem()
}

func (i DefaultPrivilegArray) ToDefaultPrivilegArrayOutput() DefaultPrivilegArrayOutput {
	return i.ToDefaultPrivilegArrayOutputWithContext(context.Background())
}

func (i DefaultPrivilegArray) ToDefaultPrivilegArrayOutputWithContext(ctx context.Context) DefaultPrivilegArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegArrayOutput)
}

func (i DefaultPrivilegArray) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultPrivileg] {
	return pulumix.Output[[]*DefaultPrivileg]{
		OutputState: i.ToDefaultPrivilegArrayOutputWithContext(ctx).OutputState,
	}
}

// DefaultPrivilegMapInput is an input type that accepts DefaultPrivilegMap and DefaultPrivilegMapOutput values.
// You can construct a concrete instance of `DefaultPrivilegMapInput` via:
//
//	DefaultPrivilegMap{ "key": DefaultPrivilegArgs{...} }
type DefaultPrivilegMapInput interface {
	pulumi.Input

	ToDefaultPrivilegMapOutput() DefaultPrivilegMapOutput
	ToDefaultPrivilegMapOutputWithContext(context.Context) DefaultPrivilegMapOutput
}

type DefaultPrivilegMap map[string]DefaultPrivilegInput

func (DefaultPrivilegMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPrivileg)(nil)).Elem()
}

func (i DefaultPrivilegMap) ToDefaultPrivilegMapOutput() DefaultPrivilegMapOutput {
	return i.ToDefaultPrivilegMapOutputWithContext(context.Background())
}

func (i DefaultPrivilegMap) ToDefaultPrivilegMapOutputWithContext(ctx context.Context) DefaultPrivilegMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultPrivilegMapOutput)
}

func (i DefaultPrivilegMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultPrivileg] {
	return pulumix.Output[map[string]*DefaultPrivileg]{
		OutputState: i.ToDefaultPrivilegMapOutputWithContext(ctx).OutputState,
	}
}

type DefaultPrivilegOutput struct{ *pulumi.OutputState }

func (DefaultPrivilegOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPrivileg)(nil)).Elem()
}

func (o DefaultPrivilegOutput) ToDefaultPrivilegOutput() DefaultPrivilegOutput {
	return o
}

func (o DefaultPrivilegOutput) ToDefaultPrivilegOutputWithContext(ctx context.Context) DefaultPrivilegOutput {
	return o
}

func (o DefaultPrivilegOutput) ToOutput(ctx context.Context) pulumix.Output[*DefaultPrivileg] {
	return pulumix.Output[*DefaultPrivileg]{
		OutputState: o.OutputState,
	}
}

// The database to grant default privileges for this role.
func (o DefaultPrivilegOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type, schema).
func (o DefaultPrivilegOutput) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.StringOutput { return v.ObjectType }).(pulumi.StringOutput)
}

// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
func (o DefaultPrivilegOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The list of privileges to apply as default privileges. An empty list could be provided to revoke all default privileges for this role.
func (o DefaultPrivilegOutput) Privileges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.StringArrayOutput { return v.Privileges }).(pulumi.StringArrayOutput)
}

// The name of the role to which grant default privileges on.
func (o DefaultPrivilegOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The database schema to set default privileges for this role.
func (o DefaultPrivilegOutput) Schema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.StringPtrOutput { return v.Schema }).(pulumi.StringPtrOutput)
}

// Permit the grant recipient to grant it to others
func (o DefaultPrivilegOutput) WithGrantOption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultPrivileg) pulumi.BoolPtrOutput { return v.WithGrantOption }).(pulumi.BoolPtrOutput)
}

type DefaultPrivilegArrayOutput struct{ *pulumi.OutputState }

func (DefaultPrivilegArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultPrivileg)(nil)).Elem()
}

func (o DefaultPrivilegArrayOutput) ToDefaultPrivilegArrayOutput() DefaultPrivilegArrayOutput {
	return o
}

func (o DefaultPrivilegArrayOutput) ToDefaultPrivilegArrayOutputWithContext(ctx context.Context) DefaultPrivilegArrayOutput {
	return o
}

func (o DefaultPrivilegArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultPrivileg] {
	return pulumix.Output[[]*DefaultPrivileg]{
		OutputState: o.OutputState,
	}
}

func (o DefaultPrivilegArrayOutput) Index(i pulumi.IntInput) DefaultPrivilegOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultPrivileg {
		return vs[0].([]*DefaultPrivileg)[vs[1].(int)]
	}).(DefaultPrivilegOutput)
}

type DefaultPrivilegMapOutput struct{ *pulumi.OutputState }

func (DefaultPrivilegMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultPrivileg)(nil)).Elem()
}

func (o DefaultPrivilegMapOutput) ToDefaultPrivilegMapOutput() DefaultPrivilegMapOutput {
	return o
}

func (o DefaultPrivilegMapOutput) ToDefaultPrivilegMapOutputWithContext(ctx context.Context) DefaultPrivilegMapOutput {
	return o
}

func (o DefaultPrivilegMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultPrivileg] {
	return pulumix.Output[map[string]*DefaultPrivileg]{
		OutputState: o.OutputState,
	}
}

func (o DefaultPrivilegMapOutput) MapIndex(k pulumi.StringInput) DefaultPrivilegOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultPrivileg {
		return vs[0].(map[string]*DefaultPrivileg)[vs[1].(string)]
	}).(DefaultPrivilegOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrivilegInput)(nil)).Elem(), &DefaultPrivileg{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrivilegArrayInput)(nil)).Elem(), DefaultPrivilegArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultPrivilegMapInput)(nil)).Elem(), DefaultPrivilegMap{})
	pulumi.RegisterOutputType(DefaultPrivilegOutput{})
	pulumi.RegisterOutputType(DefaultPrivilegArrayOutput{})
	pulumi.RegisterOutputType(DefaultPrivilegMapOutput{})
}
