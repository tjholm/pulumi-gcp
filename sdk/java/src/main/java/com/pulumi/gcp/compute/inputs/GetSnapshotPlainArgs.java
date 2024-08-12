// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSnapshotPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSnapshotPlainArgs Empty = new GetSnapshotPlainArgs();

    /**
     * A filter to retrieve the compute snapshot.
     * See [API filter parameter documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots/list#body.QUERY_PARAMETERS.filter) for reference.
     * If multiple compute snapshot match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
     * If you want to use a regular expression, use the `eq` (equal) or `ne` (not equal) operator against a single un-parenthesized expression with or without quotes or against multiple parenthesized expressions. Example `sourceDisk eq &#39;.*(.*&#47;data-disk$).*&#39;`. More details for golang Snapshots list call filters [here](https://pkg.go.dev/google.golang.org/api/compute/v1#SnapshotsListCall.Filter).
     * 
     */
    @Import(name="filter")
    private @Nullable String filter;

    /**
     * @return A filter to retrieve the compute snapshot.
     * See [API filter parameter documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots/list#body.QUERY_PARAMETERS.filter) for reference.
     * If multiple compute snapshot match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
     * If you want to use a regular expression, use the `eq` (equal) or `ne` (not equal) operator against a single un-parenthesized expression with or without quotes or against multiple parenthesized expressions. Example `sourceDisk eq &#39;.*(.*&#47;data-disk$).*&#39;`. More details for golang Snapshots list call filters [here](https://pkg.go.dev/google.golang.org/api/compute/v1#SnapshotsListCall.Filter).
     * 
     */
    public Optional<String> filter() {
        return Optional.ofNullable(this.filter);
    }

    /**
     * If `filter` is provided, ensures the most recent snapshot is returned when multiple compute snapshot match.
     * 
     * ***
     * 
     */
    @Import(name="mostRecent")
    private @Nullable Boolean mostRecent;

    /**
     * @return If `filter` is provided, ensures the most recent snapshot is returned when multiple compute snapshot match.
     * 
     * ***
     * 
     */
    public Optional<Boolean> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }

    /**
     * The name of the compute snapshot. One of `name` or `filter` must be provided.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the compute snapshot. One of `name` or `filter` must be provided.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable String project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    private GetSnapshotPlainArgs() {}

    private GetSnapshotPlainArgs(GetSnapshotPlainArgs $) {
        this.filter = $.filter;
        this.mostRecent = $.mostRecent;
        this.name = $.name;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSnapshotPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSnapshotPlainArgs $;

        public Builder() {
            $ = new GetSnapshotPlainArgs();
        }

        public Builder(GetSnapshotPlainArgs defaults) {
            $ = new GetSnapshotPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filter A filter to retrieve the compute snapshot.
         * See [API filter parameter documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots/list#body.QUERY_PARAMETERS.filter) for reference.
         * If multiple compute snapshot match, either adjust the filter or specify `most_recent`. One of `name` or `filter` must be provided.
         * If you want to use a regular expression, use the `eq` (equal) or `ne` (not equal) operator against a single un-parenthesized expression with or without quotes or against multiple parenthesized expressions. Example `sourceDisk eq &#39;.*(.*&#47;data-disk$).*&#39;`. More details for golang Snapshots list call filters [here](https://pkg.go.dev/google.golang.org/api/compute/v1#SnapshotsListCall.Filter).
         * 
         * @return builder
         * 
         */
        public Builder filter(@Nullable String filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param mostRecent If `filter` is provided, ensures the most recent snapshot is returned when multiple compute snapshot match.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(@Nullable Boolean mostRecent) {
            $.mostRecent = mostRecent;
            return this;
        }

        /**
         * @param name The name of the compute snapshot. One of `name` or `filter` must be provided.
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
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        public GetSnapshotPlainArgs build() {
            return $;
        }
    }

}
