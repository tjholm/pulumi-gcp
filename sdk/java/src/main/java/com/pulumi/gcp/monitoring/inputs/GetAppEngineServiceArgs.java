// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAppEngineServiceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppEngineServiceArgs Empty = new GetAppEngineServiceArgs();

    /**
     * The ID of the App Engine module underlying this
     * service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
     * 
     */
    @Import(name="moduleId", required=true)
    private Output<String> moduleId;

    /**
     * @return The ID of the App Engine module underlying this
     * service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
     * 
     */
    public Output<String> moduleId() {
        return this.moduleId;
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private GetAppEngineServiceArgs() {}

    private GetAppEngineServiceArgs(GetAppEngineServiceArgs $) {
        this.moduleId = $.moduleId;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppEngineServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppEngineServiceArgs $;

        public Builder() {
            $ = new GetAppEngineServiceArgs();
        }

        public Builder(GetAppEngineServiceArgs defaults) {
            $ = new GetAppEngineServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param moduleId The ID of the App Engine module underlying this
         * service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
         * 
         * @return builder
         * 
         */
        public Builder moduleId(Output<String> moduleId) {
            $.moduleId = moduleId;
            return this;
        }

        /**
         * @param moduleId The ID of the App Engine module underlying this
         * service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
         * 
         * @return builder
         * 
         */
        public Builder moduleId(String moduleId) {
            return moduleId(Output.of(moduleId));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetAppEngineServiceArgs build() {
            $.moduleId = Objects.requireNonNull($.moduleId, "expected parameter 'moduleId' to be non-null");
            return $;
        }
    }

}
