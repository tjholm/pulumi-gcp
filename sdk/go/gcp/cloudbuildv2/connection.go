// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudbuildv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Beta only: The Cloudbuildv2 Connection resource
//
// ## Example Usage
// ### GitHub Connection
// Creates a Connection to github.com
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudbuildv2"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/secretmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := secretmanager.NewSecret(ctx, "github-token-secret", &secretmanager.SecretArgs{
//				SecretId: pulumi.String("github-token-secret"),
//				Replication: &secretmanager.SecretReplicationArgs{
//					Automatic: pulumi.Bool(true),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = secretmanager.NewSecretVersion(ctx, "github-token-secret-version", &secretmanager.SecretVersionArgs{
//				Secret:     github_token_secret.ID(),
//				SecretData: readFileOrPanic("my-github-token.txt"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			p4sa_secretAccessor, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/secretmanager.secretAccessor",
//						Members: []string{
//							"serviceAccount:service-123456789@gcp-sa-cloudbuild.iam.gserviceaccount.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = secretmanager.NewSecretIamPolicy(ctx, "policy", &secretmanager.SecretIamPolicyArgs{
//				SecretId:   github_token_secret.SecretId,
//				PolicyData: *pulumi.String(p4sa_secretAccessor.PolicyData),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = cloudbuildv2.NewConnection(ctx, "my-connection", &cloudbuildv2.ConnectionArgs{
//				Location: pulumi.String("us-west1"),
//				GithubConfig: &cloudbuildv2.ConnectionGithubConfigArgs{
//					AppInstallationId: pulumi.Int(123123),
//					AuthorizerCredential: &cloudbuildv2.ConnectionGithubConfigAuthorizerCredentialArgs{
//						OauthTokenSecretVersion: github_token_secret_version.ID(),
//					},
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
// # Connection can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:cloudbuildv2/connection:Connection default projects/{{project}}/locations/{{location}}/connections/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudbuildv2/connection:Connection default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:cloudbuildv2/connection:Connection default {{location}}/{{name}}
//
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Allows clients to store small amounts of arbitrary data.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Configuration for connections to github.com.
	GithubConfig ConnectionGithubConfigPtrOutput `pulumi:"githubConfig"`
	// Configuration for connections to an instance of GitHub Enterprise.
	GithubEnterpriseConfig ConnectionGithubEnterpriseConfigPtrOutput `pulumi:"githubEnterpriseConfig"`
	// Output only. Installation state of the Connection.
	InstallationStates ConnectionInstallationStateArrayOutput `pulumi:"installationStates"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. Set to true when the connection is being set up or updated in the background.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Connection
	err := ctx.RegisterResource("gcp:cloudbuildv2/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("gcp:cloudbuildv2/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Allows clients to store small amounts of arbitrary data.
	Annotations map[string]string `pulumi:"annotations"`
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime *string `pulumi:"createTime"`
	// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	Disabled *bool `pulumi:"disabled"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Configuration for connections to github.com.
	GithubConfig *ConnectionGithubConfig `pulumi:"githubConfig"`
	// Configuration for connections to an instance of GitHub Enterprise.
	GithubEnterpriseConfig *ConnectionGithubEnterpriseConfig `pulumi:"githubEnterpriseConfig"`
	// Output only. Installation state of the Connection.
	InstallationStates []ConnectionInstallationState `pulumi:"installationStates"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. Set to true when the connection is being set up or updated in the background.
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type ConnectionState struct {
	// Allows clients to store small amounts of arbitrary data.
	Annotations pulumi.StringMapInput
	// Output only. Server assigned timestamp for when the connection was created.
	CreateTime pulumi.StringPtrInput
	// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	Disabled pulumi.BoolPtrInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Configuration for connections to github.com.
	GithubConfig ConnectionGithubConfigPtrInput
	// Configuration for connections to an instance of GitHub Enterprise.
	GithubEnterpriseConfig ConnectionGithubEnterpriseConfigPtrInput
	// Output only. Installation state of the Connection.
	InstallationStates ConnectionInstallationStateArrayInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. Set to true when the connection is being set up or updated in the background.
	Reconciling pulumi.BoolPtrInput
	// Output only. Server assigned timestamp for when the connection was updated.
	UpdateTime pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Allows clients to store small amounts of arbitrary data.
	Annotations map[string]string `pulumi:"annotations"`
	// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	Disabled *bool `pulumi:"disabled"`
	// Configuration for connections to github.com.
	GithubConfig *ConnectionGithubConfig `pulumi:"githubConfig"`
	// Configuration for connections to an instance of GitHub Enterprise.
	GithubEnterpriseConfig *ConnectionGithubEnterpriseConfig `pulumi:"githubEnterpriseConfig"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Allows clients to store small amounts of arbitrary data.
	Annotations pulumi.StringMapInput
	// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	Disabled pulumi.BoolPtrInput
	// Configuration for connections to github.com.
	GithubConfig ConnectionGithubConfigPtrInput
	// Configuration for connections to an instance of GitHub Enterprise.
	GithubEnterpriseConfig ConnectionGithubEnterpriseConfigPtrInput
	// The location for the resource
	Location pulumi.StringInput
	// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//	ConnectionArray{ ConnectionArgs{...} }
