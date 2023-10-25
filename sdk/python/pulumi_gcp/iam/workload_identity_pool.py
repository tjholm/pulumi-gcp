# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WorkloadIdentityPoolArgs', 'WorkloadIdentityPool']

@pulumi.input_type
class WorkloadIdentityPoolArgs:
    def __init__(__self__, *,
                 workload_identity_pool_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WorkloadIdentityPool resource.
        :param pulumi.Input[str] workload_identity_pool_id: The ID to use for the pool, which becomes the final component of the resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
               
               
               - - -
        :param pulumi.Input[str] description: A description of the pool. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
               existing tokens to access resources. If the pool is re-enabled, existing tokens grant
               access again.
        :param pulumi.Input[str] display_name: A display name for the pool. Cannot exceed 32 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        WorkloadIdentityPoolArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            workload_identity_pool_id=workload_identity_pool_id,
            description=description,
            disabled=disabled,
            display_name=display_name,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             disabled: Optional[pulumi.Input[bool]] = None,
             display_name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if workload_identity_pool_id is None and 'workloadIdentityPoolId' in kwargs:
            workload_identity_pool_id = kwargs['workloadIdentityPoolId']
        if workload_identity_pool_id is None:
            raise TypeError("Missing 'workload_identity_pool_id' argument")
        if display_name is None and 'displayName' in kwargs:
            display_name = kwargs['displayName']

        _setter("workload_identity_pool_id", workload_identity_pool_id)
        if description is not None:
            _setter("description", description)
        if disabled is not None:
            _setter("disabled", disabled)
        if display_name is not None:
            _setter("display_name", display_name)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> pulumi.Input[str]:
        """
        The ID to use for the pool, which becomes the final component of the resource name. This
        value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        `gcp-` is reserved for use by Google, and may not be specified.


        - - -
        """
        return pulumi.get(self, "workload_identity_pool_id")

    @workload_identity_pool_id.setter
    def workload_identity_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workload_identity_pool_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the pool. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        access again.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A display name for the pool. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

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
class _WorkloadIdentityPoolState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WorkloadIdentityPool resources.
        :param pulumi.Input[str] description: A description of the pool. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
               existing tokens to access resources. If the pool is re-enabled, existing tokens grant
               access again.
        :param pulumi.Input[str] display_name: A display name for the pool. Cannot exceed 32 characters.
        :param pulumi.Input[str] name: The resource name of the pool as
               `projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: The state of the pool.
               * STATE_UNSPECIFIED: State unspecified.
               * ACTIVE: The pool is active, and may be used in Google Cloud policies.
               * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after
               approximately 30 days. You can restore a soft-deleted pool using
               UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until it is
               permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or
               use existing tokens to access resources. If the pool is undeleted, existing tokens grant
               access again.
        :param pulumi.Input[str] workload_identity_pool_id: The ID to use for the pool, which becomes the final component of the resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
               
               
               - - -
        """
        _WorkloadIdentityPoolState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            disabled=disabled,
            display_name=display_name,
            name=name,
            project=project,
            state=state,
            workload_identity_pool_id=workload_identity_pool_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             disabled: Optional[pulumi.Input[bool]] = None,
             display_name: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[str]] = None,
             workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if display_name is None and 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if workload_identity_pool_id is None and 'workloadIdentityPoolId' in kwargs:
            workload_identity_pool_id = kwargs['workloadIdentityPoolId']

        if description is not None:
            _setter("description", description)
        if disabled is not None:
            _setter("disabled", disabled)
        if display_name is not None:
            _setter("display_name", display_name)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if state is not None:
            _setter("state", state)
        if workload_identity_pool_id is not None:
            _setter("workload_identity_pool_id", workload_identity_pool_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the pool. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        access again.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A display name for the pool. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the pool as
        `projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}`.
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
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the pool.
        * STATE_UNSPECIFIED: State unspecified.
        * ACTIVE: The pool is active, and may be used in Google Cloud policies.
        * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after
        approximately 30 days. You can restore a soft-deleted pool using
        UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until it is
        permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or
        use existing tokens to access resources. If the pool is undeleted, existing tokens grant
        access again.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for the pool, which becomes the final component of the resource name. This
        value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        `gcp-` is reserved for use by Google, and may not be specified.


        - - -
        """
        return pulumi.get(self, "workload_identity_pool_id")

    @workload_identity_pool_id.setter
    def workload_identity_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workload_identity_pool_id", value)


class WorkloadIdentityPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a collection of external workload identities. You can define IAM policies to
        grant these identities access to Google Cloud resources.

        To get more information about WorkloadIdentityPool, see:

        * [API documentation](https://cloud.google.com/iam/docs/reference/rest/v1/projects.locations.workloadIdentityPools)
        * How-to Guides
            * [Managing workload identity pools](https://cloud.google.com/iam/docs/manage-workload-identity-pools-providers#pools)

        ## Example Usage
        ### Iam Workload Identity Pool Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.iam.WorkloadIdentityPool("example", workload_identity_pool_id="example-pool")
        ```
        ### Iam Workload Identity Pool Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.iam.WorkloadIdentityPool("example",
            description="Identity pool for automated test",
            disabled=True,
            display_name="Name of pool",
            workload_identity_pool_id="example-pool")
        ```

        ## Import

        WorkloadIdentityPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{project}}/{{workload_identity_pool_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{workload_identity_pool_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the pool. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
               existing tokens to access resources. If the pool is re-enabled, existing tokens grant
               access again.
        :param pulumi.Input[str] display_name: A display name for the pool. Cannot exceed 32 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] workload_identity_pool_id: The ID to use for the pool, which becomes the final component of the resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
               
               
               - - -
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkloadIdentityPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a collection of external workload identities. You can define IAM policies to
        grant these identities access to Google Cloud resources.

        To get more information about WorkloadIdentityPool, see:

        * [API documentation](https://cloud.google.com/iam/docs/reference/rest/v1/projects.locations.workloadIdentityPools)
        * How-to Guides
            * [Managing workload identity pools](https://cloud.google.com/iam/docs/manage-workload-identity-pools-providers#pools)

        ## Example Usage
        ### Iam Workload Identity Pool Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.iam.WorkloadIdentityPool("example", workload_identity_pool_id="example-pool")
        ```
        ### Iam Workload Identity Pool Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.iam.WorkloadIdentityPool("example",
            description="Identity pool for automated test",
            disabled=True,
            display_name="Name of pool",
            workload_identity_pool_id="example-pool")
        ```

        ## Import

        WorkloadIdentityPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{project}}/{{workload_identity_pool_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{workload_identity_pool_id}}
        ```

        :param str resource_name: The name of the resource.
        :param WorkloadIdentityPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkloadIdentityPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            WorkloadIdentityPoolArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkloadIdentityPoolArgs.__new__(WorkloadIdentityPoolArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["project"] = project
            if workload_identity_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_identity_pool_id'")
            __props__.__dict__["workload_identity_pool_id"] = workload_identity_pool_id
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
        super(WorkloadIdentityPool, __self__).__init__(
            'gcp:iam/workloadIdentityPool:WorkloadIdentityPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            workload_identity_pool_id: Optional[pulumi.Input[str]] = None) -> 'WorkloadIdentityPool':
        """
        Get an existing WorkloadIdentityPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the pool. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
               existing tokens to access resources. If the pool is re-enabled, existing tokens grant
               access again.
        :param pulumi.Input[str] display_name: A display name for the pool. Cannot exceed 32 characters.
        :param pulumi.Input[str] name: The resource name of the pool as
               `projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: The state of the pool.
               * STATE_UNSPECIFIED: State unspecified.
               * ACTIVE: The pool is active, and may be used in Google Cloud policies.
               * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after
               approximately 30 days. You can restore a soft-deleted pool using
               UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until it is
               permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or
               use existing tokens to access resources. If the pool is undeleted, existing tokens grant
               access again.
        :param pulumi.Input[str] workload_identity_pool_id: The ID to use for the pool, which becomes the final component of the resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
               
               
               - - -
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkloadIdentityPoolState.__new__(_WorkloadIdentityPoolState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["state"] = state
        __props__.__dict__["workload_identity_pool_id"] = workload_identity_pool_id
        return WorkloadIdentityPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the pool. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        access again.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A display name for the pool. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the pool as
        `projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}`.
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
    def state(self) -> pulumi.Output[str]:
        """
        The state of the pool.
        * STATE_UNSPECIFIED: State unspecified.
        * ACTIVE: The pool is active, and may be used in Google Cloud policies.
        * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after
        approximately 30 days. You can restore a soft-deleted pool using
        UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until it is
        permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or
        use existing tokens to access resources. If the pool is undeleted, existing tokens grant
        access again.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> pulumi.Output[str]:
        """
        The ID to use for the pool, which becomes the final component of the resource name. This
        value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        `gcp-` is reserved for use by Google, and may not be specified.


        - - -
        """
        return pulumi.get(self, "workload_identity_pool_id")

