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

// The `Subscription` resource creates and manages a publication on a PostgreSQL
// server.
//
// ## Postgres documentation
//
// - https://www.postgresql.org/docs/current/sql-createsubscription.html
type Subscription struct {
	pulumi.CustomResourceState

	// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
	Conninfo pulumi.StringOutput `pulumi:"conninfo"`
	// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
	CreateSlot pulumi.BoolPtrOutput `pulumi:"createSlot"`
	// Which database to create the subscription on. Defaults to provider database.
	Database pulumi.StringOutput `pulumi:"database"`
	// The name of the publication.
	Name pulumi.StringOutput `pulumi:"name"`
	// Names of the publications on the publisher to subscribe to
	Publications pulumi.StringArrayOutput `pulumi:"publications"`
	// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
	SlotName pulumi.StringPtrOutput `pulumi:"slotName"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Conninfo == nil {
		return nil, errors.New("invalid value for required argument 'Conninfo'")
	}
	if args.Publications == nil {
		return nil, errors.New("invalid value for required argument 'Publications'")
	}
	if args.Conninfo != nil {
		args.Conninfo = pulumi.ToSecret(args.Conninfo).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"conninfo",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Subscription
	err := ctx.RegisterResource("postgresql:index/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("postgresql:index/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
	Conninfo *string `pulumi:"conninfo"`
	// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
	CreateSlot *bool `pulumi:"createSlot"`
	// Which database to create the subscription on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// The name of the publication.
	Name *string `pulumi:"name"`
	// Names of the publications on the publisher to subscribe to
	Publications []string `pulumi:"publications"`
	// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
	SlotName *string `pulumi:"slotName"`
}

type SubscriptionState struct {
	// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
	Conninfo pulumi.StringPtrInput
	// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
	CreateSlot pulumi.BoolPtrInput
	// Which database to create the subscription on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// The name of the publication.
	Name pulumi.StringPtrInput
	// Names of the publications on the publisher to subscribe to
	Publications pulumi.StringArrayInput
	// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
	SlotName pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
	Conninfo string `pulumi:"conninfo"`
	// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
	CreateSlot *bool `pulumi:"createSlot"`
	// Which database to create the subscription on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// The name of the publication.
	Name *string `pulumi:"name"`
	// Names of the publications on the publisher to subscribe to
	Publications []string `pulumi:"publications"`
	// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
	SlotName *string `pulumi:"slotName"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
	Conninfo pulumi.StringInput
	// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
	CreateSlot pulumi.BoolPtrInput
	// Which database to create the subscription on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// The name of the publication.
	Name pulumi.StringPtrInput
	// Names of the publications on the publisher to subscribe to
	Publications pulumi.StringArrayInput
	// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
	SlotName pulumi.StringPtrInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

func (i *Subscription) ToOutput(ctx context.Context) pulumix.Output[*Subscription] {
	return pulumix.Output[*Subscription]{
		OutputState: i.ToSubscriptionOutputWithContext(ctx).OutputState,
	}
}

// SubscriptionArrayInput is an input type that accepts SubscriptionArray and SubscriptionArrayOutput values.
// You can construct a concrete instance of `SubscriptionArrayInput` via:
//
//	SubscriptionArray{ SubscriptionArgs{...} }
type SubscriptionArrayInput interface {
	pulumi.Input

	ToSubscriptionArrayOutput() SubscriptionArrayOutput
	ToSubscriptionArrayOutputWithContext(context.Context) SubscriptionArrayOutput
}

type SubscriptionArray []SubscriptionInput

func (SubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subscription)(nil)).Elem()
}

func (i SubscriptionArray) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return i.ToSubscriptionArrayOutputWithContext(context.Background())
}

func (i SubscriptionArray) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionArrayOutput)
}

func (i SubscriptionArray) ToOutput(ctx context.Context) pulumix.Output[[]*Subscription] {
	return pulumix.Output[[]*Subscription]{
		OutputState: i.ToSubscriptionArrayOutputWithContext(ctx).OutputState,
	}
}

// SubscriptionMapInput is an input type that accepts SubscriptionMap and SubscriptionMapOutput values.
// You can construct a concrete instance of `SubscriptionMapInput` via:
//
//	SubscriptionMap{ "key": SubscriptionArgs{...} }
type SubscriptionMapInput interface {
	pulumi.Input

	ToSubscriptionMapOutput() SubscriptionMapOutput
	ToSubscriptionMapOutputWithContext(context.Context) SubscriptionMapOutput
}

type SubscriptionMap map[string]SubscriptionInput

func (SubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subscription)(nil)).Elem()
}

func (i SubscriptionMap) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return i.ToSubscriptionMapOutputWithContext(context.Background())
}

func (i SubscriptionMap) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionMapOutput)
}

func (i SubscriptionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Subscription] {
	return pulumix.Output[map[string]*Subscription]{
		OutputState: i.ToSubscriptionMapOutputWithContext(ctx).OutputState,
	}
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*Subscription] {
	return pulumix.Output[*Subscription]{
		OutputState: o.OutputState,
	}
}

// The connection string to the publisher. It should follow the [keyword/value format](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)
func (o SubscriptionOutput) Conninfo() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Conninfo }).(pulumi.StringOutput)
}

// Specifies whether the command should create the replication slot on the publisher. Default behavior is true
func (o SubscriptionOutput) CreateSlot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.CreateSlot }).(pulumi.BoolPtrOutput)
}

// Which database to create the subscription on. Defaults to provider database.
func (o SubscriptionOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

// The name of the publication.
func (o SubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Names of the publications on the publisher to subscribe to
func (o SubscriptionOutput) Publications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringArrayOutput { return v.Publications }).(pulumi.StringArrayOutput)
}

// Name of the replication slot to use. The default behavior is to use the name of the subscription for the slot name
func (o SubscriptionOutput) SlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.SlotName }).(pulumi.StringPtrOutput)
}

type SubscriptionArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subscription)(nil)).Elem()
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutput() SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) ToSubscriptionArrayOutputWithContext(ctx context.Context) SubscriptionArrayOutput {
	return o
}

func (o SubscriptionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Subscription] {
	return pulumix.Output[[]*Subscription]{
		OutputState: o.OutputState,
	}
}

func (o SubscriptionArrayOutput) Index(i pulumi.IntInput) SubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Subscription {
		return vs[0].([]*Subscription)[vs[1].(int)]
	}).(SubscriptionOutput)
}

type SubscriptionMapOutput struct{ *pulumi.OutputState }

func (SubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subscription)(nil)).Elem()
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutput() SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) ToSubscriptionMapOutputWithContext(ctx context.Context) SubscriptionMapOutput {
	return o
}

func (o SubscriptionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Subscription] {
	return pulumix.Output[map[string]*Subscription]{
		OutputState: o.OutputState,
	}
}

func (o SubscriptionMapOutput) MapIndex(k pulumi.StringInput) SubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Subscription {
		return vs[0].(map[string]*Subscription)[vs[1].(string)]
	}).(SubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionInput)(nil)).Elem(), &Subscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionArrayInput)(nil)).Elem(), SubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionMapInput)(nil)).Elem(), SubscriptionMap{})
	pulumi.RegisterOutputType(SubscriptionOutput{})
	pulumi.RegisterOutputType(SubscriptionArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionMapOutput{})
}
