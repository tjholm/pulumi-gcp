# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 clusters: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]] clusters: A block of cluster configuration options. This can be specified at least once, and up to 4 times.
               See structure below.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow this provider to destroy the instance. Unless this field is set to false
               in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
               It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
               and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
               `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
               is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        :param pulumi.Input[str] name: The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        if clusters is not None:
            pulumi.set(__self__, "clusters", clusters)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if instance_type is not None:
            warnings.warn("""It is recommended to leave this field unspecified since the distinction between \"DEVELOPMENT\" and \"PRODUCTION\" instances is going away, and all instances will become \"PRODUCTION\" instances. This means that new and existing \"DEVELOPMENT\" instances will be converted to \"PRODUCTION\" instances. It is recommended for users to use \"PRODUCTION\" instances in any case, since a 1-node \"PRODUCTION\" instance is functionally identical to a \"DEVELOPMENT\" instance, but without the accompanying restrictions.""", DeprecationWarning)
            pulumi.log.warn("""instance_type is deprecated: It is recommended to leave this field unspecified since the distinction between \"DEVELOPMENT\" and \"PRODUCTION\" instances is going away, and all instances will become \"PRODUCTION\" instances. This means that new and existing \"DEVELOPMENT\" instances will be converted to \"PRODUCTION\" instances. It is recommended for users to use \"PRODUCTION\" instances in any case, since a 1-node \"PRODUCTION\" instance is functionally identical to a \"DEVELOPMENT\" instance, but without the accompanying restrictions.""")
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def clusters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]]]:
        """
        A block of cluster configuration options. This can be specified at least once, and up to 4 times.
        See structure below.
        """
        return pulumi.get(self, "clusters")

    @clusters.setter
    def clusters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]]]):
        pulumi.set(self, "clusters", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to allow this provider to destroy the instance. Unless this field is set to false
        in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
        and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
        `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
        is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 clusters: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]] clusters: A block of cluster configuration options. This can be specified at least once, and up to 4 times.
               See structure below.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow this provider to destroy the instance. Unless this field is set to false
               in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
               It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
               and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
               `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
               is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        :param pulumi.Input[str] name: The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        if clusters is not None:
            pulumi.set(__self__, "clusters", clusters)
        if deletion_protection is not None:
            pulumi.set(__self__, "deletion_protection", deletion_protection)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if instance_type is not None:
            warnings.warn("""It is recommended to leave this field unspecified since the distinction between \"DEVELOPMENT\" and \"PRODUCTION\" instances is going away, and all instances will become \"PRODUCTION\" instances. This means that new and existing \"DEVELOPMENT\" instances will be converted to \"PRODUCTION\" instances. It is recommended for users to use \"PRODUCTION\" instances in any case, since a 1-node \"PRODUCTION\" instance is functionally identical to a \"DEVELOPMENT\" instance, but without the accompanying restrictions.""", DeprecationWarning)
            pulumi.log.warn("""instance_type is deprecated: It is recommended to leave this field unspecified since the distinction between \"DEVELOPMENT\" and \"PRODUCTION\" instances is going away, and all instances will become \"PRODUCTION\" instances. This means that new and existing \"DEVELOPMENT\" instances will be converted to \"PRODUCTION\" instances. It is recommended for users to use \"PRODUCTION\" instances in any case, since a 1-node \"PRODUCTION\" instance is functionally identical to a \"DEVELOPMENT\" instance, but without the accompanying restrictions.""")
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def clusters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]]]:
        """
        A block of cluster configuration options. This can be specified at least once, and up to 4 times.
        See structure below.
        """
        return pulumi.get(self, "clusters")

    @clusters.setter
    def clusters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceClusterArgs']]]]):
        pulumi.set(self, "clusters", value)

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to allow this provider to destroy the instance. Unless this field is set to false
        in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        """
        return pulumi.get(self, "deletion_protection")

    @deletion_protection.setter
    def deletion_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_protection", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
        and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
        `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
        is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clusters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceClusterArgs']]]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## +---

        subcategory: "Cloud Bigtable"
        layout: "google"
        page_title: "Google: bigtable.Instance"
        sidebar_current: "docs-google-bigtable-instance"
        description: |-
          Creates a Google Bigtable instance.
        ---

        # bigtable.Instance

        Creates a Google Bigtable instance. For more information see:

        * [API documentation](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/bigtable/docs)

        ## Example Usage
        ### Simple Instance

        ```python
        import pulumi
        import pulumi_gcp as gcp

        production_instance = gcp.bigtable.Instance("production-instance",
            clusters=[gcp.bigtable.InstanceClusterArgs(
                cluster_id="tf-instance-cluster",
                num_nodes=1,
                storage_type="HDD",
            )],
            labels={
                "my-label": "prod-label",
            })
        ```
        ### Replicated Instance

        ```python
        import pulumi
        import pulumi_gcp as gcp

        production_instance = gcp.bigtable.Instance("production-instance",
            clusters=[
                gcp.bigtable.InstanceClusterArgs(
                    cluster_id="tf-instance-cluster1",
                    num_nodes=1,
                    storage_type="HDD",
                    zone="us-central1-c",
                ),
                gcp.bigtable.InstanceClusterArgs(
                    autoscaling_config=gcp.bigtable.InstanceClusterAutoscalingConfigArgs(
                        cpu_target=50,
                        max_nodes=3,
                        min_nodes=1,
                    ),
                    cluster_id="tf-instance-cluster2",
                    storage_type="HDD",
                    zone="us-central1-b",
                ),
            ],
            labels={
                "my-label": "prod-label",
            })
        ```

        ## Import

        Bigtable Instances can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:bigtable/instance:Instance default projects/{{project}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigtable/instance:Instance default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigtable/instance:Instance default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceClusterArgs']]]] clusters: A block of cluster configuration options. This can be specified at least once, and up to 4 times.
               See structure below.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow this provider to destroy the instance. Unless this field is set to false
               in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
               It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
               and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
               `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
               is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        :param pulumi.Input[str] name: The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InstanceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## +---

        subcategory: "Cloud Bigtable"
        layout: "google"
        page_title: "Google: bigtable.Instance"
        sidebar_current: "docs-google-bigtable-instance"
        description: |-
          Creates a Google Bigtable instance.
        ---

        # bigtable.Instance

        Creates a Google Bigtable instance. For more information see:

        * [API documentation](https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/bigtable/docs)

        ## Example Usage
        ### Simple Instance

        ```python
        import pulumi
        import pulumi_gcp as gcp

        production_instance = gcp.bigtable.Instance("production-instance",
            clusters=[gcp.bigtable.InstanceClusterArgs(
                cluster_id="tf-instance-cluster",
                num_nodes=1,
                storage_type="HDD",
            )],
            labels={
                "my-label": "prod-label",
            })
        ```
        ### Replicated Instance

        ```python
        import pulumi
        import pulumi_gcp as gcp

        production_instance = gcp.bigtable.Instance("production-instance",
            clusters=[
                gcp.bigtable.InstanceClusterArgs(
                    cluster_id="tf-instance-cluster1",
                    num_nodes=1,
                    storage_type="HDD",
                    zone="us-central1-c",
                ),
                gcp.bigtable.InstanceClusterArgs(
                    autoscaling_config=gcp.bigtable.InstanceClusterAutoscalingConfigArgs(
                        cpu_target=50,
                        max_nodes=3,
                        min_nodes=1,
                    ),
                    cluster_id="tf-instance-cluster2",
                    storage_type="HDD",
                    zone="us-central1-b",
                ),
            ],
            labels={
                "my-label": "prod-label",
            })
        ```

        ## Import

        Bigtable Instances can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:bigtable/instance:Instance default projects/{{project}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigtable/instance:Instance default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigtable/instance:Instance default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clusters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceClusterArgs']]]]] = None,
                 deletion_protection: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["clusters"] = clusters
            __props__.__dict__["deletion_protection"] = deletion_protection
            __props__.__dict__["display_name"] = display_name
            if instance_type is not None and not opts.urn:
                warnings.warn("""It is recommended to leave this field unspecified since the distinction between \"DEVELOPMENT\" and \"PRODUCTION\" instances is going away, and all instances will become \"PRODUCTION\" instances. This means that new and existing \"DEVELOPMENT\" instances will be converted to \"PRODUCTION\" instances. It is recommended for users to use \"PRODUCTION\" instances in any case, since a 1-node \"PRODUCTION\" instance is functionally identical to a \"DEVELOPMENT\" instance, but without the accompanying restrictions.""", DeprecationWarning)
                pulumi.log.warn("""instance_type is deprecated: It is recommended to leave this field unspecified since the distinction between \"DEVELOPMENT\" and \"PRODUCTION\" instances is going away, and all instances will become \"PRODUCTION\" instances. This means that new and existing \"DEVELOPMENT\" instances will be converted to \"PRODUCTION\" instances. It is recommended for users to use \"PRODUCTION\" instances in any case, since a 1-node \"PRODUCTION\" instance is functionally identical to a \"DEVELOPMENT\" instance, but without the accompanying restrictions.""")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
        super(Instance, __self__).__init__(
            'gcp:bigtable/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            clusters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceClusterArgs']]]]] = None,
            deletion_protection: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceClusterArgs']]]] clusters: A block of cluster configuration options. This can be specified at least once, and up to 4 times.
               See structure below.
        :param pulumi.Input[bool] deletion_protection: Whether or not to allow this provider to destroy the instance. Unless this field is set to false
               in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
               It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
               and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
               `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
               is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        :param pulumi.Input[str] name: The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["clusters"] = clusters
        __props__.__dict__["deletion_protection"] = deletion_protection
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def clusters(self) -> pulumi.Output[Sequence['outputs.InstanceCluster']]:
        """
        A block of cluster configuration options. This can be specified at least once, and up to 4 times.
        See structure below.
        """
        return pulumi.get(self, "clusters")

    @property
    @pulumi.getter(name="deletionProtection")
    def deletion_protection(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to allow this provider to destroy the instance. Unless this field is set to false
        in the statefile, a `pulumi destroy` or `pulumi up` that would delete the instance will fail.
        """
        return pulumi.get(self, "deletion_protection")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[Optional[str]]:
        """
        The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        It is recommended to leave this field unspecified since the distinction between `"DEVELOPMENT"` and `"PRODUCTION"` instances is going away,
        and all instances will become `"PRODUCTION"` instances. This means that new and existing `"DEVELOPMENT"` instances will be converted to
        `"PRODUCTION"` instances. It is recommended for users to use `"PRODUCTION"` instances in any case, since a 1-node `"PRODUCTION"` instance
        is functionally identical to a `"DEVELOPMENT"` instance, but without the accompanying restrictions.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to the resource. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

