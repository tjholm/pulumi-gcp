// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateAdvancedMachineFeature;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateConfidentialInstanceConfig;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateDisk;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateGuestAccelerator;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateNetworkInterface;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateNetworkPerformanceConfig;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateReservationAffinity;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateScheduling;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateServiceAccount;
import com.pulumi.gcp.compute.outputs.GetInstanceTemplateShieldedInstanceConfig;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceTemplateResult {
    private List<GetInstanceTemplateAdvancedMachineFeature> advancedMachineFeatures;
    /**
     * @return Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     * 
     */
    private Boolean canIpForward;
    /**
     * @return Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
     * 
     */
    private List<GetInstanceTemplateConfidentialInstanceConfig> confidentialInstanceConfigs;
    /**
     * @return A brief description of this resource.
     * 
     */
    private String description;
    /**
     * @return Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     * 
     */
    private List<GetInstanceTemplateDisk> disks;
    private Map<String,String> effectiveLabels;
    /**
     * @return Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
     * 
     */
    private Boolean enableDisplay;
    private @Nullable String filter;
    /**
     * @return List of the type and count of accelerator cards attached to the instance. Structure documented below.
     * 
     */
    private List<GetInstanceTemplateGuestAccelerator> guestAccelerators;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A brief description to use for instances
     * created from this template.
     * 
     */
    private String instanceDescription;
    /**
     * @return (Optional) A set of ket/value label pairs to assign to disk created from
     * this template
     * 
     */
    private Map<String,String> labels;
    /**
     * @return The machine type to create.
     * 
     */
    private String machineType;
    /**
     * @return Metadata key/value pairs to make available from
     * within instances created from this template.
     * 
     */
    private Map<String,Object> metadata;
    /**
     * @return The unique fingerprint of the metadata.
     * 
     */
    private String metadataFingerprint;
    /**
     * @return An alternative to using the
     * startup-script metadata key, mostly to match the compute_instance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     * 
     */
    private String metadataStartupScript;
    /**
     * @return Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     * 
     */
    private String minCpuPlatform;
    private @Nullable Boolean mostRecent;
    /**
     * @return The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     * 
     */
    private @Nullable String name;
    /**
     * @return Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     * 
     */
    private String namePrefix;
    /**
     * @return The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.  s
     * 
     */
    private List<GetInstanceTemplateNetworkInterface> networkInterfaces;
    /**
     * @return The network performance configuration setting
     * for the instance, if set. Structure is documented below.
     * 
     */
    private List<GetInstanceTemplateNetworkPerformanceConfig> networkPerformanceConfigs;
    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    private @Nullable String project;
    private Map<String,String> pulumiLabels;
    /**
     * @return An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     * 
     */
    private String region;
    private List<GetInstanceTemplateReservationAffinity> reservationAffinities;
    private Map<String,String> resourceManagerTags;
    /**
     * @return (Optional) -- A list of short names of resource policies to attach to this disk for automatic snapshot creations. Currently a max of 1 resource policy is supported.
     * 
     */
    private List<String> resourcePolicies;
    /**
     * @return The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     * 
     */
    private List<GetInstanceTemplateScheduling> schedulings;
    /**
     * @return The URI of the created resource.
     * 
     */
    private String selfLink;
    /**
     * @return A special URI of the created resource that uniquely identifies this instance template with the following format: `projects/{{project}}/global/instanceTemplates/{{name}}?uniqueId={{uniqueId}}`
     * Referencing an instance template via this attribute prevents Time of Check to Time of Use attacks when the instance template resides in a shared/untrusted environment.
     * 
     */
    private @Nullable String selfLinkUnique;
    /**
     * @return Service account to attach to the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceTemplateServiceAccount> serviceAccounts;
    /**
     * @return Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     * 
     */
    private List<GetInstanceTemplateShieldedInstanceConfig> shieldedInstanceConfigs;
    /**
     * @return Tags to attach to the instance.
     * 
     */
    private List<String> tags;
    /**
     * @return The unique fingerprint of the tags.
     * 
     */
    private String tagsFingerprint;

    private GetInstanceTemplateResult() {}
    public List<GetInstanceTemplateAdvancedMachineFeature> advancedMachineFeatures() {
        return this.advancedMachineFeatures;
    }
    /**
     * @return Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     * 
     */
    public Boolean canIpForward() {
        return this.canIpForward;
    }
    /**
     * @return Enable [Confidential Mode](https://cloud.google.com/compute/confidential-vm/docs/about-cvm) on this VM. Structure is documented below
     * 
     */
    public List<GetInstanceTemplateConfidentialInstanceConfig> confidentialInstanceConfigs() {
        return this.confidentialInstanceConfigs;
    }
    /**
     * @return A brief description of this resource.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     * 
     */
    public List<GetInstanceTemplateDisk> disks() {
        return this.disks;
    }
    public Map<String,String> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * @return Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allow_stopping_for_update` must be set to true in order to update this field.
     * 
     */
    public Boolean enableDisplay() {
        return this.enableDisplay;
    }
    public Optional<String> filter() {
        return Optional.ofNullable(this.filter);
    }
    /**
     * @return List of the type and count of accelerator cards attached to the instance. Structure documented below.
     * 
     */
    public List<GetInstanceTemplateGuestAccelerator> guestAccelerators() {
        return this.guestAccelerators;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A brief description to use for instances
     * created from this template.
     * 
     */
    public String instanceDescription() {
        return this.instanceDescription;
    }
    /**
     * @return (Optional) A set of ket/value label pairs to assign to disk created from
     * this template
     * 
     */
    public Map<String,String> labels() {
        return this.labels;
    }
    /**
     * @return The machine type to create.
     * 
     */
    public String machineType() {
        return this.machineType;
    }
    /**
     * @return Metadata key/value pairs to make available from
     * within instances created from this template.
     * 
     */
    public Map<String,Object> metadata() {
        return this.metadata;
    }
    /**
     * @return The unique fingerprint of the metadata.
     * 
     */
    public String metadataFingerprint() {
        return this.metadataFingerprint;
    }
    /**
     * @return An alternative to using the
     * startup-script metadata key, mostly to match the compute_instance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     * 
     */
    public String metadataStartupScript() {
        return this.metadataStartupScript;
    }
    /**
     * @return Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     * 
     */
    public String minCpuPlatform() {
        return this.minCpuPlatform;
    }
    public Optional<Boolean> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }
    /**
     * @return The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     * 
     */
    public String namePrefix() {
        return this.namePrefix;
    }
    /**
     * @return The URL of the network attachment that this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.  s
     * 
     */
    public List<GetInstanceTemplateNetworkInterface> networkInterfaces() {
        return this.networkInterfaces;
    }
    /**
     * @return The network performance configuration setting
     * for the instance, if set. Structure is documented below.
     * 
     */
    public List<GetInstanceTemplateNetworkPerformanceConfig> networkPerformanceConfigs() {
        return this.networkPerformanceConfigs;
    }
    /**
     * @return The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }
    public Map<String,String> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * @return An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     * 
     */
    public String region() {
        return this.region;
    }
    public List<GetInstanceTemplateReservationAffinity> reservationAffinities() {
        return this.reservationAffinities;
    }
    public Map<String,String> resourceManagerTags() {
        return this.resourceManagerTags;
    }
    /**
     * @return (Optional) -- A list of short names of resource policies to attach to this disk for automatic snapshot creations. Currently a max of 1 resource policy is supported.
     * 
     */
    public List<String> resourcePolicies() {
        return this.resourcePolicies;
    }
    /**
     * @return The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     * 
     */
    public List<GetInstanceTemplateScheduling> schedulings() {
        return this.schedulings;
    }
    /**
     * @return The URI of the created resource.
     * 
     */
    public String selfLink() {
        return this.selfLink;
    }
    /**
     * @return A special URI of the created resource that uniquely identifies this instance template with the following format: `projects/{{project}}/global/instanceTemplates/{{name}}?uniqueId={{uniqueId}}`
     * Referencing an instance template via this attribute prevents Time of Check to Time of Use attacks when the instance template resides in a shared/untrusted environment.
     * 
     */
    public Optional<String> selfLinkUnique() {
        return Optional.ofNullable(this.selfLinkUnique);
    }
    /**
     * @return Service account to attach to the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceTemplateServiceAccount> serviceAccounts() {
        return this.serviceAccounts;
    }
    /**
     * @return Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shielded_instance_config` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     * 
     */
    public List<GetInstanceTemplateShieldedInstanceConfig> shieldedInstanceConfigs() {
        return this.shieldedInstanceConfigs;
    }
    /**
     * @return Tags to attach to the instance.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return The unique fingerprint of the tags.
     * 
     */
    public String tagsFingerprint() {
        return this.tagsFingerprint;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetInstanceTemplateAdvancedMachineFeature> advancedMachineFeatures;
        private Boolean canIpForward;
        private List<GetInstanceTemplateConfidentialInstanceConfig> confidentialInstanceConfigs;
        private String description;
        private List<GetInstanceTemplateDisk> disks;
        private Map<String,String> effectiveLabels;
        private Boolean enableDisplay;
        private @Nullable String filter;
        private List<GetInstanceTemplateGuestAccelerator> guestAccelerators;
        private String id;
        private String instanceDescription;
        private Map<String,String> labels;
        private String machineType;
        private Map<String,Object> metadata;
        private String metadataFingerprint;
        private String metadataStartupScript;
        private String minCpuPlatform;
        private @Nullable Boolean mostRecent;
        private @Nullable String name;
        private String namePrefix;
        private List<GetInstanceTemplateNetworkInterface> networkInterfaces;
        private List<GetInstanceTemplateNetworkPerformanceConfig> networkPerformanceConfigs;
        private @Nullable String project;
        private Map<String,String> pulumiLabels;
        private String region;
        private List<GetInstanceTemplateReservationAffinity> reservationAffinities;
        private Map<String,String> resourceManagerTags;
        private List<String> resourcePolicies;
        private List<GetInstanceTemplateScheduling> schedulings;
        private String selfLink;
        private @Nullable String selfLinkUnique;
        private List<GetInstanceTemplateServiceAccount> serviceAccounts;
        private List<GetInstanceTemplateShieldedInstanceConfig> shieldedInstanceConfigs;
        private List<String> tags;
        private String tagsFingerprint;
        public Builder() {}
        public Builder(GetInstanceTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.advancedMachineFeatures = defaults.advancedMachineFeatures;
    	      this.canIpForward = defaults.canIpForward;
    	      this.confidentialInstanceConfigs = defaults.confidentialInstanceConfigs;
    	      this.description = defaults.description;
    	      this.disks = defaults.disks;
    	      this.effectiveLabels = defaults.effectiveLabels;
    	      this.enableDisplay = defaults.enableDisplay;
    	      this.filter = defaults.filter;
    	      this.guestAccelerators = defaults.guestAccelerators;
    	      this.id = defaults.id;
    	      this.instanceDescription = defaults.instanceDescription;
    	      this.labels = defaults.labels;
    	      this.machineType = defaults.machineType;
    	      this.metadata = defaults.metadata;
    	      this.metadataFingerprint = defaults.metadataFingerprint;
    	      this.metadataStartupScript = defaults.metadataStartupScript;
    	      this.minCpuPlatform = defaults.minCpuPlatform;
    	      this.mostRecent = defaults.mostRecent;
    	      this.name = defaults.name;
    	      this.namePrefix = defaults.namePrefix;
    	      this.networkInterfaces = defaults.networkInterfaces;
    	      this.networkPerformanceConfigs = defaults.networkPerformanceConfigs;
    	      this.project = defaults.project;
    	      this.pulumiLabels = defaults.pulumiLabels;
    	      this.region = defaults.region;
    	      this.reservationAffinities = defaults.reservationAffinities;
    	      this.resourceManagerTags = defaults.resourceManagerTags;
    	      this.resourcePolicies = defaults.resourcePolicies;
    	      this.schedulings = defaults.schedulings;
    	      this.selfLink = defaults.selfLink;
    	      this.selfLinkUnique = defaults.selfLinkUnique;
    	      this.serviceAccounts = defaults.serviceAccounts;
    	      this.shieldedInstanceConfigs = defaults.shieldedInstanceConfigs;
    	      this.tags = defaults.tags;
    	      this.tagsFingerprint = defaults.tagsFingerprint;
        }

        @CustomType.Setter
        public Builder advancedMachineFeatures(List<GetInstanceTemplateAdvancedMachineFeature> advancedMachineFeatures) {
            if (advancedMachineFeatures == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "advancedMachineFeatures");
            }
            this.advancedMachineFeatures = advancedMachineFeatures;
            return this;
        }
        public Builder advancedMachineFeatures(GetInstanceTemplateAdvancedMachineFeature... advancedMachineFeatures) {
            return advancedMachineFeatures(List.of(advancedMachineFeatures));
        }
        @CustomType.Setter
        public Builder canIpForward(Boolean canIpForward) {
            if (canIpForward == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "canIpForward");
            }
            this.canIpForward = canIpForward;
            return this;
        }
        @CustomType.Setter
        public Builder confidentialInstanceConfigs(List<GetInstanceTemplateConfidentialInstanceConfig> confidentialInstanceConfigs) {
            if (confidentialInstanceConfigs == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "confidentialInstanceConfigs");
            }
            this.confidentialInstanceConfigs = confidentialInstanceConfigs;
            return this;
        }
        public Builder confidentialInstanceConfigs(GetInstanceTemplateConfidentialInstanceConfig... confidentialInstanceConfigs) {
            return confidentialInstanceConfigs(List.of(confidentialInstanceConfigs));
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder disks(List<GetInstanceTemplateDisk> disks) {
            if (disks == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "disks");
            }
            this.disks = disks;
            return this;
        }
        public Builder disks(GetInstanceTemplateDisk... disks) {
            return disks(List.of(disks));
        }
        @CustomType.Setter
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            if (effectiveLabels == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "effectiveLabels");
            }
            this.effectiveLabels = effectiveLabels;
            return this;
        }
        @CustomType.Setter
        public Builder enableDisplay(Boolean enableDisplay) {
            if (enableDisplay == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "enableDisplay");
            }
            this.enableDisplay = enableDisplay;
            return this;
        }
        @CustomType.Setter
        public Builder filter(@Nullable String filter) {

            this.filter = filter;
            return this;
        }
        @CustomType.Setter
        public Builder guestAccelerators(List<GetInstanceTemplateGuestAccelerator> guestAccelerators) {
            if (guestAccelerators == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "guestAccelerators");
            }
            this.guestAccelerators = guestAccelerators;
            return this;
        }
        public Builder guestAccelerators(GetInstanceTemplateGuestAccelerator... guestAccelerators) {
            return guestAccelerators(List.of(guestAccelerators));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder instanceDescription(String instanceDescription) {
            if (instanceDescription == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "instanceDescription");
            }
            this.instanceDescription = instanceDescription;
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "labels");
            }
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder machineType(String machineType) {
            if (machineType == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "machineType");
            }
            this.machineType = machineType;
            return this;
        }
        @CustomType.Setter
        public Builder metadata(Map<String,Object> metadata) {
            if (metadata == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "metadata");
            }
            this.metadata = metadata;
            return this;
        }
        @CustomType.Setter
        public Builder metadataFingerprint(String metadataFingerprint) {
            if (metadataFingerprint == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "metadataFingerprint");
            }
            this.metadataFingerprint = metadataFingerprint;
            return this;
        }
        @CustomType.Setter
        public Builder metadataStartupScript(String metadataStartupScript) {
            if (metadataStartupScript == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "metadataStartupScript");
            }
            this.metadataStartupScript = metadataStartupScript;
            return this;
        }
        @CustomType.Setter
        public Builder minCpuPlatform(String minCpuPlatform) {
            if (minCpuPlatform == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "minCpuPlatform");
            }
            this.minCpuPlatform = minCpuPlatform;
            return this;
        }
        @CustomType.Setter
        public Builder mostRecent(@Nullable Boolean mostRecent) {

            this.mostRecent = mostRecent;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder namePrefix(String namePrefix) {
            if (namePrefix == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "namePrefix");
            }
            this.namePrefix = namePrefix;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaces(List<GetInstanceTemplateNetworkInterface> networkInterfaces) {
            if (networkInterfaces == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "networkInterfaces");
            }
            this.networkInterfaces = networkInterfaces;
            return this;
        }
        public Builder networkInterfaces(GetInstanceTemplateNetworkInterface... networkInterfaces) {
            return networkInterfaces(List.of(networkInterfaces));
        }
        @CustomType.Setter
        public Builder networkPerformanceConfigs(List<GetInstanceTemplateNetworkPerformanceConfig> networkPerformanceConfigs) {
            if (networkPerformanceConfigs == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "networkPerformanceConfigs");
            }
            this.networkPerformanceConfigs = networkPerformanceConfigs;
            return this;
        }
        public Builder networkPerformanceConfigs(GetInstanceTemplateNetworkPerformanceConfig... networkPerformanceConfigs) {
            return networkPerformanceConfigs(List.of(networkPerformanceConfigs));
        }
        @CustomType.Setter
        public Builder project(@Nullable String project) {

            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            if (pulumiLabels == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "pulumiLabels");
            }
            this.pulumiLabels = pulumiLabels;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder reservationAffinities(List<GetInstanceTemplateReservationAffinity> reservationAffinities) {
            if (reservationAffinities == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "reservationAffinities");
            }
            this.reservationAffinities = reservationAffinities;
            return this;
        }
        public Builder reservationAffinities(GetInstanceTemplateReservationAffinity... reservationAffinities) {
            return reservationAffinities(List.of(reservationAffinities));
        }
        @CustomType.Setter
        public Builder resourceManagerTags(Map<String,String> resourceManagerTags) {
            if (resourceManagerTags == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "resourceManagerTags");
            }
            this.resourceManagerTags = resourceManagerTags;
            return this;
        }
        @CustomType.Setter
        public Builder resourcePolicies(List<String> resourcePolicies) {
            if (resourcePolicies == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "resourcePolicies");
            }
            this.resourcePolicies = resourcePolicies;
            return this;
        }
        public Builder resourcePolicies(String... resourcePolicies) {
            return resourcePolicies(List.of(resourcePolicies));
        }
        @CustomType.Setter
        public Builder schedulings(List<GetInstanceTemplateScheduling> schedulings) {
            if (schedulings == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "schedulings");
            }
            this.schedulings = schedulings;
            return this;
        }
        public Builder schedulings(GetInstanceTemplateScheduling... schedulings) {
            return schedulings(List.of(schedulings));
        }
        @CustomType.Setter
        public Builder selfLink(String selfLink) {
            if (selfLink == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "selfLink");
            }
            this.selfLink = selfLink;
            return this;
        }
        @CustomType.Setter
        public Builder selfLinkUnique(@Nullable String selfLinkUnique) {

            this.selfLinkUnique = selfLinkUnique;
            return this;
        }
        @CustomType.Setter
        public Builder serviceAccounts(List<GetInstanceTemplateServiceAccount> serviceAccounts) {
            if (serviceAccounts == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "serviceAccounts");
            }
            this.serviceAccounts = serviceAccounts;
            return this;
        }
        public Builder serviceAccounts(GetInstanceTemplateServiceAccount... serviceAccounts) {
            return serviceAccounts(List.of(serviceAccounts));
        }
        @CustomType.Setter
        public Builder shieldedInstanceConfigs(List<GetInstanceTemplateShieldedInstanceConfig> shieldedInstanceConfigs) {
            if (shieldedInstanceConfigs == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "shieldedInstanceConfigs");
            }
            this.shieldedInstanceConfigs = shieldedInstanceConfigs;
            return this;
        }
        public Builder shieldedInstanceConfigs(GetInstanceTemplateShieldedInstanceConfig... shieldedInstanceConfigs) {
            return shieldedInstanceConfigs(List.of(shieldedInstanceConfigs));
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "tags");
            }
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder tagsFingerprint(String tagsFingerprint) {
            if (tagsFingerprint == null) {
              throw new MissingRequiredPropertyException("GetInstanceTemplateResult", "tagsFingerprint");
            }
            this.tagsFingerprint = tagsFingerprint;
            return this;
        }
        public GetInstanceTemplateResult build() {
            final var _resultValue = new GetInstanceTemplateResult();
            _resultValue.advancedMachineFeatures = advancedMachineFeatures;
            _resultValue.canIpForward = canIpForward;
            _resultValue.confidentialInstanceConfigs = confidentialInstanceConfigs;
            _resultValue.description = description;
            _resultValue.disks = disks;
            _resultValue.effectiveLabels = effectiveLabels;
            _resultValue.enableDisplay = enableDisplay;
            _resultValue.filter = filter;
            _resultValue.guestAccelerators = guestAccelerators;
            _resultValue.id = id;
            _resultValue.instanceDescription = instanceDescription;
            _resultValue.labels = labels;
            _resultValue.machineType = machineType;
            _resultValue.metadata = metadata;
            _resultValue.metadataFingerprint = metadataFingerprint;
            _resultValue.metadataStartupScript = metadataStartupScript;
            _resultValue.minCpuPlatform = minCpuPlatform;
            _resultValue.mostRecent = mostRecent;
            _resultValue.name = name;
            _resultValue.namePrefix = namePrefix;
            _resultValue.networkInterfaces = networkInterfaces;
            _resultValue.networkPerformanceConfigs = networkPerformanceConfigs;
            _resultValue.project = project;
            _resultValue.pulumiLabels = pulumiLabels;
            _resultValue.region = region;
            _resultValue.reservationAffinities = reservationAffinities;
            _resultValue.resourceManagerTags = resourceManagerTags;
            _resultValue.resourcePolicies = resourcePolicies;
            _resultValue.schedulings = schedulings;
            _resultValue.selfLink = selfLink;
            _resultValue.selfLinkUnique = selfLinkUnique;
            _resultValue.serviceAccounts = serviceAccounts;
            _resultValue.shieldedInstanceConfigs = shieldedInstanceConfigs;
            _resultValue.tags = tags;
            _resultValue.tagsFingerprint = tagsFingerprint;
            return _resultValue;
        }
    }
}
