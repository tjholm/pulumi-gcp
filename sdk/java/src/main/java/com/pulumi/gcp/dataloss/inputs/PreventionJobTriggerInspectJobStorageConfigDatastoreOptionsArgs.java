// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dataloss.inputs.PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindArgs;
import com.pulumi.gcp.dataloss.inputs.PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdArgs;
import java.util.Objects;


public final class PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs Empty = new PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs();

    /**
     * A representation of a Datastore kind.
     * Structure is documented below.
     * 
     */
    @Import(name="kind", required=true)
    private Output<PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindArgs> kind;

    /**
     * @return A representation of a Datastore kind.
     * Structure is documented below.
     * 
     */
    public Output<PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindArgs> kind() {
        return this.kind;
    }

    /**
     * Datastore partition ID. A partition ID identifies a grouping of entities. The grouping
     * is always by project and namespace, however the namespace ID may be empty.
     * Structure is documented below.
     * 
     */
    @Import(name="partitionId", required=true)
    private Output<PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdArgs> partitionId;

    /**
     * @return Datastore partition ID. A partition ID identifies a grouping of entities. The grouping
     * is always by project and namespace, however the namespace ID may be empty.
     * Structure is documented below.
     * 
     */
    public Output<PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdArgs> partitionId() {
        return this.partitionId;
    }

    private PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs() {}

    private PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs(PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs $) {
        this.kind = $.kind;
        this.partitionId = $.partitionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs $;

        public Builder() {
            $ = new PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs();
        }

        public Builder(PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs defaults) {
            $ = new PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kind A representation of a Datastore kind.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder kind(Output<PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindArgs> kind) {
            $.kind = kind;
            return this;
        }

        /**
         * @param kind A representation of a Datastore kind.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder kind(PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsKindArgs kind) {
            return kind(Output.of(kind));
        }

        /**
         * @param partitionId Datastore partition ID. A partition ID identifies a grouping of entities. The grouping
         * is always by project and namespace, however the namespace ID may be empty.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder partitionId(Output<PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdArgs> partitionId) {
            $.partitionId = partitionId;
            return this;
        }

        /**
         * @param partitionId Datastore partition ID. A partition ID identifies a grouping of entities. The grouping
         * is always by project and namespace, however the namespace ID may be empty.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder partitionId(PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsPartitionIdArgs partitionId) {
            return partitionId(Output.of(partitionId));
        }

        public PreventionJobTriggerInspectJobStorageConfigDatastoreOptionsArgs build() {
            $.kind = Objects.requireNonNull($.kind, "expected parameter 'kind' to be non-null");
            $.partitionId = Objects.requireNonNull($.partitionId, "expected parameter 'partitionId' to be non-null");
            return $;
        }
    }

}
