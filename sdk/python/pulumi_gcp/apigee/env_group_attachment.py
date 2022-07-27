# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnvGroupAttachmentArgs', 'EnvGroupAttachment']

@pulumi.input_type
class EnvGroupAttachmentArgs:
    def __init__(__self__, *,
                 envgroup_id: pulumi.Input[str],
                 environment: pulumi.Input[str]):
        """
        The set of arguments for constructing a EnvGroupAttachment resource.
        :param pulumi.Input[str] envgroup_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        """
        pulumi.set(__self__, "envgroup_id", envgroup_id)
        pulumi.set(__self__, "environment", environment)

    @property
    @pulumi.getter(name="envgroupId")
    def envgroup_id(self) -> pulumi.Input[str]:
        """
        The Apigee environment group associated with the Apigee environment,
        in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        """
        return pulumi.get(self, "envgroup_id")

    @envgroup_id.setter
    def envgroup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "envgroup_id", value)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        The resource ID of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)


@pulumi.input_type
class _EnvGroupAttachmentState:
    def __init__(__self__, *,
                 envgroup_id: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnvGroupAttachment resources.
        :param pulumi.Input[str] envgroup_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        :param pulumi.Input[str] name: The name of the newly created attachment (output parameter).
        """
        if envgroup_id is not None:
            pulumi.set(__self__, "envgroup_id", envgroup_id)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="envgroupId")
    def envgroup_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Apigee environment group associated with the Apigee environment,
        in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        """
        return pulumi.get(self, "envgroup_id")

    @envgroup_id.setter
    def envgroup_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "envgroup_id", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the newly created attachment (output parameter).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class EnvGroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 envgroup_id: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An `Environment Group attachment` in Apigee.

        To get more information about EnvgroupAttachment, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups.attachments/create)
        * How-to Guides
            * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

        ## Example Usage
        ### Apigee Environment Group Attachment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project",
            project_id="tf-test",
            org_id="",
            billing_account="")
        apigee = gcp.projects.Service("apigee",
            project=project.project_id,
            service="apigee.googleapis.com")
        compute = gcp.projects.Service("compute",
            project=project.project_id,
            service="compute.googleapis.com")
        servicenetworking = gcp.projects.Service("servicenetworking",
            project=project.project_id,
            service="servicenetworking.googleapis.com")
        apigee_network = gcp.compute.Network("apigeeNetwork", project=project.project_id,
        opts=pulumi.ResourceOptions(depends_on=[compute]))
        apigee_range = gcp.compute.GlobalAddress("apigeeRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=apigee_network.id,
            project=project.project_id)
        apigee_vpc_connection = gcp.servicenetworking.Connection("apigeeVpcConnection",
            network=apigee_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[apigee_range.name],
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        apigee_org = gcp.apigee.Organization("apigeeOrg",
            analytics_region="us-central1",
            project_id=project.project_id,
            authorized_network=apigee_network.id,
            opts=pulumi.ResourceOptions(depends_on=[
                    apigee_vpc_connection,
                    apigee,
                ]))
        apigee_envgroup = gcp.apigee.EnvGroup("apigeeEnvgroup",
            org_id=apigee_org.id,
            hostnames=["abc.foo.com"])
        apigee_env = gcp.apigee.Environment("apigeeEnv",
            org_id=apigee_org.id,
            description="Apigee Environment",
            display_name="tf-test")
        env_group_attachment = gcp.apigee.EnvGroupAttachment("envGroupAttachment",
            envgroup_id=apigee_envgroup.id,
            environment=apigee_env.name)
        ```

        ## Import

        EnvgroupAttachment can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/attachments/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] envgroup_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An `Environment Group attachment` in Apigee.

        To get more information about EnvgroupAttachment, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.envgroups.attachments/create)
        * How-to Guides
            * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

        ## Example Usage
        ### Apigee Environment Group Attachment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project",
            project_id="tf-test",
            org_id="",
            billing_account="")
        apigee = gcp.projects.Service("apigee",
            project=project.project_id,
            service="apigee.googleapis.com")
        compute = gcp.projects.Service("compute",
            project=project.project_id,
            service="compute.googleapis.com")
        servicenetworking = gcp.projects.Service("servicenetworking",
            project=project.project_id,
            service="servicenetworking.googleapis.com")
        apigee_network = gcp.compute.Network("apigeeNetwork", project=project.project_id,
        opts=pulumi.ResourceOptions(depends_on=[compute]))
        apigee_range = gcp.compute.GlobalAddress("apigeeRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=apigee_network.id,
            project=project.project_id)
        apigee_vpc_connection = gcp.servicenetworking.Connection("apigeeVpcConnection",
            network=apigee_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[apigee_range.name],
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        apigee_org = gcp.apigee.Organization("apigeeOrg",
            analytics_region="us-central1",
            project_id=project.project_id,
            authorized_network=apigee_network.id,
            opts=pulumi.ResourceOptions(depends_on=[
                    apigee_vpc_connection,
                    apigee,
                ]))
        apigee_envgroup = gcp.apigee.EnvGroup("apigeeEnvgroup",
            org_id=apigee_org.id,
            hostnames=["abc.foo.com"])
        apigee_env = gcp.apigee.Environment("apigeeEnv",
            org_id=apigee_org.id,
            description="Apigee Environment",
            display_name="tf-test")
        env_group_attachment = gcp.apigee.EnvGroupAttachment("envGroupAttachment",
            envgroup_id=apigee_envgroup.id,
            environment=apigee_env.name)
        ```

        ## Import

        EnvgroupAttachment can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/attachments/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/envGroupAttachment:EnvGroupAttachment default {{envgroup_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param EnvGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 envgroup_id: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvGroupAttachmentArgs.__new__(EnvGroupAttachmentArgs)

            if envgroup_id is None and not opts.urn:
                raise TypeError("Missing required property 'envgroup_id'")
            __props__.__dict__["envgroup_id"] = envgroup_id
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            __props__.__dict__["name"] = None
        super(EnvGroupAttachment, __self__).__init__(
            'gcp:apigee/envGroupAttachment:EnvGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            envgroup_id: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'EnvGroupAttachment':
        """
        Get an existing EnvGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] envgroup_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        :param pulumi.Input[str] name: The name of the newly created attachment (output parameter).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvGroupAttachmentState.__new__(_EnvGroupAttachmentState)

        __props__.__dict__["envgroup_id"] = envgroup_id
        __props__.__dict__["environment"] = environment
        __props__.__dict__["name"] = name
        return EnvGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="envgroupId")
    def envgroup_id(self) -> pulumi.Output[str]:
        """
        The Apigee environment group associated with the Apigee environment,
        in the format `organizations/{{org_name}}/envgroups/{{envgroup_name}}`.
        """
        return pulumi.get(self, "envgroup_id")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        The resource ID of the environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the newly created attachment (output parameter).
        """
        return pulumi.get(self, "name")

