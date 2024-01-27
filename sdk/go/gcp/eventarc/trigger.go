// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventarc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Eventarc Trigger resource
//
// ## Example Usage
// ### Basic
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/eventarc"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudrun.NewService(ctx, "default", &cloudrun.ServiceArgs{
//				Location: pulumi.String("europe-west1"),
//				Metadata: &cloudrun.ServiceMetadataArgs{
//					Namespace: pulumi.String("my-project-name"),
//				},
//				Template: &cloudrun.ServiceTemplateArgs{
//					Spec: &cloudrun.ServiceTemplateSpecArgs{
//						Containers: cloudrun.ServiceTemplateSpecContainerArray{
//							&cloudrun.ServiceTemplateSpecContainerArgs{
//								Image: pulumi.String("gcr.io/cloudrun/hello"),
//								Ports: cloudrun.ServiceTemplateSpecContainerPortArray{
//									&cloudrun.ServiceTemplateSpecContainerPortArgs{
//										ContainerPort: pulumi.Int(8080),
//									},
//								},
//							},
//						},
//						ContainerConcurrency: pulumi.Int(50),
//						TimeoutSeconds:       pulumi.Int(100),
//					},
//				},
//				Traffics: cloudrun.ServiceTrafficArray{
//					&cloudrun.ServiceTrafficArgs{
//						Percent:        pulumi.Int(100),
//						LatestRevision: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eventarc.NewTrigger(ctx, "primary", &eventarc.TriggerArgs{
//				Location: pulumi.String("europe-west1"),
//				MatchingCriterias: eventarc.TriggerMatchingCriteriaArray{
//					&eventarc.TriggerMatchingCriteriaArgs{
//						Attribute: pulumi.String("type"),
//						Value:     pulumi.String("google.cloud.pubsub.topic.v1.messagePublished"),
//					},
//				},
//				Destination: &eventarc.TriggerDestinationArgs{
//					CloudRunService: &eventarc.TriggerDestinationCloudRunServiceArgs{
//						Service: _default.Name,
//						Region:  pulumi.String("europe-west1"),
//					},
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pubsub.NewTopic(ctx, "foo", nil)
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
// Trigger can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/triggers/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Trigger can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:eventarc/trigger:Trigger default projects/{{project}}/locations/{{location}}/triggers/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:eventarc/trigger:Trigger default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:eventarc/trigger:Trigger default {{location}}/{{name}}
//
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel pulumi.StringPtrOutput `pulumi:"channel"`
	// Output only. The reason(s) why a trigger is in FAILED state.
	Conditions pulumi.StringMapOutput `pulumi:"conditions"`
	// Output only. The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Required. Destination specifies where the events should be sent to.
	Destination TriggerDestinationOutput `pulumi:"destination"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.MapOutput `pulumi:"effectiveLabels"`
	// Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType pulumi.StringOutput `pulumi:"eventDataContentType"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriterias TriggerMatchingCriteriaArrayOutput `pulumi:"matchingCriterias"`
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapOutput `pulumi:"pulumiLabels"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	ServiceAccount pulumi.StringPtrOutput `pulumi:"serviceAccount"`
	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport TriggerTransportOutput `pulumi:"transport"`
	// Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.MatchingCriterias == nil {
		return nil, errors.New("invalid value for required argument 'MatchingCriterias'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trigger
	err := ctx.RegisterResource("gcp:eventarc/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("gcp:eventarc/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel *string `pulumi:"channel"`
	// Output only. The reason(s) why a trigger is in FAILED state.
	Conditions map[string]string `pulumi:"conditions"`
	// Output only. The creation time.
	CreateTime *string `pulumi:"createTime"`
	// Required. Destination specifies where the events should be sent to.
	Destination *TriggerDestination `pulumi:"destination"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]interface{} `pulumi:"effectiveLabels"`
	// Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType *string `pulumi:"eventDataContentType"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriterias []TriggerMatchingCriteria `pulumi:"matchingCriterias"`
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels map[string]interface{} `pulumi:"pulumiLabels"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport *TriggerTransport `pulumi:"transport"`
	// Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid *string `pulumi:"uid"`
	// Output only. The last-modified time.
	UpdateTime *string `pulumi:"updateTime"`
}

type TriggerState struct {
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel pulumi.StringPtrInput
	// Output only. The reason(s) why a trigger is in FAILED state.
	Conditions pulumi.StringMapInput
	// Output only. The creation time.
	CreateTime pulumi.StringPtrInput
	// Required. Destination specifies where the events should be sent to.
	Destination TriggerDestinationPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.MapInput
	// Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType pulumi.StringPtrInput
	// Optional. User labels attached to the triggers that can be used to group resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriterias TriggerMatchingCriteriaArrayInput
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	PulumiLabels pulumi.MapInput
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	ServiceAccount pulumi.StringPtrInput
	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport TriggerTransportPtrInput
	// Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	Uid pulumi.StringPtrInput
	// Output only. The last-modified time.
	UpdateTime pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel *string `pulumi:"channel"`
	// Required. Destination specifies where the events should be sent to.
	Destination TriggerDestination `pulumi:"destination"`
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType *string `pulumi:"eventDataContentType"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriterias []TriggerMatchingCriteria `pulumi:"matchingCriterias"`
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport *TriggerTransport `pulumi:"transport"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel pulumi.StringPtrInput
	// Required. Destination specifies where the events should be sent to.
	Destination TriggerDestinationInput
	// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
	EventDataContentType pulumi.StringPtrInput
	// Optional. User labels attached to the triggers that can be used to group resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriterias TriggerMatchingCriteriaArrayInput
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	ServiceAccount pulumi.StringPtrInput
	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport TriggerTransportPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//	TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//	TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

// Optional. The name of the channel associated with the trigger in `projects/{project}/locations/{location}/channels/{channel}` format. You must provide a channel to receive events from Eventarc SaaS partners.
func (o TriggerOutput) Channel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Channel }).(pulumi.StringPtrOutput)
}

// Output only. The reason(s) why a trigger is in FAILED state.
func (o TriggerOutput) Conditions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.Conditions }).(pulumi.StringMapOutput)
}

// Output only. The creation time.
func (o TriggerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Required. Destination specifies where the events should be sent to.
func (o TriggerOutput) Destination() TriggerDestinationOutput {
	return o.ApplyT(func(v *Trigger) TriggerDestinationOutput { return v.Destination }).(TriggerDestinationOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o TriggerOutput) EffectiveLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.MapOutput { return v.EffectiveLabels }).(pulumi.MapOutput)
}

// Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
func (o TriggerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. EventDataContentType specifies the type of payload in MIME format that is expected from the CloudEvent data field. This is set to `application/json` if the value is not defined.
func (o TriggerOutput) EventDataContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.EventDataContentType }).(pulumi.StringOutput)
}

// Optional. User labels attached to the triggers that can be used to group resources.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o TriggerOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
func (o TriggerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
func (o TriggerOutput) MatchingCriterias() TriggerMatchingCriteriaArrayOutput {
	return o.ApplyT(func(v *Trigger) TriggerMatchingCriteriaArrayOutput { return v.MatchingCriterias }).(TriggerMatchingCriteriaArrayOutput)
}

// Required. The resource name of the trigger. Must be unique within the location on the project.
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o TriggerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o TriggerOutput) PulumiLabels() pulumi.MapOutput {
	return o.ApplyT(func(v *Trigger) pulumi.MapOutput { return v.PulumiLabels }).(pulumi.MapOutput)
}

// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
func (o TriggerOutput) ServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.ServiceAccount }).(pulumi.StringPtrOutput)
}

// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
func (o TriggerOutput) Transport() TriggerTransportOutput {
	return o.ApplyT(func(v *Trigger) TriggerTransportOutput { return v.Transport }).(TriggerTransportOutput)
}

// Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
func (o TriggerOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The last-modified time.
func (o TriggerOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].([]*Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].(map[string]*Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerArrayInput)(nil)).Elem(), TriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerMapInput)(nil)).Elem(), TriggerMap{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}
