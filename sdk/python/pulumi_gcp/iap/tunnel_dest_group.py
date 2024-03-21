# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TunnelDestGroupArgs', 'TunnelDestGroup']

@pulumi.input_type
class TunnelDestGroupArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[str],
                 cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fqdns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TunnelDestGroup resource.
        :param pulumi.Input[str] group_name: Unique tunnel destination group name.
               
               
               - - -
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidrs: List of CIDRs that this group applies to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fqdns: List of FQDNs that this group applies to.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the tunnel group. Must be the same as the network resources in the group.
        """
        pulumi.set(__self__, "group_name", group_name)
        if cidrs is not None:
            pulumi.set(__self__, "cidrs", cidrs)
        if fqdns is not None:
            pulumi.set(__self__, "fqdns", fqdns)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        Unique tunnel destination group name.


        - - -
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter
    def cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of CIDRs that this group applies to.
        """
        return pulumi.get(self, "cidrs")

    @cidrs.setter
    def cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidrs", value)

    @property
    @pulumi.getter
    def fqdns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of FQDNs that this group applies to.
        """
        return pulumi.get(self, "fqdns")

    @fqdns.setter
    def fqdns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "fqdns", value)

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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the tunnel group. Must be the same as the network resources in the group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _TunnelDestGroupState:
    def __init__(__self__, *,
                 cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fqdns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TunnelDestGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidrs: List of CIDRs that this group applies to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fqdns: List of FQDNs that this group applies to.
        :param pulumi.Input[str] group_name: Unique tunnel destination group name.
               
               
               - - -
        :param pulumi.Input[str] name: Full resource name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the tunnel group. Must be the same as the network resources in the group.
        """
        if cidrs is not None:
            pulumi.set(__self__, "cidrs", cidrs)
        if fqdns is not None:
            pulumi.set(__self__, "fqdns", fqdns)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def cidrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of CIDRs that this group applies to.
        """
        return pulumi.get(self, "cidrs")

    @cidrs.setter
    def cidrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidrs", value)

    @property
    @pulumi.getter
    def fqdns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of FQDNs that this group applies to.
        """
        return pulumi.get(self, "fqdns")

    @fqdns.setter
    def fqdns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "fqdns", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique tunnel destination group name.


        - - -
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Full resource name.
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

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the tunnel group. Must be the same as the network resources in the group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class TunnelDestGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fqdns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Tunnel destination groups represent resources that have the same tunnel access restrictions.

        To get more information about TunnelDestGroup, see:

        * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.iap_tunnel.locations.destGroups)
        * How-to Guides
            * [Set up IAP TCP forwarding with an IP address or hostname in a Google Cloud or non-Google Cloud environment](https://cloud.google.com/iap/docs/tcp-by-host)

        ## Example Usage

        ### Iap Destgroup

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gcp as gcp

        dest_group = gcp.iap.TunnelDestGroup("dest_group",
            region="us-central1",
            group_name="testgroup_75223",
            cidrs=[
                "10.1.0.0/16",
                "192.168.10.0/24",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        TunnelDestGroup can be imported using any of these accepted formats:

        * `projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`

        * `{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`

        * `{{project}}/{{region}}/{{group_name}}`

        * `{{region}}/destGroups/{{group_name}}`

        * `{{region}}/{{group_name}}`

        * `{{group_name}}`

        When using the `pulumi import` command, TunnelDestGroup can be imported using one of the formats above. For example:

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{project}}/{{region}}/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{region}}/destGroups/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{region}}/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{group_name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidrs: List of CIDRs that this group applies to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fqdns: List of FQDNs that this group applies to.
        :param pulumi.Input[str] group_name: Unique tunnel destination group name.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the tunnel group. Must be the same as the network resources in the group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TunnelDestGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Tunnel destination groups represent resources that have the same tunnel access restrictions.

        To get more information about TunnelDestGroup, see:

        * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.iap_tunnel.locations.destGroups)
        * How-to Guides
            * [Set up IAP TCP forwarding with an IP address or hostname in a Google Cloud or non-Google Cloud environment](https://cloud.google.com/iap/docs/tcp-by-host)

        ## Example Usage

        ### Iap Destgroup

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gcp as gcp

        dest_group = gcp.iap.TunnelDestGroup("dest_group",
            region="us-central1",
            group_name="testgroup_75223",
            cidrs=[
                "10.1.0.0/16",
                "192.168.10.0/24",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        TunnelDestGroup can be imported using any of these accepted formats:

        * `projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`

        * `{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}`

        * `{{project}}/{{region}}/{{group_name}}`

        * `{{region}}/destGroups/{{group_name}}`

        * `{{region}}/{{group_name}}`

        * `{{group_name}}`

        When using the `pulumi import` command, TunnelDestGroup can be imported using one of the formats above. For example:

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{project}}/iap_tunnel/locations/{{region}}/destGroups/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{project}}/{{region}}/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{region}}/destGroups/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{region}}/{{group_name}}
        ```

        ```sh
        $ pulumi import gcp:iap/tunnelDestGroup:TunnelDestGroup default {{group_name}}
        ```

        :param str resource_name: The name of the resource.
        :param TunnelDestGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TunnelDestGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 fqdns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TunnelDestGroupArgs.__new__(TunnelDestGroupArgs)

            __props__.__dict__["cidrs"] = cidrs
            __props__.__dict__["fqdns"] = fqdns
            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["name"] = None
        super(TunnelDestGroup, __self__).__init__(
            'gcp:iap/tunnelDestGroup:TunnelDestGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            fqdns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'TunnelDestGroup':
        """
        Get an existing TunnelDestGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidrs: List of CIDRs that this group applies to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fqdns: List of FQDNs that this group applies to.
        :param pulumi.Input[str] group_name: Unique tunnel destination group name.
               
               
               - - -
        :param pulumi.Input[str] name: Full resource name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the tunnel group. Must be the same as the network resources in the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TunnelDestGroupState.__new__(_TunnelDestGroupState)

        __props__.__dict__["cidrs"] = cidrs
        __props__.__dict__["fqdns"] = fqdns
        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        return TunnelDestGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidrs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of CIDRs that this group applies to.
        """
        return pulumi.get(self, "cidrs")

    @property
    @pulumi.getter
    def fqdns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of FQDNs that this group applies to.
        """
        return pulumi.get(self, "fqdns")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        Unique tunnel destination group name.


        - - -
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Full resource name.
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

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region of the tunnel group. Must be the same as the network resources in the group.
        """
        return pulumi.get(self, "region")

