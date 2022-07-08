// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.GetInstanceSchedulingNodeAffinity;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInstanceScheduling {
    /**
     * @return Specifies if the instance should be
     * restarted if it was terminated by Compute Engine (not a user).
     * 
     */
    private final Boolean automaticRestart;
    private final Integer minNodeCpus;
    private final List<GetInstanceSchedulingNodeAffinity> nodeAffinities;
    /**
     * @return Describes maintenance behavior for the
     * instance. One of `MIGRATE` or `TERMINATE`, for more info, read
     * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options)
     * 
     */
    private final String onHostMaintenance;
    /**
     * @return Whether the instance is preemptible.
     * 
     */
    private final Boolean preemptible;
    /**
     * @return (Beta) Describe the type of preemptible VM.
     * 
     */
    private final String provisioningModel;

    @CustomType.Constructor
    private GetInstanceScheduling(
        @CustomType.Parameter("automaticRestart") Boolean automaticRestart,
        @CustomType.Parameter("minNodeCpus") Integer minNodeCpus,
        @CustomType.Parameter("nodeAffinities") List<GetInstanceSchedulingNodeAffinity> nodeAffinities,
        @CustomType.Parameter("onHostMaintenance") String onHostMaintenance,
        @CustomType.Parameter("preemptible") Boolean preemptible,
        @CustomType.Parameter("provisioningModel") String provisioningModel) {
        this.automaticRestart = automaticRestart;
        this.minNodeCpus = minNodeCpus;
        this.nodeAffinities = nodeAffinities;
        this.onHostMaintenance = onHostMaintenance;
        this.preemptible = preemptible;
        this.provisioningModel = provisioningModel;
    }

    /**
     * @return Specifies if the instance should be
     * restarted if it was terminated by Compute Engine (not a user).
     * 
     */
    public Boolean automaticRestart() {
        return this.automaticRestart;
    }
    public Integer minNodeCpus() {
        return this.minNodeCpus;
    }
    public List<GetInstanceSchedulingNodeAffinity> nodeAffinities() {
        return this.nodeAffinities;
    }
    /**
     * @return Describes maintenance behavior for the
     * instance. One of `MIGRATE` or `TERMINATE`, for more info, read
     * [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options)
     * 
     */
    public String onHostMaintenance() {
        return this.onHostMaintenance;
    }
    /**
     * @return Whether the instance is preemptible.
     * 
     */
    public Boolean preemptible() {
        return this.preemptible;
    }
    /**
     * @return (Beta) Describe the type of preemptible VM.
     * 
     */
    public String provisioningModel() {
        return this.provisioningModel;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceScheduling defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean automaticRestart;
        private Integer minNodeCpus;
        private List<GetInstanceSchedulingNodeAffinity> nodeAffinities;
        private String onHostMaintenance;
        private Boolean preemptible;
        private String provisioningModel;

        public Builder() {
    	      // Empty
        }

        public Builder(GetInstanceScheduling defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.automaticRestart = defaults.automaticRestart;
    	      this.minNodeCpus = defaults.minNodeCpus;
    	      this.nodeAffinities = defaults.nodeAffinities;
    	      this.onHostMaintenance = defaults.onHostMaintenance;
    	      this.preemptible = defaults.preemptible;
    	      this.provisioningModel = defaults.provisioningModel;
        }

        public Builder automaticRestart(Boolean automaticRestart) {
            this.automaticRestart = Objects.requireNonNull(automaticRestart);
            return this;
        }
        public Builder minNodeCpus(Integer minNodeCpus) {
            this.minNodeCpus = Objects.requireNonNull(minNodeCpus);
            return this;
        }
        public Builder nodeAffinities(List<GetInstanceSchedulingNodeAffinity> nodeAffinities) {
            this.nodeAffinities = Objects.requireNonNull(nodeAffinities);
            return this;
        }
        public Builder nodeAffinities(GetInstanceSchedulingNodeAffinity... nodeAffinities) {
            return nodeAffinities(List.of(nodeAffinities));
        }
        public Builder onHostMaintenance(String onHostMaintenance) {
            this.onHostMaintenance = Objects.requireNonNull(onHostMaintenance);
            return this;
        }
        public Builder preemptible(Boolean preemptible) {
            this.preemptible = Objects.requireNonNull(preemptible);
            return this;
        }
        public Builder provisioningModel(String provisioningModel) {
            this.provisioningModel = Objects.requireNonNull(provisioningModel);
            return this;
        }        public GetInstanceScheduling build() {
            return new GetInstanceScheduling(automaticRestart, minNodeCpus, nodeAffinities, onHostMaintenance, preemptible, provisioningModel);
        }
    }
}
