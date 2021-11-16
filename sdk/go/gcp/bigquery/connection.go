// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A connection allows BigQuery connections to external data sources..
//
// To get more information about Connection, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/bigqueryconnection/rest/v1beta1/projects.locations.connections/create)
// * How-to Guides
//     * [Cloud SQL federated queries](https://cloud.google.com/bigquery/docs/cloud-sql-federated-queries)
//
// > **Warning:** All arguments including `cloud_sql.credential.password` will be stored in the raw
// state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
// ### Bigquery Connection Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance, err := sql.NewDatabaseInstance(ctx, "instance", &sql.DatabaseInstanceArgs{
// 			DatabaseVersion: pulumi.String("POSTGRES_11"),
// 			Region:          pulumi.String("us-central1"),
// 			Settings: &sql.DatabaseInstanceSettingsArgs{
// 				Tier: pulumi.String("db-f1-micro"),
// 			},
// 			DeletionProtection: pulumi.Bool(true),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		db, err := sql.NewDatabase(ctx, "db", &sql.DatabaseArgs{
// 			Instance: instance.Name,
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		pwd, err := random.NewRandomPassword(ctx, "pwd", &random.RandomPasswordArgs{
// 			Length:  pulumi.Int(16),
// 			Special: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		user, err := sql.NewUser(ctx, "user", &sql.UserArgs{
// 			Instance: instance.Name,
// 			Password: pwd.Result,
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewConnection(ctx, "connection", &bigquery.ConnectionArgs{
// 			FriendlyName: pulumi.String("👋"),
// 			Description:  pulumi.String("a riveting description"),
// 			CloudSql: &bigquery.ConnectionCloudSqlArgs{
// 				InstanceId: instance.ConnectionName,
// 				Database:   db.Name,
// 				Type:       pulumi.String("POSTGRES"),
// 				Credential: &bigquery.ConnectionCloudSqlCredentialArgs{
// 					Username: user.Name,
// 					Password: user.Password,
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Bigquery Connection Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/sql"
// 	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance, err := sql.NewDatabaseInstance(ctx, "instance", &sql.DatabaseInstanceArgs{
// 			DatabaseVersion: pulumi.String("POSTGRES_11"),
// 			Region:          pulumi.String("us-central1"),
// 			Settings: &sql.DatabaseInstanceSettingsArgs{
// 				Tier: pulumi.String("db-f1-micro"),
// 			},
// 			DeletionProtection: pulumi.Bool(true),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		db, err := sql.NewDatabase(ctx, "db", &sql.DatabaseArgs{
// 			Instance: instance.Name,
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		pwd, err := random.NewRandomPassword(ctx, "pwd", &random.RandomPasswordArgs{
// 			Length:  pulumi.Int(16),
// 			Special: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		user, err := sql.NewUser(ctx, "user", &sql.UserArgs{
// 			Instance: instance.Name,
// 			Password: pwd.Result,
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewConnection(ctx, "connection", &bigquery.ConnectionArgs{
// 			ConnectionId: pulumi.String("my-connection"),
// 			Location:     pulumi.String("US"),
// 			FriendlyName: pulumi.String("👋"),
// 			Description:  pulumi.String("a riveting description"),
// 			CloudSql: &bigquery.ConnectionCloudSqlArgs{
// 				InstanceId: instance.ConnectionName,
// 				Database:   db.Name,
// 				Type:       pulumi.String("POSTGRES"),
// 				Credential: &bigquery.ConnectionCloudSqlCredentialArgs{
// 					Username: user.Name,
// 					Password: user.Password,
// 				},
// 			},
// 		}, pulumi.Provider(google_beta))
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
// Connection can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:bigquery/connection:Connection default projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/connection:Connection default {{project}}/{{location}}/{{connection_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/connection:Connection default {{location}}/{{connection_id}}
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Cloud SQL properties.
	// Structure is documented below.
	CloudSql ConnectionCloudSqlOutput `pulumi:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// A descriptive description for the connection
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A descriptive name for the connection
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// True if the connection has credential assigned.
	HasCredential pulumi.BoolOutput `pulumi:"hasCredential"`
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudSql == nil {
		return nil, errors.New("invalid value for required argument 'CloudSql'")
	}
	var resource Connection
	err := ctx.RegisterResource("gcp:bigquery/connection:Connection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:bigquery/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Cloud SQL properties.
	// Structure is documented below.
	CloudSql *ConnectionCloudSql `pulumi:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	ConnectionId *string `pulumi:"connectionId"`
	// A descriptive description for the connection
	Description *string `pulumi:"description"`
	// A descriptive name for the connection
	FriendlyName *string `pulumi:"friendlyName"`
	// True if the connection has credential assigned.
	HasCredential *bool `pulumi:"hasCredential"`
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location *string `pulumi:"location"`
	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConnectionState struct {
	// Cloud SQL properties.
	// Structure is documented below.
	CloudSql ConnectionCloudSqlPtrInput
	// Optional connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrInput
	// A descriptive description for the connection
	Description pulumi.StringPtrInput
	// A descriptive name for the connection
	FriendlyName pulumi.StringPtrInput
	// True if the connection has credential assigned.
	HasCredential pulumi.BoolPtrInput
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location pulumi.StringPtrInput
	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Cloud SQL properties.
	// Structure is documented below.
	CloudSql ConnectionCloudSql `pulumi:"cloudSql"`
	// Optional connection id that should be assigned to the created connection.
	ConnectionId *string `pulumi:"connectionId"`
	// A descriptive description for the connection
	Description *string `pulumi:"description"`
	// A descriptive name for the connection
	FriendlyName *string `pulumi:"friendlyName"`
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Cloud SQL properties.
	// Structure is documented below.
	CloudSql ConnectionCloudSqlInput
	// Optional connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrInput
	// A descriptive description for the connection
	Description pulumi.StringPtrInput
	// A descriptive name for the connection
	FriendlyName pulumi.StringPtrInput
	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
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
	return reflect.TypeOf((*Connection)(nil))
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

func (i *Connection) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

type ConnectionPtrInput interface {
	pulumi.Input

	ToConnectionPtrOutput() ConnectionPtrOutput
	ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput
}

type connectionPtrType ConnectionArgs

func (*connectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil))
}

func (i *connectionPtrType) ToConnectionPtrOutput() ConnectionPtrOutput {
	return i.ToConnectionPtrOutputWithContext(context.Background())
}

func (i *connectionPtrType) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPtrOutput)
}

// ConnectionArrayInput is an input type that accepts ConnectionArray and ConnectionArrayOutput values.
// You can construct a concrete instance of `ConnectionArrayInput` via:
//
//          ConnectionArray{ ConnectionArgs{...} }
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
//          ConnectionMap{ "key": ConnectionArgs{...} }
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
	return reflect.TypeOf((*Connection)(nil))
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o.ToConnectionPtrOutputWithContext(context.Background())
}

func (o ConnectionOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Connection) *Connection {
		return &v
	}).(ConnectionPtrOutput)
}

type ConnectionPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil))
}

