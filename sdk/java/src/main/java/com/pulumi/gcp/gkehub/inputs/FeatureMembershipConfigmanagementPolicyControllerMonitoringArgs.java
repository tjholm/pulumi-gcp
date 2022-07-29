// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs Empty = new FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs();

    @Import(name="backends")
    private @Nullable Output<List<String>> backends;

    public Optional<Output<List<String>>> backends() {
        return Optional.ofNullable(this.backends);
    }

    private FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs() {}

    private FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs(FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs $) {
        this.backends = $.backends;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs $;

        public Builder() {
            $ = new FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs();
        }

        public Builder(FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs defaults) {
            $ = new FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs(Objects.requireNonNull(defaults));
        }

        public Builder backends(@Nullable Output<List<String>> backends) {
            $.backends = backends;
            return this;
        }

        public Builder backends(List<String> backends) {
            return backends(Output.of(backends));
        }

        public Builder backends(String... backends) {
            return backends(List.of(backends));
        }

        public FeatureMembershipConfigmanagementPolicyControllerMonitoringArgs build() {
            return $;
        }
    }

}
