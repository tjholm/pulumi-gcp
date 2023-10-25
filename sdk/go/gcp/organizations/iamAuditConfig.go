// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Allows management of audit logging config for a given service for a Google Cloud Platform Organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewIamAuditConfig(ctx, "config", &organizations.IamAuditConfigArgs{
//				AuditLogConfigs: organizations.IamAuditConfigAuditLogConfigArray{
//					&organizations.IamAuditConfigAuditLogConfigArgs{
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@hashicorp.com"),
//						},
//						LogType: pulumi.String("DATA_READ"),
//					},
//				},
//				OrgId:   pulumi.String("your-organization-id"),
//				Service: pulumi.String("allServices"),
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
// IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//
//	$ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig config "your-organization-id foo.googleapis.com"
//
// ```
type IamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIamAuditConfig(ctx *pulumi.Context,
	name string, args *IamAuditConfigArgs, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamAuditConfig
	err := ctx.RegisterResource("gcp:organizations/iamAuditConfig:IamAuditConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamAuditConfig gets an existing IamAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamAuditConfigState, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	var resource IamAuditConfig
	err := ctx.ReadResource("gcp:organizations/iamAuditConfig:IamAuditConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamAuditConfig resources.
type iamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag *string `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId *string `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service *string `pulumi:"service"`
}

type IamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// The etag of iam policy
	Etag pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringPtrInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringPtrInput
}

func (IamAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigState)(nil)).Elem()
}

type iamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId string `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamAuditConfig resource.
type IamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringInput
}

func (IamAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigArgs)(nil)).Elem()
}

type IamAuditConfigInput interface {
	pulumi.Input

	ToIamAuditConfigOutput() IamAuditConfigOutput
	ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput
}

func (*IamAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAuditConfig)(nil)).Elem()
}

func (i *IamAuditConfig) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return i.ToIamAuditConfigOutputWithContext(context.Background())
}

func (i *IamAuditConfig) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigOutput)
}

func (i *IamAuditConfig) ToOutput(ctx context.Context) pulumix.Output[*IamAuditConfig] {
	return pulumix.Output[*IamAuditConfig]{
		OutputState: i.ToIamAuditConfigOutputWithContext(ctx).OutputState,
	}
}

// IamAuditConfigArrayInput is an input type that accepts IamAuditConfigArray and IamAuditConfigArrayOutput values.
// You can construct a concrete instance of `IamAuditConfigArrayInput` via:
//
//	IamAuditConfigArray{ IamAuditConfigArgs{...} }
type IamAuditConfigArrayInput interface {
	pulumi.Input

	ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput
	ToIamAuditConfigArrayOutputWithContext(context.Context) IamAuditConfigArrayOutput
}

type IamAuditConfigArray []IamAuditConfigInput

func (IamAuditConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamAuditConfig)(nil)).Elem()
}

func (i IamAuditConfigArray) ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput {
	return i.ToIamAuditConfigArrayOutputWithContext(context.Background())
}

func (i IamAuditConfigArray) ToIamAuditConfigArrayOutputWithContext(ctx context.Context) IamAuditConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigArrayOutput)
}

func (i IamAuditConfigArray) ToOutput(ctx context.Context) pulumix.Output[[]*IamAuditConfig] {
	return pulumix.Output[[]*IamAuditConfig]{
		OutputState: i.ToIamAuditConfigArrayOutputWithContext(ctx).OutputState,
	}
}

// IamAuditConfigMapInput is an input type that accepts IamAuditConfigMap and IamAuditConfigMapOutput values.
// You can construct a concrete instance of `IamAuditConfigMapInput` via:
//
//	IamAuditConfigMap{ "key": IamAuditConfigArgs{...} }
type IamAuditConfigMapInput interface {
	pulumi.Input

	ToIamAuditConfigMapOutput() IamAuditConfigMapOutput
	ToIamAuditConfigMapOutputWithContext(context.Context) IamAuditConfigMapOutput
}

type IamAuditConfigMap map[string]IamAuditConfigInput

func (IamAuditConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamAuditConfig)(nil)).Elem()
}

func (i IamAuditConfigMap) ToIamAuditConfigMapOutput() IamAuditConfigMapOutput {
	return i.ToIamAuditConfigMapOutputWithContext(context.Background())
}

func (i IamAuditConfigMap) ToIamAuditConfigMapOutputWithContext(ctx context.Context) IamAuditConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigMapOutput)
}

func (i IamAuditConfigMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IamAuditConfig] {
	return pulumix.Output[map[string]*IamAuditConfig]{
		OutputState: i.ToIamAuditConfigMapOutputWithContext(ctx).OutputState,
	}
}

type IamAuditConfigOutput struct{ *pulumi.OutputState }

func (IamAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAuditConfig)(nil)).Elem()
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return o
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return o
}

func (o IamAuditConfigOutput) ToOutput(ctx context.Context) pulumix.Output[*IamAuditConfig] {
	return pulumix.Output[*IamAuditConfig]{
		OutputState: o.OutputState,
	}
}

// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
func (o IamAuditConfigOutput) AuditLogConfigs() IamAuditConfigAuditLogConfigArrayOutput {
	return o.ApplyT(func(v *IamAuditConfig) IamAuditConfigAuditLogConfigArrayOutput { return v.AuditLogConfigs }).(IamAuditConfigAuditLogConfigArrayOutput)
}

// The etag of iam policy
func (o IamAuditConfigOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAuditConfig) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The numeric ID of the organization in which you want to manage the audit logging config.
func (o IamAuditConfigOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAuditConfig) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
func (o IamAuditConfigOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *IamAuditConfig) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

type IamAuditConfigArrayOutput struct{ *pulumi.OutputState }

func (IamAuditConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamAuditConfig)(nil)).Elem()
}

func (o IamAuditConfigArrayOutput) ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput {
	return o
}

func (o IamAuditConfigArrayOutput) ToIamAuditConfigArrayOutputWithContext(ctx context.Context) IamAuditConfigArrayOutput {
	return o
}

func (o IamAuditConfigArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IamAuditConfig] {
	return pulumix.Output[[]*IamAuditConfig]{
		OutputState: o.OutputState,
	}
}

func (o IamAuditConfigArrayOutput) Index(i pulumi.IntInput) IamAuditConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamAuditConfig {
		return vs[0].([]*IamAuditConfig)[vs[1].(int)]
	}).(IamAuditConfigOutput)
}

type IamAuditConfigMapOutput struct{ *pulumi.OutputState }

func (IamAuditConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamAuditConfig)(nil)).Elem()
}

func (o IamAuditConfigMapOutput) ToIamAuditConfigMapOutput() IamAuditConfigMapOutput {
	return o
}

func (o IamAuditConfigMapOutput) ToIamAuditConfigMapOutputWithContext(ctx context.Context) IamAuditConfigMapOutput {
	return o
}

func (o IamAuditConfigMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IamAuditConfig] {
	return pulumix.Output[map[string]*IamAuditConfig]{
		OutputState: o.OutputState,
	}
}

func (o IamAuditConfigMapOutput) MapIndex(k pulumi.StringInput) IamAuditConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamAuditConfig {
		return vs[0].(map[string]*IamAuditConfig)[vs[1].(string)]
	}).(IamAuditConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamAuditConfigInput)(nil)).Elem(), &IamAuditConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamAuditConfigArrayInput)(nil)).Elem(), IamAuditConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamAuditConfigMapInput)(nil)).Elem(), IamAuditConfigMap{})
	pulumi.RegisterOutputType(IamAuditConfigOutput{})
	pulumi.RegisterOutputType(IamAuditConfigArrayOutput{})
	pulumi.RegisterOutputType(IamAuditConfigMapOutput{})
}
