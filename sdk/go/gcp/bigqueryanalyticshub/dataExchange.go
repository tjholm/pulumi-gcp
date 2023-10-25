// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigqueryanalyticshub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A Bigquery Analytics Hub data exchange
//
// To get more information about DataExchange, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/analytics-hub/rest/v1/projects.locations.dataExchanges)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/bigquery/docs/analytics-hub-introduction)
//
// ## Example Usage
// ### Bigquery Analyticshub Data Exchange Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigqueryanalyticshub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigqueryanalyticshub.NewDataExchange(ctx, "dataExchange", &bigqueryanalyticshub.DataExchangeArgs{
//				DataExchangeId: pulumi.String("my_data_exchange"),
//				Description:    pulumi.String("example data exchange"),
//				DisplayName:    pulumi.String("my_data_exchange"),
//				Location:       pulumi.String("US"),
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
// ## Import
//
// # DataExchange can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchange:DataExchange default projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchange:DataExchange default {{project}}/{{location}}/{{data_exchange_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchange:DataExchange default {{location}}/{{data_exchange_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigqueryanalyticshub/dataExchange:DataExchange default {{data_exchange_id}}
//
// ```
type DataExchange struct {
	pulumi.CustomResourceState

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId pulumi.StringOutput `pulumi:"dataExchangeId"`
	// Description of the data exchange.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	//
	// ***
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Documentation describing the data exchange.
	Documentation pulumi.StringPtrOutput `pulumi:"documentation"`
	// Base64 encoded image representing the data exchange.
	Icon pulumi.StringPtrOutput `pulumi:"icon"`
	// Number of listings contained in the data exchange.
	ListingCount pulumi.IntOutput `pulumi:"listingCount"`
	// The name of the location this data exchange.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the data exchange, for example:
	// "projects/myproject/locations/US/dataExchanges/123"
	Name pulumi.StringOutput `pulumi:"name"`
	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact pulumi.StringPtrOutput `pulumi:"primaryContact"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewDataExchange registers a new resource with the given unique name, arguments, and options.
func NewDataExchange(ctx *pulumi.Context,
	name string, args *DataExchangeArgs, opts ...pulumi.ResourceOption) (*DataExchange, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataExchangeId == nil {
		return nil, errors.New("invalid value for required argument 'DataExchangeId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataExchange
	err := ctx.RegisterResource("gcp:bigqueryanalyticshub/dataExchange:DataExchange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataExchange gets an existing DataExchange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataExchange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExchangeState, opts ...pulumi.ResourceOption) (*DataExchange, error) {
	var resource DataExchange
	err := ctx.ReadResource("gcp:bigqueryanalyticshub/dataExchange:DataExchange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataExchange resources.
type dataExchangeState struct {
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId *string `pulumi:"dataExchangeId"`
	// Description of the data exchange.
	Description *string `pulumi:"description"`
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	//
	// ***
	DisplayName *string `pulumi:"displayName"`
	// Documentation describing the data exchange.
	Documentation *string `pulumi:"documentation"`
	// Base64 encoded image representing the data exchange.
	Icon *string `pulumi:"icon"`
	// Number of listings contained in the data exchange.
	ListingCount *int `pulumi:"listingCount"`
	// The name of the location this data exchange.
	Location *string `pulumi:"location"`
	// The resource name of the data exchange, for example:
	// "projects/myproject/locations/US/dataExchanges/123"
	Name *string `pulumi:"name"`
	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact *string `pulumi:"primaryContact"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type DataExchangeState struct {
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId pulumi.StringPtrInput
	// Description of the data exchange.
	Description pulumi.StringPtrInput
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	//
	// ***
	DisplayName pulumi.StringPtrInput
	// Documentation describing the data exchange.
	Documentation pulumi.StringPtrInput
	// Base64 encoded image representing the data exchange.
	Icon pulumi.StringPtrInput
	// Number of listings contained in the data exchange.
	ListingCount pulumi.IntPtrInput
	// The name of the location this data exchange.
	Location pulumi.StringPtrInput
	// The resource name of the data exchange, for example:
	// "projects/myproject/locations/US/dataExchanges/123"
	Name pulumi.StringPtrInput
	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DataExchangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeState)(nil)).Elem()
}

