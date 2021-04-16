// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Denotes one policy tag in a taxonomy.
//
// To get more information about PolicyTag, see:
//
// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1beta1/projects.locations.taxonomies.policyTags)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
//
// ## Example Usage
// ### Data Catalog Taxonomies Policy Tag Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/datacatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myTaxonomy, err := datacatalog.NewTaxonomy(ctx, "myTaxonomy", &datacatalog.TaxonomyArgs{
// 			Region:      pulumi.String("us"),
// 			DisplayName: pulumi.String("taxonomy_display_name"),
// 			Description: pulumi.String("A collection of policy tags"),
// 			ActivatedPolicyTypes: pulumi.StringArray{
// 				pulumi.String("FINE_GRAINED_ACCESS_CONTROL"),
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datacatalog.NewPolicyTag(ctx, "basicPolicyTag", &datacatalog.PolicyTagArgs{
// 			Taxonomy:    myTaxonomy.ID(),
// 			DisplayName: pulumi.String("Low security"),
// 			Description: pulumi.String("A policy tag normally associated with low security items"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Data Catalog Taxonomies Policy Tag Child Policies
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/datacatalog"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myTaxonomy, err := datacatalog.NewTaxonomy(ctx, "myTaxonomy", &datacatalog.TaxonomyArgs{
// 			Region:      pulumi.String("us"),
// 			DisplayName: pulumi.String("taxonomy_display_name"),
// 			Description: pulumi.String("A collection of policy tags"),
// 			ActivatedPolicyTypes: pulumi.StringArray{
// 				pulumi.String("FINE_GRAINED_ACCESS_CONTROL"),
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		parentPolicy, err := datacatalog.NewPolicyTag(ctx, "parentPolicy", &datacatalog.PolicyTagArgs{
// 			Taxonomy:    myTaxonomy.ID(),
// 			DisplayName: pulumi.String("High"),
// 			Description: pulumi.String("A policy tag category used for high security access"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		childPolicy, err := datacatalog.NewPolicyTag(ctx, "childPolicy", &datacatalog.PolicyTagArgs{
// 			Taxonomy:        myTaxonomy.ID(),
// 			DisplayName:     pulumi.String("ssn"),
// 			Description:     pulumi.String("A hash of the users ssn"),
// 			ParentPolicyTag: parentPolicy.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datacatalog.NewPolicyTag(ctx, "childPolicy2", &datacatalog.PolicyTagArgs{
// 			Taxonomy:        myTaxonomy.ID(),
// 			DisplayName:     pulumi.String("dob"),
// 			Description:     pulumi.String("The users date of birth"),
// 			ParentPolicyTag: parentPolicy.ID(),
// 		}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
// 			childPolicy,
// 		}))
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
// PolicyTag can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:datacatalog/policyTag:PolicyTag default {{name}}
// ```
type PolicyTag struct {
	pulumi.CustomResourceState

	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags pulumi.StringArrayOutput `pulumi:"childPolicyTags"`
	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name of this policy tag, whose format is:
	// "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	ParentPolicyTag pulumi.StringPtrOutput `pulumi:"parentPolicyTag"`
	// Taxonomy the policy tag is associated with
	Taxonomy pulumi.StringOutput `pulumi:"taxonomy"`
}

// NewPolicyTag registers a new resource with the given unique name, arguments, and options.
func NewPolicyTag(ctx *pulumi.Context,
	name string, args *PolicyTagArgs, opts ...pulumi.ResourceOption) (*PolicyTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Taxonomy == nil {
		return nil, errors.New("invalid value for required argument 'Taxonomy'")
	}
	var resource PolicyTag
	err := ctx.RegisterResource("gcp:datacatalog/policyTag:PolicyTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyTag gets an existing PolicyTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyTagState, opts ...pulumi.ResourceOption) (*PolicyTag, error) {
	var resource PolicyTag
	err := ctx.ReadResource("gcp:datacatalog/policyTag:PolicyTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyTag resources.
type policyTagState struct {
	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags []string `pulumi:"childPolicyTags"`
	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description *string `pulumi:"description"`
	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName *string `pulumi:"displayName"`
	// Resource name of this policy tag, whose format is:
	// "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
	Name *string `pulumi:"name"`
	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	ParentPolicyTag *string `pulumi:"parentPolicyTag"`
	// Taxonomy the policy tag is associated with
	Taxonomy *string `pulumi:"taxonomy"`
}

type PolicyTagState struct {
	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags pulumi.StringArrayInput
	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description pulumi.StringPtrInput
	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringPtrInput
	// Resource name of this policy tag, whose format is:
	// "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
	Name pulumi.StringPtrInput
	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	ParentPolicyTag pulumi.StringPtrInput
	// Taxonomy the policy tag is associated with
	Taxonomy pulumi.StringPtrInput
}

func (PolicyTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagState)(nil)).Elem()
}

type policyTagArgs struct {
	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description *string `pulumi:"description"`
	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName string `pulumi:"displayName"`
	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	ParentPolicyTag *string `pulumi:"parentPolicyTag"`
	// Taxonomy the policy tag is associated with
	Taxonomy string `pulumi:"taxonomy"`
}

// The set of arguments for constructing a PolicyTag resource.
type PolicyTagArgs struct {
	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description pulumi.StringPtrInput
	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName pulumi.StringInput
	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	ParentPolicyTag pulumi.StringPtrInput
	// Taxonomy the policy tag is associated with
	Taxonomy pulumi.StringInput
}

func (PolicyTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyTagArgs)(nil)).Elem()
}

type PolicyTagInput interface {
	pulumi.Input

	ToPolicyTagOutput() PolicyTagOutput
	ToPolicyTagOutputWithContext(ctx context.Context) PolicyTagOutput
}

func (*PolicyTag) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTag)(nil))
}

