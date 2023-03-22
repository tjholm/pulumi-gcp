// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A repository for storing artifacts
//
// To get more information about Repository, see:
//
// * [API documentation](https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/artifact-registry/docs/overview)
//
// ## Example Usage
// ### Artifact Registry Repository Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.NewRepository(ctx, "my-repo", &artifactregistry.RepositoryArgs{
//				Description:  pulumi.String("example docker repository"),
//				Format:       pulumi.String("DOCKER"),
//				Location:     pulumi.String("us-central1"),
//				RepositoryId: pulumi.String("my-repository"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Artifact Registry Repository Cmek
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			cryptoKey, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
//				CryptoKeyId: pulumi.String("kms-key"),
//				Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
//				Member:      pulumi.String(fmt.Sprintf("serviceAccount:service-%v@gcp-sa-artifactregistry.iam.gserviceaccount.com", project.Number)),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactregistry.NewRepository(ctx, "my-repo", &artifactregistry.RepositoryArgs{
//				Location:     pulumi.String("us-central1"),
//				RepositoryId: pulumi.String("my-repository"),
//				Description:  pulumi.String("example docker repository with cmek"),
//				Format:       pulumi.String("DOCKER"),
//				KmsKeyName:   pulumi.String("kms-key"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				cryptoKey,
//			}))
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
//	$ pulumi import gcp:artifactregistry/repository:Repository default projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:artifactregistry/repository:Repository default {{project}}/{{location}}/{{repository_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:artifactregistry/repository:Repository default {{location}}/{{repository_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:artifactregistry/repository:Repository default {{repository_id}}
//
// ```
type Repository struct {
	pulumi.CustomResourceState

	// The time when the repository was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The format of packages that are stored in the repository. Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	Format pulumi.StringOutput `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location pulumi.StringOutput `pulumi:"location"`
	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig RepositoryMavenConfigPtrOutput `pulumi:"mavenConfig"`
	// The name of the repository, for example:
	// "projects/p1/locations/us-central1/repositories/repo1"
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// The time when the repository was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	var resource Repository
	err := ctx.RegisterResource("gcp:artifactregistry/repository:Repository", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:artifactregistry/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// The time when the repository was created.
	CreateTime *string `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description *string `pulumi:"description"`
	// The format of packages that are stored in the repository. Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	Format *string `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location *string `pulumi:"location"`
	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig *RepositoryMavenConfig `pulumi:"mavenConfig"`
	// The name of the repository, for example:
	// "projects/p1/locations/us-central1/repositories/repo1"
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId *string `pulumi:"repositoryId"`
	// The time when the repository was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type RepositoryState struct {
	// The time when the repository was created.
	CreateTime pulumi.StringPtrInput
	// The user-provided description of the repository.
	Description pulumi.StringPtrInput
	// The format of packages that are stored in the repository. Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	Format pulumi.StringPtrInput
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrInput
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels pulumi.StringMapInput
	// The name of the location this repository is located in.
	Location pulumi.StringPtrInput
	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig RepositoryMavenConfigPtrInput
	// The name of the repository, for example:
	// "projects/p1/locations/us-central1/repositories/repo1"
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId pulumi.StringPtrInput
	// The time when the repository was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// The user-provided description of the repository.
	Description *string `pulumi:"description"`
	// The format of packages that are stored in the repository. Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	Format string `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location *string `pulumi:"location"`
	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig *RepositoryMavenConfig `pulumi:"mavenConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId string `pulumi:"repositoryId"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The user-provided description of the repository.
	Description pulumi.StringPtrInput
	// The format of packages that are stored in the repository. Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	Format pulumi.StringInput
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrInput
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels pulumi.StringMapInput
	// The name of the location this repository is located in.
	Location pulumi.StringPtrInput
	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig RepositoryMavenConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId pulumi.StringInput
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

// The time when the repository was created.
func (o RepositoryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The user-provided description of the repository.
func (o RepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The format of packages that are stored in the repository. Supported formats
// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
// You can only create alpha formats if you are a member of the
// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
func (o RepositoryOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// The Cloud KMS resource name of the customer managed encryption key that’s
// used to encrypt the contents of the Repository. Has the form:
// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
// This value may not be changed after the Repository has been created.
func (o RepositoryOutput) KmsKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.KmsKeyName }).(pulumi.StringPtrOutput)
}

// Labels with user-defined metadata.
// This field may contain up to 64 entries. Label keys and values may be no
// longer than 63 characters. Label keys must begin with a lowercase letter
// and may only contain lowercase letters, numeric characters, underscores,
// and dashes.
func (o RepositoryOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the location this repository is located in.
func (o RepositoryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// MavenRepositoryConfig is maven related repository details.
// Provides additional configuration details for repositories of the maven
// format type.
// Structure is documented below.
func (o RepositoryOutput) MavenConfig() RepositoryMavenConfigPtrOutput {
	return o.ApplyT(func(v *Repository) RepositoryMavenConfigPtrOutput { return v.MavenConfig }).(RepositoryMavenConfigPtrOutput)
}

// The name of the repository, for example:
// "projects/p1/locations/us-central1/repositories/repo1"
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RepositoryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The last part of the repository name, for example:
// "repo1"
func (o RepositoryOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// The time when the repository was last updated.
func (o RepositoryOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
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
