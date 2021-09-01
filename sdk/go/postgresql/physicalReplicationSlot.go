// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ``PhysicalReplicationSlot`` resource creates and manages a physical replication slot on a PostgreSQL
// server. This is useful to setup a cross datacenter replication, with Patroni for example, or permit
// any stand-by cluster to replicate physically data.
type PhysicalReplicationSlot struct {
	pulumi.CustomResourceState

	// The name of the replication slot.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewPhysicalReplicationSlot registers a new resource with the given unique name, arguments, and options.
func NewPhysicalReplicationSlot(ctx *pulumi.Context,
	name string, args *PhysicalReplicationSlotArgs, opts ...pulumi.ResourceOption) (*PhysicalReplicationSlot, error) {
	if args == nil {
		args = &PhysicalReplicationSlotArgs{}
	}

	var resource PhysicalReplicationSlot
	err := ctx.RegisterResource("postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPhysicalReplicationSlot gets an existing PhysicalReplicationSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPhysicalReplicationSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PhysicalReplicationSlotState, opts ...pulumi.ResourceOption) (*PhysicalReplicationSlot, error) {
	var resource PhysicalReplicationSlot
	err := ctx.ReadResource("postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PhysicalReplicationSlot resources.
type physicalReplicationSlotState struct {
	// The name of the replication slot.
	Name *string `pulumi:"name"`
}

type PhysicalReplicationSlotState struct {
	// The name of the replication slot.
	Name pulumi.StringPtrInput
}

func (PhysicalReplicationSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*physicalReplicationSlotState)(nil)).Elem()
}

type physicalReplicationSlotArgs struct {
	// The name of the replication slot.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a PhysicalReplicationSlot resource.
type PhysicalReplicationSlotArgs struct {
	// The name of the replication slot.
	Name pulumi.StringPtrInput
}

func (PhysicalReplicationSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*physicalReplicationSlotArgs)(nil)).Elem()
}

type PhysicalReplicationSlotInput interface {
	pulumi.Input

	ToPhysicalReplicationSlotOutput() PhysicalReplicationSlotOutput
	ToPhysicalReplicationSlotOutputWithContext(ctx context.Context) PhysicalReplicationSlotOutput
}

func (*PhysicalReplicationSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*PhysicalReplicationSlot)(nil))
}

func (i *PhysicalReplicationSlot) ToPhysicalReplicationSlotOutput() PhysicalReplicationSlotOutput {
	return i.ToPhysicalReplicationSlotOutputWithContext(context.Background())
}

func (i *PhysicalReplicationSlot) ToPhysicalReplicationSlotOutputWithContext(ctx context.Context) PhysicalReplicationSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalReplicationSlotOutput)
}

func (i *PhysicalReplicationSlot) ToPhysicalReplicationSlotPtrOutput() PhysicalReplicationSlotPtrOutput {
	return i.ToPhysicalReplicationSlotPtrOutputWithContext(context.Background())
}

func (i *PhysicalReplicationSlot) ToPhysicalReplicationSlotPtrOutputWithContext(ctx context.Context) PhysicalReplicationSlotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalReplicationSlotPtrOutput)
}

type PhysicalReplicationSlotPtrInput interface {
	pulumi.Input

	ToPhysicalReplicationSlotPtrOutput() PhysicalReplicationSlotPtrOutput
	ToPhysicalReplicationSlotPtrOutputWithContext(ctx context.Context) PhysicalReplicationSlotPtrOutput
}

type physicalReplicationSlotPtrType PhysicalReplicationSlotArgs

func (*physicalReplicationSlotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PhysicalReplicationSlot)(nil))
}

func (i *physicalReplicationSlotPtrType) ToPhysicalReplicationSlotPtrOutput() PhysicalReplicationSlotPtrOutput {
	return i.ToPhysicalReplicationSlotPtrOutputWithContext(context.Background())
}

func (i *physicalReplicationSlotPtrType) ToPhysicalReplicationSlotPtrOutputWithContext(ctx context.Context) PhysicalReplicationSlotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalReplicationSlotPtrOutput)
}

// PhysicalReplicationSlotArrayInput is an input type that accepts PhysicalReplicationSlotArray and PhysicalReplicationSlotArrayOutput values.
// You can construct a concrete instance of `PhysicalReplicationSlotArrayInput` via:
//
//          PhysicalReplicationSlotArray{ PhysicalReplicationSlotArgs{...} }
type PhysicalReplicationSlotArrayInput interface {
	pulumi.Input

	ToPhysicalReplicationSlotArrayOutput() PhysicalReplicationSlotArrayOutput
	ToPhysicalReplicationSlotArrayOutputWithContext(context.Context) PhysicalReplicationSlotArrayOutput
}

type PhysicalReplicationSlotArray []PhysicalReplicationSlotInput

func (PhysicalReplicationSlotArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PhysicalReplicationSlot)(nil))
}

func (i PhysicalReplicationSlotArray) ToPhysicalReplicationSlotArrayOutput() PhysicalReplicationSlotArrayOutput {
	return i.ToPhysicalReplicationSlotArrayOutputWithContext(context.Background())
}