func (o ConnectionPtrOutput) ToConnectionPtrOutput() ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) ToConnectionPtrOutputWithContext(ctx context.Context) ConnectionPtrOutput {
	return o
}

func (o ConnectionPtrOutput) Elem() ConnectionOutput {
	return o.ApplyT(func(v *Connection) Connection {
		if v != nil {
			return *v
		}
		var ret Connection
		return ret
	}).(ConnectionOutput)
}

type ConnectionArrayOutput struct{ *pulumi.OutputState }

func (ConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Connection)(nil))
}

func (o ConnectionArrayOutput) ToConnectionArrayOutput() ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) ToConnectionArrayOutputWithContext(ctx context.Context) ConnectionArrayOutput {
	return o
}

func (o ConnectionArrayOutput) Index(i pulumi.IntInput) ConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Connection {
		return vs[0].([]Connection)[vs[1].(int)]
	}).(ConnectionOutput)
}

type ConnectionMapOutput struct{ *pulumi.OutputState }

func (ConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Connection)(nil))
}

func (o ConnectionMapOutput) ToConnectionMapOutput() ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) ToConnectionMapOutputWithContext(ctx context.Context) ConnectionMapOutput {
	return o
}

func (o ConnectionMapOutput) MapIndex(k pulumi.StringInput) ConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Connection {
		return vs[0].(map[string]Connection)[vs[1].(string)]
	}).(ConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionOutput{})
	pulumi.RegisterOutputType(ConnectionPtrOutput{})
	pulumi.RegisterOutputType(ConnectionArrayOutput{})
	pulumi.RegisterOutputType(ConnectionMapOutput{})
}
