// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logging.NewProjectSink(ctx, "my-sink", &logging.ProjectSinkArgs{
// 			Destination:          pulumi.String("pubsub.googleapis.com/projects/my-project/topics/instance-activity"),
// 			Filter:               pulumi.String("resource.type = gce_instance AND severity >= WARNING"),
// 			UniqueWriterIdentity: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// A more complete example follows: this creates a compute instance, as well as a log sink that logs all activity to a
// cloud storage bucket. Because we are using `uniqueWriterIdentity`, we must grant it access to the bucket. Note that
// this grant requires the "Project IAM Admin" IAM role (`roles/resourcemanager.projectIamAdmin`) granted to the credentials
// used with this provider.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewInstance(ctx, "my-logged-instance", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			Zone:        pulumi.String("us-central1-a"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String("debian-cloud/debian-9"),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						nil,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucket(ctx, "log-bucket", &storage.BucketArgs{
// 			Location: pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewProjectSink(ctx, "instance-sink", &logging.ProjectSinkArgs{
// 			Description: pulumi.String("some explanation on what this is"),
// 			Destination: log_bucket.Name.ApplyT(func(name string) (string, error) {
// 				return fmt.Sprintf("%v%v", "storage.googleapis.com/", name), nil
// 			}).(pulumi.StringOutput),
// 			Filter: my_logged_instance.InstanceId.ApplyT(func(instanceId string) (string, error) {
// 				return fmt.Sprintf("%v%v%v", "resource.type = gce_instance AND resource.labels.instance_id = \"", instanceId, "\""), nil
// 			}).(pulumi.StringOutput),
// 			UniqueWriterIdentity: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMBinding(ctx, "log-writer", &projects.IAMBindingArgs{
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/storage.objectCreator"),
// 			Members: pulumi.StringArray{
// 				instance_sink.WriterIdentity,
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
// The following example uses `exclusions` to filter logs that will not be exported. In this example logs are exported to a [log bucket](https://cloud.google.com/logging/docs/buckets) and there are 2 exclusions configured
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logging.NewProjectSink(ctx, "log-bucket", &logging.ProjectSinkArgs{
// 			Destination: pulumi.String("logging.googleapis.com/projects/my-project/locations/global/buckets/_Default"),
// 			Exclusions: logging.ProjectSinkExclusionArray{
// 				&logging.ProjectSinkExclusionArgs{
// 					Description: pulumi.String("Exclude logs from namespace-1 in k8s"),
// 					Filter:      pulumi.String("resource.type = k8s_container resource.labels.namespace_name=\"namespace-1\" "),
// 					Name:        pulumi.String("nsexcllusion1"),
// 				},
// 				&logging.ProjectSinkExclusionArgs{
// 					Description: pulumi.String("Exclude logs from namespace-2 in k8s"),
// 					Filter:      pulumi.String("resource.type = k8s_container resource.labels.namespace_name=\"namespace-2\" "),
// 					Name:        pulumi.String("nsexcllusion2"),
// 				},
// 			},
// 			UniqueWriterIdentity: pulumi.Bool(true),
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
// Project-level logging sinks can be imported using their URI, e.g.
//
// ```sh
//  $ pulumi import gcp:logging/projectSink:ProjectSink my_sink projects/my-project/sinks/my-sink
// ```
type ProjectSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions ProjectSinkExclusionArrayOutput `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrOutput `pulumi:"uniqueWriterIdentity"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewProjectSink registers a new resource with the given unique name, arguments, and options.
func NewProjectSink(ctx *pulumi.Context,
	name string, args *ProjectSinkArgs, opts ...pulumi.ResourceOption) (*ProjectSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	var resource ProjectSink
	err := ctx.RegisterResource("gcp:logging/projectSink:ProjectSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectSink gets an existing ProjectSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectSinkState, opts ...pulumi.ResourceOption) (*ProjectSink, error) {
	var resource ProjectSink
	err := ctx.ReadResource("gcp:logging/projectSink:ProjectSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectSink resources.
type projectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *ProjectSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination *string `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ProjectSinkExclusion `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type ProjectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsPtrInput
	// A description of this exclusion.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringPtrInput
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions ProjectSinkExclusionArrayInput
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (ProjectSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSinkState)(nil)).Elem()
}

type projectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *ProjectSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this exclusion.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination string `pulumi:"destination"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []ProjectSinkExclusion `pulumi:"exclusions"`
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project *string `pulumi:"project"`
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
}

// The set of arguments for constructing a ProjectSink resource.
type ProjectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions ProjectSinkBigqueryOptionsPtrInput
	// A description of this exclusion.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringInput
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions ProjectSinkExclusionArrayInput
	// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project pulumi.StringPtrInput
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	// `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity pulumi.BoolPtrInput
}

func (ProjectSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectSinkArgs)(nil)).Elem()
}

type ProjectSinkInput interface {
	pulumi.Input

	ToProjectSinkOutput() ProjectSinkOutput
	ToProjectSinkOutputWithContext(ctx context.Context) ProjectSinkOutput
}

func (*ProjectSink) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectSink)(nil)).Elem()
}

func (i *ProjectSink) ToProjectSinkOutput() ProjectSinkOutput {
	return i.ToProjectSinkOutputWithContext(context.Background())
}

func (i *ProjectSink) ToProjectSinkOutputWithContext(ctx context.Context) ProjectSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSinkOutput)
}

