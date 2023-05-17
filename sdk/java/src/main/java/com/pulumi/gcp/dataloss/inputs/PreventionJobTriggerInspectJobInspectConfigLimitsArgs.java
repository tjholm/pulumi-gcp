// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dataloss.inputs.PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArgs;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PreventionJobTriggerInspectJobInspectConfigLimitsArgs extends com.pulumi.resources.ResourceArgs {

    public static final PreventionJobTriggerInspectJobInspectConfigLimitsArgs Empty = new PreventionJobTriggerInspectJobInspectConfigLimitsArgs();

    /**
     * Configuration of findings limit given for specified infoTypes.
     * Structure is documented below.
     * 
     */
    @Import(name="maxFindingsPerInfoTypes")
    private @Nullable Output<List<PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArgs>> maxFindingsPerInfoTypes;

    /**
     * @return Configuration of findings limit given for specified infoTypes.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArgs>>> maxFindingsPerInfoTypes() {
        return Optional.ofNullable(this.maxFindingsPerInfoTypes);
    }

    /**
     * Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
     * 
     */
    @Import(name="maxFindingsPerItem")
    private @Nullable Output<Integer> maxFindingsPerItem;

    /**
     * @return Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
     * 
     */
    public Optional<Output<Integer>> maxFindingsPerItem() {
        return Optional.ofNullable(this.maxFindingsPerItem);
    }

    /**
     * Max number of findings that will be returned per request/job. The maximum returned is 2000.
     * 
     */
    @Import(name="maxFindingsPerRequest")
    private @Nullable Output<Integer> maxFindingsPerRequest;

    /**
     * @return Max number of findings that will be returned per request/job. The maximum returned is 2000.
     * 
     */
    public Optional<Output<Integer>> maxFindingsPerRequest() {
        return Optional.ofNullable(this.maxFindingsPerRequest);
    }

    private PreventionJobTriggerInspectJobInspectConfigLimitsArgs() {}

    private PreventionJobTriggerInspectJobInspectConfigLimitsArgs(PreventionJobTriggerInspectJobInspectConfigLimitsArgs $) {
        this.maxFindingsPerInfoTypes = $.maxFindingsPerInfoTypes;
        this.maxFindingsPerItem = $.maxFindingsPerItem;
        this.maxFindingsPerRequest = $.maxFindingsPerRequest;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PreventionJobTriggerInspectJobInspectConfigLimitsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PreventionJobTriggerInspectJobInspectConfigLimitsArgs $;

        public Builder() {
            $ = new PreventionJobTriggerInspectJobInspectConfigLimitsArgs();
        }

        public Builder(PreventionJobTriggerInspectJobInspectConfigLimitsArgs defaults) {
            $ = new PreventionJobTriggerInspectJobInspectConfigLimitsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param maxFindingsPerInfoTypes Configuration of findings limit given for specified infoTypes.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerInfoTypes(@Nullable Output<List<PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArgs>> maxFindingsPerInfoTypes) {
            $.maxFindingsPerInfoTypes = maxFindingsPerInfoTypes;
            return this;
        }

        /**
         * @param maxFindingsPerInfoTypes Configuration of findings limit given for specified infoTypes.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerInfoTypes(List<PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArgs> maxFindingsPerInfoTypes) {
            return maxFindingsPerInfoTypes(Output.of(maxFindingsPerInfoTypes));
        }

        /**
         * @param maxFindingsPerInfoTypes Configuration of findings limit given for specified infoTypes.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerInfoTypes(PreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeArgs... maxFindingsPerInfoTypes) {
            return maxFindingsPerInfoTypes(List.of(maxFindingsPerInfoTypes));
        }

        /**
         * @param maxFindingsPerItem Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerItem(@Nullable Output<Integer> maxFindingsPerItem) {
            $.maxFindingsPerItem = maxFindingsPerItem;
            return this;
        }

        /**
         * @param maxFindingsPerItem Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerItem(Integer maxFindingsPerItem) {
            return maxFindingsPerItem(Output.of(maxFindingsPerItem));
        }

        /**
         * @param maxFindingsPerRequest Max number of findings that will be returned per request/job. The maximum returned is 2000.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerRequest(@Nullable Output<Integer> maxFindingsPerRequest) {
            $.maxFindingsPerRequest = maxFindingsPerRequest;
            return this;
        }

        /**
         * @param maxFindingsPerRequest Max number of findings that will be returned per request/job. The maximum returned is 2000.
         * 
         * @return builder
         * 
         */
        public Builder maxFindingsPerRequest(Integer maxFindingsPerRequest) {
            return maxFindingsPerRequest(Output.of(maxFindingsPerRequest));
        }

        public PreventionJobTriggerInspectJobInspectConfigLimitsArgs build() {
            return $;
        }
    }

}
