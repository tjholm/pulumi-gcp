// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Billing account logging exclusions can be imported using their URI, e.g.
//
// ```sh
//  $ pulumi import gcp:logging/billingAccountExclusion:BillingAccountExclusion my_exclusion billingAccounts/my-billing_account/exclusions/my-exclusion
// ```
type BillingAccountExclusion struct {
	pulumi.CustomResourceState

	// The billing account to create the exclusion for.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// A human-readable description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// The name of the logging exclusion.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewBillingAccountExclusion registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountExclusion(ctx *pulumi.Context,
	name string, args *BillingAccountExclusionArgs, opts ...pulumi.ResourceOption) (*BillingAccountExclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccount'")
	}
	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	var resource BillingAccountExclusion
	err := ctx.RegisterResource("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountExclusion gets an existing BillingAccountExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountExclusionState, opts ...pulumi.ResourceOption) (*BillingAccountExclusion, error) {
	var resource BillingAccountExclusion
	err := ctx.ReadResource("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountExclusion resources.
type billingAccountExclusionState struct {
	// The billing account to create the exclusion for.
	BillingAccount *string `pulumi:"billingAccount"`
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
}

type BillingAccountExclusionState struct {
	// The billing account to create the exclusion for.
	BillingAccount pulumi.StringPtrInput
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
}

func (BillingAccountExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountExclusionState)(nil)).Elem()
}

type billingAccountExclusionArgs struct {
	// The billing account to create the exclusion for.
	BillingAccount string `pulumi:"billingAccount"`
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter string `pulumi:"filter"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a BillingAccountExclusion resource.
type BillingAccountExclusionArgs struct {
	// The billing account to create the exclusion for.
	BillingAccount pulumi.StringInput
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to
	// false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
	// write a filter.
	Filter pulumi.StringInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
}

func (BillingAccountExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountExclusionArgs)(nil)).Elem()
}

type BillingAccountExclusionInput interface {
	pulumi.Input

	ToBillingAccountExclusionOutput() BillingAccountExclusionOutput
	ToBillingAccountExclusionOutputWithContext(ctx context.Context) BillingAccountExclusionOutput
}

func (*BillingAccountExclusion) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountExclusion)(nil)).Elem()
}

func (i *BillingAccountExclusion) ToBillingAccountExclusionOutput() BillingAccountExclusionOutput {
	return i.ToBillingAccountExclusionOutputWithContext(context.Background())
}

func (i *BillingAccountExclusion) ToBillingAccountExclusionOutputWithContext(ctx context.Context) BillingAccountExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountExclusionOutput)
}

// BillingAccountExclusionArrayInput is an input type that accepts BillingAccountExclusionArray and BillingAccountExclusionArrayOutput values.
// You can construct a concrete instance of `BillingAccountExclusionArrayInput` via:
//
//          BillingAccountExclusionArray{ BillingAccountExclusionArgs{...} }
type BillingAccountExclusionArrayInput interface {
	pulumi.Input

	ToBillingAccountExclusionArrayOutput() BillingAccountExclusionArrayOutput
	ToBillingAccountExclusionArrayOutputWithContext(context.Context) BillingAccountExclusionArrayOutput
}

type BillingAccountExclusionArray []BillingAccountExclusionInput

func (BillingAccountExclusionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BillingAccountExclusion)(nil)).Elem()
}

func (i BillingAccountExclusionArray) ToBillingAccountExclusionArrayOutput() BillingAccountExclusionArrayOutput {
	return i.ToBillingAccountExclusionArrayOutputWithContext(context.Background())
}

func (i BillingAccountExclusionArray) ToBillingAccountExclusionArrayOutputWithContext(ctx context.Context) BillingAccountExclusionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountExclusionArrayOutput)
}

// BillingAccountExclusionMapInput is an input type that accepts BillingAccountExclusionMap and BillingAccountExclusionMapOutput values.
// You can construct a concrete instance of `BillingAccountExclusionMapInput` via:
//
//          BillingAccountExclusionMap{ "key": BillingAccountExclusionArgs{...} }
type BillingAccountExclusionMapInput interface {
	pulumi.Input

	ToBillingAccountExclusionMapOutput() BillingAccountExclusionMapOutput
	ToBillingAccountExclusionMapOutputWithContext(context.Context) BillingAccountExclusionMapOutput
}

type BillingAccountExclusionMap map[string]BillingAccountExclusionInput

func (BillingAccountExclusionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BillingAccountExclusion)(nil)).Elem()
}

func (i BillingAccountExclusionMap) ToBillingAccountExclusionMapOutput() BillingAccountExclusionMapOutput {
	return i.ToBillingAccountExclusionMapOutputWithContext(context.Background())
}

func (i BillingAccountExclusionMap) ToBillingAccountExclusionMapOutputWithContext(ctx context.Context) BillingAccountExclusionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountExclusionMapOutput)
}

type BillingAccountExclusionOutput struct{ *pulumi.OutputState }

func (BillingAccountExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountExclusion)(nil)).Elem()
}

func (o BillingAccountExclusionOutput) ToBillingAccountExclusionOutput() BillingAccountExclusionOutput {
	return o
}

func (o BillingAccountExclusionOutput) ToBillingAccountExclusionOutputWithContext(ctx context.Context) BillingAccountExclusionOutput {
	return o
}

type BillingAccountExclusionArrayOutput struct{ *pulumi.OutputState }

func (BillingAccountExclusionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BillingAccountExclusion)(nil)).Elem()
}

func (o BillingAccountExclusionArrayOutput) ToBillingAccountExclusionArrayOutput() BillingAccountExclusionArrayOutput {
	return o
}

func (o BillingAccountExclusionArrayOutput) ToBillingAccountExclusionArrayOutputWithContext(ctx context.Context) BillingAccountExclusionArrayOutput {
	return o
}

func (o BillingAccountExclusionArrayOutput) Index(i pulumi.IntInput) BillingAccountExclusionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BillingAccountExclusion {
		return vs[0].([]*BillingAccountExclusion)[vs[1].(int)]
	}).(BillingAccountExclusionOutput)
}

type BillingAccountExclusionMapOutput struct{ *pulumi.OutputState }

func (BillingAccountExclusionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BillingAccountExclusion)(nil)).Elem()
}

func (o BillingAccountExclusionMapOutput) ToBillingAccountExclusionMapOutput() BillingAccountExclusionMapOutput {
	return o
}

func (o BillingAccountExclusionMapOutput) ToBillingAccountExclusionMapOutputWithContext(ctx context.Context) BillingAccountExclusionMapOutput {
	return o
}

func (o BillingAccountExclusionMapOutput) MapIndex(k pulumi.StringInput) BillingAccountExclusionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BillingAccountExclusion {
		return vs[0].(map[string]*BillingAccountExclusion)[vs[1].(string)]
	}).(BillingAccountExclusionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountExclusionInput)(nil)).Elem(), &BillingAccountExclusion{})
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountExclusionArrayInput)(nil)).Elem(), BillingAccountExclusionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountExclusionMapInput)(nil)).Elem(), BillingAccountExclusionMap{})
	pulumi.RegisterOutputType(BillingAccountExclusionOutput{})
	pulumi.RegisterOutputType(BillingAccountExclusionArrayOutput{})
	pulumi.RegisterOutputType(BillingAccountExclusionMapOutput{})
}
