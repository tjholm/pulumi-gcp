// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logs-based metric can also be used to extract values from logs and create a a distribution
// of the values. The distribution records the statistics of the extracted values along with
// an optional histogram of the values as specified by the bucket options.
//
// To get more information about Metric, see:
//
// * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics/create)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/logging/docs/apis)
//
// ## Example Usage
// ### Logging Metric Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logging.NewMetric(ctx, "loggingMetric", &logging.MetricArgs{
// 			BucketOptions: &logging.MetricBucketOptionsArgs{
// 				LinearBuckets: &logging.MetricBucketOptionsLinearBucketsArgs{
// 					NumFiniteBuckets: pulumi.Int(3),
// 					Offset:           pulumi.Float64(1),
// 					Width:            pulumi.Float64(1),
// 				},
// 			},
// 			Filter: pulumi.String("resource.type=gae_app AND severity>=ERROR"),
// 			LabelExtractors: pulumi.StringMap{
// 				"mass": pulumi.String("EXTRACT(jsonPayload.request)"),
// 				"sku":  pulumi.String("EXTRACT(jsonPayload.id)"),
// 			},
// 			MetricDescriptor: &logging.MetricMetricDescriptorArgs{
// 				DisplayName: pulumi.String("My metric"),
// 				Labels: logging.MetricMetricDescriptorLabelArray{
// 					&logging.MetricMetricDescriptorLabelArgs{
// 						Description: pulumi.String("amount of matter"),
// 						Key:         pulumi.String("mass"),
// 						ValueType:   pulumi.String("STRING"),
// 					},
// 					&logging.MetricMetricDescriptorLabelArgs{
// 						Description: pulumi.String("Identifying number for item"),
// 						Key:         pulumi.String("sku"),
// 						ValueType:   pulumi.String("INT64"),
// 					},
// 				},
// 				MetricKind: pulumi.String("DELTA"),
// 				Unit:       pulumi.String("1"),
// 				ValueType:  pulumi.String("DISTRIBUTION"),
// 			},
// 			ValueExtractor: pulumi.String("EXTRACT(jsonPayload.request)"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Logging Metric Counter Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logging.NewMetric(ctx, "loggingMetric", &logging.MetricArgs{
// 			Filter: pulumi.String("resource.type=gae_app AND severity>=ERROR"),
// 			MetricDescriptor: &logging.MetricMetricDescriptorArgs{
// 				MetricKind: pulumi.String("DELTA"),
// 				ValueType:  pulumi.String("INT64"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Logging Metric Counter Labels
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logging.NewMetric(ctx, "loggingMetric", &logging.MetricArgs{
// 			Filter: pulumi.String("resource.type=gae_app AND severity>=ERROR"),
// 			LabelExtractors: pulumi.StringMap{
// 				"mass": pulumi.String("EXTRACT(jsonPayload.request)"),
// 			},
// 			MetricDescriptor: &logging.MetricMetricDescriptorArgs{
// 				Labels: logging.MetricMetricDescriptorLabelArray{
// 					&logging.MetricMetricDescriptorLabelArgs{
// 						Description: pulumi.String("amount of matter"),
// 						Key:         pulumi.String("mass"),
// 						ValueType:   pulumi.String("STRING"),
// 					},
// 				},
// 				MetricKind: pulumi.String("DELTA"),
// 				ValueType:  pulumi.String("INT64"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Metric can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:logging/metric:Metric default {{project}} {{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:logging/metric:Metric default {{name}}
// ```
type Metric struct {
	pulumi.CustomResourceState

	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions MetricBucketOptionsPtrOutput `pulumi:"bucketOptions"`
	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	LabelExtractors pulumi.StringMapOutput `pulumi:"labelExtractors"`
	// The metric descriptor associated with the logs-based metric.
	// Structure is documented below.
	MetricDescriptor MetricMetricDescriptorOutput `pulumi:"metricDescriptor"`
	// The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
	// Metric identifiers are limited to 100 characters and can include only the following
	// characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
	// character (/) denotes a hierarchy of name pieces, and it cannot be the first character
	// of the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor pulumi.StringPtrOutput `pulumi:"valueExtractor"`
}

// NewMetric registers a new resource with the given unique name, arguments, and options.
func NewMetric(ctx *pulumi.Context,
	name string, args *MetricArgs, opts ...pulumi.ResourceOption) (*Metric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.MetricDescriptor == nil {
		return nil, errors.New("invalid value for required argument 'MetricDescriptor'")
	}
	var resource Metric
	err := ctx.RegisterResource("gcp:logging/metric:Metric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetric gets an existing Metric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricState, opts ...pulumi.ResourceOption) (*Metric, error) {
	var resource Metric
	err := ctx.ReadResource("gcp:logging/metric:Metric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Metric resources.
type metricState struct {
	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions *MetricBucketOptions `pulumi:"bucketOptions"`
	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description *string `pulumi:"description"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter *string `pulumi:"filter"`
	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	LabelExtractors map[string]string `pulumi:"labelExtractors"`
	// The metric descriptor associated with the logs-based metric.
	// Structure is documented below.
	MetricDescriptor *MetricMetricDescriptor `pulumi:"metricDescriptor"`
	// The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
	// Metric identifiers are limited to 100 characters and can include only the following
	// characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
	// character (/) denotes a hierarchy of name pieces, and it cannot be the first character
	// of the name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor *string `pulumi:"valueExtractor"`
}

type MetricState struct {
	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions MetricBucketOptionsPtrInput
	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description pulumi.StringPtrInput
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter pulumi.StringPtrInput
	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	LabelExtractors pulumi.StringMapInput
	// The metric descriptor associated with the logs-based metric.
	// Structure is documented below.
	MetricDescriptor MetricMetricDescriptorPtrInput
	// The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
	// Metric identifiers are limited to 100 characters and can include only the following
	// characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
	// character (/) denotes a hierarchy of name pieces, and it cannot be the first character
	// of the name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor pulumi.StringPtrInput
}

func (MetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricState)(nil)).Elem()
}

type metricArgs struct {
	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions *MetricBucketOptions `pulumi:"bucketOptions"`
	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description *string `pulumi:"description"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter string `pulumi:"filter"`
	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	LabelExtractors map[string]string `pulumi:"labelExtractors"`
	// The metric descriptor associated with the logs-based metric.
	// Structure is documented below.
	MetricDescriptor MetricMetricDescriptor `pulumi:"metricDescriptor"`
	// The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
	// Metric identifiers are limited to 100 characters and can include only the following
	// characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
	// character (/) denotes a hierarchy of name pieces, and it cannot be the first character
	// of the name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor *string `pulumi:"valueExtractor"`
}

// The set of arguments for constructing a Metric resource.
type MetricArgs struct {
	// The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
	// describes the bucket boundaries used to create a histogram of the extracted values.
	// Structure is documented below.
	BucketOptions MetricBucketOptionsPtrInput
	// A description of this metric, which is used in documentation. The maximum length of the
	// description is 8000 characters.
	Description pulumi.StringPtrInput
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
	// is used to match log entries.
	Filter pulumi.StringInput
	// A map from a label key string to an extractor expression which is used to extract data from a log
	// entry field and assign as the label value. Each label key specified in the LabelDescriptor must
	// have an associated extractor expression in this map. The syntax of the extractor expression is
	// the same as for the valueExtractor field.
	LabelExtractors pulumi.StringMapInput
	// The metric descriptor associated with the logs-based metric.
	// Structure is documented below.
	MetricDescriptor MetricMetricDescriptorInput
	// The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
	// Metric identifiers are limited to 100 characters and can include only the following
	// characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
	// character (/) denotes a hierarchy of name pieces, and it cannot be the first character
	// of the name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A valueExtractor is required when using a distribution logs-based metric to extract the values to
	// record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
	// REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
	// the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
	// (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
	// log entry field. The value of the field is converted to a string before applying the regex. It is an
	// error to specify a regex that does not include exactly one capture group.
	ValueExtractor pulumi.StringPtrInput
}

func (MetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricArgs)(nil)).Elem()
}

type MetricInput interface {
	pulumi.Input

	ToMetricOutput() MetricOutput
	ToMetricOutputWithContext(ctx context.Context) MetricOutput
}

func (*Metric) ElementType() reflect.Type {
	return reflect.TypeOf((*Metric)(nil))
}

func (i *Metric) ToMetricOutput() MetricOutput {
	return i.ToMetricOutputWithContext(context.Background())
}

func (i *Metric) ToMetricOutputWithContext(ctx context.Context) MetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricOutput)
}

func (i *Metric) ToMetricPtrOutput() MetricPtrOutput {
	return i.ToMetricPtrOutputWithContext(context.Background())
}

func (i *Metric) ToMetricPtrOutputWithContext(ctx context.Context) MetricPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricPtrOutput)
}