func (i PhysicalReplicationSlotArray) ToPhysicalReplicationSlotArrayOutputWithContext(ctx context.Context) PhysicalReplicationSlotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalReplicationSlotArrayOutput)
}

// PhysicalReplicationSlotMapInput is an input type that accepts PhysicalReplicationSlotMap and PhysicalReplicationSlotMapOutput values.
// You can construct a concrete instance of `PhysicalReplicationSlotMapInput` via:
//
//          PhysicalReplicationSlotMap{ "key": PhysicalReplicationSlotArgs{...} }
type PhysicalReplicationSlotMapInput interface {
	pulumi.Input

	ToPhysicalReplicationSlotMapOutput() PhysicalReplicationSlotMapOutput
	ToPhysicalReplicationSlotMapOutputWithContext(context.Context) PhysicalReplicationSlotMapOutput
}

type PhysicalReplicationSlotMap map[string]PhysicalReplicationSlotInput

func (PhysicalReplicationSlotMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PhysicalReplicationSlot)(nil))
}

func (i PhysicalReplicationSlotMap) ToPhysicalReplicationSlotMapOutput() PhysicalReplicationSlotMapOutput {
	return i.ToPhysicalReplicationSlotMapOutputWithContext(context.Background())
}

func (i PhysicalReplicationSlotMap) ToPhysicalReplicationSlotMapOutputWithContext(ctx context.Context) PhysicalReplicationSlotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhysicalReplicationSlotMapOutput)
}

type PhysicalReplicationSlotOutput struct {
	*pulumi.OutputState
}

func (PhysicalReplicationSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PhysicalReplicationSlot)(nil))
}

func (o PhysicalReplicationSlotOutput) ToPhysicalReplicationSlotOutput() PhysicalReplicationSlotOutput {
	return o
}

func (o PhysicalReplicationSlotOutput) ToPhysicalReplicationSlotOutputWithContext(ctx context.Context) PhysicalReplicationSlotOutput {
	return o
}

func (o PhysicalReplicationSlotOutput) ToPhysicalReplicationSlotPtrOutput() PhysicalReplicationSlotPtrOutput {
	return o.ToPhysicalReplicationSlotPtrOutputWithContext(context.Background())
}

func (o PhysicalReplicationSlotOutput) ToPhysicalReplicationSlotPtrOutputWithContext(ctx context.Context) PhysicalReplicationSlotPtrOutput {
	return o.ApplyT(func(v PhysicalReplicationSlot) *PhysicalReplicationSlot {
		return &v
	}).(PhysicalReplicationSlotPtrOutput)
}

type PhysicalReplicationSlotPtrOutput struct {
	*pulumi.OutputState
}

func (PhysicalReplicationSlotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhysicalReplicationSlot)(nil))
}

func (o PhysicalReplicationSlotPtrOutput) ToPhysicalReplicationSlotPtrOutput() PhysicalReplicationSlotPtrOutput {
	return o
}

func (o PhysicalReplicationSlotPtrOutput) ToPhysicalReplicationSlotPtrOutputWithContext(ctx context.Context) PhysicalReplicationSlotPtrOutput {
	return o
}

type PhysicalReplicationSlotArrayOutput struct{ *pulumi.OutputState }

func (PhysicalReplicationSlotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PhysicalReplicationSlot)(nil))
}

func (o PhysicalReplicationSlotArrayOutput) ToPhysicalReplicationSlotArrayOutput() PhysicalReplicationSlotArrayOutput {
	return o
}

func (o PhysicalReplicationSlotArrayOutput) ToPhysicalReplicationSlotArrayOutputWithContext(ctx context.Context) PhysicalReplicationSlotArrayOutput {
	return o
}

func (o PhysicalReplicationSlotArrayOutput) Index(i pulumi.IntInput) PhysicalReplicationSlotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PhysicalReplicationSlot {
		return vs[0].([]PhysicalReplicationSlot)[vs[1].(int)]
	}).(PhysicalReplicationSlotOutput)
}

type PhysicalReplicationSlotMapOutput struct{ *pulumi.OutputState }

func (PhysicalReplicationSlotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PhysicalReplicationSlot)(nil))
}

func (o PhysicalReplicationSlotMapOutput) ToPhysicalReplicationSlotMapOutput() PhysicalReplicationSlotMapOutput {
	return o
}

func (o PhysicalReplicationSlotMapOutput) ToPhysicalReplicationSlotMapOutputWithContext(ctx context.Context) PhysicalReplicationSlotMapOutput {
	return o
}

func (o PhysicalReplicationSlotMapOutput) MapIndex(k pulumi.StringInput) PhysicalReplicationSlotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PhysicalReplicationSlot {
		return vs[0].(map[string]PhysicalReplicationSlot)[vs[1].(string)]
	}).(PhysicalReplicationSlotOutput)
}

func init() {
	pulumi.RegisterOutputType(PhysicalReplicationSlotOutput{})
	pulumi.RegisterOutputType(PhysicalReplicationSlotPtrOutput{})
	pulumi.RegisterOutputType(PhysicalReplicationSlotArrayOutput{})
	pulumi.RegisterOutputType(PhysicalReplicationSlotMapOutput{})
}