type ConnectionArrayInput interface {
	pulumi.Input

	ToConnectionArrayOutput() ConnectionArrayOutput
	ToConnectionArrayOutputWithContext(context.Context) ConnectionArrayOutput
}

type ConnectionArray []ConnectionInput

func (ConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (i ConnectionArray) ToConnectionArrayOutput() ConnectionArrayOutput {
	return i.ToConnectionArrayOutputWithContext(context.Background())
}

func (i ConnectionArray) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionArrayOutput)
}

// ConnectionMapInput is an input type that accepts ConnectionMap and ConnectionMapOutput values.
// You can construct a concrete instance of `ConnectionMapInput` via:
//
//	ConnectionMap{ "key": ConnectionArgs{...} }
type ConnectionMapInput interface {
	pulumi.Input

	ToConnectionMapOutput() ConnectionMapOutput
	ToConnectionMapOutputWithContext(context.Context) ConnectionMapOutput
}

type ConnectionMap map[string]ConnectionInput

func (ConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (i ConnectionMap) ToConnectionMapOutput() ConnectionMapOutput {
	return i.ToConnectionMapOutputWithContext(context.Background())
}

func (i ConnectionMap) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMapOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// Allows clients to store small amounts of arbitrary data.
func (o ConnectionOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Output only. Server assigned timestamp for when the connection was created.
func (o ConnectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
func (o ConnectionOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o ConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Configuration for connections to github.com.
func (o ConnectionOutput) GithubConfig() ConnectionGithubConfigPtrOutput {
	return o.ApplyT(func(v *Connection) ConnectionGithubConfigPtrOutput { return v.GithubConfig }).(ConnectionGithubConfigPtrOutput)
}

// Configuration for connections to an instance of GitHub Enterprise.
func (o ConnectionOutput) GithubEnterpriseConfig() ConnectionGithubEnterpriseConfigPtrOutput {
	return o.ApplyT(func(v *Connection) ConnectionGithubEnterpriseConfigPtrOutput { return v.GithubEnterpriseConfig }).(ConnectionGithubEnterpriseConfigPtrOutput)
}

// Output only. Installation state of the Connection.
func (o ConnectionOutput) InstallationStates() ConnectionInstallationStateArrayOutput {
	return o.ApplyT(func(v *Connection) ConnectionInstallationStateArrayOutput { return v.InstallationStates }).(ConnectionInstallationStateArrayOutput)
}

// The location for the resource
func (o ConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o ConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Set to true when the connection is being set up or updated in the background.
func (o ConnectionOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Output only. Server assigned timestamp for when the connection was updated.
func (o ConnectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Connection)(nil)).Elem()
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].([]*Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Connection)(nil)).Elem()
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Connection {
		return vs[0].(map[string]*Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionArrayInput)(nil)).Elem(), ConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionMapInput)(nil)).Elem(), ConnectionMap{})
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
