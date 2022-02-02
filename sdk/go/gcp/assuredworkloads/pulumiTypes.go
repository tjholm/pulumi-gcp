// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package assuredworkloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadKmsSettings struct {
	// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
	NextRotationTime string `pulumi:"nextRotationTime"`
	// Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
	RotationPeriod string `pulumi:"rotationPeriod"`
}

// WorkloadKmsSettingsInput is an input type that accepts WorkloadKmsSettingsArgs and WorkloadKmsSettingsOutput values.
// You can construct a concrete instance of `WorkloadKmsSettingsInput` via:
//
//          WorkloadKmsSettingsArgs{...}
type WorkloadKmsSettingsInput interface {
	pulumi.Input

	ToWorkloadKmsSettingsOutput() WorkloadKmsSettingsOutput
	ToWorkloadKmsSettingsOutputWithContext(context.Context) WorkloadKmsSettingsOutput
}

type WorkloadKmsSettingsArgs struct {
	// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
	NextRotationTime pulumi.StringInput `pulumi:"nextRotationTime"`
	// Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
	RotationPeriod pulumi.StringInput `pulumi:"rotationPeriod"`
}

func (WorkloadKmsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadKmsSettings)(nil)).Elem()
}

func (i WorkloadKmsSettingsArgs) ToWorkloadKmsSettingsOutput() WorkloadKmsSettingsOutput {
	return i.ToWorkloadKmsSettingsOutputWithContext(context.Background())
}

func (i WorkloadKmsSettingsArgs) ToWorkloadKmsSettingsOutputWithContext(ctx context.Context) WorkloadKmsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadKmsSettingsOutput)
}

func (i WorkloadKmsSettingsArgs) ToWorkloadKmsSettingsPtrOutput() WorkloadKmsSettingsPtrOutput {
	return i.ToWorkloadKmsSettingsPtrOutputWithContext(context.Background())
}

func (i WorkloadKmsSettingsArgs) ToWorkloadKmsSettingsPtrOutputWithContext(ctx context.Context) WorkloadKmsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadKmsSettingsOutput).ToWorkloadKmsSettingsPtrOutputWithContext(ctx)
}

// WorkloadKmsSettingsPtrInput is an input type that accepts WorkloadKmsSettingsArgs, WorkloadKmsSettingsPtr and WorkloadKmsSettingsPtrOutput values.
// You can construct a concrete instance of `WorkloadKmsSettingsPtrInput` via:
//
//          WorkloadKmsSettingsArgs{...}
//
//  or:
//
//          nil
type WorkloadKmsSettingsPtrInput interface {
	pulumi.Input

	ToWorkloadKmsSettingsPtrOutput() WorkloadKmsSettingsPtrOutput
	ToWorkloadKmsSettingsPtrOutputWithContext(context.Context) WorkloadKmsSettingsPtrOutput
}

type workloadKmsSettingsPtrType WorkloadKmsSettingsArgs

func WorkloadKmsSettingsPtr(v *WorkloadKmsSettingsArgs) WorkloadKmsSettingsPtrInput {
	return (*workloadKmsSettingsPtrType)(v)
}

func (*workloadKmsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadKmsSettings)(nil)).Elem()
}

func (i *workloadKmsSettingsPtrType) ToWorkloadKmsSettingsPtrOutput() WorkloadKmsSettingsPtrOutput {
	return i.ToWorkloadKmsSettingsPtrOutputWithContext(context.Background())
}

func (i *workloadKmsSettingsPtrType) ToWorkloadKmsSettingsPtrOutputWithContext(ctx context.Context) WorkloadKmsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadKmsSettingsPtrOutput)
}

type WorkloadKmsSettingsOutput struct{ *pulumi.OutputState }

func (WorkloadKmsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadKmsSettings)(nil)).Elem()
}

func (o WorkloadKmsSettingsOutput) ToWorkloadKmsSettingsOutput() WorkloadKmsSettingsOutput {
	return o
}

func (o WorkloadKmsSettingsOutput) ToWorkloadKmsSettingsOutputWithContext(ctx context.Context) WorkloadKmsSettingsOutput {
	return o
}

func (o WorkloadKmsSettingsOutput) ToWorkloadKmsSettingsPtrOutput() WorkloadKmsSettingsPtrOutput {
	return o.ToWorkloadKmsSettingsPtrOutputWithContext(context.Background())
}

func (o WorkloadKmsSettingsOutput) ToWorkloadKmsSettingsPtrOutputWithContext(ctx context.Context) WorkloadKmsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadKmsSettings) *WorkloadKmsSettings {
		return &v
	}).(WorkloadKmsSettingsPtrOutput)
}

// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
func (o WorkloadKmsSettingsOutput) NextRotationTime() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadKmsSettings) string { return v.NextRotationTime }).(pulumi.StringOutput)
}

// Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
func (o WorkloadKmsSettingsOutput) RotationPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v WorkloadKmsSettings) string { return v.RotationPeriod }).(pulumi.StringOutput)
}

type WorkloadKmsSettingsPtrOutput struct{ *pulumi.OutputState }

func (WorkloadKmsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadKmsSettings)(nil)).Elem()
}

func (o WorkloadKmsSettingsPtrOutput) ToWorkloadKmsSettingsPtrOutput() WorkloadKmsSettingsPtrOutput {
	return o
}

func (o WorkloadKmsSettingsPtrOutput) ToWorkloadKmsSettingsPtrOutputWithContext(ctx context.Context) WorkloadKmsSettingsPtrOutput {
	return o
}

func (o WorkloadKmsSettingsPtrOutput) Elem() WorkloadKmsSettingsOutput {
	return o.ApplyT(func(v *WorkloadKmsSettings) WorkloadKmsSettings {
		if v != nil {
			return *v
		}
		var ret WorkloadKmsSettings
		return ret
	}).(WorkloadKmsSettingsOutput)
}

// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
func (o WorkloadKmsSettingsPtrOutput) NextRotationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadKmsSettings) *string {
		if v == nil {
			return nil
		}
		return &v.NextRotationTime
	}).(pulumi.StringPtrOutput)
}

// Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
func (o WorkloadKmsSettingsPtrOutput) RotationPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadKmsSettings) *string {
		if v == nil {
			return nil
		}
		return &v.RotationPeriod
	}).(pulumi.StringPtrOutput)
}

type WorkloadResource struct {
	// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
	ResourceId *int `pulumi:"resourceId"`
	// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	ResourceType *string `pulumi:"resourceType"`
}

// WorkloadResourceInput is an input type that accepts WorkloadResourceArgs and WorkloadResourceOutput values.
// You can construct a concrete instance of `WorkloadResourceInput` via:
//
//          WorkloadResourceArgs{...}
type WorkloadResourceInput interface {
	pulumi.Input

	ToWorkloadResourceOutput() WorkloadResourceOutput
	ToWorkloadResourceOutputWithContext(context.Context) WorkloadResourceOutput
}

type WorkloadResourceArgs struct {
	// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
	ResourceId pulumi.IntPtrInput `pulumi:"resourceId"`
	// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
}

func (WorkloadResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadResource)(nil)).Elem()
}

func (i WorkloadResourceArgs) ToWorkloadResourceOutput() WorkloadResourceOutput {
	return i.ToWorkloadResourceOutputWithContext(context.Background())
}

func (i WorkloadResourceArgs) ToWorkloadResourceOutputWithContext(ctx context.Context) WorkloadResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadResourceOutput)
}

// WorkloadResourceArrayInput is an input type that accepts WorkloadResourceArray and WorkloadResourceArrayOutput values.
// You can construct a concrete instance of `WorkloadResourceArrayInput` via:
//
//          WorkloadResourceArray{ WorkloadResourceArgs{...} }
type WorkloadResourceArrayInput interface {
	pulumi.Input

	ToWorkloadResourceArrayOutput() WorkloadResourceArrayOutput
	ToWorkloadResourceArrayOutputWithContext(context.Context) WorkloadResourceArrayOutput
}

type WorkloadResourceArray []WorkloadResourceInput

func (WorkloadResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadResource)(nil)).Elem()
}

func (i WorkloadResourceArray) ToWorkloadResourceArrayOutput() WorkloadResourceArrayOutput {
	return i.ToWorkloadResourceArrayOutputWithContext(context.Background())
}

func (i WorkloadResourceArray) ToWorkloadResourceArrayOutputWithContext(ctx context.Context) WorkloadResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadResourceArrayOutput)
}

type WorkloadResourceOutput struct{ *pulumi.OutputState }

func (WorkloadResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadResource)(nil)).Elem()
}

func (o WorkloadResourceOutput) ToWorkloadResourceOutput() WorkloadResourceOutput {
	return o
}

func (o WorkloadResourceOutput) ToWorkloadResourceOutputWithContext(ctx context.Context) WorkloadResourceOutput {
	return o
}

// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
func (o WorkloadResourceOutput) ResourceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkloadResource) *int { return v.ResourceId }).(pulumi.IntPtrOutput)
}

// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
func (o WorkloadResourceOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadResource) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type WorkloadResourceArrayOutput struct{ *pulumi.OutputState }

func (WorkloadResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadResource)(nil)).Elem()
}

func (o WorkloadResourceArrayOutput) ToWorkloadResourceArrayOutput() WorkloadResourceArrayOutput {
	return o
}

func (o WorkloadResourceArrayOutput) ToWorkloadResourceArrayOutputWithContext(ctx context.Context) WorkloadResourceArrayOutput {
	return o
}

func (o WorkloadResourceArrayOutput) Index(i pulumi.IntInput) WorkloadResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkloadResource {
		return vs[0].([]WorkloadResource)[vs[1].(int)]
	}).(WorkloadResourceOutput)
}