type dataExchangeArgs struct {
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId string `pulumi:"dataExchangeId"`
	// Description of the data exchange.
	Description *string `pulumi:"description"`
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	//
	// ***
	DisplayName string `pulumi:"displayName"`
	// Documentation describing the data exchange.
	Documentation *string `pulumi:"documentation"`
	// Base64 encoded image representing the data exchange.
	Icon *string `pulumi:"icon"`
	// The name of the location this data exchange.
	Location string `pulumi:"location"`
	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact *string `pulumi:"primaryContact"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a DataExchange resource.
type DataExchangeArgs struct {
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId pulumi.StringInput
	// Description of the data exchange.
	Description pulumi.StringPtrInput
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	//
	// ***
	DisplayName pulumi.StringInput
	// Documentation describing the data exchange.
	Documentation pulumi.StringPtrInput
	// Base64 encoded image representing the data exchange.
	Icon pulumi.StringPtrInput
	// The name of the location this data exchange.
	Location pulumi.StringInput
	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DataExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeArgs)(nil)).Elem()
}

type DataExchangeInput interface {
	pulumi.Input

	ToDataExchangeOutput() DataExchangeOutput
	ToDataExchangeOutputWithContext(ctx context.Context) DataExchangeOutput
}

func (*DataExchange) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchange)(nil)).Elem()
}

func (i *DataExchange) ToDataExchangeOutput() DataExchangeOutput {
	return i.ToDataExchangeOutputWithContext(context.Background())
}

func (i *DataExchange) ToDataExchangeOutputWithContext(ctx context.Context) DataExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeOutput)
}

func (i *DataExchange) ToOutput(ctx context.Context) pulumix.Output[*DataExchange] {
	return pulumix.Output[*DataExchange]{
		OutputState: i.ToDataExchangeOutputWithContext(ctx).OutputState,
	}
}

// DataExchangeArrayInput is an input type that accepts DataExchangeArray and DataExchangeArrayOutput values.
// You can construct a concrete instance of `DataExchangeArrayInput` via:
//
//	DataExchangeArray{ DataExchangeArgs{...} }
type DataExchangeArrayInput interface {
	pulumi.Input

	ToDataExchangeArrayOutput() DataExchangeArrayOutput
	ToDataExchangeArrayOutputWithContext(context.Context) DataExchangeArrayOutput
}

type DataExchangeArray []DataExchangeInput

func (DataExchangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataExchange)(nil)).Elem()
}

func (i DataExchangeArray) ToDataExchangeArrayOutput() DataExchangeArrayOutput {
	return i.ToDataExchangeArrayOutputWithContext(context.Background())
}

func (i DataExchangeArray) ToDataExchangeArrayOutputWithContext(ctx context.Context) DataExchangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeArrayOutput)
}

func (i DataExchangeArray) ToOutput(ctx context.Context) pulumix.Output[[]*DataExchange] {
	return pulumix.Output[[]*DataExchange]{
		OutputState: i.ToDataExchangeArrayOutputWithContext(ctx).OutputState,
	}
}

// DataExchangeMapInput is an input type that accepts DataExchangeMap and DataExchangeMapOutput values.
// You can construct a concrete instance of `DataExchangeMapInput` via:
//
//	DataExchangeMap{ "key": DataExchangeArgs{...} }
type DataExchangeMapInput interface {
	pulumi.Input

	ToDataExchangeMapOutput() DataExchangeMapOutput
	ToDataExchangeMapOutputWithContext(context.Context) DataExchangeMapOutput
}

type DataExchangeMap map[string]DataExchangeInput

func (DataExchangeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataExchange)(nil)).Elem()
}

func (i DataExchangeMap) ToDataExchangeMapOutput() DataExchangeMapOutput {
	return i.ToDataExchangeMapOutputWithContext(context.Background())
}

func (i DataExchangeMap) ToDataExchangeMapOutputWithContext(ctx context.Context) DataExchangeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeMapOutput)
}

func (i DataExchangeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DataExchange] {
	return pulumix.Output[map[string]*DataExchange]{
		OutputState: i.ToDataExchangeMapOutputWithContext(ctx).OutputState,
	}
}

type DataExchangeOutput struct{ *pulumi.OutputState }

func (DataExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchange)(nil)).Elem()
}

func (o DataExchangeOutput) ToDataExchangeOutput() DataExchangeOutput {
	return o
}

func (o DataExchangeOutput) ToDataExchangeOutputWithContext(ctx context.Context) DataExchangeOutput {
	return o
}

func (o DataExchangeOutput) ToOutput(ctx context.Context) pulumix.Output[*DataExchange] {
	return pulumix.Output[*DataExchange]{
		OutputState: o.OutputState,
	}
}

// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
func (o DataExchangeOutput) DataExchangeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.DataExchangeId }).(pulumi.StringOutput)
}

// Description of the data exchange.
func (o DataExchangeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
//
// ***
func (o DataExchangeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Documentation describing the data exchange.
func (o DataExchangeOutput) Documentation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringPtrOutput { return v.Documentation }).(pulumi.StringPtrOutput)
}

// Base64 encoded image representing the data exchange.
func (o DataExchangeOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringPtrOutput { return v.Icon }).(pulumi.StringPtrOutput)
}

// Number of listings contained in the data exchange.
func (o DataExchangeOutput) ListingCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.IntOutput { return v.ListingCount }).(pulumi.IntOutput)
}

// The name of the location this data exchange.
func (o DataExchangeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the data exchange, for example:
// "projects/myproject/locations/US/dataExchanges/123"
func (o DataExchangeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Email or URL of the primary point of contact of the data exchange.
func (o DataExchangeOutput) PrimaryContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringPtrOutput { return v.PrimaryContact }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DataExchangeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type DataExchangeArrayOutput struct{ *pulumi.OutputState }

func (DataExchangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataExchange)(nil)).Elem()
}

func (o DataExchangeArrayOutput) ToDataExchangeArrayOutput() DataExchangeArrayOutput {
	return o
}

func (o DataExchangeArrayOutput) ToDataExchangeArrayOutputWithContext(ctx context.Context) DataExchangeArrayOutput {
	return o
}

func (o DataExchangeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DataExchange] {
	return pulumix.Output[[]*DataExchange]{
		OutputState: o.OutputState,
	}
}

func (o DataExchangeArrayOutput) Index(i pulumi.IntInput) DataExchangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataExchange {
		return vs[0].([]*DataExchange)[vs[1].(int)]
	}).(DataExchangeOutput)
}

type DataExchangeMapOutput struct{ *pulumi.OutputState }

func (DataExchangeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataExchange)(nil)).Elem()
}

func (o DataExchangeMapOutput) ToDataExchangeMapOutput() DataExchangeMapOutput {
	return o
}

func (o DataExchangeMapOutput) ToDataExchangeMapOutputWithContext(ctx context.Context) DataExchangeMapOutput {
	return o
}

func (o DataExchangeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DataExchange] {
	return pulumix.Output[map[string]*DataExchange]{
		OutputState: o.OutputState,
	}
}

func (o DataExchangeMapOutput) MapIndex(k pulumi.StringInput) DataExchangeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataExchange {
		return vs[0].(map[string]*DataExchange)[vs[1].(string)]
	}).(DataExchangeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeInput)(nil)).Elem(), &DataExchange{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeArrayInput)(nil)).Elem(), DataExchangeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeMapInput)(nil)).Elem(), DataExchangeMap{})
	pulumi.RegisterOutputType(DataExchangeOutput{})
	pulumi.RegisterOutputType(DataExchangeArrayOutput{})
	pulumi.RegisterOutputType(DataExchangeMapOutput{})
}
