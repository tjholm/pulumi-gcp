# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'QueueAppEngineRoutingOverride',
    'QueueIamBindingCondition',
    'QueueIamMemberCondition',
    'QueueRateLimits',
    'QueueRetryConfig',
    'QueueStackdriverLoggingConfig',
]

@pulumi.output_type
class QueueAppEngineRoutingOverride(dict):
    def __init__(__self__, *,
                 host: Optional[str] = None,
                 instance: Optional[str] = None,
                 service: Optional[str] = None,
                 version: Optional[str] = None):
        """
        :param str host: -
               The host that the task is sent to.
        :param str instance: App instance.
               By default, the task is sent to an instance which is available when the task is attempted.
        :param str service: App service.
               By default, the task is sent to the service which is the default service when the task is attempted.
        :param str version: App version.
               By default, the task is sent to the version which is the default version when the task is attempted.
        """
        if host is not None:
            pulumi.set(__self__, "host", host)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        -
        The host that the task is sent to.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def instance(self) -> Optional[str]:
        """
        App instance.
        By default, the task is sent to an instance which is available when the task is attempted.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def service(self) -> Optional[str]:
        """
        App service.
        By default, the task is sent to the service which is the default service when the task is attempted.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        App version.
        By default, the task is sent to the version which is the default version when the task is attempted.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class QueueIamBindingCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")


@pulumi.output_type
class QueueIamMemberCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")


@pulumi.output_type
class QueueRateLimits(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxBurstSize":
            suggest = "max_burst_size"
        elif key == "maxConcurrentDispatches":
            suggest = "max_concurrent_dispatches"
        elif key == "maxDispatchesPerSecond":
            suggest = "max_dispatches_per_second"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in QueueRateLimits. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        QueueRateLimits.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        QueueRateLimits.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_burst_size: Optional[int] = None,
                 max_concurrent_dispatches: Optional[int] = None,
                 max_dispatches_per_second: Optional[float] = None):
        """
        :param int max_burst_size: -
               The max burst size.
               Max burst size limits how fast tasks in queue are processed when many tasks are
               in the queue and the rate is high. This field allows the queue to have a high
               rate so processing starts shortly after a task is enqueued, but still limits
               resource usage when many tasks are enqueued in a short period of time.
        :param int max_concurrent_dispatches: The maximum number of concurrent tasks that Cloud Tasks allows to
               be dispatched for this queue. After this threshold has been
               reached, Cloud Tasks stops dispatching tasks until the number of
               concurrent requests decreases.
        :param float max_dispatches_per_second: The maximum rate at which tasks are dispatched from this queue.
               If unspecified when the queue is created, Cloud Tasks will pick the default.
        """
        if max_burst_size is not None:
            pulumi.set(__self__, "max_burst_size", max_burst_size)
        if max_concurrent_dispatches is not None:
            pulumi.set(__self__, "max_concurrent_dispatches", max_concurrent_dispatches)
        if max_dispatches_per_second is not None:
            pulumi.set(__self__, "max_dispatches_per_second", max_dispatches_per_second)

    @property
    @pulumi.getter(name="maxBurstSize")
    def max_burst_size(self) -> Optional[int]:
        """
        -
        The max burst size.
        Max burst size limits how fast tasks in queue are processed when many tasks are
        in the queue and the rate is high. This field allows the queue to have a high
        rate so processing starts shortly after a task is enqueued, but still limits
        resource usage when many tasks are enqueued in a short period of time.
        """
        return pulumi.get(self, "max_burst_size")

    @property
    @pulumi.getter(name="maxConcurrentDispatches")
    def max_concurrent_dispatches(self) -> Optional[int]:
        """
        The maximum number of concurrent tasks that Cloud Tasks allows to
        be dispatched for this queue. After this threshold has been
        reached, Cloud Tasks stops dispatching tasks until the number of
        concurrent requests decreases.
        """
        return pulumi.get(self, "max_concurrent_dispatches")

    @property
    @pulumi.getter(name="maxDispatchesPerSecond")
    def max_dispatches_per_second(self) -> Optional[float]:
        """
        The maximum rate at which tasks are dispatched from this queue.
        If unspecified when the queue is created, Cloud Tasks will pick the default.
        """
        return pulumi.get(self, "max_dispatches_per_second")


