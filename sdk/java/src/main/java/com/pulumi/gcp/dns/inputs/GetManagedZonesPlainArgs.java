// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dns.inputs.GetManagedZonesManagedZone;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetManagedZonesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetManagedZonesPlainArgs Empty = new GetManagedZonesPlainArgs();

    @Import(name="managedZones")
    private @Nullable List<GetManagedZonesManagedZone> managedZones;

    public Optional<List<GetManagedZonesManagedZone>> managedZones() {
        return Optional.ofNullable(this.managedZones);
    }

    @Import(name="project")
    private @Nullable String project;

    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    private GetManagedZonesPlainArgs() {}

    private GetManagedZonesPlainArgs(GetManagedZonesPlainArgs $) {
        this.managedZones = $.managedZones;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetManagedZonesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetManagedZonesPlainArgs $;

        public Builder() {
            $ = new GetManagedZonesPlainArgs();
        }

        public Builder(GetManagedZonesPlainArgs defaults) {
            $ = new GetManagedZonesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder managedZones(@Nullable List<GetManagedZonesManagedZone> managedZones) {
            $.managedZones = managedZones;
            return this;
        }

        public Builder managedZones(GetManagedZonesManagedZone... managedZones) {
            return managedZones(List.of(managedZones));
        }

        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        public GetManagedZonesPlainArgs build() {
            return $;
        }
    }

}
