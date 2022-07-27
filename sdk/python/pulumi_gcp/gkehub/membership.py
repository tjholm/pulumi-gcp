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

__all__ = ['MembershipArgs', 'Membership']

@pulumi.input_type
class MembershipArgs:
    def __init__(__self__, *,
                 membership_id: pulumi.Input[str],
                 authority: Optional[pulumi.Input['MembershipAuthorityArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input['MembershipEndpointArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Membership resource.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input['MembershipAuthorityArgs'] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        :param pulumi.Input['MembershipEndpointArgs'] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "membership_id", membership_id)
        if authority is not None:
            pulumi.set(__self__, "authority", authority)
        if description is not None:
            warnings.warn("""This field is unavailable in the GA provider and will be removed from the beta provider in a future release.""", DeprecationWarning)
            pulumi.log.warn("""description is deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.""")
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> pulumi.Input[str]:
        """
        The client-provided identifier of the membership.
        """
        return pulumi.get(self, "membership_id")

    @membership_id.setter
    def membership_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "membership_id", value)

    @property
    @pulumi.getter
    def authority(self) -> Optional[pulumi.Input['MembershipAuthorityArgs']]:
        """
        Authority encodes how Google will recognize identities from this Membership.
        See the workload identity documentation for more details:
        https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
        Structure is documented below.
        """
        return pulumi.get(self, "authority")

    @authority.setter
    def authority(self, value: Optional[pulumi.Input['MembershipAuthorityArgs']]):
        pulumi.set(self, "authority", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input['MembershipEndpointArgs']]:
        """
        If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
        Structure is documented below.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input['MembershipEndpointArgs']]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to apply to this membership.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

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
class _MembershipState:
    def __init__(__self__, *,
                 authority: Optional[pulumi.Input['MembershipAuthorityArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input['MembershipEndpointArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Membership resources.
        :param pulumi.Input['MembershipAuthorityArgs'] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        :param pulumi.Input['MembershipEndpointArgs'] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input[str] name: The unique identifier of the membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if authority is not None:
            pulumi.set(__self__, "authority", authority)
        if description is not None:
            warnings.warn("""This field is unavailable in the GA provider and will be removed from the beta provider in a future release.""", DeprecationWarning)
            pulumi.log.warn("""description is deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.""")
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if membership_id is not None:
            pulumi.set(__self__, "membership_id", membership_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def authority(self) -> Optional[pulumi.Input['MembershipAuthorityArgs']]:
        """
        Authority encodes how Google will recognize identities from this Membership.
        See the workload identity documentation for more details:
        https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
        Structure is documented below.
        """
        return pulumi.get(self, "authority")

    @authority.setter
    def authority(self, value: Optional[pulumi.Input['MembershipAuthorityArgs']]):
        pulumi.set(self, "authority", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input['MembershipEndpointArgs']]:
        """
        If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
        Structure is documented below.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input['MembershipEndpointArgs']]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to apply to this membership.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client-provided identifier of the membership.
        """
        return pulumi.get(self, "membership_id")

    @membership_id.setter
    def membership_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "membership_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the membership.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class Membership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authority: Optional[pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[pulumi.InputType['MembershipEndpointArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Membership contains information about a member cluster.

        To get more information about Membership, see:

        * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.memberships)
        * How-to Guides
            * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)

        ## Example Usage
        ### Gkehub Membership Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.container.Cluster("primary",
            initial_node_count=1,
            location="us-central1-a")
        membership = gcp.gkehub.Membership("membership",
            endpoint=gcp.gkehub.MembershipEndpointArgs(
                gke_cluster=gcp.gkehub.MembershipEndpointGkeClusterArgs(
                    resource_link=primary.id.apply(lambda id: f"//container.googleapis.com/{id}"),
                ),
            ),
            membership_id="basic")
        ```
        ### Gkehub Membership Issuer

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.container.Cluster("primary",
            location="us-central1-a",
            initial_node_count=1,
            workload_identity_config=gcp.container.ClusterWorkloadIdentityConfigArgs(
                workload_pool="my-project-name.svc.id.goog",
            ))
        membership = gcp.gkehub.Membership("membership",
            membership_id="basic",
            endpoint=gcp.gkehub.MembershipEndpointArgs(
                gke_cluster=gcp.gkehub.MembershipEndpointGkeClusterArgs(
                    resource_link=primary.id,
                ),
            ),
            authority=gcp.gkehub.MembershipAuthorityArgs(
                issuer=primary.id.apply(lambda id: f"https://container.googleapis.com/v1/{id}"),
            ))
        ```

        ## Import

        Membership can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gkehub/membership:Membership default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        :param pulumi.Input[pulumi.InputType['MembershipEndpointArgs']] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Membership contains information about a member cluster.

        To get more information about Membership, see:

        * [API documentation](https://cloud.google.com/anthos/multicluster-management/reference/rest/v1/projects.locations.memberships)
        * How-to Guides
            * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)

        ## Example Usage
        ### Gkehub Membership Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.container.Cluster("primary",
            initial_node_count=1,
            location="us-central1-a")
        membership = gcp.gkehub.Membership("membership",
            endpoint=gcp.gkehub.MembershipEndpointArgs(
                gke_cluster=gcp.gkehub.MembershipEndpointGkeClusterArgs(
                    resource_link=primary.id.apply(lambda id: f"//container.googleapis.com/{id}"),
                ),
            ),
            membership_id="basic")
        ```
        ### Gkehub Membership Issuer

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.container.Cluster("primary",
            location="us-central1-a",
            initial_node_count=1,
            workload_identity_config=gcp.container.ClusterWorkloadIdentityConfigArgs(
                workload_pool="my-project-name.svc.id.goog",
            ))
        membership = gcp.gkehub.Membership("membership",
            membership_id="basic",
            endpoint=gcp.gkehub.MembershipEndpointArgs(
                gke_cluster=gcp.gkehub.MembershipEndpointGkeClusterArgs(
                    resource_link=primary.id,
                ),
            ),
            authority=gcp.gkehub.MembershipAuthorityArgs(
                issuer=primary.id.apply(lambda id: f"https://container.googleapis.com/v1/{id}"),
            ))
        ```

        ## Import

        Membership can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gkehub/membership:Membership default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param MembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authority: Optional[pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[pulumi.InputType['MembershipEndpointArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MembershipArgs.__new__(MembershipArgs)

            __props__.__dict__["authority"] = authority
            if description is not None and not opts.urn:
                warnings.warn("""This field is unavailable in the GA provider and will be removed from the beta provider in a future release.""", DeprecationWarning)
                pulumi.log.warn("""description is deprecated: This field is unavailable in the GA provider and will be removed from the beta provider in a future release.""")
            __props__.__dict__["description"] = description
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["labels"] = labels
            if membership_id is None and not opts.urn:
                raise TypeError("Missing required property 'membership_id'")
            __props__.__dict__["membership_id"] = membership_id
            __props__.__dict__["project"] = project
            __props__.__dict__["name"] = None
        super(Membership, __self__).__init__(
            'gcp:gkehub/membership:Membership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authority: Optional[pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[pulumi.InputType['MembershipEndpointArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            membership_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Membership':
        """
        Get an existing Membership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        :param pulumi.Input[pulumi.InputType['MembershipEndpointArgs']] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input[str] name: The unique identifier of the membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MembershipState.__new__(_MembershipState)

        __props__.__dict__["authority"] = authority
        __props__.__dict__["description"] = description
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["labels"] = labels
        __props__.__dict__["membership_id"] = membership_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        return Membership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authority(self) -> pulumi.Output[Optional['outputs.MembershipAuthority']]:
        """
        Authority encodes how Google will recognize identities from this Membership.
        See the workload identity documentation for more details:
        https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
        Structure is documented below.
        """
        return pulumi.get(self, "authority")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional['outputs.MembershipEndpoint']]:
        """
        If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
        Structure is documented below.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to apply to this membership.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> pulumi.Output[str]:
        """
        The client-provided identifier of the membership.
        """
        return pulumi.get(self, "membership_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the membership.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

