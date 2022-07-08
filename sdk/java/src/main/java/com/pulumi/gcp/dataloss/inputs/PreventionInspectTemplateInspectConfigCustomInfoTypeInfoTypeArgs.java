// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataloss.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs extends com.pulumi.resources.ResourceArgs {

    public static final PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs Empty = new PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs();

    /**
     * Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
     * or `projects/project-id/storedInfoTypes/432452342`.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
     * or `projects/project-id/storedInfoTypes/432452342`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs() {}

    private PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs(PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs $;

        public Builder() {
            $ = new PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs();
        }

        public Builder(PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs defaults) {
            $ = new PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
         * or `projects/project-id/storedInfoTypes/432452342`.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
         * or `projects/project-id/storedInfoTypes/432452342`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
