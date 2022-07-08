// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class InstanceAcceleratorConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceAcceleratorConfigArgs Empty = new InstanceAcceleratorConfigArgs();

    /**
     * Count of cores of this accelerator.
     * 
     */
    @Import(name="coreCount", required=true)
    private Output<Integer> coreCount;

    /**
     * @return Count of cores of this accelerator.
     * 
     */
    public Output<Integer> coreCount() {
        return this.coreCount;
    }

    /**
     * Type of this accelerator.
     * Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of this accelerator.
     * Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private InstanceAcceleratorConfigArgs() {}

    private InstanceAcceleratorConfigArgs(InstanceAcceleratorConfigArgs $) {
        this.coreCount = $.coreCount;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceAcceleratorConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceAcceleratorConfigArgs $;

        public Builder() {
            $ = new InstanceAcceleratorConfigArgs();
        }

        public Builder(InstanceAcceleratorConfigArgs defaults) {
            $ = new InstanceAcceleratorConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param coreCount Count of cores of this accelerator.
         * 
         * @return builder
         * 
         */
        public Builder coreCount(Output<Integer> coreCount) {
            $.coreCount = coreCount;
            return this;
        }

        /**
         * @param coreCount Count of cores of this accelerator.
         * 
         * @return builder
         * 
         */
        public Builder coreCount(Integer coreCount) {
            return coreCount(Output.of(coreCount));
        }

        /**
         * @param type Type of this accelerator.
         * Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of this accelerator.
         * Possible values are `ACCELERATOR_TYPE_UNSPECIFIED`, `NVIDIA_TESLA_K80`, `NVIDIA_TESLA_P100`, `NVIDIA_TESLA_V100`, `NVIDIA_TESLA_P4`, `NVIDIA_TESLA_T4`, `NVIDIA_TESLA_T4_VWS`, `NVIDIA_TESLA_P100_VWS`, `NVIDIA_TESLA_P4_VWS`, `NVIDIA_TESLA_A100`, `TPU_V2`, and `TPU_V3`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public InstanceAcceleratorConfigArgs build() {
            $.coreCount = Objects.requireNonNull($.coreCount, "expected parameter 'coreCount' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
