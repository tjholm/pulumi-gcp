// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.GetInstanceAdvancedMachineFeature;
import com.pulumi.gcp.compute.outputs.GetInstanceAttachedDisk;
import com.pulumi.gcp.compute.outputs.GetInstanceBootDisk;
import com.pulumi.gcp.compute.outputs.GetInstanceConfidentialInstanceConfig;
import com.pulumi.gcp.compute.outputs.GetInstanceGuestAccelerator;
import com.pulumi.gcp.compute.outputs.GetInstanceNetworkInterface;
import com.pulumi.gcp.compute.outputs.GetInstanceNetworkPerformanceConfig;
import com.pulumi.gcp.compute.outputs.GetInstanceParam;
import com.pulumi.gcp.compute.outputs.GetInstanceReservationAffinity;
import com.pulumi.gcp.compute.outputs.GetInstanceScheduling;
import com.pulumi.gcp.compute.outputs.GetInstanceScratchDisk;
import com.pulumi.gcp.compute.outputs.GetInstanceServiceAccount;
import com.pulumi.gcp.compute.outputs.GetInstanceShieldedInstanceConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetInstanceResult {
    private List<GetInstanceAdvancedMachineFeature> advancedMachineFeatures;
    private Boolean allowStoppingForUpdate;
    /**
     * @return List of disks attached to the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceAttachedDisk> attachedDisks;
    /**
     * @return The boot disk for the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceBootDisk> bootDisks;
    /**
     * @return Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     * 
     */
    private Boolean canIpForward;
    private List<GetInstanceConfidentialInstanceConfig> confidentialInstanceConfigs;
    /**
     * @return The CPU platform used by this instance.
     * 
     */
    private String cpuPlatform;
    /**
     * @return The current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).`,
     * 
     */
    private String currentStatus;
    /**
     * @return Whether deletion protection is enabled on this instance.
     * 
     */
    private Boolean deletionProtection;
    /**
     * @return A brief description of the resource.
     * 
     */
    private String description;
    private String desiredStatus;
    /**
     * @return Whether the instance has virtual displays enabled.
     * 
     */
    private Boolean enableDisplay;
    /**
     * @return List of the type and count of accelerator cards attached to the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceGuestAccelerator> guestAccelerators;
    private String hostname;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The server-assigned unique identifier of this instance.
     * 
     */
    private String instanceId;
    /**
     * @return The unique fingerprint of the labels.
     * 
     */
    private String labelFingerprint;
    /**
     * @return A set of key/value label pairs assigned to the disk.
     * 
     */
    private Map<String,String> labels;
    /**
     * @return The machine type to create.
     * 
     */
    private String machineType;
    /**
     * @return Metadata key/value pairs made available within the instance.
     * 
     */
    private Map<String,String> metadata;
    /**
     * @return The unique fingerprint of the metadata.
     * 
     */
    private String metadataFingerprint;
    private String metadataStartupScript;
    /**
     * @return The minimum CPU platform specified for the VM instance.
     * 
     */
    private String minCpuPlatform;
    private @Nullable String name;
    /**
     * @return The networks attached to the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceNetworkInterface> networkInterfaces;
    /**
     * @return The network performance configuration setting for the instance, if set. Structure is documented below.
     * 
     */
    private List<GetInstanceNetworkPerformanceConfig> networkPerformanceConfigs;
    private List<GetInstanceParam> params;
    private @Nullable String project;
    private List<GetInstanceReservationAffinity> reservationAffinities;
    private List<String> resourcePolicies;
    /**
     * @return The scheduling strategy being used by the instance. Structure is documented below
     * 
     */
    private List<GetInstanceScheduling> schedulings;
    /**
     * @return The scratch disks attached to the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceScratchDisk> scratchDisks;
    /**
     * @return The URI of the created resource.
     * 
     */
    private @Nullable String selfLink;
    /**
     * @return The service account to attach to the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceServiceAccount> serviceAccounts;
    /**
     * @return The shielded vm config being used by the instance. Structure is documented below.
     * 
     */
    private List<GetInstanceShieldedInstanceConfig> shieldedInstanceConfigs;
    /**
     * @return The list of tags attached to the instance.
     * 
     */
    private List<String> tags;
    /**
     * @return The unique fingerprint of the tags.
     * 
     */
    private String tagsFingerprint;
    private @Nullable String zone;

    private GetInstanceResult() {}
    public List<GetInstanceAdvancedMachineFeature> advancedMachineFeatures() {
        return this.advancedMachineFeatures;
    }
    public Boolean allowStoppingForUpdate() {
        return this.allowStoppingForUpdate;
    }
    /**
     * @return List of disks attached to the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceAttachedDisk> attachedDisks() {
        return this.attachedDisks;
    }
    /**
     * @return The boot disk for the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceBootDisk> bootDisks() {
        return this.bootDisks;
    }
    /**
     * @return Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     * 
     */
    public Boolean canIpForward() {
        return this.canIpForward;
    }
    public List<GetInstanceConfidentialInstanceConfig> confidentialInstanceConfigs() {
        return this.confidentialInstanceConfigs;
    }
    /**
     * @return The CPU platform used by this instance.
     * 
     */
    public String cpuPlatform() {
        return this.cpuPlatform;
    }
    /**
     * @return The current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).`,
     * 
     */
    public String currentStatus() {
        return this.currentStatus;
    }
    /**
     * @return Whether deletion protection is enabled on this instance.
     * 
     */
    public Boolean deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * @return A brief description of the resource.
     * 
     */
    public String description() {
        return this.description;
    }
    public String desiredStatus() {
        return this.desiredStatus;
    }
    /**
     * @return Whether the instance has virtual displays enabled.
     * 
     */
    public Boolean enableDisplay() {
        return this.enableDisplay;
    }
    /**
     * @return List of the type and count of accelerator cards attached to the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceGuestAccelerator> guestAccelerators() {
        return this.guestAccelerators;
    }
    public String hostname() {
        return this.hostname;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The server-assigned unique identifier of this instance.
     * 
     */
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return The unique fingerprint of the labels.
     * 
     */
    public String labelFingerprint() {
        return this.labelFingerprint;
    }
    /**
     * @return A set of key/value label pairs assigned to the disk.
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
     * @return Metadata key/value pairs made available within the instance.
     * 
     */
    public Map<String,String> metadata() {
        return this.metadata;
    }
    /**
     * @return The unique fingerprint of the metadata.
     * 
     */
    public String metadataFingerprint() {
        return this.metadataFingerprint;
    }
    public String metadataStartupScript() {
        return this.metadataStartupScript;
    }
    /**
     * @return The minimum CPU platform specified for the VM instance.
     * 
     */
    public String minCpuPlatform() {
        return this.minCpuPlatform;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The networks attached to the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceNetworkInterface> networkInterfaces() {
        return this.networkInterfaces;
    }
    /**
     * @return The network performance configuration setting for the instance, if set. Structure is documented below.
     * 
     */
    public List<GetInstanceNetworkPerformanceConfig> networkPerformanceConfigs() {
        return this.networkPerformanceConfigs;
    }
    public List<GetInstanceParam> params() {
        return this.params;
    }
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }
    public List<GetInstanceReservationAffinity> reservationAffinities() {
        return this.reservationAffinities;
    }
    public List<String> resourcePolicies() {
        return this.resourcePolicies;
    }
    /**
     * @return The scheduling strategy being used by the instance. Structure is documented below
     * 
     */
    public List<GetInstanceScheduling> schedulings() {
        return this.schedulings;
    }
    /**
     * @return The scratch disks attached to the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceScratchDisk> scratchDisks() {
        return this.scratchDisks;
    }
    /**
     * @return The URI of the created resource.
     * 
     */
    public Optional<String> selfLink() {
        return Optional.ofNullable(this.selfLink);
    }
    /**
     * @return The service account to attach to the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceServiceAccount> serviceAccounts() {
        return this.serviceAccounts;
    }
    /**
     * @return The shielded vm config being used by the instance. Structure is documented below.
     * 
     */
    public List<GetInstanceShieldedInstanceConfig> shieldedInstanceConfigs() {
        return this.shieldedInstanceConfigs;
    }
    /**
     * @return The list of tags attached to the instance.
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
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetInstanceAdvancedMachineFeature> advancedMachineFeatures;
        private Boolean allowStoppingForUpdate;
        private List<GetInstanceAttachedDisk> attachedDisks;
        private List<GetInstanceBootDisk> bootDisks;
        private Boolean canIpForward;
        private List<GetInstanceConfidentialInstanceConfig> confidentialInstanceConfigs;
        private String cpuPlatform;
        private String currentStatus;
        private Boolean deletionProtection;
        private String description;
        private String desiredStatus;
        private Boolean enableDisplay;
        private List<GetInstanceGuestAccelerator> guestAccelerators;
        private String hostname;
        private String id;
        private String instanceId;
        private String labelFingerprint;
        private Map<String,String> labels;
        private String machineType;
        private Map<String,String> metadata;
        private String metadataFingerprint;
        private String metadataStartupScript;
        private String minCpuPlatform;
        private @Nullable String name;
        private List<GetInstanceNetworkInterface> networkInterfaces;
        private List<GetInstanceNetworkPerformanceConfig> networkPerformanceConfigs;
        private List<GetInstanceParam> params;
        private @Nullable String project;
        private List<GetInstanceReservationAffinity> reservationAffinities;
        private List<String> resourcePolicies;
        private List<GetInstanceScheduling> schedulings;
        private List<GetInstanceScratchDisk> scratchDisks;
        private @Nullable String selfLink;
        private List<GetInstanceServiceAccount> serviceAccounts;
        private List<GetInstanceShieldedInstanceConfig> shieldedInstanceConfigs;
        private List<String> tags;
        private String tagsFingerprint;
        private @Nullable String zone;
        public Builder() {}
        public Builder(GetInstanceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.advancedMachineFeatures = defaults.advancedMachineFeatures;
    	      this.allowStoppingForUpdate = defaults.allowStoppingForUpdate;
    	      this.attachedDisks = defaults.attachedDisks;
    	      this.bootDisks = defaults.bootDisks;
    	      this.canIpForward = defaults.canIpForward;
    	      this.confidentialInstanceConfigs = defaults.confidentialInstanceConfigs;
    	      this.cpuPlatform = defaults.cpuPlatform;
    	      this.currentStatus = defaults.currentStatus;
    	      this.deletionProtection = defaults.deletionProtection;
    	      this.description = defaults.description;
    	      this.desiredStatus = defaults.desiredStatus;
    	      this.enableDisplay = defaults.enableDisplay;
    	      this.guestAccelerators = defaults.guestAccelerators;
    	      this.hostname = defaults.hostname;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
    	      this.labelFingerprint = defaults.labelFingerprint;
    	      this.labels = defaults.labels;
    	      this.machineType = defaults.machineType;
    	      this.metadata = defaults.metadata;
    	      this.metadataFingerprint = defaults.metadataFingerprint;
    	      this.metadataStartupScript = defaults.metadataStartupScript;
    	      this.minCpuPlatform = defaults.minCpuPlatform;
    	      this.name = defaults.name;
    	      this.networkInterfaces = defaults.networkInterfaces;
    	      this.networkPerformanceConfigs = defaults.networkPerformanceConfigs;
    	      this.params = defaults.params;
    	      this.project = defaults.project;
    	      this.reservationAffinities = defaults.reservationAffinities;
    	      this.resourcePolicies = defaults.resourcePolicies;
    	      this.schedulings = defaults.schedulings;
    	      this.scratchDisks = defaults.scratchDisks;
    	      this.selfLink = defaults.selfLink;
    	      this.serviceAccounts = defaults.serviceAccounts;
    	      this.shieldedInstanceConfigs = defaults.shieldedInstanceConfigs;
    	      this.tags = defaults.tags;
    	      this.tagsFingerprint = defaults.tagsFingerprint;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder advancedMachineFeatures(List<GetInstanceAdvancedMachineFeature> advancedMachineFeatures) {
            this.advancedMachineFeatures = Objects.requireNonNull(advancedMachineFeatures);
            return this;
        }
        public Builder advancedMachineFeatures(GetInstanceAdvancedMachineFeature... advancedMachineFeatures) {
            return advancedMachineFeatures(List.of(advancedMachineFeatures));
        }
        @CustomType.Setter
        public Builder allowStoppingForUpdate(Boolean allowStoppingForUpdate) {
            this.allowStoppingForUpdate = Objects.requireNonNull(allowStoppingForUpdate);
            return this;
        }
        @CustomType.Setter
        public Builder attachedDisks(List<GetInstanceAttachedDisk> attachedDisks) {
            this.attachedDisks = Objects.requireNonNull(attachedDisks);
            return this;
        }
        public Builder attachedDisks(GetInstanceAttachedDisk... attachedDisks) {
            return attachedDisks(List.of(attachedDisks));
        }
        @CustomType.Setter
        public Builder bootDisks(List<GetInstanceBootDisk> bootDisks) {
            this.bootDisks = Objects.requireNonNull(bootDisks);
            return this;
        }
        public Builder bootDisks(GetInstanceBootDisk... bootDisks) {
            return bootDisks(List.of(bootDisks));
        }
        @CustomType.Setter
        public Builder canIpForward(Boolean canIpForward) {
            this.canIpForward = Objects.requireNonNull(canIpForward);
            return this;
        }
        @CustomType.Setter
        public Builder confidentialInstanceConfigs(List<GetInstanceConfidentialInstanceConfig> confidentialInstanceConfigs) {
            this.confidentialInstanceConfigs = Objects.requireNonNull(confidentialInstanceConfigs);
            return this;
        }
        public Builder confidentialInstanceConfigs(GetInstanceConfidentialInstanceConfig... confidentialInstanceConfigs) {
            return confidentialInstanceConfigs(List.of(confidentialInstanceConfigs));
        }
        @CustomType.Setter
        public Builder cpuPlatform(String cpuPlatform) {
            this.cpuPlatform = Objects.requireNonNull(cpuPlatform);
            return this;
        }
        @CustomType.Setter
        public Builder currentStatus(String currentStatus) {
            this.currentStatus = Objects.requireNonNull(currentStatus);
            return this;
        }
        @CustomType.Setter
        public Builder deletionProtection(Boolean deletionProtection) {
            this.deletionProtection = Objects.requireNonNull(deletionProtection);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder desiredStatus(String desiredStatus) {
            this.desiredStatus = Objects.requireNonNull(desiredStatus);
            return this;
        }
        @CustomType.Setter
        public Builder enableDisplay(Boolean enableDisplay) {
            this.enableDisplay = Objects.requireNonNull(enableDisplay);
            return this;
        }
        @CustomType.Setter
        public Builder guestAccelerators(List<GetInstanceGuestAccelerator> guestAccelerators) {
            this.guestAccelerators = Objects.requireNonNull(guestAccelerators);
            return this;
        }
        public Builder guestAccelerators(GetInstanceGuestAccelerator... guestAccelerators) {
            return guestAccelerators(List.of(guestAccelerators));
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            this.hostname = Objects.requireNonNull(hostname);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            this.instanceId = Objects.requireNonNull(instanceId);
            return this;
        }
        @CustomType.Setter
        public Builder labelFingerprint(String labelFingerprint) {
            this.labelFingerprint = Objects.requireNonNull(labelFingerprint);
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            this.labels = Objects.requireNonNull(labels);
            return this;
        }
        @CustomType.Setter
        public Builder machineType(String machineType) {
            this.machineType = Objects.requireNonNull(machineType);
            return this;
        }
        @CustomType.Setter
        public Builder metadata(Map<String,String> metadata) {
            this.metadata = Objects.requireNonNull(metadata);
            return this;
        }
        @CustomType.Setter
        public Builder metadataFingerprint(String metadataFingerprint) {
            this.metadataFingerprint = Objects.requireNonNull(metadataFingerprint);
            return this;
        }
        @CustomType.Setter
        public Builder metadataStartupScript(String metadataStartupScript) {
            this.metadataStartupScript = Objects.requireNonNull(metadataStartupScript);
            return this;
        }
        @CustomType.Setter
        public Builder minCpuPlatform(String minCpuPlatform) {
            this.minCpuPlatform = Objects.requireNonNull(minCpuPlatform);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaces(List<GetInstanceNetworkInterface> networkInterfaces) {
            this.networkInterfaces = Objects.requireNonNull(networkInterfaces);
            return this;
        }
        public Builder networkInterfaces(GetInstanceNetworkInterface... networkInterfaces) {
            return networkInterfaces(List.of(networkInterfaces));
        }
        @CustomType.Setter
        public Builder networkPerformanceConfigs(List<GetInstanceNetworkPerformanceConfig> networkPerformanceConfigs) {
            this.networkPerformanceConfigs = Objects.requireNonNull(networkPerformanceConfigs);
            return this;
        }
        public Builder networkPerformanceConfigs(GetInstanceNetworkPerformanceConfig... networkPerformanceConfigs) {
            return networkPerformanceConfigs(List.of(networkPerformanceConfigs));
        }
        @CustomType.Setter
        public Builder params(List<GetInstanceParam> params) {
            this.params = Objects.requireNonNull(params);
            return this;
        }
        public Builder params(GetInstanceParam... params) {
            return params(List.of(params));
        }
        @CustomType.Setter
        public Builder project(@Nullable String project) {
            this.project = project;
            return this;
        }
        @CustomType.Setter
        public Builder reservationAffinities(List<GetInstanceReservationAffinity> reservationAffinities) {
            this.reservationAffinities = Objects.requireNonNull(reservationAffinities);
            return this;
        }
        public Builder reservationAffinities(GetInstanceReservationAffinity... reservationAffinities) {
            return reservationAffinities(List.of(reservationAffinities));
        }
        @CustomType.Setter
        public Builder resourcePolicies(List<String> resourcePolicies) {
            this.resourcePolicies = Objects.requireNonNull(resourcePolicies);
            return this;
        }
        public Builder resourcePolicies(String... resourcePolicies) {
            return resourcePolicies(List.of(resourcePolicies));
        }
        @CustomType.Setter
        public Builder schedulings(List<GetInstanceScheduling> schedulings) {
            this.schedulings = Objects.requireNonNull(schedulings);
            return this;
        }
        public Builder schedulings(GetInstanceScheduling... schedulings) {
            return schedulings(List.of(schedulings));
        }
        @CustomType.Setter
        public Builder scratchDisks(List<GetInstanceScratchDisk> scratchDisks) {
            this.scratchDisks = Objects.requireNonNull(scratchDisks);
            return this;
        }
        public Builder scratchDisks(GetInstanceScratchDisk... scratchDisks) {
            return scratchDisks(List.of(scratchDisks));
        }
        @CustomType.Setter
        public Builder selfLink(@Nullable String selfLink) {
            this.selfLink = selfLink;
            return this;
        }
        @CustomType.Setter
        public Builder serviceAccounts(List<GetInstanceServiceAccount> serviceAccounts) {
            this.serviceAccounts = Objects.requireNonNull(serviceAccounts);
            return this;
        }
        public Builder serviceAccounts(GetInstanceServiceAccount... serviceAccounts) {
            return serviceAccounts(List.of(serviceAccounts));
        }
        @CustomType.Setter
        public Builder shieldedInstanceConfigs(List<GetInstanceShieldedInstanceConfig> shieldedInstanceConfigs) {
            this.shieldedInstanceConfigs = Objects.requireNonNull(shieldedInstanceConfigs);
            return this;
        }
        public Builder shieldedInstanceConfigs(GetInstanceShieldedInstanceConfig... shieldedInstanceConfigs) {
            return shieldedInstanceConfigs(List.of(shieldedInstanceConfigs));
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder tagsFingerprint(String tagsFingerprint) {
            this.tagsFingerprint = Objects.requireNonNull(tagsFingerprint);
            return this;
        }
        @CustomType.Setter
        public Builder zone(@Nullable String zone) {
            this.zone = zone;
            return this;
        }
        public GetInstanceResult build() {
            final var o = new GetInstanceResult();
            o.advancedMachineFeatures = advancedMachineFeatures;
            o.allowStoppingForUpdate = allowStoppingForUpdate;
            o.attachedDisks = attachedDisks;
            o.bootDisks = bootDisks;
            o.canIpForward = canIpForward;
            o.confidentialInstanceConfigs = confidentialInstanceConfigs;
            o.cpuPlatform = cpuPlatform;
            o.currentStatus = currentStatus;
            o.deletionProtection = deletionProtection;
            o.description = description;
            o.desiredStatus = desiredStatus;
            o.enableDisplay = enableDisplay;
            o.guestAccelerators = guestAccelerators;
            o.hostname = hostname;
            o.id = id;
            o.instanceId = instanceId;
            o.labelFingerprint = labelFingerprint;
            o.labels = labels;
            o.machineType = machineType;
            o.metadata = metadata;
            o.metadataFingerprint = metadataFingerprint;
            o.metadataStartupScript = metadataStartupScript;
            o.minCpuPlatform = minCpuPlatform;
            o.name = name;
            o.networkInterfaces = networkInterfaces;
            o.networkPerformanceConfigs = networkPerformanceConfigs;
            o.params = params;
            o.project = project;
            o.reservationAffinities = reservationAffinities;
            o.resourcePolicies = resourcePolicies;
            o.schedulings = schedulings;
            o.scratchDisks = scratchDisks;
            o.selfLink = selfLink;
            o.serviceAccounts = serviceAccounts;
            o.shieldedInstanceConfigs = shieldedInstanceConfigs;
            o.tags = tags;
            o.tagsFingerprint = tagsFingerprint;
            o.zone = zone;
            return o;
        }
    }
}
