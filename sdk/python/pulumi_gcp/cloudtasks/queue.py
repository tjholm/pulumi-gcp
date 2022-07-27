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

__all__ = ['QueueArgs', 'Queue']

@pulumi.input_type
class QueueArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 app_engine_routing_override: Optional[pulumi.Input['QueueAppEngineRoutingOverrideArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rate_limits: Optional[pulumi.Input['QueueRateLimitsArgs']] = None,
                 retry_config: Optional[pulumi.Input['QueueRetryConfigArgs']] = None,
                 stackdriver_logging_config: Optional[pulumi.Input['QueueStackdriverLoggingConfigArgs']] = None):
        """
        The set of arguments for constructing a Queue resource.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input['QueueAppEngineRoutingOverrideArgs'] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue
               Structure is documented below.
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['QueueRateLimitsArgs'] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.
               Structure is documented below.
        :param pulumi.Input['QueueRetryConfigArgs'] retry_config: Settings that determine the retry behavior.
               Structure is documented below.
        :param pulumi.Input['QueueStackdriverLoggingConfigArgs'] stackdriver_logging_config: Configuration options for writing logs to Stackdriver Logging.
               Structure is documented below.
        """
        pulumi.set(__self__, "location", location)
        if app_engine_routing_override is not None:
            pulumi.set(__self__, "app_engine_routing_override", app_engine_routing_override)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if rate_limits is not None:
            pulumi.set(__self__, "rate_limits", rate_limits)
        if retry_config is not None:
            pulumi.set(__self__, "retry_config", retry_config)
        if stackdriver_logging_config is not None:
            pulumi.set(__self__, "stackdriver_logging_config", stackdriver_logging_config)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location of the queue
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="appEngineRoutingOverride")
    def app_engine_routing_override(self) -> Optional[pulumi.Input['QueueAppEngineRoutingOverrideArgs']]:
        """
        Overrides for task-level appEngineRouting. These settings apply only
        to App Engine tasks in this queue
        Structure is documented below.
        """
        return pulumi.get(self, "app_engine_routing_override")

    @app_engine_routing_override.setter
    def app_engine_routing_override(self, value: Optional[pulumi.Input['QueueAppEngineRoutingOverrideArgs']]):
        pulumi.set(self, "app_engine_routing_override", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The queue name.
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
    @pulumi.getter(name="rateLimits")
    def rate_limits(self) -> Optional[pulumi.Input['QueueRateLimitsArgs']]:
        """
        Rate limits for task dispatches.
        The queue's actual dispatch rate is the result of:
        * Number of tasks in the queue
        * User-specified throttling: rateLimits, retryConfig, and the queue's state.
        * System throttling due to 429 (Too Many Requests) or 503 (Service
        Unavailable) responses from the worker, high error rates, or to
        smooth sudden large traffic spikes.
        Structure is documented below.
        """
        return pulumi.get(self, "rate_limits")

    @rate_limits.setter
    def rate_limits(self, value: Optional[pulumi.Input['QueueRateLimitsArgs']]):
        pulumi.set(self, "rate_limits", value)

    @property
    @pulumi.getter(name="retryConfig")
    def retry_config(self) -> Optional[pulumi.Input['QueueRetryConfigArgs']]:
        """
        Settings that determine the retry behavior.
        Structure is documented below.
        """
        return pulumi.get(self, "retry_config")

    @retry_config.setter
    def retry_config(self, value: Optional[pulumi.Input['QueueRetryConfigArgs']]):
        pulumi.set(self, "retry_config", value)

    @property
    @pulumi.getter(name="stackdriverLoggingConfig")
    def stackdriver_logging_config(self) -> Optional[pulumi.Input['QueueStackdriverLoggingConfigArgs']]:
        """
        Configuration options for writing logs to Stackdriver Logging.
        Structure is documented below.
        """
        return pulumi.get(self, "stackdriver_logging_config")

    @stackdriver_logging_config.setter
    def stackdriver_logging_config(self, value: Optional[pulumi.Input['QueueStackdriverLoggingConfigArgs']]):
        pulumi.set(self, "stackdriver_logging_config", value)


@pulumi.input_type
class _QueueState:
    def __init__(__self__, *,
                 app_engine_routing_override: Optional[pulumi.Input['QueueAppEngineRoutingOverrideArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rate_limits: Optional[pulumi.Input['QueueRateLimitsArgs']] = None,
                 retry_config: Optional[pulumi.Input['QueueRetryConfigArgs']] = None,
                 stackdriver_logging_config: Optional[pulumi.Input['QueueStackdriverLoggingConfigArgs']] = None):
        """
        Input properties used for looking up and filtering Queue resources.
        :param pulumi.Input['QueueAppEngineRoutingOverrideArgs'] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['QueueRateLimitsArgs'] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.
               Structure is documented below.
        :param pulumi.Input['QueueRetryConfigArgs'] retry_config: Settings that determine the retry behavior.
               Structure is documented below.
        :param pulumi.Input['QueueStackdriverLoggingConfigArgs'] stackdriver_logging_config: Configuration options for writing logs to Stackdriver Logging.
               Structure is documented below.
        """
        if app_engine_routing_override is not None:
            pulumi.set(__self__, "app_engine_routing_override", app_engine_routing_override)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if rate_limits is not None:
            pulumi.set(__self__, "rate_limits", rate_limits)
        if retry_config is not None:
            pulumi.set(__self__, "retry_config", retry_config)
        if stackdriver_logging_config is not None:
            pulumi.set(__self__, "stackdriver_logging_config", stackdriver_logging_config)

    @property
    @pulumi.getter(name="appEngineRoutingOverride")
    def app_engine_routing_override(self) -> Optional[pulumi.Input['QueueAppEngineRoutingOverrideArgs']]:
        """
        Overrides for task-level appEngineRouting. These settings apply only
        to App Engine tasks in this queue
        Structure is documented below.
        """
        return pulumi.get(self, "app_engine_routing_override")

    @app_engine_routing_override.setter
    def app_engine_routing_override(self, value: Optional[pulumi.Input['QueueAppEngineRoutingOverrideArgs']]):
        pulumi.set(self, "app_engine_routing_override", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the queue
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The queue name.
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
    @pulumi.getter(name="rateLimits")
    def rate_limits(self) -> Optional[pulumi.Input['QueueRateLimitsArgs']]:
        """
        Rate limits for task dispatches.
        The queue's actual dispatch rate is the result of:
        * Number of tasks in the queue
        * User-specified throttling: rateLimits, retryConfig, and the queue's state.
        * System throttling due to 429 (Too Many Requests) or 503 (Service
        Unavailable) responses from the worker, high error rates, or to
        smooth sudden large traffic spikes.
        Structure is documented below.
        """
        return pulumi.get(self, "rate_limits")

    @rate_limits.setter
    def rate_limits(self, value: Optional[pulumi.Input['QueueRateLimitsArgs']]):
        pulumi.set(self, "rate_limits", value)

    @property
    @pulumi.getter(name="retryConfig")
    def retry_config(self) -> Optional[pulumi.Input['QueueRetryConfigArgs']]:
        """
        Settings that determine the retry behavior.
        Structure is documented below.
        """
        return pulumi.get(self, "retry_config")

    @retry_config.setter
    def retry_config(self, value: Optional[pulumi.Input['QueueRetryConfigArgs']]):
        pulumi.set(self, "retry_config", value)

    @property
    @pulumi.getter(name="stackdriverLoggingConfig")
    def stackdriver_logging_config(self) -> Optional[pulumi.Input['QueueStackdriverLoggingConfigArgs']]:
        """
        Configuration options for writing logs to Stackdriver Logging.
        Structure is documented below.
        """
        return pulumi.get(self, "stackdriver_logging_config")

    @stackdriver_logging_config.setter
    def stackdriver_logging_config(self, value: Optional[pulumi.Input['QueueStackdriverLoggingConfigArgs']]):
        pulumi.set(self, "stackdriver_logging_config", value)


class Queue(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_engine_routing_override: Optional[pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rate_limits: Optional[pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']]] = None,
                 retry_config: Optional[pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']]] = None,
                 stackdriver_logging_config: Optional[pulumi.Input[pulumi.InputType['QueueStackdriverLoggingConfigArgs']]] = None,
                 __props__=None):
        """
        A named resource to which messages are sent by publishers.

        > **Warning:** This resource requires an App Engine application to be created on the
        project you're provisioning it on. If you haven't already enabled it, you
        can create a `appengine.Application` resource to do so. This
        resource's location will be the same as the App Engine location specified.

        ## Example Usage
        ### Queue Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.cloudtasks.Queue("default", location="us-central1")
        ```
        ### Cloud Tasks Queue Advanced

        ```python
        import pulumi
        import pulumi_gcp as gcp

        advanced_configuration = gcp.cloudtasks.Queue("advancedConfiguration",
            app_engine_routing_override=gcp.cloudtasks.QueueAppEngineRoutingOverrideArgs(
                instance="test",
                service="worker",
                version="1.0",
            ),
            location="us-central1",
            rate_limits=gcp.cloudtasks.QueueRateLimitsArgs(
                max_concurrent_dispatches=3,
                max_dispatches_per_second=2,
            ),
            retry_config=gcp.cloudtasks.QueueRetryConfigArgs(
                max_attempts=5,
                max_backoff="3s",
                max_doublings=1,
                max_retry_duration="4s",
                min_backoff="2s",
            ),
            stackdriver_logging_config=gcp.cloudtasks.QueueStackdriverLoggingConfigArgs(
                sampling_ratio=0.9,
            ))
        ```

        ## Import

        Queue can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudtasks/queue:Queue default projects/{{project}}/locations/{{location}}/queues/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudtasks/queue:Queue default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudtasks/queue:Queue default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']] retry_config: Settings that determine the retry behavior.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['QueueStackdriverLoggingConfigArgs']] stackdriver_logging_config: Configuration options for writing logs to Stackdriver Logging.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QueueArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A named resource to which messages are sent by publishers.

        > **Warning:** This resource requires an App Engine application to be created on the
        project you're provisioning it on. If you haven't already enabled it, you
        can create a `appengine.Application` resource to do so. This
        resource's location will be the same as the App Engine location specified.

        ## Example Usage
        ### Queue Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.cloudtasks.Queue("default", location="us-central1")
        ```
        ### Cloud Tasks Queue Advanced

        ```python
        import pulumi
        import pulumi_gcp as gcp

        advanced_configuration = gcp.cloudtasks.Queue("advancedConfiguration",
            app_engine_routing_override=gcp.cloudtasks.QueueAppEngineRoutingOverrideArgs(
                instance="test",
                service="worker",
                version="1.0",
            ),
            location="us-central1",
            rate_limits=gcp.cloudtasks.QueueRateLimitsArgs(
                max_concurrent_dispatches=3,
                max_dispatches_per_second=2,
            ),
            retry_config=gcp.cloudtasks.QueueRetryConfigArgs(
                max_attempts=5,
                max_backoff="3s",
                max_doublings=1,
                max_retry_duration="4s",
                min_backoff="2s",
            ),
            stackdriver_logging_config=gcp.cloudtasks.QueueStackdriverLoggingConfigArgs(
                sampling_ratio=0.9,
            ))
        ```

        ## Import

        Queue can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudtasks/queue:Queue default projects/{{project}}/locations/{{location}}/queues/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudtasks/queue:Queue default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudtasks/queue:Queue default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param QueueArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QueueArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_engine_routing_override: Optional[pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rate_limits: Optional[pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']]] = None,
                 retry_config: Optional[pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']]] = None,
                 stackdriver_logging_config: Optional[pulumi.Input[pulumi.InputType['QueueStackdriverLoggingConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QueueArgs.__new__(QueueArgs)

            __props__.__dict__["app_engine_routing_override"] = app_engine_routing_override
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["rate_limits"] = rate_limits
            __props__.__dict__["retry_config"] = retry_config
            __props__.__dict__["stackdriver_logging_config"] = stackdriver_logging_config
        super(Queue, __self__).__init__(
            'gcp:cloudtasks/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_engine_routing_override: Optional[pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            rate_limits: Optional[pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']]] = None,
            retry_config: Optional[pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']]] = None,
            stackdriver_logging_config: Optional[pulumi.Input[pulumi.InputType['QueueStackdriverLoggingConfigArgs']]] = None) -> 'Queue':
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['QueueAppEngineRoutingOverrideArgs']] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only
               to App Engine tasks in this queue
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['QueueRateLimitsArgs']] rate_limits: Rate limits for task dispatches.
               The queue's actual dispatch rate is the result of:
               * Number of tasks in the queue
               * User-specified throttling: rateLimits, retryConfig, and the queue's state.
               * System throttling due to 429 (Too Many Requests) or 503 (Service
               Unavailable) responses from the worker, high error rates, or to
               smooth sudden large traffic spikes.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['QueueRetryConfigArgs']] retry_config: Settings that determine the retry behavior.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['QueueStackdriverLoggingConfigArgs']] stackdriver_logging_config: Configuration options for writing logs to Stackdriver Logging.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QueueState.__new__(_QueueState)

        __props__.__dict__["app_engine_routing_override"] = app_engine_routing_override
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["rate_limits"] = rate_limits
        __props__.__dict__["retry_config"] = retry_config
        __props__.__dict__["stackdriver_logging_config"] = stackdriver_logging_config
        return Queue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appEngineRoutingOverride")
    def app_engine_routing_override(self) -> pulumi.Output[Optional['outputs.QueueAppEngineRoutingOverride']]:
        """
        Overrides for task-level appEngineRouting. These settings apply only
        to App Engine tasks in this queue
        Structure is documented below.
        """
        return pulumi.get(self, "app_engine_routing_override")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the queue
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The queue name.
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
    @pulumi.getter(name="rateLimits")
    def rate_limits(self) -> pulumi.Output['outputs.QueueRateLimits']:
        """
        Rate limits for task dispatches.
        The queue's actual dispatch rate is the result of:
        * Number of tasks in the queue
        * User-specified throttling: rateLimits, retryConfig, and the queue's state.
        * System throttling due to 429 (Too Many Requests) or 503 (Service
        Unavailable) responses from the worker, high error rates, or to
        smooth sudden large traffic spikes.
        Structure is documented below.
        """
        return pulumi.get(self, "rate_limits")

    @property
    @pulumi.getter(name="retryConfig")
    def retry_config(self) -> pulumi.Output['outputs.QueueRetryConfig']:
        """
        Settings that determine the retry behavior.
        Structure is documented below.
        """
        return pulumi.get(self, "retry_config")

    @property
    @pulumi.getter(name="stackdriverLoggingConfig")
    def stackdriver_logging_config(self) -> pulumi.Output[Optional['outputs.QueueStackdriverLoggingConfig']]:
        """
        Configuration options for writing logs to Stackdriver Logging.
        Structure is documented below.
        """
        return pulumi.get(self, "stackdriver_logging_config")

