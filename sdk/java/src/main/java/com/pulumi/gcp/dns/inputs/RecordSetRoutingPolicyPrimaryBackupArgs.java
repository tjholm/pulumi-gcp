// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dns.inputs.RecordSetRoutingPolicyPrimaryBackupBackupGeoArgs;
import com.pulumi.gcp.dns.inputs.RecordSetRoutingPolicyPrimaryBackupPrimaryArgs;
import java.lang.Boolean;
import java.lang.Double;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RecordSetRoutingPolicyPrimaryBackupArgs extends com.pulumi.resources.ResourceArgs {

    public static final RecordSetRoutingPolicyPrimaryBackupArgs Empty = new RecordSetRoutingPolicyPrimaryBackupArgs();

    /**
     * The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
     * Structure is document above.
     * 
     */
    @Import(name="backupGeos", required=true)
    private Output<List<RecordSetRoutingPolicyPrimaryBackupBackupGeoArgs>> backupGeos;

    /**
     * @return The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
     * Structure is document above.
     * 
     */
    public Output<List<RecordSetRoutingPolicyPrimaryBackupBackupGeoArgs>> backupGeos() {
        return this.backupGeos;
    }

    /**
     * Specifies whether to enable fencing for backup geo queries.
     * 
     */
    @Import(name="enableGeoFencingForBackups")
    private @Nullable Output<Boolean> enableGeoFencingForBackups;

    /**
     * @return Specifies whether to enable fencing for backup geo queries.
     * 
     */
    public Optional<Output<Boolean>> enableGeoFencingForBackups() {
        return Optional.ofNullable(this.enableGeoFencingForBackups);
    }

    /**
     * The list of global primary targets to be health checked.
     * Structure is document below.
     * 
     */
    @Import(name="primary", required=true)
    private Output<RecordSetRoutingPolicyPrimaryBackupPrimaryArgs> primary;

    /**
     * @return The list of global primary targets to be health checked.
     * Structure is document below.
     * 
     */
    public Output<RecordSetRoutingPolicyPrimaryBackupPrimaryArgs> primary() {
        return this.primary;
    }

    /**
     * Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
     * 
     */
    @Import(name="trickleRatio")
    private @Nullable Output<Double> trickleRatio;

    /**
     * @return Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
     * 
     */
    public Optional<Output<Double>> trickleRatio() {
        return Optional.ofNullable(this.trickleRatio);
    }

    private RecordSetRoutingPolicyPrimaryBackupArgs() {}

    private RecordSetRoutingPolicyPrimaryBackupArgs(RecordSetRoutingPolicyPrimaryBackupArgs $) {
        this.backupGeos = $.backupGeos;
        this.enableGeoFencingForBackups = $.enableGeoFencingForBackups;
        this.primary = $.primary;
        this.trickleRatio = $.trickleRatio;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RecordSetRoutingPolicyPrimaryBackupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RecordSetRoutingPolicyPrimaryBackupArgs $;

        public Builder() {
            $ = new RecordSetRoutingPolicyPrimaryBackupArgs();
        }

        public Builder(RecordSetRoutingPolicyPrimaryBackupArgs defaults) {
            $ = new RecordSetRoutingPolicyPrimaryBackupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupGeos The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
         * Structure is document above.
         * 
         * @return builder
         * 
         */
        public Builder backupGeos(Output<List<RecordSetRoutingPolicyPrimaryBackupBackupGeoArgs>> backupGeos) {
            $.backupGeos = backupGeos;
            return this;
        }

        /**
         * @param backupGeos The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
         * Structure is document above.
         * 
         * @return builder
         * 
         */
        public Builder backupGeos(List<RecordSetRoutingPolicyPrimaryBackupBackupGeoArgs> backupGeos) {
            return backupGeos(Output.of(backupGeos));
        }

        /**
         * @param backupGeos The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
         * Structure is document above.
         * 
         * @return builder
         * 
         */
        public Builder backupGeos(RecordSetRoutingPolicyPrimaryBackupBackupGeoArgs... backupGeos) {
            return backupGeos(List.of(backupGeos));
        }

        /**
         * @param enableGeoFencingForBackups Specifies whether to enable fencing for backup geo queries.
         * 
         * @return builder
         * 
         */
        public Builder enableGeoFencingForBackups(@Nullable Output<Boolean> enableGeoFencingForBackups) {
            $.enableGeoFencingForBackups = enableGeoFencingForBackups;
            return this;
        }

        /**
         * @param enableGeoFencingForBackups Specifies whether to enable fencing for backup geo queries.
         * 
         * @return builder
         * 
         */
        public Builder enableGeoFencingForBackups(Boolean enableGeoFencingForBackups) {
            return enableGeoFencingForBackups(Output.of(enableGeoFencingForBackups));
        }

        /**
         * @param primary The list of global primary targets to be health checked.
         * Structure is document below.
         * 
         * @return builder
         * 
         */
        public Builder primary(Output<RecordSetRoutingPolicyPrimaryBackupPrimaryArgs> primary) {
            $.primary = primary;
            return this;
        }

        /**
         * @param primary The list of global primary targets to be health checked.
         * Structure is document below.
         * 
         * @return builder
         * 
         */
        public Builder primary(RecordSetRoutingPolicyPrimaryBackupPrimaryArgs primary) {
            return primary(Output.of(primary));
        }

        /**
         * @param trickleRatio Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
         * 
         * @return builder
         * 
         */
        public Builder trickleRatio(@Nullable Output<Double> trickleRatio) {
            $.trickleRatio = trickleRatio;
            return this;
        }

        /**
         * @param trickleRatio Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
         * 
         * @return builder
         * 
         */
        public Builder trickleRatio(Double trickleRatio) {
            return trickleRatio(Output.of(trickleRatio));
        }

        public RecordSetRoutingPolicyPrimaryBackupArgs build() {
            $.backupGeos = Objects.requireNonNull($.backupGeos, "expected parameter 'backupGeos' to be non-null");
            $.primary = Objects.requireNonNull($.primary, "expected parameter 'primary' to be non-null");
            return $;
        }
    }

}
