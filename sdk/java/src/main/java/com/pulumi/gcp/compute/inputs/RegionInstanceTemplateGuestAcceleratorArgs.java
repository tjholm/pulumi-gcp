// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class RegionInstanceTemplateGuestAcceleratorArgs extends com.pulumi.resources.ResourceArgs {

    public static final RegionInstanceTemplateGuestAcceleratorArgs Empty = new RegionInstanceTemplateGuestAcceleratorArgs();

    /**
     * The number of the guest accelerator cards exposed to this instance.
     * 
     */
    @Import(name="count", required=true)
    private Output<Integer> count;

    /**
     * @return The number of the guest accelerator cards exposed to this instance.
     * 
     */
    public Output<Integer> count() {
        return this.count;
    }

    /**
     * The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private RegionInstanceTemplateGuestAcceleratorArgs() {}

    private RegionInstanceTemplateGuestAcceleratorArgs(RegionInstanceTemplateGuestAcceleratorArgs $) {
        this.count = $.count;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionInstanceTemplateGuestAcceleratorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionInstanceTemplateGuestAcceleratorArgs $;

        public Builder() {
            $ = new RegionInstanceTemplateGuestAcceleratorArgs();
        }

        public Builder(RegionInstanceTemplateGuestAcceleratorArgs defaults) {
            $ = new RegionInstanceTemplateGuestAcceleratorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param count The number of the guest accelerator cards exposed to this instance.
         * 
         * @return builder
         * 
         */
        public Builder count(Output<Integer> count) {
            $.count = count;
            return this;
        }

        /**
         * @param count The number of the guest accelerator cards exposed to this instance.
         * 
         * @return builder
         * 
         */
        public Builder count(Integer count) {
            return count(Output.of(count));
        }

        /**
         * @param type The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public RegionInstanceTemplateGuestAcceleratorArgs build() {
            if ($.count == null) {
                throw new MissingRequiredPropertyException("RegionInstanceTemplateGuestAcceleratorArgs", "count");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("RegionInstanceTemplateGuestAcceleratorArgs", "type");
            }
            return $;
        }
    }

}
