// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datastream

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Datastream Connection Profile Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datastream"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datastream.NewConnectionProfile(ctx, "default", &datastream.ConnectionProfileArgs{
//				ConnectionProfileId: pulumi.String("my-profile"),
//				DisplayName:         pulumi.String("Connection profile"),
//				GcsProfile: &datastream.ConnectionProfileGcsProfileArgs{
//					Bucket:   pulumi.String("my-bucket"),
//					RootPath: pulumi.String("/path"),
//				},
//				Location: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Datastream Connection Profile Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datastream"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := datastream.NewConnectionProfile(ctx, "default", &datastream.ConnectionProfileArgs{
//				ConnectionProfileId: pulumi.String("my-profile"),
//				DisplayName:         pulumi.String("Connection profile"),
//				ForwardSshConnectivity: &datastream.ConnectionProfileForwardSshConnectivityArgs{
//					Hostname: pulumi.String("google.com"),
//					Password: pulumi.String("swordfish"),
//					Port:     pulumi.Int(8022),
//					Username: pulumi.String("my-user"),
//				},
//				GcsProfile: &datastream.ConnectionProfileGcsProfileArgs{
//					Bucket:   pulumi.String("my-bucket"),
//					RootPath: pulumi.String("/path"),
//				},
//				Location: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Datastream Connection Profile Postgres
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/datastream"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			instance, err := sql.NewDatabaseInstance(ctx, "instance", &sql.DatabaseInstanceArgs{
//				DatabaseVersion: pulumi.String("POSTGRES_14"),
//				Region:          pulumi.String("us-central1"),
//				Settings: &sql.DatabaseInstanceSettingsArgs{
//					Tier: pulumi.String("db-f1-micro"),
//					IpConfiguration: &sql.DatabaseInstanceSettingsIpConfigurationArgs{
//						AuthorizedNetworks: sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetworkArray{
//							&sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetworkArgs{
//								Value: pulumi.String("34.71.242.81"),
//							},
//							&sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetworkArgs{
//								Value: pulumi.String("34.72.28.29"),
//							},
//							&sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetworkArgs{
//								Value: pulumi.String("34.67.6.157"),
//							},
//							&sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetworkArgs{
//								Value: pulumi.String("34.67.234.134"),
//							},
//							&sql.DatabaseInstanceSettingsIpConfigurationAuthorizedNetworkArgs{
//								Value: pulumi.String("34.72.239.218"),
//							},
//						},
//					},
//				},
//				DeletionProtection: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			db, err := sql.NewDatabase(ctx, "db", &sql.DatabaseArgs{
//				Instance: instance.Name,
//			})
//			if err != nil {
//				return err
//			}
//			pwd, err := random.NewRandomPassword(ctx, "pwd", &random.RandomPasswordArgs{
//				Length:  pulumi.Int(16),
//				Special: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			user, err := sql.NewUser(ctx, "user", &sql.UserArgs{
//				Instance: instance.Name,
//				Password: pwd.Result,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = datastream.NewConnectionProfile(ctx, "default", &datastream.ConnectionProfileArgs{
//				DisplayName:         pulumi.String("Connection profile"),
//				Location:            pulumi.String("us-central1"),
//				ConnectionProfileId: pulumi.String("my-profile"),
//				PostgresqlProfile: &datastream.ConnectionProfilePostgresqlProfileArgs{
//					Hostname: instance.PublicIpAddress,
//					Username: user.Name,
//					Password: user.Password,
//					Database: db.Name,
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
// # ConnectionProfile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:datastream/connectionProfile:ConnectionProfile default projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:datastream/connectionProfile:ConnectionProfile default {{project}}/{{location}}/{{connection_profile_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:datastream/connectionProfile:ConnectionProfile default {{location}}/{{connection_profile_id}}
//
// ```
type ConnectionProfile struct {
	pulumi.CustomResourceState

	// The connection profile identifier.
	ConnectionProfileId pulumi.StringOutput `pulumi:"connectionProfileId"`
	// Display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSshConnectivity ConnectionProfileForwardSshConnectivityPtrOutput `pulumi:"forwardSshConnectivity"`
	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile ConnectionProfileGcsProfilePtrOutput `pulumi:"gcsProfile"`
	// Labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location pulumi.StringOutput `pulumi:"location"`
	// MySQL database profile.
	// Structure is documented below.
	MysqlProfile ConnectionProfileMysqlProfilePtrOutput `pulumi:"mysqlProfile"`
	// The resource's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Oracle database profile.
	// Structure is documented below.
	OracleProfile ConnectionProfileOracleProfilePtrOutput `pulumi:"oracleProfile"`
	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile ConnectionProfilePostgresqlProfilePtrOutput `pulumi:"postgresqlProfile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConnectionProfile registers a new resource with the given unique name, arguments, and options.
func NewConnectionProfile(ctx *pulumi.Context,
	name string, args *ConnectionProfileArgs, opts ...pulumi.ResourceOption) (*ConnectionProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProfileId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource ConnectionProfile
	err := ctx.RegisterResource("gcp:datastream/connectionProfile:ConnectionProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionProfile gets an existing ConnectionProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionProfileState, opts ...pulumi.ResourceOption) (*ConnectionProfile, error) {
	var resource ConnectionProfile
	err := ctx.ReadResource("gcp:datastream/connectionProfile:ConnectionProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionProfile resources.
type connectionProfileState struct {
	// The connection profile identifier.
	ConnectionProfileId *string `pulumi:"connectionProfileId"`
	// Display name.
	DisplayName *string `pulumi:"displayName"`
	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSshConnectivity *ConnectionProfileForwardSshConnectivity `pulumi:"forwardSshConnectivity"`
	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile *ConnectionProfileGcsProfile `pulumi:"gcsProfile"`
	// Labels.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location *string `pulumi:"location"`
	// MySQL database profile.
	// Structure is documented below.
	MysqlProfile *ConnectionProfileMysqlProfile `pulumi:"mysqlProfile"`
	// The resource's name.
	Name *string `pulumi:"name"`
	// Oracle database profile.
	// Structure is documented below.
	OracleProfile *ConnectionProfileOracleProfile `pulumi:"oracleProfile"`
	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile *ConnectionProfilePostgresqlProfile `pulumi:"postgresqlProfile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConnectionProfileState struct {
	// The connection profile identifier.
	ConnectionProfileId pulumi.StringPtrInput
	// Display name.
	DisplayName pulumi.StringPtrInput
	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSshConnectivity ConnectionProfileForwardSshConnectivityPtrInput
	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile ConnectionProfileGcsProfilePtrInput
	// Labels.
	Labels pulumi.StringMapInput
	// The name of the location this repository is located in.
	Location pulumi.StringPtrInput
	// MySQL database profile.
	// Structure is documented below.
	MysqlProfile ConnectionProfileMysqlProfilePtrInput
	// The resource's name.
	Name pulumi.StringPtrInput
	// Oracle database profile.
	// Structure is documented below.
	OracleProfile ConnectionProfileOracleProfilePtrInput
	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile ConnectionProfilePostgresqlProfilePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConnectionProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionProfileState)(nil)).Elem()
}

