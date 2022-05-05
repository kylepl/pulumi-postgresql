// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ``GrantRole`` resource creates and manages membership in a role to one or more other roles in a non-authoritative way.
//
// When using ``GrantRole`` resource it is likely because the PostgreSQL role you are modifying was created outside of this provider.
//
// > **Note:** This resource needs PostgreSQL version 9 or above.
type GrantRole struct {
	pulumi.CustomResourceState

	// The name of the role that is added to `role`.
	GrantRole pulumi.StringOutput `pulumi:"grantRole"`
	// The name of the role that is granted a new membership.
	Role pulumi.StringOutput `pulumi:"role"`
	// Giving ability to grant membership to others or not for `role`. (Default: false)
	WithAdminOption pulumi.BoolPtrOutput `pulumi:"withAdminOption"`
}

// NewGrantRole registers a new resource with the given unique name, arguments, and options.
func NewGrantRole(ctx *pulumi.Context,
	name string, args *GrantRoleArgs, opts ...pulumi.ResourceOption) (*GrantRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GrantRole == nil {
		return nil, errors.New("invalid value for required argument 'GrantRole'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource GrantRole
	err := ctx.RegisterResource("postgresql:index/grantRole:GrantRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrantRole gets an existing GrantRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrantRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantRoleState, opts ...pulumi.ResourceOption) (*GrantRole, error) {
	var resource GrantRole
	err := ctx.ReadResource("postgresql:index/grantRole:GrantRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GrantRole resources.
type grantRoleState struct {
	// The name of the role that is added to `role`.
	GrantRole *string `pulumi:"grantRole"`
	// The name of the role that is granted a new membership.
	Role *string `pulumi:"role"`
	// Giving ability to grant membership to others or not for `role`. (Default: false)
	WithAdminOption *bool `pulumi:"withAdminOption"`
}

type GrantRoleState struct {
	// The name of the role that is added to `role`.
	GrantRole pulumi.StringPtrInput
	// The name of the role that is granted a new membership.
	Role pulumi.StringPtrInput
	// Giving ability to grant membership to others or not for `role`. (Default: false)
	WithAdminOption pulumi.BoolPtrInput
}

func (GrantRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantRoleState)(nil)).Elem()
}

type grantRoleArgs struct {
	// The name of the role that is added to `role`.
	GrantRole string `pulumi:"grantRole"`
	// The name of the role that is granted a new membership.
	Role string `pulumi:"role"`
	// Giving ability to grant membership to others or not for `role`. (Default: false)
	WithAdminOption *bool `pulumi:"withAdminOption"`
}

// The set of arguments for constructing a GrantRole resource.
type GrantRoleArgs struct {
	// The name of the role that is added to `role`.
	GrantRole pulumi.StringInput
	// The name of the role that is granted a new membership.
	Role pulumi.StringInput
	// Giving ability to grant membership to others or not for `role`. (Default: false)
	WithAdminOption pulumi.BoolPtrInput
}

func (GrantRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantRoleArgs)(nil)).Elem()
}

type GrantRoleInput interface {
	pulumi.Input

	ToGrantRoleOutput() GrantRoleOutput
	ToGrantRoleOutputWithContext(ctx context.Context) GrantRoleOutput
}

func (*GrantRole) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantRole)(nil)).Elem()
}

func (i *GrantRole) ToGrantRoleOutput() GrantRoleOutput {
	return i.ToGrantRoleOutputWithContext(context.Background())
}

func (i *GrantRole) ToGrantRoleOutputWithContext(ctx context.Context) GrantRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantRoleOutput)
}

// GrantRoleArrayInput is an input type that accepts GrantRoleArray and GrantRoleArrayOutput values.
// You can construct a concrete instance of `GrantRoleArrayInput` via:
//
//          GrantRoleArray{ GrantRoleArgs{...} }
type GrantRoleArrayInput interface {
	pulumi.Input

	ToGrantRoleArrayOutput() GrantRoleArrayOutput
	ToGrantRoleArrayOutputWithContext(context.Context) GrantRoleArrayOutput
}

