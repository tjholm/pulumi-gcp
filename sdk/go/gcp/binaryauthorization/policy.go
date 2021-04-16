// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A policy for container image binary authorization.
//
// To get more information about Policy, see:
//
// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/binary-authorization/)
//
// ## Example Usage
// ### Binary Authorization Policy Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/containeranalysis"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		note, err := containeranalysis.NewNote(ctx, "note", &containeranalysis.NoteArgs{
// 			AttestationAuthority: &containeranalysis.NoteAttestationAuthorityArgs{
// 				Hint: &containeranalysis.NoteAttestationAuthorityHintArgs{
// 					HumanReadableName: pulumi.String("My attestor"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		attestor, err := binaryauthorization.NewAttestor(ctx, "attestor", &binaryauthorization.AttestorArgs{
// 			AttestationAuthorityNote: &binaryauthorization.AttestorAttestationAuthorityNoteArgs{
// 				NoteReference: note.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = binaryauthorization.NewPolicy(ctx, "policy", &binaryauthorization.PolicyArgs{
// 			AdmissionWhitelistPatterns: binaryauthorization.PolicyAdmissionWhitelistPatternArray{
// 				&binaryauthorization.PolicyAdmissionWhitelistPatternArgs{
// 					NamePattern: pulumi.String("gcr.io/google_containers/*"),
// 				},
// 			},
// 			DefaultAdmissionRule: &binaryauthorization.PolicyDefaultAdmissionRuleArgs{
// 				EvaluationMode:  pulumi.String("ALWAYS_ALLOW"),
// 				EnforcementMode: pulumi.String("ENFORCED_BLOCK_AND_AUDIT_LOG"),
// 			},
// 			ClusterAdmissionRules: binaryauthorization.PolicyClusterAdmissionRuleArray{
// 				&binaryauthorization.PolicyClusterAdmissionRuleArgs{
// 					Cluster:         pulumi.String("us-central1-a.prod-cluster"),
// 					EvaluationMode:  pulumi.String("REQUIRE_ATTESTATION"),
// 					EnforcementMode: pulumi.String("ENFORCED_BLOCK_AND_AUDIT_LOG"),
// 					RequireAttestationsBies: pulumi.StringArray{
// 						attestor.Name,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Binary Authorization Policy Global Evaluation
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/containeranalysis"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		note, err := containeranalysis.NewNote(ctx, "note", &containeranalysis.NoteArgs{
// 			AttestationAuthority: &containeranalysis.NoteAttestationAuthorityArgs{
// 				Hint: &containeranalysis.NoteAttestationAuthorityHintArgs{
// 					HumanReadableName: pulumi.String("My attestor"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		attestor, err := binaryauthorization.NewAttestor(ctx, "attestor", &binaryauthorization.AttestorArgs{
// 			AttestationAuthorityNote: &binaryauthorization.AttestorAttestationAuthorityNoteArgs{
// 				NoteReference: note.Name,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = binaryauthorization.NewPolicy(ctx, "policy", &binaryauthorization.PolicyArgs{
// 			DefaultAdmissionRule: &binaryauthorization.PolicyDefaultAdmissionRuleArgs{
// 				EvaluationMode:  pulumi.String("REQUIRE_ATTESTATION"),
// 				EnforcementMode: pulumi.String("ENFORCED_BLOCK_AND_AUDIT_LOG"),
// 				RequireAttestationsBies: pulumi.StringArray{
// 					attestor.Name,
// 				},
// 			},
// 			GlobalPolicyEvaluationMode: pulumi.String("ENABLE"),
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
// Policy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
// ```
//
// ```sh
//  $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
// ```
type Policy struct {
	pulumi.CustomResourceState

	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns PolicyAdmissionWhitelistPatternArrayOutput `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules PolicyClusterAdmissionRuleArrayOutput `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRuleOutput `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are `ENABLE` and `DISABLE`.
	GlobalPolicyEvaluationMode pulumi.StringOutput `pulumi:"globalPolicyEvaluationMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAdmissionRule == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAdmissionRule'")
	}
	var resource Policy
	err := ctx.RegisterResource("gcp:binaryauthorization/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("gcp:binaryauthorization/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns []PolicyAdmissionWhitelistPattern `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules []PolicyClusterAdmissionRule `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule *PolicyDefaultAdmissionRule `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description *string `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are `ENABLE` and `DISABLE`.
	GlobalPolicyEvaluationMode *string `pulumi:"globalPolicyEvaluationMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type PolicyState struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns PolicyAdmissionWhitelistPatternArrayInput
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules PolicyClusterAdmissionRuleArrayInput
	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRulePtrInput
	// A descriptive comment.
	Description pulumi.StringPtrInput
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are `ENABLE` and `DISABLE`.
	GlobalPolicyEvaluationMode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns []PolicyAdmissionWhitelistPattern `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules []PolicyClusterAdmissionRule `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRule `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description *string `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are `ENABLE` and `DISABLE`.
	GlobalPolicyEvaluationMode *string `pulumi:"globalPolicyEvaluationMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.
	// Structure is documented below.
	AdmissionWhitelistPatterns PolicyAdmissionWhitelistPatternArrayInput
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules PolicyClusterAdmissionRuleArrayInput
	// Default admission rule for a cluster without a per-cluster admission
	// rule.
	// Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRuleInput
	// A descriptive comment.
	Description pulumi.StringPtrInput
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	// Possible values are `ENABLE` and `DISABLE`.
	GlobalPolicyEvaluationMode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i *Policy) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput
}

type policyPtrType PolicyArgs

func (*policyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (i *policyPtrType) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPtrType) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//          PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Policy)(nil))
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//          PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Policy)(nil))
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o.ToPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o.ApplyT(func(v Policy) *Policy {
		return &v
	}).(PolicyPtrOutput)
}

type PolicyPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Policy)(nil))
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Policy {
		return vs[0].([]Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Policy)(nil))
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Policy {
		return vs[0].(map[string]Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
