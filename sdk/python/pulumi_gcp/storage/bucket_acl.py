# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BucketACLArgs', 'BucketACL']

@pulumi.input_type
class BucketACLArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 default_acl: Optional[pulumi.Input[str]] = None,
                 predefined_acl: Optional[pulumi.Input[str]] = None,
                 role_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a BucketACL resource.
        :param pulumi.Input[str] bucket: The name of the bucket it applies to.
        :param pulumi.Input[str] default_acl: Configure this ACL to be the default ACL.
        :param pulumi.Input[str] predefined_acl: The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        pulumi.set(__self__, "bucket", bucket)
        if default_acl is not None:
            pulumi.set(__self__, "default_acl", default_acl)
        if predefined_acl is not None:
            pulumi.set(__self__, "predefined_acl", predefined_acl)
        if role_entities is not None:
            pulumi.set(__self__, "role_entities", role_entities)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of the bucket it applies to.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="defaultAcl")
    def default_acl(self) -> Optional[pulumi.Input[str]]:
        """
        Configure this ACL to be the default ACL.
        """
        return pulumi.get(self, "default_acl")

    @default_acl.setter
    def default_acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_acl", value)

    @property
    @pulumi.getter(name="predefinedAcl")
    def predefined_acl(self) -> Optional[pulumi.Input[str]]:
        """
        The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        """
        return pulumi.get(self, "predefined_acl")

    @predefined_acl.setter
    def predefined_acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predefined_acl", value)

    @property
    @pulumi.getter(name="roleEntities")
    def role_entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        return pulumi.get(self, "role_entities")

    @role_entities.setter
    def role_entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_entities", value)


@pulumi.input_type
class _BucketACLState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 default_acl: Optional[pulumi.Input[str]] = None,
                 predefined_acl: Optional[pulumi.Input[str]] = None,
                 role_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering BucketACL resources.
        :param pulumi.Input[str] bucket: The name of the bucket it applies to.
        :param pulumi.Input[str] default_acl: Configure this ACL to be the default ACL.
        :param pulumi.Input[str] predefined_acl: The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if default_acl is not None:
            pulumi.set(__self__, "default_acl", default_acl)
        if predefined_acl is not None:
            pulumi.set(__self__, "predefined_acl", predefined_acl)
        if role_entities is not None:
            pulumi.set(__self__, "role_entities", role_entities)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket it applies to.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="defaultAcl")
    def default_acl(self) -> Optional[pulumi.Input[str]]:
        """
        Configure this ACL to be the default ACL.
        """
        return pulumi.get(self, "default_acl")

    @default_acl.setter
    def default_acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_acl", value)

    @property
    @pulumi.getter(name="predefinedAcl")
    def predefined_acl(self) -> Optional[pulumi.Input[str]]:
        """
        The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        """
        return pulumi.get(self, "predefined_acl")

    @predefined_acl.setter
    def predefined_acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "predefined_acl", value)

    @property
    @pulumi.getter(name="roleEntities")
    def role_entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        return pulumi.get(self, "role_entities")

    @role_entities.setter
    def role_entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_entities", value)


class BucketACL(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 default_acl: Optional[pulumi.Input[str]] = None,
                 predefined_acl: Optional[pulumi.Input[str]] = None,
                 role_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        Example creating an ACL on a bucket with one owner, and one reader.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_store = gcp.storage.Bucket("image-store", location="EU")
        image_store_acl = gcp.storage.BucketACL("image-store-acl",
            bucket=image_store.name,
            role_entities=[
                "OWNER:user-my.email@gmail.com",
                "READER:group-mygroup",
            ])
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket it applies to.
        :param pulumi.Input[str] default_acl: Configure this ACL to be the default ACL.
        :param pulumi.Input[str] predefined_acl: The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketACLArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        Example creating an ACL on a bucket with one owner, and one reader.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_store = gcp.storage.Bucket("image-store", location="EU")
        image_store_acl = gcp.storage.BucketACL("image-store-acl",
            bucket=image_store.name,
            role_entities=[
                "OWNER:user-my.email@gmail.com",
                "READER:group-mygroup",
            ])
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param BucketACLArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketACLArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 default_acl: Optional[pulumi.Input[str]] = None,
                 predefined_acl: Optional[pulumi.Input[str]] = None,
                 role_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketACLArgs.__new__(BucketACLArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["default_acl"] = default_acl
            __props__.__dict__["predefined_acl"] = predefined_acl
            __props__.__dict__["role_entities"] = role_entities
        super(BucketACL, __self__).__init__(
            'gcp:storage/bucketACL:BucketACL',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            default_acl: Optional[pulumi.Input[str]] = None,
            predefined_acl: Optional[pulumi.Input[str]] = None,
            role_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'BucketACL':
        """
        Get an existing BucketACL resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket it applies to.
        :param pulumi.Input[str] default_acl: Configure this ACL to be the default ACL.
        :param pulumi.Input[str] predefined_acl: The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_entities: List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketACLState.__new__(_BucketACLState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["default_acl"] = default_acl
        __props__.__dict__["predefined_acl"] = predefined_acl
        __props__.__dict__["role_entities"] = role_entities
        return BucketACL(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket it applies to.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="defaultAcl")
    def default_acl(self) -> pulumi.Output[Optional[str]]:
        """
        Configure this ACL to be the default ACL.
        """
        return pulumi.get(self, "default_acl")

    @property
    @pulumi.getter(name="predefinedAcl")
    def predefined_acl(self) -> pulumi.Output[Optional[str]]:
        """
        The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
        """
        return pulumi.get(self, "predefined_acl")

    @property
    @pulumi.getter(name="roleEntities")
    def role_entities(self) -> pulumi.Output[Sequence[str]]:
        """
        List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
        """
        return pulumi.get(self, "role_entities")

