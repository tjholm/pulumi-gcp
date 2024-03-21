// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TunnelDestGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final TunnelDestGroupArgs Empty = new TunnelDestGroupArgs();

    /**
     * List of CIDRs that this group applies to.
     * 
     */
    @Import(name="cidrs")
    private @Nullable Output<List<String>> cidrs;

    /**
     * @return List of CIDRs that this group applies to.
     * 
     */
    public Optional<Output<List<String>>> cidrs() {
        return Optional.ofNullable(this.cidrs);
    }

    /**
     * List of FQDNs that this group applies to.
     * 
     */
    @Import(name="fqdns")
    private @Nullable Output<List<String>> fqdns;

    /**
     * @return List of FQDNs that this group applies to.
     * 
     */
    public Optional<Output<List<String>>> fqdns() {
        return Optional.ofNullable(this.fqdns);
    }

    /**
     * Unique tunnel destination group name.
     * 
     * ***
     * 
     */
    @Import(name="groupName", required=true)
    private Output<String> groupName;

    /**
     * @return Unique tunnel destination group name.
     * 
     * ***
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
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

    /**
     * The region of the tunnel group. Must be the same as the network resources in the group.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region of the tunnel group. Must be the same as the network resources in the group.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private TunnelDestGroupArgs() {}

    private TunnelDestGroupArgs(TunnelDestGroupArgs $) {
        this.cidrs = $.cidrs;
        this.fqdns = $.fqdns;
        this.groupName = $.groupName;
        this.project = $.project;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TunnelDestGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TunnelDestGroupArgs $;

        public Builder() {
            $ = new TunnelDestGroupArgs();
        }

        public Builder(TunnelDestGroupArgs defaults) {
            $ = new TunnelDestGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrs List of CIDRs that this group applies to.
         * 
         * @return builder
         * 
         */
        public Builder cidrs(@Nullable Output<List<String>> cidrs) {
            $.cidrs = cidrs;
            return this;
        }

        /**
         * @param cidrs List of CIDRs that this group applies to.
         * 
         * @return builder
         * 
         */
        public Builder cidrs(List<String> cidrs) {
            return cidrs(Output.of(cidrs));
        }

        /**
         * @param cidrs List of CIDRs that this group applies to.
         * 
         * @return builder
         * 
         */
        public Builder cidrs(String... cidrs) {
            return cidrs(List.of(cidrs));
        }

        /**
         * @param fqdns List of FQDNs that this group applies to.
         * 
         * @return builder
         * 
         */
        public Builder fqdns(@Nullable Output<List<String>> fqdns) {
            $.fqdns = fqdns;
            return this;
        }

        /**
         * @param fqdns List of FQDNs that this group applies to.
         * 
         * @return builder
         * 
         */
        public Builder fqdns(List<String> fqdns) {
            return fqdns(Output.of(fqdns));
        }

        /**
         * @param fqdns List of FQDNs that this group applies to.
         * 
         * @return builder
         * 
         */
        public Builder fqdns(String... fqdns) {
            return fqdns(List.of(fqdns));
        }

        /**
         * @param groupName Unique tunnel destination group name.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName Unique tunnel destination group name.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
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

        /**
         * @param region The region of the tunnel group. Must be the same as the network resources in the group.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region of the tunnel group. Must be the same as the network resources in the group.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public TunnelDestGroupArgs build() {
            if ($.groupName == null) {
                throw new MissingRequiredPropertyException("TunnelDestGroupArgs", "groupName");
            }
            return $;
        }
    }

}
