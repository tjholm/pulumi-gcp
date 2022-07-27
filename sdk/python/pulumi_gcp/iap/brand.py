# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BrandArgs', 'Brand']

@pulumi.input_type
class BrandArgs:
    def __init__(__self__, *,
                 application_title: pulumi.Input[str],
                 support_email: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Brand resource.
        :param pulumi.Input[str] application_title: Application name displayed on OAuth consent screen.
        :param pulumi.Input[str] support_email: Support email displayed on the OAuth consent screen. Can be either a
               user or group email. When a user email is specified, the caller must
               be the user with the associated email address. When a group email is
               specified, the caller can be either a user or a service account which
               is an owner of the specified group in Cloud Identity.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "application_title", application_title)
        pulumi.set(__self__, "support_email", support_email)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="applicationTitle")
    def application_title(self) -> pulumi.Input[str]:
        """
        Application name displayed on OAuth consent screen.
        """
        return pulumi.get(self, "application_title")

    @application_title.setter
    def application_title(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_title", value)

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> pulumi.Input[str]:
        """
        Support email displayed on the OAuth consent screen. Can be either a
        user or group email. When a user email is specified, the caller must
        be the user with the associated email address. When a group email is
        specified, the caller can be either a user or a service account which
        is an owner of the specified group in Cloud Identity.
        """
        return pulumi.get(self, "support_email")

    @support_email.setter
    def support_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "support_email", value)

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
class _BrandState:
    def __init__(__self__, *,
                 application_title: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_internal_only: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Brand resources.
        :param pulumi.Input[str] application_title: Application name displayed on OAuth consent screen.
        :param pulumi.Input[str] name: Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
               identification corresponds to the project number as only one brand per project can be created.
        :param pulumi.Input[bool] org_internal_only: Whether the brand is only intended for usage inside the GSuite organization only.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] support_email: Support email displayed on the OAuth consent screen. Can be either a
               user or group email. When a user email is specified, the caller must
               be the user with the associated email address. When a group email is
               specified, the caller can be either a user or a service account which
               is an owner of the specified group in Cloud Identity.
        """
        if application_title is not None:
            pulumi.set(__self__, "application_title", application_title)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_internal_only is not None:
            pulumi.set(__self__, "org_internal_only", org_internal_only)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if support_email is not None:
            pulumi.set(__self__, "support_email", support_email)

    @property
    @pulumi.getter(name="applicationTitle")
    def application_title(self) -> Optional[pulumi.Input[str]]:
        """
        Application name displayed on OAuth consent screen.
        """
        return pulumi.get(self, "application_title")

    @application_title.setter
    def application_title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_title", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
        identification corresponds to the project number as only one brand per project can be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgInternalOnly")
    def org_internal_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the brand is only intended for usage inside the GSuite organization only.
        """
        return pulumi.get(self, "org_internal_only")

    @org_internal_only.setter
    def org_internal_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "org_internal_only", value)

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
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> Optional[pulumi.Input[str]]:
        """
        Support email displayed on the OAuth consent screen. Can be either a
        user or group email. When a user email is specified, the caller must
        be the user with the associated email address. When a group email is
        specified, the caller can be either a user or a service account which
        is an owner of the specified group in Cloud Identity.
        """
        return pulumi.get(self, "support_email")

    @support_email.setter
    def support_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "support_email", value)