type WorkloadResourceSetting struct {
	// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
	ResourceId *string `pulumi:"resourceId"`
	// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	ResourceType *string `pulumi:"resourceType"`
}

// WorkloadResourceSettingInput is an input type that accepts WorkloadResourceSettingArgs and WorkloadResourceSettingOutput values.
// You can construct a concrete instance of `WorkloadResourceSettingInput` via:
//
//          WorkloadResourceSettingArgs{...}
type WorkloadResourceSettingInput interface {
	pulumi.Input

	ToWorkloadResourceSettingOutput() WorkloadResourceSettingOutput
	ToWorkloadResourceSettingOutputWithContext(context.Context) WorkloadResourceSettingOutput
}

type WorkloadResourceSettingArgs struct {
	// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
}

func (WorkloadResourceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadResourceSetting)(nil)).Elem()
}

func (i WorkloadResourceSettingArgs) ToWorkloadResourceSettingOutput() WorkloadResourceSettingOutput {
	return i.ToWorkloadResourceSettingOutputWithContext(context.Background())
}

func (i WorkloadResourceSettingArgs) ToWorkloadResourceSettingOutputWithContext(ctx context.Context) WorkloadResourceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadResourceSettingOutput)
}

// WorkloadResourceSettingArrayInput is an input type that accepts WorkloadResourceSettingArray and WorkloadResourceSettingArrayOutput values.
// You can construct a concrete instance of `WorkloadResourceSettingArrayInput` via:
//
//          WorkloadResourceSettingArray{ WorkloadResourceSettingArgs{...} }
type WorkloadResourceSettingArrayInput interface {
	pulumi.Input

	ToWorkloadResourceSettingArrayOutput() WorkloadResourceSettingArrayOutput
	ToWorkloadResourceSettingArrayOutputWithContext(context.Context) WorkloadResourceSettingArrayOutput
}

type WorkloadResourceSettingArray []WorkloadResourceSettingInput

func (WorkloadResourceSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadResourceSetting)(nil)).Elem()
}

func (i WorkloadResourceSettingArray) ToWorkloadResourceSettingArrayOutput() WorkloadResourceSettingArrayOutput {
	return i.ToWorkloadResourceSettingArrayOutputWithContext(context.Background())
}

func (i WorkloadResourceSettingArray) ToWorkloadResourceSettingArrayOutputWithContext(ctx context.Context) WorkloadResourceSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadResourceSettingArrayOutput)
}

type WorkloadResourceSettingOutput struct{ *pulumi.OutputState }

func (WorkloadResourceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadResourceSetting)(nil)).Elem()
}

func (o WorkloadResourceSettingOutput) ToWorkloadResourceSettingOutput() WorkloadResourceSettingOutput {
	return o
}

func (o WorkloadResourceSettingOutput) ToWorkloadResourceSettingOutputWithContext(ctx context.Context) WorkloadResourceSettingOutput {
	return o
}

// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
func (o WorkloadResourceSettingOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadResourceSetting) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
func (o WorkloadResourceSettingOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadResourceSetting) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type WorkloadResourceSettingArrayOutput struct{ *pulumi.OutputState }

func (WorkloadResourceSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadResourceSetting)(nil)).Elem()
}

func (o WorkloadResourceSettingArrayOutput) ToWorkloadResourceSettingArrayOutput() WorkloadResourceSettingArrayOutput {
	return o
}

func (o WorkloadResourceSettingArrayOutput) ToWorkloadResourceSettingArrayOutputWithContext(ctx context.Context) WorkloadResourceSettingArrayOutput {
	return o
}

func (o WorkloadResourceSettingArrayOutput) Index(i pulumi.IntInput) WorkloadResourceSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkloadResourceSetting {
		return vs[0].([]WorkloadResourceSetting)[vs[1].(int)]
	}).(WorkloadResourceSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadKmsSettingsInput)(nil)).Elem(), WorkloadKmsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadKmsSettingsPtrInput)(nil)).Elem(), WorkloadKmsSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadResourceInput)(nil)).Elem(), WorkloadResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadResourceArrayInput)(nil)).Elem(), WorkloadResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadResourceSettingInput)(nil)).Elem(), WorkloadResourceSettingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadResourceSettingArrayInput)(nil)).Elem(), WorkloadResourceSettingArray{})
	pulumi.RegisterOutputType(WorkloadKmsSettingsOutput{})
	pulumi.RegisterOutputType(WorkloadKmsSettingsPtrOutput{})
	pulumi.RegisterOutputType(WorkloadResourceOutput{})
	pulumi.RegisterOutputType(WorkloadResourceArrayOutput{})
	pulumi.RegisterOutputType(WorkloadResourceSettingOutput{})
	pulumi.RegisterOutputType(WorkloadResourceSettingArrayOutput{})
}
