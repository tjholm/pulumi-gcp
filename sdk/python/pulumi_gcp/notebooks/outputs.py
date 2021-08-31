# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'EnvironmentContainerImage',
    'EnvironmentVmImage',
    'InstanceAcceleratorConfig',
    'InstanceContainerImage',
    'InstanceIamBindingCondition',
    'InstanceIamMemberCondition',
    'InstanceReservationAffinity',
    'InstanceShieldedInstanceConfig',
    'InstanceVmImage',
]

@pulumi.output_type
class EnvironmentContainerImage(dict):
    def __init__(__self__, *,
                 repository: str,
                 tag: Optional[str] = None):
        """
        :param str repository: The path to the container image repository.
               For example: gcr.io/{project_id}/{imageName}
        :param str tag: The tag of the container image. If not specified, this defaults to the latest tag.
        """
        pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def repository(self) -> str:
        """
        The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        The tag of the container image. If not specified, this defaults to the latest tag.
        """
        return pulumi.get(self, "tag")


@pulumi.output_type
class EnvironmentVmImage(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "imageFamily":
            suggest = "image_family"
        elif key == "imageName":
            suggest = "image_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EnvironmentVmImage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EnvironmentVmImage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EnvironmentVmImage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 project: str,
                 image_family: Optional[str] = None,
                 image_name: Optional[str] = None):
        """
        :param str project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param str image_family: Use this VM image family to find the image; the newest image in this family will be used.
        :param str image_name: Use VM image name to find the image.
        """
        pulumi.set(__self__, "project", project)
        if image_family is not None:
            pulumi.set(__self__, "image_family", image_family)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[str]:
        """
        Use this VM image family to find the image; the newest image in this family will be used.
        """
        return pulumi.get(self, "image_family")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        """
        Use VM image name to find the image.
        """
        return pulumi.get(self, "image_name")