// ProjectSinkArrayInput is an input type that accepts ProjectSinkArray and ProjectSinkArrayOutput values.
// You can construct a concrete instance of `ProjectSinkArrayInput` via:
//
//          ProjectSinkArray{ ProjectSinkArgs{...} }
type ProjectSinkArrayInput interface {
	pulumi.Input

	ToProjectSinkArrayOutput() ProjectSinkArrayOutput
	ToProjectSinkArrayOutputWithContext(context.Context) ProjectSinkArrayOutput
}

type ProjectSinkArray []ProjectSinkInput

func (ProjectSinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectSink)(nil)).Elem()
}

func (i ProjectSinkArray) ToProjectSinkArrayOutput() ProjectSinkArrayOutput {
	return i.ToProjectSinkArrayOutputWithContext(context.Background())
}

func (i ProjectSinkArray) ToProjectSinkArrayOutputWithContext(ctx context.Context) ProjectSinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSinkArrayOutput)
}

// ProjectSinkMapInput is an input type that accepts ProjectSinkMap and ProjectSinkMapOutput values.
// You can construct a concrete instance of `ProjectSinkMapInput` via:
//
//          ProjectSinkMap{ "key": ProjectSinkArgs{...} }
type ProjectSinkMapInput interface {
	pulumi.Input

	ToProjectSinkMapOutput() ProjectSinkMapOutput
	ToProjectSinkMapOutputWithContext(context.Context) ProjectSinkMapOutput
}

type ProjectSinkMap map[string]ProjectSinkInput

func (ProjectSinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectSink)(nil)).Elem()
}

func (i ProjectSinkMap) ToProjectSinkMapOutput() ProjectSinkMapOutput {
	return i.ToProjectSinkMapOutputWithContext(context.Background())
}

func (i ProjectSinkMap) ToProjectSinkMapOutputWithContext(ctx context.Context) ProjectSinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectSinkMapOutput)
}

type ProjectSinkOutput struct{ *pulumi.OutputState }

func (ProjectSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectSink)(nil)).Elem()
}

func (o ProjectSinkOutput) ToProjectSinkOutput() ProjectSinkOutput {
	return o
}

func (o ProjectSinkOutput) ToProjectSinkOutputWithContext(ctx context.Context) ProjectSinkOutput {
	return o
}

// Options that affect sinks exporting data to BigQuery. Structure documented below.
func (o ProjectSinkOutput) BigqueryOptions() ProjectSinkBigqueryOptionsOutput {
	return o.ApplyT(func(v *ProjectSink) ProjectSinkBigqueryOptionsOutput { return v.BigqueryOptions }).(ProjectSinkBigqueryOptionsOutput)
}

// A description of this exclusion.
func (o ProjectSinkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination of the sink (or, in other words, where logs are written to). Can be a
// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
// The writer associated with the sink must have access to write to the above resource.
func (o ProjectSinkOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// If set to True, then this exclusion is disabled and it does not exclude any log entries.
func (o ProjectSinkOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
func (o ProjectSinkOutput) Exclusions() ProjectSinkExclusionArrayOutput {
	return o.ApplyT(func(v *ProjectSink) ProjectSinkExclusionArrayOutput { return v.Exclusions }).(ProjectSinkExclusionArrayOutput)
}

// An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
// write a filter.
func (o ProjectSinkOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

// A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
func (o ProjectSinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project to create the sink in. If omitted, the project associated with the provider is
// used.
func (o ProjectSinkOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Whether or not to create a unique identity associated with this sink. If `false`
// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
// then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
// `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
func (o ProjectSinkOutput) UniqueWriterIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.BoolPtrOutput { return v.UniqueWriterIdentity }).(pulumi.BoolPtrOutput)
}

// The identity associated with this sink. This identity must be granted write access to the
// configured `destination`.
func (o ProjectSinkOutput) WriterIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectSink) pulumi.StringOutput { return v.WriterIdentity }).(pulumi.StringOutput)
}

type ProjectSinkArrayOutput struct{ *pulumi.OutputState }

func (ProjectSinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectSink)(nil)).Elem()
}

func (o ProjectSinkArrayOutput) ToProjectSinkArrayOutput() ProjectSinkArrayOutput {
	return o
}

func (o ProjectSinkArrayOutput) ToProjectSinkArrayOutputWithContext(ctx context.Context) ProjectSinkArrayOutput {
	return o
}

func (o ProjectSinkArrayOutput) Index(i pulumi.IntInput) ProjectSinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectSink {
		return vs[0].([]*ProjectSink)[vs[1].(int)]
	}).(ProjectSinkOutput)
}

type ProjectSinkMapOutput struct{ *pulumi.OutputState }

func (ProjectSinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectSink)(nil)).Elem()
}

func (o ProjectSinkMapOutput) ToProjectSinkMapOutput() ProjectSinkMapOutput {
	return o
}

func (o ProjectSinkMapOutput) ToProjectSinkMapOutputWithContext(ctx context.Context) ProjectSinkMapOutput {
	return o
}

func (o ProjectSinkMapOutput) MapIndex(k pulumi.StringInput) ProjectSinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectSink {
		return vs[0].(map[string]*ProjectSink)[vs[1].(string)]
	}).(ProjectSinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectSinkInput)(nil)).Elem(), &ProjectSink{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectSinkArrayInput)(nil)).Elem(), ProjectSinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectSinkMapInput)(nil)).Elem(), ProjectSinkMap{})
	pulumi.RegisterOutputType(ProjectSinkOutput{})
	pulumi.RegisterOutputType(ProjectSinkArrayOutput{})
	pulumi.RegisterOutputType(ProjectSinkMapOutput{})
}
