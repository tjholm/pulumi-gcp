// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firestore

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// In Cloud Firestore, the unit of storage is the document. A document is a lightweight record
// that contains fields, which map to values. Each document is identified by a name.
//
// To get more information about Document, see:
//
// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/firestore/docs/manage-data/add-data)
//
// > **Warning:** This resource creates a Firestore Document on a project that already has
// Firestore enabled. If you haven't already enabled it, you can create a
// `appengine.Application` resource with `databaseType` set to
// `"CLOUD_FIRESTORE"` to do so. Your Firestore location will be the same as
// the App Engine location specified.
//
// ## Example Usage
// ### Firestore Document Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firestore.NewDocument(ctx, "mydoc", &firestore.DocumentArgs{
//				Collection: pulumi.String("somenewcollection"),
//				DocumentId: pulumi.String(fmt.Sprintf("my-doc-%vrandom_suffix}", "%{")),
//				Fields:     pulumi.String("{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}"),
//				Project:    pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Firestore Document Nested Document
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/firestore"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mydoc, err := firestore.NewDocument(ctx, "mydoc", &firestore.DocumentArgs{
//				Collection: pulumi.String("somenewcollection"),
//				DocumentId: pulumi.String(fmt.Sprintf("my-doc-%vrandom_suffix}", "%{")),
//				Fields:     pulumi.String("{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}"),
//				Project:    pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			subDocument, err := firestore.NewDocument(ctx, "subDocument", &firestore.DocumentArgs{
//				Collection: mydoc.Path.ApplyT(func(path string) (string, error) {
//					return fmt.Sprintf("%v/subdocs", path), nil
//				}).(pulumi.StringOutput),
//				DocumentId: pulumi.String("bitcoinkey"),
//				Fields:     pulumi.String("{\"something\":{\"mapValue\":{\"fields\":{\"ayo\":{\"stringValue\":\"val2\"}}}}}"),
//				Project:    pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firestore.NewDocument(ctx, "subSubDocument", &firestore.DocumentArgs{
//				Collection: subDocument.Path.ApplyT(func(path string) (string, error) {
//					return fmt.Sprintf("%v/subsubdocs", path), nil
//				}).(pulumi.StringOutput),
//				DocumentId: pulumi.String("asecret"),
//				Fields:     pulumi.String("{\"something\":{\"mapValue\":{\"fields\":{\"secret\":{\"stringValue\":\"hithere\"}}}}}"),
//				Project:    pulumi.String("my-project-name"),
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
// # Document can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:firestore/document:Document default {{name}}
//
// ```
type Document struct {
	pulumi.CustomResourceState

	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection pulumi.StringOutput `pulumi:"collection"`
	// Creation timestamp in RFC3339 format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrOutput `pulumi:"database"`
	// The client-assigned document ID to use for this document during creation.
	DocumentId pulumi.StringOutput `pulumi:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	Fields pulumi.StringOutput `pulumi:"fields"`
	// A server defined name for this index. Format:
	// `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
	Name pulumi.StringOutput `pulumi:"name"`
	// A relative path to the collection this document exists within
	Path pulumi.StringOutput `pulumi:"path"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Last update timestamp in RFC3339 format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDocument registers a new resource with the given unique name, arguments, and options.
func NewDocument(ctx *pulumi.Context,
	name string, args *DocumentArgs, opts ...pulumi.ResourceOption) (*Document, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Collection == nil {
		return nil, errors.New("invalid value for required argument 'Collection'")
	}
	if args.DocumentId == nil {
		return nil, errors.New("invalid value for required argument 'DocumentId'")
	}
	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	var resource Document
	err := ctx.RegisterResource("gcp:firestore/document:Document", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocument gets an existing Document resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocument(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentState, opts ...pulumi.ResourceOption) (*Document, error) {
	var resource Document
	err := ctx.ReadResource("gcp:firestore/document:Document", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Document resources.
type documentState struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection *string `pulumi:"collection"`
	// Creation timestamp in RFC3339 format.
	CreateTime *string `pulumi:"createTime"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database *string `pulumi:"database"`
	// The client-assigned document ID to use for this document during creation.
	DocumentId *string `pulumi:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	Fields *string `pulumi:"fields"`
	// A server defined name for this index. Format:
	// `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
	Name *string `pulumi:"name"`
	// A relative path to the collection this document exists within
	Path *string `pulumi:"path"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Last update timestamp in RFC3339 format.
	UpdateTime *string `pulumi:"updateTime"`
}

type DocumentState struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection pulumi.StringPtrInput
	// Creation timestamp in RFC3339 format.
	CreateTime pulumi.StringPtrInput
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrInput
	// The client-assigned document ID to use for this document during creation.
	DocumentId pulumi.StringPtrInput
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	Fields pulumi.StringPtrInput
	// A server defined name for this index. Format:
	// `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
	Name pulumi.StringPtrInput
	// A relative path to the collection this document exists within
	Path pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Last update timestamp in RFC3339 format.
	UpdateTime pulumi.StringPtrInput
}

func (DocumentState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentState)(nil)).Elem()
}