func (i *PolicyTag) ToPolicyTagOutput() PolicyTagOutput {
	return i.ToPolicyTagOutputWithContext(context.Background())
}

func (i *PolicyTag) ToPolicyTagOutputWithContext(ctx context.Context) PolicyTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagOutput)
}

func (i *PolicyTag) ToPolicyTagPtrOutput() PolicyTagPtrOutput {
	return i.ToPolicyTagPtrOutputWithContext(context.Background())
}

func (i *PolicyTag) ToPolicyTagPtrOutputWithContext(ctx context.Context) PolicyTagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagPtrOutput)
}

type PolicyTagPtrInput interface {
	pulumi.Input

	ToPolicyTagPtrOutput() PolicyTagPtrOutput
	ToPolicyTagPtrOutputWithContext(ctx context.Context) PolicyTagPtrOutput
}

type policyTagPtrType PolicyTagArgs

func (*policyTagPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTag)(nil))
}

func (i *policyTagPtrType) ToPolicyTagPtrOutput() PolicyTagPtrOutput {
	return i.ToPolicyTagPtrOutputWithContext(context.Background())
}

func (i *policyTagPtrType) ToPolicyTagPtrOutputWithContext(ctx context.Context) PolicyTagPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagPtrOutput)
}

// PolicyTagArrayInput is an input type that accepts PolicyTagArray and PolicyTagArrayOutput values.
// You can construct a concrete instance of `PolicyTagArrayInput` via:
//
//          PolicyTagArray{ PolicyTagArgs{...} }
type PolicyTagArrayInput interface {
	pulumi.Input

	ToPolicyTagArrayOutput() PolicyTagArrayOutput
	ToPolicyTagArrayOutputWithContext(context.Context) PolicyTagArrayOutput
}

type PolicyTagArray []PolicyTagInput

func (PolicyTagArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PolicyTag)(nil))
}

func (i PolicyTagArray) ToPolicyTagArrayOutput() PolicyTagArrayOutput {
	return i.ToPolicyTagArrayOutputWithContext(context.Background())
}

func (i PolicyTagArray) ToPolicyTagArrayOutputWithContext(ctx context.Context) PolicyTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagArrayOutput)
}

// PolicyTagMapInput is an input type that accepts PolicyTagMap and PolicyTagMapOutput values.
// You can construct a concrete instance of `PolicyTagMapInput` via:
//
//          PolicyTagMap{ "key": PolicyTagArgs{...} }
type PolicyTagMapInput interface {
	pulumi.Input

	ToPolicyTagMapOutput() PolicyTagMapOutput
	ToPolicyTagMapOutputWithContext(context.Context) PolicyTagMapOutput
}

type PolicyTagMap map[string]PolicyTagInput

func (PolicyTagMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PolicyTag)(nil))
}

func (i PolicyTagMap) ToPolicyTagMapOutput() PolicyTagMapOutput {
	return i.ToPolicyTagMapOutputWithContext(context.Background())
}

func (i PolicyTagMap) ToPolicyTagMapOutputWithContext(ctx context.Context) PolicyTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTagMapOutput)
}

type PolicyTagOutput struct {
	*pulumi.OutputState
}

func (PolicyTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTag)(nil))
}

func (o PolicyTagOutput) ToPolicyTagOutput() PolicyTagOutput {
	return o
}

func (o PolicyTagOutput) ToPolicyTagOutputWithContext(ctx context.Context) PolicyTagOutput {
	return o
}

func (o PolicyTagOutput) ToPolicyTagPtrOutput() PolicyTagPtrOutput {
	return o.ToPolicyTagPtrOutputWithContext(context.Background())
}

func (o PolicyTagOutput) ToPolicyTagPtrOutputWithContext(ctx context.Context) PolicyTagPtrOutput {
	return o.ApplyT(func(v PolicyTag) *PolicyTag {
		return &v
	}).(PolicyTagPtrOutput)
}

type PolicyTagPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyTagPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTag)(nil))
}

func (o PolicyTagPtrOutput) ToPolicyTagPtrOutput() PolicyTagPtrOutput {
	return o
}

func (o PolicyTagPtrOutput) ToPolicyTagPtrOutputWithContext(ctx context.Context) PolicyTagPtrOutput {
	return o
}

type PolicyTagArrayOutput struct{ *pulumi.OutputState }

func (PolicyTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyTag)(nil))
}

func (o PolicyTagArrayOutput) ToPolicyTagArrayOutput() PolicyTagArrayOutput {
	return o
}

func (o PolicyTagArrayOutput) ToPolicyTagArrayOutputWithContext(ctx context.Context) PolicyTagArrayOutput {
	return o
}

func (o PolicyTagArrayOutput) Index(i pulumi.IntInput) PolicyTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyTag {
		return vs[0].([]PolicyTag)[vs[1].(int)]
	}).(PolicyTagOutput)
}

type PolicyTagMapOutput struct{ *pulumi.OutputState }

func (PolicyTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PolicyTag)(nil))
}

func (o PolicyTagMapOutput) ToPolicyTagMapOutput() PolicyTagMapOutput {
	return o
}

func (o PolicyTagMapOutput) ToPolicyTagMapOutputWithContext(ctx context.Context) PolicyTagMapOutput {
	return o
}

func (o PolicyTagMapOutput) MapIndex(k pulumi.StringInput) PolicyTagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PolicyTag {
		return vs[0].(map[string]PolicyTag)[vs[1].(string)]
	}).(PolicyTagOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyTagOutput{})
	pulumi.RegisterOutputType(PolicyTagPtrOutput{})
	pulumi.RegisterOutputType(PolicyTagArrayOutput{})
	pulumi.RegisterOutputType(PolicyTagMapOutput{})
}
