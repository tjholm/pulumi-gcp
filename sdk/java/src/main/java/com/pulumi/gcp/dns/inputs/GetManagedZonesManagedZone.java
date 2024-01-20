// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetManagedZonesManagedZone extends com.pulumi.resources.InvokeArgs {

    public static final GetManagedZonesManagedZone Empty = new GetManagedZonesManagedZone();

    @Import(name="description", required=true)
    private String description;

    public String description() {
        return this.description;
    }

    @Import(name="dnsName", required=true)
    private String dnsName;

    public String dnsName() {
        return this.dnsName;
    }

    @Import(name="id", required=true)
    private String id;

    public String id() {
        return this.id;
    }

    @Import(name="managedZoneId", required=true)
    private Integer managedZoneId;

    public Integer managedZoneId() {
        return this.managedZoneId;
    }

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    @Import(name="nameServers", required=true)
    private List<String> nameServers;

    public List<String> nameServers() {
        return this.nameServers;
    }

    @Import(name="project", required=true)
    private String project;

    public String project() {
        return this.project;
    }

    @Import(name="visibility", required=true)
    private String visibility;

    public String visibility() {
        return this.visibility;
    }

    private GetManagedZonesManagedZone() {}

    private GetManagedZonesManagedZone(GetManagedZonesManagedZone $) {
        this.description = $.description;
        this.dnsName = $.dnsName;
        this.id = $.id;
        this.managedZoneId = $.managedZoneId;
        this.name = $.name;
        this.nameServers = $.nameServers;
        this.project = $.project;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetManagedZonesManagedZone defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetManagedZonesManagedZone $;

        public Builder() {
            $ = new GetManagedZonesManagedZone();
        }

        public Builder(GetManagedZonesManagedZone defaults) {
            $ = new GetManagedZonesManagedZone(Objects.requireNonNull(defaults));
        }

        public Builder description(String description) {
            $.description = description;
            return this;
        }

        public Builder dnsName(String dnsName) {
            $.dnsName = dnsName;
            return this;
        }

        public Builder id(String id) {
            $.id = id;
            return this;
        }

        public Builder managedZoneId(Integer managedZoneId) {
            $.managedZoneId = managedZoneId;
            return this;
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public Builder nameServers(List<String> nameServers) {
            $.nameServers = nameServers;
            return this;
        }

        public Builder nameServers(String... nameServers) {
            return nameServers(List.of(nameServers));
        }

        public Builder project(String project) {
            $.project = project;
            return this;
        }

        public Builder visibility(String visibility) {
            $.visibility = visibility;
            return this;
        }

        public GetManagedZonesManagedZone build() {
            if ($.description == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "description");
            }
            if ($.dnsName == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "dnsName");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "id");
            }
            if ($.managedZoneId == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "managedZoneId");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "name");
            }
            if ($.nameServers == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "nameServers");
            }
            if ($.project == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "project");
            }
            if ($.visibility == null) {
                throw new MissingRequiredPropertyException("GetManagedZonesManagedZone", "visibility");
            }
            return $;
        }
    }

}
