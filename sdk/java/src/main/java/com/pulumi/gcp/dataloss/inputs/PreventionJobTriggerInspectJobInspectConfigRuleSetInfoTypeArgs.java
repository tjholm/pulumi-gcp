// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs extends com.pulumi.resources.ResourceArgs {

    public static final PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs Empty = new PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs();

    /**
     * Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
     * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Version of the information type to use. By default, the version is set to stable.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the information type to use. By default, the version is set to stable.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs() {}

    private PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs(PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs $) {
        this.name = $.name;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs $;

        public Builder() {
            $ = new PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs();
        }

        public Builder(PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs defaults) {
            $ = new PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
         * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
         * at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param version Version of the information type to use. By default, the version is set to stable.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the information type to use. By default, the version is set to stable.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public PreventionJobTriggerInspectJobInspectConfigRuleSetInfoTypeArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