type MetricPtrInput interface {
	pulumi.Input

	ToMetricPtrOutput() MetricPtrOutput
	ToMetricPtrOutputWithContext(ctx context.Context) MetricPtrOutput
}

type metricPtrType MetricArgs

func (*metricPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Metric)(nil))
}

func (i *metricPtrType) ToMetricPtrOutput() MetricPtrOutput {
	return i.ToMetricPtrOutputWithContext(context.Background())
}

func (i *metricPtrType) ToMetricPtrOutputWithContext(ctx context.Context) MetricPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricPtrOutput)
}

// MetricArrayInput is an input type that accepts MetricArray and MetricArrayOutput values.
// You can construct a concrete instance of `MetricArrayInput` via:
//
//          MetricArray{ MetricArgs{...} }
type MetricArrayInput interface {
	pulumi.Input

	ToMetricArrayOutput() MetricArrayOutput
	ToMetricArrayOutputWithContext(context.Context) MetricArrayOutput
}

type MetricArray []MetricInput

func (MetricArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Metric)(nil))
}

func (i MetricArray) ToMetricArrayOutput() MetricArrayOutput {
	return i.ToMetricArrayOutputWithContext(context.Background())
}

func (i MetricArray) ToMetricArrayOutputWithContext(ctx context.Context) MetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricArrayOutput)
}

