// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.DiskAsyncReplicationSecondaryDiskArgs;
import java.lang.String;
import java.util.Objects;


public final class DiskAsyncReplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final DiskAsyncReplicationArgs Empty = new DiskAsyncReplicationArgs();

    /**
     * The primary disk (source of replication).
     * 
     */
    @Import(name="primaryDisk", required=true)
    private Output<String> primaryDisk;

    /**
     * @return The primary disk (source of replication).
     * 
     */
    public Output<String> primaryDisk() {
        return this.primaryDisk;
    }

    /**
     * The secondary disk (target of replication). You can specify only one value. Structure is documented below.
     * 
     */
    @Import(name="secondaryDisk", required=true)
    private Output<DiskAsyncReplicationSecondaryDiskArgs> secondaryDisk;

    /**
     * @return The secondary disk (target of replication). You can specify only one value. Structure is documented below.
     * 
     */
    public Output<DiskAsyncReplicationSecondaryDiskArgs> secondaryDisk() {
        return this.secondaryDisk;
    }

    private DiskAsyncReplicationArgs() {}

    private DiskAsyncReplicationArgs(DiskAsyncReplicationArgs $) {
        this.primaryDisk = $.primaryDisk;
        this.secondaryDisk = $.secondaryDisk;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DiskAsyncReplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DiskAsyncReplicationArgs $;

        public Builder() {
            $ = new DiskAsyncReplicationArgs();
        }

        public Builder(DiskAsyncReplicationArgs defaults) {
            $ = new DiskAsyncReplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param primaryDisk The primary disk (source of replication).
         * 
         * @return builder
         * 
         */
        public Builder primaryDisk(Output<String> primaryDisk) {
            $.primaryDisk = primaryDisk;
            return this;
        }

        /**
         * @param primaryDisk The primary disk (source of replication).
         * 
         * @return builder
         * 
         */
        public Builder primaryDisk(String primaryDisk) {
            return primaryDisk(Output.of(primaryDisk));
        }

        /**
         * @param secondaryDisk The secondary disk (target of replication). You can specify only one value. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaryDisk(Output<DiskAsyncReplicationSecondaryDiskArgs> secondaryDisk) {
            $.secondaryDisk = secondaryDisk;
            return this;
        }

        /**
         * @param secondaryDisk The secondary disk (target of replication). You can specify only one value. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaryDisk(DiskAsyncReplicationSecondaryDiskArgs secondaryDisk) {
            return secondaryDisk(Output.of(secondaryDisk));
        }

        public DiskAsyncReplicationArgs build() {
            $.primaryDisk = Objects.requireNonNull($.primaryDisk, "expected parameter 'primaryDisk' to be non-null");
            $.secondaryDisk = Objects.requireNonNull($.secondaryDisk, "expected parameter 'secondaryDisk' to be non-null");
            return $;
        }
    }

}
