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

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 default_admission_rule: pulumi.Input['PolicyDefaultAdmissionRuleArgs'],
                 admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]]] = None,
                 cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input['PolicyDefaultAdmissionRuleArgs'] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
               
               Identifier format: `{{location}}.{{clusterId}}`.
               A location is either a compute zone (e.g. `us-central1-a`) or a region
               (e.g. `us-central1`).
               Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
               Possible values are: `ENABLE`, `DISABLE`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "default_admission_rule", default_admission_rule)
        if admission_whitelist_patterns is not None:
            pulumi.set(__self__, "admission_whitelist_patterns", admission_whitelist_patterns)
        if cluster_admission_rules is not None:
            pulumi.set(__self__, "cluster_admission_rules", cluster_admission_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if global_policy_evaluation_mode is not None:
            pulumi.set(__self__, "global_policy_evaluation_mode", global_policy_evaluation_mode)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="defaultAdmissionRule")
    def default_admission_rule(self) -> pulumi.Input['PolicyDefaultAdmissionRuleArgs']:
        """
        Default admission rule for a cluster without a per-cluster admission
        rule.
        Structure is documented below.
        """
        return pulumi.get(self, "default_admission_rule")

    @default_admission_rule.setter
    def default_admission_rule(self, value: pulumi.Input['PolicyDefaultAdmissionRuleArgs']):
        pulumi.set(self, "default_admission_rule", value)

    @property
    @pulumi.getter(name="admissionWhitelistPatterns")
    def admission_whitelist_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]]]:
        """
        A whitelist of image patterns to exclude from admission rules. If an
        image's name matches a whitelist pattern, the image's admission
        requests will always be permitted regardless of your admission rules.
        Structure is documented below.
        """
        return pulumi.get(self, "admission_whitelist_patterns")

    @admission_whitelist_patterns.setter
    def admission_whitelist_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]]]):
        pulumi.set(self, "admission_whitelist_patterns", value)

    @property
    @pulumi.getter(name="clusterAdmissionRules")
    def cluster_admission_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]]]:
        """
        Per-cluster admission rules. An admission rule specifies either that
        all container images used in a pod creation request must be attested
        to by one or more attestors, that all pod creations will be allowed,
        or that all pod creations will be denied. There can be at most one
        admission rule per cluster spec.

        Identifier format: `{{location}}.{{clusterId}}`.
        A location is either a compute zone (e.g. `us-central1-a`) or a region
        (e.g. `us-central1`).
        Structure is documented below.
        """
        return pulumi.get(self, "cluster_admission_rules")

    @cluster_admission_rules.setter
    def cluster_admission_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]]]):
        pulumi.set(self, "cluster_admission_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive comment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="globalPolicyEvaluationMode")
    def global_policy_evaluation_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Controls the evaluation of a Google-maintained global admission policy
        for common system-level images. Images not covered by the global
        policy will be subject to the project admission policy.
        Possible values are: `ENABLE`, `DISABLE`.
        """
        return pulumi.get(self, "global_policy_evaluation_mode")

    @global_policy_evaluation_mode.setter
    def global_policy_evaluation_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_policy_evaluation_mode", value)

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
class _PolicyState:
    def __init__(__self__, *,
                 admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]]] = None,
                 cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]]] = None,
                 default_admission_rule: Optional[pulumi.Input['PolicyDefaultAdmissionRuleArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Policy resources.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
               
               Identifier format: `{{location}}.{{clusterId}}`.
               A location is either a compute zone (e.g. `us-central1-a`) or a region
               (e.g. `us-central1`).
               Structure is documented below.
        :param pulumi.Input['PolicyDefaultAdmissionRuleArgs'] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.
               Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
               Possible values are: `ENABLE`, `DISABLE`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if admission_whitelist_patterns is not None:
            pulumi.set(__self__, "admission_whitelist_patterns", admission_whitelist_patterns)
        if cluster_admission_rules is not None:
            pulumi.set(__self__, "cluster_admission_rules", cluster_admission_rules)
        if default_admission_rule is not None:
            pulumi.set(__self__, "default_admission_rule", default_admission_rule)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if global_policy_evaluation_mode is not None:
            pulumi.set(__self__, "global_policy_evaluation_mode", global_policy_evaluation_mode)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="admissionWhitelistPatterns")
    def admission_whitelist_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]]]:
        """
        A whitelist of image patterns to exclude from admission rules. If an
        image's name matches a whitelist pattern, the image's admission
        requests will always be permitted regardless of your admission rules.
        Structure is documented below.
        """
        return pulumi.get(self, "admission_whitelist_patterns")

    @admission_whitelist_patterns.setter
    def admission_whitelist_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyAdmissionWhitelistPatternArgs']]]]):
        pulumi.set(self, "admission_whitelist_patterns", value)

    @property
    @pulumi.getter(name="clusterAdmissionRules")
    def cluster_admission_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]]]:
        """
        Per-cluster admission rules. An admission rule specifies either that
        all container images used in a pod creation request must be attested
        to by one or more attestors, that all pod creations will be allowed,
        or that all pod creations will be denied. There can be at most one
        admission rule per cluster spec.

        Identifier format: `{{location}}.{{clusterId}}`.
        A location is either a compute zone (e.g. `us-central1-a`) or a region
        (e.g. `us-central1`).
        Structure is documented below.
        """
        return pulumi.get(self, "cluster_admission_rules")

    @cluster_admission_rules.setter
    def cluster_admission_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicyClusterAdmissionRuleArgs']]]]):
        pulumi.set(self, "cluster_admission_rules", value)

    @property
    @pulumi.getter(name="defaultAdmissionRule")
    def default_admission_rule(self) -> Optional[pulumi.Input['PolicyDefaultAdmissionRuleArgs']]:
        """
        Default admission rule for a cluster without a per-cluster admission
        rule.
        Structure is documented below.
        """
        return pulumi.get(self, "default_admission_rule")

    @default_admission_rule.setter
    def default_admission_rule(self, value: Optional[pulumi.Input['PolicyDefaultAdmissionRuleArgs']]):
        pulumi.set(self, "default_admission_rule", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive comment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="globalPolicyEvaluationMode")
    def global_policy_evaluation_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Controls the evaluation of a Google-maintained global admission policy
        for common system-level images. Images not covered by the global
        policy will be subject to the project admission policy.
        Possible values are: `ENABLE`, `DISABLE`.
        """
        return pulumi.get(self, "global_policy_evaluation_mode")

    @global_policy_evaluation_mode.setter
    def global_policy_evaluation_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "global_policy_evaluation_mode", value)

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


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]]] = None,
                 cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]]] = None,
                 default_admission_rule: Optional[pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A policy for container image binary authorization.

        To get more information about Policy, see:

        * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/binary-authorization/)

        ## Example Usage
        ### Binary Authorization Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority=gcp.containeranalysis.NoteAttestationAuthorityArgs(
            hint=gcp.containeranalysis.NoteAttestationAuthorityHintArgs(
                human_readable_name="My attestor",
            ),
        ))
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note=gcp.binaryauthorization.AttestorAttestationAuthorityNoteArgs(
            note_reference=note.name,
        ))
        policy = gcp.binaryauthorization.Policy("policy",
            admission_whitelist_patterns=[gcp.binaryauthorization.PolicyAdmissionWhitelistPatternArgs(
                name_pattern="gcr.io/google_containers/*",
            )],
            default_admission_rule=gcp.binaryauthorization.PolicyDefaultAdmissionRuleArgs(
                evaluation_mode="ALWAYS_ALLOW",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
            ),
            cluster_admission_rules=[gcp.binaryauthorization.PolicyClusterAdmissionRuleArgs(
                cluster="us-central1-a.prod-cluster",
                evaluation_mode="REQUIRE_ATTESTATION",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
                require_attestations_bies=[attestor.name],
            )])
        ```
        ### Binary Authorization Policy Global Evaluation

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority=gcp.containeranalysis.NoteAttestationAuthorityArgs(
            hint=gcp.containeranalysis.NoteAttestationAuthorityHintArgs(
                human_readable_name="My attestor",
            ),
        ))
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note=gcp.binaryauthorization.AttestorAttestationAuthorityNoteArgs(
            note_reference=note.name,
        ))
        policy = gcp.binaryauthorization.Policy("policy",
            default_admission_rule=gcp.binaryauthorization.PolicyDefaultAdmissionRuleArgs(
                evaluation_mode="REQUIRE_ATTESTATION",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
                require_attestations_bies=[attestor.name],
            ),
            global_policy_evaluation_mode="ENABLE")
        ```

        ## Import

        Policy can be imported using any of these accepted formats* `projects/{{project}}` * `{{project}}` When using the `pulumi import` command, Policy can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
        ```

        ```sh
         $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
               
               Identifier format: `{{location}}.{{clusterId}}`.
               A location is either a compute zone (e.g. `us-central1-a`) or a region
               (e.g. `us-central1`).
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.
               Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
               Possible values are: `ENABLE`, `DISABLE`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A policy for container image binary authorization.

        To get more information about Policy, see:

        * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/binary-authorization/)

        ## Example Usage
        ### Binary Authorization Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority=gcp.containeranalysis.NoteAttestationAuthorityArgs(
            hint=gcp.containeranalysis.NoteAttestationAuthorityHintArgs(
                human_readable_name="My attestor",
            ),
        ))
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note=gcp.binaryauthorization.AttestorAttestationAuthorityNoteArgs(
            note_reference=note.name,
        ))
        policy = gcp.binaryauthorization.Policy("policy",
            admission_whitelist_patterns=[gcp.binaryauthorization.PolicyAdmissionWhitelistPatternArgs(
                name_pattern="gcr.io/google_containers/*",
            )],
            default_admission_rule=gcp.binaryauthorization.PolicyDefaultAdmissionRuleArgs(
                evaluation_mode="ALWAYS_ALLOW",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
            ),
            cluster_admission_rules=[gcp.binaryauthorization.PolicyClusterAdmissionRuleArgs(
                cluster="us-central1-a.prod-cluster",
                evaluation_mode="REQUIRE_ATTESTATION",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
                require_attestations_bies=[attestor.name],
            )])
        ```
        ### Binary Authorization Policy Global Evaluation

        ```python
        import pulumi
        import pulumi_gcp as gcp

        note = gcp.containeranalysis.Note("note", attestation_authority=gcp.containeranalysis.NoteAttestationAuthorityArgs(
            hint=gcp.containeranalysis.NoteAttestationAuthorityHintArgs(
                human_readable_name="My attestor",
            ),
        ))
        attestor = gcp.binaryauthorization.Attestor("attestor", attestation_authority_note=gcp.binaryauthorization.AttestorAttestationAuthorityNoteArgs(
            note_reference=note.name,
        ))
        policy = gcp.binaryauthorization.Policy("policy",
            default_admission_rule=gcp.binaryauthorization.PolicyDefaultAdmissionRuleArgs(
                evaluation_mode="REQUIRE_ATTESTATION",
                enforcement_mode="ENFORCED_BLOCK_AND_AUDIT_LOG",
                require_attestations_bies=[attestor.name],
            ),
            global_policy_evaluation_mode="ENABLE")
        ```

        ## Import

        Policy can be imported using any of these accepted formats* `projects/{{project}}` * `{{project}}` When using the `pulumi import` command, Policy can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
        ```

        ```sh
         $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]]] = None,
                 cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]]] = None,
                 default_admission_rule: Optional[pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyArgs.__new__(PolicyArgs)

            __props__.__dict__["admission_whitelist_patterns"] = admission_whitelist_patterns
            __props__.__dict__["cluster_admission_rules"] = cluster_admission_rules
            if default_admission_rule is None and not opts.urn:
                raise TypeError("Missing required property 'default_admission_rule'")
            __props__.__dict__["default_admission_rule"] = default_admission_rule
            __props__.__dict__["description"] = description
            __props__.__dict__["global_policy_evaluation_mode"] = global_policy_evaluation_mode
            __props__.__dict__["project"] = project
        super(Policy, __self__).__init__(
            'gcp:binaryauthorization/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admission_whitelist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]]] = None,
            cluster_admission_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]]] = None,
            default_admission_rule: Optional[pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            global_policy_evaluation_mode: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyAdmissionWhitelistPatternArgs']]]] admission_whitelist_patterns: A whitelist of image patterns to exclude from admission rules. If an
               image's name matches a whitelist pattern, the image's admission
               requests will always be permitted regardless of your admission rules.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyClusterAdmissionRuleArgs']]]] cluster_admission_rules: Per-cluster admission rules. An admission rule specifies either that
               all container images used in a pod creation request must be attested
               to by one or more attestors, that all pod creations will be allowed,
               or that all pod creations will be denied. There can be at most one
               admission rule per cluster spec.
               
               Identifier format: `{{location}}.{{clusterId}}`.
               A location is either a compute zone (e.g. `us-central1-a`) or a region
               (e.g. `us-central1`).
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['PolicyDefaultAdmissionRuleArgs']] default_admission_rule: Default admission rule for a cluster without a per-cluster admission
               rule.
               Structure is documented below.
        :param pulumi.Input[str] description: A descriptive comment.
        :param pulumi.Input[str] global_policy_evaluation_mode: Controls the evaluation of a Google-maintained global admission policy
               for common system-level images. Images not covered by the global
               policy will be subject to the project admission policy.
               Possible values are: `ENABLE`, `DISABLE`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyState.__new__(_PolicyState)

        __props__.__dict__["admission_whitelist_patterns"] = admission_whitelist_patterns
        __props__.__dict__["cluster_admission_rules"] = cluster_admission_rules
        __props__.__dict__["default_admission_rule"] = default_admission_rule
        __props__.__dict__["description"] = description
        __props__.__dict__["global_policy_evaluation_mode"] = global_policy_evaluation_mode
        __props__.__dict__["project"] = project
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="admissionWhitelistPatterns")
    def admission_whitelist_patterns(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyAdmissionWhitelistPattern']]]:
        """
        A whitelist of image patterns to exclude from admission rules. If an
        image's name matches a whitelist pattern, the image's admission
        requests will always be permitted regardless of your admission rules.
        Structure is documented below.
        """
        return pulumi.get(self, "admission_whitelist_patterns")

    @property
    @pulumi.getter(name="clusterAdmissionRules")
    def cluster_admission_rules(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyClusterAdmissionRule']]]:
        """
        Per-cluster admission rules. An admission rule specifies either that
        all container images used in a pod creation request must be attested
        to by one or more attestors, that all pod creations will be allowed,
        or that all pod creations will be denied. There can be at most one
        admission rule per cluster spec.

        Identifier format: `{{location}}.{{clusterId}}`.
        A location is either a compute zone (e.g. `us-central1-a`) or a region
        (e.g. `us-central1`).
        Structure is documented below.
        """
        return pulumi.get(self, "cluster_admission_rules")

    @property
    @pulumi.getter(name="defaultAdmissionRule")
    def default_admission_rule(self) -> pulumi.Output['outputs.PolicyDefaultAdmissionRule']:
        """
        Default admission rule for a cluster without a per-cluster admission
        rule.
        Structure is documented below.
        """
        return pulumi.get(self, "default_admission_rule")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A descriptive comment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="globalPolicyEvaluationMode")
    def global_policy_evaluation_mode(self) -> pulumi.Output[str]:
        """
        Controls the evaluation of a Google-maintained global admission policy
        for common system-level images. Images not covered by the global
        policy will be subject to the project admission policy.
        Possible values are: `ENABLE`, `DISABLE`.
        """
        return pulumi.get(self, "global_policy_evaluation_mode")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

