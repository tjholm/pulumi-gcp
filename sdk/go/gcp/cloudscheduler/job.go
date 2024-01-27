// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudscheduler

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A scheduled job that can publish a PubSub message or an HTTP request
// every X interval of time, using a crontab format string.
//
// To get more information about Job, see:
//
// * [API documentation](https://cloud.google.com/scheduler/docs/reference/rest/)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/scheduler/)
//
// ## Example Usage
// ### Scheduler Job App Engine
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudscheduler"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudscheduler.NewJob(ctx, "job", &cloudscheduler.JobArgs{
//				AppEngineHttpTarget: &cloudscheduler.JobAppEngineHttpTargetArgs{
//					AppEngineRouting: &cloudscheduler.JobAppEngineHttpTargetAppEngineRoutingArgs{
//						Instance: pulumi.String("my-instance-001"),
//						Service:  pulumi.String("web"),
//						Version:  pulumi.String("prod"),
//					},
//					HttpMethod:  pulumi.String("POST"),
//					RelativeUri: pulumi.String("/ping"),
//				},
//				AttemptDeadline: pulumi.String("320s"),
//				Description:     pulumi.String("test app engine job"),
//				RetryConfig: &cloudscheduler.JobRetryConfigArgs{
//					MaxDoublings:       pulumi.Int(2),
//					MaxRetryDuration:   pulumi.String("10s"),
//					MinBackoffDuration: pulumi.String("1s"),
//					RetryCount:         pulumi.Int(3),
//				},
//				Schedule: pulumi.String("*/4 * * * *"),
//				TimeZone: pulumi.String("Europe/London"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Scheduler Job Oauth
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudscheduler"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := compute.GetDefaultServiceAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudscheduler.NewJob(ctx, "job", &cloudscheduler.JobArgs{
//				Description:     pulumi.String("test http job"),
//				Schedule:        pulumi.String("*/8 * * * *"),
//				TimeZone:        pulumi.String("America/New_York"),
//				AttemptDeadline: pulumi.String("320s"),
//				HttpTarget: &cloudscheduler.JobHttpTargetArgs{
//					HttpMethod: pulumi.String("GET"),
//					Uri:        pulumi.String("https://cloudscheduler.googleapis.com/v1/projects/my-project-name/locations/us-west1/jobs"),
//					OauthToken: &cloudscheduler.JobHttpTargetOauthTokenArgs{
//						ServiceAccountEmail: *pulumi.String(_default.Email),
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
// ### Scheduler Job Oidc
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudscheduler"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := compute.GetDefaultServiceAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudscheduler.NewJob(ctx, "job", &cloudscheduler.JobArgs{
//				Description:     pulumi.String("test http job"),
//				Schedule:        pulumi.String("*/8 * * * *"),
//				TimeZone:        pulumi.String("America/New_York"),
//				AttemptDeadline: pulumi.String("320s"),
//				HttpTarget: &cloudscheduler.JobHttpTargetArgs{
//					HttpMethod: pulumi.String("GET"),
//					Uri:        pulumi.String("https://example.com/ping"),
//					OidcToken: &cloudscheduler.JobHttpTargetOidcTokenArgs{
//						ServiceAccountEmail: *pulumi.String(_default.Email),
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
//
// ## Import
//
// Job can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/jobs/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Job can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:cloudscheduler/job:Job default projects/{{project}}/locations/{{region}}/jobs/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudscheduler/job:Job default {{project}}/{{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudscheduler/job:Job default {{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudscheduler/job:Job default {{name}}
//
// ```
type Job struct {
	pulumi.CustomResourceState

	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// Structure is documented below.
	AppEngineHttpTarget JobAppEngineHttpTargetPtrOutput `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	//   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline pulumi.StringPtrOutput `pulumi:"attemptDeadline"`
	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// HTTP target.
	// If the job providers a httpTarget the cron will
	// send a request to the targeted url
	// Structure is documented below.
	HttpTarget JobHttpTargetPtrOutput `pulumi:"httpTarget"`
	// The name of the job.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused pulumi.BoolOutput `pulumi:"paused"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// Structure is documented below.
	PubsubTarget JobPubsubTargetPtrOutput `pulumi:"pubsubTarget"`
	// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
	Region pulumi.StringOutput `pulumi:"region"`
	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// Structure is documented below.
	RetryConfig JobRetryConfigPtrOutput `pulumi:"retryConfig"`
	// Describes the schedule on which the job will be executed.
	Schedule pulumi.StringPtrOutput `pulumi:"schedule"`
	// State of the job.
	State pulumi.StringOutput `pulumi:"state"`
	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	TimeZone pulumi.StringPtrOutput `pulumi:"timeZone"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		args = &JobArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("gcp:cloudscheduler/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("gcp:cloudscheduler/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// Structure is documented below.
	AppEngineHttpTarget *JobAppEngineHttpTarget `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	//   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline *string `pulumi:"attemptDeadline"`
	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	Description *string `pulumi:"description"`
	// HTTP target.
	// If the job providers a httpTarget the cron will
	// send a request to the targeted url
	// Structure is documented below.
	HttpTarget *JobHttpTarget `pulumi:"httpTarget"`
	// The name of the job.
	//
	// ***
	Name *string `pulumi:"name"`
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused *bool `pulumi:"paused"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// Structure is documented below.
	PubsubTarget *JobPubsubTarget `pulumi:"pubsubTarget"`
	// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
	Region *string `pulumi:"region"`
	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// Structure is documented below.
	RetryConfig *JobRetryConfig `pulumi:"retryConfig"`
	// Describes the schedule on which the job will be executed.
	Schedule *string `pulumi:"schedule"`
	// State of the job.
	State *string `pulumi:"state"`
	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	TimeZone *string `pulumi:"timeZone"`
}

type JobState struct {
	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// Structure is documented below.
	AppEngineHttpTarget JobAppEngineHttpTargetPtrInput
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	//   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline pulumi.StringPtrInput
	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	Description pulumi.StringPtrInput
	// HTTP target.
	// If the job providers a httpTarget the cron will
	// send a request to the targeted url
	// Structure is documented below.
	HttpTarget JobHttpTargetPtrInput
	// The name of the job.
	//
	// ***
	Name pulumi.StringPtrInput
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// Structure is documented below.
	PubsubTarget JobPubsubTargetPtrInput
	// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
	Region pulumi.StringPtrInput
	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// Structure is documented below.
	RetryConfig JobRetryConfigPtrInput
	// Describes the schedule on which the job will be executed.
	Schedule pulumi.StringPtrInput
	// State of the job.
	State pulumi.StringPtrInput
	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	TimeZone pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// Structure is documented below.
	AppEngineHttpTarget *JobAppEngineHttpTarget `pulumi:"appEngineHttpTarget"`
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	//   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline *string `pulumi:"attemptDeadline"`
	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	Description *string `pulumi:"description"`
	// HTTP target.
	// If the job providers a httpTarget the cron will
	// send a request to the targeted url
	// Structure is documented below.
	HttpTarget *JobHttpTarget `pulumi:"httpTarget"`
	// The name of the job.
	//
	// ***
	Name *string `pulumi:"name"`
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused *bool `pulumi:"paused"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// Structure is documented below.
	PubsubTarget *JobPubsubTarget `pulumi:"pubsubTarget"`
	// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
	Region *string `pulumi:"region"`
	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// Structure is documented below.
	RetryConfig *JobRetryConfig `pulumi:"retryConfig"`
	// Describes the schedule on which the job will be executed.
	Schedule *string `pulumi:"schedule"`
	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// Structure is documented below.
	AppEngineHttpTarget JobAppEngineHttpTargetPtrInput
	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	//   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	AttemptDeadline pulumi.StringPtrInput
	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	Description pulumi.StringPtrInput
	// HTTP target.
	// If the job providers a httpTarget the cron will
	// send a request to the targeted url
	// Structure is documented below.
	HttpTarget JobHttpTargetPtrInput
	// The name of the job.
	//
	// ***
	Name pulumi.StringPtrInput
	// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
	Paused pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// Structure is documented below.
	PubsubTarget JobPubsubTargetPtrInput
	// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
	Region pulumi.StringPtrInput
	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// Structure is documented below.
	RetryConfig JobRetryConfigPtrInput
	// Describes the schedule on which the job will be executed.
	Schedule pulumi.StringPtrInput
	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	TimeZone pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//	JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//	JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// App Engine HTTP target.
// If the job providers a App Engine HTTP target the cron will
// send a request to the service instance
// Structure is documented below.
func (o JobOutput) AppEngineHttpTarget() JobAppEngineHttpTargetPtrOutput {
	return o.ApplyT(func(v *Job) JobAppEngineHttpTargetPtrOutput { return v.AppEngineHttpTarget }).(JobAppEngineHttpTargetPtrOutput)
}

// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
// The allowed duration for this deadline is:
//   - For HTTP targets, between 15 seconds and 30 minutes.
//   - For App Engine HTTP targets, between 15 seconds and 24 hours.
//   - **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
//     A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
func (o JobOutput) AttemptDeadline() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.AttemptDeadline }).(pulumi.StringPtrOutput)
}

