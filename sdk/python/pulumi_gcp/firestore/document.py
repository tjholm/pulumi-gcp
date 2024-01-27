# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DocumentArgs', 'Document']

@pulumi.input_type
class DocumentArgs:
    def __init__(__self__, *,
                 collection: pulumi.Input[str],
                 document_id: pulumi.Input[str],
                 fields: pulumi.Input[str],
                 database: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Document resource.
        :param pulumi.Input[str] collection: The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        :param pulumi.Input[str] document_id: The client-assigned document ID to use for this document during creation.
               
               
               - - -
        :param pulumi.Input[str] fields: The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "collection", collection)
        pulumi.set(__self__, "document_id", document_id)
        pulumi.set(__self__, "fields", fields)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def collection(self) -> pulumi.Input[str]:
        """
        The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        """
        return pulumi.get(self, "collection")

    @collection.setter
    def collection(self, value: pulumi.Input[str]):
        pulumi.set(self, "collection", value)

    @property
    @pulumi.getter(name="documentId")
    def document_id(self) -> pulumi.Input[str]:
        """
        The client-assigned document ID to use for this document during creation.


        - - -
        """
        return pulumi.get(self, "document_id")

    @document_id.setter
    def document_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "document_id", value)

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Input[str]:
        """
        The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: pulumi.Input[str]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _DocumentState:
    def __init__(__self__, *,
                 collection: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 document_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Document resources.
        :param pulumi.Input[str] collection: The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 format.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] document_id: The client-assigned document ID to use for this document during creation.
               
               
               - - -
        :param pulumi.Input[str] fields: The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        :param pulumi.Input[str] name: A server defined name for this index. Format:
               `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
        :param pulumi.Input[str] path: A relative path to the collection this document exists within
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: Last update timestamp in RFC3339 format.
        """
        if collection is not None:
            pulumi.set(__self__, "collection", collection)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if document_id is not None:
            pulumi.set(__self__, "document_id", document_id)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def collection(self) -> Optional[pulumi.Input[str]]:
        """
        The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        """
        return pulumi.get(self, "collection")

    @collection.setter
    def collection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collection", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 format.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="documentId")
    def document_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client-assigned document ID to use for this document during creation.


        - - -
        """
        return pulumi.get(self, "document_id")

    @document_id.setter
    def document_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document_id", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[str]]:
        """
        The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A server defined name for this index. Format:
        `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        A relative path to the collection this document exists within
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last update timestamp in RFC3339 format.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Document(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 document_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        In Cloud Firestore, the unit of storage is the document. A document is a lightweight record
        that contains fields, which map to values. Each document is identified by a name.

        To get more information about Document, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/manage-data/add-data)

        > **Warning:** This resource creates a Firestore Document on a project that already has
        a Firestore database. If you haven't already created it, you may
        create a `firestore.Database` resource with `type` set to
        `"FIRESTORE_NATIVE"` and `location_id` set to your chosen location.
        If you wish to use App Engine, you may instead create a
        `appengine.Application` resource with `database_type` set to
        `"CLOUD_FIRESTORE"`. Your Firestore location will be the same as
        the App Engine location specified.
        Note: The surface does not support configurable database id. Only `(default)`
        is allowed for the database parameter.

        ## Example Usage
        ### Firestore Document Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        mydoc = gcp.firestore.Document("mydoc",
            collection="somenewcollection",
            document_id="my-doc-id",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"akey\\":{\\"stringValue\\":\\"avalue\\"}}}}}",
            project="my-project-name")
        ```
        ### Firestore Document Nested Document

        ```python
        import pulumi
        import pulumi_gcp as gcp

        mydoc = gcp.firestore.Document("mydoc",
            collection="somenewcollection",
            document_id="my-doc-id",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"akey\\":{\\"stringValue\\":\\"avalue\\"}}}}}",
            project="my-project-name")
        sub_document = gcp.firestore.Document("subDocument",
            collection=mydoc.path.apply(lambda path: f"{path}/subdocs"),
            document_id="bitcoinkey",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"ayo\\":{\\"stringValue\\":\\"val2\\"}}}}}",
            project="my-project-name")
        sub_sub_document = gcp.firestore.Document("subSubDocument",
            collection=sub_document.path.apply(lambda path: f"{path}/subsubdocs"),
            document_id="asecret",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"secret\\":{\\"stringValue\\":\\"hithere\\"}}}}}",
            project="my-project-name")
        ```

        ## Import

        Document can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Document can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:firestore/document:Document default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] document_id: The client-assigned document ID to use for this document during creation.
               
               
               - - -
        :param pulumi.Input[str] fields: The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DocumentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        In Cloud Firestore, the unit of storage is the document. A document is a lightweight record
        that contains fields, which map to values. Each document is identified by a name.

        To get more information about Document, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/manage-data/add-data)

        > **Warning:** This resource creates a Firestore Document on a project that already has
        a Firestore database. If you haven't already created it, you may
        create a `firestore.Database` resource with `type` set to
        `"FIRESTORE_NATIVE"` and `location_id` set to your chosen location.
        If you wish to use App Engine, you may instead create a
        `appengine.Application` resource with `database_type` set to
        `"CLOUD_FIRESTORE"`. Your Firestore location will be the same as
        the App Engine location specified.
        Note: The surface does not support configurable database id. Only `(default)`
        is allowed for the database parameter.

        ## Example Usage
        ### Firestore Document Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        mydoc = gcp.firestore.Document("mydoc",
            collection="somenewcollection",
            document_id="my-doc-id",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"akey\\":{\\"stringValue\\":\\"avalue\\"}}}}}",
            project="my-project-name")
        ```
        ### Firestore Document Nested Document

        ```python
        import pulumi
        import pulumi_gcp as gcp

        mydoc = gcp.firestore.Document("mydoc",
            collection="somenewcollection",
            document_id="my-doc-id",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"akey\\":{\\"stringValue\\":\\"avalue\\"}}}}}",
            project="my-project-name")
        sub_document = gcp.firestore.Document("subDocument",
            collection=mydoc.path.apply(lambda path: f"{path}/subdocs"),
            document_id="bitcoinkey",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"ayo\\":{\\"stringValue\\":\\"val2\\"}}}}}",
            project="my-project-name")
        sub_sub_document = gcp.firestore.Document("subSubDocument",
            collection=sub_document.path.apply(lambda path: f"{path}/subsubdocs"),
            document_id="asecret",
            fields="{\\"something\\":{\\"mapValue\\":{\\"fields\\":{\\"secret\\":{\\"stringValue\\":\\"hithere\\"}}}}}",
            project="my-project-name")
        ```

        ## Import

        Document can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Document can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:firestore/document:Document default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param DocumentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DocumentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 document_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DocumentArgs.__new__(DocumentArgs)

            if collection is None and not opts.urn:
                raise TypeError("Missing required property 'collection'")
            __props__.__dict__["collection"] = collection
            __props__.__dict__["database"] = database
            if document_id is None and not opts.urn:
                raise TypeError("Missing required property 'document_id'")
            __props__.__dict__["document_id"] = document_id
            if fields is None and not opts.urn:
                raise TypeError("Missing required property 'fields'")
            __props__.__dict__["fields"] = fields
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["path"] = None
            __props__.__dict__["update_time"] = None
        super(Document, __self__).__init__(
            'gcp:firestore/document:Document',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            collection: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            document_id: Optional[pulumi.Input[str]] = None,
            fields: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Document':
        """
        Get an existing Document resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collection: The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 format.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] document_id: The client-assigned document ID to use for this document during creation.
               
               
               - - -
        :param pulumi.Input[str] fields: The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        :param pulumi.Input[str] name: A server defined name for this index. Format:
               `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
        :param pulumi.Input[str] path: A relative path to the collection this document exists within
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: Last update timestamp in RFC3339 format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DocumentState.__new__(_DocumentState)

        __props__.__dict__["collection"] = collection
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["database"] = database
        __props__.__dict__["document_id"] = document_id
        __props__.__dict__["fields"] = fields
        __props__.__dict__["name"] = name
        __props__.__dict__["path"] = path
        __props__.__dict__["project"] = project
        __props__.__dict__["update_time"] = update_time
        return Document(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def collection(self) -> pulumi.Output[str]:
        """
        The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
        """
        return pulumi.get(self, "collection")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="documentId")
    def document_id(self) -> pulumi.Output[str]:
        """
        The client-assigned document ID to use for this document during creation.


        - - -
        """
        return pulumi.get(self, "document_id")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[str]:
        """
        The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A server defined name for this index. Format:
        `projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        A relative path to the collection this document exists within
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update timestamp in RFC3339 format.
        """
        return pulumi.get(self, "update_time")