class Brand(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_title: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        OAuth brand data. Only "Organization Internal" brands can be created
        programmatically via API. To convert it into an external brands
        please use the GCP Console.

        > **Note:** Brands can only be created once for a Google Cloud
        project and the underlying Google API doesn't not support DELETE or PATCH methods.
        Destroying a provider-managed Brand will remove it from state
        but *will not delete it from Google Cloud.*

        To get more information about Brand, see:

        * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands)
        * How-to Guides
            * [Setting up IAP Brand](https://cloud.google.com/iap/docs/tutorial-gce#set_up_iap)

        ## Example Usage
        ### Iap Brand

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project",
            project_id="tf-test",
            org_id="123456789")
        project_service = gcp.projects.Service("projectService",
            project=project.project_id,
            service="iap.googleapis.com")
        project_brand = gcp.iap.Brand("projectBrand",
            support_email="support@example.com",
            application_title="Cloud IAP protected Application",
            project=project_service.project)
        ```

        ## Import

        Brand can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iap/brand:Brand default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_title: Application name displayed on OAuth consent screen.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] support_email: Support email displayed on the OAuth consent screen. Can be either a
               user or group email. When a user email is specified, the caller must
               be the user with the associated email address. When a group email is
               specified, the caller can be either a user or a service account which
               is an owner of the specified group in Cloud Identity.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BrandArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        OAuth brand data. Only "Organization Internal" brands can be created
        programmatically via API. To convert it into an external brands
        please use the GCP Console.

        > **Note:** Brands can only be created once for a Google Cloud
        project and the underlying Google API doesn't not support DELETE or PATCH methods.
        Destroying a provider-managed Brand will remove it from state
        but *will not delete it from Google Cloud.*

        To get more information about Brand, see:

        * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands)
        * How-to Guides
            * [Setting up IAP Brand](https://cloud.google.com/iap/docs/tutorial-gce#set_up_iap)

        ## Example Usage
        ### Iap Brand

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project",
            project_id="tf-test",
            org_id="123456789")
        project_service = gcp.projects.Service("projectService",
            project=project.project_id,
            service="iap.googleapis.com")
        project_brand = gcp.iap.Brand("projectBrand",
            support_email="support@example.com",
            application_title="Cloud IAP protected Application",
            project=project_service.project)
        ```

        ## Import

        Brand can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iap/brand:Brand default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param BrandArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BrandArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_title: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BrandArgs.__new__(BrandArgs)

            if application_title is None and not opts.urn:
                raise TypeError("Missing required property 'application_title'")
            __props__.__dict__["application_title"] = application_title
            __props__.__dict__["project"] = project
            if support_email is None and not opts.urn:
                raise TypeError("Missing required property 'support_email'")
            __props__.__dict__["support_email"] = support_email
            __props__.__dict__["name"] = None
            __props__.__dict__["org_internal_only"] = None
        super(Brand, __self__).__init__(
            'gcp:iap/brand:Brand',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_title: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_internal_only: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            support_email: Optional[pulumi.Input[str]] = None) -> 'Brand':
        """
        Get an existing Brand resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_title: Application name displayed on OAuth consent screen.
        :param pulumi.Input[str] name: Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
               identification corresponds to the project number as only one brand per project can be created.
        :param pulumi.Input[bool] org_internal_only: Whether the brand is only intended for usage inside the GSuite organization only.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] support_email: Support email displayed on the OAuth consent screen. Can be either a
               user or group email. When a user email is specified, the caller must
               be the user with the associated email address. When a group email is
               specified, the caller can be either a user or a service account which
               is an owner of the specified group in Cloud Identity.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BrandState.__new__(_BrandState)

        __props__.__dict__["application_title"] = application_title
        __props__.__dict__["name"] = name
        __props__.__dict__["org_internal_only"] = org_internal_only
        __props__.__dict__["project"] = project
        __props__.__dict__["support_email"] = support_email
        return Brand(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationTitle")
    def application_title(self) -> pulumi.Output[str]:
        """
        Application name displayed on OAuth consent screen.
        """
        return pulumi.get(self, "application_title")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
        identification corresponds to the project number as only one brand per project can be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgInternalOnly")
    def org_internal_only(self) -> pulumi.Output[bool]:
        """
        Whether the brand is only intended for usage inside the GSuite organization only.
        """
        return pulumi.get(self, "org_internal_only")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> pulumi.Output[str]:
        """
        Support email displayed on the OAuth consent screen. Can be either a
        user or group email. When a user email is specified, the caller must
        be the user with the associated email address. When a group email is
        specified, the caller can be either a user or a service account which
        is an owner of the specified group in Cloud Identity.
        """
        return pulumi.get(self, "support_email")

