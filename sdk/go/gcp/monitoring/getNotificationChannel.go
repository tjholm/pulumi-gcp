// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A NotificationChannel is a medium through which an alert is delivered
// when a policy violation is detected. Examples of channels include email, SMS,
// and third-party messaging applications. Fields containing sensitive information
// like authentication tokens or contact info are only partially populated on retrieval.
//
// To get more information about NotificationChannel, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
// * How-to Guides
//   - [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
//   - [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Notification Channel Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/monitoring"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			basic, err := monitoring.LookupNotificationChannel(ctx, &monitoring.LookupNotificationChannelArgs{
//				DisplayName: pulumi.StringRef("Test Notification Channel"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = monitoring.NewAlertPolicy(ctx, "alertPolicy", &monitoring.AlertPolicyArgs{
//				DisplayName: pulumi.String("My Alert Policy"),
//				NotificationChannels: pulumi.StringArray{
//					*pulumi.String(basic.Name),
//				},
//				Combiner: pulumi.String("OR"),
//				Conditions: monitoring.AlertPolicyConditionArray{
//					&monitoring.AlertPolicyConditionArgs{
//						DisplayName: pulumi.String("test condition"),
//						ConditionThreshold: &monitoring.AlertPolicyConditionConditionThresholdArgs{
//							Filter:     pulumi.String("metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\""),
//							Duration:   pulumi.String("60s"),
//							Comparison: pulumi.String("COMPARISON_GT"),
//							Aggregations: monitoring.AlertPolicyConditionConditionThresholdAggregationArray{
//								&monitoring.AlertPolicyConditionConditionThresholdAggregationArgs{
//									AlignmentPeriod:  pulumi.String("60s"),
//									PerSeriesAligner: pulumi.String("ALIGN_RATE"),
//								},
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNotificationChannel(ctx *pulumi.Context, args *LookupNotificationChannelArgs, opts ...pulumi.InvokeOption) (*LookupNotificationChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationChannelResult
	err := ctx.Invoke("gcp:monitoring/getNotificationChannel:getNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotificationChannel.
type LookupNotificationChannelArgs struct {
	// The display name for this notification channel.
	DisplayName *string `pulumi:"displayName"`
	// Labels (corresponding to the
	// NotificationChannelDescriptor schema) to filter the notification channels by.
	Labels map[string]string `pulumi:"labels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The type of the notification channel.
	//
	// ***
	//
	// Other optional fields include:
	Type *string `pulumi:"type"`
	// User-provided key-value labels to filter by.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// A collection of values returned by getNotificationChannel.
type LookupNotificationChannelResult struct {
	// An optional human-readable description of this notification channel.
	Description string  `pulumi:"description"`
	DisplayName *string `pulumi:"displayName"`
	// Whether notifications are forwarded to the described channel.
	Enabled     bool `pulumi:"enabled"`
	ForceDelete bool `pulumi:"forceDelete"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Configuration fields that define the channel and its behavior.
	Labels map[string]string `pulumi:"labels"`
	// The full REST resource name for this channel. The syntax is:
	// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`.
	Name            string                                 `pulumi:"name"`
	Project         *string                                `pulumi:"project"`
	SensitiveLabels []GetNotificationChannelSensitiveLabel `pulumi:"sensitiveLabels"`
	Type            *string                                `pulumi:"type"`
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field.
	UserLabels map[string]string `pulumi:"userLabels"`
	// Indicates whether this channel has been verified or not.
	VerificationStatus string `pulumi:"verificationStatus"`
}

func LookupNotificationChannelOutput(ctx *pulumi.Context, args LookupNotificationChannelOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationChannelResult, error) {
			args := v.(LookupNotificationChannelArgs)
			r, err := LookupNotificationChannel(ctx, &args, opts...)
			var s LookupNotificationChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationChannelResultOutput)
}

// A collection of arguments for invoking getNotificationChannel.
type LookupNotificationChannelOutputArgs struct {
	// The display name for this notification channel.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// Labels (corresponding to the
	// NotificationChannelDescriptor schema) to filter the notification channels by.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The type of the notification channel.
	//
	// ***
	//
	// Other optional fields include:
	Type pulumi.StringPtrInput `pulumi:"type"`
	// User-provided key-value labels to filter by.
	UserLabels pulumi.StringMapInput `pulumi:"userLabels"`
}

func (LookupNotificationChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelArgs)(nil)).Elem()
}

// A collection of values returned by getNotificationChannel.
type LookupNotificationChannelResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationChannelResult)(nil)).Elem()
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutput() LookupNotificationChannelResultOutput {
	return o
}

func (o LookupNotificationChannelResultOutput) ToLookupNotificationChannelResultOutputWithContext(ctx context.Context) LookupNotificationChannelResultOutput {
	return o
}

func (o LookupNotificationChannelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNotificationChannelResult] {
	return pulumix.Output[LookupNotificationChannelResult]{
		OutputState: o.OutputState,
	}
}

// An optional human-readable description of this notification channel.
func (o LookupNotificationChannelResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Whether notifications are forwarded to the described channel.
func (o LookupNotificationChannelResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupNotificationChannelResultOutput) ForceDelete() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) bool { return v.ForceDelete }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNotificationChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Configuration fields that define the channel and its behavior.
func (o LookupNotificationChannelResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The full REST resource name for this channel. The syntax is:
// `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]`.
func (o LookupNotificationChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotificationChannelResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationChannelResultOutput) SensitiveLabels() GetNotificationChannelSensitiveLabelArrayOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) []GetNotificationChannelSensitiveLabel {
		return v.SensitiveLabels
	}).(GetNotificationChannelSensitiveLabelArrayOutput)
}

func (o LookupNotificationChannelResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field.
func (o LookupNotificationChannelResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

// Indicates whether this channel has been verified or not.
func (o LookupNotificationChannelResultOutput) VerificationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationChannelResult) string { return v.VerificationStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationChannelResultOutput{})
}
