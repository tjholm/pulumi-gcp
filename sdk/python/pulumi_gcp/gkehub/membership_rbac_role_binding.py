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

__all__ = ['MembershipRbacRoleBindingArgs', 'MembershipRbacRoleBinding']

@pulumi.input_type
class MembershipRbacRoleBindingArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 membership_id: pulumi.Input[str],
                 membership_rbac_role_binding_id: pulumi.Input[str],
                 role: pulumi.Input['MembershipRbacRoleBindingRoleArgs'],
                 user: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MembershipRbacRoleBinding resource.
        :param pulumi.Input[str] location: Location of the Membership
        :param pulumi.Input[str] membership_id: Id of the membership
        :param pulumi.Input[str] membership_rbac_role_binding_id: The client-provided identifier of the RBAC Role Binding.
        :param pulumi.Input['MembershipRbacRoleBindingRoleArgs'] role: Role to bind to the principal.
               Structure is documented below.
        :param pulumi.Input[str] user: Principal that is be authorized in the cluster (at least of one the oneof
               is required). Updating one will unset the other automatically.
               user is the name of the user as seen by the kubernetes cluster, example
               "alice" or "alice@domain.tld"
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "membership_id", membership_id)
        pulumi.set(__self__, "membership_rbac_role_binding_id", membership_rbac_role_binding_id)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "user", user)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        Location of the Membership
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> pulumi.Input[str]:
        """
        Id of the membership
        """
        return pulumi.get(self, "membership_id")

    @membership_id.setter
    def membership_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "membership_id", value)

    @property
    @pulumi.getter(name="membershipRbacRoleBindingId")
    def membership_rbac_role_binding_id(self) -> pulumi.Input[str]:
        """
        The client-provided identifier of the RBAC Role Binding.
        """
        return pulumi.get(self, "membership_rbac_role_binding_id")

    @membership_rbac_role_binding_id.setter
    def membership_rbac_role_binding_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "membership_rbac_role_binding_id", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input['MembershipRbacRoleBindingRoleArgs']:
        """
        Role to bind to the principal.
        Structure is documented below.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input['MembershipRbacRoleBindingRoleArgs']):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        Principal that is be authorized in the cluster (at least of one the oneof
        is required). Updating one will unset the other automatically.
        user is the name of the user as seen by the kubernetes cluster, example
        "alice" or "alice@domain.tld"
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _MembershipRbacRoleBindingState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 delete_time: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 membership_rbac_role_binding_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input['MembershipRbacRoleBindingRoleArgs']] = None,
                 states: Optional[pulumi.Input[Sequence[pulumi.Input['MembershipRbacRoleBindingStateArgs']]]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MembershipRbacRoleBinding resources.
        :param pulumi.Input[str] create_time: Time the RBAC Role Binding was created in UTC.
        :param pulumi.Input[str] delete_time: Time the RBAC Role Binding was deleted in UTC.
        :param pulumi.Input[str] location: Location of the Membership
        :param pulumi.Input[str] membership_id: Id of the membership
        :param pulumi.Input[str] membership_rbac_role_binding_id: The client-provided identifier of the RBAC Role Binding.
        :param pulumi.Input[str] name: The resource name for the RBAC Role Binding
        :param pulumi.Input['MembershipRbacRoleBindingRoleArgs'] role: Role to bind to the principal.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['MembershipRbacRoleBindingStateArgs']]] states: State of the RBAC Role Binding resource.
               Structure is documented below.
        :param pulumi.Input[str] uid: Google-generated UUID for this resource.
        :param pulumi.Input[str] update_time: Time the RBAC Role Binding was updated in UTC.
        :param pulumi.Input[str] user: Principal that is be authorized in the cluster (at least of one the oneof
               is required). Updating one will unset the other automatically.
               user is the name of the user as seen by the kubernetes cluster, example
               "alice" or "alice@domain.tld"
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if delete_time is not None:
            pulumi.set(__self__, "delete_time", delete_time)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if membership_id is not None:
            pulumi.set(__self__, "membership_id", membership_id)
        if membership_rbac_role_binding_id is not None:
            pulumi.set(__self__, "membership_rbac_role_binding_id", membership_rbac_role_binding_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if states is not None:
            pulumi.set(__self__, "states", states)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the RBAC Role Binding was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the RBAC Role Binding was deleted in UTC.
        """
        return pulumi.get(self, "delete_time")

    @delete_time.setter
    def delete_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_time", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the Membership
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the membership
        """
        return pulumi.get(self, "membership_id")

    @membership_id.setter
    def membership_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "membership_id", value)

    @property
    @pulumi.getter(name="membershipRbacRoleBindingId")
    def membership_rbac_role_binding_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client-provided identifier of the RBAC Role Binding.
        """
        return pulumi.get(self, "membership_rbac_role_binding_id")

    @membership_rbac_role_binding_id.setter
    def membership_rbac_role_binding_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "membership_rbac_role_binding_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the RBAC Role Binding
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input['MembershipRbacRoleBindingRoleArgs']]:
        """
        Role to bind to the principal.
        Structure is documented below.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input['MembershipRbacRoleBindingRoleArgs']]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def states(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MembershipRbacRoleBindingStateArgs']]]]:
        """
        State of the RBAC Role Binding resource.
        Structure is documented below.
        """
        return pulumi.get(self, "states")

    @states.setter
    def states(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MembershipRbacRoleBindingStateArgs']]]]):
        pulumi.set(self, "states", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        Google-generated UUID for this resource.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the RBAC Role Binding was updated in UTC.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        Principal that is be authorized in the cluster (at least of one the oneof
        is required). Updating one will unset the other automatically.
        user is the name of the user as seen by the kubernetes cluster, example
        "alice" or "alice@domain.tld"
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class MembershipRbacRoleBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 membership_rbac_role_binding_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingRoleArgs']]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Gkehub Membership Rbac Role Binding Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.container.Cluster("primary",
            name="basic-cluster",
            location="us-central1-a",
            initial_node_count=1,
            deletion_protection=True,
            network="default",
            subnetwork="default")
        membership = gcp.gkehub.Membership("membership",
            membership_id="tf-test-membership_8493",
            endpoint=gcp.gkehub.MembershipEndpointArgs(
                gke_cluster=gcp.gkehub.MembershipEndpointGkeClusterArgs(
                    resource_link=primary.id.apply(lambda id: f"//container.googleapis.com/{id}"),
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[primary]))
        project = gcp.organizations.get_project()
        membership_rbac_role_binding = gcp.gkehub.MembershipRbacRoleBinding("membership_rbac_role_binding",
            membership_rbac_role_binding_id="tf-test-membership-rbac-role-binding_9106",
            membership_id=membership.membership_id,
            user=f"service-{project.number}@gcp-sa-anthossupport.iam.gserviceaccount.com",
            role=gcp.gkehub.MembershipRbacRoleBindingRoleArgs(
                predefined_role="ANTHOS_SUPPORT",
            ),
            location="global",
            opts=pulumi.ResourceOptions(depends_on=[membership]))
        ```

        ## Import

        MembershipRBACRoleBinding can be imported using any of these accepted formats:

        * `projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}`

        * `{{project}}/{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}`

        * `{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}`

        When using the `pulumi import` command, MembershipRBACRoleBinding can be imported using one of the formats above. For example:

        ```sh
        $ pulumi import gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding default projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}
        ```

        ```sh
        $ pulumi import gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding default {{project}}/{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
        ```

        ```sh
        $ pulumi import gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding default {{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Location of the Membership
        :param pulumi.Input[str] membership_id: Id of the membership
        :param pulumi.Input[str] membership_rbac_role_binding_id: The client-provided identifier of the RBAC Role Binding.
        :param pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingRoleArgs']] role: Role to bind to the principal.
               Structure is documented below.
        :param pulumi.Input[str] user: Principal that is be authorized in the cluster (at least of one the oneof
               is required). Updating one will unset the other automatically.
               user is the name of the user as seen by the kubernetes cluster, example
               "alice" or "alice@domain.tld"
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MembershipRbacRoleBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Gkehub Membership Rbac Role Binding Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.container.Cluster("primary",
            name="basic-cluster",
            location="us-central1-a",
            initial_node_count=1,
            deletion_protection=True,
            network="default",
            subnetwork="default")
        membership = gcp.gkehub.Membership("membership",
            membership_id="tf-test-membership_8493",
            endpoint=gcp.gkehub.MembershipEndpointArgs(
                gke_cluster=gcp.gkehub.MembershipEndpointGkeClusterArgs(
                    resource_link=primary.id.apply(lambda id: f"//container.googleapis.com/{id}"),
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[primary]))
        project = gcp.organizations.get_project()
        membership_rbac_role_binding = gcp.gkehub.MembershipRbacRoleBinding("membership_rbac_role_binding",
            membership_rbac_role_binding_id="tf-test-membership-rbac-role-binding_9106",
            membership_id=membership.membership_id,
            user=f"service-{project.number}@gcp-sa-anthossupport.iam.gserviceaccount.com",
            role=gcp.gkehub.MembershipRbacRoleBindingRoleArgs(
                predefined_role="ANTHOS_SUPPORT",
            ),
            location="global",
            opts=pulumi.ResourceOptions(depends_on=[membership]))
        ```

        ## Import

        MembershipRBACRoleBinding can be imported using any of these accepted formats:

        * `projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}`

        * `{{project}}/{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}`

        * `{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}`

        When using the `pulumi import` command, MembershipRBACRoleBinding can be imported using one of the formats above. For example:

        ```sh
        $ pulumi import gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding default projects/{{project}}/locations/{{location}}/memberships/{{membership_id}}/rbacrolebindings/{{membership_rbac_role_binding_id}}
        ```

        ```sh
        $ pulumi import gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding default {{project}}/{{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
        ```

        ```sh
        $ pulumi import gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding default {{location}}/{{membership_id}}/{{membership_rbac_role_binding_id}}
        ```

        :param str resource_name: The name of the resource.
        :param MembershipRbacRoleBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MembershipRbacRoleBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 membership_rbac_role_binding_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingRoleArgs']]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MembershipRbacRoleBindingArgs.__new__(MembershipRbacRoleBindingArgs)

            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if membership_id is None and not opts.urn:
                raise TypeError("Missing required property 'membership_id'")
            __props__.__dict__["membership_id"] = membership_id
            if membership_rbac_role_binding_id is None and not opts.urn:
                raise TypeError("Missing required property 'membership_rbac_role_binding_id'")
            __props__.__dict__["membership_rbac_role_binding_id"] = membership_rbac_role_binding_id
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["states"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        super(MembershipRbacRoleBinding, __self__).__init__(
            'gcp:gkehub/membershipRbacRoleBinding:MembershipRbacRoleBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            delete_time: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            membership_id: Optional[pulumi.Input[str]] = None,
            membership_rbac_role_binding_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingRoleArgs']]] = None,
            states: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingStateArgs']]]]] = None,
            uid: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'MembershipRbacRoleBinding':
        """
        Get an existing MembershipRbacRoleBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Time the RBAC Role Binding was created in UTC.
        :param pulumi.Input[str] delete_time: Time the RBAC Role Binding was deleted in UTC.
        :param pulumi.Input[str] location: Location of the Membership
        :param pulumi.Input[str] membership_id: Id of the membership
        :param pulumi.Input[str] membership_rbac_role_binding_id: The client-provided identifier of the RBAC Role Binding.
        :param pulumi.Input[str] name: The resource name for the RBAC Role Binding
        :param pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingRoleArgs']] role: Role to bind to the principal.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MembershipRbacRoleBindingStateArgs']]]] states: State of the RBAC Role Binding resource.
               Structure is documented below.
        :param pulumi.Input[str] uid: Google-generated UUID for this resource.
        :param pulumi.Input[str] update_time: Time the RBAC Role Binding was updated in UTC.
        :param pulumi.Input[str] user: Principal that is be authorized in the cluster (at least of one the oneof
               is required). Updating one will unset the other automatically.
               user is the name of the user as seen by the kubernetes cluster, example
               "alice" or "alice@domain.tld"
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MembershipRbacRoleBindingState.__new__(_MembershipRbacRoleBindingState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["delete_time"] = delete_time
        __props__.__dict__["location"] = location
        __props__.__dict__["membership_id"] = membership_id
        __props__.__dict__["membership_rbac_role_binding_id"] = membership_rbac_role_binding_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        __props__.__dict__["states"] = states
        __props__.__dict__["uid"] = uid
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["user"] = user
        return MembershipRbacRoleBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time the RBAC Role Binding was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        Time the RBAC Role Binding was deleted in UTC.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Location of the Membership
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> pulumi.Output[str]:
        """
        Id of the membership
        """
        return pulumi.get(self, "membership_id")

    @property
    @pulumi.getter(name="membershipRbacRoleBindingId")
    def membership_rbac_role_binding_id(self) -> pulumi.Output[str]:
        """
        The client-provided identifier of the RBAC Role Binding.
        """
        return pulumi.get(self, "membership_rbac_role_binding_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the RBAC Role Binding
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output['outputs.MembershipRbacRoleBindingRole']:
        """
        Role to bind to the principal.
        Structure is documented below.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def states(self) -> pulumi.Output[Sequence['outputs.MembershipRbacRoleBindingState']]:
        """
        State of the RBAC Role Binding resource.
        Structure is documented below.
        """
        return pulumi.get(self, "states")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Google-generated UUID for this resource.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time the RBAC Role Binding was updated in UTC.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Principal that is be authorized in the cluster (at least of one the oneof
        is required). Updating one will unset the other automatically.
        user is the name of the user as seen by the kubernetes cluster, example
        "alice" or "alice@domain.tld"
        """
        return pulumi.get(self, "user")