type documentArgs struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection string `pulumi:"collection"`
	// The Firestore database id. Defaults to `"(default)"`.
	Database *string `pulumi:"database"`
	// The client-assigned document ID to use for this document during creation.
	DocumentId string `pulumi:"documentId"`
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	Fields string `pulumi:"fields"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Document resource.
type DocumentArgs struct {
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection pulumi.StringInput
	// The Firestore database id. Defaults to `"(default)"`.
	Database pulumi.StringPtrInput
	// The client-assigned document ID to use for this document during creation.
	DocumentId pulumi.StringInput
	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	Fields pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DocumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentArgs)(nil)).Elem()
}

type DocumentInput interface {
	pulumi.Input

	ToDocumentOutput() DocumentOutput
	ToDocumentOutputWithContext(ctx context.Context) DocumentOutput
}

func (*Document) ElementType() reflect.Type {
	return reflect.TypeOf((**Document)(nil)).Elem()
}

func (i *Document) ToDocumentOutput() DocumentOutput {
	return i.ToDocumentOutputWithContext(context.Background())
}

func (i *Document) ToDocumentOutputWithContext(ctx context.Context) DocumentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentOutput)
}

// DocumentArrayInput is an input type that accepts DocumentArray and DocumentArrayOutput values.
// You can construct a concrete instance of `DocumentArrayInput` via:
//
//	DocumentArray{ DocumentArgs{...} }
type DocumentArrayInput interface {
	pulumi.Input

	ToDocumentArrayOutput() DocumentArrayOutput
	ToDocumentArrayOutputWithContext(context.Context) DocumentArrayOutput
}

type DocumentArray []DocumentInput

func (DocumentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Document)(nil)).Elem()
}

func (i DocumentArray) ToDocumentArrayOutput() DocumentArrayOutput {
	return i.ToDocumentArrayOutputWithContext(context.Background())
}

func (i DocumentArray) ToDocumentArrayOutputWithContext(ctx context.Context) DocumentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentArrayOutput)
}

// DocumentMapInput is an input type that accepts DocumentMap and DocumentMapOutput values.
// You can construct a concrete instance of `DocumentMapInput` via:
//
//	DocumentMap{ "key": DocumentArgs{...} }
type DocumentMapInput interface {
	pulumi.Input

	ToDocumentMapOutput() DocumentMapOutput
	ToDocumentMapOutputWithContext(context.Context) DocumentMapOutput
}

type DocumentMap map[string]DocumentInput

func (DocumentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Document)(nil)).Elem()
}

func (i DocumentMap) ToDocumentMapOutput() DocumentMapOutput {
	return i.ToDocumentMapOutputWithContext(context.Background())
}

func (i DocumentMap) ToDocumentMapOutputWithContext(ctx context.Context) DocumentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentMapOutput)
}

type DocumentOutput struct{ *pulumi.OutputState }

func (DocumentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Document)(nil)).Elem()
}

func (o DocumentOutput) ToDocumentOutput() DocumentOutput {
	return o
}

func (o DocumentOutput) ToDocumentOutputWithContext(ctx context.Context) DocumentOutput {
	return o
}

// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
func (o DocumentOutput) Collection() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Collection }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 format.
func (o DocumentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The Firestore database id. Defaults to `"(default)"`.
func (o DocumentOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Document) pulumi.StringPtrOutput { return v.Database }).(pulumi.StringPtrOutput)
}

// The client-assigned document ID to use for this document during creation.
func (o DocumentOutput) DocumentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.DocumentId }).(pulumi.StringOutput)
}

// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
func (o DocumentOutput) Fields() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Fields }).(pulumi.StringOutput)
}

// A server defined name for this index. Format:
// `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
func (o DocumentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A relative path to the collection this document exists within
func (o DocumentOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o DocumentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Last update timestamp in RFC3339 format.
func (o DocumentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Document) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type DocumentArrayOutput struct{ *pulumi.OutputState }

func (DocumentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Document)(nil)).Elem()
}

func (o DocumentArrayOutput) ToDocumentArrayOutput() DocumentArrayOutput {
	return o
}

func (o DocumentArrayOutput) ToDocumentArrayOutputWithContext(ctx context.Context) DocumentArrayOutput {
	return o
}

func (o DocumentArrayOutput) Index(i pulumi.IntInput) DocumentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Document {
		return vs[0].([]*Document)[vs[1].(int)]
	}).(DocumentOutput)
}

type DocumentMapOutput struct{ *pulumi.OutputState }

func (DocumentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Document)(nil)).Elem()
}

func (o DocumentMapOutput) ToDocumentMapOutput() DocumentMapOutput {
	return o
}

func (o DocumentMapOutput) ToDocumentMapOutputWithContext(ctx context.Context) DocumentMapOutput {
	return o
}

func (o DocumentMapOutput) MapIndex(k pulumi.StringInput) DocumentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Document {
		return vs[0].(map[string]*Document)[vs[1].(string)]
	}).(DocumentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentInput)(nil)).Elem(), &Document{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentArrayInput)(nil)).Elem(), DocumentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentMapInput)(nil)).Elem(), DocumentMap{})
	pulumi.RegisterOutputType(DocumentOutput{})
	pulumi.RegisterOutputType(DocumentArrayOutput{})
	pulumi.RegisterOutputType(DocumentMapOutput{})
}