type connectionProfileArgs struct {
	// The connection profile identifier.
	ConnectionProfileId string `pulumi:"connectionProfileId"`
	// Display name.
	DisplayName string `pulumi:"displayName"`
	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSshConnectivity *ConnectionProfileForwardSshConnectivity `pulumi:"forwardSshConnectivity"`
	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile *ConnectionProfileGcsProfile `pulumi:"gcsProfile"`
	// Labels.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location string `pulumi:"location"`
	// MySQL database profile.
	// Structure is documented below.
	MysqlProfile *ConnectionProfileMysqlProfile `pulumi:"mysqlProfile"`
	// Oracle database profile.
	// Structure is documented below.
	OracleProfile *ConnectionProfileOracleProfile `pulumi:"oracleProfile"`
	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile *ConnectionProfilePostgresqlProfile `pulumi:"postgresqlProfile"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ConnectionProfile resource.
type ConnectionProfileArgs struct {
	// The connection profile identifier.
	ConnectionProfileId pulumi.StringInput
	// Display name.
	DisplayName pulumi.StringInput
	// Forward SSH tunnel connectivity.
	// Structure is documented below.
	ForwardSshConnectivity ConnectionProfileForwardSshConnectivityPtrInput
	// Cloud Storage bucket profile.
	// Structure is documented below.
	GcsProfile ConnectionProfileGcsProfilePtrInput
	// Labels.
	Labels pulumi.StringMapInput
	// The name of the location this repository is located in.
	Location pulumi.StringInput
	// MySQL database profile.
	// Structure is documented below.
	MysqlProfile ConnectionProfileMysqlProfilePtrInput
	// Oracle database profile.
	// Structure is documented below.
	OracleProfile ConnectionProfileOracleProfilePtrInput
	// PostgreSQL database profile.
	// Structure is documented below.
	PostgresqlProfile ConnectionProfilePostgresqlProfilePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConnectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionProfileArgs)(nil)).Elem()
}

type ConnectionProfileInput interface {
	pulumi.Input

	ToConnectionProfileOutput() ConnectionProfileOutput
	ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput
}

func (*ConnectionProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProfile)(nil)).Elem()
}

func (i *ConnectionProfile) ToConnectionProfileOutput() ConnectionProfileOutput {
	return i.ToConnectionProfileOutputWithContext(context.Background())
}

func (i *ConnectionProfile) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileOutput)
}

// ConnectionProfileArrayInput is an input type that accepts ConnectionProfileArray and ConnectionProfileArrayOutput values.
// You can construct a concrete instance of `ConnectionProfileArrayInput` via:
//
//	ConnectionProfileArray{ ConnectionProfileArgs{...} }
type ConnectionProfileArrayInput interface {
	pulumi.Input

	ToConnectionProfileArrayOutput() ConnectionProfileArrayOutput
	ToConnectionProfileArrayOutputWithContext(context.Context) ConnectionProfileArrayOutput
}

type ConnectionProfileArray []ConnectionProfileInput

func (ConnectionProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectionProfile)(nil)).Elem()
}

func (i ConnectionProfileArray) ToConnectionProfileArrayOutput() ConnectionProfileArrayOutput {
	return i.ToConnectionProfileArrayOutputWithContext(context.Background())
}

func (i ConnectionProfileArray) ToConnectionProfileArrayOutputWithContext(ctx context.Context) ConnectionProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileArrayOutput)
}

// ConnectionProfileMapInput is an input type that accepts ConnectionProfileMap and ConnectionProfileMapOutput values.
// You can construct a concrete instance of `ConnectionProfileMapInput` via:
//
//	ConnectionProfileMap{ "key": ConnectionProfileArgs{...} }
type ConnectionProfileMapInput interface {
	pulumi.Input

	ToConnectionProfileMapOutput() ConnectionProfileMapOutput
	ToConnectionProfileMapOutputWithContext(context.Context) ConnectionProfileMapOutput
}

type ConnectionProfileMap map[string]ConnectionProfileInput

func (ConnectionProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectionProfile)(nil)).Elem()
}

func (i ConnectionProfileMap) ToConnectionProfileMapOutput() ConnectionProfileMapOutput {
	return i.ToConnectionProfileMapOutputWithContext(context.Background())
}

func (i ConnectionProfileMap) ToConnectionProfileMapOutputWithContext(ctx context.Context) ConnectionProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileMapOutput)
}

type ConnectionProfileOutput struct{ *pulumi.OutputState }

func (ConnectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionProfile)(nil)).Elem()
}

func (o ConnectionProfileOutput) ToConnectionProfileOutput() ConnectionProfileOutput {
	return o
}

func (o ConnectionProfileOutput) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return o
}

// The connection profile identifier.
func (o ConnectionProfileOutput) ConnectionProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionProfile) pulumi.StringOutput { return v.ConnectionProfileId }).(pulumi.StringOutput)
}

// Display name.
func (o ConnectionProfileOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionProfile) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Forward SSH tunnel connectivity.
// Structure is documented below.
func (o ConnectionProfileOutput) ForwardSshConnectivity() ConnectionProfileForwardSshConnectivityPtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) ConnectionProfileForwardSshConnectivityPtrOutput {
		return v.ForwardSshConnectivity
	}).(ConnectionProfileForwardSshConnectivityPtrOutput)
}

// Cloud Storage bucket profile.
// Structure is documented below.
func (o ConnectionProfileOutput) GcsProfile() ConnectionProfileGcsProfilePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) ConnectionProfileGcsProfilePtrOutput { return v.GcsProfile }).(ConnectionProfileGcsProfilePtrOutput)
}

// Labels.
func (o ConnectionProfileOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionProfile) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the location this repository is located in.
func (o ConnectionProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// MySQL database profile.
// Structure is documented below.
func (o ConnectionProfileOutput) MysqlProfile() ConnectionProfileMysqlProfilePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) ConnectionProfileMysqlProfilePtrOutput { return v.MysqlProfile }).(ConnectionProfileMysqlProfilePtrOutput)
}

// The resource's name.
func (o ConnectionProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Oracle database profile.
// Structure is documented below.
func (o ConnectionProfileOutput) OracleProfile() ConnectionProfileOracleProfilePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) ConnectionProfileOracleProfilePtrOutput { return v.OracleProfile }).(ConnectionProfileOracleProfilePtrOutput)
}

// PostgreSQL database profile.
// Structure is documented below.
func (o ConnectionProfileOutput) PostgresqlProfile() ConnectionProfilePostgresqlProfilePtrOutput {
	return o.ApplyT(func(v *ConnectionProfile) ConnectionProfilePostgresqlProfilePtrOutput { return v.PostgresqlProfile }).(ConnectionProfilePostgresqlProfilePtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o ConnectionProfileOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionProfile) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

type ConnectionProfileArrayOutput struct{ *pulumi.OutputState }

func (ConnectionProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectionProfile)(nil)).Elem()
}

func (o ConnectionProfileArrayOutput) ToConnectionProfileArrayOutput() ConnectionProfileArrayOutput {
	return o
}

func (o ConnectionProfileArrayOutput) ToConnectionProfileArrayOutputWithContext(ctx context.Context) ConnectionProfileArrayOutput {
	return o
}

func (o ConnectionProfileArrayOutput) Index(i pulumi.IntInput) ConnectionProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConnectionProfile {
		return vs[0].([]*ConnectionProfile)[vs[1].(int)]
	}).(ConnectionProfileOutput)
}

type ConnectionProfileMapOutput struct{ *pulumi.OutputState }

func (ConnectionProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectionProfile)(nil)).Elem()
}

func (o ConnectionProfileMapOutput) ToConnectionProfileMapOutput() ConnectionProfileMapOutput {
	return o
}

func (o ConnectionProfileMapOutput) ToConnectionProfileMapOutputWithContext(ctx context.Context) ConnectionProfileMapOutput {
	return o
}

func (o ConnectionProfileMapOutput) MapIndex(k pulumi.StringInput) ConnectionProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConnectionProfile {
		return vs[0].(map[string]*ConnectionProfile)[vs[1].(string)]
	}).(ConnectionProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionProfileInput)(nil)).Elem(), &ConnectionProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionProfileArrayInput)(nil)).Elem(), ConnectionProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionProfileMapInput)(nil)).Elem(), ConnectionProfileMap{})
	pulumi.RegisterOutputType(ConnectionProfileOutput{})
	pulumi.RegisterOutputType(ConnectionProfileArrayOutput{})
	pulumi.RegisterOutputType(ConnectionProfileMapOutput{})
}
