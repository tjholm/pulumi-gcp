// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A repository (or repo) is a Git repository storing versioned source content.
//
// To get more information about Repository, see:
//
// * [API documentation](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/source-repositories/)
//
// ## Example Usage
// ### Sourcerepo Repository Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sourcerepo"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sourcerepo.NewRepository(ctx, "my-repo", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Sourcerepo Repository Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sourcerepo"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testAccount, err := serviceAccount.NewAccount(ctx, "testAccount", &serviceAccount.AccountArgs{
//				AccountId:   pulumi.String("my-account"),
//				DisplayName: pulumi.String("Test Service Account"),
//			})
//			if err != nil {
//				return err
//			}
//			topic, err := pubsub.NewTopic(ctx, "topic", nil)
//			if err != nil {
//				return err
//			}
//			_, err = sourcerepo.NewRepository(ctx, "my-repo", &sourcerepo.RepositoryArgs{
//				PubsubConfigs: sourcerepo.RepositoryPubsubConfigArray{
//					&sourcerepo.RepositoryPubsubConfigArgs{
//						Topic:               topic.ID(),
//						MessageFormat:       pulumi.String("JSON"),
//						ServiceAccountEmail: testAccount.Email,
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
// # Repository can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:sourcerepo/repository:Repository default projects/{{project}}/repos/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:sourcerepo/repository:Repository default {{name}}
//
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Resource name of the repository, of the form `{{repo}}`.
	// The repo name may contain slashes. eg, `name/with/slash`
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs RepositoryPubsubConfigArrayOutput `pulumi:"pubsubConfigs"`
	// The disk usage of the repo, in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// URL to clone the repository from Google Cloud Source Repositories.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		args = &RepositoryArgs{}
	}

	var resource Repository
	err := ctx.RegisterResource("gcp:sourcerepo/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("gcp:sourcerepo/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Resource name of the repository, of the form `{{repo}}`.
	// The repo name may contain slashes. eg, `name/with/slash`
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs []RepositoryPubsubConfig `pulumi:"pubsubConfigs"`
	// The disk usage of the repo, in bytes.
	Size *int `pulumi:"size"`
	// URL to clone the repository from Google Cloud Source Repositories.
	Url *string `pulumi:"url"`
}

type RepositoryState struct {
	// Resource name of the repository, of the form `{{repo}}`.
	// The repo name may contain slashes. eg, `name/with/slash`
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs RepositoryPubsubConfigArrayInput
	// The disk usage of the repo, in bytes.
	Size pulumi.IntPtrInput
	// URL to clone the repository from Google Cloud Source Repositories.
	Url pulumi.StringPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Resource name of the repository, of the form `{{repo}}`.
	// The repo name may contain slashes. eg, `name/with/slash`
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs []RepositoryPubsubConfig `pulumi:"pubsubConfigs"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Resource name of the repository, of the form `{{repo}}`.
	// The repo name may contain slashes. eg, `name/with/slash`
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// How this repository publishes a change in the repository through Cloud Pub/Sub.
	// Keyed by the topic names.
	// Structure is documented below.
	PubsubConfigs RepositoryPubsubConfigArrayInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//	RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//	RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// Resource name of the repository, of the form `{{repo}}`.
// The repo name may contain slashes. eg, `name/with/slash`
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RepositoryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// How this repository publishes a change in the repository through Cloud Pub/Sub.
// Keyed by the topic names.
// Structure is documented below.
func (o RepositoryOutput) PubsubConfigs() RepositoryPubsubConfigArrayOutput {
	return o.ApplyT(func(v *Repository) RepositoryPubsubConfigArrayOutput { return v.PubsubConfigs }).(RepositoryPubsubConfigArrayOutput)
}

// The disk usage of the repo, in bytes.
func (o RepositoryOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// URL to clone the repository from Google Cloud Source Repositories.
func (o RepositoryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].([]*Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].(map[string]*Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}
