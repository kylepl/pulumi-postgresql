// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type Clientcert struct {
	Cert      string `pulumi:"cert"`
	Key       string `pulumi:"key"`
	Sslinline *bool  `pulumi:"sslinline"`
}

// ClientcertInput is an input type that accepts ClientcertArgs and ClientcertOutput values.
// You can construct a concrete instance of `ClientcertInput` via:
//
//	ClientcertArgs{...}
type ClientcertInput interface {
	pulumi.Input

	ToClientcertOutput() ClientcertOutput
	ToClientcertOutputWithContext(context.Context) ClientcertOutput
}

type ClientcertArgs struct {
	Cert      pulumi.StringInput  `pulumi:"cert"`
	Key       pulumi.StringInput  `pulumi:"key"`
	Sslinline pulumi.BoolPtrInput `pulumi:"sslinline"`
}

func (ClientcertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Clientcert)(nil)).Elem()
}

func (i ClientcertArgs) ToClientcertOutput() ClientcertOutput {
	return i.ToClientcertOutputWithContext(context.Background())
}

func (i ClientcertArgs) ToClientcertOutputWithContext(ctx context.Context) ClientcertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientcertOutput)
}

func (i ClientcertArgs) ToOutput(ctx context.Context) pulumix.Output[Clientcert] {
	return pulumix.Output[Clientcert]{
		OutputState: i.ToClientcertOutputWithContext(ctx).OutputState,
	}
}

type ClientcertOutput struct{ *pulumi.OutputState }

func (ClientcertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Clientcert)(nil)).Elem()
}

func (o ClientcertOutput) ToClientcertOutput() ClientcertOutput {
	return o
}

func (o ClientcertOutput) ToClientcertOutputWithContext(ctx context.Context) ClientcertOutput {
	return o
}

func (o ClientcertOutput) ToOutput(ctx context.Context) pulumix.Output[Clientcert] {
	return pulumix.Output[Clientcert]{
		OutputState: o.OutputState,
	}
}

func (o ClientcertOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v Clientcert) string { return v.Cert }).(pulumi.StringOutput)
}

func (o ClientcertOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v Clientcert) string { return v.Key }).(pulumi.StringOutput)
}

func (o ClientcertOutput) Sslinline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Clientcert) *bool { return v.Sslinline }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientcertInput)(nil)).Elem(), ClientcertArgs{})
	pulumi.RegisterOutputType(ClientcertOutput{})
}
