// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Firebase Database Instance Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firebase.NewDatabaseInstance(ctx, "basic", &firebase.DatabaseInstanceArgs{
//				Project:    pulumi.String("my-project-name"),
//				Region:     pulumi.String("us-central1"),
//				InstanceId: pulumi.String("active-db"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firebase Database Instance Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firebase.NewDatabaseInstance(ctx, "full", &firebase.DatabaseInstanceArgs{
//				Project:      pulumi.String("my-project-name"),
//				Region:       pulumi.String("europe-west1"),
//				InstanceId:   pulumi.String("disabled-db"),
//				Type:         pulumi.String("USER_DATABASE"),
//				DesiredState: pulumi.String("DISABLED"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firebase Database Instance Default Database
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firebase"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultProject, err := organizations.NewProject(ctx, "defaultProject", &organizations.ProjectArgs{
//				ProjectId: pulumi.String("rtdb-project"),
//				OrgId:     pulumi.String("123456789"),
//				Labels: pulumi.StringMap{
//					"firebase": pulumi.String("enabled"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = firebase.NewProject(ctx, "defaultFirebase/projectProject", &firebase.ProjectArgs{
//				Project: defaultProject.ProjectId,
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			firebaseDatabase, err := projects.NewService(ctx, "firebaseDatabase", &projects.ServiceArgs{
//				Project: defaultFirebase / projectProject.Project,
//				Service: pulumi.String("firebasedatabase.googleapis.com"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = firebase.NewDatabaseInstance(ctx, "defaultDatabaseInstance", &firebase.DatabaseInstanceArgs{
//				Project:    defaultFirebase / projectProject.Project,
//				Region:     pulumi.String("us-central1"),
//				InstanceId: pulumi.String("rtdb-project-default-rtdb"),
//				Type:       pulumi.String("DEFAULT_DATABASE"),
//			}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
//				firebaseDatabase,
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
// # Instance can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default projects/{{project}}/locations/{{region}}/instances/{{instance_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default {{project}}/{{region}}/{{instance_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default {{region}}/{{instance_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:firebase/databaseInstance:DatabaseInstance default {{instance_id}}
//
// ```
type DatabaseInstance struct {
	pulumi.CustomResourceState

	// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
	// or https://{instance-id}.{region}.firebasedatabase.app in other regions.
	DatabaseUrl pulumi.StringOutput `pulumi:"databaseUrl"`
	// The intended database state.
	DesiredState pulumi.StringPtrOutput `pulumi:"desiredState"`
	// The globally unique identifier of the Firebase Realtime Database instance.
	// Instance IDs cannot be reused after deletion.
	//
	// ***
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The fully-qualified resource name of the Firebase Realtime Database, in the
	// format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
	// PROJECT_NUMBER: The Firebase project's [`ProjectNumber`](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the Firebase Realtime database resides.
	// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region pulumi.StringOutput `pulumi:"region"`
	// The current database state. Set desiredState to :DISABLED to disable the database and :ACTIVE to reenable the database
	State pulumi.StringOutput `pulumi:"state"`
	// The database type.
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
	// Default value is `USER_DATABASE`.
	// Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseInstance
	err := ctx.RegisterResource("gcp:firebase/databaseInstance:DatabaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseInstance gets an existing DatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseInstanceState, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	var resource DatabaseInstance
	err := ctx.ReadResource("gcp:firebase/databaseInstance:DatabaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseInstance resources.
type databaseInstanceState struct {
	// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
	// or https://{instance-id}.{region}.firebasedatabase.app in other regions.
	DatabaseUrl *string `pulumi:"databaseUrl"`
	// The intended database state.
	DesiredState *string `pulumi:"desiredState"`
	// The globally unique identifier of the Firebase Realtime Database instance.
	// Instance IDs cannot be reused after deletion.
	//
	// ***
	InstanceId *string `pulumi:"instanceId"`
	// The fully-qualified resource name of the Firebase Realtime Database, in the
	// format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
	// PROJECT_NUMBER: The Firebase project's [`ProjectNumber`](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the Firebase Realtime database resides.
	// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region *string `pulumi:"region"`
	// The current database state. Set desiredState to :DISABLED to disable the database and :ACTIVE to reenable the database
	State *string `pulumi:"state"`
	// The database type.
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
	// Default value is `USER_DATABASE`.
	// Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
	Type *string `pulumi:"type"`
}

type DatabaseInstanceState struct {
	// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
	// or https://{instance-id}.{region}.firebasedatabase.app in other regions.
	DatabaseUrl pulumi.StringPtrInput
	// The intended database state.
	DesiredState pulumi.StringPtrInput
	// The globally unique identifier of the Firebase Realtime Database instance.
	// Instance IDs cannot be reused after deletion.
	//
	// ***
	InstanceId pulumi.StringPtrInput
	// The fully-qualified resource name of the Firebase Realtime Database, in the
	// format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
	// PROJECT_NUMBER: The Firebase project's [`ProjectNumber`](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the Firebase Realtime database resides.
	// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region pulumi.StringPtrInput
	// The current database state. Set desiredState to :DISABLED to disable the database and :ACTIVE to reenable the database
	State pulumi.StringPtrInput
	// The database type.
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
	// Default value is `USER_DATABASE`.
	// Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
	Type pulumi.StringPtrInput
}

func (DatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceState)(nil)).Elem()
}

type databaseInstanceArgs struct {
	// The intended database state.
	DesiredState *string `pulumi:"desiredState"`
	// The globally unique identifier of the Firebase Realtime Database instance.
	// Instance IDs cannot be reused after deletion.
	//
	// ***
	InstanceId string `pulumi:"instanceId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A reference to the region where the Firebase Realtime database resides.
	// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region string `pulumi:"region"`
	// The database type.
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
	// Default value is `USER_DATABASE`.
	// Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// The intended database state.
	DesiredState pulumi.StringPtrInput
	// The globally unique identifier of the Firebase Realtime Database instance.
	// Instance IDs cannot be reused after deletion.
	//
	// ***
	InstanceId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A reference to the region where the Firebase Realtime database resides.
	// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region pulumi.StringInput
	// The database type.
	// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan.
	// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
	// Default value is `USER_DATABASE`.
	// Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
	Type pulumi.StringPtrInput
}

func (DatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceArgs)(nil)).Elem()
}

type DatabaseInstanceInput interface {
	pulumi.Input

	ToDatabaseInstanceOutput() DatabaseInstanceOutput
	ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput
}

func (*DatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (i *DatabaseInstance) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return i.ToDatabaseInstanceOutputWithContext(context.Background())
}

func (i *DatabaseInstance) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceOutput)
}

func (i *DatabaseInstance) ToOutput(ctx context.Context) pulumix.Output[*DatabaseInstance] {
	return pulumix.Output[*DatabaseInstance]{
		OutputState: i.ToDatabaseInstanceOutputWithContext(ctx).OutputState,
	}
}

// DatabaseInstanceArrayInput is an input type that accepts DatabaseInstanceArray and DatabaseInstanceArrayOutput values.
// You can construct a concrete instance of `DatabaseInstanceArrayInput` via:
//
//	DatabaseInstanceArray{ DatabaseInstanceArgs{...} }
type DatabaseInstanceArrayInput interface {
	pulumi.Input

	ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput
	ToDatabaseInstanceArrayOutputWithContext(context.Context) DatabaseInstanceArrayOutput
}

type DatabaseInstanceArray []DatabaseInstanceInput

func (DatabaseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseInstance)(nil)).Elem()
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return i.ToDatabaseInstanceArrayOutputWithContext(context.Background())
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceArrayOutput)
}

func (i DatabaseInstanceArray) ToOutput(ctx context.Context) pulumix.Output[[]*DatabaseInstance] {
	return pulumix.Output[[]*DatabaseInstance]{
		OutputState: i.ToDatabaseInstanceArrayOutputWithContext(ctx).OutputState,
	}
}

// DatabaseInstanceMapInput is an input type that accepts DatabaseInstanceMap and DatabaseInstanceMapOutput values.
// You can construct a concrete instance of `DatabaseInstanceMapInput` via:
//
//	DatabaseInstanceMap{ "key": DatabaseInstanceArgs{...} }
type DatabaseInstanceMapInput interface {
	pulumi.Input

	ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput
	ToDatabaseInstanceMapOutputWithContext(context.Context) DatabaseInstanceMapOutput
}

type DatabaseInstanceMap map[string]DatabaseInstanceInput

func (DatabaseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseInstance)(nil)).Elem()
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return i.ToDatabaseInstanceMapOutputWithContext(context.Background())
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceMapOutput)
}

func (i DatabaseInstanceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DatabaseInstance] {
	return pulumix.Output[map[string]*DatabaseInstance]{
		OutputState: i.ToDatabaseInstanceMapOutputWithContext(ctx).OutputState,
	}
}

type DatabaseInstanceOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToOutput(ctx context.Context) pulumix.Output[*DatabaseInstance] {
	return pulumix.Output[*DatabaseInstance]{
		OutputState: o.OutputState,
	}
}

// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
// or https://{instance-id}.{region}.firebasedatabase.app in other regions.
func (o DatabaseInstanceOutput) DatabaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.DatabaseUrl }).(pulumi.StringOutput)
}

// The intended database state.
func (o DatabaseInstanceOutput) DesiredState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.DesiredState }).(pulumi.StringPtrOutput)
}

// The globally unique identifier of the Firebase Realtime Database instance.
// Instance IDs cannot be reused after deletion.
//
// ***
func (o DatabaseInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The fully-qualified resource name of the Firebase Realtime Database, in the
// format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
// PROJECT_NUMBER: The Firebase project's [`ProjectNumber`](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
func (o DatabaseInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DatabaseInstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A reference to the region where the Firebase Realtime database resides.
// Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
func (o DatabaseInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current database state. Set desiredState to :DISABLED to disable the database and :ACTIVE to reenable the database
func (o DatabaseInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The database type.
// Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
// Creating user Databases is only available for projects on the Blaze plan.
// Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
// Default value is `USER_DATABASE`.
// Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
func (o DatabaseInstanceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type DatabaseInstanceArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DatabaseInstance] {
	return pulumix.Output[[]*DatabaseInstance]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseInstanceArrayOutput) Index(i pulumi.IntInput) DatabaseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseInstance {
		return vs[0].([]*DatabaseInstance)[vs[1].(int)]
	}).(DatabaseInstanceOutput)
}

type DatabaseInstanceMapOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DatabaseInstance] {
	return pulumix.Output[map[string]*DatabaseInstance]{
		OutputState: o.OutputState,
	}
}

func (o DatabaseInstanceMapOutput) MapIndex(k pulumi.StringInput) DatabaseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseInstance {
		return vs[0].(map[string]*DatabaseInstance)[vs[1].(string)]
	}).(DatabaseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceInput)(nil)).Elem(), &DatabaseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceArrayInput)(nil)).Elem(), DatabaseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceMapInput)(nil)).Elem(), DatabaseInstanceMap{})
	pulumi.RegisterOutputType(DatabaseInstanceOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceMapOutput{})
}