// MetricMapInput is an input type that accepts MetricMap and MetricMapOutput values.
// You can construct a concrete instance of `MetricMapInput` via:
//
//          MetricMap{ "key": MetricArgs{...} }
type MetricMapInput interface {
	pulumi.Input

	ToMetricMapOutput() MetricMapOutput
	ToMetricMapOutputWithContext(context.Context) MetricMapOutput
}

type MetricMap map[string]MetricInput

func (MetricMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Metric)(nil))
}

func (i MetricMap) ToMetricMapOutput() MetricMapOutput {
	return i.ToMetricMapOutputWithContext(context.Background())
}

func (i MetricMap) ToMetricMapOutputWithContext(ctx context.Context) MetricMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricMapOutput)
}

type MetricOutput struct {
	*pulumi.OutputState
}

func (MetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Metric)(nil))
}

func (o MetricOutput) ToMetricOutput() MetricOutput {
	return o
}

func (o MetricOutput) ToMetricOutputWithContext(ctx context.Context) MetricOutput {
	return o
}

func (o MetricOutput) ToMetricPtrOutput() MetricPtrOutput {
	return o.ToMetricPtrOutputWithContext(context.Background())
}

func (o MetricOutput) ToMetricPtrOutputWithContext(ctx context.Context) MetricPtrOutput {
	return o.ApplyT(func(v Metric) *Metric {
		return &v
	}).(MetricPtrOutput)
}

type MetricPtrOutput struct {
	*pulumi.OutputState
}

func (MetricPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Metric)(nil))
}

func (o MetricPtrOutput) ToMetricPtrOutput() MetricPtrOutput {
	return o
}

func (o MetricPtrOutput) ToMetricPtrOutputWithContext(ctx context.Context) MetricPtrOutput {
	return o
}

type MetricArrayOutput struct{ *pulumi.OutputState }

func (MetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Metric)(nil))
}

func (o MetricArrayOutput) ToMetricArrayOutput() MetricArrayOutput {
	return o
}

func (o MetricArrayOutput) ToMetricArrayOutputWithContext(ctx context.Context) MetricArrayOutput {
	return o
}

func (o MetricArrayOutput) Index(i pulumi.IntInput) MetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Metric {
		return vs[0].([]Metric)[vs[1].(int)]
	}).(MetricOutput)
}

type MetricMapOutput struct{ *pulumi.OutputState }

func (MetricMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Metric)(nil))
}

func (o MetricMapOutput) ToMetricMapOutput() MetricMapOutput {
	return o
}

func (o MetricMapOutput) ToMetricMapOutputWithContext(ctx context.Context) MetricMapOutput {
	return o
}

func (o MetricMapOutput) MapIndex(k pulumi.StringInput) MetricOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Metric {
		return vs[0].(map[string]Metric)[vs[1].(string)]
	}).(MetricOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricOutput{})
	pulumi.RegisterOutputType(MetricPtrOutput{})
	pulumi.RegisterOutputType(MetricArrayOutput{})
	pulumi.RegisterOutputType(MetricMapOutput{})
}
