// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceTemplatePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceTemplatePlainArgs Empty = new GetInstanceTemplatePlainArgs();

    /**
     * A filter to retrieve the instance templates.
     * See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
     * If multiple instance templates match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
     * 
     */
    @Import(name="filter")
    private @Nullable String filter;

    /**
     * @return A filter to retrieve the instance templates.
     * See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
     * If multiple instance templates match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
     * 
     */
    public Optional<String> filter() {
        return Optional.ofNullable(this.filter);
    }

    /**
     * If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name` or `filter` must be provided.
     * 
     */
    @Import(name="mostRecent")
    private @Nullable Boolean mostRecent;

    /**
     * @return If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name` or `filter` must be provided.
     * 
     */
    public Optional<Boolean> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }

    /**
     * The name of the instance template. One of `name` or `filter` must be provided.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the instance template. One of `name` or `filter` must be provided.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If `project` is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable String project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If `project` is not provided, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    private GetInstanceTemplatePlainArgs() {}

    private GetInstanceTemplatePlainArgs(GetInstanceTemplatePlainArgs $) {
        this.filter = $.filter;
        this.mostRecent = $.mostRecent;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceTemplatePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceTemplatePlainArgs $;

        public Builder() {
            $ = new GetInstanceTemplatePlainArgs();
        }

        public Builder(GetInstanceTemplatePlainArgs defaults) {
            $ = new GetInstanceTemplatePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filter A filter to retrieve the instance templates.
         * See [gcloud topic filters](https://cloud.google.com/sdk/gcloud/reference/topic/filters) for reference.
         * If multiple instance templates match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
         * 
         * @return builder
         * 
         */
        public Builder filter(@Nullable String filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param mostRecent If `filter` is provided, ensures the most recent template is returned when multiple instance templates match. One of `name` or `filter` must be provided.
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(@Nullable Boolean mostRecent) {
            $.mostRecent = mostRecent;
            return this;
        }

        /**
         * @param name The name of the instance template. One of `name` or `filter` must be provided.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If `project` is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        public GetInstanceTemplatePlainArgs build() {
            return $;
        }
    }

}
