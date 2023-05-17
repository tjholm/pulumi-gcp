// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class RegionDiskAsyncPrimaryDiskArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegionDiskAsyncPrimaryDiskArgs Empty = new RegionDiskAsyncPrimaryDiskArgs();

    /**
     * Primary disk for asynchronous disk replication.
     * 
     */
    @Import(name="disk", required=true)
    private Output<String> disk;

    /**
     * @return Primary disk for asynchronous disk replication.
     * 
     */
    public Output<String> disk() {
        return this.disk;
    }

    private RegionDiskAsyncPrimaryDiskArgs() {}

    private RegionDiskAsyncPrimaryDiskArgs(RegionDiskAsyncPrimaryDiskArgs $) {
        this.disk = $.disk;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionDiskAsyncPrimaryDiskArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionDiskAsyncPrimaryDiskArgs $;

        public Builder() {
            $ = new RegionDiskAsyncPrimaryDiskArgs();
        }

        public Builder(RegionDiskAsyncPrimaryDiskArgs defaults) {
            $ = new RegionDiskAsyncPrimaryDiskArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disk Primary disk for asynchronous disk replication.
         * 
         * @return builder
         * 
         */
        public Builder disk(Output<String> disk) {
            $.disk = disk;
            return this;
        }

        /**
         * @param disk Primary disk for asynchronous disk replication.
         * 
         * @return builder
         * 
         */
        public Builder disk(String disk) {
            return disk(Output.of(disk));
        }

        public RegionDiskAsyncPrimaryDiskArgs build() {
            $.disk = Objects.requireNonNull($.disk, "expected parameter 'disk' to be non-null");
            return $;
        }
    }

}
