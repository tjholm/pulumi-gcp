// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clouddeploy

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Cloud Deploy `DeliveryPipeline` resource
//
// ## Example Usage
// ### Canary_delivery_pipeline
// Creates a basic Cloud Deploy delivery pipeline
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/clouddeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clouddeploy.NewDeliveryPipeline(ctx, "primary", &clouddeploy.DeliveryPipelineArgs{
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("basic description"),
//				Project:     pulumi.String("my-project-name"),
//				SerialPipeline: &clouddeploy.DeliveryPipelineSerialPipelineArgs{
//					Stages: clouddeploy.DeliveryPipelineSerialPipelineStageArray{
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							DeployParameters: clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArray{
//								&clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArgs{
//									Values: pulumi.StringMap{
//										"deployParameterKey": pulumi.String("deployParameterValue"),
//									},
//									MatchTargetLabels: nil,
//								},
//							},
//							Profiles: pulumi.StringArray{
//								pulumi.String("example-profile-one"),
//								pulumi.String("example-profile-two"),
//							},
//							TargetId: pulumi.String("example-target-one"),
//						},
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							Profiles: pulumi.StringArray{},
//							TargetId: pulumi.String("example-target-two"),
//						},
//					},
//				},
//				Annotations: pulumi.StringMap{
//					"my_first_annotation":  pulumi.String("example-annotation-1"),
//					"my_second_annotation": pulumi.String("example-annotation-2"),
//				},
//				Labels: pulumi.StringMap{
//					"my_first_label":  pulumi.String("example-label-1"),
//					"my_second_label": pulumi.String("example-label-2"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Canary_service_networking_delivery_pipeline
// Creates a basic Cloud Deploy delivery pipeline
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/clouddeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clouddeploy.NewDeliveryPipeline(ctx, "primary", &clouddeploy.DeliveryPipelineArgs{
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("basic description"),
//				Project:     pulumi.String("my-project-name"),
//				SerialPipeline: &clouddeploy.DeliveryPipelineSerialPipelineArgs{
//					Stages: clouddeploy.DeliveryPipelineSerialPipelineStageArray{
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							DeployParameters: clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArray{
//								&clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArgs{
//									Values: pulumi.StringMap{
//										"deployParameterKey": pulumi.String("deployParameterValue"),
//									},
//									MatchTargetLabels: nil,
//								},
//							},
//							Profiles: pulumi.StringArray{
//								pulumi.String("example-profile-one"),
//								pulumi.String("example-profile-two"),
//							},
//							TargetId: pulumi.String("example-target-one"),
//						},
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							Profiles: pulumi.StringArray{},
//							TargetId: pulumi.String("example-target-two"),
//						},
//					},
//				},
//				Annotations: pulumi.StringMap{
//					"my_first_annotation":  pulumi.String("example-annotation-1"),
//					"my_second_annotation": pulumi.String("example-annotation-2"),
//				},
//				Labels: pulumi.StringMap{
//					"my_first_label":  pulumi.String("example-label-1"),
//					"my_second_label": pulumi.String("example-label-2"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Canaryrun_delivery_pipeline
// Creates a basic Cloud Deploy delivery pipeline
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/clouddeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clouddeploy.NewDeliveryPipeline(ctx, "primary", &clouddeploy.DeliveryPipelineArgs{
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("basic description"),
//				Project:     pulumi.String("my-project-name"),
//				SerialPipeline: &clouddeploy.DeliveryPipelineSerialPipelineArgs{
//					Stages: clouddeploy.DeliveryPipelineSerialPipelineStageArray{
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							DeployParameters: clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArray{
//								&clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArgs{
//									Values: pulumi.StringMap{
//										"deployParameterKey": pulumi.String("deployParameterValue"),
//									},
//									MatchTargetLabels: nil,
//								},
//							},
//							Profiles: pulumi.StringArray{
//								pulumi.String("example-profile-one"),
//								pulumi.String("example-profile-two"),
//							},
//							TargetId: pulumi.String("example-target-one"),
//						},
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							Profiles: pulumi.StringArray{},
//							TargetId: pulumi.String("example-target-two"),
//						},
//					},
//				},
//				Annotations: pulumi.StringMap{
//					"my_first_annotation":  pulumi.String("example-annotation-1"),
//					"my_second_annotation": pulumi.String("example-annotation-2"),
//				},
//				Labels: pulumi.StringMap{
//					"my_first_label":  pulumi.String("example-label-1"),
//					"my_second_label": pulumi.String("example-label-2"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Delivery_pipeline
// Creates a basic Cloud Deploy delivery pipeline
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/clouddeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clouddeploy.NewDeliveryPipeline(ctx, "primary", &clouddeploy.DeliveryPipelineArgs{
//				Annotations: pulumi.StringMap{
//					"my_first_annotation":  pulumi.String("example-annotation-1"),
//					"my_second_annotation": pulumi.String("example-annotation-2"),
//				},
//				Description: pulumi.String("basic description"),
//				Labels: pulumi.StringMap{
//					"my_first_label":  pulumi.String("example-label-1"),
//					"my_second_label": pulumi.String("example-label-2"),
//				},
//				Location: pulumi.String("us-west1"),
//				Project:  pulumi.String("my-project-name"),
//				SerialPipeline: &clouddeploy.DeliveryPipelineSerialPipelineArgs{
//					Stages: clouddeploy.DeliveryPipelineSerialPipelineStageArray{
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							DeployParameters: clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArray{
//								&clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArgs{
//									MatchTargetLabels: nil,
//									Values: pulumi.StringMap{
//										"deployParameterKey": pulumi.String("deployParameterValue"),
//									},
//								},
//							},
//							Profiles: pulumi.StringArray{
//								pulumi.String("example-profile-one"),
//								pulumi.String("example-profile-two"),
//							},
//							TargetId: pulumi.String("example-target-one"),
//						},
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							Profiles: pulumi.StringArray{},
//							TargetId: pulumi.String("example-target-two"),
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
// ### Verify_delivery_pipeline
// tests creating and updating a delivery pipeline with deployment verification strategy
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/clouddeploy"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clouddeploy.NewDeliveryPipeline(ctx, "primary", &clouddeploy.DeliveryPipelineArgs{
//				Location:    pulumi.String("us-west1"),
//				Description: pulumi.String("basic description"),
//				Project:     pulumi.String("my-project-name"),
//				SerialPipeline: &clouddeploy.DeliveryPipelineSerialPipelineArgs{
//					Stages: clouddeploy.DeliveryPipelineSerialPipelineStageArray{
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							DeployParameters: clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArray{
//								&clouddeploy.DeliveryPipelineSerialPipelineStageDeployParameterArgs{
//									Values: pulumi.StringMap{
//										"deployParameterKey": pulumi.String("deployParameterValue"),
//									},
//									MatchTargetLabels: nil,
//								},
//							},
//							Profiles: pulumi.StringArray{
//								pulumi.String("example-profile-one"),
//								pulumi.String("example-profile-two"),
//							},
//							TargetId: pulumi.String("example-target-one"),
//						},
//						&clouddeploy.DeliveryPipelineSerialPipelineStageArgs{
//							Profiles: pulumi.StringArray{},
//							TargetId: pulumi.String("example-target-two"),
//						},
//					},
//				},
//				Annotations: pulumi.StringMap{
//					"my_first_annotation":  pulumi.String("example-annotation-1"),
//					"my_second_annotation": pulumi.String("example-annotation-2"),
//				},
//				Labels: pulumi.StringMap{
//					"my_first_label":  pulumi.String("example-label-1"),
//					"my_second_label": pulumi.String("example-label-2"),
//				},
//			}, pulumi.Provider(google_beta))
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
// # DeliveryPipeline can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:clouddeploy/deliveryPipeline:DeliveryPipeline default projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:clouddeploy/deliveryPipeline:DeliveryPipeline default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:clouddeploy/deliveryPipeline:DeliveryPipeline default {{location}}/{{name}}
//
// ```
type DeliveryPipeline struct {
	pulumi.CustomResourceState

	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Output only. Information around the state of the Delivery Pipeline.
	Conditions DeliveryPipelineConditionArrayOutput `pulumi:"conditions"`
	// Output only. Time at which the pipeline was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
	// Terraform, other clients and services.
	EffectiveAnnotations pulumi.MapOutput `pulumi:"effectiveAnnotations"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.MapOutput `pulumi:"effectiveLabels"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapOutput `pulumi:"pulumiLabels"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline DeliveryPipelineSerialPipelinePtrOutput `pulumi:"serialPipeline"`
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended pulumi.BoolPtrOutput `pulumi:"suspended"`
	// Output only. Unique identifier of the `DeliveryPipeline`.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. Most recent time at which the pipeline was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDeliveryPipeline registers a new resource with the given unique name, arguments, and options.
func NewDeliveryPipeline(ctx *pulumi.Context,
	name string, args *DeliveryPipelineArgs, opts ...pulumi.ResourceOption) (*DeliveryPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeliveryPipeline
	err := ctx.RegisterResource("gcp:clouddeploy/deliveryPipeline:DeliveryPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryPipeline gets an existing DeliveryPipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryPipelineState, opts ...pulumi.ResourceOption) (*DeliveryPipeline, error) {
	var resource DeliveryPipeline
	err := ctx.ReadResource("gcp:clouddeploy/deliveryPipeline:DeliveryPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryPipeline resources.
type deliveryPipelineState struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Output only. Information around the state of the Delivery Pipeline.
	Conditions []DeliveryPipelineCondition `pulumi:"conditions"`
	// Output only. Time at which the pipeline was created.
	CreateTime *string `pulumi:"createTime"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
	// Terraform, other clients and services.
	EffectiveAnnotations map[string]interface{} `pulumi:"effectiveAnnotations"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]interface{} `pulumi:"effectiveLabels"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels map[string]interface{} `pulumi:"pulumiLabels"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline *DeliveryPipelineSerialPipeline `pulumi:"serialPipeline"`
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended *bool `pulumi:"suspended"`
	// Output only. Unique identifier of the `DeliveryPipeline`.
	Uid *string `pulumi:"uid"`
	// Output only. Most recent time at which the pipeline was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type DeliveryPipelineState struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Output only. Information around the state of the Delivery Pipeline.
	Conditions DeliveryPipelineConditionArrayInput
	// Output only. Time at which the pipeline was created.
	CreateTime pulumi.StringPtrInput
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
	// Terraform, other clients and services.
	EffectiveAnnotations pulumi.MapInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.MapInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapInput
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline DeliveryPipelineSerialPipelinePtrInput
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended pulumi.BoolPtrInput
	// Output only. Unique identifier of the `DeliveryPipeline`.
	Uid pulumi.StringPtrInput
	// Output only. Most recent time at which the pipeline was updated.
	UpdateTime pulumi.StringPtrInput
}

func (DeliveryPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryPipelineState)(nil)).Elem()
}

type deliveryPipelineArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline *DeliveryPipelineSerialPipeline `pulumi:"serialPipeline"`
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended *bool `pulumi:"suspended"`
}

// The set of arguments for constructing a DeliveryPipeline resource.
type DeliveryPipelineArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline DeliveryPipelineSerialPipelinePtrInput
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended pulumi.BoolPtrInput
}

func (DeliveryPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryPipelineArgs)(nil)).Elem()
}

type DeliveryPipelineInput interface {
	pulumi.Input

	ToDeliveryPipelineOutput() DeliveryPipelineOutput
	ToDeliveryPipelineOutputWithContext(ctx context.Context) DeliveryPipelineOutput
}

func (*DeliveryPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPipeline)(nil)).Elem()
}

func (i *DeliveryPipeline) ToDeliveryPipelineOutput() DeliveryPipelineOutput {
	return i.ToDeliveryPipelineOutputWithContext(context.Background())
}

func (i *DeliveryPipeline) ToDeliveryPipelineOutputWithContext(ctx context.Context) DeliveryPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPipelineOutput)
}

func (i *DeliveryPipeline) ToOutput(ctx context.Context) pulumix.Output[*DeliveryPipeline] {
	return pulumix.Output[*DeliveryPipeline]{
		OutputState: i.ToDeliveryPipelineOutputWithContext(ctx).OutputState,
	}
}

// DeliveryPipelineArrayInput is an input type that accepts DeliveryPipelineArray and DeliveryPipelineArrayOutput values.
// You can construct a concrete instance of `DeliveryPipelineArrayInput` via:
//
//	DeliveryPipelineArray{ DeliveryPipelineArgs{...} }
type DeliveryPipelineArrayInput interface {
	pulumi.Input

	ToDeliveryPipelineArrayOutput() DeliveryPipelineArrayOutput
	ToDeliveryPipelineArrayOutputWithContext(context.Context) DeliveryPipelineArrayOutput
}

type DeliveryPipelineArray []DeliveryPipelineInput

func (DeliveryPipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeliveryPipeline)(nil)).Elem()
}

func (i DeliveryPipelineArray) ToDeliveryPipelineArrayOutput() DeliveryPipelineArrayOutput {
	return i.ToDeliveryPipelineArrayOutputWithContext(context.Background())
}

func (i DeliveryPipelineArray) ToDeliveryPipelineArrayOutputWithContext(ctx context.Context) DeliveryPipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPipelineArrayOutput)
}

func (i DeliveryPipelineArray) ToOutput(ctx context.Context) pulumix.Output[[]*DeliveryPipeline] {
	return pulumix.Output[[]*DeliveryPipeline]{
		OutputState: i.ToDeliveryPipelineArrayOutputWithContext(ctx).OutputState,
	}
}

// DeliveryPipelineMapInput is an input type that accepts DeliveryPipelineMap and DeliveryPipelineMapOutput values.
// You can construct a concrete instance of `DeliveryPipelineMapInput` via:
//
//	DeliveryPipelineMap{ "key": DeliveryPipelineArgs{...} }
type DeliveryPipelineMapInput interface {
	pulumi.Input

	ToDeliveryPipelineMapOutput() DeliveryPipelineMapOutput
	ToDeliveryPipelineMapOutputWithContext(context.Context) DeliveryPipelineMapOutput
}

type DeliveryPipelineMap map[string]DeliveryPipelineInput

func (DeliveryPipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeliveryPipeline)(nil)).Elem()
}

func (i DeliveryPipelineMap) ToDeliveryPipelineMapOutput() DeliveryPipelineMapOutput {
	return i.ToDeliveryPipelineMapOutputWithContext(context.Background())
}

func (i DeliveryPipelineMap) ToDeliveryPipelineMapOutputWithContext(ctx context.Context) DeliveryPipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPipelineMapOutput)
}

func (i DeliveryPipelineMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DeliveryPipeline] {
	return pulumix.Output[map[string]*DeliveryPipeline]{
		OutputState: i.ToDeliveryPipelineMapOutputWithContext(ctx).OutputState,
	}
}

type DeliveryPipelineOutput struct{ *pulumi.OutputState }

func (DeliveryPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPipeline)(nil)).Elem()
}

func (o DeliveryPipelineOutput) ToDeliveryPipelineOutput() DeliveryPipelineOutput {
	return o
}

func (o DeliveryPipelineOutput) ToDeliveryPipelineOutputWithContext(ctx context.Context) DeliveryPipelineOutput {
	return o
}

func (o DeliveryPipelineOutput) ToOutput(ctx context.Context) pulumix.Output[*DeliveryPipeline] {
	return pulumix.Output[*DeliveryPipeline]{
		OutputState: o.OutputState,
	}
}

// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
//
// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
// Please refer to the field `effectiveAnnotations` for all of the annotations present on the resource.
func (o DeliveryPipelineOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Output only. Information around the state of the Delivery Pipeline.
func (o DeliveryPipelineOutput) Conditions() DeliveryPipelineConditionArrayOutput {
	return o.ApplyT(func(v *DeliveryPipeline) DeliveryPipelineConditionArrayOutput { return v.Conditions }).(DeliveryPipelineConditionArrayOutput)
}

// Output only. Time at which the pipeline was created.
func (o DeliveryPipelineOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the `DeliveryPipeline`. Max length is 255 characters.
func (o DeliveryPipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
// Terraform, other clients and services.
func (o DeliveryPipelineOutput) EffectiveAnnotations() pulumi.MapOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.MapOutput { return v.EffectiveAnnotations }).(pulumi.MapOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o DeliveryPipelineOutput) EffectiveLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.MapOutput { return v.EffectiveLabels }).(pulumi.MapOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o DeliveryPipelineOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o DeliveryPipelineOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
func (o DeliveryPipelineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
func (o DeliveryPipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o DeliveryPipelineOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DeliveryPipelineOutput) PulumiLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.MapOutput { return v.PulumiLabels }).(pulumi.MapOutput)
}

// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
func (o DeliveryPipelineOutput) SerialPipeline() DeliveryPipelineSerialPipelinePtrOutput {
	return o.ApplyT(func(v *DeliveryPipeline) DeliveryPipelineSerialPipelinePtrOutput { return v.SerialPipeline }).(DeliveryPipelineSerialPipelinePtrOutput)
}

// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
func (o DeliveryPipelineOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.BoolPtrOutput { return v.Suspended }).(pulumi.BoolPtrOutput)
}

// Output only. Unique identifier of the `DeliveryPipeline`.
func (o DeliveryPipelineOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. Most recent time at which the pipeline was updated.
func (o DeliveryPipelineOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryPipeline) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type DeliveryPipelineArrayOutput struct{ *pulumi.OutputState }

func (DeliveryPipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeliveryPipeline)(nil)).Elem()
}

func (o DeliveryPipelineArrayOutput) ToDeliveryPipelineArrayOutput() DeliveryPipelineArrayOutput {
	return o
}

func (o DeliveryPipelineArrayOutput) ToDeliveryPipelineArrayOutputWithContext(ctx context.Context) DeliveryPipelineArrayOutput {
	return o
}

func (o DeliveryPipelineArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DeliveryPipeline] {
	return pulumix.Output[[]*DeliveryPipeline]{
		OutputState: o.OutputState,
	}
}

func (o DeliveryPipelineArrayOutput) Index(i pulumi.IntInput) DeliveryPipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeliveryPipeline {
		return vs[0].([]*DeliveryPipeline)[vs[1].(int)]
	}).(DeliveryPipelineOutput)
}

type DeliveryPipelineMapOutput struct{ *pulumi.OutputState }

func (DeliveryPipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeliveryPipeline)(nil)).Elem()
}

func (o DeliveryPipelineMapOutput) ToDeliveryPipelineMapOutput() DeliveryPipelineMapOutput {
	return o
}

func (o DeliveryPipelineMapOutput) ToDeliveryPipelineMapOutputWithContext(ctx context.Context) DeliveryPipelineMapOutput {
	return o
}

func (o DeliveryPipelineMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DeliveryPipeline] {
	return pulumix.Output[map[string]*DeliveryPipeline]{
		OutputState: o.OutputState,
	}
}

func (o DeliveryPipelineMapOutput) MapIndex(k pulumi.StringInput) DeliveryPipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeliveryPipeline {
		return vs[0].(map[string]*DeliveryPipeline)[vs[1].(string)]
	}).(DeliveryPipelineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryPipelineInput)(nil)).Elem(), &DeliveryPipeline{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryPipelineArrayInput)(nil)).Elem(), DeliveryPipelineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryPipelineMapInput)(nil)).Elem(), DeliveryPipelineMap{})
	pulumi.RegisterOutputType(DeliveryPipelineOutput{})
	pulumi.RegisterOutputType(DeliveryPipelineArrayOutput{})
	pulumi.RegisterOutputType(DeliveryPipelineMapOutput{})
}