// A human-readable description for the job.
// This string must not contain more than 500 characters.
func (o JobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// HTTP target.
// If the job providers a httpTarget the cron will
// send a request to the targeted url
// Structure is documented below.
func (o JobOutput) HttpTarget() JobHttpTargetPtrOutput {
	return o.ApplyT(func(v *Job) JobHttpTargetPtrOutput { return v.HttpTarget }).(JobHttpTargetPtrOutput)
}

// The name of the job.
//
// ***
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sets the job to a paused state. Jobs default to being enabled when this property is not set.
func (o JobOutput) Paused() pulumi.BoolOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolOutput { return v.Paused }).(pulumi.BoolOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o JobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Pub/Sub target
// If the job providers a Pub/Sub target the cron will publish
// a message to the provided topic
// Structure is documented below.
func (o JobOutput) PubsubTarget() JobPubsubTargetPtrOutput {
	return o.ApplyT(func(v *Job) JobPubsubTargetPtrOutput { return v.PubsubTarget }).(JobPubsubTargetPtrOutput)
}

// Region where the scheduler job resides. If it is not provided, this provider will use the provider default.
func (o JobOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// By default, if a job does not complete successfully,
// meaning that an acknowledgement is not received from the handler,
// then it will be retried with exponential backoff according to the settings
// Structure is documented below.
func (o JobOutput) RetryConfig() JobRetryConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobRetryConfigPtrOutput { return v.RetryConfig }).(JobRetryConfigPtrOutput)
}

// Describes the schedule on which the job will be executed.
func (o JobOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Schedule }).(pulumi.StringPtrOutput)
}

// State of the job.
func (o JobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Specifies the time zone to be used in interpreting schedule.
// The value of this field must be a time zone name from the tz database.
func (o JobOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Job {
		return vs[0].([]*Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Job {
		return vs[0].(map[string]*Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