@pulumi.output_type
class QueueRetryConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "maxAttempts":
            suggest = "max_attempts"
        elif key == "maxBackoff":
            suggest = "max_backoff"
        elif key == "maxDoublings":
            suggest = "max_doublings"
        elif key == "maxRetryDuration":
            suggest = "max_retry_duration"
        elif key == "minBackoff":
            suggest = "min_backoff"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in QueueRetryConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        QueueRetryConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        QueueRetryConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 max_attempts: Optional[int] = None,
                 max_backoff: Optional[str] = None,
                 max_doublings: Optional[int] = None,
                 max_retry_duration: Optional[str] = None,
                 min_backoff: Optional[str] = None):
        """
        :param int max_attempts: Number of attempts per task.
               Cloud Tasks will attempt the task maxAttempts times (that is, if
               the first attempt fails, then there will be maxAttempts - 1
               retries). Must be >= -1.
               If unspecified when the queue is created, Cloud Tasks will pick
               the default.
               -1 indicates unlimited attempts.
        :param str max_backoff: A task will be scheduled for retry between minBackoff and
               maxBackoff duration after it fails, if the queue's RetryConfig
               specifies that the task should be retried.
        :param int max_doublings: The time between retries will double maxDoublings times.
               A task's retry interval starts at minBackoff, then doubles maxDoublings times,
               then increases linearly, and finally retries retries at intervals of maxBackoff
               up to maxAttempts times.
        :param str max_retry_duration: If positive, maxRetryDuration specifies the time limit for
               retrying a failed task, measured from when the task was first
               attempted. Once maxRetryDuration time has passed and the task has
               been attempted maxAttempts times, no further attempts will be
               made and the task will be deleted.
               If zero, then the task age is unlimited.
        :param str min_backoff: A task will be scheduled for retry between minBackoff and
               maxBackoff duration after it fails, if the queue's RetryConfig
               specifies that the task should be retried.
        """
        if max_attempts is not None:
            pulumi.set(__self__, "max_attempts", max_attempts)
        if max_backoff is not None:
            pulumi.set(__self__, "max_backoff", max_backoff)
        if max_doublings is not None:
            pulumi.set(__self__, "max_doublings", max_doublings)
        if max_retry_duration is not None:
            pulumi.set(__self__, "max_retry_duration", max_retry_duration)
        if min_backoff is not None:
            pulumi.set(__self__, "min_backoff", min_backoff)

    @property
    @pulumi.getter(name="maxAttempts")
    def max_attempts(self) -> Optional[int]:
        """
        Number of attempts per task.
        Cloud Tasks will attempt the task maxAttempts times (that is, if
        the first attempt fails, then there will be maxAttempts - 1
        retries). Must be >= -1.
        If unspecified when the queue is created, Cloud Tasks will pick
        the default.
        -1 indicates unlimited attempts.
        """
        return pulumi.get(self, "max_attempts")

    @property
    @pulumi.getter(name="maxBackoff")
    def max_backoff(self) -> Optional[str]:
        """
        A task will be scheduled for retry between minBackoff and
        maxBackoff duration after it fails, if the queue's RetryConfig
        specifies that the task should be retried.
        """
        return pulumi.get(self, "max_backoff")

    @property
    @pulumi.getter(name="maxDoublings")
    def max_doublings(self) -> Optional[int]:
        """
        The time between retries will double maxDoublings times.
        A task's retry interval starts at minBackoff, then doubles maxDoublings times,
        then increases linearly, and finally retries retries at intervals of maxBackoff
        up to maxAttempts times.
        """
        return pulumi.get(self, "max_doublings")

    @property
    @pulumi.getter(name="maxRetryDuration")
    def max_retry_duration(self) -> Optional[str]:
        """
        If positive, maxRetryDuration specifies the time limit for
        retrying a failed task, measured from when the task was first
        attempted. Once maxRetryDuration time has passed and the task has
        been attempted maxAttempts times, no further attempts will be
        made and the task will be deleted.
        If zero, then the task age is unlimited.
        """
        return pulumi.get(self, "max_retry_duration")

    @property
    @pulumi.getter(name="minBackoff")
    def min_backoff(self) -> Optional[str]:
        """
        A task will be scheduled for retry between minBackoff and
        maxBackoff duration after it fails, if the queue's RetryConfig
        specifies that the task should be retried.
        """
        return pulumi.get(self, "min_backoff")


@pulumi.output_type
class QueueStackdriverLoggingConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "samplingRatio":
            suggest = "sampling_ratio"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in QueueStackdriverLoggingConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        QueueStackdriverLoggingConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        QueueStackdriverLoggingConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 sampling_ratio: float):
        """
        :param float sampling_ratio: Specifies the fraction of operations to write to Stackdriver Logging.
               This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
               default and means that no operations are logged.
        """
        pulumi.set(__self__, "sampling_ratio", sampling_ratio)

    @property
    @pulumi.getter(name="samplingRatio")
    def sampling_ratio(self) -> float:
        """
        Specifies the fraction of operations to write to Stackdriver Logging.
        This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
        default and means that no operations are logged.
        """
        return pulumi.get(self, "sampling_ratio")