@pulumi.output_type
class InstanceAcceleratorConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "coreCount":
            suggest = "core_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceAcceleratorConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceAcceleratorConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceAcceleratorConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 core_count: int,
                 type: str):
        """
        :param int core_count: Count of cores of this accelerator.
        :param str type: Type of this accelerator.
               Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
        """
        pulumi.set(__self__, "core_count", core_count)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="coreCount")
    def core_count(self) -> int:
        """
        Count of cores of this accelerator.
        """
        return pulumi.get(self, "core_count")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of this accelerator.
        Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class InstanceContainerImage(dict):
    def __init__(__self__, *,
                 repository: str,
                 tag: Optional[str] = None):
        """
        :param str repository: The path to the container image repository.
               For example: gcr.io/{project_id}/{imageName}
        :param str tag: The tag of the container image. If not specified, this defaults to the latest tag.
        """
        pulumi.set(__self__, "repository", repository)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def repository(self) -> str:
        """
        The path to the container image repository.
        For example: gcr.io/{project_id}/{imageName}
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        """
        The tag of the container image. If not specified, this defaults to the latest tag.
        """
        return pulumi.get(self, "tag")


@pulumi.output_type
class InstanceIamBindingCondition(dict):
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
class InstanceIamMemberCondition(dict):
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
class InstanceReservationAffinity(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "consumeReservationType":
            suggest = "consume_reservation_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceReservationAffinity. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceReservationAffinity.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceReservationAffinity.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 consume_reservation_type: str,
                 key: Optional[str] = None,
                 values: Optional[Sequence[str]] = None):
        """
        :param str consume_reservation_type: The type of Compute Reservation.
               Possible values are `NO_RESERVATION`, `ANY_RESERVATION`, and `SPECIFIC_RESERVATION`.
        :param str key: Corresponds to the label key of reservation resource.
        :param Sequence[str] values: Corresponds to the label values of reservation resource.
        """
        pulumi.set(__self__, "consume_reservation_type", consume_reservation_type)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="consumeReservationType")
    def consume_reservation_type(self) -> str:
        """
        The type of Compute Reservation.
        Possible values are `NO_RESERVATION`, `ANY_RESERVATION`, and `SPECIFIC_RESERVATION`.
        """
        return pulumi.get(self, "consume_reservation_type")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        Corresponds to the label key of reservation resource.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def values(self) -> Optional[Sequence[str]]:
        """
        Corresponds to the label values of reservation resource.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class InstanceShieldedInstanceConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enableIntegrityMonitoring":
            suggest = "enable_integrity_monitoring"
        elif key == "enableSecureBoot":
            suggest = "enable_secure_boot"
        elif key == "enableVtpm":
            suggest = "enable_vtpm"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceShieldedInstanceConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceShieldedInstanceConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceShieldedInstanceConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enable_integrity_monitoring: Optional[bool] = None,
                 enable_secure_boot: Optional[bool] = None,
                 enable_vtpm: Optional[bool] = None):
        """
        :param bool enable_integrity_monitoring: Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the
               boot integrity of the instance. The attestation is performed against the integrity policy baseline.
               This baseline is initially derived from the implicitly trusted boot image when the instance is created.
               Enabled by default.
        :param bool enable_secure_boot: Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs
               authentic software by verifying the digital signature of all boot components, and halting the boot process
               if signature verification fails.
               Disabled by default.
        :param bool enable_vtpm: Defines whether the instance has the vTPM enabled.
               Enabled by default.
        """
        if enable_integrity_monitoring is not None:
            pulumi.set(__self__, "enable_integrity_monitoring", enable_integrity_monitoring)
        if enable_secure_boot is not None:
            pulumi.set(__self__, "enable_secure_boot", enable_secure_boot)
        if enable_vtpm is not None:
            pulumi.set(__self__, "enable_vtpm", enable_vtpm)

    @property
    @pulumi.getter(name="enableIntegrityMonitoring")
    def enable_integrity_monitoring(self) -> Optional[bool]:
        """
        Defines whether the instance has integrity monitoring enabled. Enables monitoring and attestation of the
        boot integrity of the instance. The attestation is performed against the integrity policy baseline.
        This baseline is initially derived from the implicitly trusted boot image when the instance is created.
        Enabled by default.
        """
        return pulumi.get(self, "enable_integrity_monitoring")

    @property
    @pulumi.getter(name="enableSecureBoot")
    def enable_secure_boot(self) -> Optional[bool]:
        """
        Defines whether the instance has Secure Boot enabled. Secure Boot helps ensure that the system only runs
        authentic software by verifying the digital signature of all boot components, and halting the boot process
        if signature verification fails.
        Disabled by default.
        """
        return pulumi.get(self, "enable_secure_boot")

    @property
    @pulumi.getter(name="enableVtpm")
    def enable_vtpm(self) -> Optional[bool]:
        """
        Defines whether the instance has the vTPM enabled.
        Enabled by default.
        """
        return pulumi.get(self, "enable_vtpm")


@pulumi.output_type
class InstanceVmImage(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "imageFamily":
            suggest = "image_family"
        elif key == "imageName":
            suggest = "image_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceVmImage. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceVmImage.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceVmImage.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 project: str,
                 image_family: Optional[str] = None,
                 image_name: Optional[str] = None):
        """
        :param str project: The name of the Google Cloud project that this VM image belongs to.
               Format: projects/{project_id}
        :param str image_family: Use this VM image family to find the image; the newest image in this family will be used.
        :param str image_name: Use VM image name to find the image.
        """
        pulumi.set(__self__, "project", project)
        if image_family is not None:
            pulumi.set(__self__, "image_family", image_family)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The name of the Google Cloud project that this VM image belongs to.
        Format: projects/{project_id}
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="imageFamily")
    def image_family(self) -> Optional[str]:
        """
        Use this VM image family to find the image; the newest image in this family will be used.
        """
        return pulumi.get(self, "image_family")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[str]:
        """
        Use VM image name to find the image.
        """
        return pulumi.get(self, "image_name")