type GrantRoleArray []GrantRoleInput

func (GrantRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrantRole)(nil)).Elem()
}

func (i GrantRoleArray) ToGrantRoleArrayOutput() GrantRoleArrayOutput {
	return i.ToGrantRoleArrayOutputWithContext(context.Background())
}

func (i GrantRoleArray) ToGrantRoleArrayOutputWithContext(ctx context.Context) GrantRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantRoleArrayOutput)
}

// GrantRoleMapInput is an input type that accepts GrantRoleMap and GrantRoleMapOutput values.
// You can construct a concrete instance of `GrantRoleMapInput` via:
//
//          GrantRoleMap{ "key": GrantRoleArgs{...} }
type GrantRoleMapInput interface {
	pulumi.Input

	ToGrantRoleMapOutput() GrantRoleMapOutput
	ToGrantRoleMapOutputWithContext(context.Context) GrantRoleMapOutput
}

type GrantRoleMap map[string]GrantRoleInput

func (GrantRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrantRole)(nil)).Elem()
}

func (i GrantRoleMap) ToGrantRoleMapOutput() GrantRoleMapOutput {
	return i.ToGrantRoleMapOutputWithContext(context.Background())
}

func (i GrantRoleMap) ToGrantRoleMapOutputWithContext(ctx context.Context) GrantRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantRoleMapOutput)
}

type GrantRoleOutput struct{ *pulumi.OutputState }

func (GrantRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantRole)(nil)).Elem()
}

func (o GrantRoleOutput) ToGrantRoleOutput() GrantRoleOutput {
	return o
}

func (o GrantRoleOutput) ToGrantRoleOutputWithContext(ctx context.Context) GrantRoleOutput {
	return o
}

// The name of the role that is added to `role`.
func (o GrantRoleOutput) GrantRole() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantRole) pulumi.StringOutput { return v.GrantRole }).(pulumi.StringOutput)
}

// The name of the role that is granted a new membership.
func (o GrantRoleOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantRole) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Giving ability to grant membership to others or not for `role`. (Default: false)
func (o GrantRoleOutput) WithAdminOption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GrantRole) pulumi.BoolPtrOutput { return v.WithAdminOption }).(pulumi.BoolPtrOutput)
}

type GrantRoleArrayOutput struct{ *pulumi.OutputState }

func (GrantRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrantRole)(nil)).Elem()
}

func (o GrantRoleArrayOutput) ToGrantRoleArrayOutput() GrantRoleArrayOutput {
	return o
}

func (o GrantRoleArrayOutput) ToGrantRoleArrayOutputWithContext(ctx context.Context) GrantRoleArrayOutput {
	return o
}

func (o GrantRoleArrayOutput) Index(i pulumi.IntInput) GrantRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GrantRole {
		return vs[0].([]*GrantRole)[vs[1].(int)]
	}).(GrantRoleOutput)
}

type GrantRoleMapOutput struct{ *pulumi.OutputState }

func (GrantRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrantRole)(nil)).Elem()
}

func (o GrantRoleMapOutput) ToGrantRoleMapOutput() GrantRoleMapOutput {
	return o
}

func (o GrantRoleMapOutput) ToGrantRoleMapOutputWithContext(ctx context.Context) GrantRoleMapOutput {
	return o
}

func (o GrantRoleMapOutput) MapIndex(k pulumi.StringInput) GrantRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GrantRole {
		return vs[0].(map[string]*GrantRole)[vs[1].(string)]
	}).(GrantRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrantRoleInput)(nil)).Elem(), &GrantRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantRoleArrayInput)(nil)).Elem(), GrantRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantRoleMapInput)(nil)).Elem(), GrantRoleMap{})
	pulumi.RegisterOutputType(GrantRoleOutput{})
	pulumi.RegisterOutputType(GrantRoleArrayOutput{})
	pulumi.RegisterOutputType(GrantRoleMapOutput{})
}
