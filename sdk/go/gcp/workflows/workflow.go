// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workflows

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Workflow program to be executed by Workflows.
//
// To get more information about Workflow, see:
//
// * [API documentation](https://cloud.google.com/workflows/docs/reference/rest/v1/projects.locations.workflows)
// * How-to Guides
//   - [Managing Workflows](https://cloud.google.com/workflows/docs/creating-updating-workflow)
//
// ## Example Usage
// ### Workflow Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/workflows"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testAccount, err := serviceaccount.NewAccount(ctx, "testAccount", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-account"),
//				DisplayName: pulumi.String("Test Service Account"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = workflows.NewWorkflow(ctx, "example", &workflows.WorkflowArgs{
//				Region:         pulumi.String("us-central1"),
//				Description:    pulumi.String("Magic"),
//				ServiceAccount: testAccount.ID(),
//				Labels: pulumi.StringMap{
//					"env": pulumi.String("test"),
//				},
//				SourceContents: pulumi.String(fmt.Sprintf(`# This is a sample workflow. You can replace it with your source code.
//
// #
// # This workflow does the following:
// # - reads current time and date information from an external API and stores
// #   the response in currentTime variable
// # - retrieves a list of Wikipedia articles related to the day of the week
// #   from currentTime
// # - returns the list of articles as an output of the workflow
// #
// # Note: In Terraform you need to escape the $$ or it will cause errors.
//
//   - getCurrentTime:
//     call: http.get
//     args:
//     url: https://timeapi.io/api/Time/current/zone?timeZone=Europe/Amsterdam
//     result: currentTime
//   - readWikipedia:
//     call: http.get
//     args:
//     url: https://en.wikipedia.org/w/api.php
//     query:
//     action: opensearch
//     search: %v
//     result: wikiResult
//   - returnOutput:
//     return: %v
//
// `, currentTime.Body.DayOfWeek, wikiResult.Body[1])),
//
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
// This resource does not support import.
type Workflow struct {
	pulumi.CustomResourceState

	// The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	CryptoKeyName pulumi.StringPtrOutput `pulumi:"cryptoKeyName"`
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description pulumi.StringOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// A set of key/value label pairs to assign to this Workflow.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the Workflow.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The region of the workflow.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The revision of the workflow. A new one is generated if the service account or source contents is changed.
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the uniqueId of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Workflow code to be executed. The size limit is 32KB.
	SourceContents pulumi.StringPtrOutput `pulumi:"sourceContents"`
	// State of the workflow deployment.
	State pulumi.StringOutput `pulumi:"state"`
	// The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		args = &WorkflowArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Workflow
	err := ctx.RegisterResource("gcp:workflows/workflow:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("gcp:workflows/workflow:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
	// The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	CryptoKeyName *string `pulumi:"cryptoKeyName"`
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// A set of key/value label pairs to assign to this Workflow.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Workflow.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The region of the workflow.
	Region *string `pulumi:"region"`
	// The revision of the workflow. A new one is generated if the service account or source contents is changed.
	RevisionId *string `pulumi:"revisionId"`
	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the uniqueId of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Workflow code to be executed. The size limit is 32KB.
	SourceContents *string `pulumi:"sourceContents"`
	// State of the workflow deployment.
	State *string `pulumi:"state"`
	// The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `pulumi:"updateTime"`
}

type WorkflowState struct {
	// The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	CryptoKeyName pulumi.StringPtrInput
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// A set of key/value label pairs to assign to this Workflow.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the Workflow.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	NamePrefix pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The region of the workflow.
	Region pulumi.StringPtrInput
	// The revision of the workflow. A new one is generated if the service account or source contents is changed.
	RevisionId pulumi.StringPtrInput
	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the uniqueId of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount pulumi.StringPtrInput
	// Workflow code to be executed. The size limit is 32KB.
	SourceContents pulumi.StringPtrInput
	// State of the workflow deployment.
	State pulumi.StringPtrInput
	// The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringPtrInput
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	CryptoKeyName *string `pulumi:"cryptoKeyName"`
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description *string `pulumi:"description"`
	// A set of key/value label pairs to assign to this Workflow.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Workflow.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the workflow.
	Region *string `pulumi:"region"`
	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the uniqueId of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// Workflow code to be executed. The size limit is 32KB.
	SourceContents *string `pulumi:"sourceContents"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// The KMS key used to encrypt workflow and execution data.
	// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	CryptoKeyName pulumi.StringPtrInput
	// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
	Description pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this Workflow.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Name of the Workflow.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the
	// specified prefix. If this and name are unspecified, a random value is chosen for the name.
	NamePrefix pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the workflow.
	Region pulumi.StringPtrInput
	// Name of the service account associated with the latest workflow version. This service
	// account represents the identity of the workflow and determines what permissions the workflow has.
	// Format: projects/{project}/serviceAccounts/{account} or {account}.
	// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
	// The {account} value can be the email address or the uniqueId of the service account.
	// If not provided, workflow will use the project's default service account.
	// Modifying this field for an existing workflow results in a new workflow revision.
	ServiceAccount pulumi.StringPtrInput
	// Workflow code to be executed. The size limit is 32KB.
	SourceContents pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

func (i *Workflow) ToOutput(ctx context.Context) pulumix.Output[*Workflow] {
	return pulumix.Output[*Workflow]{
		OutputState: i.ToWorkflowOutputWithContext(ctx).OutputState,
	}
}

// WorkflowArrayInput is an input type that accepts WorkflowArray and WorkflowArrayOutput values.
// You can construct a concrete instance of `WorkflowArrayInput` via:
//
//	WorkflowArray{ WorkflowArgs{...} }
type WorkflowArrayInput interface {
	pulumi.Input

	ToWorkflowArrayOutput() WorkflowArrayOutput
	ToWorkflowArrayOutputWithContext(context.Context) WorkflowArrayOutput
}

type WorkflowArray []WorkflowInput

func (WorkflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workflow)(nil)).Elem()
}

func (i WorkflowArray) ToWorkflowArrayOutput() WorkflowArrayOutput {
	return i.ToWorkflowArrayOutputWithContext(context.Background())
}

func (i WorkflowArray) ToWorkflowArrayOutputWithContext(ctx context.Context) WorkflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowArrayOutput)
}

func (i WorkflowArray) ToOutput(ctx context.Context) pulumix.Output[[]*Workflow] {
	return pulumix.Output[[]*Workflow]{
		OutputState: i.ToWorkflowArrayOutputWithContext(ctx).OutputState,
	}
}

// WorkflowMapInput is an input type that accepts WorkflowMap and WorkflowMapOutput values.
// You can construct a concrete instance of `WorkflowMapInput` via:
//
//	WorkflowMap{ "key": WorkflowArgs{...} }
type WorkflowMapInput interface {
	pulumi.Input

	ToWorkflowMapOutput() WorkflowMapOutput
	ToWorkflowMapOutputWithContext(context.Context) WorkflowMapOutput
}

type WorkflowMap map[string]WorkflowInput

func (WorkflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workflow)(nil)).Elem()
}

func (i WorkflowMap) ToWorkflowMapOutput() WorkflowMapOutput {
	return i.ToWorkflowMapOutputWithContext(context.Background())
}

func (i WorkflowMap) ToWorkflowMapOutputWithContext(ctx context.Context) WorkflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowMapOutput)
}

func (i WorkflowMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workflow] {
	return pulumix.Output[map[string]*Workflow]{
		OutputState: i.ToWorkflowMapOutputWithContext(ctx).OutputState,
	}
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToOutput(ctx context.Context) pulumix.Output[*Workflow] {
	return pulumix.Output[*Workflow]{
		OutputState: o.OutputState,
	}
}

// The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o WorkflowOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The KMS key used to encrypt workflow and execution data.
// Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
func (o WorkflowOutput) CryptoKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.CryptoKeyName }).(pulumi.StringPtrOutput)
}

// Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
func (o WorkflowOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o WorkflowOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// A set of key/value label pairs to assign to this Workflow.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o WorkflowOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the Workflow.
func (o WorkflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the
// specified prefix. If this and name are unspecified, a random value is chosen for the name.
func (o WorkflowOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o WorkflowOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o WorkflowOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The region of the workflow.
func (o WorkflowOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The revision of the workflow. A new one is generated if the service account or source contents is changed.
func (o WorkflowOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// Name of the service account associated with the latest workflow version. This service
// account represents the identity of the workflow and determines what permissions the workflow has.
// Format: projects/{project}/serviceAccounts/{account} or {account}.
// Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
// The {account} value can be the email address or the uniqueId of the service account.
// If not provided, workflow will use the project's default service account.
// Modifying this field for an existing workflow results in a new workflow revision.
func (o WorkflowOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Workflow code to be executed. The size limit is 32KB.
func (o WorkflowOutput) SourceContents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.SourceContents }).(pulumi.StringPtrOutput)
}

// State of the workflow deployment.
func (o WorkflowOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o WorkflowOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type WorkflowArrayOutput struct{ *pulumi.OutputState }

func (WorkflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workflow)(nil)).Elem()
}

func (o WorkflowArrayOutput) ToWorkflowArrayOutput() WorkflowArrayOutput {
	return o
}

func (o WorkflowArrayOutput) ToWorkflowArrayOutputWithContext(ctx context.Context) WorkflowArrayOutput {
	return o
}

func (o WorkflowArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Workflow] {
	return pulumix.Output[[]*Workflow]{
		OutputState: o.OutputState,
	}
}

func (o WorkflowArrayOutput) Index(i pulumi.IntInput) WorkflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workflow {
		return vs[0].([]*Workflow)[vs[1].(int)]
	}).(WorkflowOutput)
}

type WorkflowMapOutput struct{ *pulumi.OutputState }

func (WorkflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workflow)(nil)).Elem()
}

func (o WorkflowMapOutput) ToWorkflowMapOutput() WorkflowMapOutput {
	return o
}

func (o WorkflowMapOutput) ToWorkflowMapOutputWithContext(ctx context.Context) WorkflowMapOutput {
	return o
}

func (o WorkflowMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Workflow] {
	return pulumix.Output[map[string]*Workflow]{
		OutputState: o.OutputState,
	}
}

func (o WorkflowMapOutput) MapIndex(k pulumi.StringInput) WorkflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workflow {
		return vs[0].(map[string]*Workflow)[vs[1].(string)]
	}).(WorkflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowInput)(nil)).Elem(), &Workflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowArrayInput)(nil)).Elem(), WorkflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowMapInput)(nil)).Elem(), WorkflowMap{})
	pulumi.RegisterOutputType(WorkflowOutput{})
	pulumi.RegisterOutputType(WorkflowArrayOutput{})
	pulumi.RegisterOutputType(WorkflowMapOutput{})
}
