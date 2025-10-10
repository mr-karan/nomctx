{ config, lib, pkgs, ... }:

with lib;

let
  cfg = config.programs.nomctx;

  clusterType = types.submodule {
    options = {
      address = mkOption {
        type = types.str;
        description = "Nomad cluster address";
        example = "http://127.0.0.1:4646";
      };

      namespace = mkOption {
        type = types.nullOr types.str;
        default = null;
        description = "Default namespace for the cluster";
        example = "default";
      };

      token = mkOption {
        type = types.nullOr types.str;
        default = null;
        description = "Access token for the cluster";
      };

      httpAuth = mkOption {
        type = types.nullOr types.str;
        default = null;
        description = "HTTP basic auth credentials";
        example = "user:pass";
      };

      region = mkOption {
        type = types.nullOr types.str;
        default = null;
        description = "Nomad region";
      };

      auth = mkOption {
        type = types.nullOr (types.submodule {
          options = {
            method = mkOption {
              type = types.str;
              description = "Authentication method";
              example = "gitlab";
            };
            provider = mkOption {
              type = types.str;
              description = "Authentication provider";
              example = "nomad";
            };
          };
        });
        default = null;
        description = "SSO authentication configuration";
      };
    };
  };

  generateClusterConfig = name: cluster: ''
    cluster "${name}" {
      address = "${cluster.address}"
      ${optionalString (cluster.namespace != null) ''namespace = "${cluster.namespace}"''}
      ${optionalString (cluster.token != null) ''token = "${cluster.token}"''}
      ${optionalString (cluster.httpAuth != null) ''http_auth = "${cluster.httpAuth}"''}
      ${optionalString (cluster.region != null) ''region = "${cluster.region}"''}
      ${optionalString (cluster.auth != null) ''
      auth {
        method = "${cluster.auth.method}"
        provider = "${cluster.auth.provider}"
      }
      ''}
    }
  '';

  configFile = pkgs.writeText "nomctx-config.hcl" ''
    ${concatStringsSep "\n" (mapAttrsToList generateClusterConfig cfg.clusters)}
  '';

in {
  options.programs.nomctx = {
    enable = mkEnableOption "nomctx, a tool for switching between Nomad clusters";

    package = mkOption {
      type = types.package;
      default = pkgs.nomctx or (pkgs.callPackage ./default.nix {});
      defaultText = literalExpression "pkgs.nomctx";
      description = "The nomctx package to use";
    };

    clusters = mkOption {
      type = types.attrsOf clusterType;
      default = {};
      description = "Nomad cluster configurations";
      example = literalExpression ''
        {
          local = {
            address = "http://127.0.0.1:4646";
            namespace = "default";
          };
          prod = {
            address = "http://10.0.0.3:4646";
            namespace = "blue";
            token = "secret-token";
          };
        }
      '';
    };
  };

  config = mkIf cfg.enable {
    environment.systemPackages = [ cfg.package ];

    environment.etc."nomctx/config.hcl" = mkIf (cfg.clusters != {}) {
      source = configFile;
    };

    # Optionally create a symlink in user's config directory
    system.activationScripts.nomctx = mkIf (cfg.clusters != {}) ''
      mkdir -p /root/.config/nomctx
      ln -sf /etc/nomctx/config.hcl /root/.config/nomctx/config.hcl
    '';
  };
}
